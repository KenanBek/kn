package crawler

// Crawler is exported.
type Crawler interface {
	Crawl()
}

type source struct {
	url    string
	regexp string
}

func (s *source) Crawl() {
	panic("Not implemented")
}
