---
title: TDDL 简介及入门使用
date: 2016-08-13 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59735413-0c3c5980-9288-11e9-8f32-d8e6836e65b6.png
tags: 
    - MySQL
    - middleware 
categories:
    - storage
    - middleware
---

## TDDL简介
TDDL 是淘宝开源的一个用于访问数据库的中间件， 它集成了分库分表， 读写分离，权重调配，动态数据源配置等功能。封装 `jdbc ` 的 `DataSource`给用户提供统一的基于客户端的使用。它只是一组Jar包， 并不是单独的服务，目前已经被一些公司默默使用， 因为官方至今为止没有维护这一块的开源社区， 大多数使用也停留在各个公司自己研究的阶段。

TDDL 最新的版本是个[泄露的版本](https://github.com/loye168/tddl5), 有兴趣的同学可以下载源代码研究一下。

## 使用方法
### 引入依赖
```xml
<dependency>
    <groupId>com.taobao.tddl</groupId>
    <artifactId>tddl-matrix</artifactId>
    <version>${project.version}</version>
</dependency>
<dependency>
    <groupId>com.taobao.tddl</groupId>
    <artifactId>tddl-config-diamond</artifactId>
    <version>${project.version}</version>
</dependency>
<dependency>
    <groupId>com.taobao.tddl</groupId>
    <artifactId>tddl-parser</artifactId>
    <version>${project.version}</version>
</dependency>
<dependency>
    <groupId>com.taobao.tddl</groupId>
    <artifactId>tddl-sequence</artifactId>
    <version>${project.version}</version>
</dependency>
```
### 注册bean
```java
@Bean
public DataSource getDataSource() throws TddlException {
    TDataSource dataSource = new TDataSource();
    dataSource.setAppName("appName");
    dataSource.init();
    return dataSource;
}
```
完成以上两步之后，它的使用是跟普通的 `Datasource` 是一样的, 可以很方便的与各种ORM框架集成。



## 三层架构（可独立使用）：
### Matrix（TDataSource）
`matrix` 是整个 `datasource`的入口， 用户使用的入口。它是一种标准接口的多数据源多功能的实现。
这里控制了一个数据源所有的对应的数据库的信息， 下面可以持有多个 `group`, 而 `group`下面有可以有多个 `atom`。
它也是唯一支持分库分表实现的一层， 用户一般使用的是这一层， 如果不需要其他规则和功能也可以使用这一层。

### Group（TGroupDataSource）
`group` 是逻辑数据分组的概念， 一个 `group` 下的数据一般都是相同的， 不同的是对数据源的配置， 不同实际数据源， 对应不同的读写属性， 优先级， 权重等

### Atom（TAtomDataSource）
`atom`  是对实际物理数据库的一种抽象。 它持有的是实际的物理数据库的信息，以及这个物理数据库所相关的 `druid` 连接池的配置信息。
它隶属于 `group`。


## 不支持的点
1. 不支持跨库事务
2. 不支持很奇怪的SQL
3. 不推荐使用跨库JOIN的功能

## 其他
1. `tddl-manager` 模块是一个管理TDDL数据源的界面UI
2. `tddl-server` 是mysql协议的一个实现， 底层使用TDDL数据源
