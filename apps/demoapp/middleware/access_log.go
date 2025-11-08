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
	return func(ctx *gin.Context) {
		start := time.Now()

		requestId := getRequestId(ctx)
		ctx.Set(glog.KeyRequestId, requestId)

		path := ctx.Request.URL.Path
		ctx.Set(glog.KeyUrl, path)

		reqQuery := gincontext.GetReqQuery(ctx)
		// 截断参数
		if len(reqQuery) > reqQueryMaxLen {
			reqQuery = reqQuery[:reqQueryMaxLen]
		}

		reqBody, getBodyErr := gincontext.GetReqBody(ctx)
		if getBodyErr != nil {
			ctx.Error(getBodyErr)
		}
		reqBodySize := len(reqBody)
		if len(reqBody) > reqBodyMaxLen {
			reqBody = reqBody[:reqBodyMaxLen]
		}

		// Body writer
		respBodyWriter := &gincontext.RespWriter{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = respBodyWriter

		ctx.Next()

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
					ctx.Error(err)
				}
			}
			if len(responseBody) > respBodyMaxLen {
				responseBody = responseBody[:respBodyMaxLen]
			}
		}

		keysAndValues := []interface{}{
			glog.KeyHost, ctx.Request.Host,
			glog.KeyClientIp, gincontext.GetClientIp(ctx),
			glog.KeyHandle, ctx.HandlerName(),
			glog.KeyProto, ctx.Request.Proto,
			glog.KeyRefer, ctx.Request.Referer(),
			glog.KeyHeader, gincontext.GetHeader(ctx),
			glog.KeyCookie, gincontext.GetCookie(ctx),
			glog.KeyUrl, path,
			glog.KeyMethod, ctx.Request.Method,
			glog.KeyHttpStatusCode, ctx.Writer.Status(),
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
			glog.KeyRequestErr, ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
		}
		glog.Infow(ctx, glog.MsgFlagNotice, keysAndValues...)
	}
}

func getRequestId(ctx *gin.Context) string {
	requestId := ctx.Request.Header.Get(glog.KeyRequestId)
	if requestId == "" {
		requestId = ctx.GetString(glog.KeyRequestId)
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
	return func(ctx *gin.Context) {
		traceId, spanId, traceFlags := getTraceInfo(ctx, tr)
		ctx.Set(glog.KeyTraceId, traceId)
		ctx.Set(glog.KeyTraceFlags, traceFlags)
		ctx.Set(glog.KeySpanId, spanId)
		ctx.Set(glog.KeyRequestId, ctx.Request.Header.Get(glog.KeyRequestId))
		glog.Info(ctx, "[middleware]")
		ctx.Next()
	}
}

func getTraceInfo(ctx *gin.Context, tracer oteltrace.Tracer) (string, string, string) {
	traceId := ctx.Request.Header.Get(glog.KeyTraceId)
	spanId := ctx.Request.Header.Get(glog.KeySpanId)
	traceFlags := ctx.Request.Header.Get(glog.KeyTraceFlags)

	if traceId == "" {
		traceId = ctx.GetString(glog.KeyTraceId)
		spanId = ctx.GetString(glog.KeySpanId)
		traceFlags = ctx.GetString(glog.KeyTraceFlags)
	}

	rCtx := ctx.Request.Context()
	if traceId == "" {
		spanCtx, span := tracer.Start(rCtx, ctx.Request.URL.Path)
		defer span.End()
		traceId = span.SpanContext().TraceID().String()
		spanId = span.SpanContext().SpanID().String()
		traceFlags = span.SpanContext().TraceFlags().String()
		ctx.Request = ctx.Request.WithContext(spanCtx)
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
		_, span := tracer.Start(spanCtx, ctx.Request.URL.Path)
		defer span.End()
		traceId = span.SpanContext().TraceID().String()
		spanId = span.SpanContext().SpanID().String()
		traceFlags = span.SpanContext().TraceFlags().String()
		ctx.Request = ctx.Request.WithContext(spanCtx)
	}
	return traceId, spanId, traceFlags
}
