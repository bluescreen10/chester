package chester

import (
	"math/rand/v2"
	"testing"
)

// TestBookNeverPlaysAnIllegalMove walks the opening book exhaustively for a
// few plies and checks that every move it offers is one the move generator
// agrees is legal.
//
// This is the test that catches a bad Polyglot conversion. The book file
// writes castling as the king capturing its own rook and numbers promotion
// pieces its own way; get either wrong in internal/cmd/book and the moves
// that come back out are ones LegalMoves never produces.
func TestBookNeverPlaysAnIllegalMove(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))

	const (
		walks = 300
		plies = 12
	)

	castlesSeen := 0

	for w := 0; w < walks; w++ {
		p, err := ParseFEN(DefaultFEN)
		if err != nil {
			t.Fatal(err)
		}

		line := ""
		for ply := 0; ply < plies; ply++ {
			m, ok := BookMove(p, r.IntN)
			if !ok {
				break
			}

			var legal []Move
			legal, _ = LegalMoves(legal, p)

			found := false
			for _, l := range legal {
				if l == m {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("book offered illegal move %s (bits %#x) after%s\nfen: %s",
					m, uint16(m), line, p.FEN())
			}

			if p.mailbox[m.From()] == King {
				if d := int(m.To()) - int(m.From()); d == 2 || d == -2 {
					castlesSeen++
				}
			}

			line += " " + m.String()
			p.Do(m)
		}
	}

	// The conversion is only exercised if the walk actually reaches a castle.
	if castlesSeen == 0 {
		t.Error("no castling move was reached; the test is not covering the conversion")
	}
	t.Logf("%d castling moves decoded across %d walks", castlesSeen, walks)
}
