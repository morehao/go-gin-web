package router

import "github.com/gin-gonic/gin"

func RegisterRouter(routerGroup *gin.RouterGroup) {
	exampleRouter(routerGroup)
	userRouter(routerGroup)
}
