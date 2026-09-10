package chester

import (
	"context"
	"errors"
	"math"
	"math/rand/v2"
	"time"
)

const (
	// Inf is a value representing positive infinity for search scores.
	Inf = 1_000_000_000
	// MateScore is the base score for a checkmate. The actual score is
	// adjusted by ply to favor shorter mates.
	MateScore = 1_000_000
)

// Evaluation holds the result of a search at a given depth.
type Evaluation struct {
	// Search depth reached.
	Depth int

	// Best move in pure algebraic coordinate notation (e.g. "e2e4").
	Best Move

	// Centipawn score from the side to move perspective
	Score int

	// Nodes is the total number of positions visited by the search so far,
	// including quiescence nodes.
	Nodes int64

	// QNodes is the subset of Nodes that were visited during quiescence.
	QNodes int64
}

// EvalFunc defines the signature for a function that performs a static
// evaluation of a [Position].
//
// It returns a score in centipawns (100 units = 1 pawn) from the perspective
// of the side to move. A positive value indicates an advantage for the
// active player, while a negative value indicates a disadvantage.
type EvalFunc func(*Position) int

// SearchOptions defines the constraints and heuristics used by the search engine
// to determine the best move. It allows for limiting the search by time,
// node count, or recursion depth.
type SearchOptions struct {
	// MaxTime is the maximum duration the search is allowed to run.
	// If the timer expires, the search returns the best move found
	// from the last fully completed depth.
	MaxTime time.Duration

	// MaxNodes is the maximum number of positions (nodes) the engine
	// will visit before aborting the search.
	MaxNodes int64

	// MaxDepth is the maximum number of plies (half-moves) to search.
	MaxDepth int

	// Moves is an optional list of specific moves to search. If empty,
	// the engine considers all legal moves in the position.
	Moves []Move

	// EvalFunc is the evaluation function used to score leaf nodes
	// in the search tree.
	EvalFunc EvalFunc

	// Optionally you can pass a transposition table to be used
	TranspositionTable *TranspositionTable

	// History holds the Zobrist hashes of the positions that preceded the
	// one being searched, oldest first and excluding it. Supplying it lets
	// the search see a repetition of a position that was actually reached
	// earlier in the game, not just one created inside the search tree.
	History []uint64

	// DisableBook suppresses the built-in opening book, so that every move
	// comes from the search.
	//
	// Self-play testing needs this. With the book on, both sides answer the
	// opening from the same table and the games measure the book rather than
	// the change being tested.
	DisableBook bool
}

var (
	// errMaxNodesReached is returned when the search is aborted because the
	// total number of visited nodes exceeds [SearchOptions.MaxNodes].
	errMaxNodesReached = errors.New("max nodes reached")

	// errContextCancelled is returned when the search is aborted due to
	// a timeout or a manual cancellation of the [context.Context].
	errContextCancelled = errors.New("context cancelled")
)

// defaultSearchOptions provides a sensible baseline for the engine.
var defaultSearchOptions = &SearchOptions{
	MaxTime:  time.Duration(20 * time.Second),
	MaxNodes: math.MaxInt64,
	MaxDepth: 7,
	EvalFunc: EvalPesto,
}

// searchCtx tracks the state and constraints of a single search execution.
// It embeds [context.Context] for cancellation signaling and maintains
// counters for performance monitoring.
type searchCtx struct {

	// Context is used to signal search abortion (timeout or manual).
	context.Context

	// tranposition table
	tt *TranspositionTable

	// maxNodes is the hard limit for total nodes allowed for this search.
	maxNodes int64

	// nodes tracks the total number of positions visited during
	// the main negamax search.
	nodes int64

	// qnodes tracks the number of positions visited specifically
	// during the quiescence search.
	qnodes int64

	// killers holds, per ply, up to two quiet moves that caused a beta
	// cutoff at that ply. A move that refutes one line often refutes its
	// siblings, so they are tried immediately after the captures.
	killers [maxPly][2]Move

	// history scores quiet moves by how often they have caused a beta
	// cutoff anywhere in the tree, indexed by [color][from][to]. It is the
	// only ordering signal available for quiet moves that are neither the
	// transposition table move nor a killer.
	history [Color(2)][64][64]int32

	// stack holds the Zobrist hash of every position preceding the one
	// being searched: the game history supplied by the caller, followed by
	// the positions along the current search path. See isRepetition.
	stack []uint64
}

