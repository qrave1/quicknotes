package factory

import "github.com/qrave1/logger-wrapper/logrus"

func provideLogger() logrus.Logger {
	return logrus.NewDefaultLogrusWrapper()
}
