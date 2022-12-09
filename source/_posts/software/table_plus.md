---
title: 简而美的 国产MySQL GUI 工具 Tables+
date: 2022-11-19 15:14:11
toc: true
img: https://user-images.githubusercontent.com/7270177/202841457-415afc25-5d4a-4ddf-a568-8f5aa08012e9.png
cover: true
coverImg: /medias/featureimages/22.jpg
keywords: mysql
tags:
    - mysql
    - electron
categories:
    - software
---

Tables+ 是一款面向大众程序员开放的， 一款MySQL客户端， 它使用electron编写， 以保证此软件的跨平台可用， 另外，它是一款纯个人打造软件， 也希望大家多予以支持，以便此软件能够更好的发展下去。
    
 MySQL已经毫无疑问的成为了中小网站的首选存储方案，MySQL的工具，市面上现有的MySQL客户端还是比较多的。市面上现有的一些MySQL工具想必大家也不陌生。但毫无疑问，这些都是几十年的老工具了。很少会给人焕然一新的感觉，而且很多都已经与现在流行的界面元素脱节， Tables+ 的出现，无疑是弥补了这些缺失， 更重要的是它是一款相对完整的国产 MySQL客户端软件。

## 功能列表
下面通过一些典型的软件界面截图，来介绍此软件的功能, 附软件[下载地址](https://github.com/gridsx/gridsx.github.io/releases/tag/v1.0.0)
[备用下载地址](https://github.com/gridsx/gridsx.github.io/releases/tag/v1.0.0)


### 连接管理
![conn_mgr.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/2bd8ab64a35f4af982773cb269bc1ce8~tplv-k3u1fbpfcp-watermark.image?)
您可以通过连接管理->添加MySQL连接，进行MySQL连接信息的录入。

![home.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/2e9e4454702c4d61b1471def4d2140c3~tplv-k3u1fbpfcp-watermark.image?)
 录入完毕后，如果您需要删除或者修改，可以在软件首页，去删除这些已经录入的连接信息。 
 如果您录入了非常多的连接信息， 您可以通过首页右上角的搜索按钮，对您录入的连接进行搜索过滤。

![query.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/be0bcaa8a1b34206a620a9f58a280bd6~tplv-k3u1fbpfcp-watermark.image?)
点击首页卡片， 可以轻松进入数据库查询界面， 这里您可以输入SQL进行查询，查看执行计划， 也可以保存和格式化您的SQL。

![tables.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d92473e2b59e4cffb14ee2cc62511e42~tplv-k3u1fbpfcp-watermark.image?)
点击左侧表列表，您将看到此库下面的所有的表， 如果您的表数量过多，还可以进行分组管理哦，这是一个非常方便的功能， 这里可以进行常见的双击查看表数据，删除表， 新建表等操作， 也可以对打开的表的数据进行修改保存。


![design.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/3644067fe88345d1952fb53b6e3b1d1e~tplv-k3u1fbpfcp-watermark.image?)

在设计表界面， 您可以随意的修改表的列， 索引， 以及其他表选项， 以便满足您的表设计需求。


![user_sql.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/811d94f3b6754106bd7459a282afca62~tplv-k3u1fbpfcp-watermark.image?)
最后， 如果您想查看自己保存的SQL在哪，可以点击用户SQL，查看自己保存的SQL。
如果您想查看SQL内容，可以直接点击这里的卡片，进入查询页面，即可查看和查询。

![lang.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/e9bb421d847546df87604c2681706262~tplv-k3u1fbpfcp-watermark.image?)
最后， 如果您想对此软件的外观，和语言进行一定的修改，您可以通过文件-> 设置， 找到您能设置的一些选项。

## 设计理念：简单和美
Tables+ 对于多余的功能，绝对不设计， 比如不常用的函数， 不常用的存储过程， 以及触发器等等。
这些已经被很多互联网公司弃用， 因此，在此软件中不会特别的支持这些已经过时了的功能。 软件本身采用了比较现代的界面设计语言， 很多功能都经过精心调试和设计。


## 软件 RoadMap:

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7bdde8cb2f164dd0a73481ce44be1921~tplv-k3u1fbpfcp-watermark.image?)

##  贡献

如果您对此软件有疑问，或者有兴趣， 可以 [邮件](mailto://winjeg@qq.com) 我，说明您的意向。
