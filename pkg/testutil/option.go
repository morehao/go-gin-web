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

// WithHeader 设置 HTTP 请求头
func WithHeader(key, value string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.Header != nil {
			ctx.Request.Header.Set(key, value)
		}
	}
}

// WithHeaders 批量设置 HTTP 请求头
func WithHeaders(headers map[string]string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.Header != nil {
			for k, v := range headers {
				ctx.Request.Header.Set(k, v)
			}
		}
	}
}

// WithMethod 设置 HTTP 请求方法
func WithMethod(method string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil {
			ctx.Request.Method = method
		}
	}
}

// WithURL 设置请求 URL
func WithURL(url string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.URL != nil {
			ctx.Request.URL.Path = url
		}
	}
}

// WithQueryParam 添加查询参数
func WithQueryParam(key, value string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.URL != nil {
			q := ctx.Request.URL.Query()
			q.Add(key, value)
			ctx.Request.URL.RawQuery = q.Encode()
		}
	}
}

// WithQueryParams 批量添加查询参数
func WithQueryParams(params map[string]string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.URL != nil {
			q := ctx.Request.URL.Query()
			for k, v := range params {
				q.Add(k, v)
			}
			ctx.Request.URL.RawQuery = q.Encode()
		}
	}
}

// WithContentType 设置 Content-Type
func WithContentType(contentType string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.Header != nil {
			ctx.Request.Header.Set("Content-Type", contentType)
		}
	}
}

// WithJSON 设置 Content-Type 为 application/json
func WithJSON() Option {
	return WithContentType("application/json")
}

// WithFormData 设置 Content-Type 为 application/x-www-form-urlencoded
func WithFormData() Option {
	return WithContentType("application/x-www-form-urlencoded")
}

// WithAuth 设置 Authorization 请求头
func WithAuth(token string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil && ctx.Request.Header != nil {
			ctx.Request.Header.Set("Authorization", token)
		}
	}
}

// WithBearerToken 设置 Bearer Token
func WithBearerToken(token string) Option {
	return WithAuth("Bearer " + token)
}

// WithClientIP 设置客户端 IP
func WithClientIP(ip string) Option {
	return func(ctx *gin.Context) {
		if ctx.Request != nil {
			ctx.Request.RemoteAddr = ip
		}
	}
}

