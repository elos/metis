package metis

import "strings"

func SplitSnakeCase(s string) []string {
	return strings.Split(s, "_")
}

// this_is_camel => thisIsCamel
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

func Initials(s string) []string {
	splits := SplitSnakeCase(s)
	for i, s := range splits {
		splits[i] = string(s[0])
	}
	return splits
}

func AppendStrings(v ...string) {
	n := ""
	for _, s := range v {
		n += s
	}
}

func IsMul(r *Link) bool {
	return r.Multiplicity == Mul
}

func IsID(t *Trait) bool {
	return t.Type == ID
}
