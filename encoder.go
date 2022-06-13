package logger

import (
	"bytes"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var (
	blue   = []byte("\033[36m")
	red    = []byte("\033[31m")
	yellow = []byte("\033[33m")
)

// encoder

type ColorJsonEncoder struct {
	zapcore.Encoder

	config zapcore.EncoderConfig
}

type ColorConsoleEncoder struct {
	zapcore.Encoder

	config zapcore.EncoderConfig
}

func withColorRender(level zapcore.Level, buf *buffer.Buffer) *buffer.Buffer {
	buffer := new(bytes.Buffer)
	if level >= zap.PanicLevel {
		buffer.Write(red)
		buffer.Write(buf.Bytes())
		buffer.Write(red)
	} else if level >= zap.WarnLevel {
		buffer.Write(yellow)
		buffer.Write(buf.Bytes())
		buffer.Write(yellow)
	} else {
		buffer.Write(blue)
		buffer.Write(buf.Bytes())
		buffer.Write(blue)
	}

	buf.Reset()
	buf.Write(buffer.Bytes())
	return buf
}

func NewColorJsonEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	return ColorJsonEncoder{
		Encoder: zapcore.NewJSONEncoder(config),
		config:  config,
	}
}

func NewColorConsoleEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	return ColorJsonEncoder{
		Encoder: zapcore.NewConsoleEncoder(config),
		config:  config,
	}
}

func (enc ColorJsonEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	return withColorRender(entry.Level, buf), nil
}

func (enc ColorConsoleEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	return withColorRender(entry.Level, buf), nil
}
