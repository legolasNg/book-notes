// echo2输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 每次迭代，range产生一对值：索引和索引处元素的值
	// go不允许存在无用的临时变量，不然会出现编译错误。解决方案是使用空标识符
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
