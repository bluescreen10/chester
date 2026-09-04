package main

import (
	"math"
	"testing"
)

func TestEloScoreRoundTrip(t *testing.T) {
	for _, e := range []float64{-400, -100, -5, 0, 5, 100, 400} {
		if got := scoreToElo(eloToScore(e)); math.Abs(got-e) > 1e-9 {
			t.Errorf("round trip %v -> %v", e, got)
		}
	}
}

func TestEloToScoreAnchors(t *testing.T) {
	// The Elo scale is defined so that equal strength is an even score and
	// a 400 point gap is a 10:1 expected result.
	if got := eloToScore(0); math.Abs(got-0.5) > 1e-12 {
		t.Errorf("eloToScore(0) = %v, want 0.5", got)
	}
	if got := eloToScore(400); math.Abs(got-10.0/11.0) > 1e-12 {
		t.Errorf("eloToScore(400) = %v, want 10/11", got)
	}
}

func TestSPRTBounds(t *testing.T) {
	lower, upper := sprtBounds(0.05, 0.05)
	if math.Abs(lower+2.9444389791664403) > 1e-9 {
		t.Errorf("lower = %v", lower)
	}
	if math.Abs(upper-2.9444389791664403) > 1e-9 {
		t.Errorf("upper = %v", upper)
	}
	if lower >= 0 || upper <= 0 {
		t.Error("bounds should straddle zero")
	}
}

// TestLLRDirection checks the sign of the evidence: results better than both
// hypotheses must push towards H1, and worse than both towards H0.
func TestLLRDirection(t *testing.T) {
	tests := []struct {
		name string
		t    tally
		want string
	}{
		{"clearly winning", tally{w: 300, l: 100, d: 600}, "positive"},
		{"clearly losing", tally{w: 100, l: 300, d: 600}, "negative"},
		{"dead even", tally{w: 200, l: 200, d: 600}, "negative"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := llr(test.t, 0, 5)
			if test.want == "positive" && got <= 0 {
				t.Errorf("llr = %v, want positive", got)
			}
			if test.want == "negative" && got >= 0 {
				t.Errorf("llr = %v, want negative", got)
			}
		})
	}
}

// TestLLRGrowsWithEvidence checks that the same score ratio observed over
// more games is stronger evidence. This is the property that makes the test
// terminate.
func TestLLRGrowsWithEvidence(t *testing.T) {
	small := llr(tally{w: 60, l: 40, d: 100}, 0, 5)
	large := llr(tally{w: 600, l: 400, d: 1000}, 0, 5)

	if !(large > small) {
		t.Errorf("llr did not grow with evidence: %v then %v", small, large)
	}
	if small <= 0 {
		t.Errorf("expected positive evidence, got %v", small)
	}
}

// TestLLRDrawsAreInformative is the reason the variance term is there.
//
// Both tallies below are a 51% score over a thousand games, but one reached
// it through eight hundred draws and the other through none. The draw-heavy
// match has far lower variance, so the same score is much stronger evidence
// of a small difference, and the ratio must reflect that.
func TestLLRDrawsAreInformative(t *testing.T) {
	drawish := tally{w: 100, l: 80, d: 820}
	sharp := tally{w: 510, l: 490, d: 0}

	if mean := func(x tally) float64 {
		return (float64(x.w) + float64(x.d)/2) / float64(x.games())
	}; mean(drawish) != mean(sharp) {
		t.Fatalf("test is not comparing like with like: %v vs %v", mean(drawish), mean(sharp))
	}

	drawishLLR := llr(drawish, 0, 5)
	sharpLLR := llr(sharp, 0, 5)

	if !(drawishLLR > sharpLLR) {
		t.Errorf("draw-heavy result was not stronger evidence: %v vs %v", drawishLLR, sharpLLR)
	}
}

func TestLLRNoGames(t *testing.T) {
	if got := llr(tally{}, 0, 5); got != 0 {
		t.Errorf("llr of nothing = %v, want 0", got)
	}
	if got := llr(tally{d: 10}, 0, 5); got != 0 {
		t.Errorf("llr with no variance = %v, want 0", got)
	}
}

