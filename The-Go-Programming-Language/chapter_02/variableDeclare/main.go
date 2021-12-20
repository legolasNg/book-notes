package main

import (
	"fmt"
	"math/rand"
	"os"
)

func varDeclare() {
	// var声明创建一个具体类型的变量，然后给它附加一个名字，设置它的初始值
	// 类型和表达式部分可以省略一个，但是不能都省略
	var name string = "test"
	fmt.Printf("name is: %s\n", name)

	// 如果表达式省略，其初始值对应于类型的零值
	// 零值 - 数字：0，布尔值：false，字符串：""，对于接口和引用类型(slice、指针、map、通道、函数)：nil
	// 对于一个像数组或结构体这样的复合结构，零值是其所有元素或成员的零值
	var num int
	var str string
	var boolean bool
	var inf interface{}
	var sli []int
	fmt.Printf("num is: %d\n", num)
	fmt.Printf("str is: %s\n", str)
	fmt.Printf("boolean is: %t\n", boolean)
	fmt.Printf("inf is: %v\n", inf)
	fmt.Printf("sli is: %v\n", sli)

	// 如果类型省略，它的类型将由初始化表达式决定
	var sli2 = make([]int, 10, 20)
	var sli3 = *new([]int)
	fmt.Printf("sli2 is: %v\n", sli2)
	fmt.Printf("sli3 is: %v\n", sli3)

	// 声明一个变量列表
	var i, j, k int					// int, int, int
	var b, f, s = true, 2.3, "four"	// bool, float64, string
	fmt.Printf("i = %d, j = %d, k = %d,\n", i, j, k)
	fmt.Printf("b = %t, f = %g, k = %s,\n", b, f, s)

	// 变量可以通过调用返回多个值的函数进行初始化
	var file, err = os.Open("./README.md")
	fmt.Printf("file = %v, err = %v\n", file, err)
}

func shortVarDeclare() {
	// 短变量声明，必须在函数中使用
	// name := expression
	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Printf("freq is %g, t is %g\n", freq, t);

	// 多个变量可以用短变量声明的方式声明和初始化
	i, j := 0, 1
	fmt.Printf("i = %d, j = %d\n", i, j)
	// := 表示声明，= 表示赋值
	i, j = j, i

	// 短变量声明也可以用来调用返回多个值的函数
	file, err := os.Open("./README.md")
	fmt.Printf("file = %v, err = %v\n", file, err)

	// 短变量声明不需要声明所有在左边的变量
	// 如果一些变量在同一个词法块中声明，那么对于那些变量，短变量的行等同于赋值
	in, err := os.Open("./README.md")
	out, err := os.Create("./README.md")
	fmt.Printf("in = %v, err = %v\n", in, err)
	fmt.Printf("in = %v, err = %v\n", out, err)

	// 短变量声明最少声明一个新变量，否则代码编译将无法通过
	f, err := os.Open("./README.md")
	fmt.Printf("f = %v, err = %v\n", f, err)
	/*
	f, err := os.Create("./Out.md")
	fmt.Printf("f = %v, err = %v\n", f, err)
	*/
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++	// 递增p所指向的值，p自身保持不变
	return *p
}

func pointerDeclare() {
	// 指针的值是一个变量的地址
	// 一个指针指示值所保存的位置，不是所有的值都有地址，但是所有的变量都有
	// 使用指针，可以在无须知道变量名字的情况下，间接读取或者更新变量的值
	var x int = 1
	p := &x				// p是整形指针，指向x
	fmt.Printf("value of pointer: %d\n", *p)
	*p = 2
	fmt.Printf("x = %d\n", x)

	// 指针类型的零值是nil
	// 指针是可以比较的，两个指针当且仅当指向同一个变量或者两者都是nil的情况下才相等
	var i, j int
	fmt.Println(&i == &i, &i == &j, &i == nil)

	var pointer = f()
	fmt.Println(pointer == f())

	// 传递一个指针参数给函数，能够让函数更新间接传递的变量值
	v := 1
	incr(&v)				// 副作用：现在v等于2
	fmt.Println(incr(&v))	// 现在v等于3
}

func newFunc() {

}

func main() {
	varDeclare()

	shortVarDeclare()

	pointerDeclare()
}
