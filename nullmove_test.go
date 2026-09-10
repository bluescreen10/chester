package chester

import (
	"math/rand/v2"
	"testing"
)

// TestDoNullHashMatchesParsedPosition is the test that matters for DoNull.
//
// A wrong Zobrist update does not crash and does not change any move: it just
// makes the null move's position collide with unrelated entries in the
// transposition table, poisoning results in a way that looks like bad play.
// So the hash is checked against the one ParseFEN computes for the same
// position reached independently.
func TestDoNullHashMatchesParsedPosition(t *testing.T) {
	tests := []struct {
		name string
		fen  string
		want string // the same position with the turn passed
	}{
		{
			name: "no en passant right",
			fen:  "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
			want: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R b KQkq - 1 1",
		},
		{
			name: "en passant right that a pawn could use",
			// A white pawn on b5 could capture c6, so the en passant key is
			// part of the hash and passing has to fold it back out.
			fen:  "4k3/8/8/1Pp5/8/8/8/4K3 w - c6 0 1",
			want: "4k3/8/8/1Pp5/8/8/8/4K3 b - - 1 1",
		},
		{
			name: "en passant right no pawn could use",
			// No white pawn is adjacent to c5, so by the Polyglot convention
			// the key was never folded in and must not be folded out.
			fen:  "4k3/8/8/2p5/8/8/8/4K3 w - c6 0 1",
			want: "4k3/8/8/2p5/8/8/8/4K3 b - - 1 1",
		},
		{
			name: "black to move",
			fen:  "r1bqkbnr/pppp1ppp/2n5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 3 3",
			want: "r1bqkbnr/pppp1ppp/2n5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 4 4",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}
			want, err := ParseFEN(test.want)
			if err != nil {
				t.Fatal(err)
			}

			p.DoNull()

			if p.hash != want.hash {
				t.Errorf("hash after DoNull = %#x, want %#x\ngot  %s\nwant %s",
					p.hash, want.hash, p.FEN(), want.FEN())
			}
			if p.active != want.active {
				t.Errorf("side to move = %v, want %v", p.active, want.active)
			}
			if p.EnPassantTarget() != SQ_NULL {
				t.Errorf("en passant target survived: %v", p.EnPassantTarget())
			}
		})
	}
}

// TestDoNullLeavesPiecesAlone checks that passing changes nothing on the
// board. The evaluation accumulator is carried across a null move untouched,
// which is only sound because of this.
func TestDoNullLeavesPiecesAlone(t *testing.T) {
	r := rand.New(rand.NewPCG(11, 13))

	for _, root := range evalPositions(t) {
		p := *root
		for ply := 0; ply < 30; ply++ {
			before := p
			acc := NewPestoState(&p)

			after := p
			after.DoNull()

			if after.pieces != before.pieces || after.allPieces != before.allPieces {
				t.Fatalf("DoNull moved a piece\nfen %s", before.FEN())
			}
			if after.mailbox != before.mailbox {
				t.Fatalf("DoNull changed the mailbox\nfen %s", before.FEN())
			}
			if got := NewPestoState(&after); got != acc {
				t.Fatalf("evaluation state changed across a null move: %+v vs %+v", got, acc)
			}

			var moves []Move
			moves, _ = LegalMoves(moves, &p)
			if len(moves) == 0 {
				break
			}
			p.Do(moves[r.IntN(len(moves))])
		}
	}
}

func TestHasNonPawnMaterial(t *testing.T) {
	tests := []struct {
		fen               string
		white, blackWants bool
	}{
		{"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", true, true},
		{"4k3/pppppppp/8/8/8/8/PPPPPPPP/4K3 w - - 0 1", false, false},
		{"4k3/pppppppp/8/8/8/8/PPPPPPPP/R3K3 w - - 0 1", true, false},
		{"4k2r/pppppppp/8/8/8/8/PPPPPPPP/4K3 w - - 0 1", false, true},
		{"4k3/8/8/8/8/8/8/4K3 w - - 0 1", false, false},
	}

	for _, test := range tests {
		p, err := ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}
		if got := p.HasNonPawnMaterial(White); got != test.white {
			t.Errorf("%s: white = %v, want %v", test.fen, got, test.white)
		}
		if got := p.HasNonPawnMaterial(Black); got != test.blackWants {
			t.Errorf("%s: black = %v, want %v", test.fen, got, test.blackWants)
		}
	}
}

