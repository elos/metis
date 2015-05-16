package metis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/elos/metis/templates"
)

var (
	ImportPath = "github.com/elos/metis"
	Path       = templates.PackagePath(ImportPath)
)

// Parses a glob pattern (i.e., "./models/*json") and returns
// a slice of said models, or an error if the pattern doesn't
// match any files
func ParseGlob(pattern string) ([]*Model, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("metis: pattern matches no files: %#q", pattern)
	}

	return ParseModelFiles(filenames...), nil
}

// ParseModelFiles
func ParseModelFiles(filenames ...string) []*Model {
	models := make([]*Model, len(filenames))
	for i := range filenames {
		models[i] = ParseModelFile(filenames[i])
	}
	return models
}

func ParseModelFile(filepath string) *Model {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	md := ModelDef{}
	err = json.Unmarshal(input, &md)
	if err != nil {
		log.Fatal(err)
	}

	if err := md.Valid(); err != nil {
		panic(err)
	}
	return md.Model()
}
