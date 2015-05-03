package metis

// Metis string literals for primitive definitions
var primitiveLiterals = map[string]Primitive{
	"boolean":       Boolean,
	"integer":       Integer,
	"string":        String,
	"datetime":      DateTime,
	"[]boolean":     BooleanList,
	"[]integer":     IntegerList,
	"[]string":      StringList,
	"[]datetime":    DateTimeList,
	"id":            ID,
	"[]id":          IDList,
	"map[string]id": StringIDMap,
}

var InvPrimitiveLiterals = map[Primitive]string{
	Boolean:      "boolean",
	Integer:      "integer",
	String:       "string",
	DateTime:     "datetime",
	BooleanList:  "[]boolean",
	IntegerList:  "[]integer",
	StringList:   "[]string",
	DateTimeList: "[]datetime",
	ID:           "id",
	IDList:       "[]id",
	StringIDMap:  "map[string]id",
}

var multiplicityLiterals = map[string]Multiplicity{
	"one": One,
	"mul": Mul,
}
