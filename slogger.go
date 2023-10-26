package logger

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
)

func New(levelStr string, toJSON bool) *slog.Logger {
	level := strToLevel(levelStr)
	opts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler
	handler = tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
	})
	if toJSON {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func strToLevel(s string) (level slog.Level) {
	switch strings.ToUpper(s) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	}

	return
}
