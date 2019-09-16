package crawler

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"

	"kn/pkg/model"
)

// Crawl is exported.
func Crawl(sourceLink model.Link) {
	collector := colly.NewCollector()

	articleRegexp, err := regexp.Compile(sourceLink.ArticleRegexp)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while compiling regexp"))
	}

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if articleRegexp.MatchString(link) {
			log.Println("Found article", link)
		} else {
			log.Println("Not article", link)
		}
	})

	collector.Visit(sourceLink.SourceURL)
}