// SearchBestMove initiates an asynchronous search for the best move.
// Returns a channel for evaluations and a function to cancel the search.
func SearchBestMove(p *Position, opts *SearchOptions) (chan Evaluation, context.CancelFunc) {
	if opts == nil {
		opts = defaultSearchOptions
	}

	ch := make(chan Evaluation)
	ctx, cancel := context.WithCancel(context.Background())

	if opts.MaxTime != 0 {
		ctx, cancel = context.WithTimeout(ctx, opts.MaxTime)
	}

	if opts.MaxNodes == 0 {
		opts.MaxNodes = math.MaxInt
	}

	go func() {
		defer close(ch)

		if !opts.DisableBook {
			if entries, ok := book[p.hash]; ok {
				move := pickMove(entries, rand.IntN)
				ch <- Evaluation{
					Depth: 1,
					Best:  move,
				}
				return
			}
		}

		// The move buffer is shared by every ply: each node carves its own
		// slice off the tail of its parent's. Sizing it for maxPly plies of
		// maxMoves keeps the search allocation free at any reachable depth.
		rootMoves := make([]Move, 0, maxPly*maxMoves)
		newPos := Position{}

		rootMoves, _ = LegalMoves(rootMoves, p)
		if len(opts.Moves) > 0 {
			rootMoves = filterMoves(rootMoves, opts.Moves)
		}

		count := len(rootMoves)
		if count == 0 {
			return
		}

		maxDepth := opts.MaxDepth
		if maxDepth > maxPly-1 {
			maxDepth = maxPly - 1
		}

		searchCtx := &searchCtx{
			Context:  ctx,
			maxNodes: opts.MaxNodes,
			tt:       opts.TranspositionTable,

			// The root position itself terminates the repetition stack: it
			// precedes every position the search will visit.
			stack: append(append(make([]uint64, 0, len(opts.History)+maxPly+1), opts.History...), p.hash),
		}

		// Everything already in the table belongs to an earlier search from
		// here on, which is what lets stale entries be evicted ahead of
		// shallow ones computed for the position actually on the board.
		opts.TranspositionTable.NewSearch()

		// The root state is the only one computed from scratch. Every node
		// below it derives its own from its parent's.
		rootAcc := NewPestoState(p)

		// Order the root moves once before the first iteration. Later
		// iterations reuse the previous iteration's best move instead,
		// which is a far stronger signal than any static ordering.
		var rootScores [maxMoves]int32
		searchCtx.scoreMoves(&rootScores, rootMoves, p, Move(0), 0)
		for i := range rootMoves {
			pickNextMove(rootMoves, &rootScores, i)
		}

	loop:
		// iterative deepening
		for depth := 1; depth <= maxDepth; depth++ {
			bestMoveAtDepth := rootMoves[0]
			bestScoreAtDepth := -Inf
			alpha := -Inf
			beta := Inf

			// evaluate each root move
			for _, m := range rootMoves {

				newPos = *p
				newPos.Do(m)

				score, err := negamax(searchCtx, &newPos, rootAcc.Updated(p, &newPos), rootMoves[count:], -beta, -alpha, depth-1, 1)
				if err != nil {
					break loop
				}
				score = -score

				if score > bestScoreAtDepth {
					bestScoreAtDepth = score
					bestMoveAtDepth = m

					if score > alpha {
						alpha = score
					}
				}

			}

			// Search this iteration's best move first at the next depth.
			// This is the main reason iterative deepening pays for itself:
			// an immediate cutoff on move one shrinks the entire tree.
			moveToFront(rootMoves, bestMoveAtDepth)

			// inform the current evaluation
			ch <- Evaluation{
				Depth:  depth,
				Best:   bestMoveAtDepth,
				Score:  bestScoreAtDepth,
				Nodes:  searchCtx.nodes,
				QNodes: searchCtx.qnodes,
			}

			// check for context cancellation
			select {
			case <-ctx.Done():
				break loop
			default:
			}
		}
	}()

	return ch, cancel
}

// filterMoves returns a subset of allMoves that are also present in wantMoves.
// It preserves the order of moves as they appear in wantMoves, provided they
// are legal (exist in allMoves).
func filterMoves(allMoves []Move, wantMoves []Move) []Move {
	existing := make(map[Move]bool)

	for _, m := range allMoves {
		existing[m] = true
	}

	var j int
	for _, m := range wantMoves {
		if _, ok := existing[m]; ok {
			allMoves[j] = m
			j++
		}
	}

	return allMoves[:j]
}

