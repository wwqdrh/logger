package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	encoderConfig zapcore.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		// CallerKey:     "caller", // if switch, hidden
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		// EncodeLevel:    zapcore.CapitalColorLevelEncoder, //这里可以指定颜色,不过只能处理Level的颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
)

var (
	// 日志级别
	debugPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.DebugLevel
	})
	infoPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.InfoLevel
	})
	warnPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.WarnLevel
	})
	errorPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.ErrorLevel
	})
	dPanicPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.DPanicLevel
	})
	panicPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.PanicLevel
	})
	fatalPriority = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev >= zap.FatalLevel
	})
)

// 动态修改
var (
	debugAtomicLevel  = zap.NewAtomicLevelAt(zap.DebugLevel)
	infoAtomicLevel   = zap.NewAtomicLevelAt(zap.InfoLevel)
	warnAtomicLevel   = zap.NewAtomicLevelAt(zap.WarnLevel)
	errorAtomicLevel  = zap.NewAtomicLevelAt(zap.ErrorLevel)
	dPanicAtomicLevel = zap.NewAtomicLevelAt(zap.DPanicLevel)
	panicAtomicLevel  = zap.NewAtomicLevelAt(zap.PanicLevel)
	fatalAtomicLevel  = zap.NewAtomicLevelAt(zap.FatalLevel)
)

func getPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zap.DebugLevel:
		return debugPriority
	case zap.InfoLevel:
		return infoPriority
	case zap.WarnLevel:
		return warnPriority
	case zap.ErrorLevel:
		return errorPriority
	case zap.DPanicLevel:
		return dPanicPriority
	case zap.PanicLevel:
		return panicPriority
	case zap.FatalLevel:
		return fatalPriority
	default:
		return infoPriority
	}
}

func getAtomicPriority(level zapcore.Level) zap.AtomicLevel {
	switch level {
	case zap.DebugLevel:
		return debugAtomicLevel
	case zap.InfoLevel:
		return infoAtomicLevel
	case zap.WarnLevel:
		return warnAtomicLevel
	case zap.ErrorLevel:
		return errorAtomicLevel
	case zap.DPanicLevel:
		return dPanicAtomicLevel
	case zap.PanicLevel:
		return panicAtomicLevel
	case zap.FatalLevel:
		return fatalAtomicLevel
	default:
		return infoAtomicLevel
	}
}
