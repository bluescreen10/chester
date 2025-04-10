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
}

func startUCI(logger *slog.Logger) {
	pos, _ := Parse(DefaultFEN)
	uci := &UCIServer{logger: logger, pos: pos, bestMove: Move{}}
	uci.Start()
}

func (s *UCIServer) Start() {
	s.logger.Info("Starting UCI server...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		s.logger.Debug(fmt.Sprintf("Received: %s", line))

		args := strings.Split(line, " ")
		s.logger.Debug(fmt.Sprintf("Command: %s Args: %v", args[0], args[1:]))

		switch args[0] {
		case "quit":
			s.logger.Info("Quitting UCI server...")
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
		default:
			s.logger.Error(fmt.Sprintf("Unknown command: %s", args[0]))
		}
	}
}

func (s *UCIServer) Write(msg []byte) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.logger.Debug(fmt.Sprintf("Sending: %s", msg))
	n, err := os.Stdout.Write(msg)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error writing to stdout: %s", err))
	}
	return n, err
}

func (s *UCIServer) handleUCI() {
	fmt.Fprintf(s, "id name Chester")
	fmt.Fprintf(s, "id author Mariano Wahlmann")
	fmt.Fprintf(s, "uciok")
}

func (s *UCIServer) handleUCINewGame() {
	s.resetPosition()
}

func (s *UCIServer) handlePosition(args []string) {
	if len(args) < 2 {
		s.logger.Error("Position command requires at least 2 arguments")
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
			s.logger.Error(fmt.Sprintf("Error parsing FEN: %s", err))
			return
		}
		s.pos = pos
		s.bestMove = Move{}
		args = args[6:]
	default:
		s.logger.Error(fmt.Sprintf("Unknown position argument: %s", args[1]))
	}

	p := &s.pos

	s.logger.Debug(fmt.Sprintf("Args: %v", args))

	if args[0] == "moves" {
		for _, m := range args[1:] {
			move, err := ParseMove(m, *p)
			if err != nil {
				s.logger.Error(fmt.Sprintf("Error parsing move: %s", err))
				s.logger.Error(s.pos.String())
				return
			}
			p.Do(move)
			s.logger.Debug(fmt.Sprintf("Position: %s", s.pos.String()))
		}
	}
	s.pos = *p

	s.bestMove = Move{}
}

func (s *UCIServer) handleGo(args []string) {
	if len(args) < 2 {
		s.logger.Error("Go command requires at least 2 arguments")
		return
	}

	var moves []Move
	LegalMoves(&moves, &s.pos)

	if len(moves) == 0 {
		s.logger.Error("No legal moves available")
		return
	}

	fmt.Printf("info depth 1 pv %s\n", moves[0])

	s.logger.Debug(fmt.Sprintf("Legal moves: %v", moves))
	s.bestMove = moves[0]
}

func (s *UCIServer) handleIsReady() {
	fmt.Fprintf(s, "readyok")
}

func (s *UCIServer) handleStop() {
	if s.bestMove != (Move{}) {
		fmt.Fprintf(s, "bestmove %s\n", s.bestMove)
	}
}

func (s *UCIServer) handlePerft(args []string) {
	depth := 6

	if len(args) != 0 {
		if d, err := strconv.Atoi(args[0]); err == nil {
			depth = d
		} else {
			s.logger.Error(fmt.Sprintf("Error parsing depth: %s", err))
			return
		}
	}

	go func() {
		start := time.Now()
		output := bytes.Buffer{}
		nodes := Perft(s.pos, depth, s)
		duration := time.Since(start)
		fmt.Fprintf(s, "%s\nperft %d in %s\n", output.String(), nodes, duration)
	}()
}

func (s *UCIServer) handleCPUProfile(args []string) {
	if s.isCPUProfiling {
		fmt.Fprintf(s, "CPU profiling stopped\n")
		pprof.StopCPUProfile()
		return
	}

	filename := "default.pgo"
	if len(args) != 0 {
		filename = args[0]
	}

	file, err := os.Create(filename)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error creating profile file: %s", err))
		return
	}

	fmt.Fprintf(s, "CPU profiling started\n")
	pprof.StartCPUProfile(file)
	s.CPUProfileFile = file
	s.isCPUProfiling = true
}

func (s *UCIServer) resetPosition() {
	pos, err := Parse(DefaultFEN)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error parsing FEN: %s", err))
	}
	s.pos = pos
	s.bestMove = Move{}
}
