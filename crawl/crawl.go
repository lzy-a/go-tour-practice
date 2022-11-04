package crawl

import (
	"crawler/links"
	"fmt"
	"log"
)

// Crawl函数实现爬取内容
func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list

}
