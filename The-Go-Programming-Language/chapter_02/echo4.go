// echo4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 创建标识变量，有三个参数：标识名字，变量默认值，以及当用户提供非法标识、非法参数亦或-h和-help参数时输出的消息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	// 变量sep和n是指向标识变量的指针，必须通过*sep和*n来访问
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}