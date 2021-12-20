package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 练习1.1: 输出命令的名字
func exercise11() {
	println(os.Args[0])

	println(strings.Join(os.Args[0:], " "))

	println(strings.Join(os.Args, " "))
}

// 练习1.2：输出参数的索引和值，每行一个
func exercise12() {
	s := ""
	sep := " "
	for index, arg := range os.Args {
		s = strconv.Itoa(index) + sep + arg
		println(s)
	}
}

// 练习1.3: 测量低效程序和使用strings.Join的程序在执行时间上的差异
func exercise13() {
	var begin, middle, end, duration int64
	begin = time.Now().UnixNano()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	middle = time.Now().UnixNano()
	duration = middle - begin

	fmt.Println(strings.Join(os.Args[1:], " "))

	end = time.Now().UnixNano()
	duration = end - middle

	println("string concat", begin, middle, duration)
	println("string join", middle, end, duration)
}

func main() {
	exercise11()

	exercise12()

	exercise13()
}
