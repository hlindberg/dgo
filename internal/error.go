package internal

import (
	"github.com/lyraproj/dgo/dgo"
)

type (
	errw struct {
		error
	}

	errType int
)

const DefaultErrorType = errType(0)

func (t errType) Type() dgo.Type {
	return &metaType{t}
}

func (t errType) Equals(other interface{}) bool {
	return t == DefaultErrorType
}

func (t errType) HashCode() int {
	return int(t.TypeIdentifier())
}

func (t errType) Assignable(other dgo.Type) bool {
	if DefaultErrorType == other {
		return true
	}
	return CheckAssignableTo(nil, other, t)
}

func (t errType) Instance(value interface{}) bool {
	_, ok := value.(error)
	return ok
}

func (t errType) String() string {
	return TypeString(t)
}

func (t errType) TypeIdentifier() dgo.TypeIdentifier {
	return dgo.IdError
}

func (e *errw) Equals(other interface{}) bool {
	if oe, ok := other.(*errw); ok {
		return e.error.Error() == oe.error.Error()
	}
	if oe, ok := other.(error); ok {
		return e.error.Error() == oe.Error()
	}
	return false
}

func (e *errw) HashCode() int {
	return stringHash(e.error.Error())
}

func (e *errw) Error() string {
	return e.error.Error()
}

func (e *errw) String() string {
	return e.error.Error()
}

func (e *errw) Unwrap() error {
	if u, ok := e.error.(interface {
		Unwrap() error
	}); ok {
		return u.Unwrap()
	}
	return nil
}

func (e *errw) Type() dgo.Type {
	return DefaultErrorType
}
