package generic

import (
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

type Parser struct {
}

func (p *Parser) GetLinks(url string) ([]string, error) {
	var urls []string
	collector := colly.NewCollector()

	collector.OnHTML("a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		urls = append(urls, url)
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, errors.Wrap(err, "Error on page visit")
	}

	return urls, nil
}
