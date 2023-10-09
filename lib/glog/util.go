package glog

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// util key
const (
	ContextKeyRequestID = "requestId"
	ContextKeyLogID     = "logId"
	ContextKeyNoLog     = "_no_log"
	ContextKeyUri       = "_uri"
	zapLoggerAddr       = "_zap_addr"
	sugaredLoggerAddr   = "_sugared_addr"
	customerFieldKey    = "__customerFields"
)

// header key
const (
	HeaderKeyTraceID    = "Trace-Id"
	HeaderKeyLogID      = "X_LOGID"
	HeaderKeyLowerLogID = "x_logid"
)

// GetLogID 兼容虚拟机调用项目logId串联问题
func GetLogID(ctx *gin.Context) string {
	if ctx == nil {
		return genLogID()
	}

	// 如果上下文有logId，则直接返回
	if logID := ctx.GetString(ContextKeyLogID); logID != "" {
		return logID
	}

	// 从header中获取
	var logID string
	if ctx.Request != nil && ctx.Request.Header != nil {
		logID = ctx.GetHeader(HeaderKeyLogID)
		if logID == "" {
			logID = ctx.GetHeader(HeaderKeyLowerLogID)
		}
	}

	if logID == "" {
		logID = genLogID()
	}

	ctx.Set(ContextKeyLogID, logID)
	return logID
}

func GetRequestID(ctx *gin.Context) string {
	if ctx == nil {
		return genRequestID()
	}

	// 如果上下文有requestId，则直接返回
	if requestId := ctx.GetString(ContextKeyRequestID); requestId != "" {
		return requestId
	}

	// 从header中获取
	var requestID string
	if ctx.Request != nil && ctx.Request.Header != nil {
		requestID = ctx.Request.Header.Get(HeaderKeyTraceID)
	}

	// 新生成
	if requestID == "" {
		requestID = genRequestID()
	}

	ctx.Set(ContextKeyRequestID, requestID)
	return requestID
}

func genLogID() (requestId string) {
	// 随机生成
	usec := uint64(time.Now().UnixNano())
	requestId = strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
	return requestId
}

var generator = NewRand(time.Now().UnixNano())

func genRequestID() string {
	// 生成 uint64的随机数, 并转换成16进制表示方式
	number := uint64(generator.Int63())
	traceID := fmt.Sprintf("%016x", number)

	var buffer bytes.Buffer
	buffer.WriteString(traceID)
	buffer.WriteString(":")
	buffer.WriteString(traceID)
	buffer.WriteString(":0:1")
	return buffer.String()
}

func GetRequestUri(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetString(ContextKeyUri)
}

// // 用户自定义Notice
// func AddNotice(ctx *gin.Context, key string, val interface{}) {
// 	if meta, ok := metadata.CtxFromGinContext(ctx); ok {
// 		if n := metadata.Value(meta, metadata.Notice); n != nil {
// 			if _, ok = n.(map[string]interface{}); ok {
// 				notices := n.(map[string]interface{})
// 				notices[key] = val
// 			}
// 		}
// 	}
// }

// a new method for customer notice
func AddField(c *gin.Context, field ...Field) {
	customerFields := GetCustomerFields(c)
	if customerFields == nil {
		customerFields = field
	} else {
		customerFields = append(customerFields, field...)
	}

	c.Set(customerFieldKey, customerFields)
}

// 获得所有用户自定义的Field
func GetCustomerFields(c *gin.Context) (customerFields []Field) {
	if v, exist := c.Get(customerFieldKey); exist {
		customerFields, _ = v.([]Field)
	}
	return customerFields
}

// server.log 中打印出用户添加的所有Field
func PrintFields(ctx *gin.Context) {
	fields := GetCustomerFields(ctx)
	zapLogger(ctx).Info("notice", fields...)
}

func SetNoLogFlag(ctx *gin.Context) {
	ctx.Set(ContextKeyNoLog, true)
}

func SetLogFlag(ctx *gin.Context) {
	ctx.Set(ContextKeyNoLog, false)
}

type LockedSource struct {
	mut sync.Mutex
	src rand.Source
}

// NewRand returns a rand.Rand that is threadsafe.
func NewRand(seed int64) *rand.Rand {
	return rand.New(&LockedSource{src: rand.NewSource(seed)})
}

func (r *LockedSource) Int63() (n int64) {
	r.mut.Lock()
	n = r.src.Int63()
	r.mut.Unlock()
	return
}

// Seed implements Seed() of Source
func (r *LockedSource) Seed(seed int64) {
	r.mut.Lock()
	r.src.Seed(seed)
	r.mut.Unlock()
}
