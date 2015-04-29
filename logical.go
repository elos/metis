package metis

type (
	Primitive int

	Multiplicity int

	Trait struct {
		Name string
		Type Primitive
	}

	Link struct {
		Name, Other string
		Kind        Multiplicity
	}

	Model struct {
		Kind, Plural, Version string

		Traits map[string]Trait
		Links  map[string]Link
	}

	Producer interface {
		WriteFile(string)
	}
)

// Primitive Data Types
const (
	Boolean Primitive = iota
	Integer
	String
	DateTime

	BooleanList
	IntegerList
	StringList
	DateTimeList

	ID
)

// Primitive Multiplicities
const (
	Mul Multiplicity = iota
	One
)
