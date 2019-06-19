---
title: Mysql 主从同步相关的知识
date: 2016-04-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---


同步时mysql -> backup 

```bash
mkdir -p /backup/mysql
tar -izxvf /backup/hins1851643_data_20170301022021.tar.gz -C /backup/mysql/
cp /home/wenwen/backup-my.cnf /backup/mysql/backup-my.cnf
innobackupex --user=root --password --defaults-file=/etc/mysql/my.cnf --apply-log /backup/mysql
chown -R mysql.mysql /var/lib/mysql
chown -R mysql.mysql /backup/mysql
```
replicate.conf

```ini
[mysqld]
server-id       = 2

## same on the slave
log-bin         = mysql-bin
innodb_buffer_pool_size=512M
innodb_flush_log_at_trx_commit  = 1
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
log_bin_trust_function_creators=1

sync_binlog     = 1
#binlog-format='ROW'

relay-log=mysqld-relay-bin

read_only=1

master-info-repository=file
relay-log-info_repository=file
binlog-format=ROW
innodb_checksum_algorithm=innodb
innodb_data_file_path=ibdata1:200M:autoextend
innodb_log_files_in_group=2
innodb_log_file_size=524288000


#relay_log_info_repository = TABLE
#master_info_repository    = TABLE
relay_log_recovery        = on

### replicate settings
replicate-ignore-db=performance_schema
replicate-ignore-db=mysql
replicate-do-db=fenshua123
#replicate-rewrite-db="maindata->diydata"
```


```sql
CHANGE MASTER TO MASTER_HOST='10.1.8.160', MASTER_USER='rdsinner',  MASTER_PORT = 3306, MASTER_LOG_POS=120, MASTER_PASSWORD='the_password', MASTER_LOG_FILE='mysql-bin.002458';
```

## 主从同步

1. SLAVE  切换命令： CHANGE MASTER TO MASTER_HOST='10.1.6.111', MASTER_USER='fenshuaprod', MASTER_LOG_POS=120, MASTER_PASSWORD='TUzu4515', MASTER_LOG_FILE='mysql-bin.000001', relay-log=mysqld-relay-bin;
2. Master 重启在先， Slave重启在后，即可保持复制关系， 如果Slave 重启在先， 则需要在Slave 上手动 start slave；才可以维持复制关系
3. MySql 支持引入其他配置文件， 用 !include /filepath 即可，但注意引入的文件要标明在那个section 下面，否则很容易就会抛异常 
4.检查relay_log_info_repository是否修改成功。
	show variables where variable_name in  ('relay_log_info_repository','master_info_repository');
5. 设置表只读 lock table t_depart_info read;  
6. 设置表名忽略大小写 lower_case_table_names=1
7. 设置库只读read-only
8. sudo innobackupex --user=root --password --defaults-file=/data/backup-my.cnf --tables-file=/data/site.cnf --ibbackup=xtrabackup_56 --apply-log /data/mysql




