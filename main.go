package main

import (
	"github.com/gin-gonic/gin"
	"go-web/component/zlog"
	"go-web/middleware"
	"go-web/router/routerHttp"
	"log"
)

func main() {
	zlog.InitLog(zlog.LogConfig{})
	HttpServer()
}

func HttpServer() {
	engine := gin.New()
	engine.Use(middleware.AccessLog(middleware.LoggerConfig{}))
	routerGroup := engine.Group("/go-web")
	routerHttp.Backend(routerGroup)
	if err := engine.Run(":9090"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
