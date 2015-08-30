package metis

import "strings"

// The following are helper functions which bridge the linguistical gaps between
// English and various computer languages.

func eq(one, two []string) bool {
	for i, s := range one {
		if s != two[i] {
			return false
		}
	}

	return true
}

// The identity function for a string
func id(s string) string {
	return s
}

// The apply functional equivalent for strings
func apply(s []string, f func(string) string) []string {
	ns := make([]string, len(s))
	for i, ss := range s {
		ns[i] = f(ss)
	}
	return ns
}

// The reduce functional equivalent for strings
func reduce(start string, strings []string, f func(string, string) string) string {
	ns := start
	apply(strings, func(s string) string {
		ns = f(ns, s)
		return s
	})
	return ns
}

// SplitSnakeCase splits a string based on the "_" character
// e.g., this_is_snake → []string{"this", "is", "snake"}
func SplitSnakeCase(s string) []string {
	return strings.Split(s, "_")
}

// CamelCase turns a string from snake_case to camelCase
// e.g., this_is_camel →  thisIsCamel
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
// useful for templates, e.g., {{append metis.Name metis.LinkName}}
func AppendStrings(v ...string) string {
	return reduce("", v, func(x, y string) string {
		return x + y
	})
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
