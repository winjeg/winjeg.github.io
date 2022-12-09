---
title: Centos7 安装配置 mysql5.6
date: 2016-05-13 15:14:11

toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---

### 安装
```
wget http://dev.mysql.com/get/mysql57-community-release-el7-8.noarch.rpm
sudo yum -y localinstall mysql57-community-release-el7-8.noarch.rpm
sudo yum -y install mysql-community-server
sudo systemctl start mysqld
sudo systemctl enable mysqld
sudo systemctl daemon-reload

```

### 找到 root 密码
```
vim /var/log/mysqld.log 
```
### 更改root密码
```sql
set password=PASSWOR('djaskd');
```

### 更改数据文件及binlog 的位置（统一用）
```
systemctl stop mysqld
mkdir -p /db/mysql
mv /var/lib/mysql /db/mysql/
mv /db/mysql/mysql /db/mysql/data
chown -R mysql:mysql /db/mysql
```

### my.cnf 配置
```
[client]
port        = 3306
socket      = /var/run/mysqld/mysqld.sock

[mysqld_safe]
socket      = /var/run/mysqld/mysqld.sock
nice        = 0

[mysqld]
user        = mysql
pid-file    = /var/run/mysqld/mysqld.pid
socket      = /var/run/mysqld/mysqld.sock
port        = 3306
basedir     = /usr
datadir     = /db/mysql/data
tmpdir      = /tmp
skip-external-locking
skip-name-resolve
default-storage-engine=INNODB
character-set-server=utf8
collation-server=utf8_general_ci
lower_case_table_names=1
bind-address        = 0.0.0.0
max_allowed_packet  = 16M
thread_stack        = 192K
thread_cache_size       = 8
max_connections        = 1000
query_cache_limit   = 1M
query_cache_size        = 16M
log_error = /var/log/mysqld.log

server-id       = 1
binlog_format   = ROW
log_bin         = /db/mysql/mysql-bin.log
expire_logs_days    = 10
max_binlog_size         = 500M

```
### 启动mysql server
```
systemctl start mysqld
```


### ubuntu 问题

我想，你一定是从seo/seo.html" target="_blank">搜索引擎搜索这个标题进来的！你一定是想改变mysql默认安装的数据目录！
你已经修改了my.cnf中的datadir的值
首先是查看数据库日志
mysqld started
[Warning] Can't create test file xxx.lower-test 
[Warning] Can't create test file xxx.lower-test 
/usr/libexec/mysqld: Can't change dir to '/xxx' (Errcode: 13) 
[ERROR] Aborting
 
你已经chown和chmod了数次新数据目录或者其父路径的属主和权限
你无数次地试图service mysql start，或者 /etc/init.d/mysql start，以及mysql_install_db！
恭喜你看见这篇文章，我在被系统坑了几个小时之后，找到了解决的方法。
这个原因有二，其中任意的一个原因都会造成你被系统告知这个warning。如果你不是一个专业的linux系统安全工程师，或者你只是个PHP程序员，并没有对系统安全有深入的研究，你就不会太容易找到它的答案。
第一，selinux，记得当年念书时，字符界面安装redhat（很古老的操作系统么。。。）的时候，有这么一个选项，通常大家都听取前辈的建议，改变默认值以不安装它。但如果你恰好要操作的这台机器开着selinux，它确实能够使你的mysql无法在新目标位置进行mysql_install_db的操作，并爆出标题所示的警告。一个简单的解决办法是使用命令暂时关闭selinux，以便让你的操作可以继续下去
```
setenforce 0
```
但最好使用一个永久方法，以便在重启后继续不要这货。
修改/etc/selinux/config文件中设置SELINUX=disabled ，然后重启或等待下次重启。
第二，apparmor，这个坑爹货和selinux一样的坑爹，它也对mysql所能使用的目录权限做了限制
在 /etc/apparmor.d/usr.sbin.mysqld 这个文件中，有这两行，规定了mysql使用的数据文件路径权限

```
/var/lib/mysql/ r,
/var/lib/mysql/** rwk,
```
你一定看到了，/var/lib/mysql/就是之前mysql安装的数据文件默认路径，apparmor控制这里mysqld可以使用的目录的权限
我想把数据文件移动到/data/mysql下，那么为了使mysqld可以使用/data/mysql这个目录，照上面那两条，增加下面这两条就可以了
```
/data/mysql/ r,
/data/mysql/** rwk,
```
重启apparmor，/etc/inid.d/apparmor restart
之后，就可以顺利地干你想干的事儿了！