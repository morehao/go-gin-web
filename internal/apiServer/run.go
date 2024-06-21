package apiServer

import (
	"context"
	"go-gin-web/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/glog"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	initLogErr := glog.InitZapLogger(&glog.LoggerConfig{
		ServiceName: "go-gin-web",
		Level:       "debug",
		InConsole:   true,
		LogDir:      "./log",
		ExtraKeys:   []string{glog.KeyFERequestId, glog.KeyTraceId, glog.KeySpanId, glog.KeyTraceFlags},
	})
	if initLogErr != nil {
		panic(initLogErr)
	}

	// Ping test
	r.GET("/ping", middleware.AccessLog(), func(c *gin.Context) {
		// traceInfo := glog.GetTraceInfo(c)

		// c.Set(glog.KeyTraceId, traceInfo.TraceId)
		// c.Set(glog.KeySpanId, traceInfo.SpanId)
		// c.Set(glog.KeyTraceFlags, traceInfo.TraceFlags)
		glog.Info(c, "ping1")
		glog.Info(c, "ping2")
		glog.Infof(c, "ping%d", 3)
		glog.Warn(c, "ping4")
		glog.Errorf(c, "ping%d", 5)
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		glog.Info(c, "user1")
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	// }))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func Run() {
	router := setupRouter()
	glog.Info(context.Background(), "apiServer run success, port:8080")
	if err := router.Run(":8080"); err != nil {
		glog.Error(context.Background(), "apiServer run fail, port:8080")
	}
}
