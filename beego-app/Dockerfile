# Build Golang binary
# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
WORKDIR /srv/beego-app

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build main.go
# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
WORKDIR /srv/beego-app
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /srv/beego-app .
# 复制编译阶段里的config文件夹到目标目录
#OPY --from=builder /srv/beego-app/conf ./conf
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata
# 需暴露的端口
EXPOSE 8080
# 可外挂的目录
#VOLUME ["/go/kingProject/config","/go/kingProject/log"]

CMD ["./main"]

#docker build -t wangsenwangsen/beego-app:v1 .
#docker run -p 8080:8080 --name=beego-app  wangsenwangsen/beego-app:v1
#docker login
#docker images
#docker tag beego-app:latest  wangsenwangsen/beego-app:v1
#docker push wangsenwangsen/beego-app:v1
