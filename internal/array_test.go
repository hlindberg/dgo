package internal_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/lyraproj/dgo/dgo"
	require "github.com/lyraproj/dgo/dgo_test"
	"github.com/lyraproj/dgo/newtype"
	"github.com/lyraproj/dgo/typ"
	"github.com/lyraproj/dgo/vf"
)

func ExampleArray() {
	tp := newtype.Array()
	fmt.Println(tp.Equals(typ.Array))
	fmt.Println(tp.Instance(vf.Values(`hello`, 42)))
	fmt.Println(tp.Instance(42))
	// Output:
	// true
	// true
	// false
}

func ExampleArray_min() {
	tp := newtype.Array(2)
	fmt.Println(tp.Instance(vf.Values(`hello`, 42)))
	fmt.Println(tp.Instance(vf.Values(`hello`)))
	// Output:
	// true
	// false
}

func ExampleArray_type() {
	tp := newtype.Array(typ.String)
	fmt.Println(tp.Instance(vf.Values(`hello`)))
	fmt.Println(tp.Instance(vf.Values(42)))
	// Output:
	// true
	// false
}

func ExampleArray_min_max() {
	tp := newtype.Array(1, 2)
	fmt.Println(tp.Instance(vf.Values(`hello`, 42)))
	fmt.Println(tp.Instance(vf.Values(`hello`, 42, `word`)))
	// Output:
	// true
	// false
}

func ExampleArray_type_min() {
	// Create a new array type with a minimum size of 2
	tp := newtype.Array(typ.String, 2)
	fmt.Println(tp.Instance(vf.Values(`hello`, `word`)))
	fmt.Println(tp.Instance(vf.Values(`hello`)))
	// Output:
	// true
	// false
}

func ExampleArray_type_min_max() {
	tp := newtype.Array(typ.String, 2, 3)
	fmt.Println(tp.Instance(vf.Values(`hello`, `word`)))
	// Output: true
}

func TestArray_max_min(t *testing.T) {
	tp := newtype.Array(2, 1)
	require.Equal(t, tp.Min(), 1)
	require.Equal(t, tp.Max(), 2)
}

func TestArray_negative_min(t *testing.T) {
	tp := newtype.Array(-2, 3)
	require.Equal(t, tp.Min(), 0)
	require.Equal(t, tp.Max(), 3)
}

func TestArray_negative_min_max(t *testing.T) {
	tp := newtype.Array(-2, -3)
	require.Equal(t, tp.Min(), 0)
	require.Equal(t, tp.Max(), 0)
}

func TestArray_explicit_unbounded(t *testing.T) {
	tp := newtype.Array(0, math.MaxInt64)
	require.Equal(t, tp, typ.Array)
	require.True(t, tp.Unbounded())
}

func TestArray_badOneArg(t *testing.T) {
	require.Panic(t, func() { newtype.Array(`bad`) }, `illegal argument 1`)
}

func TestArray_badTwoArg(t *testing.T) {
	require.Panic(t, func() { newtype.Array(`bad`, 2) }, `illegal argument 1`)
	require.Panic(t, func() { newtype.Array(typ.String, `bad`) }, `illegal argument 2`)
}

func TestArray_badThreeArg(t *testing.T) {
	require.Panic(t, func() { newtype.Array(`bad`, 2, 2) }, `illegal argument 1`)
	require.Panic(t, func() { newtype.Array(typ.String, `bad`, 2) }, `illegal argument 2`)
	require.Panic(t, func() { newtype.Array(typ.String, 2, `bad`) }, `illegal argument 3`)
}

func TestArray_badArgCount(t *testing.T) {
	require.Panic(t, func() { newtype.Array(typ.String, 2, 2, true) }, `illegal number of arguments`)
}

func TestArrayType(t *testing.T) {
	tp := newtype.Array()
	v := vf.Strings(`a`, `b`)
	require.Instance(t, tp, v)
	require.Assignable(t, tp, newtype.AnyOf(newtype.Array(newtype.String(5, 5)), newtype.Array(newtype.String(8, 8))))
	require.Equal(t, typ.Any, tp.ElementType())
	require.Equal(t, 0, tp.Min())
	require.Equal(t, math.MaxInt64, tp.Max())
	require.True(t, tp.Unbounded())

	require.Instance(t, tp.Type(), tp)
	require.Equal(t, `[]any`, tp.String())

	require.NotEqual(t, 0, tp.HashCode())
	require.Equal(t, tp.HashCode(), tp.HashCode())
}

