package metis

type Multiplicity int

const (
	Mul Multiplicity = iota
	One
)

var multiplicities = map[string]Multiplicity{
	"one": One,
	"mul": Mul,
}
