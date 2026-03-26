package chester

// MoveCount pairs a root move with the number of leaf nodes reachable from it
// at the requested depth. It is the element type of the channel returned by
// Perft.
type MoveCount struct {
	Move  Move
	Count int
}

// Perft performs a performance test (node count) to the given depth.
// Returns a channel that yields MoveCount for each root move.
func Perft(p *Position, depth int) <-chan MoveCount {
	ch := make(chan MoveCount, 2)

	go func() {
		moves := make([]Move, 0, 1024)

		moves, _ = LegalMoves(moves, p)

		var newPos Position

		count := len(moves)
		for i := 0; i < count; i++ {
			m := moves[i]
			if depth == 1 {
				ch <- MoveCount{Move: m, Count: 1}
			} else {
				newPos = *p
				newPos.Do(m)
				newNodes := perft(&newPos, moves[count:], depth-1)
				ch <- MoveCount{Move: m, Count: newNodes}
			}
		}
		close(ch)
	}()
	return ch
}

// perft is the recursive inner implementation used by Perft. It reuses the
// tail of the provided moves slice as scratch space for each child position,
// avoiding allocations deeper in the tree. Bulk counting (returning
// len(moves) at depth 1 without recursing) keeps the leaf level fast.
func perft(p *Position, moves []Move, depth int) int {
	moves, _ = LegalMoves(moves, p)

	if depth == 1 {
		return len(moves)
	}

	var nodes int
	var newPos Position
	m := moves[len(moves):]
	for i := 0; i < len(moves); i++ {
		newPos = *p
		newPos.Do(moves[i])
		nodes += perft(&newPos, m, depth-1)
	}
	return nodes
}
