package logger

import (
	"log/slog"
	"os"

	"github.com/dusted-go/logging/prettylog"
)

var Log *slog.Logger

func init() {
	logLevel := os.Getenv("LOG_LEVEL")

	var level slog.Level
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
		Level:       level,
		AddSource:   true,
		ReplaceAttr: nil,
	})

	Log = slog.New(prettyHandler)
}
