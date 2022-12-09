---
title: Big table 笔记
date: 2015-09-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags: 
    - bigtable
categories:
    - storage
    - bigtable
---

Big Table 是一种稀疏的多维分布式的有序Map

Big Table 的基本数据结构
![image](https://user-images.githubusercontent.com/7270177/59737028-a56e6e80-928e-11e9-8509-18415344cfcc.png)


Tablet Server 的服务方式
![image](https://user-images.githubusercontent.com/7270177/59737043-b4552100-928e-11e9-9da6-85a4d186ebc6.png)
Compaction 往往发生是为了合并一些数据，节省内存空间(memtable)， 与Log文件的空间需要进行Compation

其他一些基础组件
![image](https://user-images.githubusercontent.com/7270177/59737081-d0f15900-928e-11e9-9e8f-8715a5742311.png)

当一些基础的SSTable分开或者合并的时候， 读写仍然可以同步进行
