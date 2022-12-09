---
title: tidis 笔记
date: 2016-08-13 15:14:11
toc: true
# img: https://user-images.githubusercontent.com/7270177/59735413-0c3c5980-9288-11e9-8f32-d8e6836e65b6.png
tags: 
    - tidis
    - redis
    - database
categories:
    - storage
    - redis
---


## 策略
1. 懒删除
2. 用线程去扫所有的key以便删除该删的
3. 如果启用过期机制，则需要多一个key用来存过期时间
4. 不启用过期则比较快

##  缺少
1. 监控与统计信息
2. 充足的压测信息
3. 单测不够
4. key分析， 可以在扫描的时候做掉