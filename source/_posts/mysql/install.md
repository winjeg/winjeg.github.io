---
title: Centos7 安装配置 mysql5.6
toc: true
thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
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
