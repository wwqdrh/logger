package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	basicJsonEncoder = zapcore.NewJSONEncoder(encoderConfig)
	colorJsonEncoder = NewColorJsonEncoder(encoderConfig)
)

// 输出到日志中的不加颜色
// 控制台中的根据color属性判断
func NewLogger(options ...option) *zap.Logger {
	opt := NewLoggerOption()
	for _, item := range options {
		item(opt)
	}

	var coreArr []zapcore.Core

	priority := getPriority(opt.Level)
	if opt.LogPath != "" {
		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   opt.LogPath,       //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    opt.LogMaxSize,    //文件大小限制,单位MB
			MaxBackups: opt.LogMaxBackups, //最大保留日志文件数量
			MaxAge:     opt.LogMaxAge,     //日志文件保留天数
			Compress:   opt.LogCompress,   //是否压缩处理
		})
		coreArr = append(coreArr, zapcore.NewCore(basicJsonEncoder, fileWriteSyncer, priority))
	}

	if opt.Color {
		coreArr = append(coreArr, zapcore.NewCore(colorJsonEncoder, zapcore.AddSync(os.Stdout), priority))
	} else {
		coreArr = append(coreArr, zapcore.NewCore(basicJsonEncoder, zapcore.AddSync(os.Stdout), priority))
	}

	return zap.New(zapcore.NewTee(coreArr...))
}
