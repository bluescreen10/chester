package chester

// Move ordering is what makes alpha-beta pruning effective. Searching the
// best move first lets every remaining move at that node be refuted by a
// null-window-sized amount of work, which is the difference between an
// effective branching factor of b^(3/4) and b^(1/2). Everything here exists
// to get a likely-best move to the front of the list as cheaply as possible.
//
// Moves are ranked into disjoint score bands, best first:
//
//	transposition table move
//	queen promotions        (ordered by what they capture)
//	captures                (ordered by MVV-LVA)
//	knight promotions       (not dominated by the queen: =N+ forks)
//	killer moves            (quiet moves that refuted a sibling at this ply)
//	quiet moves             (ordered by the history heuristic)
//	rook and bishop promotions (only ever played to avoid stalemate)

const (
	// maxPly is the deepest ply the main search will ever reach, and the
	// size of every per-ply table. Search depth is clamped to stay below it.
	maxPly = 128

	// maxMoves is an upper bound on the number of legal moves in any chess
	// position (the true maximum is 218). Score buffers are sized to it so
	// they can live on the stack.
	maxMoves = 256
)

// Score bands for move ordering. The gaps between bands are wide enough that
// no within-band bonus can promote a move into the band above it.
const (
	scoreTTMove      int32 = 1 << 28
	scoreQueenPromo  int32 = 1 << 26
	scoreCapture     int32 = 1 << 24
	scoreKnightPromo int32 = 1 << 23
	scoreKiller1     int32 = 1 << 22
	scoreKiller2     int32 = scoreKiller1 - 1
	scoreUnderPromo  int32 = -(1 << 22)

	// maxHistory bounds the magnitude of a history score. Quiet moves score
	// in [-maxHistory, maxHistory], well inside the gap between the killer
	// and under-promotion bands.
	maxHistory int32 = 1 << 14
)

// pieceOrder ranks piece types by value for MVV-LVA ordering. It is indexed
// by [Piece], including King and Empty, so that a raw mailbox lookup never
// needs a bounds check or a branch. A king can never be captured and an
// empty square is not a capture, so both rank zero.
var pieceOrder = [7]int32{
	Pawn:   1,
	Knight: 3,
	Bishop: 3,
	Rook:   5,
	Queen:  9,
	King:   0,
	Empty:  0,
}

// promoBonus ranks a promotion by how useful the new piece actually is, on the
// same scale mvvLva uses for captured material.
//
// It is deliberately not the nominal value of the piece. A queen is a rook
// plus a bishop, so a rook or bishop promotion is never better than a queen
// except when the queen would stalemate -- which makes them worth less in
// practice than a knight promotion, even though a rook is nominally worth
// more than a knight. Ranking them by material would try the two useless
// promotions before the useful one.
var promoBonus = [7]int32{
	Queen:  9 * 16,
	Knight: 3 * 16,
	Rook:   1 * 16,
	Bishop: 1 * 16,
}

// mvvLva scores a capture by Most Valuable Victim, Least Valuable Attacker:
// capturing a queen with a pawn is tried before capturing a pawn with a
// queen. The victim dominates because winning material matters far more than
// which piece does the winning.
func mvvLva(victim, attacker Piece) int32 {
	return pieceOrder[victim]*16 - pieceOrder[attacker]
}

// victimOf returns the piece that move m captures in position p, or Empty if
// m is not a capture. En passant is handled separately because the captured
// pawn does not sit on the destination square.
func victimOf(p *Position, m Move) Piece {
	to := m.To()
	if victim := p.mailbox[to]; victim != Empty {
		return victim
	}
	if p.mailbox[m.From()] == Pawn && to == p.enPassantTarget {
		return Pawn
	}
	return Empty
}

// isQuiet reports whether m neither captures nor promotes. Only quiet moves
// are recorded as killers or scored by the history heuristic: a capture is
// already ordered by the material it wins, and rewarding it would drown out
// the signal for the moves that have nothing else to rank them by.
func isQuiet(p *Position, m Move) bool {
	return !m.IsPromotion() && victimOf(p, m) == Empty
}

