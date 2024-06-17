package main

import (
	"go-web/router/routerHttp"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// zlog.InitLog(&zlog.LogConfig{})
	HttpServer()
}

func HttpServer() {
	engine := gin.New()
	// engine.Use(middleware.AccessLog(middleware.LoggerConfig{}))
	routerGroup := engine.Group("/go-web")
	routerHttp.Backend(routerGroup)
	if err := engine.Run(":9090"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