// moveToFront rotates m to the front of moves, preserving the relative order
// of the moves it displaces. It is a no-op if m is not present.
func moveToFront(moves []Move, m Move) {
	for i, cur := range moves {
		if cur == m {
			copy(moves[1:i+1], moves[:i])
			moves[0] = m
			return
		}
	}
}

// negamax performs a recursive negamax search with Alpha-Beta pruning from
// position p. It returns the best score achievable and the corresponding
// move. alpha and beta are the current window bounds; depth is the remaining
// plies to search.
//
// Negamax assumes both players maximize their score, with evaluation
// always from the side-to-move perspective.
// Checkmate is detected when the side to move is in check with no legal
// moves; stalemate when there are no legal moves and the king is not in
// check. Both are handled before recursing so that eval is never called on
// a terminal position.
func negamax(ctx *searchCtx, p *Position, acc PestoState, moves []Move, alpha, beta, depth, ply int) (int, error) {

	// A position that already occurred on this path, or earlier in the game,
	// is a draw. This is checked before the transposition table because the
	// table is path independent and cannot know about repetitions.
	if ctx.isRepetition(p) {
		return drawScore, nil
	}

	// tranposition table enabled
	var ttMove Move
	if ctx.tt != nil {
		if entry, ok := ctx.tt.get(p.hash); ok {
			// The entry move is worth having even when the entry depth is
			// too shallow to cut off: ordering it first is where most of the
			// table's value comes from.
			ttMove = entry.move

			if int(entry.depth) >= depth {
				score := scoreFromTT(int(entry.score), ply)
				switch {
				case entry.flag == exact:
					return score, nil
				case entry.flag == lowerBound && score >= beta:
					return score, nil
				case entry.flag == upperBound && score <= alpha:
					return score, nil
				}
			}
		}
	}

	if depth == 0 {
		return quiescence(ctx, p, acc, moves, alpha, beta)
	}

	moves, inCheck := LegalMoves(moves, p)
	count := len(moves)

	if count == 0 {
		if inCheck {
			return -MateScore + ply, nil
		} else {
			return drawScore, nil
		}
	}

	// Fifty-move rule. Checked after move generation so that a checkmate
	// delivered on the hundredth half-move still takes precedence over it.
	if p.halfMoves >= 100 {
		return drawScore, nil
	}

	var scores [maxMoves]int32
	ctx.scoreMoves(&scores, moves, p, ttMove, ply)

	originalAlpha := alpha
	bestScore := -Inf
	bestMove := Move(0)

	// Make this position visible to the subtree below it as a repetition
	// candidate. Every path out of the loop pops it again.
	ctx.stack = append(ctx.stack, p.hash)

	var newPos Position

	for i := range moves {
		// Bring the best remaining move to the front. Doing this lazily
		// avoids sorting the tail of a list that a cutoff never reaches.
		pickNextMove(moves, &scores, i)
		m := moves[i]

		// abort if we exceed the number of nodes
		ctx.nodes++
		if ctx.nodes > ctx.maxNodes {
			ctx.pop()
			return 0, errMaxNodesReached
		}

		// abort if context has been cancelled
		if ctx.nodes%2048 == 0 {
			select {
			case <-ctx.Done():
				ctx.pop()
				return 0, errContextCancelled
			default:
			}
		}

		newPos = *p
		newPos.Do(m)
		score, err := negamax(ctx, &newPos, acc.Updated(p, &newPos), moves[count:], -beta, -alpha, depth-1, ply+1)

		if err != nil {
			ctx.pop()
			return 0, err
		}

		score = -score

		if score > bestScore {
			bestScore = score
			bestMove = m
		}

		if score > alpha {
			alpha = score
		}

		if alpha >= beta {
			// A quiet move that refutes this node is likely to refute its
			// siblings too, so remember it before bailing out.
			if isQuiet(p, m) {
				ctx.updateQuietHeuristics(p, moves, i, depth, ply)
			}
			break
		}
	}

	ctx.pop()

	// transposition table enabled
	if ctx.tt != nil {
		flag := exact
		if bestScore <= originalAlpha {
			flag = upperBound
		} else if bestScore >= beta {
			flag = lowerBound
		}

		// Which record to give up is the table's decision: it can see what is
		// in the bucket now, whereas the entry probed on the way in is a
		// snapshot from before this whole subtree ran.
		ctx.tt.set(p.hash, ttEntry{
			score: int32(scoreToTT(bestScore, ply)),
			move:  bestMove,
			depth: int8(depth),
			flag:  flag,
		})
	}
	return bestScore, nil
}

