package context

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type RespWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w RespWriter) WriteString(s string) (int, error) {
	if w.Body != nil {
		w.Body.WriteString(s)
	}
	return w.ResponseWriter.WriteString(s)
}

func (w RespWriter) Write(b []byte) (int, error) {
	if w.Body != nil {
		w.Body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}
