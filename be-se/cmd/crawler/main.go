package main

import (
	"kn/se/internal/app/crawler"
	"log"
)

func main() {
	log.Println("Application: kn/se/crawler")

	sl := crawler.NewJSONSourceLoader("assets/initial_sources.json")
	c := crawler.NewCrawler(sl)
	c.Crawl()
}
