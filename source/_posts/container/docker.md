---
title: docker 简易笔记
date: 2018-04-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - docker
  - container
categories:
  - container
  - docker
---



## 基本概念

三个基本概念： Image, Contrainer, Repository
### Image
只读模板，预装一些东西

### Container
启动、开始、停止、删除 -- 隔离
启动时，在镜像上层创建一层可写层

### Repository
用来存放images

## Docker 命令

```
sudo 

# 启动docker服务
systemctl start docker.service

# 列出本地镜像
docker images

# 搜索镜像
docker search +name

docker run -t -i centos /bin/bash

docker rm 
docker rmi
docker ps
docker save filename
docker load --input filename
docker logs 查看容器的输出信息
docker ps   查看容器的状态信息

```


## Docker file
```
# 以centos latest 作为基础
FROM  centos:latest

# 维护者
MAINTAINER wenwen

# 执行命令
RUN yum install vi

#

ADD /home/wenwen/1.txt

EXPOSE port

CMD nano

#

docker build -t "" {dir}
```

##  docker-registry