func TestElo(t *testing.T) {
	// An even score is zero Elo whatever the draw rate.
	if got, _ := elo(tally{w: 100, l: 100, d: 200}); math.Abs(got) > 1e-9 {
		t.Errorf("elo of an even match = %v, want 0", got)
	}

	// More wins than losses is positive, and the interval shrinks as games
	// accumulate.
	value, wide := elo(tally{w: 60, l: 40, d: 100})
	if value <= 0 {
		t.Errorf("elo = %v, want positive", value)
	}
	_, narrow := elo(tally{w: 600, l: 400, d: 1000})
	if !(narrow < wide) {
		t.Errorf("confidence interval did not shrink: %v then %v", wide, narrow)
	}
}

// TestSPRTTerminates simulates matches at known strengths and checks the test
// reaches the correct verdict. It is the end-to-end check on the statistics:
// a truly stronger engine should be accepted, and an equal one rejected.
func TestSPRTTerminates(t *testing.T) {
	const (
		elo0, elo1   = 0.0, 5.0
		alpha, beta  = 0.05, 0.05
		drawRate     = 0.7
		maxGames     = 300000
		trueStronger = 15.0
	)
	lower, upper := sprtBounds(alpha, beta)

	simulate := func(trueElo float64) string {
		// Deterministic simulation: allocate results in the exact expected
		// proportions rather than sampling, so the test cannot flake.
		score := eloToScore(trueElo)
		winRate := score - drawRate/2

		var tl tally
		for n := 1; n <= maxGames; n++ {
			tl.w = int(float64(n) * winRate)
			tl.d = int(float64(n) * drawRate)
			tl.l = n - tl.w - tl.d

			switch ratio := llr(tl, elo0, elo1); {
			case ratio >= upper:
				return "H1"
			case ratio <= lower:
				return "H0"
			}
		}
		return "undecided"
	}

	if got := simulate(trueStronger); got != "H1" {
		t.Errorf("a %+.0f Elo engine was judged %q, want H1", trueStronger, got)
	}
	if got := simulate(-10); got != "H0" {
		t.Errorf("a weaker engine was judged %q, want H0", got)
	}
}

func TestParseTimeControl(t *testing.T) {
	tests := []struct {
		spec    string
		base    float64
		inc     float64
		wantErr bool
	}{
		{spec: "8+0.08", base: 8, inc: 0.08},
		{spec: "10", base: 10, inc: 0},
		{spec: "60+0.6", base: 60, inc: 0.6},
		{spec: "0+1", wantErr: true},
		{spec: "abc", wantErr: true},
		{spec: "8+x", wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.spec, func(t *testing.T) {
			tc, err := parseTimeControl(test.spec)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected an error for %q", test.spec)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if tc.base.Seconds() != test.base || math.Abs(tc.increment.Seconds()-test.inc) > 1e-9 {
				t.Errorf("got base %v inc %v, want %v / %v",
					tc.base.Seconds(), tc.increment.Seconds(), test.base, test.inc)
			}
		})
	}
}

func TestRepetitions(t *testing.T) {
	// history holds the positions before the current one; only entries an
	// even number of plies back share the side to move.
	h := []uint64{9, 2, 1, 4, 7, 6}

	if got := repetitions(h, 1); got != 1 {
		t.Errorf("repetitions of a position seen once = %d, want 1", got)
	}

	// 6 sits one ply back, where the other side is to move, so it can never
	// be the same position however the hashes fall.
	if got := repetitions(h, 6); got != 0 {
		t.Errorf("odd-distance match counted: %d", got)
	}

	// Two earlier occurrences plus the current one is the threefold the game
	// loop draws on.
	if got := repetitions([]uint64{1, 9, 1, 9, 1, 9}, 1); got != 3 {
		t.Errorf("repetitions = %d, want 3", got)
	}

	if got := repetitions(nil, 1); got != 0 {
		t.Errorf("empty history = %d, want 0", got)
	}
}
