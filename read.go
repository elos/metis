package metis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"

	"github.com/elos/httpserver/templates"
)

func Read(filepath string) Model {
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

var (
	metisImportPath = "github.com/elos/gen/metis"
	metisPath       = templates.PackagePath(metisImportPath)
	templatesDir    = filepath.Join(metisPath, "templates")
)

type GoModel struct {
	Model
}

var goTypeFor = func(t DataType) string {
	return gotypes[t]
}

func MakeGo(m Model) []byte {
	t := template.New("file.tmpl")

	t.Funcs(map[string]interface{}{
		"goTypeFor": goTypeFor,
	})

	pattern := filepath.Join(templatesDir, "go/*.tmpl")
	tmpl := template.Must(t.ParseGlob(pattern))

	var buf bytes.Buffer
	err := tmpl.Lookup("File").Execute(&buf, GoModel{m})
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
