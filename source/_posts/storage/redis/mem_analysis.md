---
title: redis 内存分析工具 `rma4go`
date: 2018-11-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - redis
  - database
categories:
  - storage
  - database
---

## redis 简介
redis是一个很有名的内存型数据库，这里不做详细介绍。而`rma4go` (redis memory analyzer for golang) 是一个redis的内存分析工具，这个工具的主要作用是针对运行时期的redis进行内存的分析，统计redis中key的分布情况， 各种数据类型的使用情况，key的size，大key的数量及分布, key的过期状况分布等一些有助于定位redis使用问题的工具，希望这能够给应用开发者提供便利排查生产中所遇到的实际问题。


## `rma4go`的应用场景
redis是目前很流行的一个内存型数据库，很多企业都在使用。 但由于业界并没有很多对于redis使用上的规范，或者是有一些规范并没有被很好的遵循， 存在很多redis使用上的问题，我这边就列举一些例子：
1. redis 存用满了, 不知道key的分布情况，不知道来源于那个应用
2. redis 被block了，不知道什么原因导致的block，是哪个应用里的什么key的操作导致的
3. 想迁移redis数据，或者调整一些设置，但不知道要不要对redis里的数据进行保留，以及不知道什么业务在使用等
4. redis的key的过期情况不明朗， 不知道哪些东西可以删除或者调整
其实上面的一些问题是我随便列举出来的一些，并不是所有的存在的问题，相信也有很多其他场景同样会用到这样的一个redis内存分析工具`rma4go`

## rma4go的具体功能

### 数据维度
对于key的分析我们这个工具会提供如下几个维度的数据：

- key的数量分布维度
- key的过期分布维度
- key的类型分布维度
- key对应的的数据的大小分布维度
- key的前缀分布维度
- 慢key与大key的维度

当然以后如果发现有更好的纬度也会添加进去，目前先以这几个纬度为主

### 数据类型设计

```go
type RedisStat struct {
	All     KeyStat `json:"all"`
	String  KeyStat `json:"string"`
	Hash    KeyStat `json:"hash"`
	Set     KeyStat `json:"set"`
	List    KeyStat `json:"list"`
	ZSet    KeyStat `json:"zset"`
	Other   KeyStat `json:"other"`
	BigKeys KeyStat `json:"bigKeys"`
}

// distributions of keys of all prefixes
type Distribution struct {
	KeyPattern string `json:"pattern"`
	Metrics
}

// basic metrics of a group of key
type Metrics struct {
	KeyCount       int64 `json:"keyCount"`
	KeySize        int64 `json:"keySize"`
	DataSize       int64 `json:"dataSize"`
	KeyNeverExpire int64 `json:"neverExpire"`
	ExpireInHour   int64 `json:"expireInHour"`  // >= 0h < 1h
	ExpireInDay    int64 `json:"expireInDay"`   // >= 1h < 24h
	ExpireInWeek   int64 `json:"expireInWeek"`  // >= 1d < 7d
	ExpireOutWeek  int64 `json:"expireOutWeek"` // >= 7d
}

```

### 实现细节


#### key元信息

```
type KeyMeta struct {
	Key      string
	KeySize  int64
	DataSize int64
	Ttl      int64
	Type     string
}

```
众所周知， redis里的所有的数据基本都是由key的， 也是根据key进行操作的，那么对redis里的key进行分析我们必须要记录下来这个key的信息才可以做到， 我们能记录的信息正如以上结构中的一样， key本身， key的大小， 数据的大小， 过期时间以及key的类型。这些信息是我们对key进行分析的一个基础信息，都可以通过一些简单的redis命令就可以取到。

#### 遍历redis所有key
要对一个redis进行完整的key分析， 我们就需要有办法能够访问到所有key的源信息， 所幸redis提供了 `scan`这么一种方式可以比较轻量的遍历所有的key，访问到相应的key的元信息。
这样对于redis而言， 进行在线key分析的时候造成的压力也不会非常大，当然key分析不能再QPS高峰期进行， 需要在redis资源余量允许的情况下进行分析。

另外由于redis本身的一个内存清理机制，有25%的过期占用可以在分析key的时候被清理掉， 因此这个分析工具同时兼具了清理一部分内存的作用， 如果redis里面存在过期的而且存在于内存里面的key的话。

#### 对记录的信息进行分析与汇总
有了遍历所有key的方法， 又有了元数据， 剩下的事情就是把这些数据进行聚合汇总， 这个主要是一个算法上的工作，
最难的部分要数这个key聚合的部分了， 这里面有很多取舍， 由于作者我本人不是专攻算法的， 而且没有找到合适的库， 因此只能动手自己想了一种方式。 基本的思路是：

