package test

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
)

type Option func(ctx *gin.Context)

func OptUid(uid uint) Option {
	return func(ctx *gin.Context) {
		ctx.Set(gincontext.UserID, uid)
	}
}
