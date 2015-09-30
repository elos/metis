package metis

// Here we define types for the metis logical core.

type (
	// A Trait has a name, which you use to refer to the attribute and a
	// metis primitive type (e.g. string, int, boolean)
	Trait struct {
		Name string
		Type Primitive
	}

	// A Relation has a name, a multiplicity (one or mul), a singular form
	// (iff multiplicity == mul), a codomain (the space of models which can be
	// assigned to this relation), and an inverse (the name of the corresponding
	// model's relation to this model).
	Relation struct {
		Name         string
		Multiplicity Multiplicity
		Singular     string
		Codomain     string
		Inverse      string
	}

	// A Model has a kind (it's name), a space (it's plural - but also
	// its physical domain), domains (the physcial and abstract spaces
	// this model implements), and a list of traits and a list of relations.
	Model struct {
		Kind      string
		Space     string
		Domains   []string
		Traits    map[string]*Trait
		Relations map[string]*Relation
		*Schema
	}

	// A Schema has a list of models, a list of spaces and a list of domains.
	// Plus a version. A virtual domain can be thought of as an interface and
	// is any domain referenced in a model that has no physical space counterpart.
	Schema struct {
		Version string
		Spaces  map[string]string
		Domains map[string]bool
		Models  map[string]*Model
	}
)

// Primitive is the pure base type of the value of trait
type Primitive int

// Primitive trait data types.
// We list respective go types in comments for clarity
// metis _is_ in implemented in Go after all :)
const (
	Boolean  Primitive = iota // bool
	Integer                   // int
	Float                     // float
	String                    // string
	DateTime                  // time.Time

	BooleanList  // []bool
	IntegerList  // []int
	StringList   // []string
	DateTimeList // []time.Time

	ID     // string
	IDList //[]string

	StringIDMap  // map[string]id
	IntegerIDMap // map[integer]id
)

// Multiplicity represents the size of a Relation
type Multiplicity int

const (
	Mul Multiplicity = iota // has_many
	One                     // belongs_to (because the id is on the struct)
)
