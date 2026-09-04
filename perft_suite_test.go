package chester_test

import (
	"testing"

	"github.com/bluescreen10/chester"
)

// TestPerftSuite runs the six standard perft positions. They are chosen to
// exercise the move generator's edge cases -- castling through attacked
// squares, promotions, discovered checks and the en passant pin -- which a
// perft of the starting position alone never reaches.
//
// The expected node counts are the published values for these positions. A
// mismatch means the generator is producing illegal moves or missing legal
// ones, so this is the test to reach for before trusting any search result.
func TestPerftSuite(t *testing.T) {
	tests := []struct {
		name  string
		fen   string
		nodes []int // nodes[i] is perft(i+1)
	}{
		{
			name:  "position 1 (initial)",
			fen:   "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			nodes: []int{20, 400, 8902, 197281, 4865609},
		},
		{
			name:  "position 2 (kiwipete)",
			fen:   "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			nodes: []int{48, 2039, 97862, 4085603},
		},
		{
			name:  "position 3 (en passant pin)",
			fen:   "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",
			nodes: []int{14, 191, 2812, 43238, 674624, 11030083},
		},
		{
			name:  "position 4 (promotions)",
			fen:   "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1",
			nodes: []int{6, 264, 9467, 422333},
		},
		{
			name:  "position 5",
			fen:   "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
			nodes: []int{44, 1486, 62379, 2103487},
		},
		{
			name:  "position 6",
			fen:   "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
			nodes: []int{46, 2079, 89890, 3894594},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i, want := range test.nodes {
				depth := i + 1

				p, err := chester.ParseFEN(test.fen)
				if err != nil {
					t.Fatalf("parsing fen: %s", err)
				}

				got := 0
				for mc := range chester.Perft(p, depth) {
					got += mc.Count
				}

				if got != want {
					t.Errorf("perft(%d) = %d, want %d", depth, got, want)
				}
			}
		})
	}
}
