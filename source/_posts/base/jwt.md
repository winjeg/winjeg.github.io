---
title: JWT 基本原理笔记
date: 2020-10-23 19:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/96998573-3f6eb700-1566-11eb-8c16-ca031f895b69.png
tags:
  - base
  - service
categories:
  - base
---

jwt 简介：

jwt 是json web token 的缩写， 主要用来做用户授权或者session登录的事情。


JWT 是一种比较流程的协议， 它是由三段式组成的

1. header 
2. body 
3. sign

> 说明：
1. header 包含了此JWT用的加密算法， body 则为具体的payload.
2. 签名则是由body header经过base64 加密后再进行生成的类似hash/md5 值的东西, 主要是为了防止jwt被串改， 保证请求的合法性

