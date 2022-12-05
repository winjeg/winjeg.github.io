---
title: Mysql 其他杂记
date: 2019-05-13 15:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59422808-880b5180-8e03-11e9-9dfe-ff8a9a024be7.png
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---

## 事务查询


```sql
SELECT @@tx_isolation
```


```sql
SELECT 
  r.`trx_id` waiting_trx_id,
  r.`trx_mysql_thread_id` waiting_thread,
  r.`trx_query` waiting_query,
  b.`trx_id` bolcking_trx_id,
  b.`trx_mysql_thread_id` blocking_thread,
  b.`trx_query` block_query 
FROM
  information_schema.`INNODB_LOCK_WAITS` w 
  INNER JOIN information_schema.`INNODB_TRX` b 
    ON b.`trx_id` = w.`blocking_trx_id` 
  INNER JOIN information_schema.`INNODB_TRX` r 
    ON r.`trx_id` = w.`requesting_trx_id` ;
 
```
## 查询外键

```sql
select concat('alter table ',table_name,' drop foreign key ',constraint_name,';') 
from information_schema.key_column_usage
where constraint_schema = 'dbname' and referenced_table_name = 'tbName';


SELECT concat('alter table ',table_schema,'.',table_name,' DROP FOREIGN KEY ',constraint_name,';')
FROM information_schema.table_constraints
WHERE constraint_type='FOREIGN KEY'
AND table_schema='dbname';

SELECT TABLE_NAME FROM information_schema.`TABLES`
WHERE TABLE_TYPE="BASE TABLE"
AND TABLE_SCHEMA="dbname"
```

原因是由于RDS的备份文件中包含了RDS的主从复制关系，需要把这些主从复制关系清理掉，清理方法：

```sql
truncate table  slave_relay_log_info;

truncate table  mysql.slave_master_info;	
```


### Mysql DUMP 表
```
mysqldump -u root -h 127.0.0.1  dbName msgattach  > msgattach.sql
```


###  MySQL 客户端
从表现上来看分两种：
1. 客户端执行完query之后， 直接返回， 并开始用rows.next 去取数据， 其中很快
2. 如果在客户端取数据的时候打上断点， 服务端发送完数据之后， 客户端只能取到部分数据， 不能取到全部数据， 怀疑客户端有一个缓冲区，不断接收服务端的数据并刷新
3. 如果客户端收到数据后，不打断点， 则可以获取全部数据（前提是建立在客户端执行速度比较快的情况下）

底层代码里面来看， mysql服务端发送数据给客户端的时候， 客户端会把数据存在一个 默认4096大小的buffer 里面， 从这个buffer里面再读到rows里面
如果buffer里的数据没有被及时消费掉， 那么连接上面传送过来的数据会丢失掉。

MySQL 本身的问题对于数据量传输不完的时候有个write_timeout,  出现错误之后， 会返回EOF 相关的Error
```
[mysql] 2019/04/25 14:31:27 packets.go:72: unexpected EOF
[mysql] 2019/04/25 14:31:27 packets.go:393: busy buffer

```


### 实用脚本

#### Mysql 查看表的外键 
```sql
select 
TABLE_NAME,COLUMN_NAME,CONSTRAINT_NAME, REFERENCED_TABLE_NAME,REFERENCED_COLUMN_NAME 
from INFORMATION_SCHEMA.KEY_COLUMN_USAGE 
 WHERE CONSTRAINT_NAME != 'PRIMARY' AND REFERENCED_TABLE_NAME IS NOT NULL 
```

#### 测试稳定情况
```bash
for i in `seq 1 100`; do 
    mysql -u username -h host -P 3306 -ppassword -e "use dbName; select askid from ask limit 1"; echo "==============$i"; sleep 1;
done
```