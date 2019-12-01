package crawler

import "fmt"

// Crawler is exported.
type Crawler interface {
	Crawl()
}

type Instance struct {
	sl SourceLoader
}

func (i *Instance) Crawl() {
	srcs, err := i.sl.Load()
	if err != nil {
		fmt.Println(err)
	}

}
