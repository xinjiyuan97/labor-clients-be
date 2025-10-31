# 构建阶段
FROM golang:1.24-alpine as builder

# 设置工作目录
WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache git

# 复制go.mod和go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建二进制文件
# CGO_ENABLED=0 禁用CGO，生成静态二进制文件
# GOOS=linux 目标操作系统为Linux
# -a 强制重新构建所有包
# -installsuffix cgo 使用不同的安装目录名
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 运行阶段
FROM alpine:latest

# 安装ca-certificates用于HTTPS请求
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY conf ./conf

# 暴露端口
EXPOSE 8888

# 运行应用
CMD ["/app/main", "-mode", "server", "-env", "prod"]

