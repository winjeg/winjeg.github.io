---
title: Golang 简介
date: 2017-09-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - golang
categories:
  - lang
---
## IDE 选择
### vscode

### goland （收费）
需要装的插件列表
1. material theme UI
2. m

jetbrains 出品， 省去了很多的配置麻烦，而且用IDEA的用户可以无缝切换到这个的使用。


## 常见问题

### SQL查询的Rows不关闭的时候回造成阻塞的问题

对此的英文解释大概意思是， 如果不关闭rows， 缓冲区就有未读完的数据， 然后就造成了连接阻塞，影响接下来的查询， 最终会导致连接池的连接全部用满。

```
db.Query(..) returns a rows struct which must be closed or completely iterated (with for rows.Next()). Otherwise there is unread data on the connection stream and the connection blocks.
```
如果需要执行但不需要结果可以用 `db.Exec(..) `

```
Simple solution: use db.Exec(..) instead. Does it work then?
```

gorm 的问题

```go
db.Raw.Row()
```
  如果这个row不被消费则会造成这个问题
如果不想消费

```go
db.Exec()
```

## golang 初始化顺序
![image](https://user-images.githubusercontent.com/7270177/59737472-62ad9600-9290-11e9-92e7-54556e4618de.png)
