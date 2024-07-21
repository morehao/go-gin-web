package app

import (
	"context"
	"fmt"
	"go-gin-web/internal/app/helper"
	"go-gin-web/internal/app/router"
	"go-gin-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
)

func Run() {
	helper.SetRootDir("/internal/app")
	helper.PreInit()
	helper.InitDbClient()
	defer helper.Close()
	if helper.Config.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	routerGroup := engine.Group("/app")
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", helper.Config.Server.Port)); err != nil {
		glog.Error(context.Background(), fmt.Sprintf("%s run fail, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	} else {
		glog.Info(context.Background(), fmt.Sprintf("%s run success, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	}
}
