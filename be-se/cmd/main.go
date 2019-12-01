package main

import (
	"log"

	"kn/se/pkg/crawl"
	"kn/se/pkg/source"
)

func main() {
	log.Println("KN")

	sourceLinks := source.Parse()
	for _, sourceLink := range sourceLinks {
		crawl.CrawlSourceLinks(sourceLink)
	}
}
