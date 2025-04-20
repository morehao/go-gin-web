package ctrExample

import (
	"go-gin-web/internal/app/service/svcExample"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcontext/ginrender"
)

type ExampleCtr interface {
	FormatData(c *gin.Context)
}

type exampleCtr struct {
	exampleSvc svcExample.ExampleSvc
}

var _ ExampleCtr = (*exampleCtr)(nil)

func NewExampleCtr() ExampleCtr {
	return &exampleCtr{
		exampleSvc: svcExample.NewExampleSvc(),
	}
}

func (ctr *exampleCtr) FormatData(c *gin.Context) {
	res := ctr.exampleSvc.FormatData(c)

	ginrender.SuccessWithFormat(c, res)
}
