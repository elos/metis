package metis

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
	"id":         ID,
}

var multiplicityLiterals = map[string]Multiplicity{
	"one": One,
	"mul": Mul,
}
