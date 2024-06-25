package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	p, err := Parse(DefaultFEN)
	if err != nil {
		panic(err)
	}
	depth := 5
	fmt.Printf("Node(depth = %d): %d\n", depth, perft(p, depth, true))
}

var printboard = false

func perft(p Position, depth int, print bool) int {
	if depth == 0 {
		return 1
	}
	var nodes int
	for _, m := range LegalMoves(p) {
		p.DoMove(m)
		newNodes := perft(p, depth-1, false)
		if print {
			fmt.Printf("%s: %d\n", m, newNodes)
		}
		nodes += newNodes
		p.UndoMove(m)
	}
	return nodes
}
