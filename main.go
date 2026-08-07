// go-pingdom 对多个 URL 做 HTTP 健康检查，报告每个的可用性与响应耗时。
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

// check 对单个 URL 发 GET，返回状态行或错误。
func check(u string, timeout time.Duration) (string, error) {
	client := &http.Client{Timeout: timeout}
	start := time.Now()
	resp, err := client.Get(u)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return fmt.Sprintf("%d %v", resp.StatusCode, time.Since(start).Round(time.Millisecond)), nil
}

func main() {
	timeout := flag.Duration("t", 5*time.Second, "超时时间")
	flag.Parse()
	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "用法: go-pingdom <url1> <url2> ...")
		os.Exit(1)
	}
	for _, u := range urls {
		res, err := check(u, *timeout)
		if err != nil {
			fmt.Printf("FAIL  %s  %v\n", u, err)
			continue
		}
		fmt.Printf("OK    %s  %s\n", u, res)
	}
}
