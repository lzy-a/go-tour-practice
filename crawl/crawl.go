package crawl

import (
	"crawler/links"
	"fmt"
	"log"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
