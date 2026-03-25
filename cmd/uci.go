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

func startUCI() {
	pos, _ := chester.ParseFEN(chester.DefaultFEN)
	uci := &UCIServer{pos: pos, tt: chester.NewTranspositionTable(64 * 1024 * 1024)}
	uci.Start()
}

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

func (s *UCIServer) Write(msg []byte) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return os.Stdout.Write(msg)
}

func (s *UCIServer) WriteString(msg string, args ...any) {
	s.Write([]byte(fmt.Sprintf(msg+"\n", args...)))
}

func (s *UCIServer) debug(msg string, args ...any) {
	if !s.isDebugLogging {
		return
	}
	s.info("debug: "+msg, args...)
}

func (s *UCIServer) info(msg string, args ...any) {
	s.WriteString("info string "+msg, args...)
}

func (s *UCIServer) error(msg string, args ...any) {
	s.info("error: "+msg, args...)
}

func (s *UCIServer) handleUCI() {
	s.WriteString("id name %s", BotName)
	s.WriteString("id author %s", Author)
	s.WriteString("uciok")
}

func (s *UCIServer) handleUCINewGame() {
	s.resetPosition()
}

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

func (s *UCIServer) handleIsReady() {
	s.WriteString("readyok")
}

func (s *UCIServer) handleStop() {
	if s.stopFunc != nil {
		s.stopFunc()
		s.stopFunc = nil
	}
}

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

func (s *UCIServer) resetPosition() {
	pos, err := chester.ParseFEN(chester.DefaultFEN)
	if err != nil {
		s.error("error parsing fen: %s", err)
	}
	s.pos = pos
}

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
