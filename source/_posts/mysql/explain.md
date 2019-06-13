---
title: 使用Explain了解SQL语句的执行计划
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - sql
  - MySQL
  - database
categories:
  - storage
  - database
---

## 语法与含义
语义上讲EXPLAIN 、DESCRIBE 、 DESC 是一样的
但习惯上用法不太一样， DESC 经常用于描述一个表或者列的信息
EXPLAIN经常用于描述一个查询执行计划

```sql
explain select 1;
desc select 1;
describe select 1;
explain format=json select 1;
```


### explain_type
EXTENDED -> 显示额外信息(filtered info)
PARTITIONS ->  显示所用的分区 (目前没用到)
FORMAT -> 显示结果的格式 (JSON/TRADITIONAL, 分表表示用JSON或者表格形式)

Explain支持 SELECT/INSERT/UPDATE/DELETE/REPLACE 等语句

### Explain 结果集中需要特别注意的列
 | 类型            | 说明                             |
|---------------|--------------------------------|
| ALL           | 全表扫描                           |
| INDEX         | 使用覆盖索引，且全表扫描的时候                |
| RANGE         | 使用索引并且是在索引范围内查找的时候             |
| REF           | 在查询中直接用索引匹配单个值的时候              |
| EQ_REF        | 用索引查找，而且HIT到index的时候           |
| CONST, SYSTEM | MYSQL 可以把有些查询转换成常量比如用主键查找 一行记录 |
| NULL          | 如SELECT 1 等不需要表的查询             |

*对于上表而言， 越往下， 效率越高*


![type1](https://qhyxpicoss.kujiale.com/2018/12/17/LQLZ2NIKAQBZMZASAAAAADY8_931x497.png)

![type2](https://qhyxpicoss.kujiale.com/2018/12/17/LQL3LDQKAQBZOUTLAAAAAAY8_891x593.png)

### Key
Key 是指MySQL最终决定回用哪个索引到这个执行计划中， 一般有没有用到索引及用到的是哪个索引就看此列。
### Ref
REF 记录了 在Key列记录中的索引中所用的列，或常量
### Rows 与 Filtered
ROWS列记录了查找需要查找的元素，所需要扫描的行数
FILTERED的列显示的是针对表里面符合某个条件的记录的百分比所做的一个悲观的估算

### Extra
- `USING INDEX`
是不是用了覆盖索引
- `USING WHERE`
MYSQL会在检索行后根据WHERE条件进行过滤
- `USING TEMPORARY`
在对查询结果进行排序的时候是否用了临时表
- `USING FILE SORT`
MYSQL对查询结果进行排序，不用索引排序

## MySQL Visual Explain
Visual explain 是MySQL Workbench 带的一项功能，其他工具很少会有此项功能，它可以比较直观的告诉你整个查询的执行计划，及查取过程，是分析SQL语句不可或缺的利器。
在 Visual Explain中从蓝到红的效率是递减的，越偏向红效率就越低，越偏向蓝效率就越高。
下图是个例子
![visual explain](https://qhyxpicoss.kujiale.com/2018/12/17/LQL3LDQKAQBZOUTLAAAAAAQ8_690x857.png)
