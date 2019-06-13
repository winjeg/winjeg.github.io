package main

func main() {
	// 定义
	a := map[int]string {
		1: "a",
	}
	// 取出元素
	b = a[1]

	// 判断是否包含
	_, ok := a[1] // 包含ok为true， 不包含为false

	//添加新元素
	a[2] = "b"

	// 遍历
	for k, v := range a  {
		println(k)	// key
		println(v) 	// value
	}

}