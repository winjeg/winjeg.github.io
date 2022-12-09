---
title: redis 简谈redis监控
date: 2017-12-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - redis
  - database
categories:
  - storage
  - database
---

### 简介
由于redis是一个内存型的数据库，关注的点势必跟MySQL等非内存型的数据库相差比较多， 因此有必要单独对redis的监控关注项进行梳理。
此外由于redis的单线程模型， 以及速度要求非常高，因此对于redis的监控需要因事制宜。

### 监控项

#### 内存
由于redis是内存型数据库，因此对于内存的需求是最大的需求， 监控项里面最重要的也是对于内存的监控，一般出问题，大概率是内存满掉的问题。
一般内存满了之后会导致一系列的问题，比如逐出了不该逐出的key、写不进数据、 超时阻塞等问题。 对于内存的监控是至关重要的。
一般一个合理的范围是在 30% 到70%之间。而超过了80%就需要报警和升级了。
```bash
10.1.9.164:7300> info memory
# Memory
used_memory:108476200
used_memory_human:103.45M
used_memory_rss:129523712
used_memory_rss_human:123.52M
used_memory_peak:109672760
used_memory_peak_human:104.59M
total_system_memory:16658898944
total_system_memory_human:15.51G
used_memory_lua:37888
used_memory_lua_human:37.00K
maxmemory:4294967296
maxmemory_human:4.00G
maxmemory_policy:volatile-lru
mem_fragmentation_ratio:1.19
mem_allocator:jemalloc-4.0.3
10.1.9.164:7300>
```

### 大Key， 慢查询
另外一个比较常见的导致线上问题的一大因素是对于redis的不合理的使用，如 大key与慢查询
一般而言redis的key与数据越大， 导致的查询时长就越长， 加上redis本身是单线程的模型，因此这类查询往往会影响到其他查询的正常进行， 对大key的控制一方面可以通过中间件或者proxy 等手段来截断或者拒绝， 一方面也是需要使用一定的手段如 redis的 `slowlog get`命令来查看。
有了能获取的方法，那自然就可以对这些东西进行监控并合理的进行报警。 一般来讲redis模型中， 还是以尽量不出现慢查询为宜， 一旦出现慢查询就应该立即报警。
```
10.1.9.164:7300> slowlog get
1) 1) (integer) 155
   2) (integer) 1545025300
   3) (integer) 19422
   4) 1) "COMMAND"
2) 1) (integer) 154
   2) (integer) 1544754674
   3) (integer) 11571
   4) 1) "scan"
      2) "26005"
      3) "MATCH"
      4) "*"
      5) "COUNT"
      6) "10000"

```

### 连接数
redis在与客户端进行通信也是维护了连接的， 这些连接用来处理服务端与redis server之间的命令发送与数据发送等等 另外redis的最大连接数是有上限的 可以通过命令 `config get maxclients`来取得， 也可以使用`info clients`来看连接的使用情况，一般而言，连接数只要使用不满其实是不会出现太多的问题的， 但一定用满了，则会导致应用端错误。 因此对于连接数的监控也是以比率为标准比较好， 建议合理的范围是 < 60%, 对于超过80%的情况应给予报警


```
127.0.0.1:6379> info clients
#Clients
connected_clients:621
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:0
127.0.0.1:6379>

127.0.0.1:6379> CONFIG GET maxclients
    ##1) "maxclients"
    ##2) "10000"
127.0.0.1:6379>

```
#### CPU
CPU对于redis来说也是比较重要的需要关注的一个项， 一般跟QPS直接相关， QPS越高，CPU使用的也就越高。
对于redis服务来讲， CPU 跟其他服务的要求差别也不大， 不正常区间是 <70% 超过80%则需要报警。
```
10.1.9.164:7300> info cpu
# CPU
used_cpu_sys:88088.66
used_cpu_user:39641.02
used_cpu_sys_children:40.78
used_cpu_user_children:131.68
10.1.9.164:7300>
```

#### QPS
QPS 即每秒钟Redis接受的请求数量，不同的Key的大小支持的QPS对于Redis是不同的， 因此单一的对redis的QPS进行监控可能意义不大。
如果要做redis QPS的监控应该与历史同期相比较，是否有剧增等情况。 还有一个可以参照点就是经验值， 一般这个配置的redis可以承受多少的QPS。 通过经验值与历史对比两个标准来决定是否需要报警。一般是历史同期的两倍，且比高峰期高的时候，是应该报警的。

QPS 可以通过计算两个时间间隔内执行的总命令数量来计算出来
```
10.1.9.164:7300> info  stats
# Stats
total_connections_received:215774
total_commands_processed:2980658450
instantaneous_ops_per_sec:50
total_net_input_bytes:1135673367887
total_net_output_bytes:594951090244
instantaneous_input_kbps:30.66
instantaneous_output_kbps:13.65
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:161591
evicted_keys:0
keyspace_hits:112552664
keyspace_misses:809125690
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:3349
migrate_cached_sockets:0
10.1.9.164:7300>
```

#### 网络IO
网络IO就是redis在处理命令的时候占用的网络的流入与流出的量， 一般而言， 网络IO是较少关注的一个点， 但在网络达到带宽的上限的时候我们还是应该重视起来， 因为这个时候其他的东西是会阻塞的， 一般推荐的报警值为超过网络带宽的 80%