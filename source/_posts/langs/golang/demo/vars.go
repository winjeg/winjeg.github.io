package main

import "fmt"

// 简单全局变量
var a = 10
var f = 1.0

var (
	b = 0X123	// 十六进制
	c = "word"
)

// 复杂全局变量
var (
	// 数组
	slice = []string {
		"element1",
	}

	// map
	mapx = map[int]string {
		1:"abc",
		2:"def",
	}
)


func main() {

	g := "a"	// 定义并初始化一个变量

	fmt.Println(len(mapx))
	fmt.Println(len(slice))
	fmt.Println(a, b, c, f，g)
}