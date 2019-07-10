---
title: yaml 高级语法笔记
toc: true
date: 2016-03-13 15:14:11
tags: [yaml, gramma]
categories: [gramma]
---

1. 缩进表示层级
2. 行内表示法数组用中括号，属性用花括号， 用逗号隔开
3. 非行内， 数组用短横线， 属性正常
4. !! 强制转换为字符串
5. 字符串 
    单引号和双引号都可以使用，双引号不会对特殊字符转义。
    字符串可以写成多行，从第二行开始，必须有一个单空格缩进。换行符会被转为空格。
    多行字符串可以使用|保留换行符，也可以使用>折叠换行（去掉转换的空格）。
    +表示保留文字块末尾的换行，-表示删除字符串末尾的换行

    锚点&和别名*，可以用来引用。
    &用于起别名，放在： 后面， 
    defaults: &defaults
    adapter:  postgres
    host:     localhost

    development:
    database: myapp_development
    <<: *defaults
