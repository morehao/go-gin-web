package zlog

import "go.uber.org/zap"

type LogConfig struct {
	Level     string
	Path      string
	InConsole bool
	AppName   string
}

var logConfig *LogConfig

func InitLog(conf *LogConfig) *zap.SugaredLogger {
	logConfig = conf
	SugaredLogger = GetLogger()
	return SugaredLogger
}
