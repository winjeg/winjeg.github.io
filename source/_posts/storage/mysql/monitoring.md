---
title: Mysql 的监控
date: 2016-04-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - MySQL
  - database
categories:
  - storage
  - database
---

### MySQL监控的必要性
对于MySQL的稳定性而言，一个合理的监控是十分必要的， 监控是提前预知问题并解决问题的最后一道坎， 过了这道坎，服务或者应用基本就是在挂掉的边缘了。为了未雨绸缪，大部分比较知名的公司都回具备相对健全的监控体系。

### 关于监控项
 MySQL的监控项有非常多， 多大几十上百项， 这些项中的大部分都是在日常应用场景中是使用不到的， 也不能帮助大家定位MySQL的监控问题。对于日常或者生产中所遇到的问题基本上都可以由下文中的常见的几项指标来定位并排查出来。 当然对于这些监控项， 不同的监控项的重要程度也不同， 因为MySQL的第一要义在于存储数据，因此对于MySQL而言，IOPS对于MySQL而言是最重要的指标。此外，很多指标也都是相互影响的，比如慢查询80%的情况下是因为IO过多，而非CPU占用过多。
 
### 健康度
健康是一个相对的概念，就如一个人是不是好看， 你只能用相对量来形容，而不是绝对量。
但一直以来，对于MySQL 的健康度的计算方法其实是很少有人总结的，因此本文希望能把这个健康度这个概念总结出来，并给出一定科学的依据。

一个健康度的计算规则是否合理的唯一标准是它是否能真实的反映MySQL服务的健康状况，举例子来说，假设一个MySQL服务的健康度为50分， 那么这个MySQL中运行中必然会存在可以优化的地方，比如存在慢查询，连接数稍微偏多，在这种情况下虽然服务能正常服务， 但请求到达一定量的时候就可能使得这个健康度变为0分，直接导致服务挂掉。 对于一个很健康的Mysql服务，来讲， 它应该是各项指标在合理的范围内， 不存在潜在的问题。

健康度的计算也很考究， 因为你要真实反馈问题，能够通过一个值来判断整体的情况， 那么必须就得存在一个科学的计算方法， 由于各项指标的重要程度不一， 因此我们需要对各项指标加上一个`权重`， 但仅仅加权重还是不够的， 因为往往某一项指标也能决定服务的存活与否， 因此我们还需要有另外一个概念 `checkpoint`, 在不同的checkpiont各项指标的权重应该都不一样。
比如CPU在到达100%的 checkpoint的时候它的权重就应该提高， 因此引发健康度分数下降到10分以内， 甚至更低。因此我们要加的这个权重是 动态的， 根据指标的重要程度以及checkpoint 进行相应的变化。


## 监控指标
上文讲完健康度这一概念之后， 我们知道，要想完善一个监控， 必须要对我们的监控指标项目有比较清楚的了解。下面主要简单介绍一下各个参数的意义。

### IOPS
由于mysql是为了存储数据而存在的， 所有的查询直接或者间接的也是在操作存储， 而存储的硬件基础是磁盘，磁盘的一个重要的监控指标就是IO。 对于MySQL而言相关的指标就是IOPS. 顾名思义， IOPS是每秒的IO次数。 对于不同类型的磁盘的IOPS的支持情况不同， 一般的机械硬盘支持的IOPS为100左右， 像云服务提供商阿里云ECS的磁盘支持的IOPS为300左右， 一般固态硬盘支持的IOPS会有3000到5000， 对于RDS而言， IOPS是标明了总量的。

IO主要是用户的查询从磁盘存取数据来用的， 所有的查询经过解析后都回最终涉及到IO操作，IO操作的多少也决定了查询的快慢，以及每秒能满足的查询数量（QPS）

因为对于不同的存储介质而言QPS是不一样的，因此我们关注IOPS的使用率会更科学一点。

