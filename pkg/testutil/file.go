package testutil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// OpenFile 打开测试数据文件
// relPath 为相对路径，相对于调用此函数的测试文件所在目录
// 
// 使用示例：
//
//	file, err := testutil.OpenFile("testdata/test.json")
func OpenFile(relPath string) (*os.File, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get caller info")
	}
	fullPath := filepath.Join(filepath.Dir(filename), relPath)
	return os.Open(fullPath)
}

// TestFilePath 获取测试数据文件的完整路径
// relPath 为相对路径，相对于调用此函数的测试文件所在目录
// 
// 使用示例：
//
//	path := testutil.TestFilePath("testdata/test.json")
func TestFilePath(relPath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get caller info")
	}
	return filepath.Join(filepath.Dir(filename), relPath)
}

// generateRequestID 生成唯一的请求ID
func generateRequestID() string {
	timestamp := time.Now().UnixNano()
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("test-%d", timestamp)
	}
	return fmt.Sprintf("test-%d-%s", timestamp, hex.EncodeToString(b))
}

