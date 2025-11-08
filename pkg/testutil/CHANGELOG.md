# Changelog

## v1.0.0 - 优化版本

### 🎉 新特性

1. **智能配置文件查找**
   - 自动从当前目录向上查找项目根目录（包含 go.mod 的目录）
   - 自动构建配置文件的绝对路径
   - 支持在任意目录运行测试，无需担心相对路径问题

2. **增强的上下文选项**
   - `WithUserID(uid uint)` - 设置用户ID
   - `WithRequestID(requestID string)` - 设置请求ID
   - `WithKeyValue(key, value)` - 设置自定义键值对

3. **更好的错误处理**
   - 使用错误链（error wrapping）提供完整的错误上下文
   - 详细的错误信息，便于快速定位问题

### 🔧 优化改进

1. **移除冗余代码**
   - 去除 `once.Do` 内部的额外锁（once.Do 本身已线程安全）
   - 移除重复的配置初始化调用
   - 简化初始化流程

2. **改进的命名**
   - `NewContext()` 替代 `NewCtx()`
   - `WithUserID()` 替代 `OptUid()`

3. **完善的文档**
   - 详细的函数注释
   - 完整的使用示例
   - README 使用指南

### 🐛 Bug 修复

- 修复测试执行时找不到配置文件的问题
- 通过自动查找项目根目录并设置绝对路径解决相对路径问题

### 📝 使用示例

```go
package mypackage_test

import (
    "testing"
    "github.com/morehao/go-gin-web/pkg/testutil"
)

func TestMyFunction(t *testing.T) {
    // 初始化测试环境（会自动找到配置文件）
    testutil.Init(testutil.AppNameDemo)
    
    // 创建测试上下文
    ctx := testutil.NewContext(
        testutil.WithUserID(12345),
        testutil.WithRequestID("req-001"),
    )
    
    // 执行测试...
}
```

### ⚠️ Breaking Changes

- 包名从 `test` 改为 `testutil`
- `NewCtx()` 改为 `NewContext()`
- `OptUid()` 改为 `WithUserID()`

### 🔄 迁移指南

1. 更新导入路径：
   ```go
   // 旧
   import "github.com/morehao/go-gin-web/pkg/test"
   
   // 新
   import "github.com/morehao/go-gin-web/pkg/testutil"
   ```

2. 更新 API 调用：
   ```go
   // 旧
   test.Init("demoapp")
   ctx := test.NewCtx(test.OptUid(123))
   
   // 新
   testutil.Init(testutil.AppNameDemo)
   ctx := testutil.NewContext(testutil.WithUserID(123))
   ```

