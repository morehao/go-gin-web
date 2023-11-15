package zlog

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetLogId(ctx *gin.Context) (logId string) {
	if ctx == nil {
		return genLogID()
	}
	// 如果上下文有logId，则直接返回
	if logID := ctx.GetString(ContextKeyLogID); logID != "" {
		return logID
	}
	// 从header中获取
	var logID string
	if ctx.Request != nil && ctx.Request.Header != nil {
		logID = ctx.GetHeader(HeaderKeyLogID)
		if logID == "" {
			logID = ctx.GetHeader(HeaderKeyLowerLogID)
		}
	}

	if logID == "" {
		logID = genLogID()
	}

	ctx.Set(ContextKeyLogID, logID)
	return logID
}

func genLogID() (logId string) {
	// 生成纳秒时间戳
	nanosecond := uint64(time.Now().UnixNano())
	// nanosecond&0x7FFFFFFF 使用位运算与操作，将 nanosecond 的二进制表示的最高位（最高位是符号位）清零，将其转换为正整数。
	// |0x80000000 使用位运算或操作，将二进制表示的最高位设置为 1，以确保结果是一个正整数。这样做的目的是为了确保结果是正数，而不是负数。
	logId = strconv.FormatUint(nanosecond&0x7FFFFFFF|0x80000000, 10)
	return logId
}

func GetFormatRequestTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05.999999")
}

func GetRequestCost(start, end time.Time) float64 {
	return float64(end.Sub(start).Nanoseconds()/1e4) / 100.0
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getLogLevel(level string) zapcore.Level {
	logLevel := strings.ToLower(level)
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
