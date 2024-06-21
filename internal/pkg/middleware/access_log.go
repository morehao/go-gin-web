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
		// rCtx := c.Request.Context()
		// spanCtx, span := tr.Start(rCtx, c.Request.URL.Path)
		// c.Request = c.Request.WithContext(spanCtx)
		// defer span.End()
		// c.Set(glog.KeyTraceId, span.SpanContext().TraceID().String())
		// c.Set(glog.KeyTraceFlags, span.SpanContext().TraceFlags().String())
		// c.Set(glog.KeySpanId, span.SpanContext().SpanID().String())
		traceId, spanId, traceFlags := getTraceInfo(c, tr)
		c.Set(glog.KeyTraceId, traceId)
		c.Set(glog.KeyTraceFlags, traceFlags)
		c.Set(glog.KeySpanId, spanId)
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
		traceFlagsByte := byte(0)
		if traceFlags != "" {
			decoded, err := hex.DecodeString(traceFlags)
			if err == nil && len(decoded) > 0 {
				traceFlagsByte = decoded[0]
			}
		}
		spanContextCfg := oteltrace.SpanContextConfig{
			TraceID: newTraceId,
			// SpanID:     newSpanId,
			TraceFlags: oteltrace.TraceFlags(traceFlagsByte),
		}
		if spanId != "" {
			newSpanId, _ := oteltrace.SpanIDFromHex(spanId)
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
