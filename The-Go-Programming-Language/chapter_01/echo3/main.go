package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用循环，+=语句通过追加旧的字符串、空格字符串和下一个参数，生成一个新的字符串。这样使用的代价比较大
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])
}
