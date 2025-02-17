# 使用官方 Golang 镜像作为构建环境
FROM golang:1.24.0-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件（如果有）
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod download

# 复制整个项目
COPY . .

# 构建 Go 应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用最小的 alpine 镜像作为运行时环境
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/main .

# 暴露应用端口
EXPOSE 8080

# 启动应用
CMD ["./main"]
