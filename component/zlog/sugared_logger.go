package zlog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	SugaredLogger *zap.SugaredLogger
	ZapLogger     *zap.Logger
)

func GetLogger() *zap.SugaredLogger {
	if SugaredLogger == nil {
		if ZapLogger == nil {
			ZapLogger = GetZapLogger()
		}
		SugaredLogger = ZapLogger.Sugar()
	}
	return SugaredLogger
}

func sugaredLogger(ctx *gin.Context) *zap.SugaredLogger {
	if ctx == nil {
		return SugaredLogger
	}
	s := SugaredLogger.With(
		zap.String(ContextKeyLogID, GetLogId(ctx)),
		zap.String(ContextKeyIp, ctx.GetString(ContextKeyIp)),
		zap.String(ContextKeyUri, ctx.GetString(ContextKeyUri)),
	)
	return s
}

func Debug(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Debug(args...)
}

func Debugf(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Debugf(format, args...)
}

func Info(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Info(args...)
}

func Infof(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Infof(format, args...)
}

func Warn(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Warn(args...)
}

func Warnf(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Warnf(format, args...)
}

func Error(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Error(args...)
}

func Errorf(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Errorf(format, args...)
}

func Panic(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Panic(args...)
}

func Panicf(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Panicf(format, args...)
}

func Fatal(ctx *gin.Context, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Fatal(args...)
}

func Fatalf(ctx *gin.Context, format string, args ...interface{}) {
	if nilCtx(ctx) {
		return
	}
	sugaredLogger(ctx).Fatalf(format, args...)
}

func nilCtx(ctx *gin.Context) bool {
	return ctx == nil
}