// pop removes the most recently pushed position from the repetition stack.
func (ctx *searchCtx) pop() {
	ctx.stack = ctx.stack[:len(ctx.stack)-1]
}

// quiescence performs a restricted search that only considers "noisy" moves
// (captures) until a "quiet" position is reached.
//
// This is critical for avoiding the "Horizon Effect," where the engine
// might misjudge a position because the main search depth ended
// right in the middle of a piece exchange.
//
// It returns a score that represents the settled value of the position.
// If the search is interrupted by a timeout or node limit, it returns
// an error to ensure the partial result is discarded.
func quiescence(ctx *searchCtx, p *Position, acc PestoState, moves []Move, alpha, beta int) (int, error) {
	score := acc.Score(p.active, p.inactive)

	if score >= beta {
		return beta, nil
	}
	if score > alpha {
		alpha = score
	}

	// Captures and promotions. Both change material sharply enough that
	// stopping on one would leave the score mid-swing, which is the horizon
	// effect quiescence exists to avoid.
	moves, _ = NoisyMoves(moves, p)
	count := len(moves)

	// Quiescence is where most of the nodes are spent, and it is almost
	// entirely captures. Trying the most valuable victim first means a
	// losing exchange sequence is usually refuted by its first move.
	var scores [maxMoves]int32
	scoreCaptures(&scores, moves, p)

	var newPos Position

	for i := range moves {
		pickNextMove(moves, &scores, i)
		m := moves[i]

		// abort if max nodes
		ctx.nodes++
		ctx.qnodes++
		if ctx.nodes > ctx.maxNodes {
			return 0, errMaxNodesReached
		}

		// abort if context has been cancelled
		if ctx.nodes%2048 == 0 {
			select {
			case <-ctx.Done():
				return 0, errContextCancelled
			default:
			}
		}

		newPos = *p
		newPos.Do(m)

		score, err := quiescence(ctx, &newPos, acc.Updated(p, &newPos), moves[count:], -beta, -alpha)

		if err != nil {
			return 0, err
		}

		score = -score

		if score >= beta {
			return beta, nil
		}
		if score > alpha {
			alpha = score
		}
	}

	return alpha, nil
}

// eval returns a static evaluation of position p in centipawns
// using material count only. Positive values favour Attacking,
// negative values favour Defending.
//
// Piece values:
//
//	Pawn=100  Knight=300  Bishop=300  Rook=500  Queen=900
func EvalMaterial(p *Position) int {
	pawns := p.Pawns().OnesCount()
	knight := p.Knights().OnesCount()
	bishop := p.Bishops().OnesCount()
	rook := p.Rooks().OnesCount()
	queen := p.Queens().OnesCount()

	ePawns := p.EnemyPawns().OnesCount()
	eKnight := p.EnemyKnights().OnesCount()
	eBishop := p.EnemyBishops().OnesCount()
	eRook := p.EnemyRooks().OnesCount()
	eQueen := p.EnemyQueens().OnesCount()

	return (pawns + knight*3 + bishop*3 + rook*5 + queen*9 -
		ePawns - eKnight*3 - eBishop*3 - eRook*5 - eQueen*9) * 100
}

// BookMove returns a move for position p from the built-in opening book and
// reports whether the position appears in it at all.
//
// intn supplies the randomness, as math/rand/v2's rand.IntN does; passing a
// seeded generator's IntN method makes a sequence of book moves reproducible.
func BookMove(p *Position, intn func(int) int) (Move, bool) {
	entries, ok := book[p.hash]
	if !ok {
		return Move(0), false
	}
	return pickMove(entries, intn), true
}

// pickMove selects a move from a set of book entries using weighted random
// selection. Entries with higher Weight are chosen proportionally more often.
// If all weights are zero, a move is chosen uniformly at random.
func pickMove(entries []bookEntry, intn func(int) int) Move {
	var total int
	for _, m := range entries {
		total += int(m.Weight)
	}

	if total == 0 {
		return entries[len(entries)-1].Move
	}

	r := intn(total)
	for _, e := range entries {
		r -= int(e.Weight)
		if r < 0 {
			return e.Move
		}
	}
	return entries[len(entries)-1].Move
}

