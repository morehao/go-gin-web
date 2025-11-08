# TestUtil

测试工具包，提供测试环境初始化和上下文创建功能。

## 功能特性

- ✅ 自动初始化测试环境（配置、日志、数据库等）
- ✅ 提供测试用的 Gin Context 创建
- ✅ 支持多种上下文配置选项
- ✅ 优化的错误处理和日志输出
- ✅ 基于 sync.Once 的单例初始化

## 快速开始

### 基本使用

```go
package mypackage_test

import (
    "testing"
    
    "github.com/morehao/go-gin-web/pkg/testutil"
    "github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
    // 初始化测试环境（只会执行一次）
    testutil.Init(testutil.AppNameDemo)
    
    // 创建测试上下文
    ctx := testutil.NewContext()
    
    // 执行测试...
}
```

### 使用上下文选项

```go
func TestWithUserID(t *testing.T) {
    testutil.Init(testutil.AppNameDemo)
    
    // 创建带用户ID的上下文
    ctx := testutil.NewContext(
        testutil.WithUserID(12345),
    )
    
    // 执行测试...
}

func TestWithMultipleOptions(t *testing.T) {
    testutil.Init(testutil.AppNameDemo)
    
    // 使用多个配置选项
    ctx := testutil.NewContext(
        testutil.WithUserID(12345),
        testutil.WithRequestID("req-001"),
        testutil.WithKeyValue("custom_key", "custom_value"),
    )
    
    // 执行测试...
}
```

### 资源清理

```go
func TestMain(m *testing.M) {
    // 初始化
    testutil.Init(testutil.AppNameDemo)
    
    // 运行测试
    code := m.Run()
    
    // 清理资源
    testutil.Done()
    
    os.Exit(code)
}
```

## API 说明

### 初始化函数

- `Init(appName string)`: 初始化测试环境（线程安全，只会执行一次）
- `Done()`: 清理测试环境资源
- `MustInit(appName string) error`: 初始化并返回错误（用于需要更精细控制的场景）

### 上下文创建

- `NewContext(opts ...Option) *gin.Context`: 创建测试用的 Gin 上下文

### 配置选项

- `WithUserID(uid uint)`: 设置用户ID
- `WithRequestID(requestID string)`: 设置请求ID
- `WithKeyValue(key string, value interface{})`: 设置自定义键值对

## 优化改进

相比旧的 `test` 包，`testutil` 包做了以下优化：

1. **更清晰的命名**：`NewContext` 替代 `NewCtx`，`WithUserID` 替代 `OptUid`
2. **移除冗余代码**：
   - 去除了 `once.Do` 内部的额外锁（once.Do 本身已经是线程安全的）
   - 移除了 `preInit()` 中重复的 `config.InitConf()` 调用
   - 简化了配置初始化流程
3. **改进的错误处理**：
   - 使用 `fmt.Errorf` 和 `%w` 包装错误，提供更好的错误链
   - 提供更详细的错误信息，便于调试
4. **智能的配置文件查找**：
   - 自动从当前目录向上查找项目根目录（go.mod 所在目录）
   - 自动构建配置文件的绝对路径
   - 支持在任意目录运行测试，无需担心相对路径问题
5. **增强的功能**：
   - 添加更多上下文配置选项（RequestID、自定义键值对等）
   - 添加 `MustInit()` 函数用于更精细的错误控制
6. **完善的文档**：
   - 添加详细的函数注释
   - 提供完整的使用示例
   - 包含使用说明文档
7. **健壮性提升**：
   - 添加配置检查，避免空指针等错误
   - 资源初始化前检查配置是否存在
   - 更好的 panic 信息提示

## 注意事项

1. `Init()` 基于 `sync.Once` 实现，在整个测试生命周期中只会执行一次
2. 建议在 `TestMain` 中调用 `Done()` 来清理资源
3. 配置文件自动查找：
   - 从当前目录向上查找项目根目录（包含 go.mod 的目录）
   - 自动构建配置文件路径：`{项目根目录}/apps/demoapp/config/config.yaml`
   - 无需手动设置配置文件路径，支持在任意目录运行测试

