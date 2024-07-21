package demo

import (
	"context"
	"fmt"
	"go-gin-web/internal/demo/helper"
	"go-gin-web/internal/demo/router"
	"go-gin-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
)

func Run() {
	engine := gin.Default()
	helper.SetRootDir("/internal/demo")
	helper.PreInit()
	helper.InitDbClient()
	defer helper.Close()

	routerGroup := engine.Group("/demo")
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", helper.Config.Server.Port)); err != nil {
		glog.Error(context.Background(), fmt.Sprintf("demo run fail, port:%s", helper.Config.Server.Port))
	} else {
		glog.Info(context.Background(), fmt.Sprintf("demo run success, port:%s", helper.Config.Server.Port))
	}
}
