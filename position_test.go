package main_test

import (
	"testing"

	chester "github.com/bluescreen10/chester"
)

func TestParseFen(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	expectedHash := 0x463b96181691fc9c
	p, err := chester.ParseFEN(fen)
	if err != nil {
		t.Fatal(err)
	}
	if got := p.Fen(); got != fen {
		t.Fatalf("ParseFen: got %s, want %s", got, fen)
	}

	if got := p.Hash(); got != uint64(expectedHash) {
		t.Fatalf("ParseFen: invalid hash got 0x%x, want 0x%x", got, expectedHash)
	}
}

func TestUpdateWhiteCastlingRights(t *testing.T) {
	tests := []struct {
		fen      string
		move     chester.Move
		expected string
	}{
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewCastleKingSideMove(chester.SQ_E1, chester.SQ_G1),
			expected: "4k3/8/8/8/3b4/8/8/R4RK1 b - - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewCastleQueenSideMove(chester.SQ_E1, chester.SQ_C1),
			expected: "4k3/8/8/8/3b4/8/8/2KR3R b - - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewMove(chester.SQ_A1, chester.SQ_B1, chester.Rook),
			expected: "4k3/8/8/8/3b4/8/8/1R2K2R b K - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewMove(chester.SQ_H1, chester.SQ_G1, chester.Rook),
			expected: "4k3/8/8/8/3b4/8/8/R3K1R1 b Q - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R b KQ - 0 1",
			move:     chester.NewMove(chester.SQ_D4, chester.SQ_A1, chester.Bishop),
			expected: "4k3/8/8/8/8/8/8/b3K2R w K - 0 2",
		},
		{
			fen:      "4k3/8/8/8/4b3/8/8/R3K2R b KQ - 0 1",
			move:     chester.NewMove(chester.SQ_E4, chester.SQ_H1, chester.Bishop),
			expected: "4k3/8/8/8/8/8/8/R3K2b w Q - 0 2",
		},
	}

	for _, test := range tests {
		pos, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		pos.Do(test.move)
		if got := pos.Fen(); got != test.expected {
			t.Errorf("Do(%s) failed to update position expected %s got %s", test.move, test.expected, got)
		}
	}
}

func TestUpdateBlackCastlingRights(t *testing.T) {
	tests := []struct {
		fen      string
		move     chester.Move
		expected []bool
	}{}

	for _, test := range tests {
		p, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		p.Do(test.move)
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
		move            chester.Move
		enPassantTarget chester.Square
	}{
		{
			fen:             "4k3/8/8/8/3Pp3/8/8/R3K2R b KQ d3 0 1",
			after:           "4k3/8/8/8/8/3p4/8/R3K2R w KQ - 0 2",
			move:            chester.NewEnPassantMove(chester.SQ_E4, chester.SQ_D3),
			enPassantTarget: chester.SQ_NULL,
		},
		{
			fen:             "4k3/8/8/8/4p3/8/3P4/R3K2R w KQ - 0 1",
			after:           "4k3/8/8/8/3Pp3/8/8/R3K2R b KQ d3 0 1",
			move:            chester.NewDoublePushMove(chester.SQ_D2, chester.SQ_D4),
			enPassantTarget: chester.SQ_D4,
		},
		{
			fen:             "rnbqkb1r/pppppppp/7n/P7/8/8/1PPPPPPP/RNBQKBNR b KQkq - 0 3",
			after:           "rnbqkb1r/p1pppppp/7n/Pp6/8/8/1PPPPPPP/RNBQKBNR w KQkq b6 0 4",
			move:            chester.NewDoublePushMove(chester.SQ_B7, chester.SQ_B5),
			enPassantTarget: chester.SQ_B5,
		},
	}

	for _, test := range tests {
		p, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		p.Do(test.move)
		if got := p.EnPassantTarget().Square(); got != test.enPassantTarget {
			t.Errorf("Do(%s) failed to update en passant Square expected %s got %s", test.move, test.enPassantTarget, got)
		}

		if got := p.Fen(); got != test.after {
			t.Errorf("Do(%s) failed to update position expected %s got %s", test.move, test.after, got)
		}
	}
}
