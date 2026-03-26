package chester_test

import (
	"testing"

	"github.com/bluescreen10/chester"
)

func TestParseSquare(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"Valid e4", "e4", "e4", false},
		{"Valid a1", "a1", "a1", false},
		{"Valid h8", "h8", "h8", false},
		{"Invalid file", "z4", "", true},
		{"Invalid rank", "e9", "", true},
		{"Too short", "e", "", true},
		{"Too long", "e44", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := chester.ParseSquare(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("ParseSquare(%s) error = %v, wantErr %v", test.input, err, test.wantErr)
				return
			}
			if !test.wantErr {
				if got.String() != test.want {
					t.Errorf("ParseSquare(%s) = %v, want %v", test.input, got, test.want)
				}
			}
		})
	}
}

func TestSquareRankAndFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		rank  int8
		file  int8
	}{
		{"a8", "a8", 7, 0},
		{"h8", "h8", 7, 7},
		{"a1", "a1", 0, 0},
		{"h1", "h1", 0, 7},
		{"e4", "e4", 3, 4},
		{"d5", "d5", 4, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sq, err := chester.ParseSquare(test.input)
			if err != nil {
				t.Fatalf("failed to parse square %s: %v", test.input, err)
			}

			if got := sq.Rank(); got != test.rank {
				t.Errorf("%s.Rank() = %d, want %d", test.input, got, test.rank)
			}

			if got := sq.File(); got != test.file {
				t.Errorf("%s.File() = %d, want %d", test.input, got, test.file)
			}
		})
	}
}
