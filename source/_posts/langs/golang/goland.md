---
title: Golang 简介
date: 2017-09-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - golang
categories:
  - lang
---





db.Query(..) returns a rows struct which must be closed or completely iterated (with for rows.Next()). Otherwise there is unread data on the connection stream and the connection blocks.

Simple solution: use db.Exec(..) instead. Does it work then?

gorm 的问题

```
db.Raw.Row()
```
  如果这个row不被消费则会造成这个问题
如果不想消费
```
db.Exec()
```

## golang 初始化顺序
![image](https://user-images.githubusercontent.com/7270177/59737472-62ad9600-9290-11e9-92e7-54556e4618de.png)
