package crawler

import (
	"fmt"
	"kn/se/internal/domain"
	"log"

	"github.com/pkg/errors"
)

// Crawler is exported.
type Crawler interface {
	Crawl()
}

// SourceLoader is exported.
type SourceLoader interface {
	Load() ([]Source, error)
}

// Repository is exported.
type Repository interface {
	CheckLink(hash string) (bool, error)
	SaveLink(link domain.Link) (bool, error)
	IsArticle(hash string) (bool, error)
}

// Parser is exported.
type Parser interface {
	GetLinks(url string) ([]string, error)
}

// NewCrawler is exported.
func NewCrawler(sl SourceLoader) Instance {
	return Instance{sl: sl}
}

// Source is exported.
// TODO: use domain model
type Source struct {
	URL    string `bson:"source_url"     json:"source_url"`
	Regexp string `bson:"article_regexp" json:"article_regexp"`
}

// Instance is exported.
type Instance struct {
	sl SourceLoader
}

// Crawl is exported.
func (i *Instance) Crawl() {
	ss, err := i.sl.Load()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Source load error"))
	}

	for _, s := range ss {
		fmt.Printf("%v\n", s)
	}
}
