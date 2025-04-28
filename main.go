package main

import (
	"fmt"
	"os"

	_ "go-gin-web/docs"
	"go-gin-web/internal/app/middleware"
	"go-gin-web/internal/app/router"
	"go-gin-web/internal/pkg/helper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		helper.SetRootDir(workDir)
	}
	helper.PreInit()
	helper.ResourceInit()
	defer helper.Close()
	if helper.Config.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	routerGroup := engine.Group(fmt.Sprintf("/%s", helper.Config.Server.Name))
	if helper.Config.Server.Env == "dev" {
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", helper.Config.Server.Port)); err != nil {
		fmt.Println(fmt.Sprintf("%s run fail, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	} else {
		fmt.Println(fmt.Sprintf("%s run success, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	}
}
