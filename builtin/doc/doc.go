package doc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/elos/ehttp/templates"
	"github.com/elos/metis"
)

func attributesTable(m *metis.Model) string {
	var buf bytes.Buffer
	for _, t := range m.Traits {
		n := strings.Title(metis.CamelCase(t.Name))
		fmt.Fprintf(&buf, "| %s | %s | %s |\n", n, t.Name, metis.InvPrimitiveLiterals[t.Type])
	}

	for _, l := range m.Links {
		n := strings.Title(metis.CamelCase(l.Name))
		json := l.Name + "_id"
		t := "id"

		if metis.IsMul(l) {
			n += "s"
			json += "s"
			t = "[]" + t
		}

		fmt.Fprintf(&buf, "| %s | %s | %s |\n", n, json, t)
	}

	return string(buf.Bytes())
}

var (
	importPath                = filepath.Join(metis.ImportPath, "doc")
	dirPath                   = templates.PackagePath(importPath)
	DocFile    templates.Name = 0
	engine                    = templates.NewEngine(dirPath, &templates.TemplateSet{
		DocFile: []string{filepath.Join("doc.tmpl")},
	}).WithFuncMap(template.FuncMap{
		"export":          strings.Title,
		"camel":           metis.CamelCase,
		"isMul":           metis.IsMul,
		"attributesTable": attributesTable,
	})
)

type Doc struct {
	*metis.Model
	Text string
}

func (d *Doc) WriteFile(path string) {
	if err := engine.Parse(); err != nil {
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
