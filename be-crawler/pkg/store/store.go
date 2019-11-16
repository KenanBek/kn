package store

import (
	"be/pkg/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"
    "strconv"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DatabaseName = "kn"

    CollectionNameLinks = "links"
    CollectionNamePosts = "posts"
)

// Store is exported.
type Store struct {
	client   *mongo.Client
	context  context.Context
	database *mongo.Database
	Cancel   func()

	links, posts *mongo.Collection

}

// New is exported.
func New() Store {
	store := Store{}

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
		log.Fatalln(errors.Wrap(err, "Database new client error"))
	}

	// context
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	store.Cancel = cancel

	// connection
	err = client.Connect(context)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database connection error"))
	}

	// database
	db := client.Database(DatabaseName)

	// collections
	links := db.Collection(CollectionNameLinks)
	posts := db.Collection(CollectionNamePosts)

	return Store{
        // connection
		client:   client,
		context:  context,
		database: db,
        Cancel:   cancel,
        // collections
		links:    links,
		posts:    posts,
	}
}

func (store *Store) Ping() {
	err := store.client.Ping(store.context, readpref.Primary())
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database ping error"))
	}
}

func (store *Store) CheckLinkByHash(hash string) bool {
	filter := bson.M{
		"hash": hash,
	}
	result := store.links.FindOne(store.context, filter)

	// if there is no error it means we have a matching link
	return result.Err() == nil
}

func (store *Store) AddLink(link *model.Link) {
	filter := bson.M{"hash": bson.M{"$eq": link.Hash}}

	options := options.UpdateOptions{}
	options.SetUpsert(true)

	updateData := bson.M{
		"$set": bson.M{
			"hash": link.Hash,
			"uri":  link.URI,
		},
	}

	_, err := store.links.UpdateOne(store.context, filter, updateData, &options)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Database URIHash update error"))
	}
}
