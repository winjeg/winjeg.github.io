---
title: vsftpd的简介和使用
date: 2013-11-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
  - ftp
categories:
  - application
---

## VSFTPD - 一个强大的FTP服务软件
---
vsftpd是一个非常好用的ftp软件，它非常的安全高效， 可配置性特别高，在非Windows系统很受个人和企业的欢迎。  
当然本文主要关注的还是教大家如何使用它， 我们下面将从以下几个方面详细讲述vsftpd的安装与使用。  

## 安装

安装其实是非常简单的事情， 我见过的大部分主流发行版中， 都可以一句命令安装。  

### `Ubuntu` 系列
```
sudo apt-get install vsftp
```

### `Fedora/Redhat` 系列
```bash
sudo yum install vsftp
```

### `Archlinux` 系列
```bash
sudo pacman -S vsftp
```

## 启动服务
```bash
# 开启 vsftpd 服务开机启动
systemctl enable vsftpd.service
# 开启 (修改配置完毕， 请使用  `restart` 来确保配置生效)
systemctl start vsftpd.service
```

## 配置
vsftp 配置对于喜欢GUI的人来讲可能不是那么友好， 但其实配置内容是非常易读的。 而且几乎所有配置项均有详细的注释。

### 配置文件位置：`/etc/vsftpd.conf`
```bash
vim /etc/vsftpd.conf
```

### 配置文件详解
```properties
# 是否监听IPV6端口，开启后可以用IPV6的方式访问ftp站点
listen_ipv6=YES

# 是否开启FTP目录消息
dirmessage_enable=YES
# 欢迎消息
ftpd_banner=Welcome to blah FTP service.

# 是否使用本地时间， 如不使用，默认为GMT
use_localtime=YES
# 推荐使用一个独立的用户来运行ftp服务
nopriv_user=ftpsecure
# 是否允许递归列出
ls_recurse_enable=YES

# 是否允许本地用户切换到本地目录
chroot_local_user=YES
chroot_local_user=YES
chroot_list_enable=YES
# (default follows)
chroot_list_file=/etc/vsftpd.chroot_list

# SSL相关的配置
rsa_cert_file=/etc/ssl/certs/ssl-cert-snakeoil.pem
rsa_private_key_file=/etc/ssl/private/ssl-cert-snakeoil.key
ssl_enable=NO


# 禁止一些邮件地址
deny_email_enable=YES
# (default follows)
banned_email_file=/etc/vsftpd.banned_emails

# ascii 方式上传下载
ascii_upload_enable=YES
ascii_download_enable=YES

local_enable=YES
# 是否允许任意形式的写命令
write_enable=YES
# 设置本地权限 （022 是Linux基础常识，不知道什么意思可以自行搜索）
local_umask=022
# 是否允许匿名用户上传， 只在全局写开启的情况下才有效，另外Linux中也要设置好相应路径的权限

# 匿名访问相关
anonymous_enable=NO
# 是否允许本地用户登录
anon_upload_enable=YES
# 是否允许匿名用户创建目录
anon_mkdir_write_enable=YES

# 超时时间
# 会话
idle_session_timeout=600
# 数据连接
data_connection_timeout=120


# 开启上传下载日志 (下一行为日志位置)
xferlog_enable=YES
xferlog_file=/var/log/vsftpd.log
xferlog_std_format=YES

# 使用20端口连接
connect_from_port_20=YES
# 是否把上传的文件更改所有者 （下一行为所有者）
chown_uploads=YES
chown_username=whoever

```



## 一些概念
vsftpd的用户权限其实是跟Linux的权限是紧密相关的， `chmod` 等基础linux指令会直接影响ftp的读写权限等。
