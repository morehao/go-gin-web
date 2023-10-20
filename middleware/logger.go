package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-web/component/zlog"
	"go-web/utils"
	"go.uber.org/zap"
	"io"
	"strings"
	"time"
)

const (
	printRequestLen  = 10240
	printResponseLen = 10240
)

type customRespWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w customRespWriter) WriteString(s string) (int, error) {
	if w.body != nil {
		w.body.WriteString(s)
	}
	return w.ResponseWriter.WriteString(s)
}

func (w customRespWriter) Write(b []byte) (int, error) {
	if w.body != nil {
		w.body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}

// access日志打印
type LoggerConfig struct {
	// SkipPaths is a url path array which logs are not written.
	SkipPaths []string `yaml:"skipPaths"`

	// request body 最大长度展示，0表示采用默认的10240，-1表示不打印
	MaxReqBodyLen int `yaml:"maxReqBodyLen"`
	// response body 最大长度展示，0表示采用默认的10240，-1表示不打印。指定长度的时候需注意，返回的json可能被截断
	MaxRespBodyLen int `yaml:"maxRespBodyLen"`

	// 自定义Skip功能
	Skip func(ctx *gin.Context) bool
}

func AccessLog(conf LoggerConfig) gin.HandlerFunc {
	notLogged := conf.SkipPaths
	var skip map[string]struct{}
	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range notLogged {
			skip[path] = struct{}{}
		}
	}

	maxReqBodyLen := conf.MaxReqBodyLen
	if maxReqBodyLen == 0 {
		maxReqBodyLen = printRequestLen
	}

	maxRespBodyLen := conf.MaxRespBodyLen
	if maxRespBodyLen == 0 {
		maxRespBodyLen = printResponseLen
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// body writer
		blw := &customRespWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		reqBody := getReqBody(c, maxReqBodyLen)
		reqQuery := getReqQuery(c, maxReqBodyLen)

		c.Set(zlog.ContextKeyUri, path)
		_ = zlog.GetLogId(c)

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; ok {
			return
		}

		if conf.Skip != nil && conf.Skip(c) {
			return
		}

		// Stop timer
		end := time.Now()

		response := ""
		if blw.body != nil && maxRespBodyLen != -1 {
			response = blw.body.String()
			if len(response) > maxRespBodyLen {
				response = response[:maxRespBodyLen]
			}
		}

		commonFields := []zap.Field{
			zap.String(zlog.ContextKeyUid, getReqValueByKey(c, "userid")),
			zap.String("host", c.Request.Host),
			zap.String("method", c.Request.Method),
			zap.String("httpProto", c.Request.Proto),
			zap.String("handle", c.HandlerName()),
			zap.String("userAgent", c.Request.UserAgent()),
			zap.String("refer", c.Request.Referer()),
			zap.String("clientIp", utils.GetClientIp(c)),
			zap.String("cookie", getCookie(c)),
			zap.String("requestStartTime", zlog.GetFormatRequestTime(start)),
			zap.String("requestEndTime", zlog.GetFormatRequestTime(end)),
			zap.Float64("cost", zlog.GetRequestCost(start, end)),
			zap.String("requestBody", reqBody),
			zap.String("requestQuery", reqQuery),
			zap.Int("responseStatus", c.Writer.Status()),
			zap.String("response", response),
			zap.Int("bodySize", c.Writer.Size()),
			zap.String("client", c.GetHeader("X-APP")), // 请求来源app
			zap.String("reqErr", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		}

		zlog.InfoLogger(c, "notice", commonFields...)
		zlog.Errorf(c, "测试err日志")
		zlog.Warnf(c, "测试warn日志")
	}
}

// 请求参数
func getReqBody(c *gin.Context, maxReqBodyLen int) (reqBody string) {
	if maxReqBodyLen == 0 {
		return reqBody
	}

	// body中的参数
	if c.Request.Body != nil && c.ContentType() == binding.MIMEMultipartPOSTForm {
		requestBody, err := c.GetRawData()
		if err != nil {
			zlog.WarnLogger(c, "get http request body error: "+err.Error())
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		if _, err := c.MultipartForm(); err != nil {
			zlog.WarnLogger(c, "parse http request form body error: "+err.Error())
		}
		reqBody = c.Request.PostForm.Encode()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	} else if c.Request.Body != nil && c.ContentType() == "application/octet-stream" {

	} else if c.Request.Body != nil {
		requestBody, err := c.GetRawData()
		if err != nil {
			zlog.WarnLogger(c, "get http request body error: "+err.Error())
		}
		reqBody = string(requestBody)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	// 截断参数
	if len(reqBody) > maxReqBodyLen {
		reqBody = reqBody[:maxReqBodyLen]
	}

	return reqBody
}

func getReqQuery(c *gin.Context, maxReqQueryLen int) (reqQuery string) {
	if maxReqQueryLen == 0 {
		return reqQuery
	}
	reqQuery = c.Request.URL.RawQuery
	// 截断参数
	if len(reqQuery) > maxReqQueryLen {
		reqQuery = reqQuery[:maxReqQueryLen]
	}
	return reqQuery
}

// 从request body中解析特定字段作为notice key打印
func getReqValueByKey(ctx *gin.Context, k string) string {
	if vs, exist := ctx.Request.Form[k]; exist && len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func getCookie(ctx *gin.Context) string {
	cStr := ""
	for _, c := range ctx.Request.Cookies() {
		cStr += fmt.Sprintf("%s=%s&", c.Name, c.Value)
	}
	return strings.TrimRight(cStr, "&")
}