### CPU
CPU一直以来都是各种类型的服务一项比较重要的资源， 对于MySQL而言，CPU主要用来对查询进行解析，优化， 以及聚合分析等等，有非常多的地方会使用到CPU， 当一个服务的CPU占用过高的时候必然会影响到它正常提供服务的能力， 对于MySQL而言，各种规范都是不推荐讲耗费CPU的计算放在MySQL端来做的。但要支持那么多的查询以及其他的东西往往还是要消费CPU资源， 当CPU计算一个查询所涉及到的任务的时候也必然会影响到其他任务的计算。 一般而言，CPU使用率在60%以内都还算健康的状态。最合理的使用范围是30%~60%

### QPS
QPS (Query Per Seond)即每秒的查询数量，对于很多服务，都存在这么一个概念，如普通的业务服务每秒能支持多少次订单生成请求。QPS的高低与服务器资源的占用是正相关的。
过高的QPS会导致其他资源如CPU或者IOPS等资源的占用过高。而MySQL 对于QPS的支持不是确定的， 要根据具体运行的查询状况而定， 对于慢查询而言，一般运行一个或者两个就能占用非常多的CPU和IO资源， 对于效率比较高的查询（如按主键指定查询的）， 这类查询消费的CPU与IO资源相比而言非常小， 因此这个服务就可以支持更高的QPS

### 慢查询
慢查询是各类企业经常会遇到的一些问题， 慢查询的多少与数据库实例的健康状况也直接相关。虽然慢查询的数量与查询的时长是比较间接的指标， 但它也是不可忽略的一部分。
一般优秀的程序员都会控制自己的查询时长在50ms之内， 并且考虑到随着数据量增长这样的查询是不是能够保持住这个时间量级。业界也很少有对慢查询确定非常明确的标准。不同企业，要求高低不同， 一般而言我推荐将超过100-200ms的查询列入为慢查询。 慢查询是必须要优化的对象， 如果不优化，后期会不得不为慢查询付出额外的惨痛的代价。

### 活跃连接数
`活跃连接`是指在某一时刻，MySQL 与客户端之间在处理着查询请求的连接， 它可以直接反应这个MySQL服务的繁忙程度， 一般而言活跃连接越多， QPS也会越高， CPU与IO占用的也相应越多。

### 连接数
对于mysql 的连接而言， 有活跃的，自然也有不活跃的， 这些活跃的和不活跃的连接加起来基本上就是所有的连接了。因为MySQL Service 在服务端都回设置 `max_connection_num` 这个参数，
这个参数决定了允许所有连向这个MySQL服务的总的连接数。如果超出这个数字， MySQL服务端就会拒绝连接，客户端一般会报出 `get connection failed`的错误。日常环境中，不管是MySQL 还是Redis或者一些其他的知名服务， 就有连接池这个概念，连接池是为了提高连接的复用率， 降低连接的创建与销毁所消耗的资源而存在的。连接池的设置的恰当与否也往往决定了一个MySQL服务的连接的使用情况。 如果一个MySQL服务中存在大量Sleep的连接，就要考虑连接池设置的调整了。

### 内存
内存对于很多服务来讲也是非常重要的一个资源， 对于MySQL而言相对于其他的指标而言没有那么重要，因为MySQL的内存使用率，是应用程序无法直接决定的，Mysql会在很多地方使用到内存，如查询缓存， 连接缓存， 查询解析等等， 有很多一部分是为了加快查询的速度， 也有很多部分是为了维持MySQL的自身功能。一般而言MySQL服务的内存使用分配由DBA或者云服务提供商根据业界BP已经提前设置好了，是不需要更改的。 但对于内存使用的异常情况也是需要关注的， 因为它会影响到服务的正常使用。

### 网络IN/OUT 
网络IO情况主要包括两个方面， 一个是`入流量`，一个是`出流量`, 不管什么应用， 当网络带宽被占满的时候都会出现问题。因此这个也是一个重要指标， 而指标中主要关注的还是比率。

