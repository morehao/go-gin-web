package routerHttp

import "github.com/gin-gonic/gin"

func RegisterRouter(routerGroup *gin.RouterGroup) {
	registerUserRouter(routerGroup)
}
