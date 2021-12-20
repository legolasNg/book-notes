package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 使用io.Copy代替ioutil.ReadAll来复制响应内容到os.Stdout
func fetchByCopy() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 不要泄露资源

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: write %d: %v\n", n, err)
			os.Exit(1)
		}
	}
}

// 自动补全url地址
func urlCompletion(url string) string {
	urlPrefix := "http://"
	hasPrefix := strings.HasPrefix(url, urlPrefix)
	// 判断url参数是否缺失协议前缀
	if hasPrefix == false {
		url = urlPrefix + url
	}

	return url
}

// 获取请求时，自动补全url协议
func fetchHasPrefix() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(urlCompletion(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// 获取HTTP状态码
func fetchStatusCode() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "Status: %s", resp.Status)
	}
}

func main() {
	// fetchByCopy()
	// fetchHasPrefix()
	// fetchStatusCode()
}
