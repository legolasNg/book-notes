package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// 对于每个传入的请求，服务器在不同的goroutine中运行该处理函数, 这样它可以同时处理多个请求
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 处理程序回显请求的URL的路径部分
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL.PATH = %q\n", r.URL.Path)
	fmt.Printf("count: %d\n", count)

	// 如果两个并发的请求试图同时更新计数值count，可能会不一致的增加，程序会产生一个严重的竞态bug
	// sync.Mutex可以确保最多只有一个goroutine在同一时间访问变量
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Printf("count: %d\n", count)
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

// counter回显目前为止调用的次数
func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
