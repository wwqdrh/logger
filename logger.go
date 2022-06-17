package logger

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerPool = sync.Map{}

var (
	DefaultLogger *zap.Logger
)

func init() {
	DefaultLogger = NewLogger(
		WithLevel(zap.InfoLevel),
		WithEncoderLevel(""),
		WithEncoderTime(""),
		WithEncoderOut("plain"),
	)
}

// 输出到日志中的不加颜色
// 控制台中的根据color属性判断
func NewLogger(options ...option) *zap.Logger {
	opt := NewLoggerOption()
	for _, item := range options {
		item(opt)
	}

	// encoder
	config := encoderConfig
	config.LevelKey = opt.EncoderLevel
	config.TimeKey = opt.EncoderTime

	var (
		basicEncoder zapcore.Encoder
		colorEncoder zapcore.Encoder
	)
	if opt.EncoderOut == "json" {
		basicEncoder = zapcore.NewJSONEncoder(config)
		colorEncoder = NewColorJsonEncoder(config)
	} else {
		basicEncoder = zapcore.NewConsoleEncoder(config)
		colorEncoder = NewColorConsoleEncoder(config)
	}

	// 构造zap
	var coreArr []zapcore.Core
	priority := getPriority(opt.Level)
	// 是否保存到文件中
	if opt.LogPath != "" {
		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   opt.LogPath,       //日志文件存放目录，如果文件夹不存在会自动创建
			MaxSize:    opt.LogMaxSize,    //文件大小限制,单位MB
			MaxBackups: opt.LogMaxBackups, //最大保留日志文件数量
			MaxAge:     opt.LogMaxAge,     //日志文件保留天数
			Compress:   opt.LogCompress,   //是否压缩处理
		})
		coreArr = append(coreArr, zapcore.NewCore(basicEncoder, fileWriteSyncer, priority))
	}
	// 是否有flutend
	if opt.FlutendEnable {
		coreArr = append(coreArr,
			zapcore.NewCore(NewBasicJsonFlutendEncoder(config, opt.FlutendHost, opt.FlutendPort), zapcore.AddSync(os.Stdout), priority),
		)
	}
	// 控制台输出是否添加颜色
	if opt.Color {
		coreArr = append(coreArr, zapcore.NewCore(colorEncoder, zapcore.AddSync(os.Stdout), priority))
	} else {
		coreArr = append(coreArr, zapcore.NewCore(basicEncoder, zapcore.AddSync(os.Stdout), priority))
	}

	l := zap.New(zapcore.NewTee(coreArr...))
	if opt.Name != "" {
		loggerPool.Store(opt.Name, l)
	}
	return l
}

func Set(name string, Logger *zap.Logger) {
	loggerPool.Store(name, Logger)
}

func Get(name string) *zap.Logger {
	val, ok := loggerPool.Load(name)
	if !ok {
		return DefaultLogger
	}

	if v, ok := val.(*zap.Logger); !ok {
		return DefaultLogger
	} else {
		return v
	}
}
