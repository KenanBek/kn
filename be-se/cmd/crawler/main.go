package main

import (
	"kn/se/internal/app/crawler"
	"log"
)

func main() {
	log.Println("Application: kn-se-crawler")

	sl := crawler.NewJSONSourceLoader("assets/initial_sources.json")
	s := crawler.NewCollyScraper()
	r := crawler.NewMongoRepository()

	c := crawler.NewWebCrawler(sl, s, r)
	c.Crawl()
}