// TestNullMoveZugzwangGuard checks the guard that null move pruning most
// needs.
//
// In a king and pawn endgame the side to move can be losing precisely because
// it has to move. Allowed to pass, the search would conclude the position is
// fine and prune away the line that loses it. The guard is material based, so
// that is what is asserted here -- and then the position is searched to be
// sure nothing about it upsets the search.
func TestNullMoveZugzwangGuard(t *testing.T) {
	for _, fen := range []string{
		"8/8/8/4k3/8/4K3/4P3/8 w - - 0 1",
		"6k1/8/6K1/6P1/8/8/8/8 b - - 0 1",
		"8/p7/8/P7/8/8/8/k6K w - - 0 1",
	} {
		p, err := ParseFEN(fen)
		if err != nil {
			t.Fatal(err)
		}

		if p.HasNonPawnMaterial(p.active) {
			t.Errorf("%s: side to move has non-pawn material; null move would not be guarded", fen)
		}

		// The search still has to work here, guard or no guard.
		ch, _ := SearchBestMove(p, &SearchOptions{MaxDepth: 8, DisableBook: true})
		var last Evaluation
		for e := range ch {
			last = e
		}

		var legal []Move
		legal, _ = LegalMoves(legal, p)
		if !containsMove(legal, last.Best) {
			t.Errorf("%s: search returned %s, which is not legal", fen, last.Best)
		}
	}
}

// TestNullMoveKeepsTacticsFindable is the blunt safety net. Null move pruning
// returns a bound without searching the real moves, so a guard that is too
// loose shows up as tactics going missing.
func TestNullMoveKeepsTacticsFindable(t *testing.T) {
	tests := []struct {
		fen  string
		want string
	}{
		{"6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1", "d1d8"},
		{"r1bqkb1r/pppp1ppp/2n2n2/4p2Q/2B1P3/8/PPPP1PPP/RNB1K1NR w KQkq - 0 1", "h5f7"},
		{"8/8/8/1PK5/8/8/8/k7 w - - 0 1", ""},
	}

	for _, test := range tests {
		p, err := ParseFEN(test.fen)
		if err != nil {
			t.Fatal(err)
		}

		ch, _ := SearchBestMove(p, &SearchOptions{MaxDepth: 6, DisableBook: true})
		var last Evaluation
		for e := range ch {
			last = e
		}

		if test.want != "" && last.Best.String() != test.want {
			t.Errorf("%s: got %s, want %s", test.fen, last.Best, test.want)
		}
	}
}

// TestHashAgreesBetweenDoAndParseFEN is the general form of the bug the null
// move tests exposed.
//
// A position's hash must not depend on how it was reached. If the incremental
// updates in Do and the from-scratch computation behind ParseFEN disagree,
// then a position reached by playing moves and the same position loaded from
// a FEN land in different transposition table slots, and neither the table
// nor repetition detection can be trusted.
func TestHashAgreesBetweenDoAndParseFEN(t *testing.T) {
	r := rand.New(rand.NewPCG(17, 19))

	checked, withEnPassant := 0, 0
	for _, root := range evalPositions(t) {
		for game := 0; game < 40; game++ {
			p := *root

			for ply := 0; ply < 40; ply++ {
				reparsed, err := ParseFEN(p.FEN())
				if err != nil {
					t.Fatalf("%s: %s", p.FEN(), err)
				}
				if reparsed.hash != p.hash {
					t.Fatalf("hash depends on how the position was reached\nfen %s\nplayed %#x, parsed %#x",
						p.FEN(), p.hash, reparsed.hash)
				}

				if p.EnPassantTarget() != SQ_NULL {
					withEnPassant++
				}
				checked++

				var moves []Move
				moves, _ = LegalMoves(moves, &p)
				if len(moves) == 0 {
					break
				}
				p.Do(moves[r.IntN(len(moves))])
			}
		}
	}

	if withEnPassant == 0 {
		t.Error("no position with an en passant right was reached; the test is not covering the case")
	}
	t.Logf("%d positions agreed (%d of them with an en passant right)", checked, withEnPassant)
}
