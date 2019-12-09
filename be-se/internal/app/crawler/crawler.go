package crawler

import (
	"fmt"
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

func NewCrawler(sl SourceLoader) Instance {
	return Instance{sl: sl}
}

// Source is exported.
type Source struct {
	URL    string `bson:"source_url"     json:"source_url"`
	Regexp string `bson:"article_regexp" json:"article_regexp"`
}

type Instance struct {
	sl SourceLoader
}

func (i *Instance) Crawl() {
	srcs, err := i.sl.Load()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Source load error"))
	}

	for _, s := range srcs {
		fmt.Printf("%v\n", s)
	}
}
