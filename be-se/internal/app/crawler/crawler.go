package crawler

import (
	"kn/se/internal/domain"
	"log"
	"regexp"

	"github.com/pkg/errors"
)

// Crawler is exported.
type Crawler interface {
	Crawl()
}

// SourceLoader is exported.
type SourceLoader interface {
	Load() ([]domain.Source, error)
}

// Scraper is exported.
type Scraper interface {
	GetLinks(url string) ([]string, error)
}

// Repository is exported.
type Repository interface {
	HasLink(hash string) bool
	SaveLink(link *domain.Link) error
	IsArticle(hash string) bool
}

// NewWebCrawler is exported.
func NewWebCrawler(sl SourceLoader, s Scraper, r Repository) WebCrawler {
	return WebCrawler{
		sl:         sl,
		scraper:    s,
		repository: r,
	}
}

// WebCrawler is exported.
type WebCrawler struct {
	sl         SourceLoader
	scraper    Scraper
	repository Repository
}

// Crawl is exported.
func (wc *WebCrawler) Crawl() {
	ss, err := wc.sl.Load()
	if err != nil {
		log.Fatal(errors.Wrap(err, "source load error"))
	}

	for _, s := range ss {
		ar, err := regexp.Compile(s.ArticleRegexp)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "error while compiling regexp"))
		}

		urls, err := wc.scraper.GetLinks(s.SourceURL)
		if err != nil {
			log.Println(err)
		}

		for _, url := range urls {
			h := domain.Hash(url)

			hasLink := wc.repository.HasLink(h)

			if !hasLink {
				link := domain.Link{
					Hash:      h,
					URL:       url,
					IsArticle: ar.MatchString(url),
				}

				err := wc.repository.SaveLink(&link)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
