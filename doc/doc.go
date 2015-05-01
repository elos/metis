package doc

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/elos/ehttp/templates"
	"github.com/elos/gen/metis"
)

var (
	importPath                = filepath.Join(metis.ImportPath, "doc")
	dirPath                   = templates.PackagePath(importPath)
	DocFile    templates.Name = 0
	engine                    = templates.NewEngine(dirPath, &templates.TemplateSet{
		DocFile: []string{filepath.Join("doc.tmpl")},
	}).WithFuncMap(template.FuncMap{
		"export": strings.Title,
		"camel":  metis.CamelCase,
		"isMul":  metis.IsMul,
	})
)

type Doc struct {
	*metis.Model
	Text string
}

func (d *Doc) WriteFile(path string) {
	if err := engine.ParseTemplates(); err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := engine.Execute(&buf, DocFile, d); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(path, buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func MakeDoc(m *metis.Model, textPath string) *Doc {
	bytes, err := ioutil.ReadFile(textPath)
	if err != nil {
		panic(err)
	}
	return &Doc{
		Model: m,
		Text:  string(bytes),
	}
}
