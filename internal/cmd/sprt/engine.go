package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

// engine drives one UCI engine process.
//
// Output is pumped into a channel by a background goroutine rather than read
// on demand, so that waiting for a reply can be given a deadline. An engine
// that wedges has to be detectable, not merely slow: without a deadline a
// single stuck process would hang the whole match.
type engine struct {
	name string
	path string

	cmd   *exec.Cmd
	stdin io.WriteCloser
	lines chan string
	errs  chan error
}

func startEngine(path, name string) (*engine, error) {
	cmd := exec.Command(path)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("%s: stdin: %w", name, err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("%s: stdout: %w", name, err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("%s: start: %w", name, err)
	}

	e := &engine{
		name:  name,
		path:  path,
		cmd:   cmd,
		stdin: stdin,
		lines: make(chan string, 256),
		errs:  make(chan error, 1),
	}

	go func() {
		defer close(e.lines)
		scanner := bufio.NewScanner(stdout)
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		for scanner.Scan() {
			e.lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			select {
			case e.errs <- err:
			default:
			}
		}
	}()

	if err := e.handshake(); err != nil {
		e.close()
		return nil, err
	}
	return e, nil
}

// handshake completes the UCI startup exchange and turns off the opening
// book. Leaving the book on would have both engines answer the opening from
// the same table, which is exactly the source of variation the match is
// trying to control.
func (e *engine) handshake() error {
	if err := e.send("uci"); err != nil {
		return err
	}
	if _, err := e.await("uciok", 10*time.Second); err != nil {
		return err
	}
	if err := e.send("setoption name OwnBook value false"); err != nil {
		return err
	}
	return e.ready()
}

func (e *engine) ready() error {
	if err := e.send("isready"); err != nil {
		return err
	}
	_, err := e.await("readyok", 10*time.Second)
	return err
}

func (e *engine) newGame() error {
	if err := e.send("ucinewgame"); err != nil {
		return err
	}
	return e.ready()
}

func (e *engine) send(format string, args ...any) error {
	line := fmt.Sprintf(format, args...)
	if _, err := io.WriteString(e.stdin, line+"\n"); err != nil {
		return fmt.Errorf("%s: writing %q: %w", e.name, line, err)
	}
	return nil
}

// await reads output until a line whose first token is token, or the deadline
// passes. Everything before it is discarded: the protocol allows an engine to
// emit any amount of info in between.
func (e *engine) await(token string, timeout time.Duration) (string, error) {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	for {
		select {
		case line, ok := <-e.lines:
			if !ok {
				return "", fmt.Errorf("%s: exited while waiting for %q", e.name, token)
			}
			if first, _, _ := strings.Cut(strings.TrimSpace(line), " "); first == token {
				return line, nil
			}
		case err := <-e.errs:
			return "", fmt.Errorf("%s: reading output: %w", e.name, err)
		case <-deadline.C:
			return "", fmt.Errorf("%s: timed out waiting for %q", e.name, token)
		}
	}
}

// think asks for a move in the given position and returns it along with how
// long the engine took. grace is added to the engine's remaining clock to
// form the read deadline, so that a genuine loss on time is reported as such
// rather than as a protocol failure.
func (e *engine) think(fen string, moves []string, clocks clocks, grace time.Duration) (string, time.Duration, error) {
	position := "position fen " + fen
	if len(moves) > 0 {
		position += " moves " + strings.Join(moves, " ")
	}
	if err := e.send("%s", position); err != nil {
		return "", 0, err
	}

	err := e.send("go wtime %d btime %d winc %d binc %d",
		ms(clocks.white), ms(clocks.black), ms(clocks.winc), ms(clocks.binc))
	if err != nil {
		return "", 0, err
	}

	start := time.Now()
	line, err := e.await("bestmove", clocks.remaining+grace)
	elapsed := time.Since(start)
	if err != nil {
		return "", elapsed, err
	}

	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", elapsed, fmt.Errorf("%s: malformed bestmove %q", e.name, line)
	}
	return fields[1], elapsed, nil
}

func (e *engine) close() {
	if e.stdin != nil {
		_ = e.send("quit")
		_ = e.stdin.Close()
	}
	if e.cmd != nil && e.cmd.Process != nil {
		done := make(chan struct{})
		go func() { _ = e.cmd.Wait(); close(done) }()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			_ = e.cmd.Process.Kill()
			<-done
		}
	}
}

func ms(d time.Duration) int64 { return d.Milliseconds() }
