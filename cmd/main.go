package main

import (
	"fmt"
	"instagram-worker/internal/app/service/crawler"
)

func main() {
	err := crawler.Crawl("nba")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finsih")
}
