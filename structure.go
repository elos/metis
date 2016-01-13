package metis

import "strings"

// Traits creates the list of traits required to isomorphically
// represent a relation. In effect, the traits are the structure of
// the relation. You may want to know the traits which represent the
// function when you are trying to reduce the structure of a model
// to a list of name type tuples (that is to say, precisely Traits)
func Traits(r *Relation, s *Schema) (traits []*Trait) {
	traits = make([]*Trait, 1)

	if r.Multiplicity == One {
		traits[0] = &Trait{
			Name: strings.ToLower(r.Name) + "_id",
			Type: ID,
		}
	} else {
		traits[0] = &Trait{
			Name: strings.ToLower(r.Name) + "_ids",
			Type: IDList,
		}
	}

	if s != nil && !s.IsPhysical(r.Codomain) {
		traits = append(traits, &Trait{
			Name: strings.ToLower(r.Name) + "_kind",
			Type: String,
		})
	}

	return
}

// Structure returns a list of the traits which define
// the structure of the model. This is useful for establishing
// the fields or properties needed to represent a model across
// different domains
func Structure(m *Model) (traits []*Trait) {
	traits = make([]*Trait, len(m.Traits))

	i := 0
	for _, t := range m.Traits {
		traits[i] = t
		i++
	}

	for _, r := range m.Relations {
		for _, t := range Traits(r, m.Schema) {
			traits = append(traits, t)
		}
	}

	return
}
