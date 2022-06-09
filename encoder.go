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

func NewColorJsonEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	if config.ConsoleSeparator == "" {
		// Use a default delimiter of '\t' for backwards compatibility
		config.ConsoleSeparator = "\t"
	}
	return ColorJsonEncoder{
		Encoder: zapcore.NewJSONEncoder(config),
		config:  config,
	}
}

func (enc ColorJsonEncoder) Clone() zapcore.Encoder {
	return ColorJsonEncoder{
		Encoder: zapcore.NewJSONEncoder(enc.config),
		config:  enc.config,
	}
}

func (enc ColorJsonEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	buffer := new(bytes.Buffer)
	level := entry.Level
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
	return buf, nil
}