func TestSizedArrayType(t *testing.T) {
	tp := newtype.Array(typ.String)
	v := vf.Strings(`a`, `b`)
	require.Instance(t, tp, v)
	require.NotInstance(t, tp, `a`)
	require.NotAssignable(t, tp, typ.Array)
	require.Assignable(t, tp, newtype.AnyOf(newtype.Array(newtype.String(5, 5)), newtype.Array(newtype.String(8, 8))))
	require.Equal(t, tp, tp)
	require.NotEqual(t, tp, typ.Array)

	require.Instance(t, tp.Type(), tp)
	require.Equal(t, `[]string`, tp.String())

	tp = newtype.Array(typ.Integer)
	v = vf.Strings(`a`, `b`)
	require.NotInstance(t, tp, v)
	require.NotAssignable(t, tp, v.Type())

	tp = newtype.Array(newtype.AnyOf(typ.String, typ.Integer))
	v = vf.Values(`a`, 3)
	require.Instance(t, tp, v)
	require.Assignable(t, tp, v.Type())
	require.Instance(t, v.Type(), v)
	v = v.With(vf.True)
	require.NotInstance(t, tp, v)
	require.NotAssignable(t, tp, v.Type())

	tp = newtype.Array(0, 2)
	v = vf.Strings(`a`, `b`)
	require.Instance(t, tp, v)

	tp = newtype.Array(0, 1)
	require.NotInstance(t, tp, v)

	tp = newtype.Array(2, 3)
	require.Instance(t, tp, v)

	tp = newtype.Array(3, 3)
	require.NotInstance(t, tp, v)

	tp = newtype.Array(typ.String, 2, 3)
	require.Instance(t, tp, v)

	tp = newtype.Array(typ.Integer, 2, 3)
	require.NotInstance(t, tp, v)

	require.NotEqual(t, 0, tp.HashCode())
	require.NotEqual(t, tp.HashCode(), newtype.Array(typ.Integer).HashCode())
	require.Equal(t, `[2,3]int`, tp.String())

	tp = newtype.Array(newtype.IntegerRange(0, 15), 2, 3)
	require.Equal(t, `[2,3]0..15`, tp.String())

	tp = newtype.Array(newtype.Array(2, 2), 0, 10)
	require.Equal(t, `[0,10][2,2]any`, tp.String())
}

func TestExactArrayType(t *testing.T) {
	v := vf.Strings(`a`, `b`)
	tp := v.Type().(dgo.TupleType)
	require.Instance(t, tp, v)
	require.Equal(t, tp, vf.Strings(`a`, `b`).Type())
	require.NotInstance(t, tp, `a`)
	require.Assignable(t, tp, tp)
	require.NotAssignable(t, tp, typ.Array)

	require.Assignable(t, tp, newtype.Tuple(vf.String(`a`).Type(), vf.String(`b`).Type()))
	require.NotAssignable(t, tp, newtype.Tuple(vf.String(`a`).Type(), vf.String(`b`).Type(), vf.String(`c`).Type()))
	require.NotAssignable(t, tp, newtype.Tuple(vf.String(`a`).Type(), typ.String))

	require.Equal(t, 2, tp.Min())
	require.Equal(t, 2, tp.Max())
	require.False(t, tp.Unbounded())

	require.NotAssignable(t, tp, newtype.AnyOf(newtype.Array(newtype.String(5, 5)), newtype.Array(newtype.String(8, 8))))
	require.Equal(t, tp, tp)
	require.Equal(t, vf.Values(vf.String(`a`).Type(), vf.String(`b`).Type()), tp.ElementTypes())
	require.NotEqual(t, tp, typ.Array)

	require.Instance(t, tp.Type(), tp)
	require.Equal(t, `{"a","b"}`, tp.String())

	require.NotEqual(t, 0, tp.HashCode())
	require.NotEqual(t, tp.HashCode(), newtype.Array(typ.Integer).HashCode())
}

func TestArrayElementType_singleElement(t *testing.T) {
	v := vf.Strings(`hello`)
	at := v.Type()
	et := at.(dgo.ArrayType).ElementType()
	require.Assignable(t, at, at)
	tp := newtype.Array(typ.String)
	require.Assignable(t, tp, at)
	require.NotAssignable(t, at, tp)
	require.Assignable(t, et, vf.String(`hello`).Type())
	require.NotAssignable(t, et, vf.String(`hey`).Type())
	require.Instance(t, et, `hello`)
	require.NotInstance(t, et, `world`)
	require.Equal(t, et, vf.Strings(`hello`).Type().(dgo.ArrayType).ElementType())
	require.Equal(t, et.(dgo.ExactType).Value(), vf.Strings(`hello`))
	require.NotEqual(t, et, vf.Strings(`hello`).Type().(dgo.ArrayType))

	require.NotEqual(t, 0, et.HashCode())
	require.Equal(t, et.HashCode(), et.HashCode())
}

