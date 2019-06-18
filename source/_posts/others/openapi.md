---
title: 开放接口的一些小知识
date: 2018-12-13 10:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - biz
  - openapi
categories:
  -biz
---

## open 简介

open api 是什么？
Open API即开放API,也称开放接口 所谓的开放API（OpenAPI）是服务型网站常见的一种应用，网站的服务商将自己的网站服务封装成一系列API（Application Programming Interface，应用编程接口）开放出去，供第三方开发者使用，这种行为就叫做开放网站的API，所开放的API就被称作OpenAPI（开放API）。 （摘自百科，主要为了说一下概念）

## open api 设计常见问题

### 服务提供

#### 通信方式设计
常见的设计是：
必备的有：app key 与 app secret 签名
可选的有： 企业ID， 业务ID。 这些可选的东西一般用来区分业务领域或者权限而存在，确保不同调用者领域与权限的隔离。

app key 与app secret 是我见过最常见，也是必然会出现的一个设计， 在我对接过的十多种API（阿里、腾讯、小米）中， 这这些都是必备的。

##### app key 的作用
主要是为了防止请求被篡改以及用户识别。一般一个 app key 是用于做用户识别的， 服务提供方一般会存放调用方的一些信息， 并可以通过这个App key 检索到。

##### app secret 的作用
app secret 主要是用来签名请求的， 签名过的请求如果被修改之后， 则签名就会发生变化，攻击者是不可能知道这个变化的结果的，这样可以有效防止攻击者的攻击。

一般来讲，会对如下领域进行签名：
1. parameter （url 参数）
2. header （请求头）


一般的流程如下：
1. 用户在请求参数信息中加入时间信息
2. 用户使用自己的 app secret 对请求进行签名
3. 用户把签名的结果， app key 与其他参数一起放到请求里面传到服务提供方
4. 服务方根据 app key找到自己存放的 app secret， 并对请求进行签名
5. 比对签名信息是否一致， 不一致，则认认为请求受到了非法修改，直接拒绝服务。
6. 校验时间信息， 确保时间信息是在允许的范围内， 否则拒绝提供服务。



如何颁发？
一般是由服务提供方生成这样的一个键值对， 并把键值对安全的递给调用方。
以后 app secret 不会出现在网络传输中，只会用于双方的签名校验。


#### 数据格式与文档
1. 服务地址， 需要准确无误的告诉服务提供方的调用地址
2. 环境设置 （一般API提供方会提供线上与线下两种途径一个用来线上使用，一个用来调试）
3. 参数规定 （对于一个Open API一般来讲参数都是固定的， 不能随意变动，否则会引起双方不小的争执）
4. 结果格式 （结果格式，要写清楚所有出现的结果格式的可能性， 让调用者有办法可以提前对任何可能的结果进行处理）
5. 示例代码 （一段示例代码， 对于开发者是很友好的，大部分开发者喜欢看到这个）

#### 注意事项
1. 安全防范
2. 限流


#### 秘钥对设计


### 安全防范
#### 1. 存放攻击
这个是最常见的一种， 它不需要知道服务提供方与三方的秘钥， 只需要知道，双方的通信方式， 通过分析这种通信方式，来达到窃取信息，或者造成攻击的目的。
这种行为的防范方法也是有一些的：
在请求里面加上时间信息，对时间信息进行加密签名， 对时间设置可用时间段， 比如1分钟， 过了一分钟，攻击者获取到的信息就不能再用了。

#### 2. 秘钥泄露
秘钥泄露对于很多企业来说并不陌生， 一旦泄露对于双方的危害都很大， 而且如果不能及时发现会带来不小的损失， 这种人为泄漏， 很多时候并没有太好的办法。 可以做的事情比较现实的就是更换秘钥。


#### 3. 穿透攻击
调用服务方提供的Open API的三方合作者， 有可能对服务提供方造成压力过大的攻击， 这种攻击主要来源于两种途径：
1. 三方合作者应用程序问题导致请求放大， 导致的调用次数过高
2. 三方合作者把相关接口直接暴露到外部， 这样会带来潜在的问题， 在攻击者识别这样的接口之后，疯狂发起攻击， 服务提供方的服务器会承受巨大压力，甚至crash

对于穿透的攻击： 服务提供方能做的最常见的方案就是限流， 另外就是流量识别，
在流量异常的时候对API接口本身做一些限制

当然安全防范确实还有很多需要考虑的点，一般都会结合攻击的特点， 对于攻击类型进行定制防范， 对于业界常见的防范措施，一般都是要默认加入的。

## 个人的开源项目
[openapi](https://github.com/winjeg/openapi)
它主要解决的问题是： 简化服务端提供API的流程

1. 封装服务端API签名流程
2. 封装服务端颁发键值对的流程
3. 封装服务端签名校验机制
4. 提供简单的安全防范


这个工具是用golang设计的， 因此只适用于golang的项目。
由于这个工具是基于 `http.Request` 进行设计的， 因此理论上讲兼容所有的web框架， 如 `gin`, `iris`, `beego`等， 我在readme里面提供了`iris`的示例代码

 如果使用mysql来存放app key 和secret信息， 可以建如下的一个表
```sql
CREATE TABLE `app` (
  `app_key` varchar(32) NOT NULL,
  `app_secret` varchar(128) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`app_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
如果已经存在这样一个类似的表也可以不用建立，在代码中指定即可
```go
	r, err := openapi.CheckValid(req,
	// default implementation is via sql, to fetch the secrect
	    openapi.SqlSecretKeeper{
            Db:        store.GetDb(),   // 可以使用的 mysql 连接
            TableName: "app",       // 存放app key 和secrets的表名
            KeyCol:    "app_key",   // app key 的列名
            SecretCol: "app_secret", // app secret的列名
            AppKey:    k,           // 用户使用的 app key
	})
```

当然如果您已经封装好了app key 与secret的逻辑， 也可以自己实现如下接口
```go
// the interface to get the secret
type SecretKeeper interface {
	GetSecret() (string, error)
}

```

对于使用了web 框架的，只需要写一个middleware， 并启用就行了， 示例代码如下：

创建middleware
```go

// create a middle ware for iris
func OpenApiHandler(ctx iris.Context) {

    //sign header? to prevent header being modified by others
    // openapi.SignHeader(true)

	req := ctx.Request()
	// you can put the key somewhere in the header or url params
	k := ctx.URLParam("app_key")
	r, err := openapi.CheckValid(req,
	// default implementation is via sql, to fetch the secrect
	    openapi.SqlSecretKeeper{
            Db:        store.GetDb(),
            TableName: "app",       // the name of table where you store all your app  keys and  secretcs
            KeyCol:    "app_key",   // the column name of the app keys
            SecretCol: "app_secret", // the column name of the app secrets
            AppKey:    k,           // the app key that the client used
	})
	logError(err)
	if r {
	    // verfy success, continue the request
		ctx.Next()
	} else {
	    // verify fail, stop the request and return
		ctx.Text(err.Error())
		ctx.StopExecution()
		return
	}
}

```
启用middleware
```go
// use the middle ware somewhere
// so all the apis under this group should be
// called with signed result and app key
	openApiGroup := app.Party("/open")
	openApiGroup.Use(OpenApiHandler)
	{
		openApiGroup.Get("/app", func(ctx iris.Context) {
			ctx.Text("success")
		})
	}

```


是不是很简单，如果文中有误，或者缺失的内容欢迎各种批评教育。

如果您能读到这里， 我会感觉到十分荣幸，谢谢您的关注。


[谢谢支持](https://github.com/winjeg)