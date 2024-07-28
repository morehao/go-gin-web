# Build stage
# 使用 Golang 的 alpine 镜像作为基础镜像，并命名为 builder
FROM golang:alpine AS builder

# 设置工作目录为 /opt/server
WORKDIR /opt/server

# 设置 Go 包代理
ENV GOPROXY=https://goproxy.cn,direct

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 将当前构建上下文的所有文件复制到工作目录
COPY . .

# 编译 Go 源代码，并将输出的可执行文件命名为 webserver
RUN go build -o webserver /opt/server/cmd/app/main.go

# Final stage
# 使用 alpine 镜像作为基础镜像
FROM alpine

# 设置工作目录为 /opt/server
WORKDIR /opt/server

# 从 builder 阶段复制 webserver 可执行文件到工作目录
COPY --from=builder /opt/server/webserver .

# 从构建阶段复制配置文件
COPY --from=builder /opt/server/internal/app/config/docker_config.yaml /opt/server/internal/app/config/config.yaml

# 暴露 8099 端口
EXPOSE 8099

# 使用环境变量配置
ENV APP_ENV=prod

# 设置容器启动时运行的命令
CMD ["./webserver", "-env", "${APP_ENV}"]
