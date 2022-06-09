package logger

import "go.uber.org/zap/zapcore"

type LoggerOptions struct {
	Level         zapcore.Level
	Color         bool
	LogPath       string // 保存的日志文件
	LogMaxSize    int    // 文件大小限制
	LogMaxBackups int    //最大保留日志文件数量
	LogMaxAge     int    //日志文件保留天数
	LogCompress   bool   //是否压缩处理
}

type option func(*LoggerOptions)

func NewLoggerOption() *LoggerOptions {
	return &LoggerOptions{
		Level:         zapcore.InfoLevel,
		Color:         false,
		LogPath:       "",
		LogMaxSize:    1,
		LogMaxBackups: 5,
		LogMaxAge:     30,
		LogCompress:   false,
	}
}

func WithLevel(level zapcore.Level) option {
	return func(lo *LoggerOptions) {
		lo.Level = level
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
