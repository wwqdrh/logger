package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(options ...option) *zap.Logger {
	opt := NewLoggerOption()
	for _, item := range options {
		item(opt)
	}

	var coreArr []zapcore.Core
	var encoder zapcore.Encoder
	if opt.Color {
		encoder = NewColorJsonEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	priority := getPriority(opt.Level)
	if opt.LogPath != "" {
		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   opt.LogPath,       //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    opt.LogMaxSize,    //文件大小限制,单位MB
			MaxBackups: opt.LogMaxBackups, //最大保留日志文件数量
			MaxAge:     opt.LogMaxAge,     //日志文件保留天数
			Compress:   opt.LogCompress,   //是否压缩处理
		})
		coreArr = append(coreArr, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), priority)) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	} else {
		coreArr = append(coreArr, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), priority))
	}

	return zap.New(zapcore.NewTee(coreArr...))
}
