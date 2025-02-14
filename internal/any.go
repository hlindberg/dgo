package internal

import "github.com/lyraproj/dgo/dgo"

// anyType represents all possible values
type anyType int

const DefaultAnyType = anyType(0)

func (t anyType) Assignable(other dgo.Type) bool {
	return true
}

func (t anyType) Equals(other interface{}) bool {
	return t == other
}

func (t anyType) HashCode() int {
	return int(dgo.IdAny)
}

func (t anyType) Instance(value interface{}) bool {
	return true
}

func (t anyType) String() string {
	return TypeString(t)
}

func (t anyType) Type() dgo.Type {
	return &metaType{t}
}

func (t anyType) TypeIdentifier() dgo.TypeIdentifier {
	return dgo.IdAny
}
