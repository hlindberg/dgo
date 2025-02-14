package newtype

import (
	"github.com/lyraproj/dgo/dgo"
	"github.com/lyraproj/dgo/internal"
)

// IllegalAssignment returns the error that represents an assignment type constraint mismatch
func IllegalAssignment(expected dgo.Type, actual dgo.Value) dgo.Value {
	return internal.IllegalAssignment(expected, actual)
}

// IllegalSize returns the error that represents an size constraint mismatch
func IllegalSize(expected dgo.Type, size int) dgo.Value {
	return internal.IllegalSize(expected, size)
}
