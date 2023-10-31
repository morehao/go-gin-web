package routerHttp

import (
	"github.com/gin-gonic/gin"
	"go-web/controller/ctrExample"
)

func exampleRouterGroup(router *gin.RouterGroup) {
	routerGroup := router.Group("/example")
	{
		routerGroup.GET("format", ctrExample.FormatData)
	}
}
