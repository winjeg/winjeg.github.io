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
mergetool 的设置， 下面以`Kdiff`为例， 当然merge工具也有非常多， 如 `meld`， `beyond compare` 等。
其设置方法都雷同，因此此处仅仅只举出一个例子作为说明。

```bash
git config --global  merge.tool kdiff3
git config mergetool.kdiff3.path "/usr/bin/kdiff3"
```
设置好mergetool之后，以后有merge冲突的时候， `kdiff3` 会自动跳出并让你人工merge。

### Rebase
```bash
git rebase -i HEAD~4 # 合并提交记录
git:(feature1) git rebase master # rebase 到master分支
git rebase --continue # 继续rebase
git rebase —abort # 放弃rebase回到原始状态

```
在合并提交记录的时候会有如下信息打印出来
```bash
s cacc52da add: qrcode
s f072ef48 update: indexeddb hack
s 4e84901a feat: add indexedDB floder
s 8f33126c feat: add test2.js

# Rebase 5f2452b2..8f33126c onto 5f2452b2 (4 commands)
#
# Commands:
# p, pick = use commit
# r, reword = use commit, but edit the commit message
# e, edit = use commit, but stop for amending
# s, squash = use commit, but meld into previous commit
# f, fixup = like "squash", but discard this commit's log message
# x, exec = run command (the rest of the line) using shell
# d, drop = remove commit
#
# These lines can be re-ordered; they are executed from top to bottom.
#
# If you remove a line here THAT COMMIT WILL BE LOST.
#
# However, if you remove everything, the rebase will be aborted.
#
```
选择其中一种作为合并的方式， 上述是在vim（也可以是其他设置的编辑器）的一个窗口中展示的。
如果你异常退出了 `vim`  窗口，不要紧张：
```bash
git rebase --edit-todo
```
这时候会一直处在这个编辑的模式里，我们可以回去继续编辑，修改完保存一下：

```
git rebase --continue
```

与其他版本管理工具类似， 下图比较形象的展示了`git`中 `rebase`与`merge`的区别
![image](https://user-images.githubusercontent.com/7270177/60004279-e1467100-969e-11e9-9d2f-16b6d437ac74.png)

### Git命令别名
git是一个比较开放的系统， 与bash类似， git可以自定义别名来取代冗长的命令行输入如可以设置 `git st` 代替 `git status`， 使用 `git l`代替 `git log` 等等， 这些都被定义在git的配置文件中(`~/.gitconfig`)， 修改起来非常方便。 


### Git 的GPG签名设置(Windows)

安装`gpg4win` 如果没有响相应的GPG的KEY， 利用这个工具生成相应的key与配置， 记得备份。  
如果是已有备份， 可以直接用这个工具导入，非常简单。

然而仅仅这样设置还是不够的， 你需要在`Github/Gitlab`上添加相应的 `PGP PUBLIC KEY BLOCK`  
提交的时候使用如下命令， 则会自动签名。

```bash
$ git commit -S -m "change readme"
ggpg: directory '/c/Users/winjeg/.gnupg' created
igpg: keybox '/c/Users/winjeg/.gnupg/pubring.kbx' created
gpg: skipped "winjeg <winjeg@qq.com>": No secret key
gpg: signing failed: No secret key
error: gpg failed to sign the data
fatal: failed to write commit object

```

如上产生的错误则是由于Git默认的寻找签名证书的程序的路径有误。按照下面的方法进行设置。

```bash
git config --global gpg.program "C:\Program Files (x86)\GnuPG\bin\gpg.exe"
```

设置完毕再次运行， 则可以看到成功签名`commit`

```bash
winjeg@gpc MINGW64 /d/projects/go/github.com/winjeg/cloudb (master)
$ git commit -S -m "change readme"
[master eca6b52] change readme
 1 file changed, 3 insertions(+), 1 deletion(-)
```

## Git 高级用法

### Git对象
接下来，新建一个空文件test.txt。

```bash
touch test.txt
```
然后，把这个文件加入 Git 仓库，也就是为test.txt的当前内容创建一个副本。
```bash
git hash-object -w test.txt
e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
```
上面代码中，`git hash-object`命令把`test.txt`的当前内容压缩成二进制文件，存入 Git。压缩后的二进制文件，称为一个 Git 对象，保存在.git/objects目录。

这个命令还会计算当前内容的 SHA1 哈希值（长度40的字符串），作为该对象的文件名。

查看文件对象的内容
```bash
git cat-file -p 3b18e512dba79e4c8300dd08aeb37f8e728b8dad
hello world
```
### 暂存区 (`git add`)
文件保存成二进制对象以后，还需要通知 Git 哪些文件发生了变动。所有变动的文件，Git 都记录在一个区域，叫做"暂存区"（英文叫做 index 或者 stage）。等到变动告一段落，再统一把暂存区里面的文件写入正式的版本历史。
```bash
git update-index --add --cacheinfo 100644 \
3b18e512dba79e4c8300dd08aeb37f8e728b8dad test.txt
git ls-files --stage
100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad 0   test.txt
```

### Git 快照 （`commit`）
暂存区保留本次变动的文件信息，等到修改了差不多了，就要把这些信息写入历史，这就相当于生成了当前项目的一个快照（snapshot）。

项目的历史就是由不同时点的快照构成。Git 可以将项目恢复到任意一个快照。快照在 Git 里面有一个专门名词，叫做 commit，生成快照又称为完成一次提交。

下文所有提到"快照"的地方，指的就是 commit。


### Git分支
Git分支其实是指向某个快照节点的指针， 对于Git来说， 分支的创建成本是极其低廉的。另外，Git 有一个特殊指针HEAD， 总是指向当前分支的最近一次快照。另外，Git 还提供简写方式，HEAD^指向 HEAD的前一个快照（父节点），HEAD~6则是HEAD之前的第6个快照。


---
本文将不对其他内容做过多介绍, 仅仅介绍到此为止
