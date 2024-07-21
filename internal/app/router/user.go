package router

import (
	"go-gin-web/internal/app/controller/ctrUser"

	"github.com/gin-gonic/gin"
)

func userRouter(routerGroup *gin.RouterGroup) {
	userCtr := ctrUser.NewUserCtr()
	userGroup := routerGroup.Group("/user")
	{
		userGroup.GET("/get", userCtr.Get)
		userGroup.GET("/formatData", userCtr.FormatData)
	}
}
