package logger

import (
	"log/slog"

	"github.com/MinnaSync/proxy/config"
	"github.com/dusted-go/logging/prettylog"
)

var Log *slog.Logger

func init() {
	var level slog.Level
	switch config.LogLevel {
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
