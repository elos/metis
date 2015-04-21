package metis

// Turns a trait definition into a metis trait
func (td TraitDef) Process(name string) Trait {
	return Trait{
		Name: name,
		Type: primitiveLiterals[td.Type],
	}
}

// Turns a relationship definition into a metis relationship
func (rd RelationshipDef) Process(name string) Relationship {
	return Relationship{
		Name:  name,
		Kind:  multiplicities[rd.Kind],
		Other: rd.Other,
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

		Traits:        make(map[string]Trait),
		Relationships: make(map[string]Relationship),
	}

	for name, def := range md.Traits {
		m.Traits[name] = def.Process(name)
	}

	for name, def := range md.Relationships {
		m.Relationships[name] = def.Process(name)
	}

	return m
}
