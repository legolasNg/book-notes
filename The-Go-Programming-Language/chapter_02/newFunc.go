package main

import (
	"fmt"
)

func newInt() *int {
	return new(int)
}

func yetNewInt() *int {
	var dummy int
	return &dummy
}

func main() {
	// *int类型的p，指向未命名的int变量
	p := new(int)
	fmt.Printf("*p = %d, address: %p\n", *p, p)
	// 把未命名的int设置为2
	*p = 2
	fmt.Printf("*p = %d, address: %p\n", *p, p)

	// 两个newInt函数有相同的行为
	fmt.Printf("*newInt = %d\n", *newInt())
	fmt.Printf("*yetNewInt = %d\n", *yetNewInt())

	// 每次调用new返回一个具有唯一地址
	q := new(int)
	r := new(int)
	fmt.Println(q == r)

	structA := new(struct{})
	structB := new(struct{})
	fmt.Println(structA == structB)

	arrayA := new([0]int)
	arrayB := new([0]int)
	fmt.Println(arrayA == arrayB)
}
