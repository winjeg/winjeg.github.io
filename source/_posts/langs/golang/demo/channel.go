package main

// 创建一个channel
var a = make(chan int)

func main() {
	go goRoutine(a)
	// 向channel放入数据
	a <- 10
	select {
	}
}

func goRoutine(x <-chan int) {
	// 从channel里面取数据
	val := <-x
	println(val)
}