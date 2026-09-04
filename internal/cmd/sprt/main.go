// Command sprt decides whether a change to the engine made it stronger.
//
// It plays the working tree against a baseline build in a self-play match and
// applies a Sequential Probability Ratio Test to the results, stopping as soon
// as the evidence is decisive rather than after a fixed number of games.
//
// The test is needed because chess results are extremely noisy: a change worth
// five Elo shifts the expected score from 50% to about 50.7%, which thousands
// of games can fail to distinguish by eye. It is also the only way to tune the
// pruning heuristics -- reduction formulas, null-move R, futility margins --
// since those change what the search returns, so node counts stop being a
// sufficient signal.
//
// Usage:
//
//	go run ./internal/cmd/openings -n 500 -plies 10 -o openings.fen
//	go run ./internal/cmd/sprt -openings openings.fen
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var (
		baseRef     = flag.String("baseline", "HEAD", "git ref to build the baseline from")
		baselineBin = flag.String("baseline-bin", "", "use this binary as the baseline instead of building one")
		candidateBn = flag.String("candidate-bin", "", "use this binary as the candidate instead of building the working tree")
		openingFile = flag.String("openings", "openings.fen", "file of opening positions, one FEN per line")
		tcSpec      = flag.String("tc", "8+0.08", "time control, base+increment in seconds")
		concurrency = flag.Int("concurrency", 0, "games in parallel (default: cores minus one)")
		maxGames    = flag.Int("maxgames", 40000, "stop after this many games even if undecided")
		elo0        = flag.Float64("elo0", 0, "H0: the Elo the candidate is worth if the change does nothing")
		elo1        = flag.Float64("elo1", 5, "H1: the Elo the candidate is worth if the change works")
		alpha       = flag.Float64("alpha", 0.05, "probability of accepting H1 when H0 is true")
		beta        = flag.Float64("beta", 0.05, "probability of accepting H0 when H1 is true")
		maxPlies    = flag.Int("maxplies", 400, "adjudicate a game as drawn after this many plies")
		seed        = flag.Uint64("seed", 1, "seed for shuffling the openings")
	)
	flag.Parse()

	if err := run(config{
		baseRef: *baseRef, baselineBin: *baselineBin, candidateBin: *candidateBn,
		openings: *openingFile, tc: *tcSpec, concurrency: *concurrency,
		maxGames: *maxGames, elo0: *elo0, elo1: *elo1, alpha: *alpha, beta: *beta,
		maxPlies: *maxPlies, seed: *seed,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "sprt: %s\n", err)
		os.Exit(1)
	}
}

type config struct {
	baseRef, baselineBin, candidateBin string
	openings, tc                       string
	concurrency, maxGames, maxPlies    int
	elo0, elo1, alpha, beta            float64
	seed                               uint64
}

