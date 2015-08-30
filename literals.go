package metis

// Metis string literals for primitive definitions
var primitiveLiterals = map[string]Primitive{
	"boolean":        Boolean,
	"integer":        Integer,
	"float":          Float,
	"string":         String,
	"datetime":       DateTime,
	"[]boolean":      BooleanList,
	"[]integer":      IntegerList,
	"[]string":       StringList,
	"[]datetime":     DateTimeList,
	"id":             ID,
	"[]id":           IDList,
	"map[string]id":  StringIDMap,
	"map[integer]id": IntegerIDMap,
}

var InvPrimitiveLiterals = map[Primitive]string{
	Boolean:      "boolean",
	Integer:      "integer",
	Float:        "float",
	String:       "string",
	DateTime:     "datetime",
	BooleanList:  "[]boolean",
	IntegerList:  "[]integer",
	StringList:   "[]string",
	DateTimeList: "[]datetime",
	ID:           "id",
	IDList:       "[]id",
	StringIDMap:  "map[string]id",
	IntegerIDMap: "mao[integer]id",
}

var multiplicityLiterals = map[string]Multiplicity{
	"one": One,
	"mul": Mul,
}
