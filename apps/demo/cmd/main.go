package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"go-gin-web/apps/demo/config"
	"go-gin-web/apps/demo/internal/middleware"
	"go-gin-web/apps/demo/internal/router"
	_ "go-gin-web/docs"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_, workDir, _, _ := runtime.Caller(0)
	config.SetRootDir(filepath.Dir(filepath.Dir(workDir)))
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
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", config.Conf.Server.Port)); err != nil {
		fmt.Println(fmt.Sprintf("%s run fail, port:%s", config.Conf.Server.Name, config.Conf.Server.Port))
	} else {
		fmt.Println(fmt.Sprintf("%s run success, port:%s", config.Conf.Server.Name, config.Conf.Server.Port))
	}
}
