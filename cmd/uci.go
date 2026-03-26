package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bluescreen10/chester"
)

// UCIServer handles communication between the chess engine and a UCI-compliant
// GUI. It manages the engine's state, position, and search execution.
type UCIServer struct {
	mutex          sync.Mutex
	pos            *chester.Position
	bestMove       string
	isCPUProfiling bool
	CPUProfileFile *os.File
	isDebugLogging bool
	tt             *chester.TranspositionTable
	stopFunc       func()
}

// startUCI initializes a standard UCI session.
func startUCI() {
	pos, _ := chester.ParseFEN(chester.DefaultFEN)
	uci := &UCIServer{pos: pos, tt: chester.NewTranspositionTable(64 * 1024 * 1024)}
	uci.Start()
}

// Start runs the main loop of the UCI server, reading commands from standard
// input and dispatching them to the appropriate handlers.
func (s *UCIServer) Start() {
	s.info("starting uci server...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		s.debug("received: %s", line)

		args := strings.Split(line, " ")
		s.debug("command: %s args: %v", args[0], args[1:])

		switch args[0] {
		case "quit":
			s.info("quitting uci server...")
			return
		case "uci":
			s.handleUCI()
		case "ucinewgame":
			s.handleUCINewGame()
		case "position":
			s.handlePosition(args[1:])
		case "go":
			s.handleGo(args[1:])
		case "stop":
			s.handleStop()
		case "isready":
			s.handleIsReady()
		case "perft":
			s.handlePerft(args[1:])
		case "cpuprofile":
			s.handleCPUProfile(args[1:])
		case "debug":
			s.handleDebug(args[1:])
		default:
			s.error("unknown command: %s", args[0])
		}
	}
}

// Write sends a raw byte slice to standard output, which is the protocol's
// communication channel with the GUI.
func (s *UCIServer) Write(msg []byte) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return os.Stdout.Write(msg)
}

// WriteString sends a formatted message to the GUI. It automatically appends
// a newline character.
func (s *UCIServer) WriteString(msg string, args ...any) {
	s.Write([]byte(fmt.Sprintf(msg+"\n", args...)))
}

// debug logs a message to the GUI if debug logging is enabled.
func (s *UCIServer) debug(msg string, args ...any) {
	if !s.isDebugLogging {
		return
	}
	s.info("debug: "+msg, args...)
}

// info sends a standard "info string" message to the GUI.
func (s *UCIServer) info(msg string, args ...any) {
	s.WriteString("info string "+msg, args...)
}

// error sends an error message formatted as an "info string" to the GUI.
func (s *UCIServer) error(msg string, args ...any) {
	s.info("error: "+msg, args...)
}

// handleUCI responds to the "uci" command by identifying the engine and
// confirming it is ready to use the UCI protocol.
func (s *UCIServer) handleUCI() {
	s.WriteString("id name %s", BotName)
	s.WriteString("id author %s", Author)
	s.WriteString("uciok")
}

// handleUCINewGame responds to the "ucinewgame" command by resetting the
// board to the starting position.
func (s *UCIServer) handleUCINewGame() {
	s.resetPosition()
}

// handlePosition responds to the "position" command, which sets up the board
// state. It supports both "startpos" and custom "fen" strings, followed by
// an optional list of "moves" to apply.
func (s *UCIServer) handlePosition(args []string) {
	if len(args) < 1 {
		s.error("position command requires at least 2 arguments")
		return
	}

	switch args[0] {
	case "startpos":
		s.resetPosition()
		if len(args) > 1 {
			args = args[1:]
		} else {
			args = nil
		}
	case "fen":
		var i int
		for i = 1; i < len(args); i++ {
			if args[i] == "moves" {
				break
			}
		}

		fen := strings.Join(args[1:i], " ")
		pos, err := chester.ParseFEN(fen)
		if err != nil {
			s.error("error parsing fen: %s", err)
			return
		}
		s.pos = pos
		args = args[i:]
	default:
		s.error("unknown position argument: %s", args[1])
	}

	s.debug("args: %v", args)

	if len(args) > 0 && args[0] == "moves" {
		for _, m := range args[1:] {
			m = strings.TrimSpace(m)
			if m != "" {
				move, err := chester.ParseMove(m, s.pos)
				if err != nil {
					s.error("error parsing move: %s", err)
					return
				}
				s.pos.Do(move)
				s.debug("position: %s", s.pos.String())
			}
		}
	}
}

