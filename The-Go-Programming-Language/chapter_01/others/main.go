package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type声明给已有类型命名
type Point struct {
	x, y int
}
var p Point

func coinFlip() int {
	timestamp := time.Now().Unix()
	source := rand.NewSource(timestamp)
	fmt.Printf("random source: %d\n", timestamp)
	random := rand.New(source)
	return random.Intn(2)
}

func sigNum(x int) {
	// 无标签选择，等价于 switch true
	switch {
	case x > 0:
		fmt.Println("the integer is > 0")
		// case语句不会从上到下贯穿执行，使用fallthrough可以改写这个行为
		fallthrough
	case x > 1:
		fmt.Println("the integer is > 1")
	// 默认的case语句可以放在任何地方
	default:
		fmt.Println("default case")
	case x > 2:
		fmt.Println("the integer is > 2")
		fallthrough
	case x > 3:
		fmt.Println("the integer is > 3")
	}
}

func main() {
	// 控制流: if 、for 、switch
	// 多路分支控制
	switch coinFlip() {
	case 0:
		fmt.Println("tails")
	case 1:
		fmt.Println("heads")
	default:
		fmt.Println("landed on edge!")
	}

	sigNum(0)

	// 使用&操作符可以获取变量的地址
	num := 1
	fmt.Printf("address of num: %p\n", &num)
	// 使用*操作符可以获取指针引用的变量的值
	p := &num
	fmt.Printf("value of point: %d\n", *p)
}
