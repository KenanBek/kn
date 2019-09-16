package source

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"

	"be/pkg/model"
)

// Parse is exported.
func Parse() []model.SourceLink {
	log.Println("Parser started")

	file, err := os.Open("assets/initial_sources.json")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while opening file"))
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while reading file"))
	}

	var sourceLinks []model.SourceLink
	err = json.Unmarshal(bytes, &sourceLinks)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Error while parsing JSON"))
	}

	return sourceLinks
}
