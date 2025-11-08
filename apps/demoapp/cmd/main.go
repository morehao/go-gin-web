package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-gin-web/apps/demoapp/config"
	_ "github.com/morehao/go-gin-web/apps/demoapp/docs"
	"github.com/morehao/go-gin-web/apps/demoapp/middleware"
	"github.com/morehao/go-gin-web/apps/demoapp/router"
	"github.com/morehao/golib/glog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := serverInit(); err != nil {
		panic(fmt.Sprintf("server init failed, error: %v", err))
	}
	if config.Conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	defer glog.Close()

	engine := gin.New()
	engine.Use(gin.Recovery())
	routerGroup := engine.Group(fmt.Sprintf("/%s", config.Conf.Server.Name))
	if config.Conf.Server.Env == "dev" {
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("demoapp")))
	}
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", config.Conf.Server.Port)); err != nil {
		glog.Errorf(context.Background(), "%s run fail, port:%s", config.Conf.Server.Name, config.Conf.Server.Port)
		panic(err)
	} else {
		glog.Infof(context.Background(), "%s run success, port:%s", config.Conf.Server.Name, config.Conf.Server.Port)
	}
}
