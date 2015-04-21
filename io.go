package metis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/elos/httpserver/templates"
)

var (
	metisImportPath = "github.com/elos/gen/metis"
	metisPath       = templates.PackagePath(metisImportPath)
	templatesDir    = filepath.Join(metisPath, "templates")
)

func Parse(filepath string) Model {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	md := ModelDef{}
	err = json.Unmarshal(input, &md)
	if err != nil {
		log.Fatal(err)
	}

	model := md.Process()
	return model
}
