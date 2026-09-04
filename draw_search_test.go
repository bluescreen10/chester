package chester_test

import (
	"testing"
	"time"

	"github.com/bluescreen10/chester"
)

func search(t *testing.T, p *chester.Position, opts *chester.SearchOptions) chester.Evaluation {
	t.Helper()

	ch, _ := chester.SearchBestMove(p, opts)
	var last chester.Evaluation
	for e := range ch {
		last = e
	}
	return last
}

// TestFiftyMoveRule checks that an overwhelming material advantage is scored
// as a draw once the half-move clock is about to reach 100.
func TestFiftyMoveRule(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		wantDraw bool
	}{
		{
			name:     "clock at zero, ordinary win",
			fen:      "8/8/8/3k4/8/8/3K4/6QR w - - 0 1",
			wantDraw: false,
		},
		{
			name:     "clock about to expire",
			fen:      "8/8/8/3k4/8/8/3K4/6QR w - - 99 1",
			wantDraw: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := chester.ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}

			eval := search(t, p, &chester.SearchOptions{
				MaxDepth: 4,
				MaxTime:  10 * time.Second,
			})

			switch {
			case test.wantDraw && eval.Score != 0:
				t.Errorf("score = %d, want 0 (draw by fifty-move rule)", eval.Score)
			case !test.wantDraw && eval.Score <= 0:
				t.Errorf("score = %d, want a winning score", eval.Score)
			}
		})
	}
}

// TestCheckmateBeatsFiftyMoveRule checks that mate delivered on the hundredth
// half-move still counts as mate. The rule is only a draw if the game has not
// already ended.
func TestCheckmateBeatsFiftyMoveRule(t *testing.T) {
	// White mates in one with Ra8#, on the hundredth half-move.
	p, err := chester.ParseFEN("7k/8/6K1/8/8/8/8/R7 w - - 99 1")
	if err != nil {
		t.Fatal(err)
	}

	eval := search(t, p, &chester.SearchOptions{
		MaxDepth: 3,
		MaxTime:  10 * time.Second,
	})

	if eval.Score < chester.MateScore-100 {
		t.Errorf("score = %d, want a mate score; best move was %s", eval.Score, eval.Best)
	}
	if eval.Best.String() != "a1a8" {
		t.Errorf("best = %s, want a1a8", eval.Best)
	}
}

// TestRepetitionIsADraw checks that repeating a position from the game
// history is scored as a draw rather than by material.
//
// White is a rook up, so a repetition is the worst available outcome and the
// search would never choose one voluntarily. SearchOptions.Moves is used to
// force it down the repeating line, which isolates the score of that line
// from White's decision about whether to enter it.
func TestRepetitionIsADraw(t *testing.T) {
	p, err := chester.ParseFEN("7k/8/8/8/8/8/8/R5K1 w - - 0 1")
	if err != nil {
		t.Fatal(err)
	}

	var history []uint64
	play := func(move string) {
		t.Helper()
		m, err := chester.ParseMove(move, p)
		if err != nil {
			t.Fatal(err)
		}
		history = append(history, p.Hash())
		p.Do(m)
	}

	// Shuffle both kings out and back. After these four plies the position on
	// the board is identical to the starting one, and Kg1-h1 would repeat the
	// position that followed the first ply.
	play("g1h1")
	play("h8g8")
	play("h1g1")
	play("g8h8")

	repeat, err := chester.ParseMove("g1h1", p)
	if err != nil {
		t.Fatal(err)
	}
	onlyRepeat := []chester.Move{repeat}

	winning := search(t, p, &chester.SearchOptions{
		MaxDepth: 4,
		MaxTime:  10 * time.Second,
		Moves:    onlyRepeat,
	})
	if winning.Score <= 0 {
		t.Fatalf("without history: score = %d, want a winning score for a rook up", winning.Score)
	}

	drawn := search(t, p, &chester.SearchOptions{
		MaxDepth: 4,
		MaxTime:  10 * time.Second,
		Moves:    onlyRepeat,
		History:  history,
	})

	t.Logf("Kg1-h1 scored %d without game history, %d with it",
		winning.Score, drawn.Score)

	if drawn.Score != 0 {
		t.Errorf("with history: score = %d, want 0 (draw by repetition)", drawn.Score)
	}
}

// TestRepetitionIsAvoidedWhenWinning is the counterpart: the same position,
// the same history, but with every move available. A side that is winning
// must decline the repetition.
func TestRepetitionIsAvoidedWhenWinning(t *testing.T) {
	p, err := chester.ParseFEN("7k/8/8/8/8/8/8/R5K1 w - - 0 1")
	if err != nil {
		t.Fatal(err)
	}

	var history []uint64
	play := func(move string) {
		t.Helper()
		m, _ := chester.ParseMove(move, p)
		history = append(history, p.Hash())
		p.Do(m)
	}
	play("g1h1")
	play("h8g8")
	play("h1g1")
	play("g8h8")

	eval := search(t, p, &chester.SearchOptions{
		MaxDepth: 4,
		MaxTime:  10 * time.Second,
		History:  history,
	})

	if eval.Best.String() == "g1h1" {
		t.Errorf("best = g1h1, want a move that does not repeat")
	}
	if eval.Score <= 0 {
		t.Errorf("score = %d, want a winning score", eval.Score)
	}
}

// TestEnPassantLegality covers the two ways an en passant capture can be
// illegal that ordinary pin and check masks do not catch, both of which
// stem from the captured pawn not standing on the destination square.
func TestEnPassantLegality(t *testing.T) {
	tests := []struct {
		name    string
		fen     string
		move    string
		isLegal bool
	}{
		{
			name: "capture would expose the king along the rank",
			// White Ka5 and pawn b5, Black pawn c5 and rook h5. Capturing
			// b5xc6 removes both pawns from rank 5 at once, uncovering the
			// rook. Neither pawn is pinned on its own.
			fen:     "8/8/8/KPp4r/8/8/8/7k w - c6 0 1",
			move:    "b5c6",
			isLegal: false,
		},
		{
			name:    "same rank but no slider behind",
			fen:     "8/8/8/KPp5/8/8/8/7k w - c6 0 1",
			move:    "b5c6",
			isLegal: true,
		},
		{
			name:    "king is not on the rank, so no discovery is possible",
			fen:     "7r/8/8/1Pp5/8/8/8/K6k w - c6 0 1",
			move:    "b5c6",
			isLegal: true,
		},
		{
			name: "en passant captures the checking pawn",
			// The black pawn double-pushed to c5 and now checks the king on
			// d4. Taking it en passant removes the checker, so it is a legal
			// evasion even though the destination square c6 is nowhere near
			// the check.
			fen:     "8/8/8/1Pp5/3K4/8/8/7k w - c6 0 1",
			move:    "b5c6",
			isLegal: true,
		},
		{
			name: "en passant does not address an unrelated check",
			// The rook on e8 checks the king on e1. The en passant capture
			// neither takes the checker nor blocks the file, so it must not
			// be generated as an evasion.
			fen:     "4r3/8/8/1Pp5/8/8/8/4K2k w - c6 0 1",
			move:    "b5c6",
			isLegal: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, err := chester.ParseFEN(test.fen)
			if err != nil {
				t.Fatal(err)
			}

			var moves []chester.Move
			moves, _ = chester.LegalMoves(moves, p)

			found := false
			for _, m := range moves {
				if m.String() == test.move {
					found = true
					break
				}
			}

			if found != test.isLegal {
				t.Errorf("%s generated = %v, want %v", test.move, found, test.isLegal)
			}
		})
	}
}
