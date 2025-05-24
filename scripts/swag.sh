#!/bin/bash

# 接收第一个参数作为模块名，比如 demoapp
MODULE=$1

# 检查是否传了模块名
if [ -z "$MODULE" ]; then
  echo "未指定模块名称，如 demo: $0 <module>"
  exit 1
fi

# 切换到模块目录
cd "internal/apps/${MODULE}" || {
  echo "模块目录 internal/apps/${MODULE} 不存在!"
  exit 1
}

echo "当前工作目录: $(pwd)"

# 执行 swag init
swag init \
  --parseDependency \
  --parseInternal \
  -g cmd/main.go \
  -o internal/docs

# 检查执行结果
if [ $? -eq 0 ]; then
  echo "Swagger 文档生成成功! 文件位置: internal/docs"
else
  echo "Swagger 文档生成失败!"
  exit 1
fi