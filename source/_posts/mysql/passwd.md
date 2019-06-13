---
title: Mysql 如何找回Root密码
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59422808-880b5180-8e03-11e9-9dfe-ff8a9a024be7.png
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---

## 1  修改 /etc/my.cnf
加上
```
[mysqld]
skip-grant-tables
```

## 2. 重启mysql 
```
service mysql restart
```

## 3. 更改root密码
```sql
mysql> USE mysql ; 
mysql> UPDATE user SET Password = password ( 'new-password' ) WHERE User = 'root' ; 
mysql> flush privileges ; 
mysql> quit
```
如果版本高于5.6 则是如下字段存着密码
```
authentication_string
```
## 4.将MySQL的登录设置修改回来 

## 5. 重启mysql


