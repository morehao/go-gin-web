package test

import (
	"path"
	"runtime"
	"sync"

	"go-gin-web/internal/pkg/helper"

	"github.com/gin-gonic/gin"
)

var once sync.Once

func Init() {
	once.Do(func() {
		_, file, _, _ := runtime.Caller(0)
		rootDir := path.Dir(path.Dir(path.Dir(path.Dir(file))))
		helper.SetRootDir(rootDir)
		helper.ConfInit()
		helper.LogInit()
		helper.ResourceInit()
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
	helper.Close()
}
