---
title: golang 学习笔记
toc: true
tags:
	- golang
	- lang
categories:
	- lang
---

# go-doc
golang 中文文档
本文主要分三部分来帮助初学者能够顺利入坑`golang`， 第一部分主要是讲golang语法， 第二部分主要是讲golang的工程化的东西，附带会推荐一些常用的三方包。
第三部分则主要是介绍这门语言的编程思想，常见的思路。
学习一门语言，主要是学习语法、工程化、以及这门语言的编程思想。
附录则会讲一些工具的使用
## golang 基本语法
参见[demo](demo)下面的， 基本10分钟看完就学会了

## golang 工程化
golang 工程化是教大家如何配置golang项目，如何组织代码结构，以及如何调试分析，进行单元测试与性能测试等等。工程化是一门语言离不开的话题，好的工程化与工具可以令人事半功倍。

### golang 开发工具
常见的golang开发工具有非常多，但有着比较好的代码补全，以及其他提高编程体验的IDE却为数不多。目前市面上比较流行的开发工具主要有如下几个，也可能有些高手手写vim插件使用vim之类的，此类的不做说明了。
1. goland
2. vscode
3. lite ide

以上三个IDE(集成开发环境) 我更倾向于使用jetbrains 的 goland, 虽然goland目前有成吨的bug, 但由于习惯问题，我还是习惯使用，大家也可以根据自己的习惯来使用相应的IDE。后面会有专门的文章对goland下如何配置golang项目进行详细说明。


### golang 项目组织结构
一般而言一个非库型golang项目会存在一个main.go, 用作程序的主入口。
同时也会存在一些模块，跟一些配置文件。这些配置文件与模块的推荐的放置方式如下。后面也会有专题去说明我们这样放置的道理与好处。
```
module1_dir
module2_dir
others.go
main.go
config_file
readme
```

### `dep`与 `go module`
 `dep` 是 `go module` 推出之前半官方的一个依赖管理工具，它使用起来非常简单， 它的总体思想是把依赖代码集成到项目中的vendor目录去，但不需要手动管理这些依赖。这样做有一个好处：在编译期不用下载额外的依赖。但也有个坏处：会使得代码仓库变得比较庞大， 除非使用ignore 忽略这个vendor目录
2. `go module` 是golang官方推出的在 `1.11`版本后用来取代其他各种依赖管理工具的官方工具，它的思想跟`maven`的思想有点像， 与`dep`不同的就是它会把依赖不放入项目中去，而是管理到 `GOPATH`下面。这样做可以节省很多的代码仓库的空间。在进行编译的时候提前把依赖下载好，就不存在编译的时候下载依赖的问题了， 是以后的方向。

### golang 生成各个平台下的可执行文件
golang 生成其他平台的二进制文件是非常非常简单的，只需要进行简单的一个环境变量的设置即可生成 mac/windows/linux 等平台下的二进制文件， 二进制文件可以直接运行而不需要像java一样需要一个 `jre`, 这也许就是golang的一个简单哲学。

### golang 单元测试与性能测试
golang提供了简单的单元测试与性能测试的功能， 虽然简单却很强大，后面会专门文档来说明如何去写单元测试与性能测试。
golang的单元测试与性能测试的文件需要以 `_test.go`结尾
单测方法和性能测试方法需要用类似如下方法声名。
```go
import "testing"

// 单元测试需要以Test开头的函数名称
func TestGenerateQrCode(t *testing.T) {
	GenerateQrCode()
	// t.Fail()
}

// 性能测试需要以Benchmark 开头的函数名称
func BenchmarkGenerateQrCode(b *testing.B) {
	GenerateQrCode()
	// b.ReportAllocs()
}

```

### golang 内存分析
`pprof`
### golang 远程调试
`dlv`
## golang 编程思想

## 附录
### go fmt

### go doc

### go mod


