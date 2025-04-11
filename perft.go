package main

import (
	"fmt"
	"io"
)

func Perft(p *Position, depth int, output io.Writer) int {
	var nodes int
	moves := movesPool.Get().(*[]Move)
	defer movesPool.Put(moves)
	*moves = (*moves)[:0]
	LegalMoves(moves, p)

	for _, m := range *moves {
		if depth == 1 {
			fmt.Fprintf(output, "%s: 1\n", m)
			nodes++
		} else {
			newPos := *p
			Do(&newPos, m)
			newNodes := perft(&newPos, depth-1)
			fmt.Fprintf(output, "%s: %d\n", m, newNodes)
			nodes += newNodes
		}
	}
	return nodes
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

	for _, m := range *moves {
		newPos := *p
		Do(&newPos, m)
		newNodes := perft(&newPos, depth-1)
		nodes += newNodes
	}
	return nodes
}
