package chester_test

import (
	"testing"
	"time"

	"github.com/bluescreen10/chester"
)

// TestNoisyMovesIncludesPromotions covers the move set the quiescence
// search runs on. A promotion is not a capture, but it is not quiet either,
// so NoisyMoves has to return it or quiescence stops with a pawn frozen one
// square short of becoming a queen.
func TestNoisyMovesIncludesPromotions(t *testing.T) {
	tests := []struct {
		name    string
		fen     string
		want    []string // moves that must be generated
		notWant []string // moves that must not be
	}{
		{
			name: "quiet promotion onto an empty square",
			fen:  "8/P7/8/8/8/8/8/K6k w - - 0 1",
			// Queen and knight only. A queen is a rook plus a bishop, so
			// rook and bishop promotions are only ever played to avoid
			// stalemate, which quiescence never reasons about.
			want:    []string{"a7a8q", "a7a8n"},
			notWant: []string{"a7a8r", "a7a8b"},
		},
		{
			name:    "capturing promotion is still generated",
			fen:     "1n6/P7/8/8/8/8/8/K6k w - - 0 1",
			want:    []string{"a7a8q", "a7a8n", "a7b8q", "a7b8n"},
			notWant: []string{"a7b8r", "a7b8b"},
		},
		{
			name: "ordinary pawn pushes are still excluded",
			fen:  "8/8/8/8/8/8/P7/K6k w - - 0 1",
			// A push to a3 or a4 is quiet and must stay out of quiescence.
			notWant: []string{"a2a3", "a2a4"},
		},
		{
			name: "promotion that blocks a check is generated",
			// The rook on a8 checks the king on h8 along the rank; promoting
			// on b8 interposes.
			fen:  "r6K/1P6/8/8/8/8/8/7k w - - 0 1",
			want: []string{"b7b8q", "b7a8q"},
		},
		{
			name: "promotion that ignores a check is not generated",
			// The rook on e8 checks the king on e1. Promoting on a8 does
			// nothing about it.
			fen:     "4r3/P7/8/8/8/8/8/4K2k w - - 0 1",
			notWant: []string{"a7a8q", "a7a8n"},
		},
		{
			name: "pinned pawn cannot promote",
			// The rook on a1 pins the pawn on a7 against the king on a2...
			// which it does not: the pawn is between them, so pushing it off
			// the file would expose the king.
			fen:     "8/P7/8/8/8/8/K7/r6k w - - 0 1",
			notWant: []string{"a7a8q"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := chester.ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}

			var moves []chester.Move
			moves, _ = chester.NoisyMoves(moves, p)

			generated := make(map[string]bool, len(moves))
			for _, m := range moves {
				generated[m.String()] = true
			}

			for _, want := range test.want {
				if !generated[want] {
					t.Errorf("%s missing from NoisyMoves, got %v", want, keys(generated))
				}
			}
			for _, notWant := range test.notWant {
				if generated[notWant] {
					t.Errorf("%s should not be in NoisyMoves, got %v", notWant, keys(generated))
				}
			}
		})
	}
}

// TestLegalMovesKeepsAllPromotions is the counterpart: the full move
// generator must still produce all four pieces. Rook promotion is the point
// of the Saavedra position -- 1.c8=R wins where 1.c8=Q is stalemate -- and
// perft counts depend on every one of them.
func TestLegalMovesKeepsAllPromotions(t *testing.T) {
	p, err := chester.ParseFEN("1n6/P7/8/8/8/8/8/K6k w - - 0 1")
	if err != nil {
		t.Fatal(err)
	}

	var moves []chester.Move
	moves, _ = chester.LegalMoves(moves, p)

	generated := make(map[string]bool, len(moves))
	for _, m := range moves {
		generated[m.String()] = true
	}

	for _, want := range []string{
		"a7a8q", "a7a8r", "a7a8b", "a7a8n",
		"a7b8q", "a7b8r", "a7b8b", "a7b8n",
	} {
		if !generated[want] {
			t.Errorf("%s missing from LegalMoves, got %v", want, keys(generated))
		}
	}
}

// TestSaavedra checks the position rook promotion exists for: promoting to a
// queen stalemates Black, so the only winning move is 1.c8=R.
func TestSaavedra(t *testing.T) {
	p, err := chester.ParseFEN("8/8/1KP5/3r4/8/8/8/k7 w - - 0 1")
	if err != nil {
		t.Fatal(err)
	}

	eval := search(t, p, &chester.SearchOptions{
		MaxDepth: 9,
		MaxTime:  30 * time.Second,
	})

	t.Logf("best %s score %d", eval.Best, eval.Score)

	if eval.Best.String() != "c6c7" {
		t.Errorf("best = %s, want c6c7", eval.Best)
	}
	if eval.Score <= 0 {
		t.Errorf("score = %d, want a winning score", eval.Score)
	}
}

func keys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
