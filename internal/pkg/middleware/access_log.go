package middleware

import (
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func AccessLog() gin.HandlerFunc {
	traceProvider := trace.NewTracerProvider()
	otel.SetTracerProvider(traceProvider)
	tr := traceProvider.Tracer("gin-server")
	return func(c *gin.Context) {
		traceId, spanId, traceFlags := getTraceInfo(c, tr)
		c.Set(glog.KeyTraceId, traceId)
		c.Set(glog.KeyTraceFlags, traceFlags)
		c.Set(glog.KeySpanId, spanId)
		c.Set(glog.KeyFERequestId, c.Request.Header.Get(glog.KeyFERequestId))
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
