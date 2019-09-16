package main

import (
	"log"

	"kn/pkg/crawler"
	"kn/pkg/source"
	"kn/pkg/storage"
)

func main() {
	log.Println("KN")

	storage := storage.New()
	defer storage.Cancel()

	storage.Ping()

	sourceLinks := source.Parse()
	for _, sourceLink := range sourceLinks {
		crawler.Crawl(sourceLink)
	}
}
