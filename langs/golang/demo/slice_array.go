package main

// 数组与slice是有很大差别的， 下面的例子很好的证明了这一点
func main() {
	var s1 =  []int {1, 2}			// 定义一个slice
	var s2 = make([]int, 0, 2)		// 定义一个容量为2， 长度为0 的slice
	var s3 = make([]int, 2)			// 定义一个长度为2 的slice， 里面的初始元素都是0
	
	var a1  = [2]int{1, 2}		 // 定义一个数组
	var a2 = [1]int{1}
	s1 = append(s1, 1)	// ok, slice 可以append元素
	a1 = append(a1, 1)	// error, 数组不可以 append元素
	a1 = a2						// error, 同为数组， 但a1, a2 是不同的类型
	s1 = a1						// error, 数组不能赋值给 slice
	a1 = s1						// error,  slice不能赋值给数组
	s1 = nil					// 正确, slice 可以用nil来赋值
	a1 = nil					// 错误, 数组不能用nil 赋值

	// 数组与slice的截取操作比较类似， 左闭右开
	p1 := a1[1:]		// 第一个元素之后的部分
	p2 := s1[:1]		//第一个元素之前的部分
	p3 := s1[0:1]		// 从第0个元素，到第一个元素，不包含第一个元素
	fmt.Println(p1, p2, p3)	
}