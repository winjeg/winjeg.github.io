---
title: linux BootLoader Grub
date: 2014-05-13 15:14:11
toc: true
# img: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - linux
categories:
  - os
  - linux
---



   前段时间修复移动硬盘分区表的时候，把本本的分区也重新弄了下，导致ubuntu的grub找不到linux分区（由于root分区uuid改变造成的不能正常启动），切换到Linux系统时，提示错误“unknown filesystem”，进入grub rescue模式。
在Google上搜了一下，终于解决了，方法如下：
1. 首先使用set命令，查看当前配置信息；
2. 然后再使用ls命令，遍历一下所有的磁盘；
3.查找Linux操作系统的”/”分区所在的磁盘，可以使用“ls (hdx,x)/”，这里的hdx代表你的物理磁盘，如果只有一块硬盘，则x的值为0，后面一个x（也肯呢个是msdosx,是具体情况而定）代表“/”分区的编号。
执行该命令（注意，ls命令后的“/”不能少，否则会出现“bad filename”错误）后，如果结果为“unknown filesystem”，则说明不是Linux分区，继续查找，知道返回带有“ /boot”目录的分区。
4. 找到“/”挂载点所在的分区后，就可以修改启动分区了：
    grub rescue >root=(hdx,msdosx)
    grub rescue >prefix=(hdx,msdosx)/boot/grub
    grub rescue >insmod normal
    grub rescue >normal
执行完normal命令后，如果normal模块加载成功，那我们就可以看到久违的grub引导菜单了。此时，按“c”切换到grub的命令行模式,修改grub菜单：
   grub >root=(hdx,msdosx) //设置系统启动分区，在这里指向内核所在分区
   grub >prefix=(hdx,msdosx)
接下来加载Linux.mod模块，并将新的启动信息写入grub：
   grub >insmod (hdx,msdosx)/boot/gurb/linux.mod
   grub >linux /boot/vmlinuz-xxx-xxx root=/dev/sdax //里边的xxxx可以按Tab键
   grub >initrd /boot/initrd.img-xxx-xxx
5.执行boot命令，启动系统（如果系统不能启动，可以重复1-4步，多试几次）：
   grub >boot
6.正常启动系统后，在终端中输入“sudo update-grub”命令，重新生成“grub.ccfg”文件，更新grub信息，屏幕会出现“generating…”的信息。
如果没有安装grub-pc软件包，或者grub-legacy，会出现无法找到命令的错误。这时，只需安装一下grub-pc软件包即可。（注意：安装过程中会出现提示要不要新建grub到第一分区，由于我的本本第一分区是Windows系统，所以在此我选择“NO”，而是将grub建立在“/”挂载点所在的分区）
7.更新完毕之后，重启，问题解决了。如果问题还没解决，重复1-6步的同时，重新建立grub到第一硬盘mbr：
sudo grub-install /dev/sda