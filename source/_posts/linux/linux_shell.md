---
title: Linux Shell简介
date: 2014-05-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
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


### 在shell中嵌入二进制文件 (TODO, fix )
将二进制文件打包到shell脚本
之前因为要用支付宝更新浏览器插件,直接下载了一个aliedit.sh脚本,直接执行脚本,便搞定了插件的安装,正要称赞阿里的开发人员人性化了,转念一下,一个shell脚本就能搞定的安装,岂不是可以直接cat脚本就可得知支付宝监控工具的代码啦.
直接cat结果如下:
```
123456789101112131415161718192021222324 main(){SetStringsMkdirARCHIVE=`awk '/^__ARCHIVE_BELOW__/ {print NR + 1; exit 0; }' "$0"`tail -n+$ARCHIVE "$0" | tar xzvm -C $TMP_DIR > /dev/null 2>&1 3>&1if [ $? -ne 0 ]thenecho $PACKAGE_BADQuitfiCUR_DIR=`pwd`cd $TMP_DIR./install.sh#cd "$CUR_DIR"rm -rf $TMP_DIRexit 0}main#This line must be the last line of the file__ARCHIVE_BELOW__2�^��^M�.�Ɠ��jz���Y�Zi(�#;S4#C^��?*oX#���`����jW�u��_���#p��#�`<span style="font-family: 'Lohit Hindi';">י��#n�UY������,c���d��II���
```

shell脚本后面跟了一些乱码,莫非是直接加密了shell,通过阅读代码可以看出脚本后的乱码其实一个tar.gz的二进制:
```
ARCHIVE=<code>awk '/^__ARCHIVE_BELOW__/ {print NR + 1; exit 0; }' "$0"tail -n+$ARCHIVE "$0" | tar xzvm -C $TMP_DIR > /dev/null 2>&1 3>&1
```
首先是用awk获取脚本代码的开始行号,使用tail获取所有二进制码(所以脚本才会有如此注释:#This line must be the last line of the file),通过管道传给tar命令解压到制定目录.

```
oen@oen ~/code/shell/aliedit/install $ du -a4 ./README12 ./install.sh268 ./lib/libaliedit64.so244 ./lib/libaliedit32.so516 ./lib536 .
```
如上可以看到,真正执行是通过解压到临时目录的install.sh 实现的,同时真正玄机在libaliedit64下,是看不到了.
不得不说的一个偶然出现的问题：脚本执行完成之后会把临时目录删除,通过vi注释掉删除语句,结果提示包错误,但是怀疑是难道是有脚本校验,但脚本包错误提示是因为tar失败发出的,原来是二进制的乱码通过vi编辑后保存,二进制便彻底变成乱码了,所以tar解包失败.
具体使用方式如下:
```
12345678 oen@oen ~/code/shell/aliedit/install/lib $ echo "oenhan.com blog code" > example.shoen@oen ~/code/shell/aliedit/install/lib $ tar -zcvm * >> example.shexample.shlibaliedit32.solibaliedit64.sooen@oen ~/code/shell/aliedit/install/lib $ cat example.sh | head -n 3oenhan.com blog code#A#3Q#��ePO�����-@p�#��kpw��5�#�<span style="font-family: 'Lohit Hindi';">ݝ#����u
```
其实这个脚本就是一个自解压包,同理你可以把很多文件的二进制搞出来,脚本中找个命令接受转义即可.
当然还有专门的命令可以搞定在脚本嵌入二进制文件: uuencode
1 uuencode /home/oenhan.com.1.tar /home/oenhan.com.2.tar  > /home/oenhan.com.txt
将 oenhan.com.1.tar编码到 oenhan.com.txt,将来解码到 oenhan.com.2.tar.
具体实现:
首先需要写一个脚本example.sh的头:
```
12345 #!/bin/bashuudecode $0cd /home/tar xvf  oenhan.com.2.tarexit
```
然后将自解压代码编码到脚本中
```
1 uuencode /home/oenhan.com.1.tar /home/oenhan.com.2.tar >> /home/example.sh
```
如此一个自解压脚本做成了, uuencode和tar解压没有本质区别,uudecode 自己完成了tar找寻二进制代码的过程,看似很自动化却需要用户安装一个包sharutils,从简易度上得不偿失,不如用tar的方式搞定shell的二进制代码嵌入.
