package main

import (
	"github.com/wwqdrh/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Default Logger
	logger.DefaultLogger.Warn("this is a debug message")
	logger.DefaultLogger.Info("this is a info message")

	// info级别 with name
	l := logger.NewLogger(logger.WithLevel(zapcore.WarnLevel), logger.WithLogPath("./logs/info.log"), logger.WithName("info"))
	l.Warn("this is a debug message")
	l.Info("this is a info message")
	logger.Get("info").Warn("this is a debug message")
	logger.Get("info").Info("this is a info message")

	// switch debug level
	l = logger.NewLogger(logger.WithLevel(zapcore.DebugLevel), logger.WithLogPath("./logs/info2.log"), logger.WithColor(true))
	l.Debug("this is a debug message")
	l.Info("this is a info message")
	l.Warn("this is a warn message")
}
