// echo1 输出命令行参数
package main

import (
	"fmt"
	"os"
)

// 命令行参数以os包中Args名字的变量供程序访问
func main() {
	// var关键字声明string类型的变量
	// 变量可以在声明的时候初始化，如果没有明确初始化，将隐式地初始化为这个类型的空值
	var s, sep string

	// 变量os.Args是一个字符串slice
	for i := 1; i < len(os.Args); i++ {
		// slice[i]访问单个元素
		// slice[m:n]访问连续子区间，索引使用半开区间，即包含第一个索引，不包含最后一个索引
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	// slice[m:n]如果m或n缺失，默认分别为0或len(slice)
	fmt.Println(os.Args[1:len(os.Args)])
	fmt.Println(os.Args[1:])
}
