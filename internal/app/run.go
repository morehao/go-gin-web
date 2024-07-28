package app

import (
	"fmt"
	_ "go-gin-web/docs"
	"go-gin-web/internal/app/helper"
	"go-gin-web/internal/app/router"
	"go-gin-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	routerGroup := engine.Group(fmt.Sprintf("/%s", helper.Config.Server.Name))
	routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", helper.Config.Server.Port)); err != nil {
		fmt.Println(fmt.Sprintf("%s run fail, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	} else {
		fmt.Println(fmt.Sprintf("%s run success, port:%s", helper.Config.Server.Name, helper.Config.Server.Port))
	}
}
