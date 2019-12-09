package domain

// SourceLink is exported.
type SourceLink struct {
	SourceURL     string
	ArticleRegexp string
}

// Link is exported.
type Link struct {
	Hash string
	URL  string
}

// Article is exported
type Article struct {
	Hash  string
	URL   string
	Title string
}
