package zlog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	LogFileLevelStdout = "stdout"
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
		zap.String(ContextKeyIp, ctx.ClientIP()),
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
	var zapCore []zapcore.Core
	zapCore = append(zapCore, zapcore.NewCore(
		getEncoder(),
		getLogWriter("server", LogFileLevelStdout),
		zap.InfoLevel))
	core := zapcore.NewTee(zapCore...)

	// 开启堆栈跟踪
	caller := zap.WithCaller(true)

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

func getLogWriter(name, loggerType string) (ws zapcore.WriteSyncer) {
	var w io.Writer
	var err error
	director := strings.TrimSuffix("./log", "/") + "/" + time.Now().Format("20060102")
	filename := filepath.Join(director, name+".log")
	if ok, _ := PathExists(director); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(director, os.ModePerm)
	}
	w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("open log file error: " + err.Error())
	}
	// if loggerType == LogFileLevelStdout {
	// 	w = os.Stdout
	// } else {
	// 	var err error
	// 	filename := filepath.Join(strings.TrimSuffix("./log", "/"), "%Y-%m-%d", name+".log")
	// 	w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// 	if err != nil {
	// 		panic("open log file error: " + err.Error())
	// 	}
	// }
	ws = &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(w),
		Size:          256 * 1024,
		FlushInterval: 5 * time.Second,
		Clock:         nil,
	}
	return ws
}

func CloseLogger() {
	if SugaredLogger != nil {
		_ = SugaredLogger.Sync()
	}

	if ZapLogger != nil {
		_ = ZapLogger.Sync()
	}
}
