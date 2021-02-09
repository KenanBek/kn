package crawler

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type SessionLog struct {
	es *elasticsearch.Client
}

func NewSessionLog() *SessionLog {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	return &SessionLog{es: es}
}
