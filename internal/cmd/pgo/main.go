// Command pgo generates the profile-guided optimization profile that the
// build uses, at cmd/default.pgo.
//
// A PGO profile tells the compiler which call sites are hot, so it can spend
// its inlining and layout budget where it pays off. That makes the profile's
// composition part of the build: whatever the profile does not exercise, the
// compiler has no reason to optimize.
//
// Perft alone is not enough. It exercises move generation and Position.Do and
// nothing else, so a perft-only profile leaves the evaluation and search
// functions -- negamax, quiescence, EvalPesto, move ordering -- looking cold
// even though they dominate real play. This tool runs both workloads in one
// process and gives each an explicit share of the samples.
//
// Usage:
//
//	go run ./internal/cmd/pgo [flags]
//
// The default budgets weight the profile towards search, which is what the
// engine actually spends its time on during a game.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/bluescreen10/chester"
)

// perftPositions exercise move generation. They are the standard perft suite:
// between them they cover castling, promotions, en passant, discovered checks
// and pins, so the generator's branches are all represented rather than just
// the ones the opening position happens to reach.
var perftPositions = []string{
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
	"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
	"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",
	"r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
	"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
	"r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
}

// searchPositions exercise evaluation, quiescence and move ordering. They
// span the phases of a game so that the profile is not dominated by the
// branches one structure happens to take.
//
// None of them is in the opening book: a book hit returns without searching
// and would contribute no samples at all.
var searchPositions = []string{
	// Open tactical middlegames.
	"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
	"r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1",
	"2rr3k/pp3pp1/1nnqbN1p/3pN3/2pP4/2P3Q1/PPB4P/R4RK1 w - - 0 1",
	"rnbqkb1r/pp1p1ppp/2p5/4P3/2B5/8/PPP1NnPP/RNBQK2R w KQkq - 0 6",

	// Closed and manoeuvring positions, where evaluation matters more than
	// tactics and the search spends its time on quiet moves.
	"r1bqk2r/pp2bppp/2n1pn2/2pp4/3P1B2/2PBPN2/PP1N1PPP/R2QK2R w KQkq - 0 8",
	"r2q1rk1/1b1nbppp/p2ppn2/1p6/3NPP2/1BN1B3/PPP3PP/R2Q1RK1 w - - 0 12",
	"2r3k1/1p3pp1/p2p3p/P2Pp3/1PP1P3/5nP1/5P1P/2R1NK2 b - - 0 1",

	// Endgames, where the game-phase interpolation and the endgame tables
	// carry the score.
	"8/2p5/3p4/KP5r/1R4pk/8/4P1P1/8 w - - 0 1",
	"8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1",
	"6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1",
	"8/5ppp/8/8/2P5/1K6/5k2/8 w - - 0 1",
	"4k3/8/8/8/8/8/4P3/4K3 w - - 0 1",
}

func main() {
	var (
		out        = flag.String("o", filepath.Join("cmd", "default.pgo"), "output profile path")
		perftTime  = flag.Duration("perft", 4*time.Second, "CPU time budget for the move generation workload")
		searchTime = flag.Duration("search", 12*time.Second, "CPU time budget for the search and evaluation workload")
		perftDepth = flag.Int("perft-depth", 5, "perft depth per position")
		searchStep = flag.Duration("search-step", 2*time.Second, "time to spend on each search position before moving to the next")
		ttSize     = flag.Int("tt", 64, "transposition table size in MiB for the search workload")
	)
	flag.Parse()

	if err := run(*out, *perftTime, *searchTime, *perftDepth, *searchStep, *ttSize); err != nil {
		fmt.Fprintf(os.Stderr, "pgo: %s\n", err)
		os.Exit(1)
	}
}

