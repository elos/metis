package ego

import (
	"fmt"
	"strings"

	"github.com/elos/gen/metis"
)

func TypeFor(p metis.Primitive) string {
	return PrimitiveTypes[p]
}

func Attr(s string) string {
	return "E" + strings.Title(metis.CamelCase(s))
}

func Signature(m *metis.Model) string {
	return fmt.Sprintf("func (%s *%s)", metis.CamelCase(m.Kind), name(m.Kind))
}

func InterfaceFor(schema *metis.Schema, space string) string {
	if schema.IsPhysical(space) {
		return name("*" + schema.Spaces[space])
	}

	/* don't try to be magical for now
	if string(space[len(space)-1]) == "s" {
		space = string(space[:len(space)-1])
	}
	*/
	return name(space)
}
