package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		if duplicateLine(f) == true {
			fmt.Println(arg)
		}
	}
}

// 判断是否有重复行
func duplicateLine(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			return true
		}
	}
	// 忽略input.Err()中可能的错误

	return false
}
