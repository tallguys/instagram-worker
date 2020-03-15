package main

import (
	"instagram-worker/internal/app/service/crawler"

	"os"

	logger "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		logger.Error("Please input the username")
		return
	}

	username := os.Args[1]

	err := crawler.Crawl(username)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("finish")
}
