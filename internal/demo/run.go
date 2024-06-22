package demo

import (
	"context"
	"fmt"
	"go-gin-web/internal/demo/config"
	"go-gin-web/internal/demo/helper"
	"go-gin-web/internal/demo/router/routerHttp"
	"go-gin-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
)

func Run() {
	engine := gin.Default()

	helper.PreInit()
	defer glog.Close()

	routerGroup := engine.Group("/demo")
	routerGroup.Use(middleware.AccessLog())
	routerHttp.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", config.Cfg.Server.Port)); err != nil {
		glog.Error(context.Background(), fmt.Sprintf("demo run fail, port:%s", config.Cfg.Server.Port))
	} else {
		glog.Info(context.Background(), fmt.Sprintf("demo run success, port:%s", config.Cfg.Server.Port))
	}
}
