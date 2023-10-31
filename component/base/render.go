package base

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// 定义render通用类型
type Render interface {
	SetReturnCode(int)
	SetReturnMsg(string)
	SetReturnData(interface{})
	GetReturnCode() int
	GetReturnMsg() string
}

var newRender func() Render

func RegisterRender(s func() Render) {
	newRender = s
}

func newJsonRender() Render {
	if newRender == nil {
		newRender = defaultNew
	}
	return newRender()
}

func RenderJson(ctx *gin.Context, code int, msg string, data interface{}) {
	r := newJsonRender()
	r.SetReturnCode(code)
	r.SetReturnMsg(msg)
	r.SetReturnData(data)
	ctx.JSON(http.StatusOK, r)
	return
}

func RenderJsonSucc(ctx *gin.Context, data interface{}) {
	r := newJsonRender()
	r.SetReturnCode(0)
	r.SetReturnMsg("success")
	r.SetReturnData(data)
	ctx.JSON(http.StatusOK, r)
	return
}

func RenderJsonFail(ctx *gin.Context, err error) {
	r := newJsonRender()

	code, msg := -1, errors.Cause(err).Error()
	switch errors.Cause(err).(type) {
	case Error:
		code = errors.Cause(err).(Error).Code
		msg = errors.Cause(err).(Error).Msg
	default:
	}

	r.SetReturnCode(code)
	r.SetReturnMsg(msg)
	r.SetReturnData(gin.H{})
	ctx.JSON(http.StatusOK, r)

	return
}

func RenderJsonAbort(ctx *gin.Context, err error) {
	r := newJsonRender()

	switch errors.Cause(err).(type) {
	case Error:
		r.SetReturnCode(errors.Cause(err).(Error).Code)
		r.SetReturnMsg(errors.Cause(err).(Error).Msg)
		r.SetReturnData(gin.H{})
	default:
		r.SetReturnCode(-1)
		r.SetReturnMsg(errors.Cause(err).Error())
		r.SetReturnData(gin.H{})
	}
	ctx.AbortWithStatusJSON(http.StatusOK, r)

	return
}

// default render

var defaultNew = func() Render {
	return &DefaultRender{}
}

type DefaultRender struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *DefaultRender) GetReturnCode() int {
	return r.Code
}
func (r *DefaultRender) SetReturnCode(code int) {
	r.Code = code
}
func (r *DefaultRender) GetReturnMsg() string {
	return r.Msg
}
func (r *DefaultRender) SetReturnMsg(msg string) {
	r.Msg = msg
}
func (r *DefaultRender) GetReturnData() interface{} {
	return r.Data
}
func (r *DefaultRender) SetReturnData(data interface{}) {
	ResponseFormat(data)
	r.Data = data
}
