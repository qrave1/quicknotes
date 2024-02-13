package logger

import (
	"log/slog"
)

type Logger struct {
	*slog.Logger
}
