package main

import (
	"crawler/crawl"
	"fmt"
)

func main() {
	workList := make(chan []string)
	unseenLink := make(chan string)
	// unseenLink <- "https://zhuanlan.zhihu.com/"
	go func() { workList <- []string{"https://go.dev/"} }()
	for i := 0; i < 50; i++ {
		go func() {
			for url := range unseenLink {
				urlList := crawl.Crawl(url)
				go func() { workList <- urlList }()
			}

		}()
	}
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLink <- link
			}
			fmt.Println(len(seen))
		}

	}
}
