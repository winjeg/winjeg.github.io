---
title: java 笔记
toc: true
date: 2016-03-13 15:14:11
tags: java
categories: lang
---

## 线程状态

![image](https://user-images.githubusercontent.com/7270177/59737806-90470f00-9291-11e9-904b-20ab88db3e52.png)


## 高cpu 使用的代码查找
1
ps H -eo user,pid,ppid,tid,time,%cpu,cmd --sort=%cpu 

tid 是线程ID 转成十六进制去查

jstack 9002 > stack.log

2. arthas

