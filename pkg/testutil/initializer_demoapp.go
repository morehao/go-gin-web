package testutil

import (
	"fmt"
	"os"

	"github.com/morehao/go-gin-web/apps/demoapp/config"
	"github.com/morehao/go-gin-web/pkg/storages"
	"github.com/morehao/golib/glog"
)

// demoappInitializer 为 demoapp 应用提供初始化逻辑
type demoappInitializer struct {
	*baseInitializer
}

// newDemoappInitializer 创建 demoapp 应用的初始化器
func newDemoappInitializer() (Initializer, error) {
	base, err := newBaseInitializer(AppNameDemo)
	if err != nil {
		return nil, fmt.Errorf("create base initializer: %w", err)
	}

	return &demoappInitializer{
		baseInitializer: base,
	}, nil
}

// Initialize 实现 demoapp 应用的初始化逻辑
func (d *demoappInitializer) Initialize() error {
	// 1. 查找配置文件
	configPath, err := d.FindConfigPath()
	if err != nil {
		return fmt.Errorf("find config path: %w", err)
	}

	// 检查文件是否存在
	if _, err := os.Stat(configPath); err != nil {
		return fmt.Errorf("config file not found: %s, error: %w", configPath, err)
	}

	// 2. 加载配置（使用应用自己的配置加载函数）
	var panicErr interface{}
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicErr = r
			}
		}()
		config.LoadConfig(configPath)
	}()

	if panicErr != nil {
		return fmt.Errorf("load config failed: %v", panicErr)
	}

	// 3. 初始化日志
	logCfg, ok := config.Conf.Log["default"]
	if !ok {
		// 如果没有 default，尝试获取第一个配置
		for _, cfg := range config.Conf.Log {
			logCfg = cfg
			break
		}
	}
	if err := glog.InitLogger(&logCfg); err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	// 4. 初始化资源
	if err := d.initResources(); err != nil {
		return fmt.Errorf("init resources: %w", err)
	}

	return nil
}

// initResources 初始化各类资源（MySQL、Redis、ES 等）
func (d *demoappInitializer) initResources() error {
	// 初始化 MySQL
	if len(config.Conf.MysqlConfigs) > 0 {
		if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
			return fmt.Errorf("init mysql: %w", err)
		}
	}

	// 初始化 Redis
	if len(config.Conf.RedisConfigs) > 0 {
		if err := storages.InitMultiRedis(config.Conf.RedisConfigs); err != nil {
			return fmt.Errorf("init redis: %w", err)
		}
	}

	// 初始化 Elasticsearch
	if len(config.Conf.ESConfigs) > 0 {
		if err := storages.InitMultiEs(config.Conf.ESConfigs); err != nil {
			return fmt.Errorf("init elasticsearch: %w", err)
		}
	}

	return nil
}

// Close 实现 demoapp 应用的清理逻辑
func (d *demoappInitializer) Close() error {
	// 添加应用特定的清理逻辑
	
	// 执行基础清理
	return d.baseInitializer.Close()
}

