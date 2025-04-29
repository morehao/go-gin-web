package cuctx

import (
	"github.com/gin-gonic/gin"
)

const (
	UserId = "userId"
)

func GetClientIp(c *gin.Context) string {
	return c.ClientIP()
}

func GetUserID(c *gin.Context) uint64 {
	return c.GetUint64(UserId)
}
