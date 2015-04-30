package metis

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/elos/ehttp/templates"
)

var (
	ImportPath = "github.com/elos/gen/metis"
	Path       = templates.PackagePath(ImportPath)
)

func ParseModelFile(filepath string) Model {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	md := ModelDef{}
	err = json.Unmarshal(input, &md)
	if err != nil {
		log.Fatal(err)
	}

	model := md.Model()
	return *model
}
