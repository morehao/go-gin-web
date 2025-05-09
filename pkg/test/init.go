package test

import (
	"path"
	"runtime"
	"sync"

	helper2 "go-gin-web/pkg/storages"

	"github.com/gin-gonic/gin"
)

var once sync.Once

func Init() {
	once.Do(func() {
		_, file, _, _ := runtime.Caller(0)
		rootDir := path.Dir(path.Dir(path.Dir(path.Dir(file))))
		helper2.SetRootDir(rootDir)
		helper2.ConfInit()
		helper2.LogInit()
		helper2.ResourceInit()
	})
}

func NewCtx(opts ...Option) *gin.Context {
	ctx := new(gin.Context)
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

func Done() {
	helper2.Close()
}
