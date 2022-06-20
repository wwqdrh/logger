package main

import (
	"time"

	"github.com/wwqdrh/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Default Logger
	logger.DefaultLogger.Warn("this is a debug message")
	logger.DefaultLogger.Info("this is a info message")

	// info级别 with name
	l := logger.NewLogger(
		logger.WithLevel(zapcore.DebugLevel),
		logger.WithLogPath("./logs/info.log"),
		logger.WithName("info"),
		logger.WithSwitchTime(5*time.Second),
	)
	l.Debug("this is a debug message")
	l.Info("this is a info message")
	logger.Get("info").Debug("this is a debug message")
	logger.Get("info").Info("this is a info message")

	// switch debug level
	logger.Switch("info", zapcore.WarnLevel)

	l.Debug("this is a debug message")
	l.Info("this is a info message")
	logger.Get("info").Debug("this is a debug message")
	logger.Get("info").Info("this is a info message")

	time.Sleep(7 * time.Second)

	l.Debug("this is a debug message")
	l.Info("this is a info message")
	logger.Get("info").Debug("this is a debug message")
	logger.Get("info").Info("this is a info message")
}
