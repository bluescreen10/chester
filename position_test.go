package chester_test

import (
	"testing"

	"github.com/bluescreen10/chester"
)

func TestParseFEN(t *testing.T) {
	tests := []struct {
		fen  string
		hash uint64
	}{
		{
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			hash: 0x463b96181691fc9c,
		},
		{
			fen:  "rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPPKPPP/RNBQ1BNR b kq - 0 3",
			hash: 0x652a607ca3f242c1,
		},
	}

	for _, test := range tests {
		p, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		if got := p.FEN(); got != test.fen {
			t.Fatalf("ParseFEN: got %s, want %s", got, test.fen)
		}

		if got := p.Hash(); got != test.hash {
			t.Fatalf("ParseFEN: invalid hash got 0x%x, want 0x%x", got, test.hash)
		}
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
			move:     chester.NewMove(chester.SQ_E1, chester.SQ_G1),
			expected: "4k3/8/8/8/3b4/8/8/R4RK1 b - - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewMove(chester.SQ_E1, chester.SQ_C1),
			expected: "4k3/8/8/8/3b4/8/8/2KR3R b - - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewMove(chester.SQ_A1, chester.SQ_B1),
			expected: "4k3/8/8/8/3b4/8/8/1R2K2R b K - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R w KQ - 0 1",
			move:     chester.NewMove(chester.SQ_H1, chester.SQ_G1),
			expected: "4k3/8/8/8/3b4/8/8/R3K1R1 b Q - 1 1",
		},
		{
			fen:      "4k3/8/8/8/3b4/8/8/R3K2R b KQ - 0 1",
			move:     chester.NewMove(chester.SQ_D4, chester.SQ_A1),
			expected: "4k3/8/8/8/8/8/8/b3K2R w K - 0 2",
		},
		{
			fen:      "4k3/8/8/8/4b3/8/8/R3K2R b KQ - 0 1",
			move:     chester.NewMove(chester.SQ_E4, chester.SQ_H1),
			expected: "4k3/8/8/8/8/8/8/R3K2b w Q - 0 2",
		},
	}

	for _, test := range tests {
		pos, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		pos.Do(test.move)
		if got := pos.FEN(); got != test.expected {
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
			move:            chester.NewMove(chester.SQ_E4, chester.SQ_D3),
			enPassantTarget: chester.SQ_NULL,
		},
		{
			fen:             "4k3/8/8/8/4p3/8/3P4/R3K2R w KQ - 0 1",
			after:           "4k3/8/8/8/3Pp3/8/8/R3K2R b KQ d3 0 1",
			move:            chester.NewMove(chester.SQ_D2, chester.SQ_D4),
			enPassantTarget: chester.SQ_D3,
		},
		{
			fen:             "rnbqkb1r/pppppppp/7n/P7/8/8/1PPPPPPP/RNBQKBNR b KQkq - 0 3",
			after:           "rnbqkb1r/p1pppppp/7n/Pp6/8/8/1PPPPPPP/RNBQKBNR w KQkq b6 0 4",
			move:            chester.NewMove(chester.SQ_B7, chester.SQ_B5),
			enPassantTarget: chester.SQ_B6,
		},
	}

	for _, test := range tests {
		p, err := chester.ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		p.Do(test.move)
		if got := p.EnPassantTarget(); got != test.enPassantTarget {
			t.Errorf("Do(%s) failed to update en passant Square expected %s got %s", test.move, test.enPassantTarget, got)
		}

		if got := p.FEN(); got != test.after {
			t.Errorf("Do(%s) failed to update position expected %s got %s", test.move, test.after, got)
		}
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		moves []chester.Move
		fen   string
		hash  uint64
	}{
		{
			moves: []chester.Move{chester.NewMove(chester.SQ_E2, chester.SQ_E4)},
			fen:   "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			hash:  0x823c9b50fd114196,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_E2, chester.SQ_E4),
				chester.NewMove(chester.SQ_D7, chester.SQ_D5),
			},
			fen:  "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 2",
			hash: 0x0756b94461c50fb0,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_E2, chester.SQ_E4),
				chester.NewMove(chester.SQ_D7, chester.SQ_D5),
				chester.NewMove(chester.SQ_E4, chester.SQ_E5),
			},
			fen:  "rnbqkbnr/ppp1pppp/8/3pP3/8/8/PPPP1PPP/RNBQKBNR b KQkq - 0 2",
			hash: 0x662fafb965db29d4,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_E2, chester.SQ_E4),
				chester.NewMove(chester.SQ_D7, chester.SQ_D5),
				chester.NewMove(chester.SQ_E4, chester.SQ_E5),
				chester.NewMove(chester.SQ_F7, chester.SQ_F5),
			},
			fen:  "rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 3",
			hash: 0x22a48b5a8e47ff78,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_E2, chester.SQ_E4),
				chester.NewMove(chester.SQ_D7, chester.SQ_D5),
				chester.NewMove(chester.SQ_E4, chester.SQ_E5),
				chester.NewMove(chester.SQ_F7, chester.SQ_F5),
				chester.NewMove(chester.SQ_E1, chester.SQ_E2),
			},
			fen:  "rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPPKPPP/RNBQ1BNR b kq - 1 3",
			hash: 0x652a607ca3f242c1,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_E2, chester.SQ_E4),
				chester.NewMove(chester.SQ_D7, chester.SQ_D5),
				chester.NewMove(chester.SQ_E4, chester.SQ_E5),
				chester.NewMove(chester.SQ_F7, chester.SQ_F5),
				chester.NewMove(chester.SQ_E1, chester.SQ_E2),
				chester.NewMove(chester.SQ_E8, chester.SQ_F7),
			},
			fen:  "rnbq1bnr/ppp1pkpp/8/3pPp2/8/8/PPPPKPPP/RNBQ1BNR w - - 2 4",
			hash: 0x00fdd303c946bdd9,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_A2, chester.SQ_A4),
				chester.NewMove(chester.SQ_B7, chester.SQ_B5),
				chester.NewMove(chester.SQ_H2, chester.SQ_H4),
				chester.NewMove(chester.SQ_B5, chester.SQ_B4),
				chester.NewMove(chester.SQ_C2, chester.SQ_C4),
			},
			fen:  "rnbqkbnr/p1pppppp/8/8/PpP4P/8/1P1PPPP1/RNBQKBNR b KQkq c3 0 3",
			hash: 0x3c8123ea7b067637,
		},

		{
			moves: []chester.Move{
				chester.NewMove(chester.SQ_A2, chester.SQ_A4),
				chester.NewMove(chester.SQ_B7, chester.SQ_B5),
				chester.NewMove(chester.SQ_H2, chester.SQ_H4),
				chester.NewMove(chester.SQ_B5, chester.SQ_B4),
				chester.NewMove(chester.SQ_C2, chester.SQ_C4),
				chester.NewMove(chester.SQ_B4, chester.SQ_C3),
				chester.NewMove(chester.SQ_A1, chester.SQ_A3),
			},
			fen:  "rnbqkbnr/p1pppppp/8/8/P6P/R1p5/1P1PPPP1/1NBQKBNR b Kkq - 1 4",
			hash: 0x5c3f9b829b279560,
		},
	}

	for _, test := range tests {
		p, err := chester.ParseFEN(chester.DefaultFEN)
		if err != nil {
			t.Fatal(err)
		}

		for _, m := range test.moves {
			p.Do(m)
		}

		if got := p.FEN(); got != test.fen {
			t.Errorf("Do(%s) failed to update position got %s, want %s", test.moves, got, test.fen)
		}

		if got := p.Hash(); got != test.hash {
			t.Errorf("Do(%s) failed to update hash got %x, want %x", test.moves, got, test.hash)
		}
	}
}
