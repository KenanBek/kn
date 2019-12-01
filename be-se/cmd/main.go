package main

import (
	"log"

	"kn/se/internal/crawler/source"
	"kn/se/pkg/crawl"
)

func main() {
	log.Println("KN")

	sourceLinks := source.Parse()
	for _, sourceLink := range sourceLinks {
		crawl.CrawlSourceLinks(sourceLink)
	}
}
