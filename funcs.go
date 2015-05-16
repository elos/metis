package metis

import "strings"

//  SplitSnakeCase splits a string based on the "_" character
//  e.g., this_is_snake => []string{"this", "is", "snake"}
func SplitSnakeCase(s string) []string {
	return strings.Split(s, "_")
}

// CamelCase turns a string from snake_case to camelCase
// e.g., this_is_camel => thisIsCamel
func CamelCase(s string) string {
	splits := SplitSnakeCase(s)
	ns := ""
	for i, s := range splits {
		if i != 0 {
			s = strings.Title(s)
		}

		ns += s
	}

	return ns
}

// Initials returns the first letter of each word
// as defined by snake case.
// e.g., this_is_snake_case => []string{"t", "i", "s", "c"}
// Used at times to define the variable used as a struct pointer for
// methods in go i.e., u for User model
func Initials(s string) []string {
	splits := SplitSnakeCase(s)
	for i, s := range splits {
		splits[i] = string(s[0])
	}
	return splits
}

// AppendStrings joins a list of strings together
// useful for templates {{append metis.Name metis.LinkName}}
func AppendStrings(v ...string) {
	n := ""
	for _, s := range v {
		n += s
	}
}

// IsMul returns a flag whether the link's
// multiplicity == metis.Mul
func IsMul(l *Link) bool {
	return l.Multiplicity == Mul
}

// IsID returns a flag whether the trait's
// primitive type == metis.ID
func IsID(t *Trait) bool {
	return t.Type == ID
}
