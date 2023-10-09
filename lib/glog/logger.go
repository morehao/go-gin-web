package glog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type (
	Field  = zap.Field
	Logger = zap.Logger
)

var (
	Array    = zap.Array
	Bools    = zap.Bools
	Ints     = zap.Ints
	Uints    = zap.Uints
	Float64s = zap.Float64s
	Strings  = zap.Strings
	Errors   = zap.Errors

	Binary = zap.Binary
	Bool   = zap.Bool

	ByteString = zap.ByteString
	String     = zap.String

	Float64 = zap.Float64
	Float32 = zap.Float32

	Int   = zap.Int
	Int64 = zap.Int64
	Int32 = zap.Int32
	Int16 = zap.Int16
	Int8  = zap.Int8

	Uint   = zap.Uint
	Uint64 = zap.Uint64
	Uint32 = zap.Uint32
	Uint16 = zap.Uint16
	Uint8  = zap.Uint8

	Reflect       = zap.Reflect
	Namespace     = zap.Namespace
	Duration      = zap.Duration
	Object        = zap.Object
	Any           = zap.Any
	Skip          = zap.Skip()
	AddCallerSkip = zap.AddCallerSkip
)
var (
	SugaredLogger *zap.SugaredLogger
	ZapLogger     *zap.Logger
)

// log文件后缀类型
const (
	txtLogNormal    = "normal"
	txtLogWarnFatal = "warnfatal"
	txtLogStdout    = "stdout"
)

// NewLogger 新建Logger，每一次新建会同时创建x.log与x.log.wf (access.log 不会生成wf)
func newLogger() *zap.Logger {
	var infoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl <= zapcore.InfoLevel
	})

	var errorLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl >= zapcore.WarnLevel
	})

	var stdLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl >= zapcore.DebugLevel
	})

	name := "server"
	var zapCore []zapcore.Core
	if logConfig.Stdout {
		c := zapcore.NewCore(
			getEncoder(),
			getLogWriter(name, txtLogStdout),
			stdLevel)
		zapCore = append(zapCore, c)
	}

	// 仅开发环境有效，便于开发调试
	if logConfig.Log2File {
		zapCore = append(zapCore,
			zapcore.NewCore(
				getEncoder(),
				getLogWriter(name, txtLogNormal),
				infoLevel))

		zapCore = append(zapCore,
			zapcore.NewCore(
				getEncoder(),
				getLogWriter(name, txtLogWarnFatal),
				errorLevel))
	}

	// core
	core := zapcore.NewTee(zapCore...)

	// 开启开发模式，堆栈跟踪
	caller := zap.WithCaller(true)

	// 由于之前没有DPanic，同化DPanic和Panic
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields()

	// 构造日志
	logger := zap.New(core, filed, caller, development)

	return logger
}

func getLogLevel(confLevel string) (level zapcore.Level) {
	levelStr := strings.ToUpper(confLevel)
	switch levelStr {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}

func getEncoder() zapcore.Encoder {
	// time字段编码器
	timeEncoder := zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999")

	encoderCfg := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderCfg)
}

func getLogWriter(name, loggerType string) (ws zapcore.WriteSyncer) {
	var w io.Writer
	if loggerType == txtLogStdout {
		// stdOut
		w = os.Stdout
	} else {
		// 打印到 name.log[.wf] 中
		var err error
		filename := filepath.Join(strings.TrimSuffix(logConfig.Path, "/"), appendLogFileTail(name, loggerType))
		w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic("open log file error: " + err.Error())
		}
	}

	if !logConfig.BufferSwitch {
		return zapcore.AddSync(w)
	}

	// 开启缓冲区
	ws = &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(w),
		Size:          logConfig.BufferSize,
		FlushInterval: logConfig.BufferFlushInterval,
		Clock:         nil,
	}
	return ws
}

// genFilename 拼装完整文件名
func appendLogFileTail(appName, loggerType string) string {
	var tailFixed string
	switch loggerType {
	case txtLogNormal:
		tailFixed = ".log"
	case txtLogWarnFatal:
		tailFixed = ".log.wf"
	default:
		tailFixed = ".log"
	}
	return appName + tailFixed
}

func CloseLogger() {
	if SugaredLogger != nil {
		_ = SugaredLogger.Sync()
	}

	if ZapLogger != nil {
		_ = ZapLogger.Sync()
	}
}
