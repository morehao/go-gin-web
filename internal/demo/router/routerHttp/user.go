package routerHttp

import (
	"go-gin-web/internal/demo/controller/ctrUser"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(routerGroup *gin.RouterGroup) {
	userCtr := ctrUser.NewUserCtr()
	userGroup := routerGroup.Group("/user")
	{
		userGroup.GET("/get", userCtr.Get)
		userGroup.GET("/formatData", userCtr.FormatData)
	}
}
