package crawl

import (
	"crypto/sha256"
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"

	"be/pkg/model"
	"be/pkg/store"
)

// Hash is exported.
func Hash(url string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
	return hash
}

// SourceLinks is exported.
func SourceLinks(sourceLink model.SourceLink) {
	storage := store.New()
	defer storage.Cancel()

	collector := colly.NewCollector()

	postRegexp, err := regexp.Compile(sourceLink.ArticleRegexp)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while compiling regexp"))
	}

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		hash := Hash(link)
		if !storage.IsURIHashed(hash) {
			uriHash := model.URIHash{
				URI:  link,
				Hash: hash,
			}
			storage.AddURIHash(&uriHash)

			if postRegexp.MatchString(link) {
				log.Println("Found article", link)

			} else {
				log.Println("Not article", link)

			}
		} else {
			log.Println("Already processed", link)
		}
	})

	collector.Visit(sourceLink.SourceURL)
}
