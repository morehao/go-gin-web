package router

import (
	"go-gin-web/internal/app/controller/ctrExample"

	"github.com/gin-gonic/gin"
)

func exampleRouter(routerGroup *gin.RouterGroup) {
	exampleCtr := ctrExample.NewExampleCtr()
	exampleGroup := routerGroup.Group("/example")
	{
		exampleGroup.GET("/formatData", exampleCtr.FormatData)
	}
}
