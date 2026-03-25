package chester_test

import (
	"testing"
	"time"

	"github.com/bluescreen10/chester"
)

func TestSearchBestMove_Tactics(t *testing.T) {
	tests := []struct {
		name     string
		fen      string
		depth    int
		wantMove string // algebraic notation
	}{
		{
			name:     "Mate in 1 - Back Rank",
			fen:      "6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1",
			depth:    2,
			wantMove: "d1d8",
		},
		{
			name:     "Mate in 1 - Scholar's Mate",
			fen:      "r1bqkb1r/pppp1ppp/2n2n2/4p2Q/2B1P3/8/PPPP1PPP/RNB1K1NR w KQkq - 0 1",
			depth:    2,
			wantMove: "h5f7",
		},
		{
			name:     "Take Hanging Queen",
			fen:      "rnb1kbnr/pppp1ppp/8/4p3/4P3/5q2/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
			depth:    2,
			wantMove: "g1f3", // or d1f3, but f3 is the destination
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p, _ := chester.ParseFEN(test.fen)
			opts := &chester.SearchOptions{
				MaxDepth: test.depth,
				MaxTime:  time.Second * 2,
			}

			ch, _ := chester.SearchBestMove(p, opts)

			var lastEval chester.Evaluation
			for e := range ch {
				lastEval = e
			}

			if lastEval.Best.String() != test.wantMove {
				t.Errorf("%s: got move %s, want %s", test.name, lastEval.Best, test.wantMove)
			}
		})
	}
}

func TestSearchBestMove_Cancellation(t *testing.T) {
	p, _ := chester.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	opts := &chester.SearchOptions{
		MaxDepth: 20,
		MaxTime:  20 * time.Millisecond,
	}

	start := time.Now()
	ch, _ := chester.SearchBestMove(p, opts)

	var finalDepth int
	for e := range ch {
		finalDepth = e.Depth
	}
	elapsed := time.Since(start)

	if finalDepth >= 20 {
		t.Errorf("Search did not cancel, reached depth %d", finalDepth)
	}

	if elapsed > 200*time.Millisecond {
		t.Errorf("Search took too long to cancel: %v", elapsed)
	}
}

func TestSearchBestMove_MateScore(t *testing.T) {
	p, _ := chester.ParseFEN("6k1/5ppp/8/8/8/8/5PPP/3R2K1 w - - 0 1")
	opts := &chester.SearchOptions{
		MaxDepth: 2,
	}

	ch, _ := chester.SearchBestMove(p, opts)

	var lastEval chester.Evaluation
	for e := range ch {
		lastEval = e
	}

	expectedMate := chester.MateScore - 1
	if lastEval.Score != expectedMate {
		t.Errorf("Expected mate score %d, got %d", expectedMate, lastEval.Score)
	}
}
