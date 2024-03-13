package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Handler() slog.Handler
}

var DefaultLogger Logger

func init() {
	DefaultLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}
func GetLogger() Logger {
	return DefaultLogger
}
