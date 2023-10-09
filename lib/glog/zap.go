package glog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetZapLogger() (l *zap.Logger) {
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
	if t, exist := ctx.Get(zapLoggerAddr); exist {
		if l, ok := t.(*zap.Logger); ok {
			return l
		}
	}

	l := m.With(
		zap.String("logId", GetLogID(ctx)),
		zap.String("requestId", GetRequestID(ctx)),
		zap.String("module", ctx.GetString("module")),
		zap.String("localIp", ctx.GetString("localIp")),
		zap.String("uri", GetRequestUri(ctx)),
	)
	return l
}

func DebugLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Debug(msg, fields...)
}
func InfoLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Info(msg, fields...)
}

func WarnLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Warn(msg, fields...)
}

func ErrorLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Error(msg, fields...)
}

func PanicLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Panic(msg, fields...)
}

func FatalLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if ctx == nil {
		return
	}
	zapLogger(ctx).Fatal(msg, fields...)
}
