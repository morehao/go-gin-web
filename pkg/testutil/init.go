package testutil

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	initializer Initializer
	once        sync.Once
)

// Initializer 定义初始化器接口
type Initializer interface {
	Initialize() error
	Close() error
}

// InitFunc 初始化函数类型
type InitFunc func() Initializer

// 应用初始化函数映射
var initFuncMap = map[string]InitFunc{
	AppNameDemo: newDemo,
}

// Init 初始化测试环境
// appName: 应用名称，如 "demoapp"
func Init(appName string) {
	once.Do(func() {
		initFunc, ok := initFuncMap[appName]
		if !ok {
			panic(fmt.Sprintf("unknown app name: %s", appName))
		}

		initializer = initFunc()
		if err := initializer.Initialize(); err != nil {
			panic(fmt.Sprintf("initialize app %s failed: %v", appName, err))
		}
	})
}

// Done 清理测试环境资源
func Done() {
	if initializer == nil {
		panic("initializer is nil, please call Init first")
	}
	if err := initializer.Close(); err != nil {
		panic(fmt.Sprintf("close initializer failed: %v", err))
	}
}

// NewContext 创建测试用的 gin.Context
func NewContext(opts ...Option) *gin.Context {
	ctx := &gin.Context{}
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

// MustInit 初始化并返回错误（用于需要更精细控制的场景）
func MustInit(appName string) error {
	initFunc, ok := initFuncMap[appName]
	if !ok {
		return fmt.Errorf("unknown app name: %s", appName)
	}

	initializer = initFunc()
	return initializer.Initialize()
}

