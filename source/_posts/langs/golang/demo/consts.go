package main

import (
	"fmt"
)

// 全局常量
const a = "abc"

const (
	b = 1
	c = 1.1
)

func main() {
	// 局部常量
	const d = 0x123
	fmt.Println(a, b, c, d)
}