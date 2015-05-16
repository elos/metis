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
	// Primitive is a base type of a trait
	Primitive int
	// Multiplicity reps the size of a link
	Multiplicity int
	Producer     interface {
		WriteFile(string)
	}
)

// Primitive Data Types we list respective
// go types in comments for clarity
// metis _is_ in go after all :)
const (
	Boolean  Primitive = iota // bool
	Integer                   // int
	String                    // string
	DateTime                  // time.Time

	BooleanList  // []bool
	IntegerList  // []int
	StringList   // []string
	DateTimeList // []time.Time

	ID     // string
	IDList //[]string

	StringIDMap // map[string]string
)

// Primitive Multiplicities
const (
	Mul Multiplicity = iota // has_many
	One                     // belongs_to (because the id is on the struct)
)
