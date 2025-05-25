package ctrexample

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"go-gin-web/internal/apps/demoapp/dto/dtoexample"

	"github.com/gin-gonic/gin"
)

type SSECtr interface {
	Time(ctx *gin.Context)
	TimeRaw(ctx *gin.Context)
	Process(ctx *gin.Context)
	Chat(ctx *gin.Context)
	Raw(ctx *gin.Context)
}

type sseCtr struct {
}

var _ SSECtr = (*sseCtr)(nil)

func NewSSECtr() SSECtr {
	return &sseCtr{}
}

// Time 实时时间流示例
func (ctr *sseCtr) Time(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// 创建一个用于停止的通道
	clientGone := ctx.Request.Context().Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			fmt.Println("client gone")
			return false
		case t := <-ticker.C:
			_, err := fmt.Fprintf(w, "event: time\n")
			if err != nil {
				return false
			}
			_, err = fmt.Fprintf(w, "data: %s\n\n", t.Format(time.DateTime))
			if err != nil {
				return false
			}
			return true
		}
	})
}

// TimeRaw Write写入实时时间流示例
func (ctr *sseCtr) TimeRaw(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// 创建一个用于停止的通道
	clientGone := ctx.Request.Context().Done()

	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			fmt.Println("client gone")
			return false
		default:
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			sseData := fmt.Sprintf("id: %d\nevent: time\ndata: %s\n\n",
				time.Now().Unix(), currentTime)

			// 直接写入到 w
			_, err := w.Write([]byte(sseData))
			if err != nil {
				return false
			}

			// 刷新缓冲区确保数据立即发送
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}

			time.Sleep(1 * time.Second)
			return true
		}
	})
}

// Process 模拟数据处理进度示例
func (ctr *sseCtr) Process(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// 创建一个用于停止的通道
	clientGone := ctx.Request.Context().Done()

	progress := 0

	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		default:
			if progress <= 100 {
				// 发送进度更新
				ctx.SSEvent("progress", fmt.Sprintf(`{"progress": %d, "message": "Processing... %d%%"}`, progress, progress))
				progress += 10
				time.Sleep(500 * time.Millisecond)
				return true
			} else {
				// 完成时发送完成事件
				ctx.SSEvent("complete", `{"progress": 100, "message": "Task completed!"}`)
				return false
			}
		}
	})
}

// Chat 聊天消息流示例
func (ctr *sseCtr) Chat(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	messages := []string{
		"Hello! How can I help you today?",
		"I'm an AI assistant powered by SSE streaming.",
		"This demonstrates real-time message delivery.",
		"Each message arrives with a small delay.",
		"This simulates a natural conversation flow.",
		"SSE is perfect for chat applications!",
		"Thanks for trying this demo. Goodbye! 👋",
	}

	// 创建一个用于停止的通道
	clientGone := ctx.Request.Context().Done()

	messageIndex := 0

	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		default:
			if messageIndex < len(messages) {
				msg := dtoexample.SSEMessage{
					ID:    fmt.Sprintf("msg_%d", messageIndex+1),
					Event: "message",
					Data: fmt.Sprintf(`{"id": %d, "text": "%s", "timestamp": "%s"}`,
						messageIndex+1,
						messages[messageIndex],
						time.Now().Format("15:04:05")),
				}

				ctx.SSEvent("message", msg.Data)
				messageIndex++
				time.Sleep(2 * time.Second)
				return true
			} else {
				// 所有消息发送完毕
				ctx.SSEvent("end", `{"message": "Conversation ended"}`)
				return false
			}
		}
	})
}

// Raw 自定义格式的 SSE 示例
func (ctr *sseCtr) Raw(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// 创建一个用于停止的通道
	clientGone := ctx.Request.Context().Done()

	counter := 0

	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		default:
			if counter < 10 {
				// 方式3: 使用自定义结构体 + w 参数
				msg := dtoexample.SSEMessage{
					ID:    fmt.Sprintf("event_%d", counter),
					Event: "counter",
					Data:  fmt.Sprintf(`{"count": %d, "message": "This is event #%d"}`, counter, counter),
				}

				// 直接写入格式化的 SSE 数据到 w
				_, err := fmt.Fprint(w, msg.Format())
				if err != nil {
					return false
				}

				// 刷新缓冲区
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}

				counter++
				time.Sleep(1 * time.Second)
				return true
			}
			return false
		}
	})
}
