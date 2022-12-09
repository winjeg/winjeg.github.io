---
title: Golang 下的数据库使用
date: 2018-12-13 15:14:11
toc: true
img: https://user-images.githubusercontent.com/7270177/59482717-9f4a4d80-8e9c-11e9-82b0-58254e0f4c4b.png
tags:
  - golang
  - database
categories:
  - lang
---

golang 下使用数据库是几乎每个golang程序员必须经历过的一个环节， 我们在这里专门挑了两个常见的数据库的使用方式来进行简单的科普一下。  
此文可以用作使用MySQL和Postgres的笔记性的文档， 不做深入分析， 全当给大家记备。

## Golang下使用 `MySQL`
由于golang 官方的SDK中已经定义好了数据库的访问接口， 还内定了连接池连接方式等基本的数据库操作元素， 但Golang并没有实现每种数据库的访问方式。  
因此如果要正常使用数据库，数据库相关的开发者需要找到定义了数据库具体访问方式与协议的数据库的驱动， 以访问数据库。 下面我们就从驱动讲起。  

### 驱动
```
go get github.com/go-sql-driver/mysql
```
### 基础使用

#### 创建数据库访问对象 `*sql.DB`
```go

import (
    "database/sql"
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
)
const (
    USERNAME = "root"
    PASSWORD = "*******"
    NETWORK  = "tcp"
    SERVER   = "localhost"
    PORT     = 3306
    DATABASE = "blog"
)
func getDb() *sql.DB {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
    DB,err := sql.Open("mysql",dsn)
    if err != nil{
        fmt.Printf("Open mysql failed,err:%v\n",err)
        return
    }
    DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
    DB.SetMaxOpenConns(100）//设置最大连接数
    DB.SetMaxIdleConns(16) //设置闲置连接数
    return db
}

var mysqlDb = getDb()

// 用此函数，则为单例， 只在初始化的时候创建一次， 不会多创建无用的连接池
func GetDb() {
    return mysqlDb
}

```

#### 查询数据库
```go
func FetchData(db *sql.DB) {
    // 查询多行
    rows, err := db.Query("SELECT * FROM user limit 3")
    if err != nil {
        // 如果err不等于nil, 一般row 为nil， 所以你直接close 会抛出空指针异常
        log.Error(err)
        return
    }
    // 注意这个一定要close
    defer rows.Close()

    // 取数据
    for rows.Next() {
        var userId int
        var username string
        // 如果有多行， 可以把结果放到相应的数据结构中
        sErr := rows.Scan(&userId, &username)
    }

    
    // 查询单行
    row := db.QueryRow("SELECT username FROM user WHERE user_id = ?", 1)
    var username string
    sErr := row.Scan(&username)
    // deal with error and username
}

```

#### 执行DML

```go
func DML() {
    // INSERT
    result, err := db.Exec("INSERT INTO user(username, user_id) VALUES(?, ?)", username, userId)
    stmt, err := db.Prepare("INSERT INTO user(username, user_id) VALUES(?, ?)")
    sr, err := stmt.Exec(username, userId)
    // result 和 sr 中都可以拿到自动生成的ID，以及影响的行数    
}
```

## Golang下使用 `Postgres SQL`

### 驱动
```
go get github.com/lib/pq
```

### 基础使用

#### 创建数据库访问对象 `*sql.DB`

```golang
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "testuser"
	password = "testpass"
	dbname   = "test_db"
)

func getDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=test_schema",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
    }
    return db
}
// 同样你可以再使用单例模式创建一个访问方式
// 但推荐使用 sync.Once() 进行单例模式的设计

```

#### Postgres SQL中查询数据
由于其他的与MySQL一样， 因此此处实例就简单一点写明：

```golang
func QueryPostgres(db *sql.DB) {
    // MySQL 使用 "?" Postgres 使用 $num,
    // 使用 ？ postgres 会报错
    db.Query("SELECT * FROM user WHERE user_id = $1", userId)
    db.QueryRow("SELECT * FROM user WHERE user_id = $1", userId)
}

```

#### Postgres SQL 中进行DML操作
```golang
func DML(db *sql.DB) {
    // 貌似必须得用stmt 直接使用Exec 会报错
    stmt, err := db.Prepare("UPDATE user SET username= $1 WHERE user_id = $2")
    stmt.exec("Lily", 123)
}
```

## 使用数据库的注意点

golang 使用数据库的注意点都基本相同因此不再把postgres跟mysql分开讲述

### 合理创建连接池

1. 合理设定连接池参数
2. 对于一个数据库仅仅创建一个连接池

### 关闭该关闭的结果集

主要还是提到的

```golang
// 如果不关闭， 它会慢慢的耗尽连接池的连接， 并且让你无连接可用
rows.Close()
```


