package testutil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/morehao/go-gin-web/apps/demoapp/config"
	"github.com/morehao/go-gin-web/pkg/storages"
	"github.com/morehao/golib/glog"
)

// demo 实现 demoapp 的测试初始化
type demo struct{}

// newDemo 创建 demoapp 初始化器
func newDemo() Initializer {
	return &demo{}
}

// Initialize 初始化 demoapp 测试环境
func (d *demo) Initialize() error {
	// 1. 初始化配置
	if err := d.initConfig(); err != nil {
		return fmt.Errorf("init config failed: %w", err)
	}

	// 2. 初始化日志
	if err := d.initLogger(); err != nil {
		return fmt.Errorf("init logger failed: %w", err)
	}

	// 3. 初始化资源（数据库、缓存等）
	if err := d.initResources(); err != nil {
		return fmt.Errorf("init resources failed: %w", err)
	}

	return nil
}

// initConfig 初始化配置文件
func (d *demo) initConfig() error {
	// 查找项目根目录（包含 go.mod 的目录）
	projectRoot, err := d.findProjectRoot()
	if err != nil {
		return fmt.Errorf("find project root failed: %w", err)
	}

	// 构建配置文件的绝对路径
	configPath := filepath.Join(projectRoot, "apps", AppNameDemo, "config", "config.yaml")
	
	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); err != nil {
		return fmt.Errorf("config file not found: %s, error: %w", configPath, err)
	}

	// 设置环境变量，让 config.InitConf() 能找到正确的配置文件
	if err := os.Setenv("APP_CONFIG_PATH", configPath); err != nil {
		return fmt.Errorf("set config path env failed: %w", err)
	}

	config.InitConf()
	return nil
}

// findProjectRoot 查找项目根目录（包含 go.mod 文件的目录）
func (d *demo) findProjectRoot() (string, error) {
	// 从当前工作目录开始查找
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get current directory failed: %w", err)
	}

	// 向上查找 go.mod 文件
	for {
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达根目录
			return "", fmt.Errorf("project root not found (no go.mod file)")
		}
		dir = parent
	}
}

// initLogger 初始化日志
func (d *demo) initLogger() error {
	logCfg, ok := config.Conf.Log["default"]
	if !ok {
		return fmt.Errorf("default log config not found")
	}

	if err := glog.InitLogger(&logCfg); err != nil {
		return fmt.Errorf("init logger failed: %w", err)
	}

	return nil
}

// initResources 初始化各种资源
func (d *demo) initResources() error {
	// 初始化 MySQL
	if len(config.Conf.MysqlConfigs) > 0 {
		if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
			return fmt.Errorf("init mysql failed: %w", err)
		}
	}

	// 初始化 Redis
	if len(config.Conf.RedisConfigs) > 0 {
		if err := storages.InitMultiRedis(config.Conf.RedisConfigs); err != nil {
			return fmt.Errorf("init redis failed: %w", err)
		}
	}

	// 初始化 Elasticsearch
	if len(config.Conf.ESConfigs) > 0 {
		if err := storages.InitMultiEs(config.Conf.ESConfigs); err != nil {
			return fmt.Errorf("init es failed: %w", err)
		}
	}

	return nil
}

// Close 清理资源
func (d *demo) Close() error {
	glog.Close()
	return nil
}

