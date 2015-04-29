package ego

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/elos/ehttp/templates"
	"github.com/elos/gen/metis"
)

type GoModel struct {
	metis.Model
}

var (
	templatesDir = filepath.Join(metis.Path, "go", "templates")
	pattern      = filepath.Join(templatesDir, "*.tmpl")

	funcs = map[string]interface{}{
		"typeFor":       TypeFor,
		"camel":         metis.CamelCase,
		"export":        strings.Title,
		"attr":          Attr,
		"appendStrings": metis.AppendStrings,
		"dict":          templates.Dict,
		"isMul":         metis.IsMul,
		"isID":          metis.IsID,
		"firstLetter":   func(s string) string { return metis.Initials(s)[0] },
	}

	t = template.Must(template.New("file.tmpl").Funcs(funcs).ParseGlob(pattern))

	templateName = "File"
)

type GoFile struct {
	contents []byte
}

func (gf GoFile) WriteFile(path string) {
	err := ioutil.WriteFile(path, gf.contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func MakeGo(m metis.Model) metis.Producer {
	var buf bytes.Buffer
	err := t.Lookup(templateName).Execute(&buf, GoModel{m})
	if err != nil {
		log.Fatal(err)
	}
	return GoFile{contents: buf.Bytes()}
}
