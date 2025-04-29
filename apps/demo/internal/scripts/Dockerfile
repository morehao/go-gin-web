# ========================
# 第一阶段：构建 Go 应用
# ========================
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o /app/demo ./apps/demo/cmd

# ========================
# 第二阶段：构建最终镜像
# ========================
FROM scratch

# 从 alpine 镜像中复制必要的文件
COPY --from=alpine:3.19 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 设置时区
ENV TZ=Asia/Shanghai

# 从构建阶段复制二进制文件和配置文件
COPY --from=builder /app/demo /app/demo
COPY --from=builder /build/apps/demo/config/config.prod.yaml /build/apps/demo/config/config.yaml

# 设置工作目录
WORKDIR /build

# 设置配置文件路径环境变量
ENV CONFIG_PATH=/build/apps/demo/config/config.yaml

EXPOSE 8099

ENTRYPOINT ["/app/demo"]
