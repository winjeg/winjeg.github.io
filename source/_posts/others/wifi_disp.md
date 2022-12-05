---
title: Wifi DISPLAY 与Miracast那些事
date: 2021-12-13 10:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - multi-media
categories:
  - other
---

## 简介

Miracast本质就是一个基于Wi-Fi的网络应用。这个应用包括服务端和客户端。服务端和客户端必须支持RTP/RTSP等网络协议和相应的编解码技术。
Wi-Fi Display经常和Miracast联系在一起。实际上，Miracast是Wi-Fi联盟（Wi-Fi Alliance）对支持Wi-Fi Display功能的设备的认证名称。
通过Miracast认证的设备将在最大程度内保持对Wi-Fi Display功能的支持和兼容。

## Mircast 依赖的无线网卡的技术特性
1. Wi-Fi Direct，也就是Wi-Fi P2P。它支持在没有AP（Access Point, 热点， 即没有连接）的情况下，两个Wi-Fi设备直连并通信。
2. Wi-Fi Protected Setup：用于帮助用户自动配置Wi-Fi网络、添加Wi-Fi设备等。
3. 11n/WMM/WPA2：
    - 11n就是802.11n协议，它将11a和11g提供的Wi-Fi传输速率从56Mbps提升到300甚至600Mbps。
    - WMM是Wi-Fi Multimedia的缩写，是一种针对实时视音频数据的QoS服务。
    - WPA2意为Wi-Fi Protected Acess第二版，主要用来给传输的数据进行加密保护。
- Miracast一个重要功能就是支持Wi-Fi Direct。但它也考虑了无线网络环境中存在AP设备的情况下，设备之间的互联问题。读者可参考如图2所示的四种拓扑结构。
- Wi-Fi Direct：该功能由Android中的WifiP2pService来管理和控制。Wi-Fi Multimedia：为了支持Miracast，Android 4.2对MultiMedia系统也进行了修改。

发送端(Wifi Display, WFD)
只要网卡支持， 找个发送端应用就可以投屏

接收端(Wifi Sink Function)
只要网卡支持， 找个接收端应用打开，就可以设置
linux 里面甚至有可以打开的选项

参考资料
https://www.21ic.com/tougao/article/2883.html
https://www.pianshen.com/article/3595659134/
https://blog.csdn.net/shenghuo59/article/details/81981377