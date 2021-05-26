// dup1输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map存储一个键/值对集合，键可以是其值能够进行相等(==)比较的任意类型
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// 在windows下使用ctrl+Z可以退出输入
	for input.Scan() {
		counts[input.Text()]++
	}

	// map里面的键值的迭代顺序不是固定的，通常是随机的，每次运行都不一致
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