// handleGo responds to the "go" command, which initiates a search for the
// best move. It parses search constraints like time, depth, and node limits
// from the provided arguments.
func (s *UCIServer) handleGo(args []string) {
	if len(args) < 1 {
		s.error("go command requires at least 2 arguments")
		return
	}

	opts := &chester.SearchOptions{
		MaxDepth:           100,
		MaxNodes:           math.MaxInt64,
		TranspositionTable: s.tt,
	}

	var wtime, btime, winc, binc, movestogo, movetime int64

	// parse go arguments
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "depth":
			i++
			fmt.Sscanf(args[i], "%d", &opts.MaxDepth)
		case "nodes":
			i++
			fmt.Sscanf(args[i], "%d", &opts.MaxNodes)
		case "movetime":
			i++
			fmt.Sscanf(args[i], "%d", &movetime)
			opts.MaxTime = time.Duration(movetime) * time.Millisecond
		case "wtime":
			i++
			fmt.Sscanf(args[i], "%d", &wtime)
		case "btime":
			i++
			fmt.Sscanf(args[i], "%d", &btime)
		case "winc":
			i++
			fmt.Sscanf(args[i], "%d", &winc)
		case "binc":
			i++
			fmt.Sscanf(args[i], "%d", &binc)
		case "movestogo":
			i++
			fmt.Sscanf(args[i], "%d", &movestogo)
		case "infinite":
			opts.MaxTime = 0
			opts.MaxDepth = 100
		}
	}

	if opts.MaxTime == 0 && (wtime > 0 || btime > 0) {
		opts.MaxTime = calculateTimeLimit(s.pos.Active(), wtime, btime, winc, binc, movestogo)
	}

	go func() {
		pos := *s.pos
		ch, stopFunc := chester.SearchBestMove(&pos, opts)
		s.stopFunc = stopFunc
		for e := range ch {
			s.info("depth %d score cp %d pv %s", e.Depth, e.Score, e.Best)
			s.bestMove = e.Best.String()
		}

		s.stopFunc = nil
		s.WriteString("bestmove %s", s.bestMove)
	}()
}

// handleIsReady responds to the "isready" command, signaling to the GUI
// that the engine is ready for further commands.
func (s *UCIServer) handleIsReady() {
	s.WriteString("readyok")
}

// handleStop responds to the "stop" command by immediately aborting any
// ongoing search.
func (s *UCIServer) handleStop() {
	if s.stopFunc != nil {
		s.stopFunc()
		s.stopFunc = nil
	}
}

// handlePerft handles the "perft" command, which runs a performance test
// at a specified depth to count the number of nodes in the move tree.
func (s *UCIServer) handlePerft(args []string) {
	depth := 6

	if len(args) != 0 {
		if d, err := strconv.Atoi(args[0]); err == nil {
			depth = d
		} else {
			s.error("error parsing depth: %s", err)
			return
		}
	}

	go func() {
		nodes := 0
		start := time.Now()
		pos := *s.pos
		ch := chester.Perft(&pos, depth)
		for m := range ch {
			nodes += m.Count
			s.WriteString("%s: %d", m.Move, m.Count)
		}
		duration := time.Since(start)
		nps := float32(nodes) / float32(duration.Seconds())

		s.WriteString("\nNodes: %d", nodes)
		s.WriteString("Time: %s", duration.Round(time.Millisecond))
		s.WriteString("NPS: %s\n", formatNPS(nps))
	}()
}

// handleCPUProfile toggles CPU profiling for performance analysis.
func (s *UCIServer) handleCPUProfile(args []string) {
	if s.isCPUProfiling {
		s.info("cpu profiling stopped")
		pprof.StopCPUProfile()
		return
	}

	filename := "default.pgo"
	if len(args) != 0 {
		filename = args[0]
	}

	file, err := os.Create(filename)
	if err != nil {
		s.error("error creating profile file: %s", err)
		return
	}

	s.info("cpu profiling started")
	pprof.StartCPUProfile(file)
	s.CPUProfileFile = file
	s.isCPUProfiling = true
}

// handleDebug toggles detailed debug logging to the GUI.
func (s *UCIServer) handleDebug(args []string) {
	if len(args) == 0 {
		s.isDebugLogging = !s.isDebugLogging
		return
	}

	if args[0] == "off" && s.isDebugLogging {
		s.isDebugLogging = false
	}

	if args[0] == "on" && !s.isDebugLogging {
		s.isDebugLogging = true
	}
}

// resetPosition resets the board to the standard starting FEN.
func (s *UCIServer) resetPosition() {
	pos, err := chester.ParseFEN(chester.DefaultFEN)
	if err != nil {
		s.error("error parsing fen: %s", err)
	}
	s.pos = pos
}

// calculateTimeLimit determines a reasonable maximum duration for a move
// based on remaining time, increments, and moves to go.
func calculateTimeLimit(color chester.Color, wtime, btime, winc, binc, movestogo int64) time.Duration {
	var timeLeft, inc int64

	if color == chester.White {
		timeLeft, inc = wtime, winc
	} else {
		timeLeft, inc = btime, binc
	}

	if movestogo <= 0 {
		movestogo = 30
	}

	targetMs := (timeLeft / movestogo) + int64(float64(inc)*0.8)

	if targetMs > timeLeft/2 {
		targetMs = timeLeft / 2
	}

	return time.Duration(targetMs) * time.Millisecond
}

// formatNPS returns a human-readable string representation of nodes per second.
func formatNPS(nps float32) string {
	switch {
	case nps >= 1_000_000:
		return fmt.Sprintf("%.2f MNPS", nps/1_000_000)
	case nps >= 1_000:
		return fmt.Sprintf("%.1f kNPS", nps/1_000)
	default:
		return fmt.Sprintf("%.0f NPS", nps)
	}
}
