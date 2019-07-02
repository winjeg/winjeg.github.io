---
title: 在docker下运行MySQL
date: 2019-07-02 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59422808-880b5180-8e03-11e9-9dfe-ff8a9a024be7.png
tags:
  - MySQL
  - database
  - docker
categories:
  - storage
  - database
---

## 为啥要使用docker运行MySQL

熟悉docker和k8s的人都知道， 我们使用容器化技术，是为了方便我们运行某个服务，我们用docker 运行Mysql， 并不是因为mysql在docker比在物理机或者kvm上运行的更好，配置起来更简单。  
我们使用docker 主要还是因为，我们在不关注太多MySQL本身的东西的时候， 单纯想快速简单的启动一个MySQL服务的时候能够做到，分钟级别即可完成。  
这相对于传统级别的从安装到配置动则半小时到几个小时的工作量来说，已经非常简单方便了。

## 怎么使用Docker 去运行一个MySQL服务

在运行之前，你首先要装docker， 安装docker非常简单，只需要一路下一步下一步就可以完成， linux或者mac野只需要一些安装命令即可搞定。  
当然安装docker可以不仅仅用于MySQL用途， 也可以用来它方便其他任何你想需要运行的服务的部署和运行。  

### 下载镜像

```bash
docker pull mysql:5.6
```
这里mysql是镜像名称， 5.6 是mysql的版本。 一半这些镜像都默认是比较官方维护的， 在docker没有进行特殊设置的情况下可以信任。

### 运行

```bash
docker run -p 3306:3306 --name mysql -v /opt/docker_v/mysql/conf:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 -d imageID
```
上述命令有如下解释：
1. `-p 3306:3306 ` 是指把容器的3306端口映射到本地的3306端口
2. `--name` 制定容器的名称, 我们下次要操作的时候可以制定名称即可， 如 `docker start mysql`
3. `-v` 把本地文件夹与docker文件夹进行一个映射
4. `-e ` 传入容器一些环境变量， 这里传入的是MySQL运行的时候需要的root密码
5. `-d` 后台方式运行

### 常见的问题
1. docker 守护进程没有开启  
    解决方案： 开启docker后台进程 如 `service docker start`或者 `systemctl start docker.service`
2. mysql 运行的端口过小导致的没有权限  
    解决方案：请在大于1024 的端口上运行MySQL
3. 系统环境是 Windows 10 Pro，Docker 版本 18.03.1-ce，电脑开机之后第一次运行 docker run 的时候会遇到这个错误

```
C:\Program Files\Docker\Docker\Resources\bin\docker.exe: Error response from daemon: driver failed programming external connectivity on endpoint app (36577729ce7d4d2dddefb7fddd32521ea66958cf824138804b02ffb3c98452f3): Error starting userland proxy: mkdir /port/tcp:0.0.0.0:3306:tcp:172.17.0.2:3306
: input/output error.
```

解决方案： 重启docker服务




