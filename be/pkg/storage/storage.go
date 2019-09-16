package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	databaseName = "x_be"
	postsName    = "posts"
	linksName    = "links"
	errorsName   = "errors"
)

// Storage is exported.
type Storage struct {
	client               *mongo.Client
	context              context.Context
	database             *mongo.Database
	posts, links, errors *mongo.Collection

	Cancel func()
}

// New is exported.
func New() Storage {
	storage := Storage{}

	// client
	host := os.Getenv("KN_HOST_MONGODB")
	if host == "" {
		host = "localhost"
	}
	uri := fmt.Sprintf("mongodb://%s:27017", host)
	clnt, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database new client error"))
	}

	// context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	storage.Cancel = cancel

	// connection
	err = clnt.Connect(ctx)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database connection error"))
	}

	// database
	db := clnt.Database(databaseName)

	// collections
	posts := db.Collection(postsName)
	links := db.Collection(linksName)
	errors := db.Collection(errorsName)

	return Storage{
		client:   clnt,
		context:  ctx,
		Cancel:   cancel,
		database: db,
		posts:    posts,
		links:    links,
		errors:   errors,
	}
}

// Ping is exported.
func (storage *Storage) Ping() {
	err := storage.client.Ping(storage.context, readpref.Primary())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database ping error"))
	}
}

// CheckLink is exported.
func (storage *Storage) CheckLink(uri string) bool {
	result := storage.links.FindOne(storage.context, bson.M{uri: uri})

	// if there is no error it means we have a matching link
	return result.Err() == nil
}