func TestArrayElementType_multipleElements(t *testing.T) {
	v := vf.Strings(`hello`, `world`)
	at := v.Type()
	et := at.(dgo.ArrayType).ElementType()
	require.Assignable(t, at, at)
	tp := newtype.Array(typ.String)
	require.Assignable(t, tp, at)
	require.NotAssignable(t, at, tp)
	require.NotAssignable(t, et, vf.String(`hello`).Type())
	require.NotAssignable(t, et, vf.String(`world`).Type())
	require.NotAssignable(t, vf.String(`hello`).Type(), et)
	require.NotAssignable(t, vf.String(`world`).Type(), et)
	require.Assignable(t, vf.Strings(`hello`, `world`).Type(), at)
	require.Assignable(t, vf.Strings(`hello`, `world`).Type().(dgo.ArrayType).ElementType(), et)
	require.NotAssignable(t, vf.Strings(`world`, `hello`).Type().(dgo.ArrayType).ElementType(), et)

	require.Assignable(t, newtype.Array(2, 2), at)
	require.Assignable(t, newtype.Array(et, 2, 2), at)

	et.(dgo.Iterable).Each(func(v dgo.Value) {
		t.Helper()
		require.Instance(t, typ.String, v)
	})

	require.Instance(t, et.Type(), et)
	require.Equal(t, `"hello"&"world"`, et.String())
}

func TestTupleType(t *testing.T) {
	tt := newtype.Tuple()
	require.Equal(t, typ.Tuple, tt)
	require.Assignable(t, tt, typ.Array)
	require.Assignable(t, tt, typ.Tuple)
	require.Assignable(t, tt, vf.Values(`one`, 2, 3.0).Type())
	require.Equal(t, 0, tt.Min())
	require.Equal(t, math.MaxInt64, tt.Max())
	require.True(t, tt.Unbounded())

	tt = newtype.Tuple(typ.String, typ.Integer, typ.Float)
	require.Assignable(t, tt, tt)
	require.NotAssignable(t, tt, newtype.Tuple(typ.String, typ.Integer, typ.Boolean))
	require.Assignable(t, typ.Array, tt)
	require.Assignable(t, newtype.Array(0, 3), tt)

	require.Assignable(t, tt, vf.Values(`one`, 2, 3.0).Type())
	require.NotAssignable(t, tt, vf.Values(`one`, 2, 3.0, `four`).Type())
	require.NotAssignable(t, tt, vf.Values(`one`, 2, 3).Type())
	require.NotAssignable(t, tt, typ.Array)
	require.NotAssignable(t, tt, typ.Tuple)
	require.NotEqual(t, tt, typ.String)

	require.Equal(t, 3, tt.Min())
	require.Equal(t, 3, tt.Max())
	require.False(t, tt.Unbounded())

	require.Assignable(t, newtype.Array(newtype.AnyOf(typ.String, typ.Integer, typ.Float)), tt)
	require.NotAssignable(t, newtype.Array(newtype.AnyOf(typ.String, typ.Integer)), tt)
	require.Assignable(t, newtype.Array(newtype.AnyOf(typ.String, typ.Integer, typ.Float, typ.Boolean)), tt)

	require.Assignable(t, newtype.Array(newtype.AnyOf(typ.String, typ.Integer, typ.Float), 0, 3), tt)
	require.NotAssignable(t, newtype.Array(newtype.AnyOf(typ.String, typ.Integer, typ.Float), 0, 2), tt)

	okv := vf.Values(`hello`, 1, 2.0)
	require.Instance(t, typ.Tuple, okv)
	require.Instance(t, tt, okv)
	require.NotInstance(t, tt, okv.Get(0))
	require.Assignable(t, tt, okv.Type())
	require.Assignable(t, typ.Array, tt)

	okv = vf.Values(`hello`, 1, 2)
	require.NotInstance(t, tt, okv)

	okv = vf.Values(`hello`, 1, 2.0, true)
	require.NotInstance(t, tt, okv)

	okm := vf.MutableValues(tt, `world`, 2, 3.0)
	require.Panic(t, func() { okm.Add(3) }, newtype.IllegalSize(tt, 4))
	require.Panic(t, func() { okm.Set(2, 3) }, newtype.IllegalAssignment(typ.Float, vf.Value(3)))
	require.Equal(t, `world`, okm.Set(0, `earth`))

	tt = typ.Tuple
	require.Assignable(t, tt, newtype.Array(typ.String, 2, 2))
	tt = newtype.Tuple(typ.String, typ.String)
	require.Assignable(t, tt, newtype.Array(typ.String, 2, 2))
	require.NotAssignable(t, tt, newtype.AnyOf(typ.Nil, newtype.Array(typ.String, 2, 2)))
	tt = newtype.Tuple(typ.String, typ.Integer)
	require.NotAssignable(t, tt, newtype.Array(typ.String, 2, 2))

	require.Equal(t, typ.Any, typ.Tuple.ElementType())
	require.Equal(t, newtype.AllOf(typ.String, typ.Integer), tt.ElementType())
	require.Equal(t, vf.Values(typ.String, typ.Integer), tt.ElementTypes())

	require.Instance(t, tt.Type(), tt)
	require.Equal(t, `{string,int}`, tt.String())

	require.NotEqual(t, 0, tt.HashCode())
	require.Equal(t, tt.HashCode(), tt.HashCode())
}

