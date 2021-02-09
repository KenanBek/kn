// Package crawler fetches web pages and searches for article links.
//go:generate mockgen -source=crawler.go -destination=crawler_mock_test.go -package=crawler
package crawler

import (
	"kn/se/internal/domain"
	"log"
	"regexp"

	"github.com/pkg/errors"
)

// Crawler is exported.
type Crawler interface {
	Crawl() error
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
func (wc *WebCrawler) Crawl() error {
	NewSessionLog()

	// Load the source links by the given source loader.
	ss, err := wc.sl.Load()
	if err != nil {
		return errors.Wrap(err, "source load error")
	}

	// Iterate over the source links and check internal links.
	// TODO: Check each iteration withing goroutine (#).
	for _, s := range ss {
		// Compile article regexp so it can be checked faster.
		ar, err := regexp.Compile(s.ArticleRegexp)
		if err != nil {
			return errors.Wrap(err, "error while compiling regexp")
		}

		// Get all the links within the source link.
		urls, err := wc.scraper.GetLinks(s.SourceURL)
		if err != nil {
			return errors.Wrap(err, "error while getting links for source link")
		}

		log.Println("found links: ", len(urls))

		// Iterate over the found links and check for article links.
		// All found links will be checked with database and if it is already
		// processed it will be skipped.
		for _, url := range urls {
			h := domain.Hash(url)
			hasLink := wc.repository.HasLink(h)

			log.Println("checking link: ", url)

			if !hasLink {
				link := domain.Link{
					Hash:      h,
					URL:       url,
					IsArticle: ar.MatchString(url),
				}

				log.Println("saving link: ", link)

				err := wc.repository.SaveLink(&link)
				if err != nil {
					return errors.Wrap(err, "error while saving link")
				}
			}
		}
	}

	return nil
}
