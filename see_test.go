package chester

import "testing"

// TestPawnAttackersTo checks the inverted shifts against a rank-and-file
// reference. Getting a direction or a file mask wrong here would make SEE
// quietly miss pawn defenders, which are the defenders that matter most.
func TestPawnAttackersTo(t *testing.T) {
	for _, color := range []Color{White, Black} {
		for target := range Square(64) {
			got := pawnAttackersTo(target, color)

			var want Bitboard
			tr, tf := target.RankAndFile()
			for from := range Square(64) {
				fr, ff := from.RankAndFile()

				// A white pawn attacks the two squares diagonally ahead of
				// it, where ahead means towards rank 8.
				forward := 1
				if color == Black {
					forward = -1
				}
				if int(tr) == int(fr)+forward && abs(int(tf)-int(ff)) == 1 {
					want |= NewBitboardFromSquare(from)
				}
			}

			if got != want {
				t.Fatalf("%s pawns attacking %v: got %#016x, want %#016x",
					colorName(color), target, uint64(got), uint64(want))
			}
		}
	}
}

func colorName(c Color) string {
	if c == White {
		return "white"
	}
	return "black"
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// TestAttackersTo checks the combined attacker set against the move generator,
// which is independently tested by the perft suite: every legal move landing
// on a square must come from a piece the function reports as attacking it.
func TestAttackersTo(t *testing.T) {
	fens := []string{
		"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
		"rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
		"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1",
		"r2q1rk1/1b1nbppp/p2ppn2/1p6/3NPP2/1BN1B3/PPP3PP/R2Q1RK1 w - - 0 12",
	}

	for _, fen := range fens {
		p, err := ParseFEN(fen)
		if err != nil {
			t.Fatal(err)
		}

		occupied := p.Occupied()
		for _, side := range []*Position{p, mirrorToMove(p)} {
			var moves []Move
			moves, _ = LegalMoves(moves, side)

			for _, m := range moves {
				// Castling moves the king two squares without attacking the
				// destination, and en passant lands on an empty square.
				if side.mailbox[m.From()] == King {
					if d := int(m.To()) - int(m.From()); d == 2 || d == -2 {
						continue
					}
				}
				if side.mailbox[m.From()] == Pawn && m.To() == side.enPassantTarget {
					continue
				}
				if side.mailbox[m.To()] == Empty {
					continue // a quiet move is not an attack in this sense
				}

				attackers := p.attackersTo(m.To(), occupied)
				if attackers&NewBitboardFromSquare(m.From()) == 0 {
					t.Errorf("%s: %s captures but %v is not reported as attacking %v",
						fen, m, m.From(), m.To())
				}
			}
		}
	}
}

// mirrorToMove returns the same position with the other side to move, so the
// attacker set can be checked from both sides without writing every position
// out twice.
func mirrorToMove(p *Position) *Position {
	q := *p
	q.active, q.inactive = q.inactive, q.active
	q.enPassantTarget = SQ_NULL
	return &q
}

func TestSeeGE(t *testing.T) {
	tests := []struct {
		name string
		fen  string
		move string
		// the exchange's value, worked out by hand
		value int
	}{
		{
			name:  "free knight",
			fen:   "4k3/8/8/4n3/8/8/8/4RK2 w - - 0 1",
			move:  "e1e5",
			value: 320,
		},
		{
			name: "knight defended by a pawn costs a rook",
			// Rxe5 wins a knight, dxe5 wins a rook: 320 - 500.
			fen:   "4k3/8/3p4/4n3/8/8/8/4RK2 w - - 0 1",
			move:  "e1e5",
			value: -180,
		},
		{
			name: "pawn takes a defended queen",
			// exd5 wins a queen, cxd5 wins a pawn: 900 - 100.
			fen:   "4k3/8/2p5/3q4/4P3/8/8/4K3 w - - 0 1",
			move:  "e4d5",
			value: 800,
		},
		{
			name: "knight takes a defended pawn",
			// Nxe5 wins a pawn, dxe5 wins a knight: 100 - 320.
			fen:   "4k3/8/3p4/4p3/8/3N4/8/4K3 w - - 0 1",
			move:  "d3e5",
			value: -220,
		},
		{
			name: "battery on both sides",
			// Rxe3, Rxe3, Rxe3: white wins a rook, loses a rook, wins a rook.
			// Neither back rook is an attacker until the one in front of it
			// moves, so this only comes out right if x-rays are handled.
			fen:   "4r2k/8/8/8/8/4r3/4R3/4RK2 w - - 0 1",
			move:  "e2e3",
			value: 500,
		},
		{
			name:  "quiet move wins nothing",
			fen:   "4k3/8/8/8/8/8/8/4RK2 w - - 0 1",
			move:  "e1e5",
			value: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}
			m, err := ParseMove(test.move, p)
			if err != nil {
				t.Fatal(err)
			}

			var legal []Move
			legal, _ = LegalMoves(legal, p)
			if !containsMove(legal, m) {
				t.Fatalf("%s is not legal in %s", test.move, test.fen)
			}

			// seeGE is a threshold test, so the exact value is pinned by
			// bracketing it: true at the value, false one point above.
			if !seeGE(p, m, test.value) {
				t.Errorf("seeGE(%s, %d) = false, want true", test.move, test.value)
			}
			if seeGE(p, m, test.value+1) {
				t.Errorf("seeGE(%s, %d) = true, want false", test.move, test.value+1)
			}
		})
	}
}

// TestSeeGELosingCapturesAreRejected is the property quiescence will rely on.
func TestSeeGELosingCapturesAreRejected(t *testing.T) {
	p, err := ParseFEN("4k3/8/3p4/4n3/8/8/8/4RK2 w - - 0 1")
	if err != nil {
		t.Fatal(err)
	}
	m, _ := ParseMove("e1e5", p)

	if seeGE(p, m, 0) {
		t.Error("a capture that loses a rook for a knight passed a zero threshold")
	}
}