func TestMutableArray_withoutType(t *testing.T) {
	a := vf.MutableArray(nil, []dgo.Value{nil})
	require.True(t, vf.Nil == a.Get(0))
}

func TestMutableArray_maxSizeMismatch(t *testing.T) {
	a := vf.MutableArray(newtype.Array(0, 1), []dgo.Value{vf.True})
	require.Panic(t, func() { a.Add(vf.False) }, `size constraint`)
	require.Panic(t, func() { a.With(vf.False) }, `size constraint`)
	require.Panic(t, func() { a.WithAll(vf.Values(false)) }, `size constraint`)
	require.Panic(t, func() { a.WithValues(false) }, `size constraint`)
	require.Panic(t, func() { vf.MutableArray(newtype.Array(0, 1), []dgo.Value{vf.True, vf.False}) }, `size constraint`)

	a.WithAll(vf.Values()) // No panic
	a.WithValues()         // No panic
}

func TestMutableArray_minSizeMismatch(t *testing.T) {
	require.Panic(t, func() { vf.MutableArray(newtype.Array(1, 1), []dgo.Value{}) }, `size constraint`)

	a := vf.MutableArray(newtype.Array(typ.Boolean, 1, 1), []dgo.Value{vf.True})
	require.Panic(t, func() { a.Remove(0) }, `size constraint`)
	require.Panic(t, func() { a.RemoveValue(vf.True) }, `size constraint`)
	require.Panic(t, func() { a.Add(vf.True) }, `size constraint`)
	require.Panic(t, func() { a.AddAll(vf.Values(vf.True)) }, `size constraint`)
}

func TestMutableArray_elementTypeMismatch(t *testing.T) {
	require.Panic(t, func() { vf.MutableArray(newtype.Array(typ.String), []dgo.Value{vf.True}) }, `cannot be assigned`)

	a := vf.MutableArray(newtype.Array(typ.String), []dgo.Value{})
	a.Add(`hello`)
	a.AddAll(vf.Values(`hello`))
	a.AddAll(vf.Values())
	require.Panic(t, func() { a.Add(vf.True) }, `cannot be assigned`)
	require.Panic(t, func() { a.AddAll(vf.Values(vf.True)) }, `cannot be assigned`)

	a = vf.MutableValues(newtype.Tuple(typ.String, typ.Integer), `a`, 2)
	a.Set(0, `hello`)
	a.Set(1, 3)
	require.Panic(t, func() { a.Set(0, 3) }, `cannot be assigned`)
}

func TestMutableArray_tupleTypeMismatch(t *testing.T) {
	require.Panic(t, func() { vf.MutableArray(newtype.Tuple(typ.String), []dgo.Value{vf.True}) }, `cannot be assigned`)
}

func TestArray_Set(t *testing.T) {
	a := vf.MutableValues(newtype.Array(typ.Integer))
	a.Add(1)
	a.Set(0, 2)
	require.Equal(t, 2, a.Get(0))

	require.Panic(t, func() { a.Set(0, 1.0) }, `cannot be assigned`)

	f := a.Copy(true)
	require.Panic(t, func() { f.Set(0, 1) }, `Set .* frozen`)
}

