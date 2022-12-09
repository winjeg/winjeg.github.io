---
title: Linux 用户相关的一些知识
date: 2014-03-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
categories:
  - os
  - linux
---



主要问题都是安全


譬如专门为 mysql tomcat 创建一个用户名
```
sudo groupadd mysql
sudo useradd -r -g mysql mysql
```

好控制权限，只给相关文件的权限

隔离

万一有 bug/被入侵 运气好也许还能靠权限控制避免进一步的恶劣影响。有些服务用户还设置了 nologin。

话说，搭车问下 windows 上面怎么以其他用户身份运行服务啊，现在是 system 身份运行 ftp 服务感觉很有压力

我司（截至 2017 年 6 月）的某个 linux 生产环境上因 struts2 远程执行漏洞种了 DDOS 木马,由于 tomcat 以 root 运行，导致无法清除木马，后直接迁移。

组策略设定允许某用户以服务身份登陆

在服务里面可以自己改服务启动所属的用户
不过首先你要自己新建一个，要不是就用 nt 自有的

命令好像也可以改，明天上班去试试