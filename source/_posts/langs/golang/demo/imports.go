// 声明包， main包里面包含main函数即可运行， 其他包名里面的main函数则不行
package main

// 第一种import 方法， 适合仅仅有一个导入项目
import "fmt"

// 第二种导入方法， 适合多个导入项目， 导入的时候尽量把项目本身的包与系统包三方包分开，并按字母序排好
import (
	_ "net/http"		// 仅仅执行 包里面的 init 方法， 但不使用包里的方法
	alias "os/exec"		// 使用别名引用包

	"somepackage/somepackage"
)

// 在进入main方法之前执行, 的特殊方法 init
func init() {
	fmt.Println("init...")
}

// 最终输出为
// init...
// main
func main() {
	fmt.Println("main")
}