package logger

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func TestDefaultLogger(t *testing.T) {
	// Default Logger
	DefaultLogger.Info("this is a info message")
	DefaultLogger.Error("this is a error message")

	DefaultLogger.Infox("this is a %s message", []interface{}{"infox"})
	DefaultLogger.Errorx("this is a %s message", []interface{}{"errorx"})
}

func TestExampleCustomLogger(t *testing.T) {
	// info级别 with name
	l := NewLogger(
		WithLevel(zapcore.DebugLevel),
		WithLogPath("./logs/info.log"),
		WithName("info"),
		WithCaller(false),
		WithSwitchTime(2*time.Second),
	)
	l.Debug("this is a debug message")
	l.Info("this is a info message")
	l.Error("this is error")
	Get("info").Debug("this is a debug message")
	Get("info").Info("this is a info message")

	// switch debug level
	Switch("info", zapcore.WarnLevel)

	l.Debug("this is a debug message")
	l.Info("this is a info message")
	Get("info").Debug("this is a debug message")
	Get("info").Info("this is a info message")

	time.Sleep(3 * time.Second)

	l.Debug("this is a debug message")
	l.Info("this is a info message")
	Get("info").Debug("this is a debug message")
	Get("info").Info("this is a info message")
}
