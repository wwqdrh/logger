package logx

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"go.uber.org/zap"
)

// TODO: e := logger.Post(tag, data) ,应该是只能修改encoder才行了
type ZapWithFlutend struct {
	*zap.Logger
	fluentClient *fluent.Fluent
}

func NewZapWithFlutend(logger *zap.Logger, fluentdHost string, fluentdPort int) (*ZapWithFlutend, error) {
	client, err := fluent.New(fluent.Config{FluentPort: fluentdPort, FluentHost: fluentdHost})
	if err != nil {
		return nil, err
	}
	return &ZapWithFlutend{
		Logger:       logger,
		fluentClient: client,
	}, nil
}

func (l *ZapWithFlutend) Close() error {
	return l.fluentClient.Close()
}

func (log *ZapWithFlutend) Debug(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) Info(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) Warn(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) Error(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) DPanic(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.DPanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) Panic(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *ZapWithFlutend) Fatal(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}
