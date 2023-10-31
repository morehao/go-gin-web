package routerHttp

import (
	"github.com/gin-gonic/gin"
)

func Backend(router *gin.RouterGroup) {
	router = router.Group("/backend")
	userRouterGroup(router)
	exampleRouterGroup(router)
}
