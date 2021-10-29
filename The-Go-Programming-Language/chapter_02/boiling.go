// boiling输出水的沸点
// 每个文件以package声明开头，表明文件属于哪个包
package main

import "fmt"

// 声明常量，是一个包级别的声明
// 包级别的实体名字不仅对于包含其声明的源文件可见，而且对于同一个包里面的所有源文件都可见
const boilingF = 212.0

// 函数声明包含一个名字、一个参数列表(由函数调用者提供的变量)、一个可选的返回值列表，以及函数体(其中包含逻辑语句)
func main() {
	// 声明属于main函数的局部变量
	// 局部声明仅仅是在声明所在函数内部可见，并且可能对于函数中的一小块区域可见
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
