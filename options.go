package logger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type LoggerOptions struct {
	Name       string
	Level      zapcore.Level
	Color      bool
	Console    bool // 如非必要不输出到控制台，例如开启fluentd就不需要输出，除非是fluentd失败
	Switch     bool // 是否支持动态修改等级
	SwitchTime time.Duration

	// encoder config
	EncoderOut   string // json plain
	EncoderLevel string
	EncoderTime  string

	LogPath       string // 保存的日志文件
	LogMaxSize    int    // 文件大小限制
	LogMaxBackups int    //最大保留日志文件数量
	LogMaxAge     int    //日志文件保留天数
	LogCompress   bool   //是否压缩处理

	// flutend config
	FlutendEnable bool // 是否上报给fluentd
	FlutendHost   string
	FlutendPort   int
}

type option func(*LoggerOptions)

func NewLoggerOption() *LoggerOptions {
	return &LoggerOptions{
		Level:         zapcore.InfoLevel,
		Color:         true,
		Console:       true,
		Switch:        true,
		SwitchTime:    5 * time.Minute,
		EncoderOut:    "json",
		EncoderLevel:  "level",
		EncoderTime:   "time",
		LogPath:       "",
		LogMaxSize:    1,
		LogMaxBackups: 5,
		LogMaxAge:     30,
		LogCompress:   false,
	}
}

func WithName(name string) option {
	return func(lo *LoggerOptions) {
		lo.Name = name
	}
}

func WithLevel(level zapcore.Level) option {
	return func(lo *LoggerOptions) {
		lo.Level = level
	}
}

func WithSwitch(isSwitch bool) option {
	return func(lo *LoggerOptions) {
		lo.Switch = isSwitch
	}
}

func WithSwitchTime(switchTime time.Duration) option {
	return func(lo *LoggerOptions) {
		lo.SwitchTime = switchTime
	}
}

func WithLogPath(logPath string) option {
	return func(lo *LoggerOptions) {
		lo.LogPath = logPath
	}
}

func WithColor(color bool) option {
	return func(lo *LoggerOptions) {
		lo.Color = color
	}
}

func WithEncoderTime(timeKey string) option {
	return func(lo *LoggerOptions) {
		lo.EncoderTime = timeKey
	}
}

func WithEncoderLevel(levelKey string) option {
	return func(lo *LoggerOptions) {
		lo.EncoderLevel = levelKey
	}
}

func WithEncoderOut(out string) option {
	return func(lo *LoggerOptions) {
		lo.EncoderOut = out
	}
}

func WithFluentd(enable bool, host string, port int) option {
	return func(lo *LoggerOptions) {
		lo.FlutendEnable = enable
		lo.FlutendHost = host
		lo.FlutendPort = port
	}
}

func WithConsole(enable bool) option {
	return func(lo *LoggerOptions) {
		lo.Console = enable
	}
}
