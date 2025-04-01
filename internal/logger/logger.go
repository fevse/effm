package logger

import (
	"log/slog"
	"os"

	"github.com/fevse/effm/internal/config"
)

type Logger struct {
	Logger *slog.Logger
}

func NewLogger(config *config.Config) *Logger {
	logConfig := &slog.HandlerOptions{}
	switch config.LogLevel {
	case "debug":
		logConfig.Level = slog.LevelDebug
	case "info":
		logConfig.Level = slog.LevelInfo
	case "error":
		logConfig.Level = slog.LevelError
	}

	return &Logger{Logger: slog.New(slog.NewTextHandler(os.Stderr, logConfig))}
}

func (l *Logger) Debug(msg string) {
	l.Logger.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.Logger.Info(msg)
}

func (l *Logger) Error(msg string) {
	l.Logger.Error(msg)
}
