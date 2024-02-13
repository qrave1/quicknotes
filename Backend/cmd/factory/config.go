package factory

import (
	"github.com/qrave1/quicknotes/internal/config"
)

func provideConfig() *config.Config {
	return config.MustLoad()
}