func TestArray_SetType(t *testing.T) {
	a := vf.MutableValues(nil, 1, 2.0, `three`)

	at := newtype.Array(newtype.AnyOf(typ.Integer, typ.Float, typ.String))
	a.SetType(at)
	require.Same(t, at, a.Type())

	require.Panic(t, func() { a.SetType(newtype.Array(newtype.AnyOf(typ.Float, typ.String))) },
		`cannot be assigned`)

	a.Freeze()
	require.Panic(t, func() { a.SetType(at) }, `SetType .* frozen`)
}

func TestArray_recursiveFreeze(t *testing.T) {
	a := vf.Array([]dgo.Value{vf.MutableValues(nil, `b`)})
	require.True(t, a.Get(0).(dgo.Array).Frozen())
}

func TestArray_recursiveReflectiveFreeze(t *testing.T) {
	a := vf.Value(
		reflect.ValueOf([]interface{}{
			reflect.ValueOf(vf.MutableValues(nil, `b`))})).(dgo.Array)
	require.True(t, a.Get(0).(dgo.Array).Frozen())
}

func TestArray_replaceNil(t *testing.T) {
	a := vf.Array([]dgo.Value{nil})
	require.True(t, vf.Nil == a.Get(0))
}

func TestArray_fromReflected(t *testing.T) {
	a := vf.Value([]interface{}{`a`, 1, nil}).(dgo.Array)
	require.True(t, a.Frozen())
	require.True(t, vf.Nil == a.Get(2))
}

func TestArray_Add(t *testing.T) {
	a := vf.Values(`a`)
	require.Panic(t, func() { a.Add(vf.Value(`b`)) }, `Add .* frozen`)
	m := a.Copy(false)
	m.Add(vf.Value(`b`))
	require.Equal(t, vf.Values(`a`), a)
	require.Equal(t, vf.Values(`a`, `b`), m)
}

func TestArray_AddAll(t *testing.T) {
	a := vf.Values(`a`)
	require.Panic(t, func() { a.AddAll(vf.Values(`b`)) }, `AddAll .* frozen`)
	m := a.Copy(false)
	m.AddAll(vf.Values(`b`))
	require.Equal(t, vf.Values(`a`), a)
	require.Equal(t, vf.Values(`a`, `b`), m)
}

func TestArray_AddValues(t *testing.T) {
	a := vf.Values(`a`)
	require.Panic(t, func() { a.AddValues(`b`) }, `AddValues .* frozen`)
	m := a.Copy(false)
	m.AddValues(`b`)
	require.Equal(t, vf.Values(`a`), a)
	require.Equal(t, vf.Values(`a`, `b`), m)
}

func TestArray_All(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`)
	i := 0
	require.True(t, a.All(func(e dgo.Value) bool {
		i++
		return e.Equals(`a`) || e.Equals(`b`) || e.Equals(`c`)
	}))
	require.Equal(t, 3, i)

	i = 0
	require.False(t, a.All(func(e dgo.Value) bool {
		i++
		return e.Equals(`a`)
	}))
	require.Equal(t, 2, i)
}

func TestArray_Any(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`)
	i := 0
	require.True(t, a.Any(func(e dgo.Value) bool {
		i++
		return e.Equals(`b`)
	}))
	require.Equal(t, 2, i)

	i = 0
	require.False(t, a.Any(func(e dgo.Value) bool {
		i++
		return e.Equals(`d`)
	}))
	require.Equal(t, 3, i)
}

