package epy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/elos/ehttp/templates"
	"github.com/elos/metis"
)

// Mappings to equivalent python types
var PrimitiveTypes = map[metis.Primitive]string{
	metis.Boolean:      "bool",
	metis.Integer:      "int",
	metis.Float:        "float",
	metis.String:       "str",
	metis.DateTime:     "DateTime",
	metis.BooleanList:  "[]bool",
	metis.IntegerList:  "[]int",
	metis.StringList:   "[]str",
	metis.DateTimeList: "[]Datetime",
	metis.ID:           "str",
	metis.IDList:       "[]str",
	metis.StringIDMap:  "map[str]str",
	metis.IntegerIDMap: "map[int]str",
}

var (
	// templatesDir is the absolute path the epy's templates
	templatesDir = filepath.Join(metis.Path, "builtin", "epy")

	funcs = map[string]interface{}{
		"name": func(s string) string {
			return strings.Title(metis.CamelCase(s))
		},
		"field": func(v interface{}) string {
			switch v.(type) {
			case *metis.Trait:
				return v.(*metis.Trait).Name
			case *metis.Relation:
				r := v.(*metis.Relation)
				if metis.IsMul(r) {
					return r.Name + "_ids"
				} else {
					return r.Name + "_id"
				}
			default:
				panic("field not passed *Trait or *Relation")
			}
		},
		"type": func(v interface{}) string {
			switch v.(type) {
			case *metis.Trait:
				t := v.(*metis.Trait)
				return PrimitiveTypes[t.Type]
			case *metis.Relation:
				r := v.(*metis.Relation)
				if metis.IsMul(r) {
					return PrimitiveTypes[metis.String]
				} else {
					return PrimitiveTypes[metis.StringList]
				}
			default:
				panic("type not passed *Trait or *Relation")
			}
		},
		"unpack": func(t *metis.Trait, jsonVariable string) string {
			switch t.Type {
			case metis.DateTime:
				return fmt.Sprintf("dateutil.parser.parse(%s[\"%s\"])", jsonVariable, t.Name)
			case metis.DateTimeList:
				return fmt.Sprintf("map(dateutil.parser.parse, %s[\"%s\"])", jsonVariable, t.Name)
			}
			return fmt.Sprintf("%s[\"%s\"]", jsonVariable, t.Name)
		},
		"isMul": metis.IsMul,
		"kindFor": func(s *metis.Schema, r *metis.Relation) string {
			return s.Spaces[r.Codomain]
		},
	}

	File templates.Name = 1

	engine = func() *templates.Engine {
		e := templates.NewEngine(templatesDir, &templates.TemplateSet{
			File: []string{"file.tmpl"},
		}).WithFuncMap(funcs)

		if err := e.Parse(); err != nil {
			log.Fatal(err)
		}

		return e
	}()
)

func Generate(s *metis.Schema, into string) error {
	if err := s.Valid(); err != nil {
		return fmt.Errorf("epy.Generate Error: schema invalid: %s", err)
	}

	var out bytes.Buffer
	if err := engine.Execute(&out, File, s); err != nil {
		return fmt.Errorf("epy.Generate Error: templating: %s", err)
	}

	if err := ioutil.WriteFile(into, out.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
