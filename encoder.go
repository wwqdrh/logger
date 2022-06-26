package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// 包含两种属性: 颜色、输出格式
// 输出位置不是互斥，是可以组合的(由zapcore构造的时候指定)

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
}

type BasicJsonFlutendEncoder struct {
	zapcore.Encoder

	config zapcore.EncoderConfig
	client *fluent.Fluent
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
	_, err := buf.Write(buffer.Bytes())
	if err != nil {
		return nil
	}
	return buf
}

func NewColorJsonEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	return ColorJsonEncoder{
		Encoder: zapcore.NewJSONEncoder(config),
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

func NewColorConsoleEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	return ColorJsonEncoder{
		Encoder: zapcore.NewConsoleEncoder(config),
		config:  config,
	}
}

func (enc ColorConsoleEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	return withColorRender(entry.Level, buf), nil
}

func NewBasicJsonFlutendEncoder(config zapcore.EncoderConfig, host string, port int) zapcore.Encoder {
	client, err := fluent.New(fluent.Config{
		FluentHost: host,
		FluentPort: port,
	})
	if err != nil {
		return BasicJsonFlutendEncoder{
			Encoder: zapcore.NewJSONEncoder(config),
			config:  config,
			client:  nil,
		}
	} else {
		return BasicJsonFlutendEncoder{
			Encoder: zapcore.NewJSONEncoder(config),
			config:  config,
			client:  client,
		}
	}
}

func (enc BasicJsonFlutendEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	// message must be a map
	if enc.client != nil {
		message := map[string]interface{}{}
		if err := json.Unmarshal(buf.Bytes(), &message); err != nil {
			return buf, err
		}
		e := enc.client.Post(entry.Level.String(), message)
		buf.Reset()
		if e != nil {
			if _, err := buf.Write([]byte(fmt.Sprintf("[fluentd error] %s\n", e.Error()))); err != nil {
				return buf, err
			}
		}
	} else {
		old := buf.String()
		buf.Reset()
		if _, err := buf.Write([]byte("[nil fluentd]: ")); err != nil {
			return buf, err
		}
		if _, err := buf.WriteString(old); err != nil {
			return buf, err
		}
		buf = withColorRender(entry.Level, buf)
	}

	return buf, err
}
