package main

import "math"

// result is the outcome of one game, always from the candidate's point of
// view regardless of which colour it had.
type result int

const (
	loss result = iota
	draw
	win
)

// tally counts game outcomes.
type tally struct{ w, l, d int }

func (t tally) games() int { return t.w + t.l + t.d }

func (t *tally) add(r result) {
	switch r {
	case win:
		t.w++
	case loss:
		t.l++
	case draw:
		t.d++
	}
}

// eloToScore converts an Elo difference into the expected score per game,
// under the logistic model the Elo scale is defined by.
func eloToScore(elo float64) float64 {
	return 1 / (1 + math.Pow(10, -elo/400))
}

// scoreToElo is the inverse of eloToScore.
func scoreToElo(score float64) float64 {
	switch {
	case score <= 0:
		return math.Inf(-1)
	case score >= 1:
		return math.Inf(1)
	}
	return -400 * math.Log10(1/score-1)
}

// llr returns the log-likelihood ratio of the two hypotheses given the games
// played so far: that the candidate is worth elo0, against that it is worth
// elo1.
//
// This is the normal approximation to the true likelihood ratio, which is
// what makes it cheap enough to recompute after every game. It reads the
// observed score and its variance and asks which of the two hypothesised
// scores the evidence favours, scaled by how much evidence there is.
//
// The variance term is why draws matter: a result of 50 wins and 50 losses
// and one of 100 draws have the same mean, but the second is far more
// informative about a small difference, and its lower variance makes the
// ratio move faster.
func llr(t tally, elo0, elo1 float64) float64 {
	n := float64(t.games())
	if n == 0 {
		return 0
	}

	w, d := float64(t.w), float64(t.d)

	mean := (w + d/2) / n
	meanSquare := (w + d/4) / n
	variance := meanSquare - mean*mean

	// All games identical so far. There is no variance to divide by, and no
	// conclusion to draw yet either.
	if variance <= 0 {
		return 0
	}

	s0, s1 := eloToScore(elo0), eloToScore(elo1)
	return n * (s1 - s0) * (2*mean - s0 - s1) / (2 * variance)
}

// sprtBounds returns the thresholds the log-likelihood ratio is tested
// against. Crossing the upper bound accepts H1 (the candidate is at least
// elo1), crossing the lower bound accepts H0 (it is no better than elo0).
func sprtBounds(alpha, beta float64) (lower, upper float64) {
	return math.Log(beta / (1 - alpha)), math.Log((1 - beta) / alpha)
}

// elo returns the observed Elo difference and the half-width of its 95%
// confidence interval. Both are +/-Inf when every game went the same way,
// which the caller is expected to render rather than print.
func elo(t tally) (value, margin float64) {
	n := float64(t.games())
	if n == 0 {
		return 0, math.Inf(1)
	}

	w, d := float64(t.w), float64(t.d)
	mean := (w + d/2) / n
	value = scoreToElo(mean)

	meanSquare := (w + d/4) / n
	variance := meanSquare - mean*mean
	if variance <= 0 {
		return value, math.Inf(1)
	}

	// Propagate the standard error of the score through the Elo transform.
	stderr := math.Sqrt(variance / n)
	slope := 400 / (math.Ln10 * mean * (1 - mean))
	return value, 1.96 * stderr * slope
}