// scoreMoves fills scores[i] with an ordering score for moves[i]. ttMove is
// the best move previously stored for this position, or 0 if there is none.
func (ctx *searchCtx) scoreMoves(scores *[maxMoves]int32, moves []Move, p *Position, ttMove Move, ply int) {
	var killer1, killer2 Move
	if ply < maxPly {
		killer1 = ctx.killers[ply][0]
		killer2 = ctx.killers[ply][1]
	}

	us := p.active
	history := &ctx.history[us]

	for i, m := range moves {
		var score int32

		switch victim := victimOf(p, m); {
		case m == ttMove:
			score = scoreTTMove

		// A promotion is ranked by the piece that appears, not by whether it
		// happens to capture on the way. Capturing with a rook promotion is
		// still a rook promotion, and belongs last.
		case m.IsPromotion():
			switch m.PromoPiece() {
			case Queen:
				score = scoreQueenPromo + mvvLva(victim, Pawn)
			case Knight:
				score = scoreKnightPromo + mvvLva(victim, Pawn)
			default:
				score = scoreUnderPromo
			}

		case victim != Empty:
			score = scoreCapture + mvvLva(victim, p.mailbox[m.From()])
		case m == killer1:
			score = scoreKiller1
		case m == killer2:
			score = scoreKiller2
		default:
			score = history[m.From()][m.To()]
		}

		scores[i] = score
	}
}

// scoreCaptures is the quiescence-search counterpart of scoreMoves. Every
// move it sees is a capture or a promotion, so it skips the killer and
// history lookups entirely.
func scoreCaptures(scores *[maxMoves]int32, moves []Move, p *Position) {
	for i, m := range moves {
		score := mvvLva(victimOf(p, m), p.mailbox[m.From()])
		if m.IsPromotion() {
			score += promoBonus[m.PromoPiece()]
		}
		scores[i] = score
	}
}

// pickNextMove swaps the highest scoring move in moves[i:] into position i.
//
// This is a selection sort performed one step at a time rather than a full
// sort up front, because most nodes produce a beta cutoff within the first
// few moves and the remaining moves are never looked at.
func pickNextMove(moves []Move, scores *[maxMoves]int32, i int) {
	best := i
	for j := i + 1; j < len(moves); j++ {
		if scores[j] > scores[best] {
			best = j
		}
	}
	if best != i {
		moves[i], moves[best] = moves[best], moves[i]
		scores[i], scores[best] = scores[best], scores[i]
	}
}

// storeKiller records m as the most recent quiet move to cause a beta cutoff
// at the given ply, keeping the previous killer as a second candidate.
func (ctx *searchCtx) storeKiller(ply int, m Move) {
	if ply >= maxPly || ctx.killers[ply][0] == m {
		return
	}
	ctx.killers[ply][1] = ctx.killers[ply][0]
	ctx.killers[ply][0] = m
}

// updateHistory adjusts the history score of a quiet move by bonus.
//
// The update uses "history gravity": the applied change shrinks as the score
// approaches maxHistory, so scores converge instead of overflowing and old
// entries decay on their own as new evidence arrives. No periodic rescaling
// of the table is needed.
func (ctx *searchCtx) updateHistory(us Color, m Move, bonus int32) {
	if bonus > maxHistory {
		bonus = maxHistory
	} else if bonus < -maxHistory {
		bonus = -maxHistory
	}

	magnitude := bonus
	if magnitude < 0 {
		magnitude = -magnitude
	}

	h := &ctx.history[us][m.From()][m.To()]
	*h += bonus - *h*magnitude/maxHistory
}

// updateQuietHeuristics is called when quiet move m at index i causes a beta
// cutoff. It promotes m and penalises the quiet moves that were searched
// before it and failed, so that the next visit to a similar position tries
// them in a better order.
func (ctx *searchCtx) updateQuietHeuristics(p *Position, moves []Move, i, depth, ply int) {
	m := moves[i]
	ctx.storeKiller(ply, m)

	bonus := int32(depth) * int32(depth)
	if bonus > maxHistory {
		bonus = maxHistory
	}

	us := p.active
	ctx.updateHistory(us, m, bonus)

	for _, tried := range moves[:i] {
		if isQuiet(p, tried) {
			ctx.updateHistory(us, tried, -bonus)
		}
	}
}
