package router

import (
	"go-gin-web/internal/apps/demoapp/controller/ctrexample"

	"github.com/gin-gonic/gin"
)

func exampleRouter(routerGroup *gin.RouterGroup) {
	exampleCtr := ctrexample.NewExampleCtr()
	exampleGroup := routerGroup.Group("/example")
	{
		exampleGroup.GET("/formatData", exampleCtr.FormatData)
		exampleGroup.GET("/sseTime", exampleCtr.SSETime)
		exampleGroup.GET("/sseTimeRaw", exampleCtr.SSETimeRaw)
		exampleGroup.GET("/sseProcess", exampleCtr.SSEProcess)
		exampleGroup.GET("/sseChat", exampleCtr.SSEChat)
		exampleGroup.GET("/sseRaw", exampleCtr.SSERaw)
	}
}
