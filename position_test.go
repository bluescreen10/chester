package main_test

import (
	"testing"

	pawned "github.com/bluescreen10/pawned"
)

func TestParseFen(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	p, err := pawned.Parse(fen)
	if err != nil {
		t.Fatal(err)
	}
	if got := p.Fen(); got != fen {
		t.Fatalf("Parse() got %s, want %s", got, fen)
	}
}

func TestUpdateWhiteCastlingRights(t *testing.T) {
	tests := []struct {
		fen      string
		move     pawned.Move
		expected []bool
	}{
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     pawned.NewCastleKingSideMove(pawned.SQ_E1, pawned.SQ_G1),
			expected: []bool{false, false},
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     pawned.NewCastleKingSideMove(pawned.SQ_E1, pawned.SQ_G1),
			expected: []bool{false, false},
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     pawned.NewCastleQueenSideMove(pawned.SQ_E1, pawned.SQ_C1),
			expected: []bool{false, false},
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     pawned.NewMove(pawned.SQ_A1, pawned.SQ_B1, pawned.Rook),
			expected: []bool{true, false},
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     pawned.NewMove(pawned.SQ_H1, pawned.SQ_G1, pawned.Rook),
			expected: []bool{false, true},
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R b KQ - 0 1",
			move:     pawned.NewMove(pawned.SQ_D4, pawned.SQ_A1, pawned.Bishop),
			expected: []bool{true, false},
		},
		{
			fen:      "4k3/8/8/8/4b3/8/8/R3K2R b KQ - 0 1",
			move:     pawned.NewMove(pawned.SQ_E4, pawned.SQ_H1, pawned.Bishop),
			expected: []bool{false, true},
		},
	}

	for _, test := range tests {
		pos, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		pawned.Do(&pos, test.move)
		if got := pos.CanWhiteCastleKingSide(); got != test.expected[0] {
			t.Errorf("Do(%s) failed to update white king side castling rights expected %v got %v", test.move, test.expected[0], got)
		}
		if got := pos.CanWhiteCastleQueenSide(); got != test.expected[1] {
			t.Errorf("Do(%s) failed to update white queen side castling rights expected %v got %v", test.move, test.expected[1], got)
		}
	}
}

func TestUpdateBlackCastlingRights(t *testing.T) {
	tests := []struct {
		fen      string
		move     pawned.Move
		expected []bool
	}{}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		pawned.Do(&p, test.move)
		if got := p.CanBlackCastleKingSide(); got != test.expected[0] {
			t.Errorf("Do(%s) failed to update black king side castling rights expected %v got %v", test.move, test.expected[0], got)
		}
		if got := p.CanBlackCastleQueenSide(); got != test.expected[1] {
			t.Errorf("Do(%s) failed to update black queen side castling rights expected %v got %v", test.move, test.expected[1], got)
		}
	}
}

func TestEnPassant(t *testing.T) {
	tests := []struct {
		fen             string
		after           string
		move            pawned.Move
		enPassantSquare pawned.Square
	}{
		{
			fen:             "4k3/8/8/8/3Pp3/8/8/R3K2R b KQ d3 0 1",
			after:           "4k3/8/8/8/8/3p4/8/R3K2R w KQ - 0 2",
			move:            pawned.NewEnPassantMove(pawned.SQ_E4, pawned.SQ_D3),
			enPassantSquare: pawned.SQ_NULL,
		},
		{
			fen:             "4k3/8/8/8/4p3/8/3P4/R3K2R w KQ - 0 1",
			after:           "4k3/8/8/8/3Pp3/8/8/R3K2R b KQ d3 0 1",
			move:            pawned.NewDoublePushMove(pawned.SQ_D2, pawned.SQ_D4),
			enPassantSquare: pawned.SQ_D3,
		},
		{
			fen:             "rnbqkb1r/pppppppp/7n/P7/8/8/1PPPPPPP/RNBQKBNR b KQkq - 0 3",
			after:           "rnbqkb1r/p1pppppp/7n/Pp6/8/8/1PPPPPPP/RNBQKBNR w KQkq b6 0 4",
			move:            pawned.NewDoublePushMove(pawned.SQ_B7, pawned.SQ_B5),
			enPassantSquare: pawned.SQ_B6,
		},
	}

	for _, test := range tests {
		p, err := pawned.Parse(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		pawned.Do(&p, test.move)
		if got := p.EnPassant.Square(); got != test.enPassantSquare {
			t.Errorf("Do(%s) failed to update en passant square expected %s got %s", test.move, test.enPassantSquare, got)
		}

		if got := p.Fen(); got != test.after {
			t.Errorf("Do(%s) failed to update position expected %s got %s", test.move, test.after, got)
		}
	}
}
