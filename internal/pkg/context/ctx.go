package context

import (
	"github.com/gin-gonic/gin"
)

func GetClientIp(ctx *gin.Context) (clientIP string) {
	if ctx == nil {
		return clientIP
	}
	return ctx.ClientIP()
}
