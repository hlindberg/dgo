package internal_test

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"testing"

	"github.com/lyraproj/dgo/dgo"

	require "github.com/lyraproj/dgo/dgo_test"
	"github.com/lyraproj/dgo/vf"
)

func ExampleValueFromReflected() {
	v1 := vf.ValueFromReflected(reflect.ValueOf([]string{`a`, `b`}))
	v2 := vf.Value([]string{`a`, `b`})
	fmt.Println(v1.Equals(v2))
	// Output: true
}

func TestSameInstance(t *testing.T) {
	require.True(t, vf.SameInstance(vf.Nil, vf.Nil))
	require.False(t, vf.SameInstance(vf.True, vf.Nil))
	require.True(t, vf.SameInstance(vf.True, vf.True))
	require.False(t, vf.SameInstance(vf.String(`hello`), vf.String(`hello)`)))
}

func TestValue(t *testing.T) {
	s := vf.String(`a`)
	require.Same(t, s, vf.Value(s))
	require.True(t, vf.True == vf.Value(true))
	require.True(t, vf.False == vf.Value(false))
	require.True(t, vf.Value([]dgo.Value{s}).(dgo.Array).Frozen())
	require.True(t, vf.Value([]string{`a`}).(dgo.Array).Frozen())
	require.Equal(t, vf.Value([]dgo.Value{s}), vf.Value([]string{`a`}))
	require.True(t, vf.Value([]int{1}).(dgo.Array).Frozen())

	v := vf.Value(regexp.MustCompile(`.*`))
	_, ok := v.(dgo.Regexp)
	require.True(t, ok)

	v = vf.Value(int8(42))
	i, ok := v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(int16(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(int32(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(int64(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(uint8(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(uint16(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(uint32(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(uint(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(uint64(42))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	require.Panic(t, func() { vf.Value(uint(math.MaxUint64)) }, `overflows`)
	require.Panic(t, func() { vf.Value(uint64(math.MaxUint64)) }, `overflows`)

	v = vf.Value(float32(3.14))
	f, ok := v.(dgo.Float)
	require.True(t, ok)
	require.True(t, float32(3.14) == float32(f.GoFloat()))

	v = vf.Value(3.14)
	f, ok = v.(dgo.Float)
	require.True(t, ok)
	require.True(t, 3.14 == f.GoFloat())

	v = vf.Value(struct{ A int }{10})
	require.Equal(t, struct{ A int }{10}, v)
}

func TestValue_reflected(t *testing.T) {
	s := vf.String(`a`)
	require.True(t, vf.Nil == vf.Value(reflect.ValueOf(nil)))
	require.True(t, vf.Nil == vf.Value(reflect.ValueOf(([]string)(nil))))
	require.True(t, vf.Nil == vf.Value(reflect.ValueOf((map[string]string)(nil))))
	require.True(t, vf.Nil == vf.Value(reflect.ValueOf((*string)(nil))))

	require.True(t, vf.True == vf.Value(reflect.ValueOf(true)))
	require.True(t, vf.False == vf.Value(reflect.ValueOf(false)))
	require.Same(t, s, vf.Value(reflect.ValueOf(s)))
	require.True(t, vf.Value(reflect.ValueOf([]dgo.Value{s})).(dgo.Array).Frozen())
	require.True(t, vf.Value(reflect.ValueOf([]string{`a`})).(dgo.Array).Frozen())
	require.Equal(t, vf.Value(reflect.ValueOf([]dgo.Value{s})), vf.Value([]string{`a`}))
	require.True(t, vf.Value(reflect.ValueOf([]int{1})).(dgo.Array).Frozen())

	v := vf.Value(reflect.ValueOf(regexp.MustCompile(`.*`)))
	_, ok := v.(dgo.Regexp)
	require.True(t, ok)

	v = vf.Value(reflect.ValueOf(int8(42)))
	i, ok := v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(int16(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(int32(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(int64(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(uint8(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(uint16(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(uint32(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(uint(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	v = vf.Value(reflect.ValueOf(uint64(42)))
	i, ok = v.(dgo.Integer)
	require.True(t, ok)
	require.True(t, 42 == i.GoInt())

	require.Panic(t, func() { vf.Value(reflect.ValueOf(uint(math.MaxUint64))) }, `overflows`)
	require.Panic(t, func() { vf.Value(reflect.ValueOf(uint64(math.MaxUint64))) }, `overflows`)

	v = vf.Value(reflect.ValueOf(float32(3.14)))
	f, ok := v.(dgo.Float)
	require.True(t, ok)
	require.True(t, float32(3.14) == float32(f.GoFloat()))

	v = vf.Value(reflect.ValueOf(3.14))
	f, ok = v.(dgo.Float)
	require.True(t, ok)
	require.True(t, 3.14 == f.GoFloat())

	v = vf.Value(reflect.ValueOf(reflect.ValueOf))
	_, ok = v.(dgo.Native)
	require.True(t, ok)

	require.Panic(t, func() { vf.Value(reflect.ValueOf(struct{ bar int }{bar: 1}).Field(0)) }, `field or method`)
}
