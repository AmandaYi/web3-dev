FROM ubuntu:14.04

MAINTAINER zzy

RUN mkdir hello
RUN mkdir world

RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update

RUN apt-get install nginx -y

EXPOSE 80