1．在ECS服务器上安装MySQL，详细步骤可以参考如下：
http://www.centoscn.com/mysql/2014/0924/3833.html
一些关键注意点：
a.数据库的版本至少为5.6.16及以上
b.需要在my.cnf中配置的一些关键参数：
```ini
server-id ###Slave配置需要
master-info-repository=file### Slave配置需要
relay-log-info_repository=file### Slave配置需要
binlog-format=ROW### Slave配置需要
gtid-mode=on###开启GTID需要
enforce-gtid-consistency=true###开启GTID需要
innodb_data_file_path=ibdata1:200M:autoextend###使用RDS的物理备份中的backup-my.cnf参数
innodb_log_files_in_group=2###使用RDS的物理备份中的backup-my.cnf参数
innodb_log_file_size=524288000###使用RDS的物理备份中的backup-my.cnf参数
```
2.MySQL安装好后，可以使用RDS提供的物理备份文件恢复到本地MySQL中，可以参考：
http://help.aliyun.com/knowledge_detail/5973700.html?spm=5176.7114037.1996646101.1.7qe3ot&pos=1
注意：
需要将备份解压后的文件backup-my.cnf中的三个参数加到启动文件中去
```ini
innodb_checksum_algorithm=innodb
innodb_data_file_path=ibdata1:200M:autoextend
innodb_log_files_in_group=2
```
3.数据库启动后，开始设置本地数据库与RDS的同步关系
a．reset slave;####用于重置本地MySQL的复制关系，这一步操作有可能报错：
```
mysql> reset slave;
ERROR 1794 (HY000): Slave is not configured or failed to initialize properly. You must at least set –server-id to enable either a master or a slave. Additional error messages can be found in the MySQL error log.
```
原因是由于RDS的备份文件中包含了RDS的主从复制关系，需要把这些主从复制关系清理掉，清理方法：
```sql
truncate table  slave_relay_log_info;
truncate table  mysql.slave_master_info;
truncate table  mysql.slave_worker_info;
```
然后重启MySQL；
b.SET @@GLOBAL.GTID_PURGED
=’818795a2-8aa8-11e5-95b1:1-289,8da7b8ab-8aa8-11e5-95b1:1-75′;
打开备份解压文件可以看到文件xtrabackup_slave_info，其中第一行就是我们需要在本地MySQL执行的命令，他表示在备份结束时刻RDS当前GTID值’
c.
```sql
change master to
master_host=’gtid1.mysql.rds.aliyuncs.com’,
master_user=’qianyi’,master_port=3306,master_password=’qianyi’,
master_auto_position=1;
```
设置本地MySQL与RDS的复制关系，账户qianyi是在RDS控制系统中添加（注意：
同步账户不要以repl开头）；
4．测试同步关系是否正常，可以在本地MySQL执行show slave status\G查看同步状态，同时可以在RDS中插入测试一些数据，或者重启实例，观察同步情况：
```
mysql> show slave status\G;
Slave_IO_State: Queueing master event to the relay log
Master_Host: gtid1.mysql.rds.aliyuncs.com
Master_User: qianyi
Master_Port: 3306
Connect_Retry: 60
Master_Log_File: mysql-bin.000007
Read_Master_Log_Pos: 625757
Relay_Log_File: slave-relay.000002
Relay_Log_Pos: 2793
Relay_Master_Log_File: mysql-bin.000007
                Slave_IO_Running: Yes
                Slave_SQL_Running: Yes
Exec_Master_Log_Pos: 612921
Relay_Log_Space: 15829
       Seconds_Behind_Master: 57133
Master_SSL_Verify_Server_Cert: No
Master_Server_Id: 2319282016
Master_UUID: 818795a2-8aa8-11e5-95b1-6c92bf20cfcf
Master_Info_File: /data/work/mysql/data3001/mysql/master.info
SQL_Delay: 0
SQL_Remaining_Delay: NULL
Slave_SQL_Running_State: Reading event from the relay log
Master_Retry_Count: 86400
818795a2-8aa8-11e5-95b1-6c92bf20cfcf:17754-17811
Executed_Gtid_Set: 818795a2-8aa8-11e5-95b1-6c92bf20cfcf:1-17761
Auto_Position: 1
```
5.做好监控，由于采用MySQL的原生复制，所以可能会导致本地MySQL与RDS的复制出现中断，可以定时去探测  Slave_IO_Running和 Slave_SQL_Running两个状态值是否为yes，同时也需要关注本地MySQL与RDS的延迟： Seconds_Behind_Master。



删除slave_worker_info内容




```
mysql> set global sync_binlog=20 ;
Query OK, 0 rows affected (0.00 sec)
mysql> set global innodb_flush_log_at_trx_commit=2;Query OK, 0 rows affected (0.00 sec)
innodb_flush_log_at_trx_commit
```
如果innodb_flush_log_at_trx_commit设置为0，log buffer将每秒一次地写入log file中，并且log file的flush(刷到磁盘)操作同时进行.该模式下，在事务提交的时候，不会主动触发写入磁盘的操作。
如果innodb_flush_log_at_trx_commit设置为1，每次事务提交时MySQL都会把log buffer的数据写入log file，并且flush(刷到磁盘)中去.
如果innodb_flush_log_at_trx_commit设置为2，每次事务提交时MySQL都会把log buffer的数据写入log file.但是flush(刷到磁盘)操作并不会同时进行。该模式下,MySQL会每秒执行一次 flush(刷到磁盘)操作。
注意：
由于进程调度策略问题,这个“每秒执行一次 flush(刷到磁盘)操作”并不是保证100%的“每秒”。
sync_binlog
sync_binlog 的默认值是0，像操作系统刷其他文件的机制一样，MySQL不会同步到磁盘中去而是依赖操作系统来刷新binary log。
当sync_binlog =N (N>0) ，MySQL 在每写 N次 二进制日志binary log时，会使用fdatasync()函数将它的写二进制日志binary log同步到磁盘中去。
注意:
如果启用了autocommit，那么每一个语句statement就会有一次写操作；否则每个事务对应一个写操作。
而且mysql服务默认是autocommit打开的
修改参数后，slave2,3也一样可以跟上slave1的速度了
