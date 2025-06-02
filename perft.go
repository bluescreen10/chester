package main

type MoveCount struct {
	Move  Move
	Count int
}

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
				Do(&newPos, m)
				newNodes := perft(&newPos, moves[count:], depth-1)
				ch <- MoveCount{Move: m, Count: newNodes}
			}
		}
		close(ch)
	}()
	return ch
}

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
		Do(&newPos, moves[i])
		nodes += perft(&newPos, m, depth-1)
	}
	return nodes
}
