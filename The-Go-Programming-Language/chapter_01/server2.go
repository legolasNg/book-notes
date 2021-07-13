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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/count", countHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 处理程序回显请求的URL的路径部分
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL.PATH = %q\n", r.URL.Path)
	fmt.Printf("count: %d\n", count)
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