##### 压缩的算法
1. 对于每个新的key的元信息， 添加到老的key分析对象里去
2. 对这个key从后往前缩短， 去除尾部，看是否已经包含这个key的统计信息，如果包含， 则把key的信息累加上去， 如果不包含则创建一个新的纪录。
3. 当记录的个数添加到一定数量的时候， 对对象的个数进行一次压缩
    - 压缩的算法也是从字符串的末尾往字符串首部进行压缩
    - 当压缩不能增加这个pattern 的key的个数的时候使用原来的key（压缩前的key）
    - 当压缩可以增加这个pattern的key的个数的时候，进行key的合并，把pattern设置成压缩后的pattern
    - 当记录的条数超过指定的条数就循环往复，直到压缩到小于指定的条数为止
    - 如果对于key的最小长度（就算再压缩也要保留一两位）有要求， 有一些压缩到字符串的最小长度的参数可以进行调整与设置， 进行一定的取舍。
3. 直到scan完毕

##### 代码如下
```go
const (
	defaultSize = 128
	compactNum  = 30
	maxLeftNum =  150
	minKeyLenLower = 2
	minKeyLen   = 5
)


func (stat *KeyStat) compact() {
	distMap := stat.Distribution
	tmpMap := make(map[string][]string, defaultSize)
	shrinkTo := compactNum
	for k := range distMap {
		compactedKey := k
		if orgks, ok := tmpMap[compactedKey]; ok {
			orgks = append(orgks, k)
			tmpMap[compactedKey] = orgks
		} else {
			ks := make([]string, 0, defaultSize)
			ks = append(ks, k)
			tmpMap[compactedKey] = ks
		}
	}
	shrinkTo--
	for (len(tmpMap) > compactNum && shrinkTo >= minKeyLen) || (len(tmpMap) > maxLeftNum && shrinkTo >= minKeyLenLower) {
		tnMap := make(map[string][]string, defaultSize)
		for k := range tmpMap {
			// shrink
			if len(k) > shrinkTo {
				compactedKey := k[0:shrinkTo]
				if oik, ok := tnMap[compactedKey]; ok {
					oik = append(oik, tmpMap[k]...)
					tnMap[compactedKey] = oik

				} else {
					ks := make([]string, 0, defaultSize)
					ks = append(ks, tmpMap[k]...)
					tnMap[compactedKey] = ks
				}
			} else {
				tnMap[k] = tmpMap[k]
			}
		}

		// 如果此次shrink 没有使得这个集合的元素数量增加， 就使用原来的key
		for k := range tmpMap {
			if len(k) > shrinkTo {
				ck := k[0:shrinkTo]
				if len(tnMap[ck]) == len(tmpMap[k]) && len(tnMap[ck]) > 1 {
					x := make([]string, 0, defaultSize)
					tnMap[k] = append(x, tnMap[ck]...)
					delete(tnMap, ck)
				}
			}
		}
		tmpMap = tnMap
		shrinkTo --
	}

	dists := make(map[string]Distribution, defaultSize)
	for k, v := range tmpMap {
		if len(v) > 1 {
			var nd Distribution
			for _, dk := range v {
				d := distMap[dk]
				nd.KeyPattern = k + "*"
				nd.KeyCount += d.KeyCount
				nd.KeySize += d.KeySize
				nd.DataSize += d.DataSize
				nd.ExpireInHour += d.ExpireInHour
				nd.ExpireInWeek += d.ExpireInWeek
				nd.ExpireInDay += d.ExpireInDay
				nd.ExpireOutWeek += d.ExpireOutWeek
				nd.KeyNeverExpire += d.KeyNeverExpire
			}
			dists[k] = nd
		} else {
			for _, dk := range v {
				nd := distMap[dk]
				nd.KeyPattern = dk + "*"
				dists[dk] = nd
			}
		}
	}
	stat.Distribution = dists
}


```


