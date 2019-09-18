package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SourceLink is exported.
type SourceLink struct {
	ID            primitive.ObjectID `bson:"_id"            json:"id,omitempty"`
	SourceURL     string             `bson:"source_url"     json:"source_url"`
	ArticleRegexp string             `bson:"article_regexp" json:"article_regexp"`
}

// URIHash is exported.
type URIHash struct {
	ID   primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Hash string             `bson:"hash" json:"hash"`
	URI  string             `bson:"uri"  json:"uri"`
}

// Post is exported.
type Post struct {
	ID   primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Hash string             `bson:"hash" json:"hash"`
	URI  string             `bson:"uri"  json:"uri"`
}

// Page is exported.
type Page struct {
	ID   primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Hash string             `bson:"hash" json:"hash"`
	URI  string             `bson:"uri"  json:"uri"`
}
