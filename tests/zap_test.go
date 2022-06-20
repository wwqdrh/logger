package tests

import (
	"log"
	"net/http"
	"testing"
	"time"

	"go.uber.org/zap"
)

func printLog(logger *zap.Logger) {
	for i := 0; i < 10; {

		logger.Sugar().Debug("debug log show not diplay")

		logger.Sugar().Error("info log display")

		t, _ := time.ParseDuration("10s")
		time.Sleep(t)

		i++
	}
}

func TestAtomLevel(t *testing.T) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	atomLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	cfg := zap.Config{
		Level:             atomLevel,
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	go printLog(logger)

	http.HandleFunc("/", atomLevel.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
