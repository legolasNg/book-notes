// 每个源文件的开始用package声明
package main

// import声明必须跟在package后面，导入需要的包
import "fmt"

// 在main包中，函数main是特殊的，总是程序开始执行的地方
// 一个函数的声明，由func关键字、函数名、参数列表(main函数为空)、返回值列表、函数体组成
func main() {
	fmt.Println("hello, 世界")
}
