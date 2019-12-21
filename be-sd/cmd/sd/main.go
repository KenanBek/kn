package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"syscall"
)

const (
	// ENForBESdPort is exported.
	ENForBESdPort = "KN_BE_SD_PORT"
	// ENForBESd is exported.
	ENForBESd = "KN_BE_SD"
	// ENForBEApi is exported.
	ENForBEApi = "KN_BE_API"
	// ENForFEWeb is exported.
	ENForFEWeb = "KN_FE_WEB"
)

// SD struct used to keep information about internal services.
type SD struct {
	BESdPort string `json:"be_sd_port"`
	BESd     string `json:"be_sd"`
	BEApi    string `json:"be_api"`
	FEWeb    string `json:"fe_web"`
}

func NewSD() *SD {
	sdPort, _ := syscall.Getenv(ENForBESdPort)
	sdURL, _ := syscall.Getenv(ENForBESd)
	apiURL, _ := syscall.Getenv(ENForBEApi)
	webURL, _ := syscall.Getenv(ENForFEWeb)

	sd := SD{
		BESdPort: sdPort,
		BESd:     sdURL,
		BEApi:    apiURL,
		FEWeb:    webURL,
	}

	return &sd
}

var sd *SD

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
	sd = NewSD()

	http.HandleFunc("/", handlerIndex)

	log.Println("Application: kn-be-sd")
	log.Printf("Port: %s\n", sd.BESdPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", sd.BESdPort), nil))
}
