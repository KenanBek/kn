package domain

import (
	"crypto/sha256"
	"fmt"
)

// Source is exported.
type Source struct {
	SourceURL     string `bson:"source_url" json:"source_url"`
	ArticleRegexp string `bson:"article_regexp" json:"article_regexp"`
}

// Link is exported.
type Link struct {
	Hash      string `bson:"hash" json:"hash"`
	URL       string `bson:"url" json:"url"`
	IsArticle bool   `bson:"is_article" json:"is_article"`
}

// Article is exported
type Article struct {
	Hash  string `bson:"hash" json:"hash"`
	URL   string `bson:"url" json:"url"`
	Title string `bson:"title" json:"title"`
}

// Hash returns hash value of the given URL
func Hash(url string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:6]
	return hash
}
