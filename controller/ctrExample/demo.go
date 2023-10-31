package ctrExample

import (
	"github.com/gin-gonic/gin"
	"go-web/component/base"
	"go-web/service/srvExample"
)

func FormatData(ctx *gin.Context) {
	res := srvExample.FormatData(ctx)
	base.RenderJsonSucc(ctx, res)
}
