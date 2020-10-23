---
title:  web 前端一些基本知识
date: 2019-06-18 15:14:11
toc: true
tags:
    - npm
    - project
categories:
    - lang
---

## IDE
vscode 是开发前段不二的IDE, 也是近年来维护最活跃的一个编辑器
### 插件
1. Auto Close Tag
2. Auto Rename Tag
3. Beautify
4. Bracket Pair Colorizer
5. Class autocomplete for HTML
6. Code Runner
7. Css peek
8. Document this
9. Eslint
10. TsLint
11. Image Preview
12. Node.js Module Intellisense
13. Path Intellisense
### 快捷键
1. ctrl+shift+b  构建项目
2. F5 run
3. F10 Step over
4. F11 step into
5. Shift F11 step out.
6. Alt+Shift+F format

## node与nmp
### 包安装与管理工具yarn 与npm

### 淘宝npm 镜像
#### cnpm
```
 npm install -g cnpm --registry=https://registry.npm.taobao.org
```
#### 别名的方式
```
alias cnpm="npm --registry=https://registry.npm.taobao.org \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npm.taobao.org/dist \
--userconfig=$HOME/.cnpmrc"

# Or alias it in .bashrc or .zshrc
$ echo '\n#alias for cnpm\nalias cnpm="npm --registry=https://registry.npm.taobao.org \
  --cache=$HOME/.npm/.cache/cnpm \
  --disturl=https://npm.taobao.org/dist \
  --userconfig=$HOME/.cnpmrc"' >> ~/.zshrc && source ~/.zshrc
```


## 项目结构
### package.json
#### 文件介绍
package.json 是项目的总的一个配置文件， 它定义了这个项目所需要的各种模块， 以及项目的基本配置信息。
可以自动生成或者手动编写， 自动生成的方法是用node
```
node init
```
#### scripts 段
scripts 段制定了运行脚本命令的npm命令行缩写， 比如start指定了运行npm run start 的时候所需要执行的命令
```
"scripts": {
    "dev": "node build/dev-server.js",
    "lint":"eslint --ext .js,.vue src test/unit/specs"
}
```
#### dependencies
depencies and devDependencies 分别指定了项目在运行时候依赖的模块与项目开发的时候需要的一些模块， 它们都指向同一个对象， 用来管理各种依赖
```
"dependencies": {
    "vue": "^2.2.2",
    "vue-router": "^2.2.0"
  },
  "devDependencies": {
    "autoprefixer": "^6.7.2"
  }
```
#### config 字段
config 字段用于向环境变量输出值
#### engines 字段
engines 字段主要声明node 与 npm的版本
#### bin 字段
bin字段 主要是为了让一个可执行命令安装到系统的路径， 可以直接调用,
比如，要使用hello作为命令时可以这么做：

```
{ "bin" : { "hello" : "./cli.js" } }
```
这么一来，当你安装hello程序，npm会从cli.js文件创建一个到/usr/local/bin/myapp的符号链接(这使你可以直接在命令行执行hello脚本)。

