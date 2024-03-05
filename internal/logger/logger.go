package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	Logger *slog.Logger
}

func New() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &Logger{Logger: logger}
}
func (l *Logger) New() *Logger {
	return New()
}
