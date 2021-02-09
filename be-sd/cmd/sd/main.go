package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"kn/sd/internal/envd"
)

var sd *envd.SD

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	// prepare json or return error response
	js, err := json.Marshal(sd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// set content type
	w.Header().Set("Content-Type", "application/json")

	// write final response or return error response
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var err error
	sd, err = envd.NewSD()
	if err != nil {
		log.Println("error while parsing env variables %w", err)
	}

	http.HandleFunc("/", handlerIndex)

	log.Println("Application: kn-be-sd")
	log.Printf("Port: %s\n", sd.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", sd.Port), nil))
}
