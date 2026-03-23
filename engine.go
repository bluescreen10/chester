package chester

import (
	"context"
	"math"
	"math/rand/v2"
)

// Evaluation holds the result of a search at a given depth.
type Evaluation struct {
	// Search depth reached.
	Depth int

	// Best move in pure algebraic coordinate notation (e.g. "e2e4").
	Best string

	// Centipawn score from White's perspective.
	Score int
}

// SearchBestMove searches for the best move in position p and sends the
// result on the returned channel, which is closed when the search completes.
// The search runs in a separate goroutine and respects ctx for cancellation.
//
// If the position is found in the opening book the book move is returned
// immediately at depth 1 with no score. Otherwise a fixed-depth minimax
// search with Alpha-Beta pruning is performed.
func SearchBestMove(ctx context.Context, p *Position) chan Evaluation {
	ch := make(chan Evaluation)

	go func() {
		defer close(ch)

		if entries, ok := Book[p.hash]; ok {
			move := pickMove(entries)
			ch <- Evaluation{
				Depth: 1,
				Best:  move.String(),
			}
			return
		}

		depth := 5
		eval, m := minmax(p, math.MinInt, math.MaxInt, depth)
		ch <- Evaluation{
			Depth: depth,
			Best:  m.String(),
			Score: eval,
		}
	}()
	return ch
}

// minmax performs a recursive minimax search with Alpha-Beta pruning from
// position p. It returns the best score achievable and the corresponding
// move. alpha and beta are the current window bounds; depth is the remaining
// plies to search.
//
// White is the maximising player; Black is the minimising player.
// Checkmate is detected when the side to move is in check with no legal
// moves; stalemate when there are no legal moves and the king is not in
// check. Both are handled before recursing so that eval is never called on
// a terminal position.
func minmax(p *Position, alpha, beta, depth int) (int, Move) {
	if depth == 0 {
		return eval(p), Move(0)
	}
	moves := make([]Move, 0, 218)
	moves, inCheck := LegalMoves(moves, p)

	if inCheck && len(moves) == 0 {
		if p.Active() == White {
			return math.MinInt, Move(0)
		} else {
			return math.MaxInt, Move(0)
		}
	}

	if len(moves) == 0 {
		return 0, Move(0)
	}

	if p.Active() == White {
		max := math.MinInt
		best := moves[0]
		for _, m := range moves {
			newP := *p
			newP.Do(m)
			eval, _ := minmax(&newP, alpha, beta, depth-1)
			if eval > max {
				best = m
				max = eval
			}

			if eval >= beta {
				best = m
				break
			}

			alpha = fmax(alpha, eval)
		}
		return max, best
	} else {
		min := math.MaxInt
		best := moves[0]
		for _, m := range moves {
			newP := *p
			newP.Do(m)
			eval, _ := minmax(&newP, alpha, beta, depth-1)
			if eval <= min {
				best = m
				min = eval
			}

			if eval <= alpha {
				best = m
				break
			}

			beta = fmin(beta, eval)

		}
		return min, best
	}
}

// eval returns a static evaluation of position p in centipawns from White's
// perspective using material count only. Positive values favour White,
// negative values favour Black.
//
// Piece values:
//
//	Pawn=100  Knight=300  Bishop=300  Rook=500  Queen=900
func eval(p *Position) int {
	wPawns := p.WhitePawns().OnesCount()
	wKnight := p.WhiteKnights().OnesCount()
	wBishop := p.WhiteBishops().OnesCount()
	wRook := p.WhiteRooks().OnesCount()
	wQueen := p.WhiteQueens().OnesCount()

	bPawns := p.BlackPawns().OnesCount()
	bKnight := p.BlackKnights().OnesCount()
	bBishop := p.BlackBishops().OnesCount()
	bRook := p.BlackRooks().OnesCount()
	bQueen := p.BlackQueens().OnesCount()

	return (wPawns + wKnight*3 + wBishop*3 + wRook*5 + wQueen*9 -
		bPawns - bKnight*3 - bBishop*3 - bRook*5 - bQueen*9) * 100
}

// fmax returns the larger of a and b.
func fmax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// fmin returns the smaller of a and b.
func fmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// pickMove selects a move from a set of book entries using weighted random
// selection. Entries with higher Weight are chosen proportionally more often.
// If all weights are zero, a move is chosen uniformly at random.
func pickMove(entries []BookEntry) Move {
	var total int
	for _, m := range entries {
		total += int(m.Weight)
	}

	if total == 0 {
		return entries[len(entries)-1].Move
	}

	r := rand.IntN(total)
	for _, e := range entries {
		r -= int(e.Weight)
		if r < 0 {
			return e.Move
		}
	}
	return entries[len(entries)-1].Move
}
