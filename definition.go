package metis

import (
	"errors"
	"fmt"
)

type (
	// Definition for a trait
	//		"traits": [
	//			{"name": "<Name>", "type": "<Type">},
	//		]
	TraitDef struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// Definition for a link
	//		"links": [
	//			{"name": "<Name>", "multiplicity": "<Multiplicity">,
	// 			 "codomain": "<Codomain>", "inverse": "<Inverse>"},
	//		]
	LinkDef struct {
		Name         string `json:"name"`
		Multiplicity string `json:"multiplicity"`
		Codomain     string `json:"codomain"`
		Inverse      string `json:"inverse"`
	}

	// Definition for a model
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

	// Definition for a schema
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

func (td *TraitDef) Valid() error {
	if td.Name == "" {
		return errors.New("Trait definition must have a name")
	}

	_, validType := primitiveLiterals[td.Type]
	if !validType {
		return errors.New(fmt.Sprintf("Trait definition must have valid type, type %s is invalid", td.Type))
	}

	return nil
}

func (td *TraitDef) Trait() *Trait {
	return &Trait{
		Name: td.Name,
		Type: primitiveLiterals[td.Type],
	}
}

// }}}

// LinkDef: Valid() error, Link() *Link  {{{

func (ld *LinkDef) Valid() error {
	if ld.Name == "" {
		return errors.New("Link definition must have a name")
	}

	_, validMultiplicity := multiplicityLiterals[ld.Multiplicity]
	if !validMultiplicity {
		return errors.New(fmt.Sprintf("Link definition must have valid multiplicity, multiplicity %s is invalid", ld.Multiplicity))
	}

	if ld.Codomain == "" {
		return errors.New("Link definition must ahve codomain")
	}

	return nil
}

func (ld *LinkDef) Link() *Link {
	return &Link{
		Name:         ld.Name,
		Multiplicity: multiplicityLiterals[ld.Multiplicity],
		Codomain:     ld.Codomain,
		Inverse:      ld.Inverse,
	}
}

// }}}

// ModelDef: Valid() error, Model() *Model {{{

func (md *ModelDef) Valid() error {
	if md.Kind == "" {
		return errors.New("Model definition must have a kind")
	}

	if md.Space == "" {
		return errors.New("Model definition must have a space")
	}

	for i := range md.Domains {
		if md.Domains[i] == "" {
			return errors.New("Model definition domain \"\" is invalid")
		}
	}

	seenNames := make(map[string]bool)

	for _, t := range md.Traits {
		if _, seen := seenNames[t.Name]; seen {
			return fmt.Errorf("Model name clash %s", t.Name)
		}

		if err := t.Valid(); err != nil {
			return err
		}
	}

	for _, l := range md.Links {
		if _, seen := seenNames[l.Name]; seen {
			return fmt.Errorf("Model name clash %s", l.Name)
		}

		if err := l.Valid(); err != nil {
			return err
		}
	}

	return nil
}

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
