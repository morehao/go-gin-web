package ctrexample

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
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

	clientGone := ctx.Stream(func(w io.Writer) bool {
		_, err := fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.DateTime))
		if err != nil {
			return false
		}
		return true
	})
	if clientGone {
		glog.Infof(ctx, "[sseCtr.Time] Client disconnected during streaming")
	} else {
		glog.Infof(ctx, "[sseCtr.Time] Stream completed normally")
	}
	return
}

// TimeRaw Write写入实时时间流示例
func (ctr *sseCtr) TimeRaw(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	clientGone := ctx.Stream(func(w io.Writer) bool {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		sseData := fmt.Sprintf("id: %d\nevent: time\ndata: %s\n\n",
			time.Now().Unix(), currentTime)

		// 直接写入到 w
		_, err := w.Write([]byte(sseData))
		if err != nil {
			return false
		}

		time.Sleep(1 * time.Second)
		return true
	})
	if clientGone {
		glog.Infof(ctx, "[sseCtr.TimeRaw] Client disconnected during streaming")
	} else {
		glog.Infof(ctx, "[sseCtr.TimeRaw] Stream completed normally")
	}
}

// Process 模拟数据处理进度示例
func (ctr *sseCtr) Process(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	progress := 0

	clientGone := ctx.Stream(func(w io.Writer) bool {
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
	})
	if clientGone {
		glog.Infof(ctx, "[sseCtr.Process] Client disconnected during streaming")
	} else {
		glog.Infof(ctx, "[sseCtr.Process] Stream completed normally")
	}
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

	messageIndex := 0

	clientGone := ctx.Stream(func(w io.Writer) bool {
		if messageIndex < len(messages) {

			dataMsg := fmt.Sprintf(`{"id": %d, "text": "%s", "timestamp": "%s"}`,
				messageIndex+1,
				messages[messageIndex],
				time.Now().Format("15:04:05"))

			ctx.SSEvent("message", dataMsg)
			messageIndex++
			time.Sleep(2 * time.Second)
			return true
		} else {
			// 所有消息发送完毕
			ctx.SSEvent("end", `{"message": "Conversation ended"}`)
			return false
		}
	})
	if clientGone {
		glog.Infof(ctx, "[sseCtr.Chat] Client disconnected during streaming")
	} else {
		glog.Infof(ctx, "[sseCtr.Chat] Stream completed normally")
	}
}

// Raw 自定义格式的 SSE 示例
func (ctr *sseCtr) Raw(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	counter := 0

	clientGone := ctx.Stream(func(w io.Writer) bool {
		if counter < 10 {

			dataMsg := fmt.Sprintf(`{"count": %d, "message": "This is event #%d"}`, counter, counter)

			ctx.SSEvent("message", dataMsg)

			counter++
			time.Sleep(1 * time.Second)
			return true
		}
		ctx.SSEvent("message", `{"message": "Stream ended"}`)
		return false
	})
	if clientGone {
		glog.Infof(ctx, "[sseCtr.Raw] Client disconnected during streaming")
	} else {
		glog.Infof(ctx, "[sseCtr.Raw] Stream completed normally")
	}
}
