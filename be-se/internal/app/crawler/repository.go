package crawler

import (
	"context"
	"fmt"
	"kn/se/internal/domain"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	// DatabaseName is exported.
	DatabaseName = "kn-be-se"

	// CollectionNameLinks is exported.
	CollectionNameLinks = "links"

	// CollectionNameArticles is exported.
	CollectionNameArticles = "articles"
)

// MongoRepository is exported.
type MongoRepository struct {
	client   *mongo.Client
	context  context.Context
	database *mongo.Database
	Cancel   func()

	links    *mongo.Collection
	articles *mongo.Collection
}

// NewMongoRepository is exported.
func NewMongoRepository() *MongoRepository {
	repo := MongoRepository{}

	// client
	host := os.Getenv("KN_MONGODB_HOST")
	if host == "" {
		host = "localhost"
	}
	portStr := os.Getenv("KN_MONGODB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 27017
	}
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(errors.Wrap(err, "database new client error"))
	}

	// context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	repo.Cancel = cancel

	// connection
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "database connection error"))
	}

	// database
	db := client.Database(DatabaseName)

	// collections
	ls := db.Collection(CollectionNameLinks)
	as := db.Collection(CollectionNameArticles)

	return &MongoRepository{
		// connection
		client:   client,
		context:  ctx,
		database: db,
		Cancel:   cancel,
		// collections
		links:    ls,
		articles: as,
	}
}

// Ping is exported.
func (s *MongoRepository) Ping() {
	err := s.client.Ping(s.context, readpref.Primary())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "database ping error"))
	}
}

// HasLink returns true if there is no matching link found.
func (s *MongoRepository) HasLink(hash string) bool {
	filter := bson.M{
		"hash": hash,
	}
	result := s.links.FindOne(s.context, filter)

	// if there is an error then we do not have a matching link
	// hence the given link is new so result is "not has link"
	err := result.Err()
	if err != nil {
		return false
	}

	// if we get to this point that means there was no error
	// and the given link is in the database
	return true
}

// SaveLink is exported.
func (s *MongoRepository) SaveLink(link *domain.Link) error {
	filter := bson.M{"hash": bson.M{"$eq": link.Hash}}

	o := options.UpdateOptions{}
	o.SetUpsert(true)

	updateData := bson.M{
		"$set": bson.M{
			"hash":       link.Hash,
			"uri":        link.URL,
			"is_article": link.IsArticle,
		},
	}

	_, err := s.links.UpdateOne(s.context, filter, updateData, &o)
	if err != nil {
		return errors.Wrap(err, "database error on save link")
	}

	return nil
}

// IsArticle is exported.
func (s *MongoRepository) IsArticle(hash string) bool {
	return false
}
