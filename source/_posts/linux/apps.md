---
title: linux 下常用的应用软件
date: 2015-03-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
categories:
  - os
  - linux
---


## ssh 相关
ssh 是linux下最为广泛使用的远程连接工具， 也是每个linux学习者必备的技能    
ssh最常用的一些操作包括 [ssh key转换、免密登录，连接管理](./apps/ssh.md) 也是使用者最应该知悉的内容

免密码ssh登录的设置 

1.确保你有公钥私钥对，确保安装了ssh 
如果没有可以用下面的命令生成 
ssh-keygen 
如果有私钥可以用下面的命令生成公钥 
ssh-keygen -y -f id_rsa > id_rsa.pub 
 
2.要把自己的公钥添加至目标机的 
.ssh/authorized_keys 
文件中去，authorized_keys的权限是600 
 

open ssh与  windows ppk相互转换
关键工具 puttygen
由私钥生成公钥如上

ssh 连接 linux
自动补全需要bash_completion
还需要在 .ssh/config 文件中记录 host user ip三个字段
另外，如果仅需要补全自己

ssh-copy-id 会自动加免密到目标机器

## 设置hostname

```
#!/bin/bash
# setting up hostenv
# @author winjeg@qq.com
#

### setting hostname now
hostname $1
echo $1 > /etc/hostname
echo -e "127.0.0.1\t$1$2\t$1" >> /etc/hosts

```

## openvpn

openvpn 

添加  /etc/openvpn/client/xxxx.conf

systemctl enable openvpn-client@jiaxing.service

设置免密
在conf中
auth-user-pass  xxx.txt

xxx.txt 中第一行用户名，第二行密码
