package main

import (
	"github.com/wwqdrh/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	// info级别
	l := logger.NewLogger(logger.WithLevel(zapcore.WarnLevel), logger.WithLogPath("./logs/info.log"))
	l.Warn("this is a debug message")
	l.Info("this is a info message")

	// switch debug level
	l = logger.NewLogger(logger.WithLevel(zapcore.DebugLevel), logger.WithLogPath("./logs/info2.log"), logger.WithColor(true))
	l.Debug("this is a debug message")
	l.Info("this is a info message")
	l.Warn("this is a warn message")
}