func run(out string, perftBudget, searchBudget time.Duration, perftDepth int, searchStep time.Duration, ttSize int) error {
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	// Write to a temporary file and rename on success, so that a failed or
	// interrupted run cannot leave a truncated profile behind for the build
	// to pick up.
	tmp := out + ".tmp"
	file, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("creating profile: %w", err)
	}
	defer os.Remove(tmp)

	if err := pprof.StartCPUProfile(file); err != nil {
		file.Close()
		return fmt.Errorf("starting profile: %w", err)
	}

	fmt.Printf("collecting move generation samples for %s...\n", perftBudget)
	perftElapsed, perftNodes := runPerft(perftBudget, perftDepth)
	fmt.Printf("  %s, %d nodes, %s\n", perftElapsed.Round(time.Millisecond), perftNodes, formatNPS(perftNodes, perftElapsed))

	// Keep the phases from bleeding into each other: a collection triggered
	// by the first phase's garbage should not be charged to the second.
	runtime.GC()

	fmt.Printf("collecting search and evaluation samples for %s...\n", searchBudget)
	searchElapsed, searchNodes := runSearch(searchBudget, searchStep, ttSize)
	fmt.Printf("  %s, %d nodes, %s\n", searchElapsed.Round(time.Millisecond), searchNodes, formatNPS(searchNodes, searchElapsed))

	pprof.StopCPUProfile()

	if err := file.Close(); err != nil {
		return fmt.Errorf("closing profile: %w", err)
	}

	info, err := os.Stat(tmp)
	if err != nil {
		return fmt.Errorf("stat profile: %w", err)
	}
	if info.Size() == 0 {
		return fmt.Errorf("profile is empty; no samples were collected")
	}

	if err := os.Rename(tmp, out); err != nil {
		return fmt.Errorf("writing profile: %w", err)
	}

	total := perftElapsed + searchElapsed
	fmt.Printf("\nwrote %s (%d bytes)\n", out, info.Size())
	fmt.Printf("  move generation  %5.1f%% of %s\n", percent(perftElapsed, total), total.Round(time.Millisecond))
	fmt.Printf("  search and eval  %5.1f%%\n", percent(searchElapsed, total))
	fmt.Printf("\ninspect it with:  go tool pprof -top %s\n", out)
	return nil
}

// runPerft counts nodes over the perft positions, cycling through them until
// the budget is spent. The budget is checked between positions, so a run
// overshoots by at most one perft.
func runPerft(budget time.Duration, depth int) (time.Duration, int64) {
	start := time.Now()
	var nodes int64

	for i := 0; time.Since(start) < budget; i++ {
		fen := perftPositions[i%len(perftPositions)]

		p, err := chester.ParseFEN(fen)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pgo: skipping %q: %s\n", fen, err)
			continue
		}

		for mc := range chester.Perft(p, depth) {
			nodes += int64(mc.Count)
		}
	}

	return time.Since(start), nodes
}

// runSearch searches the search positions, cycling through them until the
// budget is spent. Each search is capped by step rather than by depth, which
// keeps the total close to the budget regardless of how fast the machine is,
// and lets iterative deepening decide how far it gets.
//
// One transposition table is shared across every search, as it would be
// across the moves of a real game.
func runSearch(budget, step time.Duration, ttSize int) (time.Duration, int64) {
	tt := chester.NewTranspositionTable(uint64(ttSize) * 1024 * 1024)

	start := time.Now()
	var nodes int64

	for i := 0; ; i++ {
		remaining := budget - time.Since(start)
		if remaining <= 0 {
			break
		}
		if remaining > step {
			remaining = step
		}

		fen := searchPositions[i%len(searchPositions)]

		p, err := chester.ParseFEN(fen)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pgo: skipping %q: %s\n", fen, err)
			continue
		}

		ch, cancel := chester.SearchBestMove(p, &chester.SearchOptions{
			MaxDepth:           99,
			MaxTime:            remaining,
			TranspositionTable: tt,
		})

		var last chester.Evaluation
		for e := range ch {
			last = e
		}
		cancel()

		nodes += last.Nodes
	}

	return time.Since(start), nodes
}

func percent(part, total time.Duration) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total) * 100
}

func formatNPS(nodes int64, d time.Duration) string {
	if d <= 0 {
		return "-"
	}
	return fmt.Sprintf("%.2f Mnps", float64(nodes)/d.Seconds()/1e6)
}