var mgValue = [Piece(6)]int{82, 337, 365, 477, 1025, 0}
var egValue = [Piece(6)]int{94, 281, 297, 512, 936, 0}
var mgTable [Color(2)][Piece(6)][64]int
var egTable [Color(2)][Piece(6)][64]int

var mgPawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	98, 134, 61, 95, 68, 126, 34, -11,
	-6, 7, 26, 31, 65, 56, 25, -20,
	-14, 13, 6, 21, 23, 12, 17, -23,
	-27, -2, -5, 12, 17, 6, 10, -25,
	-26, -4, -4, -10, 3, 3, 33, -12,
	-35, -1, -20, -23, -15, 24, 38, -22,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var egPawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	178, 173, 158, 134, 147, 132, 165, 187,
	94, 100, 85, 67, 56, 53, 82, 84,
	32, 24, 13, 5, -2, 4, 17, 17,
	13, 9, -3, -7, -7, -8, 3, -1,
	4, 7, -6, 1, 0, -5, -1, -8,
	13, 8, 8, 10, 13, 0, 2, -7,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var mgKnightTable = [64]int{
	-167, -89, -34, -49, 61, -97, -15, -107,
	-73, -41, 72, 36, 23, 62, 7, -17,
	-47, 60, 37, 65, 84, 129, 73, 44,
	-9, 17, 19, 53, 37, 69, 18, 22,
	-13, 4, 16, 13, 28, 19, 21, -8,
	-23, -9, 12, 10, 19, 17, 25, -16,
	-29, -53, -12, -3, -1, 18, -14, -19,
	-105, -21, -58, -33, -17, -28, -19, -23,
}

var egKnightTable = [64]int{
	-58, -38, -13, -28, -31, -27, -63, -99,
	-25, -8, -25, -2, -9, -25, -24, -52,
	-24, -20, 10, 9, -1, -9, -19, -41,
	-17, 3, 22, 22, 22, 11, 8, -18,
	-18, -6, 16, 25, 16, 17, 4, -18,
	-23, -3, -1, 15, 10, -3, -20, -22,
	-42, -20, -10, -5, -2, -20, -23, -44,
	-29, -51, -23, -15, -22, -18, -50, -64,
}

var mgBishopTable = [64]int{
	-29, 4, -82, -37, -25, -42, 7, -8,
	-26, 16, -18, -13, 30, 59, 18, -47,
	-16, 37, 43, 40, 35, 50, 37, -2,
	-4, 5, 19, 50, 37, 37, 7, -2,
	-6, 13, 13, 26, 34, 12, 10, 4,
	0, 15, 15, 15, 14, 27, 18, 10,
	4, 15, 16, 0, 7, 21, 33, 1,
	-33, -3, -14, -21, -13, -12, -39, -21,
}

var egBishopTable = [64]int{
	-14, -21, -11, -8, -7, -9, -17, -24,
	-8, -4, 7, -12, -3, -13, -4, -14,
	2, -8, 0, -1, -2, 6, 0, 4,
	-3, 9, 12, 9, 14, 10, 3, 2,
	-6, 3, 13, 19, 7, 10, -3, -9,
	-12, -3, 8, 10, 13, 3, -7, -15,
	-14, -18, -7, -1, 4, -9, -15, -27,
	-23, -9, -23, -5, -9, -16, -5, -17,
}

var mgRookTable = [64]int{
	32, 42, 32, 51, 63, 9, 31, 43,
	27, 32, 58, 62, 80, 67, 26, 44,
	-5, 19, 26, 36, 17, 45, 61, 16,
	-24, -11, 7, 26, 24, 35, -8, -20,
	-36, -26, -12, -1, 9, -7, 6, -23,
	-45, -25, -16, -17, 3, 0, -5, -33,
	-44, -16, -20, -9, -1, 11, -6, -71,
	-19, -13, 1, 17, 16, 7, -37, -26,
}

var egRookTable = [64]int{
	13, 10, 18, 15, 12, 12, 8, 5,
	11, 13, 13, 11, -3, 3, 8, 3,
	7, 7, 7, 5, 4, -3, -5, -3,
	4, 3, 13, 1, 2, 1, -1, 2,
	3, 5, 8, 4, -5, -6, -8, -11,
	-4, 0, -5, -1, -7, -12, -8, -16,
	-6, -6, 0, 2, -9, -9, -11, -3,
	-9, 2, 3, -1, -5, -13, 4, -20,
}

