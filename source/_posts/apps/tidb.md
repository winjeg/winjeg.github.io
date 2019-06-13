# tidb 相关

## 简介
TIDB是一种分布式高可用实现MySQL协议的可动态扩容的数据库
动态扩容是指计算能力与存储空间两个维度的动态扩容
tidb 由 tidb server， PDserver 与 TIKV组成
基本结构如下图
![TIDB架构](imgs/tidb-architecture.png)

## TIDB
TIDB 是一个无状态可以动态扩展的服务， 本身并不存储任何数据
### 主要功能
1. 接收SQL请求
2. 处理SQL相关的逻辑
3. 与PD Server 通信定位数据实际存储位置
4. 与TIKV实际交换数据
5. 返回数据

## PD server
PD Server 负责整个集群信息的管理与通信， 通常要部署奇数台
### 主要功能
1. 存储集群数据元信息
2. 管理与负载均衡TIKV， 并负责数据迁移
3. 分配全局唯一单调递增ID


## TIKV
TIKV 是一种分布式事务型Key-Value型存储， 它是负责整个TIDB集群所有的实际数据存储的。
它使用Raft协议来保证集群数据的一致性与故障恢复，区域是它最基本的存储单元，所有不同节点
上的区域副本构成了一个Raft的分组

