package metis

import (
	"strings"
	"testing"
)

func TestEQ(t *testing.T) {
	one := []string{"one", "two"}
	two := one

	if !eq(one, two) {
		t.Errorf("eq doesn't work, %s and %s should be equal", one, two)
	}

	two = []string{"hey"}

	if eq(one, two) {
		t.Errorf("eq doesn't work, %s and %s should not be equal", one, two)
	}
}

func TestID(t *testing.T) {
	s := "string"

	if id(s) != s {
		t.Errorf("id function doesn't return argument")
	}
}

func TestApply(t *testing.T) {
	s := []string{"one", "two", "three"}
	correct := []string{"ONE", "TWO", "THREE"}

	applied := apply(s, strings.ToUpper)

	if !eq(applied, correct) {
		t.Errorf("Apply didn't work, got %s but expected %v", applied, correct)
	}
}

func TestReduce(t *testing.T) {
	s := []string{"o", "n", "e"}
	correct := "one"

	reduced := reduce("", s, func(one, two string) string {
		return one + two
	})

	if reduced != correct {
		t.Errorf("reduced doesn't work as expected")
	}
}

func TestSplitSnakeCase(t *testing.T) {
	example := "this_is_snake"
	correct := []string{"this", "is", "snake"}

	pieces := SplitSnakeCase(example)

	if !eq(pieces, correct) {
		t.Errorf("SplitSnakeCase didn't work as expected, got %v but wanted %v", pieces, correct)
	}
}

func TestCamelCase(t *testing.T) {
	example := "this_is_snake"
	correct := "thisIsSnake"

	switched := CamelCase(example)

	if switched != correct {
		t.Errorf("CamelCase doesn't work as expected, %s should have been %s", switched, correct)
	}
}

func TestInitials(t *testing.T) {
	example := "this_is_snake"
	correct := []string{"t", "i", "s"}

	applied := Initials(example)

	if !eq(applied, correct) {
		t.Errorf("Initials didn't work as expected, got %s but wanted %s", applied, correct)
	}
}

func TestAppendStrings(t *testing.T) {
	examples := []string{"one", "two", "three"}
	correct := "onetwothree"

	applied := AppendStrings(examples...)

	if applied != correct {
		t.Errorf("AppliedStrings didn't work as expected, got %s, but wanted %s", applied, correct)
	}
}
