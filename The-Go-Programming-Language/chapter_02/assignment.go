package main

import (
	"fmt"
)

type Person struct {
	name    string
	sex     int
	address string
}

func assignment() {
	var x int
	x = 1 // 有名称的变量
	fmt.Printf("x = %d\n", x)

	bl := false
	var p *bool = &bl
	*p = true // 间接变量
	fmt.Printf("p point to %v\n", *p)

	var person Person
	person.name = "bob" // 结构体成员
	fmt.Printf("person is %v\n", person)

	var count [10]int
	var scale int = 1
	fmt.Printf("count is %v\n", count)
	count[x] = count[x] + scale // 数组或slice或map的元素
	fmt.Printf("count is %v\n", count)

	count[x] *= scale
	fmt.Printf("count is %v\n", count)
	v := 1
	fmt.Printf("v = %d\n", v)
	v++ // 递增
	fmt.Printf("after ++, v = %d\n", v)
	v-- // 递减
	fmt.Printf("after --, v = %d\n", v)
}

// 最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		// 等同于
		// temp := x % y
		// x = y
		// y = temp
	}

	return x
}

// 斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

// 多重赋值
func tupleAssignment() {
	var x int = 1
	var y int = 2
	fmt.Printf("x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("x = %d, y = %d\n", x, y)

	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a = %v\n", a)
	i, j := 8, 9
	a[i], a[j] = a[j], a[i]
	fmt.Printf("a = %v\n", a)

	g := gcd(12, 2)
	fmt.Printf("g is %d\n", g)

	fib8 := fib(8)
	fmt.Printf("fib8 is %d\n", fib8)

	m := make(map[string]string)
	m["China"] = "Beijing"
	m["Japan"] = "Tokyo"
	v, ok := m["china"] // map查询
	fmt.Printf("value is %s, result is %v\n", v, ok)

	var inter interface{}
	t, ok := inter.(int) // 类型断言
	fmt.Printf("value is %v, result is %v\n", t, ok)
	// v, ok = <-ch					// 通道接收

	// 将不需要的值赋给空标识符
	// _, err := io.Copy(dst, src)		// 丢弃字节个数
	// _, ok := x.(T)					// 检查类型但丢弃结果
}

// 可赋值性
func assignability() {
	// 复合类型的字面量表达式，隐式地给每一个元素赋值
	medals := []string{"gold", "silver", "bronze"}
	// 等同于
	// medals[0] = "gold"
	// medals[0] = "silver"
	// medals[0] = "bronze"
	fmt.Printf("medals is %v", medals)

	// 赋值只有在值对于变量类型是可赋值的时候才合法
}

func main() {
	assignment()

	tupleAssignment()

	assignability()
}
