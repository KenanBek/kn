package main

import (
	"log"

	"be/pkg/crawl"
	"be/pkg/source"
)

func main() {
	log.Println("KN")

	sourceLinks := source.Parse()
	for _, sourceLink := range sourceLinks {
		crawl.CrawlSourceLinks(sourceLink)
	}
}
