# 设置基础镜像为Golang 1.20
FROM golang:1.20

# 设置工作目录
WORKDIR /msdm

# 将本地文件复制到容器中
COPY . .

RUN go env -w GO111MODULE=auto

# 安装依赖
RUN go mod download

# 构建可执行文件
RUN go build -o msdm

# 运行可执行文件
CMD ["./msdm"]
