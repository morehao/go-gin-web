#!/bin/bash

# 接收第一个参数作为模块名，比如 demoapp
MODULE=$1

# 检查是否传了模块名
if [ -z "$MODULE" ]; then
  echo "未指定模块名称，如 demo: $0 <module>"
  exit 1
fi

APP_DIR="apps/${MODULE}"
MAIN_FILE="apps/${MODULE}/server.go"
DOCS_DIR="${APP_DIR}/docs"

# 检查目录和入口文件是否存在
if [ ! -d "$APP_DIR" ]; then
  echo "模块目录 ${APP_DIR} 不存在!"
  exit 1
fi

if [ ! -f "$MAIN_FILE" ]; then
  echo "入口文件 ${MAIN_FILE} 不存在!"
  exit 1
fi

# 切换到项目根目录（假设脚本始终从项目根运行）
echo "当前工作目录: $(pwd)"
echo "入口文件: ${MAIN_FILE}"
echo "文档输出目录: ${DOCS_DIR}"

# 执行 swag init
swag init \
  --parseDependency \
  --parseInternal \
  -g "${MAIN_FILE}" \
  -o "${DOCS_DIR}"

# 检查执行结果
if [ $? -eq 0 ]; then
  echo "Swagger 文档生成成功! 文件位置: ${DOCS_DIR}"
else
  echo "Swagger 文档生成失败!"
  exit 1
fi
