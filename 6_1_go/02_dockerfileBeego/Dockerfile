# 构建一个基于ubuntu 的docker 定制镜像
# 基础镜像
FROM ubuntu

# 镜像作者
MAINTAINER zzy

# 增加国内源
#COPY sources.list /etc/apt/

RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

# 执行命令

RUN apt-get update
RUN apt-get install gcc libc6-dev git lrzsz -y
#将go复制解压到容器中
ADD go1.20.4.linux-amd64.tar.gz /usr/local/

# 定制环境变量

ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/root/go
ENV PATH=$GOPATH/bin/:$PATH

# 下载项目

RUN go get github.com/astaxie/beego

# 增加文件

COPY test.go /root/go/src/myTest/

# 定制工作目录

WORKDIR /root/go/src/myTest/

# 对外端口

EXPOSE 8080

# 运行项目

ENTRYPOINT ["go","run","test.go"]

#把sources.list和test.go文件放到这个目录中
#构建镜像
#docker build -t go-test:v0.1 .
#运行镜像
#docker run -p 8080:8080 -itd go-test:v0.1
#访问镜像，查看效果