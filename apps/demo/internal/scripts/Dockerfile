# ========================
# 第一阶段：构建 Go 应用
# ========================
FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# 设置工作目录
WORKDIR /build

# 拷贝项目完整代码
COPY ../../ .  # 注意：路径是相对于 Dockerfile 的位置，如在 apps/demo/internal/scripts

# 编译 demo 应用
RUN go build -ldflags="-s -w" -o /app ./apps/demo/cmd

# ========================
# 第二阶段：构建最终镜像
# ========================
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app /app

USER nonroot:nonroot
ENTRYPOINT ["/app"]
