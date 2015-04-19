package metis

type DataType int

const (
	Boolean DataType = iota
	Integer
	String
	DateTime
	BooleanList
	IntegerList
	StringList
	DateTimeList
)

var datatypes = map[string]DataType{
	"boolean":    Boolean,
	"integer":    Integer,
	"string":     String,
	"datetime":   DateTime,
	"[]boolean":  BooleanList,
	"[]integer":  IntegerList,
	"[]string":   StringList,
	"[]datetime": DateTimeList,
}

var gotypes = map[DataType]string{
	Boolean:      "bool",
	Integer:      "int",
	String:       "string",
	DateTime:     "*time.Time",
	BooleanList:  "[]bool",
	IntegerList:  "[]int",
	StringList:   "[]string",
	DateTimeList: "[]*time.Time",
}

var relkinds = map[string]RelKind{
	"one": One,
	"mul": Mul,
}

type RelKind int

const (
	Mul RelKind = iota
	One
)

type TraitDef struct {
	Type string `json:"type"`
}

func (t TraitDef) Process(name string) Trait {
	return Trait{
		Name: name,
		Type: datatypes[t.Type],
	}
}

type RelationshipDef struct {
	Kind  string `json:"kind"`
	Other string `json:"other"`
}

func (r RelationshipDef) Process(name string) Relationship {
	return Relationship{
		Name:  name,
		Kind:  relkinds[r.Kind],
		Other: r.Other,
	}
}

type ModelDef struct {
	Kind    string  `json:"kind"`
	Plural  string  `json:"plural"`
	Version float64 `json:"version"`

	Traits        map[string]TraitDef
	Relationships map[string]RelationshipDef
}

func (md ModelDef) Process() Model {
	m := Model{
		Kind:          md.Kind,
		Plural:        md.Plural,
		Version:       md.Version,
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

type Trait struct {
	Name string
	Type DataType
}

type Relationship struct {
	Name  string
	Kind  RelKind
	Other string
}

type Model struct {
	Kind    string
	Plural  string
	Version float64

	Traits        map[string]Trait
	Relationships map[string]Relationship
}
