package cronlog

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	zapLogger *zap.SugaredLogger
}

// NewLogger creates a logger which implements `github.com/robfig/cron.Logger`
func NewLogger(zapLogger *zap.SugaredLogger) Logger {
	return Logger{zapLogger: zapLogger}
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.zapLogger.Infow(msg, args...)
}

func (l Logger) Error(err error, msg string, args ...interface{}) {
	l.zapLogger.Errorw(fmt.Sprintf(msg, args...), "error", err.Error())
}

func (l Logger) Flush() {
	_ = l.zapLogger.Sync()
}
