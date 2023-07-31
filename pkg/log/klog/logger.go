package klog

import (
	"flag"

	"go.uber.org/zap"
	"k8s.io/klog"
)

// InitLogger inits klog by zap logger
func InitLogger(zapLogger *zap.Logger) {
	fs := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(fs)
	defer klog.Flush()
	klog.SetOutputBySeverity("INFO", &infoLogger{logger: zapLogger})
	klog.SetOutputBySeverity("WARNING", &warnLogger{logger: zapLogger})
	klog.SetOutputBySeverity("FATAL", &fatalogger{logger: zapLogger})
	klog.SetOutputBySeverity("ERROR", &errorLogger{logger: zapLogger})
	_ = fs.Set("skip_headers", "true")
	_ = fs.Set("logtostderr", "false")
}

type infoLogger struct {
	logger *zap.Logger
}

func (l *infoLogger) Write(p []byte) (n int, err error) {
	l.logger.Info(string(p[:len(p)-1]))

	return len(p), nil
}

type warnLogger struct {
	logger *zap.Logger
}

func (l *warnLogger) Write(p []byte) (n int, err error) {
	l.logger.Warn(string(p[:len(p)-1]))

	return len(p), nil
}

type fatalogger struct {
	logger *zap.Logger
}

func (l *fatalogger) Write(p []byte) (n int, err error) {
	l.logger.Fatal(string(p[:len(p)-1]))

	return len(p), nil
}

type errorLogger struct {
	logger *zap.Logger
}

func (l *errorLogger) Write(p []byte) (n int, err error) {
	l.logger.Error(string(p[:len(p)-1]))

	return len(p), nil
}
