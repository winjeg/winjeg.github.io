package main
// 定义一个struct
type A struct {

}

// 别名
type B A

// 引用, C 就是A
type C = A

// 定义A的一个方法, 指针类型
func (*A) M1() {
}
// 定义A的一个方法， 非指针类型
func (A) M2() {
}

// 定义一个接口
type D interface {
	M2()
}
// ，由于这个接口A已经实现， 因此A 对象可以直接赋值给接口
var _ D = A(nil)
