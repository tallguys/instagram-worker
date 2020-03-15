package main

import (
	"instagram-worker/internal/app/service/crawler"

	logger "github.com/sirupsen/logrus"
)

func main() {
	err := crawler.Crawl("nba")
	if err != nil {
		logger.Error(err)
	}
	logger.Info("finish")
}
