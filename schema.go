package metis

import "fmt"

// Schema doesn't have any IO parsing, since a schema is implicitly defined by the models
// which compose it. Here's to decentralization! Because of this, the composition of a schema
// and it's verification are two separaten steps. Any set of models can be composed into a schema
// but that does not necessarily make them a valid schema. To compose models, use BuildSchema, to
// verify the schema use Valid()

func (s *Schema) IsPhysical(space string) bool {
	_, ok := s.Spaces[space]
	return ok
}

// BuildSchema creates a schema logical structures from a list
// of models. Note this function _can not_ fail but does not verify
// the formal correctness of the schema. You must check after:
//		s := BuildSchema(models...)
//      if !s.Valid() {
//			log.Fatal("Our schema is invalid")
//		}
func BuildSchema(models ...*Model) *Schema {
	// initialization & allocation
	s := new(Schema)
	s.Spaces = make(map[string]string)
	s.Domains = make(map[string]bool)
	s.Models = make(map[string]*Model)

	for _, m := range models {
		s.Models[m.Space] = m

		for i := range m.Domains {
			// We only care to tally up the valid domains, several
			// models may implement the same domain
			s.Domains[m.Domains[i]] = true
		}

		// Each space is only associated with one model, and therefore one "kind"
		s.Spaces[m.Space] = m.Kind

		// Now we can establish this model as a member of our new schema
		m.Schema = s
	}

	return s
}

// Valid verifies whether the individual models together form a logically valid
// schema. Think about this like the verification of the relationships between models.
func (s *Schema) Valid() error {
	// We need to verify uniqueness of both kinds and spaces
	seenKinds := make(map[string]bool)
	seenSpaces := make(map[string]bool)

	for _, m := range s.Models {
		if _, seen := seenKinds[m.Kind]; seen {
			return fmt.Errorf("duplicate kind %s", m.Kind)
		}

		seenKinds[m.Kind] = true

		if _, seen := seenSpaces[m.Space]; seen {
			return fmt.Errorf("duplicate space %s", m.Space)
		}

		seenSpaces[m.Space] = true

		for _, l := range m.Links {
			// there must be a codomain defined
			if _, codomainDefined := s.Domains[l.Codomain]; !codomainDefined {
				return fmt.Errorf("model %s has codomain that is undefined: %s", m.Kind, l.Codomain)
			}

			if l.Inverse == "" {
				continue // don't need to check inverses
			}

			// for a codomain to have an inverse, it must be a concrete space
			if _, concrete := s.Spaces[l.Codomain]; !concrete {
				return fmt.Errorf("model %s has codomain that is not concrete but has inverse", m.Kind)
			}

			other := s.Models[l.Codomain]

			// that concrete definition must have matching inverse
			if _, ok := other.Links[l.Inverse]; !ok {
				return fmt.Errorf("model %s, link %s has invalid inverse definition: %s", m.Kind, l.Name, l.Inverse)
			}
		}
	}

	// Formally, this schema is valid
	return nil
}
