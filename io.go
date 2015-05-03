package metis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/elos/gen/metis/templates"
)

var (
	ImportPath = "github.com/elos/gen/metis"
	Path       = templates.PackagePath(ImportPath)
)

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
