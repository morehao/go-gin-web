package glog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Buffer 对用户暴露的log配置
type Buffer struct {
	Switch        string        `yaml:"switch"`
	Size          int           `yaml:"size"`
	FlushInterval time.Duration `yaml:"flushInterval"`
}

type LogConfig struct {
	Level  string `yaml:"level"`
	Stdout bool   `yaml:"stdout"`
	Buffer Buffer `yaml:"buffer"`
	Dir    string `yaml:"dir"`
}

func (conf LogConfig) SetLogLevel() {
	logConfig.ZapLevel = getLogLevel(conf.Level)
}

func (conf LogConfig) SetLogOutput() {
	// 开发环境下默认输出到文件，支持自定义是否输出到终端
	logConfig.Log2File = true
	logConfig.Stdout = conf.Stdout
	logConfig.Path = conf.Dir

	// 目录不存在则先创建目录
	if _, err := os.Stat(logConfig.Path); os.IsNotExist(err) {
		err = os.MkdirAll(logConfig.Path, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logConfig.Path, err))
		}
	}
}

// 全局配置 仅限Init函数进行变更
var logConfig = struct {
	ZapLevel  zapcore.Level
	hookField HookFieldFunc

	// 以下变量仅对开发环境生效
	Stdout   bool
	Log2File bool
	Path     string

	// 缓冲区
	BufferSwitch        bool
	BufferSize          int
	BufferFlushInterval time.Duration
}{
	ZapLevel:  zapcore.InfoLevel,
	hookField: defaultHook,

	Stdout:   false,
	Log2File: true,
	Path:     "./log",

	// 缓冲区，如果不配置默认使用以下配置
	BufferSwitch:        true,
	BufferSize:          256 * 1024, // 256kb
	BufferFlushInterval: 5 * time.Second,
}

func InitLog(conf LogConfig) *zap.SugaredLogger {

	// 全局日志级别
	conf.SetLogLevel()
	// 日志输出方式
	conf.SetLogOutput()

	// 初始化全局logger
	SugaredLogger = GetLogger()
	return SugaredLogger
}

func RegisterHookField(hook HookFieldFunc) {
	logConfig.hookField = hook
}