func run(cfg config) error {
	tc, err := parseTimeControl(cfg.tc)
	if err != nil {
		return err
	}

	openings, err := readOpenings(cfg.openings)
	if err != nil {
		return err
	}
	rand.New(rand.NewPCG(cfg.seed, 0)).Shuffle(len(openings), func(i, j int) {
		openings[i], openings[j] = openings[j], openings[i]
	})

	workers := cfg.concurrency
	if workers <= 0 {
		if workers = runtimeCores() - 1; workers < 1 {
			workers = 1
		}
	}

	tmp, err := os.MkdirTemp("", "chester-sprt-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	basePath := cfg.baselineBin
	if basePath == "" {
		basePath = filepath.Join(tmp, "baseline")
		fmt.Printf("building baseline from %s... ", cfg.baseRef)
		if err := buildRef(cfg.baseRef, basePath); err != nil {
			fmt.Println("failed")
			return err
		}
		fmt.Println("ok")
	}

	candPath := cfg.candidateBin
	if candPath == "" {
		candPath = filepath.Join(tmp, "candidate")
		fmt.Print("building candidate from worktree... ")
		if err := build(".", candPath); err != nil {
			fmt.Println("failed")
			return err
		}
		fmt.Println("ok")
	}

	lower, upper := sprtBounds(cfg.alpha, cfg.beta)
	fmt.Printf("\n%d openings, tc %s, %d threads\n", len(openings), cfg.tc, workers)
	fmt.Printf("H0: %+.1f Elo   H1: %+.1f Elo   bounds [%.2f, %.2f]\n\n",
		cfg.elo0, cfg.elo1, lower, upper)
	fmt.Printf("%7s %6s %6s %6s %10s %8s\n", "games", "W", "L", "D", "elo", "LLR")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Each game is one opening played from one side. Consecutive pairs share
	// an opening with the colours swapped, so that an opening favouring White
	// cannot favour either engine.
	specs := make(chan int)
	go func() {
		defer close(specs)
		for i := 0; i < cfg.maxGames; i++ {
			select {
			case specs <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	results := make(chan outcome, workers)
	var wg sync.WaitGroup
	var failures []error
	var failMu sync.Mutex

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cand, err := startEngine(candPath, "candidate")
			if err != nil {
				failMu.Lock()
				failures = append(failures, err)
				failMu.Unlock()
				cancel()
				return
			}
			defer cand.close()

			base, err := startEngine(basePath, "baseline")
			if err != nil {
				failMu.Lock()
				failures = append(failures, err)
				failMu.Unlock()
				cancel()
				return
			}
			defer base.close()

			for i := range specs {
				o, err := playGame(cand, base, openings[(i/2)%len(openings)], i%2 == 0, tc, cfg.maxPlies)
				if err != nil {
					failMu.Lock()
					failures = append(failures, err)
					failMu.Unlock()
					cancel()
					return
				}
				select {
				case results <- o:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() { wg.Wait(); close(results) }()

	var t tally
	var verdict string
	reasons := map[string]int{}

	for o := range results {
		t.add(o.result)
		reasons[o.reason]++

		ratio := llr(t, cfg.elo0, cfg.elo1)

		if t.games()%20 == 0 || t.games() == 1 {
			report(t, ratio, "")
		}

		switch {
		case ratio >= upper:
			verdict = "H1 accepted"
		case ratio <= lower:
			verdict = "H0 accepted"
		}
		if verdict != "" {
			report(t, ratio, verdict)
			cancel()
			break
		}
	}

	// Drain whatever the workers produced while shutting down.
	for range results {
	}

	failMu.Lock()
	defer failMu.Unlock()
	if len(failures) > 0 && t.games() == 0 {
		return errors.Join(failures...)
	}

	summarise(t, cfg, verdict, reasons, failures)
	return nil
}

func report(t tally, ratio float64, note string) {
	value, _ := elo(t)
	fmt.Printf("%7d %6d %6d %6d %10s %8.2f  %s\n",
		t.games(), t.w, t.l, t.d, formatElo(value), ratio, note)
}

func summarise(t tally, cfg config, verdict string, reasons map[string]int, failures []error) {
	value, margin := elo(t)

	fmt.Println()
	switch verdict {
	case "H1 accepted":
		fmt.Printf("candidate is stronger (%s +/- %.1f Elo) after %d games\n",
			formatElo(value), margin, t.games())
	case "H0 accepted":
		fmt.Printf("candidate is not worth %+.1f Elo (%s +/- %.1f) after %d games\n",
			cfg.elo1, formatElo(value), margin, t.games())
	default:
		fmt.Printf("undecided after %d games (%s +/- %.1f Elo)\n",
			t.games(), formatElo(value), margin)
	}

	fmt.Println("\nhow games ended:")
	for reason, n := range reasons {
		fmt.Printf("  %-28s %d\n", reason, n)
	}

	if len(failures) > 0 {
		fmt.Printf("\n%d worker error(s); first: %v\n", len(failures), failures[0])
	}
}

func formatElo(v float64) string {
	if math.IsInf(v, 0) || math.IsNaN(v) {
		return "-"
	}
	return fmt.Sprintf("%+.1f", v)
}

// buildRef checks ref out into a throwaway worktree and builds it there, so
// that the baseline is the committed state of that ref and not whatever the
// working tree happens to contain.
func buildRef(ref, out string) error {
	dir, err := os.MkdirTemp("", "chester-baseline-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	if out, err := exec.Command("git", "worktree", "add", "--detach", dir, ref).CombinedOutput(); err != nil {
		return fmt.Errorf("git worktree add %s: %w\n%s", ref, err, out)
	}
	defer exec.Command("git", "worktree", "remove", "--force", dir).Run()

	return build(dir, out)
}

func build(dir, out string) error {
	abs, err := filepath.Abs(out)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "build", "-o", abs, "./cmd")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go build in %s: %w\n%s", dir, err, out)
	}
	return nil
}

func readOpenings(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading openings: %w (generate one with: go run ./internal/cmd/openings)", err)
	}

	var out []string
	for _, line := range strings.Split(string(data), "\n") {
		if line = strings.TrimSpace(line); line != "" && !strings.HasPrefix(line, "#") {
			out = append(out, line)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("%s contains no positions", path)
	}
	return out, nil
}

// parseTimeControl reads "base+increment" in seconds, as in "8+0.08".
func parseTimeControl(spec string) (timeControl, error) {
	base, inc, found := strings.Cut(spec, "+")
	if !found {
		inc = "0"
	}

	seconds := func(s string) (time.Duration, error) {
		v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return 0, fmt.Errorf("bad time control %q: %w", spec, err)
		}
		return time.Duration(v * float64(time.Second)), nil
	}

	b, err := seconds(base)
	if err != nil {
		return timeControl{}, err
	}
	i, err := seconds(inc)
	if err != nil {
		return timeControl{}, err
	}
	if b <= 0 {
		return timeControl{}, fmt.Errorf("time control %q has no base time", spec)
	}
	return timeControl{base: b, increment: i}, nil
}

func runtimeCores() int { return runtime.NumCPU() }
