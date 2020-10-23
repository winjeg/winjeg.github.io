---
title: TDDL 常见配置详解
date: 2016-09-13 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59735413-0c3c5980-9288-11e9-8f32-d8e6836e65b6.png
tags: 
    - MySQL
    - middleware 
categories:
    - storage
    - middleware
---


## TDDL 简介
TDDL 是淘宝开源的一个基于Client做的， 用于如下特性的一个中间件：
1. 读写分离
2. 权重调配
3. 分库分表
4. 流控/限速
5. 数据库信息配置化
6. 查询重写与优化
TDDL 最新的版本是个[泄露的版本](https://github.com/loye168/tddl5), 有兴趣的同学可以下载源代码研究一下。

TDDL的配置均以Key-Value 形式存储， 拥有 内存-磁盘-远程 三级降级措施, 官方默认存储配置的地方为`Diamond`服务器。
由于TDDL的配置规则相对复杂， 也许很多人并不能入门， 此文档就是给大家一个配置入门的方案。


## 变量列表说明
1. `${dbName}` 数据库名称
2. `${envName}` 环境名称
3. `${userName}` 访问数据库使用的用户名
4. `${ip}` mysql 服务器的IP地址
5. `${appName}` appName，数据源名称
6. `${groupName}` 分组名称
7. `${atomName}` 物理数据库别名

## 数据源的拓扑结构
### com.taobao.tddl.v1_ds-${appName}_topology
```xml
<?xml version="1.0" encoding="UTF-8"?>
<matrix xmlns="https://github.com/tddl/tddl/schema/matrix"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:schemaLocation="https://github.com/tddl/tddl/schema/matrix https://raw.github.com/tddl/tddl/master/tddl-common/src/main/resources/META-INF/matrix.xsd">
    <appName>${appName}</appName>
    <group name="${groupName}" type="mysql_jdbc">
        <atom name="${atomName}" />
    </group>
</matrix>
```
这份key-value 对， 配置的是数据源的整体的拓扑结构信息， 这里会想尽的写清楚有哪些 数据分组， 哪些物理数据库会被用到.
其中  `appName` 元素只允许存在一个， 但 `group` 元素允许存在多个， 也允许非mysql段的存在， `atom`元素只能存在于`group` 元素之中，但可以存在多个， 每个`group` 下面的 `atom` 都是在数据上是等价的（数据集合相同， 可能读写属性不同）


## 数据源的规则集合
### com.taobao.tddl.rule.le.${appName}
```xml
<beans xmlns="http://www.springframework.org/schema/beans">
    <bean id="designsnapshot_bean" class="com.taobao.tddl.rule.TableRule">
       <property name="dbNamePattern" value="designsnapshot_{0000}"/>
       <property name="dbRuleArray">
           <value>(#designid,1,64#.longValue() % 256).intdiv(64)</value>
       </property>
       <property name="tbNamePattern" value="designsnapshot_{0000}" />
       <property name="tbRuleArray">
           <value>#designid,1,64#.longValue() % 64</value>
       </property>
       <property name="allowFullTableScan" value="true"/>
   </bean>
   <bean id="vtabroot" class="com.taobao.tddl.rule.VirtualTableRoot">
       <property name="tableRules">
           <map>
               <entry key="design_snapshot" value-ref="designsnapshot_bean" />
           </map>
       </property>
       <property name="dbIndexMap">
           <map>
               <entry key="designsnapshot_sequence" value="designsnapshot_0000" />
           </map>
       </property>
       <property name="defaultDbIndex" value="tms-db-dev" />
   </bean>
</beans>
```
 
这是一个表明整个数据源拓扑结构的规则文件， 也是分库分表规则配置的地方. `vtabroot` 是总的规则入口的`bean`。下面有几个概念需要详细解释一下。
1.  `defaultDbIndex` 是默认的数据分组的设定， 当一个数据源（逻辑）有多个数据库（不同的数据库）的时候， 一个查询分发需要知道这个查询所涉及的表在哪个库上， 如果没有特殊指定， 那就会分发到这个东西指定的 `group`上。
2. `dbIndexMap` 用于指定哪些表在哪 `group`， 具体来说就是表库的映射关系， 优先级比 `defaultDbIndex` 高， 同样是在用于配置多个不同数据的数据库的时候使用的。
3. `tableRules` 用于定义虚拟表与虚拟表对应的分库分表规则的实际物理数据库关系。 key 是虚拟表名， 可以直接 SELECT FROM , value-ref 是 分库分表规则对应的bean
4. `分库分表的bean` 用于设置分库分表的规则
`dbRuleArra` 是分库规则。如例所示， 以 designid列 进行分库， 每个库64张表， 共分为 256/64个库，比如designid = 500,  244 / 64  = 3 分到第四个库上。 
`tbRuleArray` 是分表规则。进入到每个库上的ID再进行取余64决定最后落到哪个表里。 这两个字段写法并非固定， 都是支持TDDL内置的一些表达式语法的， 具体得去看代码，此处只列出了最常用的取模的表达式。


## 数据源的读写分离权重优先级等配置

### com.taobao.tddl.jdbc.group_V2.4.1_${groupName}
```
${atomName}:rwp0i0,${atomName}:rp1i1
```
逗号是各个数据库原的分隔符，冒号前面的是`atom`的名字。`rw` 是指读写, `r` 是指只读, `p` 是优先级的意思， p越大则越优先分配到。`i` index， 是指各个数据源的一个index， 此外还有个权重配置的`w`可以写在 `p` 后面 

## 数据库本身的配置
### com.taobao.tddl.atom.global.${atomName}
```
dbName=${dbName}
dbType=mysql
dbStatus=RW
ip=${ip}
port=3306
```
数据库自身的配置， 与TDDL 无关， 需要让TDDL知道数据库的地址， 端口， 数据库名称等基本信息

## 连接池的配置
# com.taobao.tddl.atom.app.${appName}.${atomName}
```
userName=${userName}
maxPoolSize=100
minPoolSize=5
idleTimeout=1
testOnBorrow=1
abandonedTimeout=60
connectionProperties=charset=utf-8
```
由于TDDL底层使用`Driud`连接池， 因此， 这里主要用于配置 `druid` 连接池

## 数据库账号与密码指定
### com.taobao.tddl.atom.passwd.${dbName}.mysql.${userName}
```
encPasswd=${password}
```
此处用于配置数据库用的密码。 由于开源出来的版本好像是没有对应的加密逻辑实现的， 因此， 这里的密码并不是加密的密码，而是明文的密码。
不过用户可以自行实现加密方法， 此处的接口TDDL是有留出来的。
