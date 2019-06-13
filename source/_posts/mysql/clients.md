---
title: 常见的MySQL客户端
date: 2016-07-14 15:14:11

toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---


## 驱动/程序连接
### java 
如果你是maven项目，则只需要引入以下依赖
```xml
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>5.1.38</version>
</dependency>
```
如果不是， 可以下载Jar包直接使用

### go
```bash
go get -u github.com/go-sql-driver/mysql
```
对于go项目直接引入依赖即可

### 其他
由于作者我不熟悉其他语言， 不在此赘述其他语言的客户端驱动是怎么样的，但是对于大部分语言来说都有比较知名的驱动。


##桌面客户端
###	MySQL Workbench (Official)
MySQL 是官方的MySQL客户端，它是开源免费的，有开源社区跟Oracle的支持，功能也是比较的完善，易用性比Navicat 稍差，功能比较强大，它自带了非常多的功能，比如数据迁移，可视化Explain， Dashboard, 及常见的查询， schema 设计等等。本文将以次工具来展示MySQL的许多功能
###	Navicat for MySQL
Navicat是一款商业软件，其功能是非常的强大，UI与易用性也是屈指可数的，是一个被广泛盗版的软件
###	PhpMyAdmin
一个基于Web的SQL客户端
###	DBForge Studio
它是一个商业软件，目前存在免费版跟收费版本，收费版本功能更多一些，界面跟以前的某些版本的Visual Studio类似，功能应该也比较强大 
###	HeidiSQL
使用Pascal写的一个SQL客户端工具，它有着非常强大的功能跟非常垃圾的用户使用体验。
###	Sequel Pro
Mac下的一个MySQL 客户端， 列出来只是给Mac用户多一个选择。
###	Eclipse与Idea 插件
这个是与自己的开发环境IDE集成的一种MySQL工具， 它往往提供的功能比较简单，通常包括语句查询、Schema建立与删除、自动补全等功能。
