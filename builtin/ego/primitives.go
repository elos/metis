package ego

import (
	m "github.com/elos/metis"
)

// Mappings to equivalent go types
var PrimitiveTypes = map[m.Primitive]string{
	m.Boolean:      "bool",
	m.Integer:      "int",
	m.Float:        "float64",
	m.String:       "string",
	m.DateTime:     "time.Time",
	m.BooleanList:  "[]bool",
	m.IntegerList:  "[]int",
	m.StringList:   "[]string",
	m.DateTimeList: "[]time.Time",
	m.ID:           "string",
	m.IDList:       "[]string",
	m.StringIDMap:  "map[string]string",
	m.IntegerIDMap: "map[int]string",
	m.JSON:         "map[string]interface{}",
}
