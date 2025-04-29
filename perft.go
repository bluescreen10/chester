package main

import "sync"

type MoveCount struct {
	Move  Move
	Count int
}

var movesPool = sync.Pool{
	New: func() any {
		var s []Move
		return &s
	},
}

func Perft(p *Position, depth int) <-chan MoveCount {
	ch := make(chan MoveCount, maxMoves)

	go func() {
		moves := movesPool.Get().(*[]Move)
		*moves = (*moves)[:0]
		LegalMoves(moves, p)

		var newPos Position

		for _, m := range *moves {
			if depth == 1 {
				ch <- MoveCount{Move: m, Count: 1}
			} else {
				newPos = *p
				Do(&newPos, m)
				newNodes := perft(&newPos, depth-1)
				ch <- MoveCount{Move: m, Count: newNodes}
			}
		}
		movesPool.Put(moves)
		close(ch)
	}()
	return ch
}

func perft(p *Position, depth int) int {
	var nodes int
	moves := movesPool.Get().(*[]Move)
	defer movesPool.Put(moves)
	*moves = (*moves)[:0]
	LegalMoves(moves, p)

	if depth == 1 {
		return len(*moves)
	}

	var newPos Position

	for _, m := range *moves {
		newPos = *p
		Do(&newPos, m)
		newNodes := perft(&newPos, depth-1)
		nodes += newNodes
	}
	return nodes
}
