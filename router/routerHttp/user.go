package routerHttp

import (
	"github.com/gin-gonic/gin"
	"go-web/controller/ctrUser"
)

func userRouterGroup(router *gin.RouterGroup) {
	accountGroup := router.Group("/account")
	{
		accountGroup.GET("health", ctrUser.Health)
	}
}
