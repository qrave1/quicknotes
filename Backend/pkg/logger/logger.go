package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: slog.New(
			slog.NewTextHandler(os.Stdout, nil),
		),
	}
}
