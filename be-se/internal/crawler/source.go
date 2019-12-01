package crawler

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// INTERFACE

// SourceLoader is exported.
type SourceLoader interface {
	Load() ([]Source, error)
}

// Source is exported.
type Source struct {
	URL string `bson:"source_url"     json:"source_url"`
	// Regexp defines regular expression to match with file
	Regexp string `bson:"article_regexp" json:"article_regexp"`
}

// IMPLEMENTATION

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
	defer file.Close()

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
