package logrus

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// NewLogger creates a logrus logger, adds hook to it and return it
func NewLogger(zapLogger *zap.Logger) *logrus.Logger {
	logger := logrus.New()
	logrus.SetOutput(ioutil.Discard)
	logrus.AddHook(newHook(zapLogger))

	return logger
}
