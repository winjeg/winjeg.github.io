---
title: Linux Shell简介
date: 2015-03-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
categories:
  - os
  - linux
---

## 简介
### 背景知识
Linux 的Shell 是Linux入门者必须绕不过的一个坎，对于很多初学者Shell是一个噩梦， 但对于很多有经验的用户，一个没有Shell的系统是最不好用的系统。 Linux、Windows、Mac这三个主流操作系统都是由Shell的存在的不管是Windows的Cmd， PowerShell， 或者是Mac的Terminal， 异或是Linux的各种Terminal。但在Linux的Terminal上你几乎可以无所不能，你可以几乎修改任何系统的设置， 也可以完成一系列复杂的操作， 比如处理图片，处理视频，也可以唤起任何图形界面的程序，只要你系统配置好了， 可以这么说，Linux可以没有图形界面（GUI）， 但不能没有Shell。个人理解Linux的Shell只要配置得当可以甩Windows一大截。

## shell 的作用
用一句话来概括，就是Shell是你与Linux进行交互的主要渠道之一，主要用来操作Linux系统的方方面面。

另外Shell有着自己的一套脚本语言， 有了这个语言，你也几乎可以做任何自动化的事情， 当然此文不会详细讲Shell编程，因为笔者认为你就算不会Shell编程也可以使用Linux， 完全没有什么问题。


## 常见的Shell
### sh
sh是几乎所有发行版必备的一个Shell, 但它可能不是默认的Shell， 但你总能唤起它。
### bash
bash是几乎所有Linux发行版都会默认的Shell， 是笔者自己会用的Shell
### zsh
zsh 是对Shell期望比较高的一帮人搞出来的一个东西， 它有着自动补全，自动纠错，还有一些自动目录跳转的功能，当然他的功能也不仅仅于此， 当你熟悉了Linux你也会对这个Shell工具十分感兴趣


## 常见的Shell软件

### gnome-terminal
gnome-terminal 是Gnome桌面环境默认自带的terminal，一般的功能都具有，比如支持多彩显示，主题设置，cursor设置， 背景，前景，字体，粗细，字符集等等

### xfce4-terminal
xfce4-terminal 是xfce4桌面环境自带的terminal， 一般情况下你只要装了xfce4桌面环境套装，这个terminal 就存在了。它的功能与Gnome的termianl很类似

### xterm
xterm 的历史比较悠久，很多发行版都自带Xterm，作为一个默认的terminal

### 其他发行版terminal
因为Terminal 的软件实在是数不胜数，因此我也在这里无法列出全部的详细的Terminal， 更没有能力进行一一点评， 选择Terminal也是要按照个人的喜好进行选择，主流发行版/桌面环境里面的Terminal一般都还是不错的，功能也都类似， 也不必刻意去找一个更好的取代品。

### TODO Gnome terminal 的配置
