package test

import (
	"go-gin-web/internal/pkg/context"

	"github.com/gin-gonic/gin"
)

type Option func(ctx *gin.Context)

func OptUid(uid uint64) Option {
	return func(ctx *gin.Context) {
		ctx.Set(context.UserId, uid)
	}
}