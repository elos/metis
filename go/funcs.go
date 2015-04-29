package ego

import (
	"strings"

	"github.com/elos/gen/metis"
)

func TypeFor(p metis.Primitive) string {
	return PrimitiveTypes[p]
}

func Attr(s string) string {
	return "E" + strings.Title(metis.CamelCase(s))
}
