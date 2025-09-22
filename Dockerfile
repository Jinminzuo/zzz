# 二开推荐阅读[如何提高项目构建效率](https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/scene/build/speed.html)
# 选择构建用基础镜像（选择原则：在包含所有用到的依赖前提下尽可能体积小）。如需更换，请到[dockerhub官方仓库](https://hub.docker.com/_/golang?tab=tags)自行选择后替换。
# ------------------------
# 构建阶段
# ------------------------
FROM golang:1.17.1-alpine3.14 AS builder

WORKDIR /app

# 只先拷贝 go.mod 和 go.sum，提前下载依赖，提高缓存命中率
COPY go.mod go.sum ./
RUN go mod download

# 再拷贝其他源码文件
COPY . .

# 编译 Go 程序
RUN GOOS=linux go build -o main .

# ------------------------
# 运行阶段
# ------------------------
FROM alpine:3.13

# 安装 CA 证书
RUN apk add --no-cache ca-certificates

WORKDIR /app

# 拷贝构建好的二进制文件
COPY --from=builder /app/main /app/

# 如果有静态文件可以一起拷贝
# COPY --from=builder /app/index.html /app/

# 容器启动命令
CMD ["/app/main"]