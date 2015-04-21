package metis

type Primitive int

// Metis Primitive Data Types
const (
	Boolean Primitive = iota
	Integer
	String
	DateTime
	BooleanList
	IntegerList
	StringList
	DateTimeList
)

// Metis string literals for primitive definitions
var primitiveLiterals = map[string]Primitive{
	"boolean":    Boolean,
	"integer":    Integer,
	"string":     String,
	"datetime":   DateTime,
	"[]boolean":  BooleanList,
	"[]integer":  IntegerList,
	"[]string":   StringList,
	"[]datetime": DateTimeList,
}

// Mappings to equivalent go types
var goPrimitiveTypes = map[Primitive]string{
	Boolean:      "bool",
	Integer:      "int",
	String:       "string",
	DateTime:     "*time.Time",
	BooleanList:  "[]bool",
	IntegerList:  "[]int",
	StringList:   "[]string",
	DateTimeList: "[]*time.Time",
}
