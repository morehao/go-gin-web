package ctrexample

import (
	"go-gin-web/internal/apps/demoapp/service/svcexample"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
)

type ExampleCtr interface {
	FormatData(c *gin.Context)
}

type exampleCtr struct {
	exampleSvc svcexample.ExampleSvc
}

var _ ExampleCtr = (*exampleCtr)(nil)

func NewExampleCtr() ExampleCtr {
	return &exampleCtr{
		exampleSvc: svcexample.NewExampleSvc(),
	}
}

func (ctr *exampleCtr) FormatData(c *gin.Context) {
	res := ctr.exampleSvc.FormatData(c)

	gincontext.SuccessWithFormat(c, res)
}