func TestArray_One(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`)
	i := 0
	require.True(t, a.One(func(e dgo.Value) bool {
		i++
		return e.Equals(`b`)
	}))
	require.Equal(t, 3, i)

	a = vf.Strings(`a`, `b`, `c`, `b`)
	i = 0
	require.False(t, a.One(func(e dgo.Value) bool {
		i++
		return e.Equals(`b`)
	}))
	require.Equal(t, 4, i)

	i = 0
	require.False(t, a.One(func(e dgo.Value) bool {
		i++
		return e.Equals(`d`)
	}))
	require.Equal(t, 4, i)
}

func TestArray_CompareTo(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`)

	c, ok := a.CompareTo(a)
	require.True(t, ok)
	require.Equal(t, 0, c)

	c, ok = a.CompareTo(vf.Nil)
	require.True(t, ok)
	require.Equal(t, 1, c)

	_, ok = a.CompareTo(vf.String(`a`))
	require.False(t, ok)

	b := vf.Strings(`a`, `b`, `c`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 0, c)

	b = vf.Strings(`a`, `b`, `c`, `d`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, -1, c)

	b = vf.Values(`a`, `b`, `d`, `d`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, -1, c)

	b = vf.Values(`a`, `b`, 3, `d`)
	_, ok = a.CompareTo(b)
	require.False(t, ok)

	b = vf.Strings(`a`, `b`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 1, c)

	b = vf.Strings(`a`, `b`, `d`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, -1, c)

	b = vf.Strings(`a`, `b`, `b`)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 1, c)

	a = vf.MutableValues(nil, `a`, `b`)
	a.Add(a)
	b = vf.MutableValues(nil, `a`, `b`)
	b.Add(b)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 0, c)

	b = vf.MutableValues(nil, `a`, `b`)
	b.Add(a)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 0, c)

	a = vf.Values(`a`, 1, nil)
	b = vf.Values(`a`, 1, nil)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 0, c)

	b = vf.Values(`a`, 1, 2)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, -1, c)

	a = vf.Values(`a`, 1, 2)
	b = vf.Values(`a`, 1, nil)
	c, ok = a.CompareTo(b)
	require.True(t, ok)
	require.Equal(t, 1, c)

	a = vf.Values(`a`, 1, 2)
	m := vf.Map(map[string]int{`a`: 1})
	_, ok = a.CompareTo(m)
	require.False(t, ok)

	a = vf.Values(`a`, 1, []int{2})
	b = vf.Values(`a`, 1, 2)
	_, ok = a.CompareTo(b)
	require.False(t, ok)

	a = vf.Values(m)
	b = vf.Values(`a`)
	_, ok = a.CompareTo(b)
	require.False(t, ok)
}

func TestArray_Copy(t *testing.T) {
	a := vf.Values(`a`, `b`, vf.MutableArray(nil, []dgo.Value{vf.String(`c`)}))
	require.Same(t, a, a.Copy(true))
	require.True(t, a.Get(2).(dgo.Freezable).Frozen())

	c := a.Copy(false)
	require.False(t, c.Frozen())
	require.NotSame(t, c, c.Copy(false))

	c = c.Copy(true)
	require.True(t, c.Frozen())
	require.Same(t, c, c.Copy(true))
}

func TestArray_Equal(t *testing.T) {
	a := vf.Values(1, 2)
	require.True(t, a.Equals(a))

	b := vf.Values(1, nil)
	require.False(t, a.Equals(b))

	a = vf.Values(1, nil)
	require.True(t, a.Equals(b))

	b = vf.Values(1, 2)
	require.False(t, a.Equals(b))

	a = vf.Values(1, []int{2})
	require.False(t, a.Equals(b))

	b = vf.Values(1, map[int]int{2: 1})
	require.False(t, a.Equals(b))

	b = vf.Values(1, `2`)
	require.False(t, a.Equals(b))

	// Values containing themselves.
	a = vf.MutableValues(nil, `2`)
	a.Add(a)

	b = vf.MutableValues(nil, `2`)
	b.Add(b)

	m := vf.MutableMap(3, nil)
	m.Put(`me`, m)
	a.Add(m)
	b.Add(m)
	require.True(t, a.Equals(b))

	require.Equal(t, a.HashCode(), b.HashCode())
}

func TestArray_Freeze(t *testing.T) {
	a := vf.MutableValues(nil, `a`, `b`, vf.MutableArray(nil, []dgo.Value{vf.String(`c`)}))
	require.False(t, a.Frozen())

	sa := a.Get(2).(dgo.Array)
	require.False(t, sa.Frozen())

	// In place recursive freeze
	a.Freeze()

	// Sub Array is frozen in place
	require.Same(t, a.Get(2), sa)
	require.True(t, a.Frozen())
	require.True(t, sa.Frozen())
}

func TestArray_FrozenEqual(t *testing.T) {
	f := vf.Values(1, 2, 3)
	require.True(t, f.Frozen(), `not frozen`)

	a := f.Copy(false)
	require.False(t, a.Frozen(), `frozen`)

	require.Equal(t, f, a)
	require.Equal(t, a, f)

	a.Freeze()
	require.True(t, a.Frozen(), `not frozen`)
	require.Same(t, a, a.Copy(true))

	b := a.Copy(false)
	require.NotSame(t, a, b)
	require.NotSame(t, b, b.Copy(true))
	require.NotSame(t, b, b.Copy(false))
}

func TestArray_IndexOf(t *testing.T) {
	a := vf.Values(1, nil, 3)
	require.Equal(t, 2, a.IndexOf(3))
	require.Equal(t, 1, a.IndexOf(nil))
	require.Equal(t, 1, a.IndexOf(vf.Nil))
}

