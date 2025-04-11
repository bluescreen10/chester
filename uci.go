package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type UCIServer struct {
	logger         *slog.Logger
	mutex          sync.Mutex
	pos            Position
	bestMove       Move
	isCPUProfiling bool
	CPUProfileFile *os.File
	isDebugLogging bool
}

func startUCI() {
	pos, _ := Parse(DefaultFEN)
	uci := &UCIServer{pos: pos, bestMove: Move{}}
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
	if len(args) < 2 {
		s.error("position command requires at least 2 arguments")
		return
	}

	switch args[1] {
	case "startpos":
		s.resetPosition()
		args = args[2:]
	case "fen":
		fen := strings.Join(args[2:6], " ")
		pos, err := Parse(fen)
		if err != nil {
			s.error("error parsing fen: %s", err)
			return
		}
		s.pos = pos
		s.bestMove = Move{}
		args = args[6:]
	default:
		s.error("unknown position argument: %s", args[1])
	}

	p := &s.pos

	s.debug("args: %v", args)

	if args[0] == "moves" {
		for _, m := range args[1:] {
			move, err := ParseMove(m, *p)
			if err != nil {
				s.error("error parsing move: %s", err)
				return
			}
			Do(p, move)
			s.debug("position: %s", s.pos.String())
		}
	}
	s.pos = *p

	s.bestMove = Move{}
}

func (s *UCIServer) handleGo(args []string) {
	if len(args) < 1 {
		s.error("go command requires at least 2 arguments")
		return
	}

	var moves []Move
	LegalMoves(&moves, &s.pos)

	if len(moves) == 0 {
		return
	}

	s.WriteString("info depth 1 pv %s", moves[0])
	s.bestMove = moves[0]
}

func (s *UCIServer) handleIsReady() {
	s.WriteString("readyok")
}

func (s *UCIServer) handleStop() {
	if s.bestMove != (Move{}) {
		s.WriteString("bestmove %s", s.bestMove)
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
		start := time.Now()
		output := bytes.Buffer{}
		nodes := Perft(&s.pos, depth, s)
		duration := time.Since(start)
		fmt.Fprintf(s, "%s\nperft %d in %s\n", output.String(), nodes, duration)
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
	pos, err := Parse(DefaultFEN)
	if err != nil {
		s.error("error parsing fen: %s", err)
	}
	s.pos = pos
	s.bestMove = Move{}
}
