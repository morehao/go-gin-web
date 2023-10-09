package zlog

import "go.uber.org/zap"

type LogConfig struct {
	Level  string `yaml:"level"`
	Stdout bool   `yaml:"stdout"`
}

func InitLog(conf LogConfig) *zap.SugaredLogger {
	SugaredLogger = GetLogger()
	return SugaredLogger
}
