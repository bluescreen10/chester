package chester

import "testing"

// TestLMRReductionNeverStarvesTheSearch checks the bound that keeps the
// reduced search meaningful. Handing negamax a depth of zero or less would
// send it straight to quiescence, so a reduction has to leave at least one
// ply behind.
func TestLMRReductionNeverStarvesTheSearch(t *testing.T) {
	for depth := lmrMinDepth; depth < maxPly; depth++ {
		for i := lmrMinMove; i < maxMoves; i++ {
			r := lmrReduction(depth, i)

			if r < 0 {
				t.Fatalf("depth %d move %d: negative reduction %d", depth, i, r)
			}
			if remaining := depth - 1 - r; remaining < 1 {
				t.Fatalf("depth %d move %d: reduction %d leaves %d plies",
					depth, i, r, remaining)
			}
		}
	}
}

// TestLMRReductionGrows pins the shape of the table: later moves and deeper
// nodes are reduced more, and the growth is gentle rather than linear.
func TestLMRReductionGrows(t *testing.T) {
	if a, b := lmrReduction(8, 4), lmrReduction(8, 20); !(b > a) {
		t.Errorf("later move not reduced more: index 4 -> %d, index 20 -> %d", a, b)
	}
	if a, b := lmrReduction(4, 8), lmrReduction(16, 8); !(b > a) {
		t.Errorf("deeper node not reduced more: depth 4 -> %d, depth 16 -> %d", a, b)
	}

	// Logarithmic, not linear: quadrupling the move index must not quadruple
	// the reduction.
	small, large := lmrReduction(12, 5), lmrReduction(12, 20)
	if large >= small*4 {
		t.Errorf("reduction grew faster than logarithmically: %d then %d", small, large)
	}
}

// TestLMRReductionClampsInputs checks the table lookup cannot run off its
// bounds, since ply and move counts are both bounded by convention rather
// than by the type system.
func TestLMRReductionClampsInputs(t *testing.T) {
	for _, tc := range [][2]int{
		{maxPly, maxMoves},
		{maxPly + 100, maxMoves + 100},
		{lmrMinDepth, maxMoves * 2},
	} {
		r := lmrReduction(tc[0], tc[1])
		if r < 0 || r > tc[0] {
			t.Errorf("lmrReduction(%d, %d) = %d, out of range", tc[0], tc[1], r)
		}
	}
}

// TestSearchStillFindsTactics is the safety net for the reductions. LMR
// searches most moves shallower than they deserve, so a formula that is too
// aggressive shows up as tactics disappearing.
func TestSearchStillFindsTactics(t *testing.T) {
	tests := []struct {
		name  string
		fen   string
		depth int
		want  string
	}{
		{"back rank mate", "6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1", 4, "d1d8"},
		{"scholar's mate", "r1bqkb1r/pppp1ppp/2n2n2/4p2Q/2B1P3/8/PPPP1PPP/RNB1K1NR w KQkq - 0 1", 4, "h5f7"},
		{"saavedra first move", "8/8/1KP5/3r4/8/8/8/k7 w - - 0 1", 9, "c6c7"},
		{"win the queen", "rnb1kbnr/pppp1ppp/8/4p3/4P3/5q2/PPPP1PPP/RNBQKBNR w KQkq - 0 1", 4, "g1f3"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}

			ch, _ := SearchBestMove(p, &SearchOptions{
				MaxDepth:           test.depth,
				DisableBook:        true,
				TranspositionTable: NewTranspositionTable(16 << 20),
			})
			var last Evaluation
			for e := range ch {
				last = e
			}

			if last.Best.String() != test.want {
				t.Errorf("got %s (score %d), want %s", last.Best, last.Score, test.want)
			}
		})
	}
}
