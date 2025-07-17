package router

import (
	"github.com/morehao/go-gin-web/internal/apps/demoapp/controller/ctrexample"

	"github.com/gin-gonic/gin"
)

func formatRouter(routerGroup *gin.RouterGroup) {
	formatCtr := ctrexample.NewFormatCtr()
	formatGroup := routerGroup.Group("/format")
	{
		formatGroup.GET("/formatRes", formatCtr.FormatRes)
	}
}

func sseRouter(routerGroup *gin.RouterGroup) {
	sseCtr := ctrexample.NewSSECtr()
	sseGroup := routerGroup.Group("/sse")
	{
		sseGroup.GET("/time", sseCtr.Time)
		sseGroup.GET("/timeRaw", sseCtr.TimeRaw)
		sseGroup.GET("/process", sseCtr.Process)
		sseGroup.GET("/chat", sseCtr.Chat)
		sseGroup.GET("/raw", sseCtr.Raw)
	}
}
