package chester

// PestoState is the incremental state of the PeSTO evaluation: the
// piece-square sums for each side and the running game phase.
//
// It belongs to the evaluation, not to [Position]. The search carries one
// alongside each position and derives the child's state from the parent's,
// which is what keeps the board representation free of anything evaluation
// specific -- and keeps a future NNUE accumulator, which is far larger, out
// of the struct that gets copied at every node.
type PestoState struct {
	mg    [2]int32
	eg    [2]int32
	phase int32
}

// NewPestoState computes the state of a position from scratch. The search
// does this once at the root; every node below it updates instead.
func NewPestoState(p *Position) PestoState {
	var s PestoState

	for color := White; color <= Black; color++ {
		own := p.allPieces[color]

		for piece := Pawn; piece <= King; piece++ {
			var sq Square

			for bb := p.pieces[piece] & own; bb != 0; {
				sq, bb = bb.PopLSB()

				s.mg[color] += int32(mgTable[color][piece][sq])
				s.eg[color] += int32(egTable[color][piece][sq])
				s.phase += int32(gamephaseInc[piece])
			}
		}
	}

	return s
}

// Updated returns the state for a position reached by playing one move, given
// the position before the move and the position after it.
//
// It works out what changed by comparing the two boards rather than by being
// told what the move did. Exclusive-or of the per-colour occupancies gives
// exactly the squares the move touched, and the two mailboxes say what stood
// on each of them before and after.
//
// Deriving it this way means there is no move classification to get wrong.
// En passant needs no special case: the captured pawn's square is simply
// another changed square. Neither does promotion, because the piece type is
// read per square rather than assumed to be preserved. Castling is two
// ordinary pairs. The cases that would otherwise have to be duplicated from
// Position.Do, and could drift out of step with it, do not exist here.
func (s PestoState) Updated(before, after *Position) PestoState {
	// The exclusive-or has to be taken per colour and then combined. A
	// capture leaves its destination occupied before and after, so the
	// combined occupancy does not change there and the square would be
	// missed.
	changed := (before.allPieces[White] ^ after.allPieces[White]) |
		(before.allPieces[Black] ^ after.allPieces[Black])

	var sq Square
	for changed != 0 {
		sq, changed = changed.PopLSB()

		if piece := before.mailbox[sq]; piece != Empty {
			color := Color(before.allPieces[Black] >> sq & 1)
			s.mg[color] -= int32(mgTable[color][piece][sq])
			s.eg[color] -= int32(egTable[color][piece][sq])
			s.phase -= int32(gamephaseInc[piece])
		}

		if piece := after.mailbox[sq]; piece != Empty {
			color := Color(after.allPieces[Black] >> sq & 1)
			s.mg[color] += int32(mgTable[color][piece][sq])
			s.eg[color] += int32(egTable[color][piece][sq])
			s.phase += int32(gamephaseInc[piece])
		}
	}

	return s
}

// Score returns the evaluation in centipawns from the perspective of active.
func (s PestoState) Score(active, inactive Color) int {
	mgScore := int(s.mg[active] - s.mg[inactive])
	egScore := int(s.eg[active] - s.eg[inactive])

	// The phase interpolates between the middlegame and endgame tables. It
	// saturates because promotions can put more material on the board than
	// the opening started with.
	mgPhase := int(s.phase)
	if mgPhase > 24 {
		mgPhase = 24
	}
	egPhase := 24 - mgPhase

	return (mgScore*mgPhase + egScore*egPhase) / 24
}
