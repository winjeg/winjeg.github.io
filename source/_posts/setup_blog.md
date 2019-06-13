---
title: 使用Hexo创建博客
toc: true
thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - blog
  - tutorial
categories:
  - other
---

在操作之前，
1. 安装好了相关的软件如： `node`
2. 拥有自己的github账号正确设置好`SSH KEY`
3. 创建一个 `username.github.io` 的repo 并开启相应的page设置

## 安装步骤
```bash
git clone git@github.com:your_name/your_name.github.io.git
git checkout -b hexo
```

### 安装 hexo
```bash
npm install -g hexo-cli
hexo init
```
由于Hexo要求必须在空文件夹中执行init操作，所以我们需要在博客文件夹以外的地方新建一个空文件夹，之后点击鼠标右键选择`Git bash Here`输入以下命令，并将命令执行完成后文件夹中的所有文件复制到`your_name.github.io`文件夹中
```
npm install
```
本地预览(可省略)
```
hexo generate
hexo server
```

### 远程部署
我们已经在本地成功建站，接下来我们要做的就是通过简单的修改配置文件使得Hexo为我们生成的静态页面能够部署到Github Pages上面。

1. 编辑username.github.io文件夹下面的_config.yml（Hexo官方文档中将其称为全局配置文件），找到deploy关键字，将其修改为
```yml
deploy:
  type: git
  repo: git@github.com:your_name/your_name.github.io.git
  branch: master
```
2. 为了将完成到Github的远程部署，我们还需要安装一个插件。
```bash
npm install hexo-deployer-git --save
```
3. 执行以下命令，完成静态页面的远程部署与博客源文件的备份
```bash
git add .
git commit -m "提交说明"
git push origin hexo
hexo generate -d
```
### 主题设置
请自行搜索github, 输入关键字 `hexo theme`选择自己喜爱的主题，并按照相关文档进行设置

### 对于其他设备上写博客

```bash
git clone your_repo
cd your_repo
npm install -g hexo-cli
npm install
npm install hexo-deployer-git --save
```

## 自动脚本编写
在源文件分支(`hexo`分支)上添加如下文件, 如果是windows 命名为`publish.cmd` 如果是linux 或者mac命名为`publish` 并加入可执行权限`chmod a+x publish`
```bash
git add .
git commit -m "add article"
git push origin hexo
hexo generate -d
```
编写完毕，运行脚本 windows,  `publish` 其他 `./publish`