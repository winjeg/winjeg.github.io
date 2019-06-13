---
title: redis 使用规范
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59422638-49759700-8e03-11e9-8bdd-dd708ca69e1f.png
tags:
  - redis
  - database
categories:
  - storage
  - database
---

## key/value规范
1. 所有的key 应该有合理的业务前缀
2. key 不应该包含特殊字符，应尽量使用常见的字母，数字， 下划线等组合
3. key 的大小应该保证在`16kb` 之内
4. value 的大小应该保证在`1M`之内， 拒绝bigkey(防止网卡流量、慢查询)
5. string类型控制在10KB以内，hash、list、set、zset元素个数不要超过5000。
6. 控制key的生命周期，redis不是垃圾桶, 能过期的key一定要设置过期时间

## 数据类型使用注意
1. hash 不推荐使用 `hkeys` 命令除非hash表的尺寸非常小， 否则推荐使用scan
2. set/zset/list 需要遍历整个列表也尽量使用`scan`与 `lindex`等命令
3. 非字符串的bigkey，不要使用del删除，使用hscan、sscan、zscan方式渐进式删除，同时要注意防止bigkey过期时间自动删除问题(例如一个200万的zset设置1小时过期，会触发del操作，造成阻塞，而且该操作不会不出现在慢查询中(latency可查))
4. 选择适合的数据类型, 注意节省内存和性能之间的平衡

## 命令使用
1. O(N)命令关注N的数量
```
例如hgetall、lrange、smembers、zrange、sinter等并非不能使用，但是需要明确N的值。有遍历的需求可以使用hscan、sscan、zscan代替。
```
2. 禁用命令
```
禁止线上使用keys、flushall、flushdb等，通过redis的rename机制禁掉命令，或者使用scan的方式渐进式处理。
```
3. 不推荐使用select,  新业务禁止使用 select
4. 用批量操作提高效率, 但要注意控制一次批量操作的元素个数(例如500以内，实际也和元素字节数有关)。
5. Redis事务功能较弱，不建议过多使用
6. 必要情况下使用monitor命令时，要注意不要长时间使用

## 客户端使用
1. 避免多个应用使用一个Redis实例, 不相干的业务拆分，公共数据做服务化
2. 使用带有连接池的数据库，可以有效控制连接，同时提高效率
3. 高并发下建议客户端添加熔断功能(例如netflix hystrix)
4. 设置合理的密码
5. 根据自身业务类型，选好maxmemory-policy(最大内存淘汰策略)，设置好过期时间

```
volatile-lru: 即超过最大内存后，在过期键中使用lru算法进行key的剔除，保证不过期数据不被删除，但是可能会出现OOM问题
allkeys-lru：根据LRU算法删除键，不管数据有没有设置超时属性，直到腾出足够空间为止。
allkeys-random：随机删除所有键，直到腾出足够空间为止。
volatile-random:随机删除过期键，直到腾出足够空间为止。
volatile-ttl：根据键值对象的ttl属性，删除最近将要过期数据。如果没有，回退到noeviction策略。
noeviction：不会剔除任何数据，拒绝所有写入操作并返回客户端错误信息"(error) OOM command not allowed when used memory"，此时Redis只响应读操作。
```
