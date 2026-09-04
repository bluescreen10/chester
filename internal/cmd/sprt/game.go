package main

import (
	"fmt"
	"time"

	"github.com/bluescreen10/chester"
)

// timeControl is a base time plus a per-move increment, as "8+0.08" means
// eight seconds each with eighty milliseconds added after every move.
type timeControl struct {
	base      time.Duration
	increment time.Duration
}

// clocks is what one side sees when it is asked to move.
type clocks struct {
	white, black time.Duration
	winc, binc   time.Duration

	// remaining is the mover's own clock, used to bound how long to wait.
	remaining time.Duration
}

// outcome is a finished game, reported from the candidate's point of view.
type outcome struct {
	result result
	reason string
	plies  int
}

// playGame plays one game between two engines and adjudicates it.
//
// The rules are applied by the engine's own library rather than reimplemented
// here: legality, checkmate, stalemate, the fifty-move rule and repetition all
// come from the same code the search uses. That keeps the arbiter honest about
// what the engine believes, and means the perft suite is also testing it.
func playGame(cand, base *engine, fen string, candWhite bool, tc timeControl, maxPlies int) (outcome, error) {
	p, err := chester.ParseFEN(fen)
	if err != nil {
		return outcome{}, fmt.Errorf("parsing opening %q: %w", fen, err)
	}

	for _, e := range []*engine{cand, base} {
		if err := e.newGame(); err != nil {
			return outcome{}, err
		}
	}

	// byColour[0] plays White. The opening FEN can have either side to move,
	// so the mover is taken from the position rather than assumed.
	var byColour [2]*engine
	if candWhite {
		byColour = [2]*engine{cand, base}
	} else {
		byColour = [2]*engine{base, cand}
	}

	// The candidate's result is the White result when it has White, and the
	// opposite when it does not.
	fromCandidate := func(whiteResult result) result {
		if candWhite {
			return whiteResult
		}
		switch whiteResult {
		case win:
			return loss
		case loss:
			return win
		}
		return draw
	}

	remaining := [2]time.Duration{tc.base, tc.base}

	var moves []string
	history := []uint64{}

	for ply := 0; ply < maxPlies; ply++ {
		var legal []chester.Move
		legal, inCheck := chester.LegalMoves(legal, p)

		side := p.Active()

		if len(legal) == 0 {
			if inCheck {
				// The side to move is mated, so the other side won.
				if side == chester.White {
					return outcome{fromCandidate(loss), "checkmate", ply}, nil
				}
				return outcome{fromCandidate(win), "checkmate", ply}, nil
			}
			return outcome{draw, "stalemate", ply}, nil
		}

		if p.HalfMoves() >= 100 {
			return outcome{draw, "fifty-move rule", ply}, nil
		}
		if repetitions(history, p.Hash()) >= 2 {
			// Two earlier occurrences plus this one is a threefold.
			return outcome{draw, "threefold repetition", ply}, nil
		}

		mover := byColour[side]
		c := clocks{
			white: remaining[chester.White], black: remaining[chester.Black],
			winc: tc.increment, binc: tc.increment,
			remaining: remaining[side],
		}

		best, elapsed, err := mover.think(fen, moves, c, 5*time.Second)
		if err != nil {
			// A crashed or wedged engine loses the game. Reporting it as an
			// error instead would abort the whole match over one bad game.
			return outcome{lossFor(mover, cand), fmt.Sprintf("%s failed: %v", mover.name, err), ply}, nil
		}

		remaining[side] -= elapsed
		if remaining[side] < 0 {
			return outcome{lossFor(mover, cand), mover.name + " lost on time", ply}, nil
		}
		remaining[side] += tc.increment

		m, err := chester.ParseMove(best, p)
		if err != nil {
			return outcome{lossFor(mover, cand), fmt.Sprintf("%s sent unparsable move %q", mover.name, best), ply}, nil
		}
		if !contains(legal, m) {
			return outcome{lossFor(mover, cand), fmt.Sprintf("%s played illegal move %s", mover.name, best), ply}, nil
		}

		history = append(history, p.Hash())
		moves = append(moves, best)
		p.Do(m)
	}

	return outcome{draw, "move limit", maxPlies}, nil
}

// lossFor reports the game as a loss for whichever engine mover is.
func lossFor(mover, cand *engine) result {
	if mover == cand {
		return loss
	}
	return win
}

// repetitions counts how many times hash already appears in history. Only
// positions an even number of plies back can match, since the other side is
// to move in the rest.
func repetitions(history []uint64, hash uint64) int {
	n := 0
	for i := len(history) - 2; i >= 0; i -= 2 {
		if history[i] == hash {
			n++
		}
	}
	return n
}

func contains(moves []chester.Move, m chester.Move) bool {
	for _, l := range moves {
		if l == m {
			return true
		}
	}
	return false
}
