package metis

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
	"strings"
)

type GoModel struct {
	Model
}

var (
	goTypeFor = func(p Primitive) string {
		return goPrimitiveTypes[p]
	}

	export = func(s string) string {
		return strings.Title(s)
	}

	attr = func(s string) string {
		return "E" + s
	}

	goTemplateFuncs = map[string]interface{}{
		"goTypeFor": goTypeFor,
		"export":    export,
		"attr":      attr,
	}

	goTemplatesPattern = filepath.Join(templatesDir, "go/*.tmpl")

	goFileTemplate = template.Must(template.New("file.tmpl").Funcs(goTemplateFuncs).ParseGlob(goTemplatesPattern))

	goFileTemplateName = "File"
)

func MakeGo(m Model) []byte {
	var buf bytes.Buffer
	err := goFileTemplate.Lookup(goFileTemplateName).Execute(&buf, GoModel{m})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
