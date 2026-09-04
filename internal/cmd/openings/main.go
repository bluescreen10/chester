// Command openings generates a file of opening positions for self-play
// testing.
//
// Two engines playing from the initial position play very nearly the same
// game every time, so a match started there measures almost nothing: the
// draw rate is enormous and the few decisive games turn on one repeated
// line. A shared set of varied openings is what makes the games independent
// enough to draw a conclusion from.
//
// The positions come from the engine's own Polyglot book, walked for a fixed
// number of plies with the book's own weights. That keeps them to real
// openings rather than random legal moves, needs nothing from the network,
// and is reproducible from a seed.
//
// Usage:
//
//	go run ./internal/cmd/openings -n 500 -plies 8 -o openings.fen
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/bluescreen10/chester"
)

func main() {
	var (
		count  = flag.Int("n", 500, "number of unique positions to generate")
		plies  = flag.Int("plies", 8, "how many book plies to play before recording the position")
		out    = flag.String("o", "openings.fen", "output file, one FEN per line")
		seed   = flag.Uint64("seed", 1, "random seed, for reproducible output")
		tries  = flag.Int("tries", 200, "give up after this many consecutive walks that produce nothing new")
		verify = flag.Bool("verify", true, "check that every generated position is legal and has moves")
	)
	flag.Parse()

	positions, err := generate(*count, *plies, *seed, *tries, *verify)
	if err != nil {
		fmt.Fprintf(os.Stderr, "openings: %s\n", err)
		os.Exit(1)
	}

	if err := write(*out, positions); err != nil {
		fmt.Fprintf(os.Stderr, "openings: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("wrote %s (%d unique positions, %d plies)\n", *out, len(positions), *plies)
	for i, fen := range positions {
		if i == 3 {
			fmt.Println("  ...")
			break
		}
		fmt.Printf("  %s\n", fen)
	}
}

// generate walks the book until it has count distinct positions, or until it
// stops finding new ones. A shallow book runs out of distinct lines well
// before an arbitrary count, so the give-up path is the normal one for large
// -n rather than an error case.
func generate(count, plies int, seed uint64, tries int, verify bool) ([]string, error) {
	r := rand.New(rand.NewPCG(seed, 0x9e3779b97f4a7c15))

	seen := make(map[string]bool, count)
	positions := make([]string, 0, count)

	stale := 0
	for len(positions) < count && stale < tries {
		fen, ok := walk(r, plies, verify)
		if !ok || seen[fen] {
			stale++
			continue
		}

		seen[fen] = true
		positions = append(positions, fen)
		stale = 0
	}

	if len(positions) == 0 {
		return nil, fmt.Errorf("no positions generated; is the book empty?")
	}
	return positions, nil
}

// walk plays plies book moves from the initial position and returns the FEN
// of where it lands. It reports false if the book ran out first, which is
// normal: the book is deeper down some lines than others.
func walk(r *rand.Rand, plies int, verify bool) (string, bool) {
	p, err := chester.ParseFEN(chester.DefaultFEN)
	if err != nil {
		return "", false
	}

	for i := 0; i < plies; i++ {
		m, ok := chester.BookMove(p, r.IntN)
		if !ok {
			return "", false
		}
		p.Do(m)
	}

	if verify {
		// A position with no legal moves is already over and cannot be used
		// as a starting point. This also re-parses the FEN, so a position
		// that does not survive a round trip is rejected rather than handed
		// to the engines.
		var moves []chester.Move
		if moves, _ = chester.LegalMoves(moves, p); len(moves) == 0 {
			return "", false
		}

		fen := p.FEN()
		if _, err := chester.ParseFEN(fen); err != nil {
			return "", false
		}
		return fen, true
	}

	return p.FEN(), true
}

func write(path string, positions []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, fen := range positions {
		if _, err := fmt.Fprintln(w, fen); err != nil {
			return err
		}
	}
	return w.Flush()
}
