package zlog

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	ormUtil "gorm.io/gorm/utils"
)

type OrmLogger struct {
	Service  string
	Addr     string
	Database string
	logger   *zap.Logger
}

type DbConfig struct {
	DataBase string `yaml:"database"`
	Addr     string `yaml:"addr"`
}

func NewOrmLog(conf *DbConfig) *OrmLogger {
	s := conf.DataBase

	return &OrmLogger{
		Service:  s,
		Addr:     conf.Addr,
		Database: conf.DataBase,
		logger:   ZapLogger.WithOptions(zap.AddCallerSkip(2)),
	}
}

// go-sql-driver error log
func (l *OrmLogger) Print(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...), l.commonFields(context.Background())...)
}

// LogMode log mode
func (l *OrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l OrmLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	m := fmt.Sprintf(msg, append([]interface{}{ormUtil.FileWithLineNum()}, data...)...)
	// 非trace日志改为debug级别输出
	l.logger.Debug(m, l.commonFields(ctx)...)
}

// Warn print warn messages
func (l OrmLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	m := fmt.Sprintf(msg, append([]interface{}{ormUtil.FileWithLineNum()}, data...)...)
	l.logger.Warn(m, l.commonFields(ctx)...)
}

// Error print error messages
func (l OrmLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	m := fmt.Sprintf(msg, append([]interface{}{ormUtil.FileWithLineNum()}, data...)...)
	l.logger.Error(m, l.commonFields(ctx)...)
}

// Trace print sql message
func (l OrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	end := time.Now()
	elapsed := end.Sub(begin)
	cost := float64(elapsed.Nanoseconds()/1e4) / 100.0

	msg := "mysql do success"
	ralCode := -0
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有找到记录不统计在请求错误中
		msg = err.Error()
		ralCode = -1
	}
	sql, rows := fc()

	fileLineNum := ormUtil.FileWithLineNum()
	fields := l.commonFields(ctx)
	fields = append(fields,
		// zap.String("msg", msg),
		zap.Int64("affectedRow", rows),
		zap.String("requestEndTime", GetFormatRequestTime(end)),
		zap.String("requestStartTime", GetFormatRequestTime(begin)),
		zap.String("file", fileLineNum),
		zap.Float64("cost", cost),
		zap.Int("ralCode", ralCode),
		zap.String("sql", sql),
	)

	l.logger.Info(msg, fields...)
}

func (l OrmLogger) commonFields(ctx context.Context) []zap.Field {
	var logID, uri string
	if c, ok := ctx.(*gin.Context); (ok && c != nil) || (!ok && !IsNil(ctx)) {
		logID, _ = ctx.Value(ContextKeyLogID).(string)
		if logID == "" {
			logID = genLogID()
		}
		uri, _ = ctx.Value(ContextKeyUri).(string)
	}
	context.WithValue(ctx, ContextKeyLogID, logID)
	fields := []zap.Field{
		zap.String(ContextKeyLogID, logID),
		zap.String("uri", uri),
		zap.String("prot", "mysql"),
		zap.String("service", l.Service),
		zap.String("addr", l.Addr),
		zap.String("db", l.Database),
	}
	return fields
}
