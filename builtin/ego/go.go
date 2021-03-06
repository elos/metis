package ego

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/elos/ehttp/templates"
	"github.com/elos/metis"
)

type GoModel struct {
	*metis.Model
}

func name(s string) string {
	return strings.Title(metis.CamelCase(s))
}

func initial(s string) string {
	return metis.Initials(s)[0]
}

func traitFieldName(t *metis.Trait) string {
	f := strings.Title(metis.CamelCase(t.Name))
	return f
}

func relationFieldName(r *metis.Relation) string {
	f := strings.Title(metis.CamelCase(r.Name)) + "Id"
	if metis.IsMul(r) {
		f += "s"
	}
	return f
}

func traitFieldTags(t *metis.Trait) string {
	tags := fmt.Sprintf("`json:\"%s\"", t.Name)
	if metis.IsID(t) {
		tags += " bson:\"_id,omitempty\"`"
	} else {
		tags += fmt.Sprintf(" bson:\"%s\"`", t.Name)
	}
	return tags
}

func relationFieldTags(r *metis.Relation) string {
	tag := r.Name + "_id"

	if metis.IsMul(r) {
		tag += "s"
	}

	tags := fmt.Sprintf("`json:\"%s\" bson:\"%s\"`", tag, tag)
	return tags
}

func relationFieldType(l *metis.Relation) string {
	tipe := "string"
	if metis.IsMul(l) {
		tipe = "[]" + tipe
	}
	return tipe
}

func virtualRelationExtraFields(l *metis.Relation) string {
	field := strings.Title(metis.CamelCase(l.Name)) + "Kind"
	tag := l.Name + "_kind"
	tags := fmt.Sprintf("`json:\"%s\" bson:\"%s\"`", tag, tag)
	return fmt.Sprintf("%s string %s\n", field, tags)
}

func typeDefinition(m *metis.Model) string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "type %s struct{\n", name(m.Kind))

	keys := make([]string, 0)
	for k := range m.Traits {
		keys = append(keys, k)
	}
	for k := range m.Relations {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		t, ok := m.Traits[keys[i]]
		if ok {
			fmt.Fprintf(&buf, "%s %s %s\n", traitFieldName(t), TypeFor(t.Type), traitFieldTags(t))
			continue
		}

		l := m.Relations[keys[i]]

		fmt.Fprintf(&buf, "%s %s %s\n", relationFieldName(l), relationFieldType(l), relationFieldTags(l))

		if !m.Schema.IsPhysical(l.Codomain) {
			fmt.Fprintf(&buf, virtualRelationExtraFields(l))
		}
	}

	fmt.Fprint(&buf, "}\n")

	return buf.String()
}

func traitAccessors(m *metis.Model) string {
	// none!
	return ""
}

type accessVar struct {
	Model *metis.Model
	This  struct {
		Relation *metis.Relation
	}
	Other struct {
		Type, Kind string
		IsPhysical bool
	}
}

func oneRelationAccessorsVar(m *metis.Model, l *metis.Relation) *accessVar {
	var oKind, oType string
	s := m.Schema
	if s.IsPhysical(l.Codomain) {
		oKind = s.Spaces[l.Codomain]
		oType = "*" + name(oKind)
	} else {
		oKind = l.Codomain
		oType = name(l.Codomain)
	}

	return &accessVar{
		m,
		struct{ Relation *metis.Relation }{Relation: l},
		struct {
			Type       string
			Kind       string
			IsPhysical bool
		}{oType, oKind, s.IsPhysical(l.Codomain)},
	}
}

func mulRelationAccessorsVar(m *metis.Model, l *metis.Relation) *accessVar {
	var oKind, oType string
	s := m.Schema
	if s.IsPhysical(l.Codomain) {
		oKind = s.Spaces[l.Codomain]
		oType = "*" + name(oKind)
	} else {
		oKind = l.Codomain
		oType = name(l.Codomain)
	}
	return &accessVar{
		m,
		struct{ Relation *metis.Relation }{Relation: l},
		struct {
			Type       string
			Kind       string
			IsPhysical bool
		}{oType, oKind, s.IsPhysical(l.Codomain)},
	}
}

var (
	templatesDir = filepath.Join(metis.Path, "builtin", "ego", "templates")
	pattern      = filepath.Join(templatesDir, "*.tmpl")

	funcs = map[string]interface{}{
		"typeFor": TypeFor,
		"camel":   metis.CamelCase,
		"export":  strings.Title,
		"attr":    Attr,
		"dict":    templates.Dict,
		"isMul":   metis.IsMul,
		"isID":    metis.IsID,
		"isVirtual": func(m *metis.Model, l *metis.Relation) bool {
			return !m.Schema.IsPhysical(l.Codomain)
		},
		"firstLetter":                func(s string) string { return metis.Initials(s)[0] },
		"typeDefinition":             typeDefinition,
		"traitAccessors":             traitAccessors,
		"oneRelationAccessorsVar":    oneRelationAccessorsVar,
		"mulRelationAccessorsVar":    mulRelationAccessorsVar,
		"initial":                    func(s string) string { return metis.Initials(s)[0] },
		"sig":                        Signature,
		"name":                       name,
		"relationFieldName":          relationFieldName,
		"relationFieldType":          relationFieldType,
		"relationFieldTags":          relationFieldTags,
		"virtualRelationExtraFields": virtualRelationExtraFields,
		"traitFieldName":             traitFieldName,
		"traitFieldTags":             traitFieldTags,
		"this": func(m *metis.Model) string {
			return metis.CamelCase(m.Kind)
		},
		"structure": metis.Structure,
	}

	File         templates.Name = 0
	Kinds        templates.Name = 1
	Constructors templates.Name = 2
	Dynamic      templates.Name = 3
	DBs          templates.Name = 4

	engine = func() *templates.Engine {
		e := templates.NewEngine(templatesDir, &templates.TemplateSet{
			File:         []string{"file.tmpl", "one_accessors.tmpl", "mul_accessors.tmpl", "bson.tmpl"},
			Kinds:        []string{"kinds.tmpl"},
			Constructors: []string{"constructors.tmpl"},
			Dynamic:      []string{"dynamic.tmpl"},
			DBs:          []string{"dbs.tmpl"},
		}).WithFuncMap(funcs)

		err := e.Parse()
		if err != nil {
			log.Fatal(err)
		}
		return e
	}()

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
	err = exec.Command("gofmt", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func MakeGo(m *metis.Model) metis.Producer {
	var buf bytes.Buffer
	err := engine.Execute(&buf, File, GoModel{m})
	if err != nil {
		log.Fatal(err)
	}
	return GoFile{contents: buf.Bytes()}
}

func WriteKindsFile(s *metis.Schema, path string) {
	var buf bytes.Buffer
	err := engine.Execute(&buf, Kinds, s)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteConstructorsFile(s *metis.Schema, path string) {
	var buf bytes.Buffer
	err := engine.Execute(&buf, Constructors, s)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteDBsFile(s *metis.Schema, path string) {
	var buf bytes.Buffer
	err := engine.Execute(&buf, DBs, s)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteDynamicFile(s *metis.Schema, path string) {
	var buf bytes.Buffer
	err := engine.Execute(&buf, Dynamic, s)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("goimports", "-w=true", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}
