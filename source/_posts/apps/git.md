---
title: 版本管理工具Git的使用
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - git
categories:
  - tools
---

## 什么是Git
Git 是一个开源的分布式版本控制系统，可以有效、高速地处理从很小到非常大的项目版本管理，是 Linus Torvalds 为了帮助管理 Linux 内核开发而开发的一个开放源码的版本控制软件， Git也是计算机界最重要的软件之一， 被广泛的适用于各大中小公司的各类项目中。


## Git的基础使用

### 暂存
```bash
git add file_name # 暂存某文件
git add . # 暂存所有变更
git reset # 取消暂存
git reset --hard # 取消本地所有未提交的更改
git checkout file # 取消某文件的更改
```

### 提交
```bash
git commit -m "the commit message"
```
提交时候，结果仍在本地，但已经属于创建了本地的一个变更集
### 推送
```bash
git push
```
推送代码到远端

### 分支管理

#### 新建分支
```bash
git branch -b new_branch
```
推送新建的分支到远端
```
git push origin new_branch
```

#### 切换分支
1. 切换到本地分支
```bash
git checkout branch_name
```
2. 切换到远程分支
```bash
git branch -b branch_name origin/branch_name # 设置本地分支与远程分支同步
git pull # 拉取远程分支代码
```

#### 合并分支
1. 无冲突合并
```bash
git merge branch_name
```

2. 有冲突合并
```bash
git merge branch_name
git mergetool
```

---
TODO

## Git 高级用法

### Rebase

### 节点操作


## Git 基本思想


## Git 工具


## Git API
