package testutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/morehao/go-gin-web/pkg/testutil"
)

// TestMain 展示标准的测试主函数写法
func TestMain(m *testing.M) {
	// 初始化测试环境
	testutil.Initialize(testutil.AppNameDemo)

	// 运行所有测试
	code := m.Run()

	// 清理资源
	testutil.Close(testutil.AppNameDemo)

	os.Exit(code)
}

// Example_basicUsage 基础用法示例
func Example_basicUsage() {
	// 创建测试上下文
	ctx := testutil.NewContext(
		testutil.WithUserID(123),
		testutil.WithRequestID("test-request-001"),
	)

	fmt.Printf("User ID: %v\n", ctx.Value("user_id"))
	// Output: User ID: 123
}

// Example_httpRequest HTTP 请求模拟示例
func Example_httpRequest() {
	ctx := testutil.NewContext(
		testutil.WithMethod("POST"),
		testutil.WithURL("/api/users"),
		testutil.WithJSON(),
		testutil.WithBearerToken("test-token"),
		testutil.WithQueryParams(map[string]string{
			"page": "1",
			"size": "10",
		}),
	)

	fmt.Printf("Method: %s\n", ctx.Request.Method)
	fmt.Printf("URL: %s\n", ctx.Request.URL.Path)
	fmt.Printf("Query: %s\n", ctx.Request.URL.RawQuery)
	// Output: Method: POST
	// URL: /api/users
	// Query: page=1&size=10
}

// Example_testFile 测试文件辅助函数示例
func Example_testFile() {
	// 获取测试文件路径
	path := testutil.TestFilePath("testdata/test.json")
	fmt.Printf("File path exists: %t\n", path != "")
	// Output: File path exists: true
}

// Example_registerApp 动态注册应用示例
func Example_registerApp() {
	// 动态注册新应用
	testutil.RegisterApp("newapp", func() (testutil.Initializer, error) {
		// 这里应该返回应用的初始化器
		// 通常你会创建一个类似 newMyappInitializer 的函数
		return nil, fmt.Errorf("not implemented")
	})
	fmt.Println("App registered")
	// Output: App registered
}

