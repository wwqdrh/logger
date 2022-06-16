package logx

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

// go-zero的适配

type ZeroWriter struct {
	logger *zap.Logger
}

func NewZeroWriter(logger *zap.Logger) (logx.Writer, error) {
	return &ZeroWriter{
		logger: logger,
	}, nil
}

func (w *ZeroWriter) Alert(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *ZeroWriter) Close() error {
	return w.logger.Sync()
}

func (w *ZeroWriter) Error(v interface{}, fields ...logx.LogField) {
	w.logger.Error(fmt.Sprint(v), toZapFields(fields...)...)
}

func (w *ZeroWriter) Info(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)
}

func (w *ZeroWriter) Severe(v interface{}) {
	w.logger.Fatal(fmt.Sprint(v))
}

func (w *ZeroWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.logger.Warn(fmt.Sprint(v), toZapFields(fields...)...)
}

func (w *ZeroWriter) Stack(v interface{}) {
	w.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (w *ZeroWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)
}

func toZapFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