var mgQueenTable = [64]int{
	-28, 0, 29, 12, 59, 44, 43, 45,
	-24, -39, -5, 1, -16, 57, 28, 54,
	-13, -17, 7, 8, 29, 56, 47, 57,
	-27, -27, -16, -16, -1, 17, -2, 1,
	-9, -26, -9, -10, -2, -4, 3, -3,
	-14, 2, -11, -2, -5, 2, 14, 5,
	-35, -8, 11, 2, 8, 15, -3, 1,
	-1, -18, -9, 10, -15, -25, -31, -50,
}

var egQueenTable = [64]int{
	-9, 22, 22, 27, 27, 19, 10, 20,
	-17, 20, 32, 41, 58, 25, 30, 0,
	-20, 6, 9, 49, 47, 35, 19, 9,
	3, 22, 24, 45, 57, 40, 57, 36,
	-18, 28, 19, 47, 31, 34, 39, 23,
	-16, -27, 15, 6, 9, 17, 10, 5,
	-22, -23, -30, -16, -16, -23, -36, -32,
	-33, -28, -22, -43, -5, -32, -20, -41,
}

var mgKingTable = [64]int{
	-65, 23, 16, -15, -56, -34, 2, 13,
	29, -1, -20, -7, -8, -4, -38, -29,
	-9, 24, 2, -16, -20, 6, 22, -22,
	-17, -20, -12, -27, -30, -25, -14, -36,
	-49, -1, -27, -39, -46, -44, -33, -51,
	-14, -14, -22, -46, -44, -30, -15, -27,
	1, 7, -8, -64, -43, -16, 9, 8,
	-15, 36, 12, -54, 8, -28, 24, 14,
}

var egKingTable = [64]int{
	-74, -35, -18, -18, -11, 15, 4, -17,
	-12, 17, 14, 17, 17, 38, 23, 11,
	10, 17, 23, 15, 20, 45, 44, 13,
	-8, 22, 24, 27, 26, 33, 26, 3,
	-18, -4, 21, 24, 27, 23, 9, -11,
	-19, -3, 11, 21, 23, 16, 7, -9,
	-27, -11, 4, 13, 14, 4, -5, -17,
	-53, -34, -21, -11, -28, -14, -24, -43,
}

var mgPestoTable = [6][64]int{
	mgPawnTable,
	mgKnightTable,
	mgBishopTable,
	mgRookTable,
	mgQueenTable,
	mgKingTable,
}

var egPestoTable = [6][64]int{
	egPawnTable,
	egKnightTable,
	egBishopTable,
	egRookTable,
	egQueenTable,
	egKingTable,
}

var gamephaseInc = [6]int{0, 1, 1, 2, 4, 0}

// FIXME: this should be moved to const.go and generated by
// the const generator.
func init() {
	for p := range Piece(6) {
		for sq := range Square(64) {
			mgTable[White][p][sq] = mgValue[p] + mgPestoTable[p][int(sq)]
			egTable[White][p][sq] = egValue[p] + egPestoTable[p][int(sq)]
			mgTable[Black][p][sq] = mgValue[p] + mgPestoTable[p][int(sq^56)]
			egTable[Black][p][sq] = egValue[p] + egPestoTable[p][int(sq^56)]
		}
	}
}

// EvalPesto calculates a static evaluation using the PeSTO method.
// Returns a score in centipawns based on PST and game phase.
//
// The board is walked one piece at a time through the piece bitboards rather
// than one square at a time. A position has at most thirty-two pieces and
// usually far fewer, so scanning all sixty-four squares spends most of its
// iterations establishing that a square is empty.
func EvalPesto(p *Position) int {
	var mg, eg [2]int
	gamePhase := 0

	for color := White; color <= Black; color++ {
		own := p.allPieces[color]

		for piece := Pawn; piece <= King; piece++ {
			var sq Square

			for bb := p.pieces[piece] & own; bb != 0; {
				sq, bb = bb.PopLSB()

				mg[color] += mgTable[color][piece][sq]
				eg[color] += egTable[color][piece][sq]
				gamePhase += gamephaseInc[piece]
			}
		}
	}

	mgScore := mg[p.active] - mg[p.inactive]
	egScore := eg[p.active] - eg[p.inactive]

	// The phase interpolates between the middlegame and endgame tables. It
	// saturates because promotions can put more material on the board than
	// the opening started with.
	mgPhase := gamePhase
	if mgPhase > 24 {
		mgPhase = 24
	}
	egPhase := 24 - mgPhase

	return (mgScore*mgPhase + egScore*egPhase) / 24
}
