package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// const DefaultFEN = "r2qkb1r/ppp2ppp/4p1b1/4P3/4p3/2PB1P2/P1P3PP/R1BQK2R w KQkq - 0 11"
//const DefaultFEN = "rnbq1bnr/ppppkppp/4pP2/8/8/8/PPPPP1PP/RNBQKBNR b KQ - 0 3"

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	fen := flag.String("fen", DefaultFEN, "FEN string")
	depth := flag.Int("depth", 1, "depth")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	p, err := Parse(*fen)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Printf("Node(depth = %d): %d\n", *depth, Perft(p, *depth, true, os.Stdout))
}
