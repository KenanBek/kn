package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

// JSONSourceLoader is exported.
type JSONSourceLoader struct {
	// Path to JSON file
	Path string
}

// NewJSONSourceLoader is exported.
func NewJSONSourceLoader(path string) *JSONSourceLoader {
	return &JSONSourceLoader{
		Path: path,
	}
}

// Load is exported.
func (jsl *JSONSourceLoader) Load() ([]Source, error) {
	file, err := os.Open(jsl.Path)
	if err != nil {
		return nil, errors.Wrap(err, "error while opening file")

	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(errors.Wrap(err, "Error on file close"))
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading file")
	}

	var srcs []Source
	err = json.Unmarshal(bytes, &srcs)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing JSON")
	}

	return srcs, nil
}
