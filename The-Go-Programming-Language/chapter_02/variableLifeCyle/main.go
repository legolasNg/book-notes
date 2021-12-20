package main

import (
	"fmt"
)

var global *int

func foo() {
	// x使用堆空间，因为在foo函数返回之后还可以从global变量访问
	var x int
	x = 1
	global = &x
}

func bar() {
	// 当bar函数返回时，变量*y变得不可访问，可回收，所以编译器可以安全地在栈上分配*y
	y := new(int)
	*y = 1
}

func main() {
	// 变量t在每次for循环的开始创建，变量x在循环的每次迭代中创建
	for t := 0; t < 10; t++ {
		x := t + 1
		fmt.Printf("t = %d, x = %d\n", t, x)
	}

	foo()
}
