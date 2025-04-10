package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"strings"
)

const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func main() {
	filename := flag.String("logfile", "uci.log", "log file")
	level := flag.String("loglevel", "warn", "log level")
	flag.Parse()

	var logfile io.Writer
	if *filename == "-" {
		logfile = os.Stderr
	} else {
		if *filename == "" {
			*filename = "uci.log"
		}
		logfile, err := os.OpenFile("uci.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer logfile.Close()
	}

	var logLevel slog.Level

	switch strings.ToLower(*level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelWarn
	}

	logger := slog.New(slog.NewTextHandler(logfile, &slog.HandlerOptions{Level: logLevel}))
	startUCI(logger)
}
