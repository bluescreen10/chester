package main

import (
	"bufio"
	"context"
	"fmt"
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
	mutex          sync.Mutex
	pos            Position
	bestMove       string
	isCPUProfiling bool
	CPUProfileFile *os.File
	isDebugLogging bool
	stopFunc       func()
}

func startUCI() {
	pos, _ := Parse(DefaultFEN)
	uci := &UCIServer{pos: pos}
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
		pos, err := Parse(fen)
		if err != nil {
			s.error("error parsing fen: %s", err)
			return
		}
		s.pos = pos
		args = args[i:]
	default:
		s.error("unknown position argument: %s", args[1])
	}

	p := &s.pos

	s.debug("args: %v", args)

	if len(args) > 0 && args[0] == "moves" {
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
}

func (s *UCIServer) handleGo(args []string) {
	if len(args) < 1 {
		s.error("go command requires at least 2 arguments")
		return
	}

	ctx, f := context.WithCancel(context.Background())
	s.stopFunc = f

	go func() {
		ch := SearchBestMove(ctx, &s.pos)
		for e := range ch {
			s.info("depth %d score cp %d pv %s", e.depth, e.score, e.best)
			s.bestMove = e.best
		}
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

	if s.bestMove != "" {
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
		nodes := 0
		start := time.Now()
		ch := Perft(&s.pos, depth)
		for m := range ch {
			nodes += m.Count
			s.WriteString("%s: %d", m.Move, m.Count)
		}
		duration := time.Since(start)
		s.WriteString("perft %d in %s\n", nodes, duration)
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
}
