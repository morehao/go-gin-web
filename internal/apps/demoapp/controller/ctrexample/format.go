package ctrexample

import (
	"go-gin-web/internal/apps/demoapp/service/svcexample"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
)

type FormatCtr interface {
	FormatRes(ctx *gin.Context)
}

type formatCtr struct {
	exampleSvc svcexample.FormatSvc
}

var _ FormatCtr = (*formatCtr)(nil)

func NewFormatCtr() FormatCtr {
	return &formatCtr{
		exampleSvc: svcexample.NewFormatSvc(),
	}
}

func (ctr *formatCtr) FormatRes(ctx *gin.Context) {
	res := ctr.exampleSvc.FormatRes(ctx)

	gincontext.SuccessWithFormat(ctx, res)
}
