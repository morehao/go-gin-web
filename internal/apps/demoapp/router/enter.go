package router

import "github.com/gin-gonic/gin"

func RegisterRouter(routerGroup *gin.RouterGroup) {
	formatRouter(routerGroup)
	sseRouter(routerGroup)
	clientRouter(routerGroup)
	userRouter(routerGroup)
}