#### 百分比监控项
对于监控项与权重的说法如下：

| 监控项       | Check Point | 权重 | 描述   |
|-----------|-------------|----|------|
| IOPS使用率   | <30%        | 0  | 健康   |
|           | 30%~60%     | 1  | 正常   |
|           | 60%~80%     | 5  | 偏高   |
|           | 80%~95%     | 50 | 非常高  |
|           | >95%        | 90 | 致命   |
| CPU使用率    | <30%        | 0  | 健康   |
|           | 30%~60%     | 1  | 一般   |
|           | 60%~80%     | 5  | 偏高   |
|           | 80%~95%     | 70 | 非常高  |
|           | >95%        | 90 | 致命   |
| 连接使用率     | <40%        | 0  | 健康   |
|           | 40%~60%     | 5  | 正常   |
|           | 60%~80%     | 30 | 偏高   |
|           | 80%~95%     | 50 | 非常高  |
|           | >95%        | 70 | 急需优化 |
| 内存使用率 | <70%        | 0  | 健康   |
|           | 70%~90%     | 10 | 偏高   |
|           | >90%        | 20 | 急需优化  |
| 网络出口带宽使用率 | <60%        | 0  | 正常   |
|           | 60~90%      | 10 | 偏高   |
|           | >90%        | 30 | 急需优化  |


#### 其他监控项
活跃连接数，慢查询与QPS三项根据配置的不同，支持的也不同， 通常情况下如下来分
活跃连接就是客户端正在使用的连接，这类数据一般需要结合历史情况和现状来看， 如果一个数据库服务的活跃连接突然陡增，比历史同期要高好几倍（排除大促等情况），那肯定是不正常的， 有可能是应用端使用的问题， 也有可能是来自外部的攻击， 这个时候就要做相应的报警，以应对突发情况。

对于慢查询，任何公司的态度都是始终如一的， 不管如何，慢查询都是不应该出现在线上的OLTP数据库里的，因为数据库出问题最多的原因也就是慢查询， 70%左右的问题都是跟慢查询直接相关的， 慢查询可以导致CPU/IOPS等的使用剧增，往往不需要几个慢查询就能就能拖垮一个数据库。因此对于慢查询的监控也是十分必要的。但通常情况下慢查询应该也要区分对待， 否则出现的频率会非常大，容易出现报警太多别人不关注的情况。我这里推荐一下查询的区分与等级。
1. `<=50ms` 正常查询，一般不需要进行优化
2. `> 50ms && <=100ms` 稍微慢一点的查询，如果能优化可以想办法优化， 但不适宜比较大规模的调用，支持的QPS会比较低
3. `> 100ms && <= 300ms` 一般的还好的慢查询，这类慢查询往往出现在一些用了索引的统计任务里，就算有一般也应该加上缓存， 由任务定时刷新， 后期应该考虑迁移到其他解决方案上
4. `> 300ms <= 2000ms` 一般的慢查询， 对于这类慢查询， 如果出现就要尽快优化掉， 而且一定要控制触发这类查询的入口，不能让用户或者攻击者可以出发这类查询。 一旦留有口子， 数据库很容易被攻击挂掉。
5. `> 2000ms` 这类查询不应该出现，出现之后应该紧急优化上线。

一般很多应用开发者会经常问DBA的一个问题是这个配置的数据库能支持多少QPS， 其实这个是一个非常错误的问法。因为对于查询快慢， 数据库的支持情况是肯定不同的， 如果应用开发者只根据主键进行查询， 那一般可以支持比较高的QPS， 而且不需要很高的配置， 但如果都是慢查询， 可能支持的QPS也就只有个位数了， 所以大家一般问这类问题的时候，DBA给的值一般是按照经验，然后假定应用的查询还可以的情况下，给出的一个相对保守的值。


