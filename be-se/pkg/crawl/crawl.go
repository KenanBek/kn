package crawl

import (
	"kn/se/internal/crawler/source"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"

	"kn/se/pkg/model"
)

// CrawlSourceLinks is exported.
func CrawlSourceLinks(sourceLink model.SourceLink) {
	storage := source.New()
	defer storage.Cancel()

	collector := colly.NewCollector()

	postRegexp, err := regexp.Compile(sourceLink.ArticleRegexp)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while compiling regexp"))
	}

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		uri := e.Attr("href")

		hash := Hash(uri)
		if !storage.CheckLinkByHash(hash) {
			link := model.Link{
				Hash: hash,
				URI:  uri,
			}
			storage.AddLink(&link)

			if postRegexp.MatchString(uri) {
				log.Println("Found article", uri)

			} else {
				log.Println("Not article", uri)

			}
		} else {
			log.Println("Already processed", uri)
		}
	})

	collector.Visit(sourceLink.SourceURL)
}
