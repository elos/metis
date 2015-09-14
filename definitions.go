package metis

import (
	"errors"
	"fmt"
)

// Definition are the JSON blueprints for the logical components of metis.
// It follows, then, that you can have a Trait definition, a link definition
// and a model definition. These definitions have rules of validity applied to
// them and each has a transformation,
// Valid: d ∈ Definitions → l ∈ Logical Components

type (
	// TraitDef is the definition for a trait
	//		"traits": [
	//			{"name": "<Name>", "type": "<Type">},
	//		]
	TraitDef struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// LinkDef is the definition for a link
	//		"links": [
	//			{"name": "<Name>", "multiplicity": "<Multiplicity">,
	// 			 "codomain": "<Codomain>", "inverse": "<Inverse>"},
	//		]
	LinkDef struct {
		Name         string `json:"name"`
		Multiplicity string `json:"multiplicity"`
		Singular     string `json:"singular"`
		Codomain     string `json:"codomain"`
		Inverse      string `json:"inverse"`
	}

	// ModelDef is the definition for a model
	//		{
	//			"kind": "<Kind>",
	//			"space": "<Space>",
	//			"domains": "<Domains>",
	//			"traits": [ ... ],
	//			"links": [ ... ],
	//		}
	ModelDef struct {
		Kind    string   `json:"kind"`
		Space   string   `json:"space"`
		Domains []string `json:"domains"`
		Traits  []*TraitDef
		Links   []*LinkDef
	}

	// SchemaDef is the definition for a schema
	//		{
	//			"version": "<Version>",
	//			"spaces": ["<space1>", ... ,  "<spaceN>"],
	//			"domains": ["<domain1>", ... , "<domainN>"],
	//		}
	SchemaDef struct {
		Version string   `json:"version"`
		Spaces  []string `json:"spaces"`
		Domains []string `json:"domains"`
	}
)

// TraitDef: Valid() error, Trait *Trait() {{{

// Valid returns an error if the trait definition is invalid
// or nil otherwise
// A trait can be invalid for 2 reasons
//	1. It does not have a name
//  2. It does not have a valid primitive type
func (td *TraitDef) Valid() error {
	if td.Name == "" {
		return errors.New("trait definition must have a name")
	}

	_, validType := primitiveLiterals[td.Type]
	if !validType {
		return fmt.Errorf("trait definition must have valid type, type %s is invalid", td.Type)
	}

	return nil
}

// Trait returns a metis.Trait built from the
// Trait definition, TraitDef
func (td *TraitDef) Trait() *Trait {
	return &Trait{
		Name: td.Name,
		Type: primitiveLiterals[td.Type],
	}
}

// }}}

// LinkDef: Valid() error, Link() *Link  {{{

// Valid returns an error if a Link definition is invalid,
// or nil otherwise
// A Link can be invalid for 4 reasons
// 1. It does not have a valid name
// 2. It has an invalid multiplicity
// 3. It lacks a codomain
// 4. It has a multiplicity of "mul," but no singular form specified
func (ld *LinkDef) Valid() error {
	if ld.Name == "" {
		return errors.New("link definition must have a name")
	}

	_, validMultiplicity := multiplicityLiterals[ld.Multiplicity]
	if !validMultiplicity {
		return fmt.Errorf("link definition must have valid multiplicity, multiplicity %s is invalid", ld.Multiplicity)
	}

	if ld.Codomain == "" {
		return errors.New("link definition must have codomain")
	}

	if multiplicityLiterals[ld.Multiplicity] == Mul {
		if ld.Singular == "" {
			return errors.New("mul link defintion must specify singular form")
		}
	}

	return nil
}

// Link returns a metis.Link that is based on the
// definition of this LinkDef
func (ld *LinkDef) Link() *Link {
	return &Link{
		Name:         ld.Name,
		Multiplicity: multiplicityLiterals[ld.Multiplicity],
		Singular:     ld.Singular,
		Codomain:     ld.Codomain,
		Inverse:      ld.Inverse,
	}
}

// }}}

// ModelDef: Valid() error, Model() *Model {{{

// Valid returns an error if the Model definition is invalid
// or nil otherwise. Valid recursively checks the validity
// of a model's traits and links
// A model can be invalid for 5 reasons
// 1. It does not have a kind
// 2. It does not have a space
// 3. It does not have at least one domain defined
// 4. It has a trait or link name clash
// 5. It has a trait or link error
func (md *ModelDef) Valid() error {
	if md.Kind == "" {
		return errors.New("model definition must have a kind")
	}

	if md.Space == "" {
		return fmt.Errorf("model definition for %s must have a space", md.Kind)
	}

	for i := range md.Domains {
		if md.Domains[i] == "" {
			return fmt.Errorf("model definition for %s has no domain", md.Kind)
		}
	}

	seenNames := make(map[string]bool)

	for _, t := range md.Traits {
		if _, seen := seenNames[t.Name]; seen {
			return fmt.Errorf("model %s has name clash %s", md.Kind, t.Name)
		}

		if err := t.Valid(); err != nil {
			return fmt.Errorf("model %s has trait error: %s", md.Kind, err.Error())
		}
	}

	for _, l := range md.Links {
		if _, seen := seenNames[l.Name]; seen {
			return fmt.Errorf("model %s name clash %s", md.Kind, l.Name)
		}

		if err := l.Valid(); err != nil {
			return fmt.Errorf("model %s has link error: %s", md.Kind, err.Error())
		}
	}

	return nil
}

// Model creates a metis.Model using the definition
// defined in the ModelDef
func (md *ModelDef) Model() *Model {
	m := &Model{
		Kind:    md.Kind,
		Space:   md.Space,
		Domains: md.Domains,
		Traits:  make(map[string]*Trait),
		Links:   make(map[string]*Link),
	}

	for _, td := range md.Traits {
		m.Traits[td.Name] = td.Trait()
	}

	for _, ld := range md.Links {
		m.Links[ld.Name] = ld.Link()
	}

	return m

}

// }}}
