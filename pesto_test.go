package chester

import (
	"math/rand/v2"
	"testing"
)

// TestPestoStateMatchesFullScan checks the invariant the whole incremental
// scheme rests on: updating the parent's state across a move must give
// exactly what computing the child's state from scratch would give.
//
// An incremental evaluation that drifts does not crash and does not fail a
// perft. It returns a plausible number that is quietly wrong, and the only
// symptom is lost games. This is the test that catches it, so it walks real
// trees rather than a list of positions: promotions, castling and en passant
// are reached by playing them.
func TestPestoStateMatchesFullScan(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 5))

	checked := 0
	for _, root := range evalPositions(t) {
		for game := 0; game < 80; game++ {
			before := *root
			acc := NewPestoState(&before)

			for ply := 0; ply < 60; ply++ {
				var moves []Move
				moves, _ = LegalMoves(moves, &before)
				if len(moves) == 0 {
					break
				}
				m := moves[r.IntN(len(moves))]

				after := before
				after.Do(m)

				got := acc.Updated(&before, &after)
				want := NewPestoState(&after)

				if got != want {
					t.Fatalf("state diverged after %s\nfrom: %s\ngot  %+v\nwant %+v",
						m, before.FEN(), got, want)
				}

				// The score has to agree too, not just the sums it is built
				// from: this is what the search actually reads.
				if gotScore, wantScore := got.Score(after.active, after.inactive), EvalPesto(&after); gotScore != wantScore {
					t.Fatalf("score diverged after %s: incremental %d, full %d\nfen: %s",
						m, gotScore, wantScore, after.FEN())
				}

				before, acc = after, got
				checked++
			}
		}
	}

	t.Logf("%d incremental updates agreed with a full rescan", checked)
}

// TestPestoStateSpecialMoves pins the move types that would need their own
// branch in a scheme that classified moves instead of diffing boards.
func TestPestoStateSpecialMoves(t *testing.T) {
	tests := []struct {
		name string
		fen  string
		move string
	}{
		{"en passant", "8/8/8/1Pp5/8/8/8/K6k w - c6 0 1", "b5c6"},
		{"white castles kingside", "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1", "e1g1"},
		{"white castles queenside", "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1", "e1c1"},
		{"black castles kingside", "r3k2r/8/8/8/8/8/8/R3K2R b KQkq - 0 1", "e8g8"},
		{"quiet promotion", "8/P7/8/8/8/8/8/K6k w - - 0 1", "a7a8q"},
		{"promotion to knight", "8/P7/8/8/8/8/8/K6k w - - 0 1", "a7a8n"},
		{"capturing promotion", "1n6/P7/8/8/8/8/8/K6k w - - 0 1", "a7b8q"},
		{"plain capture", "8/8/8/3p4/4N3/8/8/K6k w - - 0 1", "e4d6"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, err := ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}
			m, err := ParseMove(test.move, before)
			if err != nil {
				t.Fatal(err)
			}

			var legal []Move
			legal, _ = LegalMoves(legal, before)
			if !containsMove(legal, m) {
				t.Fatalf("%s is not legal in %s", test.move, test.fen)
			}

			after := *before
			after.Do(m)

			if got, want := NewPestoState(before).Updated(before, &after), NewPestoState(&after); got != want {
				t.Errorf("after %s: got %+v, want %+v", test.move, got, want)
			}
		})
	}
}

func containsMove(moves []Move, m Move) bool {
	for _, l := range moves {
		if l == m {
			return true
		}
	}
	return false
}

func BenchmarkPestoUpdated(b *testing.B) {
	p, _ := ParseFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
	var moves []Move
	moves, _ = LegalMoves(moves, p)

	afters := make([]Position, len(moves))
	for i, m := range moves {
		afters[i] = *p
		afters[i].Do(m)
	}

	acc := NewPestoState(p)
	b.ResetTimer()

	var sink PestoState
	for i := 0; i < b.N; i++ {
		sink = acc.Updated(p, &afters[i%len(afters)])
	}
	_ = sink
}
