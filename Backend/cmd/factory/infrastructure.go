package factory

import (
	"github.com/qrave1/logwrap"
	"github.com/qrave1/logwrap/logrus"
)

func provideLogger() logwrap.Logger {
	return logrus.NewDefaultPrettyWrapper()
}
