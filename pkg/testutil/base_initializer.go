package testutil

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/morehao/golib/glog"
)

// baseInitializer 提供基础的初始化器辅助功能
// 每个应用的初始化器可以嵌入此结构体以复用通用方法
type baseInitializer struct {
	appName string
}

// newBaseInitializer 创建基础初始化器
func newBaseInitializer(appName string) (*baseInitializer, error) {
	if appName == "" {
		return nil, fmt.Errorf("app name is empty")
	}

	return &baseInitializer{
		appName: appName,
	}, nil
}

// FindConfigPath 查找配置文件路径
// 使用 runtime.Caller 自动定位项目根目录
func (b *baseInitializer) FindConfigPath() (string, error) {
	// 获取当前文件路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot get runtime caller info")
	}

	// 当前文件在 pkg/testutil/ 目录下
	// 项目根目录是往上两级目录
	pkgDir := filepath.Dir(filename)                  // pkg/testutil
	projectRoot := filepath.Dir(filepath.Dir(pkgDir)) // 项目根目录

	// 构建配置文件路径：{项目根目录}/apps/{appName}/config/config.yaml
	configPath := filepath.Join(projectRoot, "apps", b.appName, "config", "config.yaml")
	return configPath, nil
}

// Close 实现基础的清理逻辑
func (b *baseInitializer) Close() error {
	// 关闭日志
	glog.Close()
	return nil
}

