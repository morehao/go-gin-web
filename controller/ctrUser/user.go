package ctrUser

import (
	"github.com/gin-gonic/gin"
	"go-web/component/zlog"
	"go-web/service/srvUser"
)

func Health(ctx *gin.Context) {
	err := srvUser.Health(ctx)
	if err != nil {
		zlog.Infof(ctx, "Health log:%s", "Health")
		return
	}
}
