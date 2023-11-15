package zlog

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetZapLogger() *zap.Logger {
	if ZapLogger == nil {
		ZapLogger = newLogger().WithOptions(zap.AddCallerSkip(1))
	}
	return ZapLogger
}

func zapLogger(ctx *gin.Context) *zap.Logger {
	m := GetZapLogger()
	if ctx == nil {
		return m
	}

	l := m.With(
		zap.String(ContextKeyLogID, GetLogId(ctx)),
		zap.String(ContextKeyIp, ctx.GetString(ContextKeyIp)),
		zap.String(ContextKeyUri, ctx.GetString(ContextKeyUri)),
	)
	return l
}

func DebugLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Debug(msg, fields...)
}
func InfoLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Info(msg, fields...)
}

func WarnLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Warn(msg, fields...)
}

func ErrorLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Error(msg, fields...)
}

func PanicLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Panic(msg, fields...)
}

func FatalLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if nilCtx(ctx) {
		return
	}
	zapLogger(ctx).Fatal(msg, fields...)
}

func newLogger() *zap.Logger {
	var zapCores []zapcore.Core
	logConfigLevel := getLogLevel(logConfig.Level)
	var infoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfigLevel && lvl <= zapcore.InfoLevel
	})

	var errorLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfigLevel && lvl >= zapcore.WarnLevel
	})

	var stdLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfigLevel && lvl >= zapcore.DebugLevel
	})
	if logConfig.InConsole {
		c := zapcore.NewCore(
			getEncoder(),
			getLogWriter(logConfig.AppName, LogFileLevelStdout),
			stdLevel)
		zapCores = append(zapCores, c)
	}

	zapCores = append(zapCores,
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(logConfig.AppName, LogFileLevelNormal),
			infoLevel))
	zapCores = append(zapCores,
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(logConfig.AppName, LogFileLevelNormal),
			errorLevel))
	zapCores = append(zapCores,
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(logConfig.AppName, LogFileLevelWarnFatal),
			errorLevel))
	core := zapcore.NewTee(zapCores...)

	// 开启开发模式，堆栈跟踪
	caller := zap.WithCaller(true)

	// 由于之前没有DPanic，同化DPanic和Panic
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields()

	// 构造logger
	logger := zap.New(core, filed, caller, development)

	return logger
}

func getEncoder() zapcore.Encoder {
	encodeTime := zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999")
	encoderCfg := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     encodeTime,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderCfg)
}

func getColorEncoder() zapcore.Encoder {
	encodeTime := zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999")
	encoderCfg := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     encodeTime,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return zapcore.NewJSONEncoder(encoderCfg)
}

func getLogWriter(name, loggerType string) (ws zapcore.WriteSyncer) {
	var w io.Writer
	if loggerType == LogFileLevelStdout {
		w = os.Stdout
	} else {
		var err error
		director := strings.TrimSuffix(logConfig.Path, "/") + "/" + time.Now().Format("20060102")
		filename := filepath.Join(director, appendLogFileTail(name, loggerType))
		if ok, _ := PathExists(director); !ok { // 判断是否有Director文件夹
			_ = os.Mkdir(director, os.ModePerm)
		}
		w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic("open log file error: " + err.Error())
		}
	}

	flushInterval := 5 * time.Second
	if loggerType == LogFileLevelStdout {
		flushInterval = 1 * time.Second
	}
	ws = &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(w),
		Size:          256 * 1024,
		FlushInterval: flushInterval,
		Clock:         nil,
	}

	return ws
}

func appendLogFileTail(appName, loggerType string) string {
	var tailFixed string
	switch loggerType {
	case LogFileLevelNormal:
		tailFixed = ".log"
	case LogFileLevelWarnFatal:
		tailFixed = "_wf.log"
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
