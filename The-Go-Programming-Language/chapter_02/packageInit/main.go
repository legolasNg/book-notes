package main

import "fmt"

// 包的初始化，从初始化包级别的变量开始，这些变量按照声明顺序初始化
// 在依赖已解析完毕的情况下，根据依赖的顺序进行
var a = b + c // 最后把a初始化为3
var b = f()   // 通过调用f()接着把b初始化为2
var c = 1     // 首先初始化为1

func f() int {
	return c + 1
}

func main() {
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
