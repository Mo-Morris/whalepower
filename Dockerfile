# 第一阶段：构建阶段
FROM m.daocloud.io/docker.io/golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 拷贝到工作目录
COPY go.mod go.sum ./


RUN go env -w GOPROXY=https://goproxy.cn,direct
# 下载依赖
RUN go mod download

# 将源代码拷贝到工作目录
COPY . .

# 编译项目
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o whalepower

# 第二阶段：运行阶段
FROM m.daocloud.io/docker.io/alpine:3.18

# 创建一个非root用户来运行应用
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 将二进制文件从编译阶段复制到运行阶段
COPY --from=builder /app/whalepower /usr/local/bin/whalepower
COPY --from=builder /app/cfg /etc/whalepower/cfg


# 切换到非root用户
USER appuser

# 运行应用
ENTRYPOINT [ "whalepower","serve","--config", "/etc/whalepower/cfg/config.yaml"]