package factory

import (
	"github.com/qrave1/quicknotes/pkg/logger"
	"log/slog"
	"os"
)

func provideLogger() *logger.Logger {
	return &logger.Logger{
		Logger: slog.New(
			slog.NewTextHandler(os.Stdout, nil),
		),
	}
}
