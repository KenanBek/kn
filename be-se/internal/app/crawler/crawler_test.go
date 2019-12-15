package crawler

import (
	"kn/se/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestWebCrawler_Crawl_NoNewLinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	link1 := "http://example.com/1"
	link2 := "http://example.com/2"

	sl := NewMockSourceLoader(ctrl)
	sl.EXPECT().Load().Return([]domain.Source{
		{
			SourceURL:     "http://example.com",
			ArticleRegexp: "re",
		},
	}, nil)

	scraper := NewMockScraper(ctrl)
	scraper.EXPECT().GetLinks("http://example.com").Return([]string{
		link1,
		link2,
	}, nil)

	repository := NewMockRepository(ctrl)
	repository.EXPECT().HasLink(domain.Hash(link1)).Return(true)
	repository.EXPECT().HasLink(domain.Hash(link2)).Return(true)

	wc := NewWebCrawler(sl, scraper, repository)
	wc.Crawl()
}
