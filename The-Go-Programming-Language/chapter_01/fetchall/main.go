package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	fmt.Printf("url: %s\n", url)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄露资源

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs 	%7d		%s", secs, nBytes, url)
}

func main() {
	start := time.Now()
	// 通道是一种允许某一例程向另一历程传递指定类型的值的通信机制
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// goroutine是一个并发执行的函数
		go fetch(url, ch) // 启动一个goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道ch接收
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
