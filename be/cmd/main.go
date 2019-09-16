package main

import (
	"log"

	"be/pkg/crawler"
	"be/pkg/source"
	"be/pkg/storage"
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
