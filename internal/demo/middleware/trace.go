package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shutdown := stdoutProvider(ctx)
		defer shutdown()
		tracer := otel.Tracer("go-gin-web")
		_, span1 := tracer.Start(ctx.Request.Context(), "root")
		glog.Infof(ctx, "trace middleware1")
		span1.End()
		_, span2 := tracer.Start(ctx.Request.Context(), "root")
		glog.Infof(ctx, "trace middleware2")
		span2.End()
		ctx.Next()
	}
}

func stdoutProvider(ctx *gin.Context) func() {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)

	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exp)
	provider.RegisterSpanProcessor(bsp)

	return func() {
		if err := provider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}
