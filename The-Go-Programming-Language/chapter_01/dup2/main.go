// dup2 打印输入中多次出现的行的个数和文本
// 它从stdin 或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		// map是一个使用make创建的数据结构的引用
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// 函数os.Open返回两个值：一个是打开的文件(*os.File)，第二个是内置的error类型的值
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 忽略input.Err()中可能的错误
}
