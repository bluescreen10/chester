package chester_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bluescreen10/chester"
)

// benchPositions is a small, fixed suite used to compare search versions.
// It mixes an opening, a tactical middlegame (Kiwipete) and an endgame so
// that a change is not tuned to a single phase of the game.
var benchPositions = []struct {
	name string
	fen  string
}{
	{"kiwipete", "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"},
	{"endgame", "8/2p5/3p4/KP5r/1R4pk/8/4P1P1/8 w - - 0 1"},
	{"midgame", "r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1"},
	{"positional", "2rr3k/pp3pp1/1nnqbN1p/3pN3/2pP4/2P3Q1/PPB4P/R4RK1 w - - 0 1"},
}

// benchDepth is the fixed depth every bench position is searched to.
// Raise it as the engine gets stronger; comparisons are only meaningful
// between runs that used the same value.
const benchDepth = 8

// TestSearchBench reports nodes and time per position at a fixed depth.
// It is not an assertion; run it with -v to compare two versions of the
// search:
//
//	go test -run TestSearchBench -v ./...
func TestSearchBench(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping bench in short mode")
	}

	var totalNodes int64
	var totalTime time.Duration

	for _, bp := range benchPositions {
		p, err := chester.ParseFEN(bp.fen)
		if err != nil {
			t.Fatalf("%s: %s", bp.name, err)
		}

		opts := &chester.SearchOptions{
			MaxDepth:           benchDepth,
			MaxTime:            60 * time.Second,
			TranspositionTable: chester.NewTranspositionTable(64 * 1024 * 1024),
		}

		start := time.Now()
		ch, _ := chester.SearchBestMove(p, opts)
		var last chester.Evaluation
		for e := range ch {
			last = e
		}
		elapsed := time.Since(start)

		totalNodes += last.Nodes
		totalTime += elapsed

		t.Logf("%-12s depth %2d  best %-6s score %7d  nodes %10d  %8s  %s",
			bp.name, last.Depth, last.Best, last.Score, last.Nodes,
			elapsed.Round(time.Millisecond), formatNPS(last.Nodes, elapsed))
	}

	t.Logf("%-12s %36s nodes %10d  %8s  %s", "TOTAL", "",
		totalNodes, totalTime.Round(time.Millisecond), formatNPS(totalNodes, totalTime))
}

func formatNPS(nodes int64, d time.Duration) string {
	if d == 0 {
		return "-"
	}
	return fmt.Sprintf("%.2f Mnps", float64(nodes)/d.Seconds()/1e6)
}
