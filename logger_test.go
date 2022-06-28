package logger

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func TestDefaultLogger(t *testing.T) {
	// Default Logger
	DefaultLogger.Info("this is a info message")
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

func TestLoggerWithFluentd(t *testing.T) {
	// info级别 with name
	l := NewLogger(
		WithLevel(zapcore.WarnLevel),
		WithLogPath("./logs/info.log"),
		WithName("info"),
		WithFluentd(true, "127.0.0.1", 24224),
		WithConsole(false),
	)
	l.Warn("this is a debug message")
	l.Info("this is a info message")
	Get("info").Warn("this is a debug message")
	Get("info").Info("this is a info message")
}
