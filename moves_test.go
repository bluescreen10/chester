package main_test

import (
	"os"
	"slices"
	"testing"

	pawned "github.com/bluescreen10/pawned"
)

func TestMoveGen(t *testing.T) {
	p, err := pawned.Parse(pawned.DefaultFEN)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		depth int
		nodes int
	}{
		{1, 20},
		{2, 400},
		{3, 8902},
		{4, 197281},
		{5, 4865609},
		{6, 119060324},
		{7, 3195901860},
	}

	for _, test := range tests {
		if got := pawned.Perft(p, test.depth, false, os.Stdout); got != test.nodes {
			t.Fatalf("Perft(%d) = %d, want %d", test.depth, got, test.nodes)
		}
	}

	// depth := 5
	// fmt.Printf("Node(depth = %d): %d\n", depth, perft(p, depth, true))
	t.Fatal("not implemented")
}

func TestLegalMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expected: []string{
				"a2a3", "a2a4", "b1a3", "b1c3", "b2b3",
				"b2b4", "c2c3", "c2c4", "d2d3", "d2d4",
				"e2e3", "e2e4", "f2f3", "f2f4", "g1f3",
				"g1h3", "g2g3", "g2g4", "h2h3", "h2h4",
			}},
		{
			fen: "rnbqkbnr/1ppppp1p/6p1/p7/1P6/N4P2/P1PPP1PP/R1BQKBNR b KQkq - 1 3",
			expected: []string{
				"a5a4", "a5b4", "a8a6", "a8a7", "b7b5",
				"b7b6", "b8a6", "b8c6", "c7c5", "c7c6",
				"d7d5", "d7d6", "e7e5", "e7e6", "f7f5",
				"f7f6", "f8g7", "f8h6", "g6g5", "g8f6",
				"g8h6", "h7h5", "h7h6",
			}},
		{
			fen: "rnbqkbnr/2pppppp/1p6/p7/1P6/2P5/P2PPPPP/RNBQKBNR w KQkq - 0 3",
			expected: []string{
				"a2a3", "a2a4", "b1a3", "b4a5", "b4b5",
				"c1a3", "c1b2", "c3c4", "d1a4", "d1b3",
				"d1c2", "d2d3", "d2d4", "e2e3", "e2e4",
				"f2f3", "f2f4", "g1f3", "g1h3", "g2g3",
				"g2g4", "h2h3", "h2h4",
			}},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}

}

func TestWhitePawnMoves(t *testing.T) {
	fen := "r1b1kb1r/1P4P1/1n3n2/2PpP2p/8/pP6/P4P1P/4K3 w - d6 0 1"
	p, err := pawned.Parse(fen)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"b3b4", "b7a8b", "b7a8n", "b7a8q", "b7a8r",
		"b7b8b", "b7b8n", "b7b8q", "b7b8r", "b7c8b",
		"b7c8n", "b7c8q", "b7c8r", "c5b6", "c5c6",
		"c5d6", "e1d1", "e1d2", "e1e2", "e1f1",
		"e5d6", "e5e6", "e5f6", "f2f3", "f2f4",
		"g7f8b", "g7f8n", "g7f8q", "g7f8r", "g7g8b",
		"g7g8n", "g7g8q", "g7g8r", "g7h8b", "g7h8n",
		"g7h8q", "g7h8r", "h2h3", "h2h4",
	}

	var moves []pawned.Move
	var got []string
	pawned.LegalMoves(&moves, &p)
	for _, move := range moves {
		got = append(got, move.String())
	}
	slices.Sort(got)

	if !slices.Equal(got, expected) {
		t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", fen, len(got), got, len(expected), expected)
	}
}

func TestBlackPawnMoves(t *testing.T) {
	fen := "4k3/p4p1p/8/Pp6/2pPp2P/1N3N2/1p4p1/R1B1KB1R b - d3 0 1"
	p, err := pawned.Parse(fen)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"a7a6", "b2a1b", "b2a1n", "b2a1q", "b2a1r",
		"b2b1b", "b2b1n", "b2b1q", "b2b1r", "b2c1b",
		"b2c1n", "b2c1q", "b2c1r", "b5b4", "c4b3",
		"c4c3", "c4d3", "e4d3", "e4e3", "e4f3",
		"e8d7", "e8d8", "e8e7", "e8f8", "f7f5",
		"f7f6", "g2f1b", "g2f1n", "g2f1q", "g2f1r",
		"g2g1b", "g2g1n", "g2g1q", "g2g1r", "g2h1b",
		"g2h1n", "g2h1q", "g2h1r", "h7h5", "h7h6",
	}

	var moves []pawned.Move
	var got []string
	pawned.LegalMoves(&moves, &p)
	for _, move := range moves {
		got = append(got, move.String())
	}
	slices.Sort(got)

	if !slices.Equal(got, expected) {
		t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", fen, len(got), got, len(expected), expected)
	}
}

func TestWhiteKnightMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "3k4/8/8/8/2p5/p1Pp4/P2P4/KN6 w - - 0 1",
			expected: []string{"b1a3"},
		},
		{
			fen:      "3k4/5p2/8/2n1N3/1bP1b3/pPp5/Kp6/8 w - - 0 1",
			expected: []string{"e5c6", "e5d3", "e5d7", "e5f3", "e5f7", "e5g4", "e5g6"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestBlackKnightMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "8/8/8/Bn6/B7/p1P2N2/K7/2k5 b - - 0 1",
			expected: []string{"b5a7", "b5c3", "b5c7", "b5d4", "b5d6"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestWhiteBishopMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "7b/4p3/4P3/8/r7/pB6/K1p5/2k5 w - - 0 1",
			expected: []string{"b3a4", "b3c2", "b3c4", "b3d5"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestBlackBishopMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "7k/6b1/4QBB1/8/8/8/8/3K4 b - - 0 1",
			expected: []string{"g7f6"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestWhiteRookMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "7k/8/4bb2/4np2/3ppKR1/6P1/8/8 w - - 0 1",
			expected: []string{"g4g5", "g4g6", "g4g7", "g4g8", "g4h4"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestBlackRookMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "3nk3/1pPrPp2/1Pp1pP2/2P1P3/3P4/3K4/8/8 b - - 0 1",
			expected: []string{"d7c7", "d7d4", "d7d5", "d7d6", "d7e7"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestWhiteQueenMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen: "8/3k4/8/8/1q6/b2r4/p7/KQ1n4 w - - 0 1",
			expected: []string{
				"a1a2", "b1a2", "b1b2", "b1b3", "b1b4",
				"b1c1", "b1c2", "b1d1", "b1d3",
			},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestBlackQueenMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen:      "7R/3k4/3q4/8/8/3Q2B1/8/K1R1R3 b - - 0 1",
			expected: []string{"d6d3", "d6d4", "d6d5"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestWhiteKingMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen: "8/3k4/8/8/8/2bb4/2PP4/R3K3 w Q - 0 1",
			expected: []string{
				"a1a2", "a1a3", "a1a4", "a1a5",
				"a1a6", "a1a7", "a1a8", "a1b1", "a1c1",
				"a1d1", "c2d3", "d2c3", "e1c1", "e1d1",
				"e1f2"},
		},
		{
			fen: "8/3k4/8/8/8/2bb4/2PP4/R3K3 w - - 0 1",
			expected: []string{
				"a1a2", "a1a3", "a1a4", "a1a5",
				"a1a6", "a1a7", "a1a8", "a1b1", "a1c1",
				"a1d1", "c2d3", "d2c3", "e1d1", "e1f2",
			},
		},
		{
			fen: "8/3k4/8/8/8/2bb4/P1PP1r1P/RN2K2R w KQ - 0 1",
			expected: []string{
				"a2a3", "a2a4", "b1a3", "b1c3", "c2d3",
				"d2c3", "e1d1", "e1f2", "h1f1", "h1g1",
				"h2h3", "h2h4",
			},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}

func TestBlackKingMoves(t *testing.T) {
	tests := []struct {
		fen      string
		expected []string
	}{
		{
			fen: "r3k2r/p1p1pRp1/p1p1p1p1/N1N1N1N1/8/8/8/4K3 b kq - 0 1",
			expected: []string{
				"a8b8", "a8c8", "a8d8", "e8c8", "e8d8",
				"h8f8", "h8g8", "h8h1", "h8h2", "h8h3",
				"h8h4", "h8h5", "h8h6", "h8h7",
			},
		},
		{
			fen: "r3k2r/p1pRp1pp/p1p1p1p1/N1N1N1N1/8/8/8/4K3 b kq - 0 1",
			expected: []string{
				"a8b8", "a8c8", "a8d8", "e8f8", "e8g8",
				"h7h5", "h7h6", "h8f8", "h8g8",
			},
		},
		{
			fen:      "r3k2r/p1pBpBpp/p1p1p1p1/N1N1N1N1/8/8/8/4K3 b kq - 0 1",
			expected: []string{"e8d8", "e8f8"},
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		var moves []pawned.Move
		var got []string
		pawned.LegalMoves(&moves, &p)
		for _, move := range moves {
			got = append(got, move.String())
		}
		slices.Sort(got)

		if !slices.Equal(got, test.expected) {
			t.Fatalf("LegalMoves(%s) got(%d) %s, want(%d) %s", test.fen, len(got), got, len(test.expected), test.expected)
		}
	}
}
