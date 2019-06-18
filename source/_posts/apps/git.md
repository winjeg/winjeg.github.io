---
title: 版本管理工具Git的使用
date: 2019-06-13 22:14:11
toc: true
thumbnail: https://user-images.githubusercontent.com/7270177/59480122-fb5ba480-8e91-11e9-9c26-7f8d6d97868a.png
tags:
  - git
categories:
  - tools
---

## 什么是Git
Git 是一个开源的分布式版本控制系统，可以有效、高速地处理从很小到非常大的项目版本管理，是 Linus Torvalds 为了帮助管理 Linux 内核开发而开发的一个开放源码的版本控制软件， Git也是计算机界最重要的软件之一， 被广泛的适用于各大中小公司的各类项目中。

小插曲: 为啥我们不提其他的版本控制系统， 因为对于一般的情况下， 有`Git`就足够了，它至少能满足99 %的人的需求。

## Git 配置
### SSH配置
1. 生成 RSA 秘钥对， 私钥自己保存， 公钥需要给
```bash
ssh-keygen -t rsa -b 4096 -C "winjeg@qq.com"
```
生成过程中使用的密码会更安全一些， 但设置就会更麻烦一些， 关于相关的设置， 希望大家“不厌其烦”， 去搜索引擎自己搜索就好。
生成结果一般有两个文件：
1. `id_rsa` 这是一个绝密的文件， 只有使用者自己知道， 其他人不能知道
2. `id_rsa.pub` 这个是一个公开的文件， 是发给外界用来安全通信的一个工具
对于Github或者Gitlab而言，均有地方添加 `public key`， 一般在 用户`settings` 菜单下

生成完毕之后， 把私钥放到相应的位置：
1. linux/mac  ~/.ssh  并设置`id_rsa`的权限 `chmod 600 ~/.ssh/id_rsa`
2. windows 用户直接把 `id_rsa` 放到 用户目录下的 .ssh 文件夹中即可

### 安装Git
1. windows 下安装
由于安装Git 比较简单，只需要去官方网站， 去下载并且按照默认步骤安装即可。
因此，此处不做更多详细的介绍。

如果想用`gpg`签名则比较复杂， 但注意一点， 如果出现 `key not avalible` 类似的， 尝试设置下gpg的位置

```bash
git config --global gpg.program "C:\Program Files (x86)\GnuPG\bin\gpg.exe"
```
2. Linux 下安装
```bash
sudo apt-get install git # debian based
sudo yum install git # redhat based
sudo pacman -S git # archlinux  based
sudo emerge git    # gentoo based
```

3. mac 下安装
我猜是：
```bash
brew install git
```

## Git的基础使用
### 新建Git 项目
####  克隆代码

```bash
git clone https://github.com/winjeg/demos-go
```

#### 新建本地项目，并关联到远程
```bash
git init repo_name   # 创建 repo_name 的文件夹， 并创建好相关的 .git 隐藏文件夹等
cd repo_name  
git remote add origin git@github.com:winjeg/repo.git # 设置远端地址(这个关系到推送的地址)
git add .  # 把当前的项目文件都暂存
git commit -m "Initial commit" # 把暂存的文件作为一次 commit  提交
git push -u origin master # 把commit push 到远程的master分支
```

经过以上步骤， 一个本地可以用的repo就建立好啦


### 拉取远端
```bash
# 拉取指定分支的变化
git fetch origin master 
# 拉取所有分支的变化
git fetch 
# 拉取所有分支的变化，并且将远端不存在的分支同步移除【推荐】
git fetch -p 
```

### 查看当前状态
```bash
git status
```
对于当前repo， 增加， 删除，修改等的状态都会被列出来
```
HEAD detached from fd07db2
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        deleted:    rops.yaml
        modified:   values.yaml

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        a.md

no changes added to commit (use "git add" and/or "git commit -a")
```

### 暂存
暂存文件是commit这些变更的前提
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
git checkout -b new_branch # 新建
git checkout new_branch # 切换到新建的分支
```

#### 删除分支
```bash
# 删除本地分支，如果本地还有未合并的代码，则不能删除
git branch -d qixiu/feature
# 强制删除本地分支
git branch -D qixiu/feature 
```

#### 推送新建的分支到远端
```bash
git push origin new_branch
```
#### 设置本地分支与远程同步
```bash
git branch --set-upstream-to=origin/<branch> hexo
```
#### 删除远程分支
```bash
# 等同于git push origin -d qixiu/feaure
git push origin :qixiu/feature
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
mergetool 的设置， 下面以`Kdiff`为例
```bash
git config --global  merge.tool kdiff3
git config mergetool.kdiff3.path "/usr/bin/kdiff3"
```
设置好mergetool之后，以后有merge冲突的时候， `kdiff3` 会自动跳出并让你人工merge

---
TODO

## Git 高级用法

### Rebase

### 节点操作


## Git 基本思想


## Git 工具


## Git API