## 在线key分析的github项目
[rma4go](https://github.com/winjeg/rma4go)
这是一个我已经写好的项目， 它使用起来非常简单

### 构建方法
1. 构建之前请确保golang sdk 已经安装， 并且版本 >=1.11.0
2. 请确保已经具备翻墙的环境， 因为它要下载一些依赖，可能来自墙外
翻墙方法如下
```bash
// linux/osx
export http_proxy=somehost:port
export https_proxy=somehost:port
// windows
set http_proxy=somehost:port
set https_proxy=somehost:port

```
3. 构建
```
git clone git@github.com:winjeg/rma4go.git
cd rma4go
go build .
```




### 使用方法
用法如下：`rma4go -h`
```
rma4go usage:
rma4go -r some_host -p 6379 -a password -d 0
======================================================
  -H string
        address of a redis (default "localhost")
  -a string
        password/auth of the redis
  -d int
        db of the redis to analyze
  -h    help content
  -p int
        port of the redis (default 6379)
  -r string
        address of a redis (default "localhost")
```

### 示例输出


```
all keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
string keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
list keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
hash keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
set keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
zset keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
other keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
big keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |

```
rendered by markdown 
total count 4004


all keys statistics

|                    PATTERN                    | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-----------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| TOP_TEN_NEW_XXXXXXXX*                         |       1 |       20 |      1529 |              0 |             0 |              0 |               0 |            1 |
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST*     |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| time-root:*                                   |      23 |      272 |       299 |              0 |             0 |              0 |               0 |           23 |
| DS_AXXXXXXXX_CORRECT*                         |       2 |       45 |        46 |              0 |             0 |              0 |               0 |            2 |
| time-2*                                       |     761 |     7528 |      9893 |              0 |             0 |              0 |               0 |          761 |
| time-level:*                                  |     537 |     8461 |      6981 |              0 |             0 |              0 |               0 |          537 |
| time-9*                                       |     102 |      901 |      1326 |              0 |             0 |              0 |               0 |          102 |
| time-7*                                       |     153 |     1372 |      1989 |              0 |             0 |              0 |               0 |          153 |
| DS_MAGIC_SUCC_2017-06-22*                     |       1 |       24 |       415 |              0 |             0 |              0 |               0 |            1 |
| tersssss*                                     |       5 |      124 |         0 |              0 |             0 |              0 |               0 |            5 |
| appoint_abcdefg_msgid*                        |       1 |       21 |         0 |              0 |             0 |              0 |               0 |            1 |
| BUSSINESSXXXXXXX_STATISTICS_NEED_CALC_RECENT* |       1 |       44 |         1 |              0 |             0 |              0 |               0 |            1 |
| switch_abcd_abcde*                            |       3 |       69 |         3 |              0 |             0 |              0 |               0 |            3 |
| abcdeferCounter_201*                          |       3 |       78 |         0 |              0 |             0 |              0 |               0 |            3 |
| diy1234567flag*                               |       1 |       14 |         1 |              0 |             0 |              0 |               0 |            1 |
| DS_PRXXBCD_LIST*                              |       1 |       15 |     17208 |              0 |             0 |              0 |               0 |            1 |
| time-4*                                       |     133 |     1194 |      1729 |              0 |             0 |              0 |               0 |          133 |
| datastatistics_switch_version0*               |       1 |       30 |         1 |              0 |             0 |              0 |               0 |            1 |
| register_count_2_201*                         |     592 |    15984 |       640 |              0 |             0 |              0 |               0 |          592 |
| canVisitNewabcdef1234PageLevels*              |       1 |       31 |         0 |              0 |             0 |              0 |               0 |            1 |
| YOUR_WEEK_VITALITY_INFO*                      |       1 |       23 |     75782 |              0 |             0 |              0 |               0 |            1 |
| time-8*                                       |     101 |      894 |      1313 |              0 |             0 |              0 |               0 |          101 |
| EXPERTS_APPOINT_INFO_MAP*                     |       1 |       24 |         0 |              0 |             0 |              0 |               0 |            1 |
| time-3*                                       |     130 |     1215 |      1690 |              0 |             0 |              0 |               0 |          130 |
| time-1*                                       |     943 |     9456 |     12259 |              0 |             0 |              0 |               0 |          943 |
| time-64*                                      |      87 |      781 |      1131 |              0 |             0 |              0 |               0 |           87 |
| time-5*                                       |     168 |     1516 |      2184 |              0 |             0 |              0 |               0 |          168 |
| total                                         |    4004 |    53422 |   7832490 |              0 |             0 |              0 |               0 |         4004 |


string keys statistics

|                    PATTERN                    | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-----------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| BUSSINESSXXXXXXX_STATISTICS_NEED_CALC_RECENT* |       1 |       44 |         1 |              0 |             0 |              0 |               0 |            1 |
| time-5*                                       |     130 |     1174 |      1690 |              0 |             0 |              0 |               0 |          130 |
| datastatistics_switch_version0*               |       1 |       30 |         1 |              0 |             0 |              0 |               0 |            1 |
| time-7*                                       |      39 |      348 |       507 |              0 |             0 |              0 |               0 |           39 |
| time-level:*                                  |     567 |     8939 |      7371 |              0 |             0 |              0 |               0 |          567 |
| diy1234567flag*                               |       1 |       14 |         1 |              0 |             0 |              0 |               0 |            1 |
| switch_abcd_abcde*                            |       3 |       69 |         3 |              0 |             0 |              0 |               0 |            3 |
| time-2*                                       |     598 |     5918 |      7774 |              0 |             0 |              0 |               0 |          598 |
| time-6*                                       |     125 |     1118 |      1625 |              0 |             0 |              0 |               0 |          125 |
| time-4*                                       |     136 |     1225 |      1768 |              0 |             0 |              0 |               0 |          136 |
| time-8*                                       |      72 |      636 |       936 |              0 |             0 |              0 |               0 |           72 |
| time-1*                                       |    1176 |    11814 |     15288 |              0 |             0 |              0 |               0 |         1176 |
| time-9*                                       |     100 |      880 |      1300 |              0 |             0 |              0 |               0 |          100 |
| time-root:*                                   |      23 |      272 |       299 |              0 |             0 |              0 |               0 |           23 |
| register_count_2_201*                         |     592 |    15984 |       640 |              0 |             0 |              0 |               0 |          592 |
| DS_AXXXXXXXX_CORRECT*                         |       1 |       20 |        20 |              0 |             0 |              0 |               0 |            1 |
| TOP_TEN_NEW_tersssss*                         |       1 |       20 |      1529 |              0 |             0 |              0 |               0 |            1 |
| time-3*                                       |     202 |     1925 |      2626 |              0 |             0 |              0 |               0 |          202 |
| total                                         |    3989 |    53042 |     46253 |              0 |             0 |              0 |               0 |         3989 |


list keys statistics

|                  PATTERN                  | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST* |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| DS_MAGIC_SUCC_2017-06-22*                 |       1 |       24 |       415 |              0 |             0 |              0 |               0 |            1 |
| DS_PRXXBCD_LIST*                          |       1 |       15 |     17208 |              0 |             0 |              0 |               0 |            1 |
| total                                     |       3 |       79 |   7710455 |              0 |             0 |              0 |               0 |            3 |


hash keys statistics

|           PATTERN            | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| tersssss_action_prepage_new* |       1 |       27 |         0 |              0 |             0 |              0 |               0 |            1 |
| YOUR_WEEK_VITALITY_INFO*     |       1 |       23 |     75782 |              0 |             0 |              0 |               0 |            1 |
| EXPERTS_APPOINT_INFO_MAP*    |       1 |       24 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2017-06-11*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssssHardTaskCounter*     |       1 |       23 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2018-04-27*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2017-09-01*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssssEasyTaskCounter*     |       1 |       23 |         0 |              0 |             0 |              0 |               0 |            1 |
| total                        |       8 |      198 |     75782 |              0 |             0 |              0 |               0 |            8 |


set keys statistics

|             PATTERN              | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|----------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| tersssss_bind_phone_phone*       |       1 |       25 |         0 |              0 |             0 |              0 |               0 |            1 |
| appoint_abcdefg_msgid*           |       1 |       21 |         0 |              0 |             0 |              0 |               0 |            1 |
| canVisitNewabcdef1234PageLevels* |       1 |       31 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssss_bind_phone_userid*      |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| total                            |       4 |      103 |         0 |              0 |             0 |              0 |               0 |            4 |


zset keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
other keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |


big keys statistics

|                  PATTERN                  | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST* |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| total                                     |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |





### 作为依赖使用
获取方法如下：
```
go get github.com/winjeg/rma4go
```

使用方法如下：

```go
func testFunc() {
	h := "localhost"
	a := ""
	p := 6379
	cli := client.BuildRedisClient(client.ConnInfo{
		Host: h,
		Auth: a,
		Port: p,
	}, cmder.GetDb())

	stat := analyzer.ScanAllKeys(cli)
    // print in command line
	stat.Print()
	// the object is ready to use
}

```

### github 维护(主要阵地)
1. 欢迎其他开发者加入
2. 欢迎提issue 反馈问题
3. 欢迎任何有意义的建议
4. 另外欢迎star，不建议fork，建议直接提交PR  ; )




