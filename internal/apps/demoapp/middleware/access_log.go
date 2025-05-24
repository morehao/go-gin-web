package middleware

import (
	"bytes"
	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/gerror"
	"github.com/morehao/golib/glog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	reqBodyMaxLen  = 10240
	respBodyMaxLen = 10240
	reqQueryMaxLen = 10240
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestId := getRequestId(c)
		c.Set(glog.KeyRequestId, requestId)

		path := c.Request.URL.Path
		c.Set(glog.KeyUri, path)

		reqQuery := gincontext.GetReqQuery(c)
		// 截断参数
		if len(reqQuery) > reqQueryMaxLen {
			reqQuery = reqQuery[:reqQueryMaxLen]
		}

		reqBody, getBodyErr := gincontext.GetReqBody(c)
		if getBodyErr != nil {
			c.Error(getBodyErr)
		}
		reqBodySize := len(reqBody)
		if len(reqBody) > reqBodyMaxLen {
			reqBody = reqBody[:reqBodyMaxLen]
		}

		// Body writer
		respBodyWriter := &gincontext.RespWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = respBodyWriter

		c.Next()

		end := time.Now()
		cost := glog.GetRequestCost(start, end)

		responseBody := ""
		var responseBodySize int
		var errInfo gerror.Error
		if respBodyWriter.Body != nil && respBodyMaxLen != -1 {
			responseBody = respBodyWriter.Body.String()
			responseBodySize = len(responseBody)
			if responseBodySize > 0 {
				if err := jsoniter.Unmarshal([]byte(responseBody), &errInfo); err != nil {
					c.Error(err)
				}
			}
			if len(responseBody) > respBodyMaxLen {
				responseBody = responseBody[:respBodyMaxLen]
			}
		}

		keysAndValues := []interface{}{
			glog.KeyHost, c.Request.Host,
			glog.KeyClientIp, gincontext.GetClientIp(c),
			glog.KeyHandle, c.HandlerName(),
			glog.KeyProto, c.Request.Proto,
			glog.KeyRefer, c.Request.Referer(),
			glog.KeyUserAgent, c.Request.UserAgent(),
			glog.KeyHeader, gincontext.GetHeader(c),
			glog.KeyCookie, gincontext.GetCookie(c),
			glog.KeyUri, path,
			glog.KeyMethod, c.Request.Method,
			glog.KeyHttpStatusCode, c.Writer.Status(),
			glog.KeyRequestQuery, reqQuery,
			glog.KeyRequestBody, reqBody,
			glog.KeyRequestBodySize, reqBodySize,
			glog.KeyResponseBody, responseBody,
			glog.KeyResponseBodySize, responseBodySize,
			glog.KeyRequestStartTime, glog.FormatRequestTime(start),
			glog.KeyRequestEndTime, glog.FormatRequestTime(end),
			glog.KeyCost, cost,
			glog.KeyErrorCode, errInfo.Code,
			glog.KeyErrorMsg, errInfo.Msg,
			glog.KeyRequestErr, c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}
		glog.Infow(c, glog.MsgFlagNotice, keysAndValues...)
	}
}

func getRequestId(c *gin.Context) string {
	requestId := c.Request.Header.Get(glog.KeyRequestId)
	if requestId == "" {
		requestId = c.GetString(glog.KeyRequestId)
	}
	if requestId == "" {
		requestId = glog.GenRequestID()
	}
	return requestId
}

func AccessLogOtel() gin.HandlerFunc {
	traceProvider := trace.NewTracerProvider()
	otel.SetTracerProvider(traceProvider)
	tr := traceProvider.Tracer("gin-server")
	return func(c *gin.Context) {
		traceId, spanId, traceFlags := getTraceInfo(c, tr)
		c.Set(glog.KeyTraceId, traceId)
		c.Set(glog.KeyTraceFlags, traceFlags)
		c.Set(glog.KeySpanId, spanId)
		c.Set(glog.KeyRequestId, c.Request.Header.Get(glog.KeyRequestId))
		glog.Info(c, "[middleware]")
		c.Next()
	}
}

func getTraceInfo(c *gin.Context, tracer oteltrace.Tracer) (string, string, string) {
	traceId := c.Request.Header.Get(glog.KeyTraceId)
	spanId := c.Request.Header.Get(glog.KeySpanId)
	traceFlags := c.Request.Header.Get(glog.KeyTraceFlags)

	if traceId == "" {
		traceId = c.GetString(glog.KeyTraceId)
		spanId = c.GetString(glog.KeySpanId)
		traceFlags = c.GetString(glog.KeyTraceFlags)
	}

	rCtx := c.Request.Context()
	if traceId == "" {
		spanCtx, span := tracer.Start(rCtx, c.Request.URL.Path)
		defer span.End()
		traceId = span.SpanContext().TraceID().String()
		spanId = span.SpanContext().SpanID().String()
		traceFlags = span.SpanContext().TraceFlags().String()
		c.Request = c.Request.WithContext(spanCtx)
	} else {
		newTraceId, _ := oteltrace.TraceIDFromHex(traceId)
		traceFlagsByte := byte(1)
		if traceFlags != "" {
			decoded, err := hex.DecodeString(traceFlags)
			if err == nil && len(decoded) > 0 {
				traceFlagsByte = decoded[0]
			}
		}
		spanContextCfg := oteltrace.SpanContextConfig{
			TraceID:    newTraceId,
			TraceFlags: oteltrace.TraceFlags(traceFlagsByte),
		}
		if spanId != "" {
			newSpanId, _ := oteltrace.SpanIDFromHex(spanId)
			spanContextCfg.SpanID = newSpanId
		} else {
			_, span := tracer.Start(oteltrace.ContextWithRemoteSpanContext(rCtx, oteltrace.NewSpanContext(spanContextCfg)), "generateSpanId")
			newSpanId := span.SpanContext().SpanID()
			span.End()
			spanContextCfg.SpanID = newSpanId
		}
		spanContext := oteltrace.NewSpanContext(spanContextCfg)
		spanCtx := oteltrace.ContextWithRemoteSpanContext(rCtx, spanContext)
		_, span := tracer.Start(spanCtx, c.Request.URL.Path)
		defer span.End()
		traceId = span.SpanContext().TraceID().String()
		spanId = span.SpanContext().SpanID().String()
		traceFlags = span.SpanContext().TraceFlags().String()
		c.Request = c.Request.WithContext(spanCtx)
	}
	return traceId, spanId, traceFlags
}
