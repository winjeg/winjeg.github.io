---
title: ssh 周边知识
date: 2013-09-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
  - ssh
categories:
  - os
  - linux
---

## SSH 密钥生成与转换
1. 确保你有公钥私钥对，确保安装了ssh 
如果没有可以用下面的命令生成 
```
ssh-keygen 
```
由私钥生成公钥
```
ssh-keygen -y -f id_rsa > id_rsa.pub 
```
2. open ssh与  windows ppk相互转换
关键工具 `puttygen.exe`

## ssh免密码登录
1. 要把自己的公钥添加至目标机的
```
.ssh/authorized_keys
```
文件中去，`authorized_keys` 的权限是 `600` 

2. `ssh-copy-id` + host
会自动加免密到目标机器如
```
ssh-copy-id user@host
```

## ssh 连接管理
###  ssh 自动补全工具
 自动补全是：ssh 连接 linux 自动补全需要bash_completion
 还需要在 `~/.ssh/config` 文件中记录 三个字段如下
 ```
 Host   alias-of-the-host
 User   username
 Ip     10.1.1.2
 ```

### ssh远程管理工具
- 在 `.ssh/config` 中配置的各种服务
- 一些terminal自带的一些ssh管理工具， 如deeepin的terminal
- remmina 远程管理工具