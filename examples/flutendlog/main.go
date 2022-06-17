package main

import (
	"github.com/wwqdrh/logger"
	"go.uber.org/zap/zapcore"
)

// if fluentd is down, the log will lock
func main() {
	// info级别 with name
	l := logger.NewLogger(
		logger.WithLevel(zapcore.WarnLevel),
		logger.WithLogPath("./logs/info.log"),
		logger.WithName("info"),
		logger.WithFluentd(true, "127.0.0.1", 3306))
	l.Warn("this is a debug message")
	l.Info("this is a info message")
	logger.Get("info").Warn("this is a debug message")
	logger.Get("info").Info("this is a info message")
}
