# Go应用程序Makefile

# 构建相关变量
APP =
BINARY = $(APP)
MAIN_DIR = ./apps/$(APP)/cmd
BUILD_DIR = ./output/build
VERSION = $(shell date +%Y%m%d%H%M%S)-$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# go命令的环境变量
GO_ENV = CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct

# Docker 相关变量
DOCKER_IMAGE = $(APP)

# 伪目标
.PHONY: all build clean run lint test swag docker-build docker-run help list-apps deps tidy

# 通用入口：清理、依赖、构建并运行
all: clean deps build run

# 定义函数：验证 APP 参数是否有效
define validate_app
	@if [ -z "$(APP)" ]; then \
		echo "❌ 请使用 APP=<名称> 指定要操作的应用程序，例如：make build APP=demo"; \
		exit 1; \
	fi
	@if [ ! -d "./apps/$(APP)" ]; then \
		echo "❌ 应用程序 '$(APP)' 不存在于 ./apps 目录下，请使用 make list-apps 查看可用应用"; \
		exit 1; \
	fi
endef

# 构建应用程序
build:
	$(call validate_app)
	@echo "正在构建应用程序 $(APP)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-X 'main.BuildVersion=$(VERSION)'" -o $(BUILD_DIR)/$(BINARY) $(MAIN_DIR)
	@echo "✅ 构建完成: $(BUILD_DIR)/$(BINARY)"

# 为特定环境构建（例如 Linux）
build-env:
	$(call validate_app)
	@echo "正在为 $(GO_ENV) 构建 $(APP)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO_ENV) go build -ldflags="-X 'main.BuildVersion=$(VERSION)'" -o $(BUILD_DIR)/$(BINARY) $(MAIN_DIR)
	@echo "✅ 构建完成: $(BUILD_DIR)/$(BINARY)"

# 清理构建产物
clean:
	@echo "🧹 正在清理构建目录..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ 已清理构建目录"

# 运行应用程序
run:
	$(call validate_app)
	@echo "🚀 正在运行应用程序 $(APP)..."
	@go run $(MAIN_DIR)

# 运行测试
test:
	$(call validate_app)
	@echo "🧪 正在运行测试..."
	@go test ./apps/$(APP)/internal/... -v

# 下载依赖项
deps:
	@echo "📦 正在下载依赖项..."
	@$(GO_ENV) go mod download
	@$(GO_ENV) go mod tidy
	@echo "✅ 依赖项已更新"

# 生成 Swagger 文档
swag:
	$(call validate_app)
	@echo "📚 正在生成 Swagger 文档..."
	@chmod +x ./scripts/swag.sh
	@./scripts/swag.sh $(APP)
	@echo "✅ Swagger 文档已生成"

# 构建 Docker 镜像
docker-build:
	$(call validate_app)
	@echo "🐳 正在构建 $(APP) 的 Docker 镜像..."
	@docker build -t $(DOCKER_IMAGE):latest -f ./apps/$(APP)/internal/scripts/Dockerfile .
	@echo "✅ Docker 镜像 $(DOCKER_IMAGE):latest 已构建完成"

# 运行 Docker 容器
docker-run:
	$(call validate_app)
	@echo "🚀 正在运行 $(APP) 容器..."
	@docker run -d \
		--name $(APP) \
		--add-host=host.docker.internal:host-gateway \
		-p 8099:8099 \
		$(DOCKER_IMAGE):latest
	@echo "✅ 容器 $(APP) 已启动，服务地址：http://localhost:8099"

# 列出所有可用的应用程序
list-apps:
	@echo "📂 可用的应用程序:"
	@ls -1 ./apps

# 运行代码检查工具
lint:
	@echo "🔍 正在运行代码检查工具..."
	@golangci-lint run ./...

# 显示帮助信息
help:
	@echo "🆘 可用命令:"
	@echo "  make                    - 清理、下载依赖并构建应用程序"
	@echo "  make build APP=<名称>    - 构建指定的应用程序"
	@echo "  make build-env APP=<名称> - 为特定环境构建"
	@echo "  make clean              - 清理构建产物"
	@echo "  make deps               - 下载依赖项"
	@echo "  make run APP=<名称>     - 运行指定的应用程序"
	@echo "  make test APP=<名称>    - 运行测试"
	@echo "  make swag APP=<名称>    - 生成 Swagger 文档"
	@echo "  make docker-build APP=<名称>  - 构建 Docker 镜像"
	@echo "  make docker-run APP=<名称> - 运行 Docker 容器"
	@echo "  make list-apps          - 列出所有可用的应用程序"
	@echo "  make lint               - 运行代码检查工具"
