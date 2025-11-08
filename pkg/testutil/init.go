package testutil

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
)

func init() {
	// 设置 gin 为测试模式
	gin.SetMode(gin.TestMode)
}

var (
	initializers = make(map[string]Initializer) // 存储所有已初始化的应用
	initOnce     = make(map[string]*sync.Once)  // 每个应用只初始化一次
	mu           sync.RWMutex                    // 保护全局状态
)

// Initializer 定义初始化器接口
type Initializer interface {
	Initialize() error
	Close() error
}

// InitializerFunc 初始化器构造函数类型
type InitializerFunc func() (Initializer, error)

// 应用初始化器映射
// 每个应用都应该在这里注册自己的初始化器构造函数
var initializerFuncMap = map[string]InitializerFunc{
	AppNameDemo: newDemoappInitializer,
}

// RegisterApp 注册新的应用初始化器
// 允许其他包动态注册自己的应用
func RegisterApp(appName string, initFunc InitializerFunc) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := initializerFuncMap[appName]; exists {
		panic(fmt.Sprintf("app %s already registered", appName))
	}
	initializerFuncMap[appName] = initFunc
}

// Initialize 初始化指定应用的测试环境
// 支持多次调用，但每个应用只会初始化一次（幂等性）
// 如果初始化失败会 panic
func Initialize(appName string) {
	mu.Lock()
	// 获取或创建该应用的 sync.Once
	once, ok := initOnce[appName]
	if !ok {
		once = &sync.Once{}
		initOnce[appName] = once
	}
	mu.Unlock()

	// 确保该应用只初始化一次
	once.Do(func() {
		mu.RLock()
		initFunc, ok := initializerFuncMap[appName]
		mu.RUnlock()

		if !ok {
			panic(fmt.Sprintf("app %s not registered", appName))
		}

		initializer, err := initFunc()
		if err != nil {
			panic(fmt.Sprintf("create initializer for app %s failed: %v", appName, err))
		}

		if err := initializer.Initialize(); err != nil {
			panic(fmt.Sprintf("initialize app %s failed: %v", appName, err))
		}

		// 存储初始化器
		mu.Lock()
		initializers[appName] = initializer
		mu.Unlock()
	})
}

// Close 清理指定应用的测试环境资源
// 如果 appName 为空字符串，则清理所有已初始化的应用
func Close(appName string) {
	if appName != "" {
		// 清理指定应用
		mu.RLock()
		initializer, ok := initializers[appName]
		mu.RUnlock()

		if !ok {
			return // 应用未初始化，无需清理
		}

		if err := initializer.Close(); err != nil {
			// 静默失败，避免影响测试
		}

		mu.Lock()
		delete(initializers, appName)
		delete(initOnce, appName)
		mu.Unlock()
	} else {
		// 清理所有应用
		mu.Lock()
		defer mu.Unlock()

		for _, initializer := range initializers {
			if err := initializer.Close(); err != nil {
				// 静默失败，避免影响测试
			}
		}

		initializers = make(map[string]Initializer)
		initOnce = make(map[string]*sync.Once)
	}
}

// NewContext 创建测试用的 gin.Context
// 自动设置基础字段（如 Request、RequestID 等）
func NewContext(opts ...Option) *gin.Context {
	// 创建基础 Context
	ctx := &gin.Context{
		Request: &http.Request{
			URL:    &url.URL{},
			Header: http.Header{},
		},
	}
	ctx.Request = ctx.Request.WithContext(context.Background())

	// 应用所有选项
	for _, opt := range opts {
		opt(ctx)
	}

	// 如果没有设置 RequestID，自动生成一个
	if _, exists := ctx.Get(glog.KeyRequestId); !exists {
		ctx.Set(glog.KeyRequestId, generateRequestID())
	}

	return ctx
}

// Init 是 Initialize 的别名，保持向后兼容
func Init(appName string) {
	Initialize(appName)
}

// Done 是 Close 的别名，保持向后兼容
func Done(appName string) {
	Close(appName)
}

