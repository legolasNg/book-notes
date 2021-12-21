package tempconv0

import "fmt"

// 命名类型的底层类型决定了它的结构和表达方式，以及它支持的内部操作集合，这些内部操作与直接使用底层类型的情况相同
func Example_legal() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
}

// 即使底层类型相同，如果命名类型不同，他们也不能使用算术表达式进行比较和合并
func Example_illegal() {
	// 编译错误: 类型不匹配
	// boilingF := CToF(BoilingC)
	// fmt.Printf("%g\n", boilingF - FreezingC)
}

// 命名类型的值可以与其相同类型的值或者底层类型相同的未命名类型的值相比较
// 但是不同命名类型的值不能直接比较
func Example_compare() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(c == 0)
	fmt.Println(c == Celsius(f))

	// var zero float64
	// 编译错误: 类型不匹配
	// fmt.Println(f >= zero)
	// fmt.Println(f >= zero)
	// fmt.Println(c >= f)
}

func Example_test() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; 不需要显式调用字符串
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"；不调用字符串
	fmt.Println(float64(c)) // "100"；不调用字符串
}
