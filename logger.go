package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerPool      = sync.Map{}
	atomicLevelPool = sync.Map{} // string=>*switchContext
	loggerLevel     = zap.InfoLevel
	loggerLevelMap  = map[string]zapcore.Level{
		"debug":  zap.DebugLevel,
		"info":   zap.InfoLevel,
		"warn":   zap.WarnLevel,
		"error":  zap.ErrorLevel,
		"dpanic": zap.DPanicLevel,
		"panic":  zap.PanicLevel,
		"fatal":  zap.FatalLevel,
	}
)

type switchContext struct {
	l        *zap.AtomicLevel
	duration time.Duration
}

var (
	DefaultLogger *ZapX
)

type ZapX struct {
	*zap.Logger
}

func NewZapX(l *zap.Logger) *ZapX {
	return &ZapX{
		Logger: l,
	}
}

// 从环境变量获取日志级别
func init() {
	logLevel := os.Getenv("LOGLEVEL")
	if val, ok := loggerLevelMap[logLevel]; !ok {
		loggerLevel = zapcore.InfoLevel // 默认info级别
	} else {
		loggerLevel = val
	}

	DefaultLogger = NewLogger(
		WithName("default"),
		WithLevel(loggerLevel),
		WithEncoderLevel(""),
		WithEncoderTime(""),
		WithEncoderOut("plain"),
	)
}

// 设置logger
func Set(name string, Logger *ZapX) {
	loggerPool.Store(name, Logger)
}

// 获取logger
func Get(name string) *ZapX {
	val, ok := loggerPool.Load(name)
	if !ok {
		return DefaultLogger
	}

	if v, ok := val.(*ZapX); !ok {
		return DefaultLogger
	} else {
		return v
	}
}

// 动态切换logger的日志级别
func Switch(name string, level zapcore.Level) {
	val, ok := atomicLevelPool.Load(name)
	if !ok {
		return
	}
	swit := val.(*switchContext)

	// 隔一段时间自动恢复
	go func(l zapcore.Level) {
		<-time.After(swit.duration)
		swit.l.SetLevel(l)
	}(swit.l.Level())

	swit.l.SetLevel(level)
}

// 输出到日志中的不加颜色
// 控制台中的根据color属性判断
// l := NewLogger(
// 	WithLevel(zapcore.DebugLevel),
// 	WithLogPath("./logs/info.log"),
// 	WithName("info"),
// 	WithCaller(false),
// 	WithSwitchTime(2*time.Second),
// )
func NewLogger(options ...option) *ZapX {
	opt := NewLoggerOption()
	for _, item := range options {
		item(opt)
	}

	if val, ok := loggerPool.Load(opt.Name); ok {
		if l, ok := val.(*ZapX); ok {
			return l
		}
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
	// 获取level
	var priority zap.LevelEnablerFunc
	if opt.Switch {
		atomicLevel := getAtomicPriority(opt.Level)
		atomicLevelPool.Store(opt.Name, &switchContext{
			l:        &atomicLevel,
			duration: opt.SwitchTime,
		})
		priority = atomicLevel.Enabled
	} else {
		priority = getPriority(opt.Level)
	}

	// 控制台输出是否添加颜色
	if opt.Console {
		if opt.Color {
			coreArr = append(coreArr, zapcore.NewCore(colorEncoder, zapcore.AddSync(os.Stdout), priority))
		} else {
			coreArr = append(coreArr, zapcore.NewCore(basicEncoder, zapcore.AddSync(os.Stdout), priority))
		}
	}
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

	zapOpts := []zap.Option{}
	if opt.Caller {
		zapOpts = append(zapOpts, zap.AddCaller())
	}
	l := zap.New(zapcore.NewTee(coreArr...), zapOpts...)
	if opt.Name != "" {
		loggerPool.Store(opt.Name, l)
	}
	return &ZapX{
		Logger: l,
	}
}

func (l *ZapX) Debugx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) Infox(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) Warnx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) Errorx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Error(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) DPanicx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.DPanic(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) Panicx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Panic(fmt.Sprintf(format, value...), fields...)
}

func (l *ZapX) Fatalx(format string, fields []zap.Field, value ...interface{}) {
	l.Logger.Fatal(fmt.Sprintf(format, value...), fields...)
}
