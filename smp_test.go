package chester_test

import (
	"testing"
	"time"

	"github.com/bluescreen10/chester"
)

var smpPositions = []string{
	"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
	"r2q1rk1/1b1nbppp/p2ppn2/1p6/3NPP2/1BN1B3/PPP3PP/R2Q1RK1 w - - 0 12",
	"8/2p5/3p4/KP5r/1R4pk/8/4P1P1/8 w - - 0 1",
	"6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1",
}

func searchWith(t testing.TB, fen string, threads, depth int) chester.Evaluation {
	t.Helper()

	p, err := chester.ParseFEN(fen)
	if err != nil {
		t.Fatal(err)
	}

	ch, _ := chester.SearchBestMove(p, &chester.SearchOptions{
		MaxDepth:           depth,
		MaxTime:            2 * time.Minute,
		DisableBook:        true,
		Threads:            threads,
		TranspositionTable: chester.NewTranspositionTable(64 << 20),
	})

	var last chester.Evaluation
	for e := range ch {
		last = e
	}
	return last
}

// TestSingleThreadIsDeterministic pins the property the whole verification
// method in this project rests on: one thread must give the same answer every
// time, so that a refactor can be checked by comparing node counts.
func TestSingleThreadIsDeterministic(t *testing.T) {
	for _, fen := range smpPositions {
		first := searchWith(t, fen, 1, 8)

		for run := 0; run < 3; run++ {
			again := searchWith(t, fen, 1, 8)

			if again.Nodes != first.Nodes || again.Best != first.Best || again.Score != first.Score {
				t.Errorf("%s: run %d differed\nfirst %s %d %d nodes\nagain %s %d %d nodes",
					fen, run, first.Best, first.Score, first.Nodes,
					again.Best, again.Score, again.Nodes)
			}
		}
	}
}

// TestThreadsAgreeOnDepth checks that adding threads does not change what the
// search concludes at a fixed depth. Lazy SMP threads share only the table, so
// node counts and the exact score will differ -- but a fixed-depth search is
// still answering the same question, and a wildly different answer would mean
// the threads are corrupting each other's state rather than cooperating.
func TestThreadsAgreeOnDepth(t *testing.T) {
	for _, fen := range smpPositions {
		one := searchWith(t, fen, 1, 8)

		for _, threads := range []int{2, 4, 8} {
			many := searchWith(t, fen, threads, 8)

			if many.Depth != one.Depth {
				t.Errorf("%s: %d threads reached depth %d, 1 thread reached %d",
					fen, threads, many.Depth, one.Depth)
			}

			// Scores may drift slightly: a helper thread can leave a bound in
			// the table that changes a cutoff. A large gap is a bug.
			if diff := many.Score - one.Score; diff > 50 || diff < -50 {
				t.Errorf("%s: %d threads scored %d, 1 thread scored %d",
					fen, threads, many.Score, one.Score)
			}

			var legal []chester.Move
			p, _ := chester.ParseFEN(fen)
			legal, _ = chester.LegalMoves(legal, p)
			found := false
			for _, m := range legal {
				if m == many.Best {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s: %d threads returned %s, which is not legal", fen, threads, many.Best)
			}
		}
	}
}

// TestThreadsSearchMore is the point of the exercise under Lazy SMP: every
// thread searches the whole root independently, so more threads visit more
// nodes in the same number of plies. If the count does not grow, the helper
// threads are not running.
//
// Root splitting was tried instead and divides one tree between the threads,
// which keeps the node count flat. It reached depth faster and measured 186
// Elo weaker head to head, so this engine shares by table rather than by
// dividing work, and a flat node count here would mean the wrong scheme is
// wired in.
func TestThreadsSearchMore(t *testing.T) {
	const fen = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"

	one := searchWith(t, fen, 1, 8)
	four := searchWith(t, fen, 4, 8)

	t.Logf("1 thread %10d nodes   4 threads %10d nodes  (%.2fx)",
		one.Nodes, four.Nodes, float64(four.Nodes)/float64(one.Nodes))

	if four.Nodes <= one.Nodes {
		t.Errorf("4 threads visited %d nodes, 1 thread visited %d; the helpers did nothing",
			four.Nodes, one.Nodes)
	}
}
