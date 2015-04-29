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

type (
	// Definition for a trait
	// { ...
	//		"trait_name": {
	//			"type": "datetime"
	//		}
	// ... }
	TraitDef struct {
		Type string `json:"type"`
	}

	// Definition for a link
	// { ...
	//		"link_name": {
	//			"other": "event",
	//			"type": "mul",
	//		}
	// ... }
	LinkDef struct {
		Kind  string `json:"kind"`
		Other string `json:"other"`
	}

	// Definition for a model
	// {
	//		"kind": "user", "plural": "users", "version": "1.0.0",
	//		"links": { ... },
	//		"traits": { ... }
	// }
	ModelDef struct {
		Kind    string `json:"kind"`
		Plural  string `json:"plural"`
		Version string `json:"version"`

		Traits map[string]TraitDef
		Links  map[string]LinkDef
	}
)

// Turns a trait definition into a metis trait
func (td TraitDef) Process(name string) Trait {
	return Trait{
		Name: name,
		Type: primitiveLiterals[td.Type],
	}
}

// Turns a link definition into a metis link
func (ld LinkDef) Process(name string) Link {
	return Link{
		Name:  name,
		Kind:  multiplicityLiterals[ld.Kind],
		Other: ld.Other,
	}
}

// Turns a model definition into a metis model, including
// transforming each of the Traits and Relationships associated
// with the transformation
func (md ModelDef) Process() Model {
	m := Model{
		Kind:    md.Kind,
		Plural:  md.Plural,
		Version: md.Version,

		Traits: make(map[string]Trait),
		Links:  make(map[string]Link),
	}

	for name, def := range md.Traits {
		m.Traits[name] = def.Process(name)
	}

	for name, def := range md.Links {
		m.Links[name] = def.Process(name)
	}

	return m
}

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

	model := md.Process()
	return model
}
