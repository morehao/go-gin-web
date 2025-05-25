package test

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var initializer Initializer
var once sync.Once
var lock sync.Mutex

func Init(appName string) {
	once.Do(func() {
		lock.Lock()
		defer lock.Unlock()
		initFunc, ok := initFuncMap[appName]
		if !ok {
			panic("unknown app name: " + appName)
		}
		initializer = initFunc()
		if err := initializer.Initialize(); err != nil {
			panic(err)
		}
	})
}

func Done() {
	if initializer == nil {
		panic("initializer is nil")
	}
	if err := initializer.Close(); err != nil {
		panic(err)
	}
}

type Initializer interface {
	Initialize() error
	Close() error
}

type InitFunc func() Initializer

var initFuncMap = map[string]InitFunc{
	AppNameDemo: newDemo,
}

func NewCtx(opts ...Option) *gin.Context {
	ctx := new(gin.Context)
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}
