package metis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/elos/ehttp/templates"
)

type Producer interface {
	WriteFile(string)
}

// The relative import path of this library
const ImportPath = "github.com/elos/metis"

// The absolute file path of this package
var Path = templates.PackagePath(ImportPath)

// ParseGlob follows a glob pattern (i.e., "./models/*.json") and returns
// a slice of said models, or an error if the pattern doesn't
// match any files
func ParseGlob(pattern string) ([]*Model, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// Consider no match an error
	if len(filenames) == 0 {
		return nil, fmt.Errorf("metis: pattern matches no files: %#q", pattern)
	}

	return ParseModelFiles(filenames...)
}

// ParseModelFiles takes an array of file names and returns the parsed models
func ParseModelFiles(filenames ...string) ([]*Model, error) {
	models := make([]*Model, len(filenames))
	for i := range filenames {
		tm, err := ParseModelFile(filenames[i])

		if err != nil {
			return models, err
		}

		models[i] = tm
	}
	return models, nil
}

// ParseModelFile takes a model definition file and parses it, returns an
// error if it the JSON was invalid or the definition was invalid
func ParseModelFile(filepath string) (*Model, error) {
	input, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	md := ModelDef{}
	if err := json.Unmarshal(input, &md); err != nil {
		return nil, err
	}

	if err := md.Valid(); err != nil {
		return nil, err
	}

	return md.Model(), nil
}
