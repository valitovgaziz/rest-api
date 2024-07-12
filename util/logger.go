package util

import (
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8)
	LevelFatal = slog.Level(12)
)

var Logger *slog.Logger

func SetLogLevel(level string) { // set log level
	programLevel := slog.LevelInfo
	if level == "debug" {
		programLevel = slog.LevelDebug
	}
	if level == "info" {
		programLevel = slog.LevelInfo
	}
	if level == "warn" {
		programLevel = slog.LevelWarn
	}
	if level == "error" {
		programLevel = slog.LevelError
	}
	if level == "trace" {
		programLevel = LevelTrace
	}
	if level == "fatal" {
		programLevel = LevelFatal
	}

	// init logger
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel}) // create handler
	slog.SetDefault(slog.New(h))
}
