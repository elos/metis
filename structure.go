package metis

type (
	Trait struct {
		Name string
		Type Primitive
	}

	Relationship struct {
		Name, Other string
		Kind        Multiplicity
	}

	Model struct {
		Kind, Plural, Version string

		Traits        map[string]Trait
		Relationships map[string]Relationship
	}
)
