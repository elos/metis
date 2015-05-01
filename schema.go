package metis

import "fmt"

func BuildSchema(models ...*Model) *Schema {
	s := new(Schema)
	s.Spaces = make(map[string]bool)
	s.Domains = make(map[string]bool)
	s.Models = make(map[string]*Model)

	for _, m := range models {
		s.Models[m.Space] = m

		for i := range m.Domains {
			s.Domains[m.Domains[i]] = true
		}

		s.Spaces[m.Space] = true
	}

	return s
}

func (s *Schema) Valid() error {
	seenSpaces := make(map[string]bool)
	seenKinds := make(map[string]bool)

	for _, m := range s.Models {
		if _, seen := seenKinds[m.Kind]; seen {
			return fmt.Errorf("Duplicate kind %s", m.Kind)
		}

		if _, seen := seenSpaces[m.Space]; seen {
			return fmt.Errorf("Duplicate space %s", m.Space)
		}

		seenSpaces[m.Space] = true

		for _, l := range m.Links {
			if _, codomainDefined := s.Domains[l.Codomain]; !codomainDefined {
				return fmt.Errorf("Model %s has codomain that is undefined: %s", m.Kind, l.Codomain)
			}

			if l.Inverse == "" {
				continue // don't need to check inverses
			}

			// for a codomain to have an inverse, it must be a concrete space
			if _, concrete := s.Spaces[l.Codomain]; !concrete {
				return fmt.Errorf("Model %s has codomain that is not concrete but has inverse", m.Kind)
			}

			other := s.Models[l.Codomain]

			// that concrete definition must have matching inverse
			if _, ok := other.Links[l.Inverse]; !ok {
				return fmt.Errorf("Model %s, link %s has invalid inverse definition: %s", m.Kind, l.Name, l.Inverse)
			}
		}
	}

	return nil
}
