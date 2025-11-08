package testutil

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
)

// Option 定义上下文配置选项
type Option func(ctx *gin.Context)

// WithUserID 设置用户ID
func WithUserID(uid uint) Option {
	return func(ctx *gin.Context) {
		ctx.Set(gincontext.UserID, uid)
	}
}

// WithRequestID 设置请求ID
func WithRequestID(requestID string) Option {
	return func(ctx *gin.Context) {
		ctx.Set(glog.KeyRequestId, requestID)
	}
}

// WithKeyValue 设置自定义键值对
func WithKeyValue(key string, value interface{}) Option {
	return func(ctx *gin.Context) {
		ctx.Set(key, value)
	}
}

