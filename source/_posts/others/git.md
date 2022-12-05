---
title: Git提交规范
date: 2021-11-12 10:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59480122-fb5ba480-8e91-11e9-9c26-7f8d6d97868a.png
tags:
  - git
categories:
  - other
---


## 简介

Git是使用最广的代码管理工具，版本控制工具，也是大家最熟知的，如果不了解Git是什么以及怎么使用的，请参考临近的一些文档。
本文档主要讨论Git提交代码的一些推荐规范。

## 一、分支选择
一个分支一般对应一个比较明确的版本， 大家需要在这个分支上开发，继承各种功能点， 是许多Commit的集合。
为了能够更好的适应Git的特点与企业级的分支管理策略， 分支的命名就显得尤为重要。 比如Gitlab可以根据分支名称特点进行设置权限级别， 有些CICD工具根据分支工具进行约束部署行为， 所以合理的分支名称，是一个合格的代码开发者的基本素质要求。
推荐的分支名称

feature 功能点分支
release/production 发布分支
test/benchmark 测试分支
比较典型的用法如：feature/weixin_register 作为微信注册的一个功能点， feature/email_register 作为邮件注册的一个功能点。
而上线的时候可以用 release/user_register 作为用户注册的功能，集成之前的微信注册与邮件注册的功能点。

而在分支保护的时候，我们也可以轻松的将 release/* 设置为保护分支，仅仅允许固定的工具或者固定的人去提交和merge, 这样就能很好的控制线上在运行的代码的质量。

## 二、标签选择
git tag 也是一个比较重要的功能，往往用作一个比较长周期的，例如中间件的迭代，如java代码， 可以用 git tag 与maven版本号保持一致， 可以很方便的回溯代码。
这里推荐的tag命名方式为：vx.x.x 其中 v 代表版本的意思， 第一个 x 代表大版本号， 第二个代表小版本， 第三个代表小的修订版本, 这种命名方式，对于一些语言，如golang，就比较友好，golang是根据 tag来读取软件的版本的。

## 三、Commit 规范
1. 一个Commit只做一件事情

这是为了可以在出现问题的情况下可以随时对不同commit进行操作，且同时最大程度的降低对其他地方造成的影响。
同时这也是非常知名的一些仓库的一些普遍做法，如 linux kernel的维护方式就是这样的。
2. commit （标签）

commit 标签是为了更好的识别与分类commit的内容， 更好的组织commit本身. 常见的commit标签如下：
- bugfix 如` bugfix:fix user name not long enough problem.`
- doc 如 `doc: update user related api doc`
- improvement 如 `improvement: change the implementation of the - algrithm reduce exec time to 1/10`
- hotfix 用于紧急修复
- task 任务
- feature 功能特点