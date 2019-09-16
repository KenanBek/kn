package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Link is type to keep information about source link.
type Link struct {
	ID            primitive.ObjectID `bson:"_id"            json:"id,omitempty"`
	SourceURL     string             `bson:"source_url"     json:"source_url"`
	ArticleRegexp string             `bson:"article_regexp" json:"article_regexp"`
	CreatedAt     time.Time          `bson:"created_at"     json:"created_at,omitempty"`
	UpdatedAt     time.Time          `bson:"updated_at"     json:"updated_at,omitempty"`
}
