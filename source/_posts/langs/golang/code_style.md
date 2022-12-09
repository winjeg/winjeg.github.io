---
title: Golang的代码风格
date: 2020-05-30 15:14:11
toc: true
img: https://user-images.githubusercontent.com/7270177/59482717-9f4a4d80-8e9c-11e9-82b0-58254e0f4c4b.png
tags:
  - golang
categories:
  - lang
---

## 一些废话（Some useless words）
想必能看到我博客的人，已经对golang有一定的了解了


### gofmt 一统天下  

gofmt 是目前golang里面用的最多的用来格式化代码风格的一个命令行工具， 很多知名项目都用它来保证自己的最基本的代码风格与官方和社区推荐的风格一致。  
使用如下命令就可以轻松格式化一个目录下的所有文件：  
```sh
gofmt -w .
```
1. gofmt 可以解决的问题

2. gofmt 不能解决的问题


### 天下之外
其实标准gofmt有很多代码风格没有规定的地方比如以下几种场景：  
- 常量命名风格
- 变量命名风格
- 函数命名风格
- 对象字段使用风格
- 值传递，还是指针传递


#### 据库对象：
1. 可空字段，使用指针，用来表示NULL
2. 对于不可空字段不使用指针类型， 用来表示这个字段一定有值


### 注释风格

golang 与某Java不同， 不喜欢多行注释，纵观Golang SDK以及 一些非常文明的项目，大部分注释均为单行注释， 虽然golang 支持以下两种注释类型。  
```golang
//  单行注释
var code = 0

/*
多行注释
 */
const someVeryNastyThing = -1
```

#####  特殊指令
在golang注释里面可以写一些特殊指令， 这个时候编译器就会处理这些指令，而不是仅仅当做注释， 这在很多场景下非常有用。