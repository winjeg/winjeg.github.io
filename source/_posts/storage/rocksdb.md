---
title:  gorocksdb 的安装与使用
date: 2018-08-13 15:14:11
toc: true
# img: https://user-images.githubusercontent.com/7270177/59735413-0c3c5980-9288-11e9-8f32-d8e6836e65b6.png
tags: 
    - rocksdb
    - database 
categories:
    - storage
    - rocksdb
---

## 安装rocksdb
[官方参考安装方法](https://github.com/facebook/rocksdb)

去下载 rocksdb [最新的发行版](https://github.com/facebook/rocksdb/releases)
如下代码:
```bash
wget https://github.com/facebook/rocksdb/archive/v5.14.2.tar.gz
tar xpf  v5.14.2.tar.gz
cd rocksdb-5.14.2/
make shared_lib -j9
```
如果编译过程中出错如下
```
util/status.cc: In static member function ‘static const char* rocksdb::Status::CopyState(const char*)’:
util/status.cc:28:15: error: ‘char* strncpy(char*, const char*, size_t)’ output truncated before terminating nul copying as many bytes from a string as its length [-Werror=stringop-truncation]
   std::strncpy(result, state, cch - 1);
   ~~~~~~~~~~~~^~~~~~~~~~~~~~~~~~~~~~~~
util/status.cc:19:18: note: length computed here
       std::strlen(state) + 1; // +1 for the null terminator
       ~~~~~~~~~~~^~~~~~~
cc1plus: all warnings being treated as errors
make: *** [Makefile:650: shared-objects/util/status.o] Error 1
```
需要打开`util/status.cc`修改第 28行改为
```c++
    std::strncpy(result, state, cch);
```
然后继续编译完成
然后执行如下命令进行安装
```bash
cd includes
# 拷贝头文件到 include目录
cp -r rocksdb /usr/lib/
sudo su
cp librocksdb.so.5.14.2 /usr/lib/
cd /usr/lib
ln -sf librocksdb.so.5.14.2 librocksdb.so
ln -sf librocksdb.so.5.14.2 librocksdb.so.5
ln -sf librocksdb.so.5.14.2 librocksdb.so.5.14
```

## 安装其他依赖
```
zlib - a library for data compression.
bzip2 - a library for data compression.
lz4 - a library for extremely fast data compression.
snappy - a library for fast data compression.
zstandard - Fast real-time compression algorithm.
```

## 安装 gorocksdb
```
CGO_CFLAGS="-I/path/to/rocksdb/include" \
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
  go get github.com/tecbot/gorocksdb
```

## 测试代码

```golang
package main

import (
    "github.com/tecbot/gorocksdb"
    "log"
)

func Test() {
    opts := gorocksdb.NewDefaultOptions()
    opts.SetCreateIfMissing(true)
    opts.SetCompression(gorocksdb.NoCompression)
    opts.SetWriteBufferSize(671088640)
    db, err := gorocksdb.OpenDb(opts, "test")
    wopt := gorocksdb.NewDefaultWriteOptions()
    if err != nil {
        log.Printf("%v\n", err)
    }
    defer db.Close()
    db.Put(wopt, []byte("data"), []byte("value"))
}

```