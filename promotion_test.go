package chester_test

import (
	"testing"

	"github.com/bluescreen10/chester"
)

// TestCaptureMovesIncludesPromotions covers the move set the quiescence
// search runs on. A promotion is not a capture, but it is not quiet either,
// so CaptureMoves has to return it or quiescence stops with a pawn frozen one
// square short of becoming a queen.
func TestCaptureMovesIncludesPromotions(t *testing.T) {
	tests := []struct {
		name    string
		fen     string
		want    []string // moves that must be generated
		notWant []string // moves that must not be
	}{
		{
			name: "quiet promotion onto an empty square",
			fen:  "8/P7/8/8/8/8/8/K6k w - - 0 1",
			want: []string{"a7a8q", "a7a8r", "a7a8b", "a7a8n"},
		},
		{
			name: "capturing promotion is still generated",
			fen:  "1n6/P7/8/8/8/8/8/K6k w - - 0 1",
			want: []string{"a7a8q", "a7b8q"},
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
			moves, _ = chester.CaptureMoves(moves, p)

			generated := make(map[string]bool, len(moves))
			for _, m := range moves {
				generated[m.String()] = true
			}

			for _, want := range test.want {
				if !generated[want] {
					t.Errorf("%s missing from CaptureMoves, got %v", want, keys(generated))
				}
			}
			for _, notWant := range test.notWant {
				if generated[notWant] {
					t.Errorf("%s should not be in CaptureMoves, got %v", notWant, keys(generated))
				}
			}
		})
	}
}

func keys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
