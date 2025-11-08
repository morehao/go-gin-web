# TestUtil

测试工具包，用于简化单元测试和集成测试的初始化工作。

## 快速开始

```go
package myservice_test

import (
    "os"
    "testing"
    "github.com/morehao/go-gin-web/pkg/testutil"
)

func TestMain(m *testing.M) {
    // 初始化测试环境
    testutil.Initialize(testutil.AppNameDemo)
    
    // 运行测试
    code := m.Run()
    
    // 清理资源
    testutil.Close(testutil.AppNameDemo)
    
    os.Exit(code)
}

func TestYourFunction(t *testing.T) {
    // 创建测试上下文
    ctx := testutil.NewContext(
        testutil.WithUserID(123),
        testutil.WithMethod("POST"),
        testutil.WithJSON(),
    )
    
    // 你的测试代码
    result := myService.DoSomething(ctx)
    // ...
}
```

## 测试上下文选项

```go
ctx := testutil.NewContext(
    // 基础选项
    testutil.WithUserID(123),
    testutil.WithRequestID("test-req-001"),
    
    // HTTP 请求
    testutil.WithMethod("POST"),
    testutil.WithURL("/api/users"),
    testutil.WithQueryParams(map[string]string{
        "page": "1",
        "size": "10",
    }),
    
    // 内容类型
    testutil.WithJSON(),
    testutil.WithFormData(),
    
    // 认证
    testutil.WithBearerToken("token"),
    testutil.WithClientIP("127.0.0.1"),
    
    // 自定义
    testutil.WithKeyValue("key", "value"),
)
```

## 测试文件辅助

```go
// 打开测试数据文件（相对于当前测试文件）
file, err := testutil.OpenFile("testdata/test.json")

// 获取测试文件路径
path := testutil.TestFilePath("testdata/config.yaml")
```

## 添加新应用

### 前置要求

**确保你的应用配置包提供了 `LoadConfig(configPath string)` 函数：**

```go
// apps/myapp/config/config.go
package config

var Conf *Config

type Config struct {
    Log          map[string]glog.LogConfig `yaml:"log"`
    MysqlConfigs []dbmysql.MysqlConfig     `yaml:"mysql_configs"`
    RedisConfigs []dbredis.RedisConfig     `yaml:"redis_configs"`
    ESConfigs    []dbes.ESConfig           `yaml:"es_configs"`
}

// InitConf 运行时使用，自动查找配置文件
func InitConf() {
    configPath := getConfigPath()
    LoadConfig(configPath)
}

// LoadConfig 从指定路径加载配置（运行时和测试环境共用）
func LoadConfig(configPath string) {
    var cfg Config
    conf.LoadConfig(configPath, &cfg)
    Conf = &cfg
}
```

### 创建初始化器

参考 `initializer_demoapp.go`，为你的应用创建独立的初始化器：

```go
// pkg/testutil/initializer_myapp.go
package testutil

import (
    "fmt"
    "os"
    
    "github.com/myproject/apps/myapp/config"
    "github.com/myproject/pkg/storages"
    "github.com/morehao/golib/glog"
)

type myappInitializer struct {
    *baseInitializer
}

func newMyappInitializer() (Initializer, error) {
    base, err := newBaseInitializer("myapp")
    if err != nil {
        return nil, fmt.Errorf("create base initializer: %w", err)
    }
    return &myappInitializer{baseInitializer: base}, nil
}

func (m *myappInitializer) Initialize() error {
    // 1. 查找配置文件
    configPath, err := m.FindConfigPath()
    if err != nil {
        return fmt.Errorf("find config path: %w", err)
    }
    
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
        for _, cfg := range config.Conf.Log {
            logCfg = cfg
            break
        }
    }
    if err := glog.InitLogger(&logCfg); err != nil {
        return fmt.Errorf("init logger: %w", err)
    }
    
    // 4. 初始化资源
    if len(config.Conf.MysqlConfigs) > 0 {
        if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
            return fmt.Errorf("init mysql: %w", err)
        }
    }
    // ... 其他资源初始化
    
    return nil
}

func (m *myappInitializer) Close() error {
    return m.baseInitializer.Close()
}
```

### 注册应用

在 `init.go` 的 `initializerFuncMap` 中添加：

```go
var initializerFuncMap = map[string]InitializerFunc{
    AppNameDemo: newDemoappInitializer,
    "myapp":     newMyappInitializer,  // 添加这行
}
```

就这样！新应用就可以使用了：

```go
testutil.Initialize("myapp")
```

## 架构设计

```
testutil/
├── init.go                   # 核心初始化逻辑和全局 API
├── base_initializer.go       # 基础辅助功能（查找配置文件等）
├── initializer_demoapp.go    # demoapp 的独立初始化器
├── initializer_xxx.go        # 其他应用的独立初始化器
├── option.go                 # 上下文选项函数（20+）
├── file.go                   # 测试文件辅助函数
├── constant.go               # 常量定义
└── README.md                 # 文档
```

### 设计原则

1. **与应用配置复用**:
   - 每个应用使用自己的 `config.LoadConfig()`
   - 运行时和测试环境使用相同的配置加载逻辑
   - 无需 `CommonConfig` 这样的中间层

2. **差异化实现**:
   - `baseInitializer` 只提供辅助方法（查找配置文件路径）
   - 每个应用完全控制自己的初始化逻辑
   - 类似 testutils 的 `kecoreInitializer` 和 `kechatInitializer`

3. **选项模式**:
   - `NewContext()` 使用函数选项模式
   - 20+ 实用选项函数
   - 灵活且易于扩展

4. **单例模式**:
   - 每个应用只初始化一次
   - 线程安全的全局状态管理

## 配置文件要求

配置文件位置：`apps/{appName}/config/config.yaml`

```yaml
# 日志配置
log:
  default:
    service: myapp
    level: info
    writer: file
    dir: ../../../log

# MySQL 配置（可选）
mysql_configs:
  - addr: 127.0.0.1:3306
    user: root
    password: password
    database: mydb

# Redis 配置（可选）
redis_configs:
  - service: myapp
    addr: 127.0.0.1:6379

# ES 配置（可选）
es_configs:
  - service: myapp
    addr: http://127.0.0.1:9200
```

