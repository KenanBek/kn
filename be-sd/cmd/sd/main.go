package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Application: kn-be-sd URL Input: %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Application: kn-be-sd")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
