---
title: JWT 基本原理笔记
date: 2021-11-12 10:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/96998573-3f6eb700-1566-11eb-8c16-ca031f895b69.png
tags:
  - web
categories:
  - other
---



## jwt 简介：

jwt 是json web token 的缩写， 主要用来做用户授权或者session登录的事情。
JWT 是一种比较流程的协议， 它是由三段式组成的

1. header
2. body
3. sign

---

1. header 包含了此JWT用的加密算法， body 则为具体的payload.
2. 签名则是由body header经过base64 加密后再进行生成的类似hash/md5 值的东西, 主要是为了防止jwt被串改， 保证请求的合法性
## 安全相关
由于JWT的信息并非经过加密的信息，因此JWT本身不适合存一些比较敏感的信息，比如密码或者secret之类的东西。
JWT 存储的信息是确定用户身份而用的， 一般不应该被浏览器中的其他JS读取到， 因此我们可以设置JWT cookie 的 httpOnly 属性为true这样一些基于JS的攻击对TOKEN的窃取就不会生效。

## jwt浏览器与服务端的交互流程
jwt信息一般会被放在浏览器的cookie中， 设置一定的有效日期， 每次客户端发起请求的时候都会吧这个cookie信息发给服务端， 服务端本身可以获取到这个cookie信息，并经过签名校验和解密，可以得到用户信息。