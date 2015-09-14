package metis

import (
	"errors"
	"fmt"
)

// These definition structures are the JSON blueprints for the logical components
// of metis. It follows, then, that you can have trait, link, and model defintions.
// A definition is not necessarily a valid metis construct, and therefore is treated
// as a intermediate stage. We define a function on each definition to check its
// validity and a transformation to construct a well-formed logical component.

type (
	// TraitDef is the definition for a trait.
	//		"traits": [
	//			{"name": "<Name>", "type": "<Type">},
	//		]
	TraitDef struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// RelationDef is the definition for a relation.
	//		"relations": [
	//			{"name": "<Name>", "multiplicity": "<Multiplicity">,
	// 			 "codomain": "<Codomain>", "inverse": "<Inverse>"},
	//		]
	RelationDef struct {
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
	//			"relations": [ ... ],
	//		}
	ModelDef struct {
		Kind      string   `json:"kind"`
		Space     string   `json:"space"`
		Domains   []string `json:"domains"`
		Traits    []*TraitDef
		Relations []*RelationDef
	}
)

// TraitDef: Valid() error, Trait *Trait() {{{

// Valid returns an error if the trait definition is invalid, otherwise nil.
// A trait can be invalid for 2 reasons
//	1. It lacks a name
//  2. It lacks a valid primitive type
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

// Trait constructs and returns a metis.Trait built from the defintion
func (td *TraitDef) Trait() *Trait {
	return &Trait{
		Name: td.Name,
		Type: primitiveLiterals[td.Type],
	}
}

// }}}

// RelationDef: Valid() error, Link() *Link  {{{

// Valid returns an error if a Link definition is invalid, otherwise nil.
// A Link can be invalid for 4 reasons
//  1. It lacks a name
//  2. It lacks a valid multiplicity
//  3. It lacks a codomain
//  4. It lakcs a singular form, despite having a multiplicity of "mul"
func (ld *RelationDef) Valid() error {
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

// Relation constructs and returns a metis.Link built from the definition
func (ld *RelationDef) Relation() *Relation {
	return &Relation{
		Name:         ld.Name,
		Multiplicity: multiplicityLiterals[ld.Multiplicity],
		Singular:     ld.Singular,
		Codomain:     ld.Codomain,
		Inverse:      ld.Inverse,
	}
}

// }}}

// ModelDef: Valid() error, Model() *Model {{{

// Valid returns an error if the Model definition is invalid, otherwise nil.
// Note: Valid checks the validity of the traits and links of the model
// A model can be invalid for 5 reasons
//  1. It lacks a kind
//  2. It lacks a space
//  3. It lacks at least one domain
//  4. It has a trait or link name clash
//  5. It has a trait or link error
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

	for _, l := range md.Relations {
		if _, seen := seenNames[l.Name]; seen {
			return fmt.Errorf("model %s name clash %s", md.Kind, l.Name)
		}

		if err := l.Valid(); err != nil {
			return fmt.Errorf("model %s has link error: %s", md.Kind, err.Error())
		}
	}

	return nil
}

// Model constructs and returns a metis.Model built from the definition
// Note: This procedure will also construct the traits and links of the model
func (md *ModelDef) Model() *Model {
	m := &Model{
		Kind:      md.Kind,
		Space:     md.Space,
		Domains:   md.Domains,
		Traits:    make(map[string]*Trait),
		Relations: make(map[string]*Relation),
	}

	for _, td := range md.Traits {
		m.Traits[td.Name] = td.Trait()
	}

	for _, ld := range md.Relations {
		m.Relations[ld.Name] = ld.Relation()
	}

	return m

}

// }}}
