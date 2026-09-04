package chester

// drawScore is the value of a drawn position, from the perspective of the
// side to move.
const drawScore = 0

// mateThreshold is the lowest absolute score that still denotes a forced
// mate. Any score at or beyond it encodes a distance to mate and must be
// adjusted when it crosses in or out of the transposition table.
const mateThreshold = MateScore - maxPly

// isRepetition reports whether position p has already occurred, either
// earlier on the current search path or earlier in the game.
//
// The stack holds the hash of every position preceding p, so only entries an
// even number of plies back can match: those are the positions with the same
// side to move. The scan stops at the half-move clock because a pawn move or
// a capture is irreversible, and no position before it can ever recur.
//
// A single repetition is treated as a draw rather than waiting for the third
// occurrence. That is the standard approximation: if a position can be
// reached twice, the side that benefits can nearly always reach it a third
// time, and detecting it early prunes the entire repeating subtree.
func (ctx *searchCtx) isRepetition(p *Position) bool {
	n := len(ctx.stack)

	limit := int(p.halfMoves)
	if limit > n {
		limit = n
	}

	for i := 2; i <= limit; i += 2 {
		if ctx.stack[n-i] == p.hash {
			return true
		}
	}
	return false
}

// scoreToTT converts a score for storage in the transposition table.
//
// A mate score encodes "mate in N plies from the root", but a table entry is
// path independent: the same position can be reached at a different ply and
// the stored distance would then be wrong. Storing the distance relative to
// the node itself, and converting back on retrieval, keeps mate distances
// correct however the position is reached.
func scoreToTT(score, ply int) int {
	switch {
	case score >= mateThreshold:
		return score + ply
	case score <= -mateThreshold:
		return score - ply
	default:
		return score
	}
}

// scoreFromTT is the inverse of scoreToTT, converting a stored score back
// into one relative to the root of the current search.
func scoreFromTT(score, ply int) int {
	switch {
	case score >= mateThreshold:
		return score - ply
	case score <= -mateThreshold:
		return score + ply
	default:
		return score
	}
}
