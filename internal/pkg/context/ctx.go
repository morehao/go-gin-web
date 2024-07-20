package context

import (
	"github.com/gin-gonic/gin"
)

const (
	UserId = "userId"
)

func GetClientIp(ctx *gin.Context) (clientIP string) {
	if ctx == nil {
		return clientIP
	}
	return ctx.ClientIP()
}

func GetUserId(ctx *gin.Context) (userId uint64) {
	return ctx.GetUint64(UserId)
}
