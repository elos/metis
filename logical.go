package metis

type (
	Trait struct {
		Name string
		Type Primitive
	}

	Link struct {
		Name         string
		Multiplicity Multiplicity
		Singular     string
		Codomain     string
		Inverse      string
	}

	Model struct {
		Kind    string
		Space   string
		Domains []string
		Traits  map[string]*Trait
		Links   map[string]*Link
		*Schema
	}

	Schema struct {
		Version string
		Spaces  map[string]string
		Domains map[string]bool
		Models  map[string]*Model
	}
)

type (
	Primitive    int
	Multiplicity int
	Producer     interface {
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
	IDList

	StringIDMap
)

// Primitive Multiplicities
const (
	Mul Multiplicity = iota
	One
)