func TestArray_Insert(t *testing.T) {
	a := vf.Values(`a`)
	require.Panic(t, func() { a.Insert(0, vf.Value(`b`)) }, `Insert .* frozen`)
	m := a.Copy(false)
	m.Insert(0, vf.Value(`b`))
	require.Equal(t, vf.Values(`a`), a)
	require.Equal(t, vf.Values(`b`, `a`), m)
}

func TestArray_Map(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`)
	require.Equal(t, vf.Strings(`d`, `e`, `f`), a.Map(func(e dgo.Value) interface{} {
		return string([]byte{e.String()[0] + 3})
	}))
	require.Equal(t, vf.Values(vf.Nil, vf.Nil, vf.Nil), a.Map(func(e dgo.Value) interface{} {
		return nil
	}))
}

func TestArray_MapTo(t *testing.T) {
	require.Equal(t, vf.Integers(97), vf.Strings(`a`).MapTo(nil, func(e dgo.Value) interface{} {
		return int64(e.String()[0])
	}))

	at := newtype.Array(typ.Integer, 2, 3)
	b := vf.Strings(`a`, `b`, `c`).MapTo(at, func(e dgo.Value) interface{} {
		return int64(e.String()[0])
	})

	require.Equal(t, at, b.Type())
	require.Equal(t, vf.Integers(97, 98, 99), b)

	require.Panic(t, func() {
		vf.Strings(`a`, `b`, `c`, `d`).MapTo(at, func(e dgo.Value) interface{} {
			return int64(e.String()[0])
		})
	}, `size constraint`)

	require.Panic(t, func() {
		vf.Strings(`a`).MapTo(at, func(e dgo.Value) interface{} {
			return int64(e.String()[0])
		})
	}, `size constraint`)

	oat := newtype.Array(newtype.AnyOf(typ.Nil, typ.Integer), 0, 3)
	require.Equal(t, vf.Values(97, 98, vf.Nil), vf.Strings(`a`, `b`, `c`).MapTo(oat, func(e dgo.Value) interface{} {
		if e.Equals(`c`) {
			return nil
		}
		return vf.Integer(int64(e.String()[0]))
	}))
	require.Panic(t, func() {
		vf.Strings(`a`, `b`, `c`).MapTo(at, func(e dgo.Value) interface{} {
			if e.Equals(`c`) {
				return nil
			}
			return int64(e.String()[0])
		})
	}, `cannot be assigned`)
}

func TestArray_Pop(t *testing.T) {
	a := vf.Strings(`a`, `b`)
	require.Panic(t, func() { a.Pop() }, `Pop .* frozen`)
	a = a.Copy(false)
	l, ok := a.Pop()
	require.True(t, ok)
	require.Equal(t, `b`, l)
	require.Equal(t, a, vf.Strings(`a`))
	l, ok = a.Pop()
	require.True(t, ok)
	require.Equal(t, `a`, l)
	require.Equal(t, 0, a.Len())
	_, ok = a.Pop()
	require.False(t, ok)
}

func TestArray_Reduce(t *testing.T) {
	a := vf.Integers(1, 2, 3)
	require.Equal(t, 6, a.Reduce(nil, func(memo, v dgo.Value) interface{} {
		if memo == vf.Nil {
			return v
		}
		return memo.(dgo.Integer).GoInt() + v.(dgo.Integer).GoInt()
	}))

	require.Equal(t, vf.Nil, a.Reduce(nil, func(memo, v dgo.Value) interface{} {
		return nil
	}))
}

func TestArray_Remove(t *testing.T) {
	s := vf.Integers(1, 2, 3, 4, 5)
	a := s.Copy(false)
	a.Remove(0)
	require.Equal(t, vf.Integers(2, 3, 4, 5), a)

	a = s.Copy(false)
	a.Remove(4)
	require.Equal(t, vf.Integers(1, 2, 3, 4), a)

	a = s.Copy(false)
	a.Remove(2)
	require.Equal(t, vf.Integers(1, 2, 4, 5), a)

	require.Panic(t, func() { s.Remove(3) }, `Remove .* frozen`)
}

func TestArray_RemoveValue(t *testing.T) {
	s := vf.Integers(1, 2, 3, 4, 5)
	a := s.Copy(false)
	require.True(t, a.RemoveValue(vf.Integer(1)))
	require.Equal(t, vf.Integers(2, 3, 4, 5), a)

	a = s.Copy(false)
	require.True(t, a.RemoveValue(vf.Integer(5)))
	require.Equal(t, vf.Integers(1, 2, 3, 4), a)

	a = s.Copy(false)
	require.True(t, a.RemoveValue(vf.Integer(3)))
	require.Equal(t, vf.Integers(1, 2, 4, 5), a)

	a = s.Copy(false)
	require.False(t, a.RemoveValue(vf.Integer(0)))
	require.Equal(t, vf.Integers(1, 2, 3, 4, 5), a)

	require.Panic(t, func() { s.RemoveValue(vf.Integer(3)) }, `RemoveValue .* frozen`)
}

func TestArray_Reject(t *testing.T) {
	require.Equal(t, vf.Values(1, 2, 4, 5), vf.Values(1, 2, vf.Nil, 4, 5).Reject(func(e dgo.Value) bool {
		return e == vf.Nil
	}))
}

func TestArray_SameValues(t *testing.T) {
	require.True(t, vf.Values().SameValues(vf.Values()))
	require.True(t, vf.Values(1, 2, 3).SameValues(vf.Values(3, 2, 1)))
	require.False(t, vf.Values(1, 2, 4).SameValues(vf.Values(3, 2, 1)))
	require.False(t, vf.Values(1, 2).SameValues(vf.Values(3, 2, 1)))
}

func TestArray_Select(t *testing.T) {
	require.Equal(t, vf.Values(1, 2, 4, 5), vf.Values(1, 2, vf.Nil, 4, 5).Select(func(e dgo.Value) bool {
		return e != vf.Nil
	}))
}

func TestArray_Sort(t *testing.T) {
	a := vf.Strings(`some`, `arbitrary`, `unsorted`, `words`)
	b := a.Sort()
	require.NotEqual(t, a, b)
	c := vf.Strings(`arbitrary`, `some`, `unsorted`, `words`)
	require.Equal(t, b, c)
	require.Equal(t, b, c.Sort())

	a = vf.Values(3.14, -4.2, 2)
	b = a.Sort()
	require.NotEqual(t, a, b)
	c = vf.Values(-4.2, 2, 3.14)
	require.Equal(t, b, c)
	require.Equal(t, b, c.Sort())

	a = vf.Strings(`the one and only`)
	require.Same(t, a, a.Sort())

	a = vf.Values(4.2, `hello`, -3.14)
	b = a.Sort()

	require.Equal(t, b, vf.Values(-3.14, 4.2, `hello`))

}

func TestArray_ToMap(t *testing.T) {
	a := vf.Strings(`a`, `b`, `c`, `d`)
	b := a.ToMap()
	require.Equal(t, b, vf.Map(map[string]string{`a`: `b`, `c`: `d`}))

	a = vf.Strings(`a`, `b`, `c`)
	b = a.ToMap()
	require.Equal(t, b, vf.Map(map[string]interface{}{`a`: `b`, `c`: nil}))
}

func TestArray_ToMapFromEntries(t *testing.T) {
	a := vf.Values(vf.Strings(`a`, `b`), vf.Strings(`c`, `d`))
	b, ok := a.ToMapFromEntries()
	require.True(t, ok)
	require.Equal(t, b, vf.Map(map[string]string{`a`: `b`, `c`: `d`}))

	a = b.Entries()
	b, ok = a.ToMapFromEntries()
	require.True(t, ok)
	require.Equal(t, b, vf.Map(map[string]string{`a`: `b`, `c`: `d`}))

	a = vf.Values(vf.Strings(`a`, `b`), `c`)
	_, ok = a.ToMapFromEntries()
	require.False(t, ok)
}

func TestArray_String(t *testing.T) {
	require.Equal(t, `[1,"two",3.1,true,null]`, vf.Values(1, "two", 3.1, true, nil).String())
}

func TestArray_Unique(t *testing.T) {
	a := vf.Strings(`and`, `some`, `more`, `arbitrary`, `unsorted`, `yes`, `unsorted`, `and`, `yes`, `arbitrary`, `words`)
	b := a.Unique()
	require.NotEqual(t, a, b)
	c := vf.Strings(`and`, `some`, `more`, `arbitrary`, `unsorted`, `yes`, `words`)
	require.Equal(t, b, c)
	require.NotSame(t, b, c)
	require.Equal(t, b, c.Unique())
	require.Same(t, c, c.Unique())

	a = vf.Strings(`the one and only`)
	require.Same(t, a, a.Unique())
}
