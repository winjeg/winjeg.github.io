---
title: MySQL Schema 设计
date: 2016-03-13 15:14:11

toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---

## Schema 简介
说到使用MySQL我们接触最早的就是Schema设计，俗称建表，这个小节主要介绍MySQL Schema的设计方法与一些基本的使用原则。

Schema 设计在用户应用设计前期是非常重要的，一般情况下它会影响到业务以后的健康程度，以及其他业务代码的设计，Schema一旦设计成型并投入使用，当数据量达到一定程度的时候将会对索引的要求会越来越高，变更Schema也将会花费更多的代价，应用设计之前，认真正确的设计Schema是非常有必要的。

下面将主要分Schema的数据类型，索引与外键，使用原则，默认值等来介绍Schema的设计，在此之前，下面是一个最常见的一个schema设计语句。

```sql
CREATE TABLE `userinfo` (
`id`  int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键' ,
`name`  varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名字' ,
`gender`  bit NULL DEFAULT NULL COMMENT '性别 0 男 1 女 NULL 未知' ,
`height`  float UNSIGNED NULL DEFAULT NULL COMMENT '身高， 单位m' ,
`account_id`  int UNSIGNED NOT NULL COMMENT '账户ID' ,
`created`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间' ,
`last_modfied`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后修改时间' ,
PRIMARY KEY (`id`),
INDEX `idx_name` (`name`) USING BTREE COMMENT '用于按姓名查找',
UNIQUE INDEX `idx_account_id` (`account_id`) USING BTREE COMMENT '用于对应某个账户的唯一ID',
INDEX `idx_heigt_gender` (`height`, `gender`) USING BTREE COMMENT '用户筛选用户特征'
)
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;
```

## 数据类型

###	整形数据： TINYINT, SMALLINT, MEDIUMINT, INT BIGINT
整形的数据类型一旦类型确定长度也就确定了，所以之前文章里面给定的那个长度是不起实际作用的，可以不指定长度只指定类型。
下表是各种整形所对应的长度
| 数据类型      | 位数 | 范围                                                       |
|-----------|----|----------------------------------------------------------|
| TINYINT   | 8  | -2^7 到 2^7-1      (-128-127)                             |
| SMALLINT  | 16 | -2^15到2^15-1     (-32768-32767)                          |
| MEDIUMINT | 24 | -2^23到2^23-1   (-8388608-8388607)                        |
| INT       | 32 | -2^31到2^31-1     (-2147483648-2147483647)                |
| BIGINT    | 64 | -2^63到2^63-1  (-9223372036854775808-9223372036854775807) |
应用场景：INT 通常用作一个实体表的主键部分，比如用户表的用户ID
TINYINT 通常用作枚举类型的数据，需要很小长度比如 status/type/deleted
通常会根据一个数据将来可能的最大量来选择类型数据， 一般使用前会做一个预估来选择数据类型。

###	实数型数据： DECIMAL， FLOAT， DOUBLE
DECIMAL 可以用于存储比BIGINT 更大的整数，计算方法MYSQL自行实现
FLOAT与DOUBLE 都可以指定精度
应用场景：实数类型通常应用与需求精度非常高的字段
可能的改进方法：在存储面积，身高等需要精度低的领域，在最大值可预估的时候，可以用整形来存取 

###	字符串类型 VARCHAR CHAR TEXT BLOB
VARCHAR可变长字符串， 需要额外字节来存储当前的长度, 用于存储，改动不太频繁，且长度变化范围比较大的字符串
CHAR 固定长字符串 用于存储相对固定的东西，如密码的MD5，以及经常要改变的列
TEXT与BLOB 分别是指用字符存储与二进制的存取
可能的改进：
如果一个列的字符串的数量基本是确定的， 比如性别： Male Female Unknown 只有固定的几种String 或者是可以遇见的几种，尽量可以使用tinyint 来代替
常用函数
```sql
CHAR_LENGT
STR_CMP
STR_TO_DATE
```
###	日期类型 TIMESTAMP, DATETIME
MYSQL支持毫秒级别的时间， MariaDb 支持微秒级的时间值
DATETIME (1001-9999)  时区无关  使8Byte的存储空间
TIMESTAMP（1970-2038）时区相关 使用4Byte的存储空间
常用函数
```sql
# 把UNIX时间戳转成TIMESTAMP
UNIX_TIMESTAMP() /FROM_UNIXTIME() 
DATE()
DATE_ADD
DATE_SUB
DATE_FORMAT
```
###	枚举类型与集合类型
此类型仅作列出， 不作实际使用， 一般这些类型可以用TINYINT 等值来代替，不建议用这些类型


## 默认值
###	NULL
NULL 是最常见的默认值，一切默认值都可以设置为NULL， 但不是所有默认值都要设置为NULL， 要根据情况具体设定
###	Empty String
空字符串，是很多char, varchar, text 的默认值的选项，如果你的应用程序没有处理NULL值的时候可以把默认值设置为这个避免空指针异常， 但一般情况下不希望MySQL给你做这类的检测。

###	CURRENT_TIMESTAMP
当前时间， 一般用于TIMESTAMP的默认值，它在使用跟理论上均优于创建TRIGGER来设置默认值

## 索引与外键
索引与外键是建立Schema 永远不能避开的话题，你可以避开外键，但你可能会想到它，但你不能避开索引。这里不会对索引进行详细的表述，具体的描述在后面小节中会详细道来。在这里你只需要记住一点，建立表的时候，一定要根据查询建立索引，最好不要建立外键，除非非常必要。

## Schema设计误区
###	太多的列
太多的列需要更大的维护成本， 对与SELECT * 的效率来说也更低

###	太多的JOIN
JOIN过多，会导致查询复杂，结果更难判断，查询效率变低

###	泛滥的外键与索引
索引与外键并不越多越好， 每个外键与索引都需要额外的维护成本

###	随意性建表
在很小的数据量的情况下， 通常存在一个汇总表，或者设置表里面比较合适]

## 范式与反范式
| 设计模式 | 优点                           | 缺点                                |
|------|------------------------------|-----------------------------------|
| 范式   | 数据量小，单实体表查询效率高               | 需要做很多的JOIN                        |
| 反范式  | 反范式把冗余信息都放在一张表里面，可以很好的避免JOIN | 数据量的增大，导致查询的速度，稍受影响； 更新须保证冗余信息的更新 |

两者要结合实际情况，可以混合使用，但一般不推荐过多的JOIN，SQL语句JOIN不能超过两张表。
## Schema 设计原则
###	简单就好，更小的通常更好
越简单的Schema所需要的存储空间就越小，对于MySQL而言需要的存储空间， IO， 运算等资源就越少，MySQL支持的QPS就越高
###	尽量避免大量NULL值的列
如果一个列有80%甚至90%以上的行都是NULL，那么这个属性值，很可能可以采取其他方法来实现，比如新建一张表，用ID去关联
###	不使用外键
外键在MySQL维护与使用中始终是一个坑，大家尽量在建表的时候避免使用外键。
###	合理建立索引
索引要根据查询而建， 并不是越多越好，索引多有可能会引起MySQL不能正确使用索引，MySQL维护索引的空间变大等等
###	写好表注释
与写代码一样一个良好的习惯就是所有MySQL的字段，及索引，表等都要有详尽的注释，注释不会占用多余空间，最好的注释是大家都能看懂的注释。
