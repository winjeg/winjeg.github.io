---
title: Golang的工程配置
date: 2018-12-13 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59482717-9f4a4d80-8e9c-11e9-82b0-58254e0f4c4b.png
tags:
  - golang
categories:
  - lang
---

# go module简介

Go 1.11 中引入了module 管理的概念，这使得我们终于有一个官方正式支持的模块管理工具了， 但由于官方工具刚出功能还不是十分完善，易用， 很多人还摸索不出来如何使用是最科学的。

## 为什么要使用 go module?

 更好的版本管理， 使得项目依赖的module版本都是确定以及稳定的（这是依赖管理的最核心的要求）
依赖管理更简单方便
dep本身不利于项目中私有仓库的包共享， dep 在拉取依赖的时候， 会把依赖放到 项目目录的vendor下面， 这样也能达到上面说的这个效果，
但是dep本身对私有仓库的支持不好，不能很好的配置ssh key拉取私有仓库代码（无论如何配置它就是不用，不是不能用就是不好用）
另外，dep对依赖的拉取是相对比较慢， 处理比较复杂的。
dep已经进入不维护状态， 而且go module 得到了官方的支持与更新的确认


## 私有仓库
### 通过git命令
```
git config --global url.git@github.com:.insteadof=https://github.com
```

### 修改git配置文件
看下 vim ~/.gitconfig 是否有生效，否则，就手动改一下
```
[url "git@github.com:"]
        insteadOf = https://github.com/
```
重启terminal生效

## go module的要求
`golang version < 1.12`
```
GO111MODULE=on
```
由于官方没有默认打开 go module 的feature 因此在设置ci的时候需要手动设置一下两个环境变量, 这是唯一的代价

`golang version >=1.12`则不需要进行任何设置


## 依赖管理

## 初始化项目
```
go mod init mod_name
```

### 添加依赖
```
go get github.com/demo/xxx
```
## 删除依赖
```
go mod edit -droprequire github.com/demo/test
```

## 更新依赖版本
```
go mod edit xxx
```
更新与删除依赖比较难用， 建议直接修改go.mod 文件与go.sum 文件， 更简单直接

## 整理依赖 
```
go mod tidy
```
这句命令会自动去除没用的依赖， 添加需要增加的依赖

## 依赖查询
```
go mod graph
```