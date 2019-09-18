package store

import (
	"be/pkg/model"
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

const (
	// DatabaseName ...
	DatabaseName = "x_be"
	// URIHashesColName ...
	URIHashesColName = "uris_hashes"
	// PostsColName ...
	PostsColName = "posts"
	// PagesColName ...
	PagesColName = "pages"
	// ErrorsColName ...
	ErrorsColName = "errors"
)

// Store is exported.
type Store struct {
	client                          *mongo.Client
	context                         context.Context
	database                        *mongo.Database
	uriHashes, posts, pages, errors *mongo.Collection

	Cancel func()
}

// New is exported.
func New() Store {
	store := Store{}

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
	store.Cancel = cancel

	// connection
	err = clnt.Connect(ctx)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database connection error"))
	}

	// database
	db := clnt.Database(DatabaseName)

	// collections
	uriHashes := db.Collection(URIHashesColName)
	posts := db.Collection(PostsColName)
	pages := db.Collection(PagesColName)
	errors := db.Collection(ErrorsColName)

	return Store{
		client:    clnt,
		context:   ctx,
		Cancel:    cancel,
		database:  db,
		uriHashes: uriHashes,
		posts:     posts,
		pages:     pages,
		errors:    errors,
	}
}

// Ping is exported.
func (store *Store) Ping() {
	err := store.client.Ping(store.context, readpref.Primary())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database ping error"))
	}
}

// IsURIHashed is exported.
func (store *Store) IsURIHashed(hash string) bool {
	filter := bson.M{
		"hash": hash,
	}
	result := store.uriHashes.FindOne(store.context, filter)

	// if there is no error it means we have a matching link
	return result.Err() == nil
}

// AddURIHash is exported.
func (store *Store) AddURIHash(uriHash *model.URIHash) {
	filter := bson.M{"hash": bson.M{"$eq": uriHash.Hash}}

	options := options.UpdateOptions{}
	options.SetUpsert(true)

	updateData := bson.M{
		"$set": bson.M{
			"hash": uriHash.Hash,
			"uri":  uriHash.URI,
		},
	}

	_, err := store.uriHashes.UpdateOne(store.context, filter, updateData, &options)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database URIHash update error"))
	}
}
