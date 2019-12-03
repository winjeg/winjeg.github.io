---
title: Golang slice 与 数组详解
date: 2017-03-13 15:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - golang
categories:
  - lang
---

## golang slice 与数组详解

对于很多人来说golang的 slice 与数组很难分清楚具体的区别，不知道什么情况下是数组什么情况下是slice，两者有什么区别，怎么 使用它们， 本文就slice与数数组的区别做出一个详细的说明。

### 定义  
#### slice 定义  
slice的定义比较灵活， 可以直接用make来定义， 也可以直接初始化好
```go
	var a = []byte{1,2,3}		// 定义一个长度为3容量为3的slice， 并初始化好数据
	var b = make([]int, 0, 3)   // 定义一个长度为0， 容量为3的slice
```
#### 数组定义  
数据在定义的时候，就指定了长度，且长度与容量相等  
```go
	var a = [3]byte{1, 2, 3} // 定义一个长度为3的数组, 并初始化好
	var b = [2]byte{}         // 定义一个长度为2的数组, 初始化为 0
```

### 使用区别

#### slice  
slice 对于golang来说使用的更普遍一些，因为其灵活性非常好，可以自动扩容， 当然功能多，意味着性能相比功能单一的数组来说较低一些  
```go
    var a = []byte{1,2,3}
	fmt.Println(append(a, byte(1)))	// ok, slice 可以被append, 且slice 会自动扩容
    b[1] = int(a[2])	// ok, slice 也可以根据下标进行操作
    len(a) // 这时长度为4, 容量为8（2倍）
```

#### 数组  
数组一旦定义了就固定了类型使用起来也只能根据下标进行操作，不同长度的数组，其类型不同，即便其基类型相同，仍然是不同的类型  
```go
// append(a, b) 报错， 因为数组不能进行append
b[1] = a[2] // ok 可以直接根据数组下标进行操作
fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))	// false 不类型不同, 所以也不能进行赋值或者转换
```

### 相互转换  
尽管如此golang提供了方便的slice与数组相互转换的机制，下面是一些示例代码  
```go
func conv() {
	a := [3]byte{1,2,3}	// 定义一个数组

	// 数组转slice
	b := a[:]

	// slice 转数组， 其实这种情况很少
	c :=[3]byte{}
	copy(c[:], b)

	// 查看结果
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
}

```

输出结果  
```
[3]uint8
[]uint8
[3]uint8
[1 2 3]
```