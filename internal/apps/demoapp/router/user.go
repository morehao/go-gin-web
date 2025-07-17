package router

import (
	"github.com/morehao/go-gin-web/internal/apps/demoapp/controller/ctruser"

	"github.com/gin-gonic/gin"
)

// userRouter 初始化用户管理路由信息
func userRouter(routerGroup *gin.RouterGroup) {
	userCtr := ctruser.NewUserCtr()
	userGroup := routerGroup.Group("user")
	{
		userGroup.POST("create", userCtr.Create)    // 新建用户管理
		userGroup.POST("delete", userCtr.Delete)    // 删除用户管理
		userGroup.POST("update", userCtr.Update)    // 更新用户管理
		userGroup.GET("detail", userCtr.Detail)     // 根据ID获取用户管理
		userGroup.GET("pageList", userCtr.PageList) // 获取用户管理列表
	}
}
