package domain

import (
	"crypto/sha256"
	"fmt"
)

// SourceLink is exported.
type SourceLink struct {
	SourceURL     string
	ArticleRegexp string
}

// Link is exported.
type Link struct {
	Hash      string
	URL       string
	IsArticle bool
}

// Article is exported
type Article struct {
	Hash  string
	URL   string
	Title string
}

// Hash returns hash value of the given URL
func Hash(url string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
	return hash
}
