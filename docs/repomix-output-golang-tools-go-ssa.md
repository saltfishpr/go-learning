This file is a merged representation of a subset of the codebase, containing specifically included files, combined into a single document by Repomix. The content has been processed where security check has been disabled.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Only files matching these patterns are included: go/ssa/**
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Security check has been disabled - content may contain sensitive information

## Additional Info

# Directory Structure
```
go/
  ssa/
    interp/
      testdata/
        fixedbugs/
          issue52342.go
          issue52835.go
          issue55086.go
          issue55115.go
          issue66783.go
          issue69298.go
          issue69929.go
        src/
          encoding/
            encoding.go
          errors/
            errors.go
          fmt/
            fmt.go
          io/
            io.go
          log/
            log.go
          math/
            math.go
          os/
            os.go
          reflect/
            deepequal.go
            reflect.go
          runtime/
            runtime.go
          sort/
            sort.go
          strconv/
            strconv.go
          strings/
            strings.go
          sync/
            sync.go
          time/
            time.go
          unicode/
            utf8/
              utf8.go
          unsafe/
            unsafe.go
        boundmeth.go
        complit.go
        convert.go
        coverage.go
        deepequal.go
        defer.go
        fieldprom.go
        forvarlifetime_go122.go
        forvarlifetime_old.go
        ifaceconv.go
        ifaceprom.go
        initorder.go
        methprom.go
        minmax.go
        mrvchain.go
        range.go
        rangefunc.go
        rangeoverint.go
        rangevarlifetime_go122.go
        rangevarlifetime_old.go
        recover.go
        reflect.go
        slice2array.go
        slice2arrayptr.go
        static.go
        typeassert.go
        width32.go
        zeros.go
      external.go
      interp_test.go
      interp.go
      map.go
      ops.go
      rangefunc_test.go
      reflect.go
      value.go
    ssautil/
      testdata/
        switches.txtar
      deprecated_test.go
      deprecated.go
      load_test.go
      load.go
      switch_test.go
      switch.go
      visit.go
    testdata/
      fixedbugs/
        issue66783a.go
        issue66783b.go
      src/
        bytes/
          bytes.go
        context/
          context.go
        encoding/
          json/
            json.go
          xml/
            xml.go
          encoding.go
        errors/
          errors.go
        fmt/
          fmt.go
        io/
          io.go
        log/
          log.go
        math/
          math.go
        os/
          os.go
        reflect/
          reflect.go
        runtime/
          runtime.go
        sort/
          sort.go
        strconv/
          strconv.go
        strings/
          strings.go
        sync/
          atomic/
            atomic.go
          sync.go
        time/
          time.go
        unsafe/
          unsafe.go
        README.txt
      indirect.txtar
      objlookup.go
      structconv.go
      valueforexpr.go
    block.go
    blockopt.go
    builder_generic_test.go
    builder_test.go
    builder.go
    const_test.go
    const.go
    create.go
    doc.go
    dom_test.go
    dom.go
    emit.go
    example_test.go
    func.go
    instantiate_test.go
    instantiate.go
    lift.go
    lvalue.go
    methods_test.go
    methods.go
    mode.go
    print.go
    sanity.go
    source_test.go
    source.go
    ssa.go
    stdlib_test.go
    subst_test.go
    subst.go
    task.go
    testutil_test.go
    TODO
    typeset.go
    util.go
    wrappers.go
```

# Files

## File: go/ssa/interp/testdata/fixedbugs/issue52342.go
```go
package main

func main() {
	var d byte

	d = 1
	d <<= 256
	if d != 0 {
		panic(d)
	}

	d = 1
	d >>= 256
	if d != 0 {
		panic(d)
	}
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue52835.go
```go
package main

var called bool

type I interface {
	Foo()
}

type A struct{}

func (a A) Foo() {
	called = true
}

func lambda[X I]() func() func() {
	return func() func() {
		var x X
		return x.Foo
	}
}

func main() {
	lambda[A]()()()
	if !called {
		panic(called)
	}
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue55086.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func a() (r string) {
	s := "initial"
	var p *struct{ i int }
	defer func() {
		recover()
		r = s
	}()

	s, p.i = "set", 2 // s must be set before p.i panics
	return "unreachable"
}

func b() (r string) {
	s := "initial"
	fn := func() []int { panic("") }
	defer func() {
		recover()
		r = s
	}()

	s, fn()[0] = "set", 2 // fn() panics before any assignment occurs
	return "unreachable"
}

func c() (r string) {
	s := "initial"
	var p map[int]int
	defer func() {
		recover()
		r = s
	}()

	s, p[0] = "set", 2 //s must be set before p[0] index panics"
	return "unreachable"
}

func d() (r string) {
	s := "initial"
	var p map[int]int
	defer func() {
		recover()
		r = s
	}()
	fn := func() int { panic("") }

	s, p[0] = "set", fn() // fn() panics before s is set
	return "unreachable"
}

func e() (r string) {
	s := "initial"
	p := map[int]int{}
	defer func() {
		recover()
		r = s
	}()
	fn := func() int { panic("") }

	s, p[fn()] = "set", 0 // fn() panics before any assignment occurs
	return "unreachable"
}

func f() (r string) {
	s := "initial"
	p := []int{}
	defer func() {
		recover()
		r = s
	}()

	s, p[1] = "set", 0 // p[1] panics after s is set
	return "unreachable"
}

func g() (r string) {
	s := "initial"
	p := map[any]any{}
	defer func() {
		recover()
		r = s
	}()
	var i any = func() {}
	s, p[i] = "set", 0 // p[i] panics after s is set
	return "unreachable"
}

func h() (r string) {
	fail := false
	defer func() {
		recover()
		if fail {
			r = "fail"
		} else {
			r = "success"
		}
	}()

	type T struct{ f int }
	var p *struct{ *T }

	// The implicit "p.T" operand should be evaluated in phase 1 (and panic),
	// before the "fail = true" assignment in phase 2.
	fail, p.f = true, 0
	return "unreachable"
}

func main() {
	for _, test := range []struct {
		fn   func() string
		want string
		desc string
	}{
		{a, "set", "s must be set before p.i panics"},
		{b, "initial", "p() panics before s is set"},
		{c, "set", "s must be set before p[0] index panics"},
		{d, "initial", "fn() panics before s is set"},
		{e, "initial", "fn() panics before s is set"},
		{f, "set", "p[1] panics after s is set"},
		{g, "set", "p[i] panics after s is set"},
		{h, "success", "p.T panics before fail is set"},
	} {
		if test.fn() != test.want {
			panic(test.desc)
		}
	}
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue55115.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "reflect"

func main() {
	type MyByte byte
	type MyRune rune
	type MyString string

	a := []MyByte{'a', 'b', 'c'}
	if s := string(a); s != "abc" {
		panic(s)
	}

	b := []MyRune{'五', '五'}
	if s := string(b); s != "五五" {
		panic(s)
	}

	c := []MyByte{'l', 'o', 'r', 'e', 'm'}
	if s := MyString(c); s != MyString("lorem") {
		panic(s)
	}

	d := "lorem"
	if a := []MyByte(d); !reflect.DeepEqual(a, []MyByte{'l', 'o', 'r', 'e', 'm'}) {
		panic(a)
	}

	e := 42
	if s := MyString(e); s != "*" {
		panic(s)
	}
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue66783.go
```go
package main

import "fmt"

func Fn[N any]() (any, any, any) {
	// Very recursive type to exercise substitution.
	type t[x any, ignored *N] struct {
		f  x
		g  N
		nx *t[x, *N]
		nn *t[N, *N]
	}
	n := t[N, *N]{}
	s := t[string, *N]{}
	i := t[int, *N]{}
	return n, s, i
}

func main() {

	sn, ss, si := Fn[string]()
	in, is, ii := Fn[int]()

	for i, t := range []struct {
		x, y any
		want bool
	}{
		{sn, ss, true},  // main.t[string;string,*string] == main.t[string;string,*string]
		{sn, si, false}, // main.t[string;string,*string] != main.t[string;int,*string]
		{sn, in, false}, // main.t[string;string,*string] != main.t[int;int,*int]
		{sn, is, false}, // main.t[string;string,*string] != main.t[int;string,*int]
		{sn, ii, false}, // main.t[string;string,*string] != main.t[int;int,*int]

		{ss, si, false}, // main.t[string;string,*string] != main.t[string;int,*string]
		{ss, in, false}, // main.t[string;string,*string] != main.t[int;int,*int]
		{ss, is, false}, // main.t[string;string,*string] != main.t[int;string,*int]
		{ss, ii, false}, // main.t[string;string,*string] != main.t[int;int,*int]

		{si, in, false}, // main.t[string;int,*string] != main.t[int;int,*int]
		{si, is, false}, // main.t[string;int,*string] != main.t[int;string,*int]
		{si, ii, false}, // main.t[string;int,*string] != main.t[int;int,*int]

		{in, is, false}, // main.t[int;int,*int] != main.t[int;string,*int]
		{in, ii, true},  // main.t[int;int,*int] == main.t[int;int,*int]

		{is, ii, false}, // main.t[int;string,*int] != main.t[int;int,*int]
	} {
		x, y, want := t.x, t.y, t.want
		if got := x == y; got != want {
			msg := fmt.Sprintf("(case %d) %T == %T. got %v. wanted %v", i, x, y, got, want)
			panic(msg)
		}
	}
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue69298.go
```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

type Seq[V any] func(yield func(V) bool)

func AppendSeq[Slice ~[]E, E any](s Slice, seq Seq[E]) Slice {
	for v := range seq {
		s = append(s, v)
	}
	return s
}

func main() {
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}

	s := AppendSeq([]int{1, 2}, seq)
	fmt.Println(s)
}
```

## File: go/ssa/interp/testdata/fixedbugs/issue69929.go
```go
package main

// This is a regression test for a bug (#69929) in
// the SSA interpreter in which it would not execute phis in parallel.
//
// The insert function below has interdependent phi nodes:
//
//	  entry:
//		t0 = *root       // t0 is x or y before loop
//		jump test
//	  body:
//		print(t5)      // t5 is x at loop entry
//		t3 = t5.Child    // t3 is x after loop
//		jump test
//	  test:
//		t5 = phi(t0, t3) // t5 is x at loop entry
//		t6 = phi(t0, t5) // t6 is y at loop entry
//		if t5 != nil goto body else done
//	  done:
//		print(t6)
//		return
//
// The two phis:
//
//	t5 = phi(t0, t3)
//	t6 = phi(t0, t5)
//
// must be executed in parallel as if they were written in Go
// as:
//
//	t5, t6 = phi(t0, t3), phi(t0, t5)
//
// with the second phi node observing the original, not
// updated, value of t5. (In more complex examples, the phi
// nodes may be mutually recursive, breaking partial solutions
// based on simple reordering of the phi instructions. See the
// Briggs paper for detail.)
//
// The correct behavior is print(1, root); print(2, root); print(3, root).
// The previous incorrect behavior had print(2, nil).

func main() {
	insert()
	print(3, root)
}

var root = new(node)

type node struct{ child *node }

func insert() {
	x := root
	y := x
	for x != nil {
		y = x
		print(1, y)
		x = x.child
	}
	print(2, y)
}

func print(order int, ptr *node) {
	println(order, ptr)
	if ptr != root {
		panic(ptr)
	}
}
```

## File: go/ssa/interp/testdata/src/encoding/encoding.go
```go
package encoding

type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}
type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}

type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}
type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
```

## File: go/ssa/interp/testdata/src/errors/errors.go
```go
package errors

func New(text string) error { return errorString{text} }

type errorString struct{ s string }

func (e errorString) Error() string { return e.s }
```

## File: go/ssa/interp/testdata/src/fmt/fmt.go
```go
package fmt

import (
	"errors"
	"strings"
)

func Sprint(args ...interface{}) string

func Sprintln(args ...interface{}) string {
	return Sprint(args...) + "\n"
}

func Print(args ...interface{}) (int, error) {
	var n int
	for i, arg := range args {
		if i > 0 {
			print(" ")
			n++
		}
		msg := Sprint(arg)
		n += len(msg)
		print(msg)
	}
	return n, nil
}

func Println(args ...interface{}) {
	Print(args...)
	println()
}

// formatting is too complex to fake
// handle the bare minimum needed for tests

func Printf(format string, args ...interface{}) (int, error) {
	msg := Sprintf(format, args...)
	print(msg)
	return len(msg), nil
}

func Sprintf(format string, args ...interface{}) string {
	// handle extremely simple cases that appear in tests.
	if len(format) == 0 {
		return ""
	}
	switch {
	case strings.HasPrefix("%v", format) || strings.HasPrefix("%s", format):
		return Sprint(args[0]) + Sprintf(format[2:], args[1:]...)
	case !strings.HasPrefix("%", format):
		return format[:1] + Sprintf(format[1:], args...)
	default:
		panic("unsupported format string for testing Sprintf")
	}
}

func Errorf(format string, args ...interface{}) error {
	msg := Sprintf(format, args...)
	return errors.New(msg)
}
```

## File: go/ssa/interp/testdata/src/io/io.go
```go
package io

import "errors"

var EOF = errors.New("EOF")
```

## File: go/ssa/interp/testdata/src/log/log.go
```go
package log

import (
	"fmt"
	"os"
)

func Println(v ...interface{}) {
	fmt.Println(v...)
}
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func Fatalln(v ...interface{}) {
	Println(v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Printf(format, v...)
	os.Exit(1)
}
```

## File: go/ssa/interp/testdata/src/math/math.go
```go
package math

func Copysign(float64, float64) float64

func NaN() float64

func Inf(int) float64

func IsNaN(float64) bool

func Float64bits(float64) uint64

func Signbit(x float64) bool {
	return Float64bits(x)&(1<<63) != 0
}

func Sqrt(x float64) float64
```

## File: go/ssa/interp/testdata/src/os/os.go
```go
package os

func Getenv(string) string

func Exit(int)
```

## File: go/ssa/interp/testdata/src/reflect/deepequal.go
```go
package reflect

// Not an actual implementation of DeepEqual. This is a model that supports
// the bare minimum needed to get through testing interp.
//
// Does not handle cycles.
//
// Note: unclear if reflect.go can support this.
func DeepEqual(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := ValueOf(x)
	v2 := ValueOf(y)

	return deepValueEqual(v1, v2, make(map[visit]bool))
}

// Key for the visitedMap in deepValueEqual.
type visit struct {
	a1, a2 uintptr
	typ    Type
}

func deepValueEqual(v1, v2 Value, visited map[visit]bool) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}

	// Short circuit on reference types that can lead to cycles in comparison.
	switch v1.Kind() {
	case Pointer, Map, Slice, Interface:
		k := visit{v1.Pointer(), v2.Pointer(), v1.Type()} // Not safe for moving GC.
		if visited[k] {
			// The comparison algorithm assumes that all checks in progress are true when it reencounters them.
			return true
		}
		visited[k] = true
	}

	switch v1.Kind() {
	case Array:
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited) {
				return false
			}
		}
		return true
	case Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i), visited) {
				return false
			}
		}
		return true
	case Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited)
	case Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return deepValueEqual(v1.Elem(), v2.Elem(), visited)
	case Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i), visited) {
				return false
			}
		}
		return true
	case Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(val1, val2, visited) {
				return false
			}
		}
		return true
	case Func:
		return v1.IsNil() && v2.IsNil()
	default:
		// Normal equality suffices
		return v1.Interface() == v2.Interface() // try interface comparison as a fallback.
	}
}
```

## File: go/ssa/interp/testdata/src/reflect/reflect.go
```go
package reflect

type Type interface {
	String() string
	Kind() Kind
	Elem() Type
}

type Value struct {
}

func (Value) String() string

func (Value) Elem() Value
func (Value) Kind() Kind
func (Value) Int() int64
func (Value) IsValid() bool
func (Value) IsNil() bool
func (Value) Len() int
func (Value) Pointer() uintptr
func (Value) Index(i int) Value
func (Value) Type() Type
func (Value) Field(int) Value
func (Value) MapIndex(Value) Value
func (Value) MapKeys() []Value
func (Value) NumField() int
func (Value) Interface() interface{}

func SliceOf(Type) Type

func TypeOf(interface{}) Type

func ValueOf(interface{}) Value

type Kind uint

// Constants need to be kept in sync with the actual definitions for comparisons in tests.
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

const Ptr = Pointer
```

## File: go/ssa/interp/testdata/src/runtime/runtime.go
```go
package runtime

// An errorString represents a runtime error described by a single string.
type errorString string

func (e errorString) RuntimeError() {}

func (e errorString) Error() string {
	return "runtime error: " + string(e)
}

func Breakpoint()

type Error interface {
	error
	RuntimeError()
}

func GC()
```

## File: go/ssa/interp/testdata/src/sort/sort.go
```go
package sort

func Strings(x []string)
func Ints(x []int)
func Float64s(x []float64)
```

## File: go/ssa/interp/testdata/src/strconv/strconv.go
```go
package strconv

func Itoa(i int) string
func Atoi(s string) (int, error)

func FormatFloat(float64, byte, int, int) string
```

## File: go/ssa/interp/testdata/src/strings/strings.go
```go
package strings

func Replace(s, old, new string, n int) string

func Index(haystack, needle string) int

func Contains(haystack, needle string) bool {
	return Index(haystack, needle) >= 0
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func EqualFold(s, t string) bool
func ToLower(s string) string

type Builder struct {
	s string
}

func (b *Builder) WriteString(s string) (int, error) {
	b.s += s
	return len(s), nil
}
func (b *Builder) String() string { return b.s }
```

## File: go/ssa/interp/testdata/src/sync/sync.go
```go
package sync

// Rudimentary implementation of a mutex for interp tests.
type Mutex struct {
	c chan int // Mutex is held when held c!=nil and is empty. Access is guarded by g.
}

func (m *Mutex) Lock() {
	c := ch(m)
	<-c
}

func (m *Mutex) Unlock() {
	c := ch(m)
	c <- 1
}

// sequentializes Mutex.c access.
var g = make(chan int, 1)

func init() {
	g <- 1
}

// ch initializes the m.c field if needed and returns it.
func ch(m *Mutex) chan int {
	<-g
	defer func() {
		g <- 1
	}()
	if m.c == nil {
		m.c = make(chan int, 1)
		m.c <- 1
	}
	return m.c
}
```

## File: go/ssa/interp/testdata/src/time/time.go
```go
package time

type Duration int64

func Sleep(Duration)
```

## File: go/ssa/interp/testdata/src/unicode/utf8/utf8.go
```go
package utf8

func DecodeRuneInString(string) (rune, int)

func DecodeRune(b []byte) (rune, int) {
	return DecodeRuneInString(string(b))
}

const RuneError = '\uFFFD'
```

## File: go/ssa/interp/testdata/src/unsafe/unsafe.go
```go
package unsafe
```

## File: go/ssa/interp/testdata/boundmeth.go
```go
// Tests of bound method closures.

package main

import (
	"errors"
	"fmt"
)

func assert(b bool) {
	if !b {
		panic("oops")
	}
}

type I int

func (i I) add(x int) int {
	return int(i) + x
}

func valueReceiver() {
	var three I = 3
	assert(three.add(5) == 8)
	var add3 func(int) int = three.add
	assert(add3(5) == 8)
}

type S struct{ x int }

func (s *S) incr() {
	s.x++
}

func (s *S) get() int {
	return s.x
}

func pointerReceiver() {
	ps := new(S)
	incr := ps.incr
	get := ps.get
	assert(get() == 0)
	incr()
	incr()
	incr()
	assert(get() == 3)
}

func addressibleValuePointerReceiver() {
	var s S
	incr := s.incr
	get := s.get
	assert(get() == 0)
	incr()
	incr()
	incr()
	assert(get() == 3)
}

type S2 struct {
	S
}

func promotedReceiver() {
	var s2 S2
	incr := s2.incr
	get := s2.get
	assert(get() == 0)
	incr()
	incr()
	incr()
	assert(get() == 3)
}

func anonStruct() {
	var s struct{ S }
	incr := s.incr
	get := s.get
	assert(get() == 0)
	incr()
	incr()
	incr()
	assert(get() == 3)
}

func typeCheck() {
	var i interface{}
	i = (*S).incr
	_ = i.(func(*S)) // type assertion: receiver type prepended to params

	var s S
	i = s.incr
	_ = i.(func()) // type assertion: receiver type disappears
}

type errString string

func (err errString) Error() string {
	return string(err)
}

// Regression test for a builder crash.
func regress1(x error) func() string {
	return x.Error
}

// Regression test for b/7269:
// taking the value of an interface method performs a nil check.
func nilInterfaceMethodValue() {
	err := errors.New("ok")
	f := err.Error
	if got := f(); got != "ok" {
		panic(got)
	}

	err = nil
	if got := f(); got != "ok" {
		panic(got)
	}

	defer func() {
		r := fmt.Sprint(recover())
		// runtime panic string varies across toolchains
		if r != "interface conversion: interface is nil, not error" &&
			r != "runtime error: invalid memory address or nil pointer dereference" &&
			r != "method value: interface is nil" {
			panic("want runtime panic from nil interface method value, got " + r)
		}
	}()
	f = err.Error // runtime panic: err is nil
	panic("unreachable")
}

func main() {
	valueReceiver()
	pointerReceiver()
	addressibleValuePointerReceiver()
	promotedReceiver()
	anonStruct()
	typeCheck()

	if e := regress1(errString("hi"))(); e != "hi" {
		panic(e)
	}

	nilInterfaceMethodValue()
}
```

## File: go/ssa/interp/testdata/complit.go
```go
package main

// Tests of composite literals.

import "fmt"

// Map literals.
// TODO(adonovan): we can no longer print maps
// until the interpreter supports (reflect.Value).MapRange.
func _() {
	type M map[int]int
	m1 := []*M{{1: 1}, &M{2: 2}}
	want := "map[1:1] map[2:2]"
	if got := fmt.Sprint(*m1[0], *m1[1]); got != want {
		panic(got)
	}
	m2 := []M{{1: 1}, M{2: 2}}
	if got := fmt.Sprint(m2[0], m2[1]); got != want {
		panic(got)
	}
}

// Nonliteral keys in composite literal.
func init() {
	const zero int = 1
	var v = []int{1 + zero: 42}
	if x := fmt.Sprint(v); x != "[0 0 42]" {
		panic(x)
	}
}

// Test for in-place initialization.
func init() {
	// struct
	type S struct {
		a, b int
	}
	s := S{1, 2}
	s = S{b: 3}
	if s.a != 0 {
		panic("s.a != 0")
	}
	if s.b != 3 {
		panic("s.b != 3")
	}
	s = S{}
	if s.a != 0 {
		panic("s.a != 0")
	}
	if s.b != 0 {
		panic("s.b != 0")
	}

	// array
	type A [4]int
	a := A{2, 4, 6, 8}
	a = A{1: 6, 2: 4}
	if a[0] != 0 {
		panic("a[0] != 0")
	}
	if a[1] != 6 {
		panic("a[1] != 6")
	}
	if a[2] != 4 {
		panic("a[2] != 4")
	}
	if a[3] != 0 {
		panic("a[3] != 0")
	}
	a = A{}
	if a[0] != 0 {
		panic("a[0] != 0")
	}
	if a[1] != 0 {
		panic("a[1] != 0")
	}
	if a[2] != 0 {
		panic("a[2] != 0")
	}
	if a[3] != 0 {
		panic("a[3] != 0")
	}
}

// Regression test for https://golang.org/issue/10127:
// composite literal clobbers destination before reading from it.
func init() {
	// map
	{
		type M map[string]int
		m := M{"x": 1, "y": 2}
		m = M{"x": m["y"], "y": m["x"]}
		if m["x"] != 2 || m["y"] != 1 {
			panic(fmt.Sprint(m))
		}

		n := M{"x": 3}
		m, n = M{"x": n["x"]}, M{"x": m["x"]} // parallel assignment
		if got := fmt.Sprint(m["x"], n["x"]); got != "3 2" {
			panic(got)
		}
	}

	// struct
	{
		type T struct{ x, y, z int }
		t := T{x: 1, y: 2, z: 3}

		t = T{x: t.y, y: t.z, z: t.x} // all fields
		if got := fmt.Sprint(t); got != "{2 3 1}" {
			panic(got)
		}

		t = T{x: t.y, y: t.z + 3} // not all fields
		if got := fmt.Sprint(t); got != "{3 4 0}" {
			panic(got)
		}

		u := T{x: 5, y: 6, z: 7}
		t, u = T{x: u.x}, T{x: t.x} // parallel assignment
		if got := fmt.Sprint(t, u); got != "{5 0 0} {3 0 0}" {
			panic(got)
		}
	}

	// array
	{
		a := [3]int{0: 1, 1: 2, 2: 3}

		a = [3]int{0: a[1], 1: a[2], 2: a[0]} //  all elements
		if got := fmt.Sprint(a); got != "[2 3 1]" {
			panic(got)
		}

		a = [3]int{0: a[1], 1: a[2] + 3} //  not all elements
		if got := fmt.Sprint(a); got != "[3 4 0]" {
			panic(got)
		}

		b := [3]int{0: 5, 1: 6, 2: 7}
		a, b = [3]int{0: b[0]}, [3]int{0: a[0]} // parallel assignment
		if got := fmt.Sprint(a, b); got != "[5 0 0] [3 0 0]" {
			panic(got)
		}
	}

	// slice
	{
		s := []int{0: 1, 1: 2, 2: 3}

		s = []int{0: s[1], 1: s[2], 2: s[0]} //  all elements
		if got := fmt.Sprint(s); got != "[2 3 1]" {
			panic(got)
		}

		s = []int{0: s[1], 1: s[2] + 3} //  not all elements
		if got := fmt.Sprint(s); got != "[3 4]" {
			panic(got)
		}

		t := []int{0: 5, 1: 6, 2: 7}
		s, t = []int{0: t[0]}, []int{0: s[0]} // parallel assignment
		if got := fmt.Sprint(s, t); got != "[5] [3]" {
			panic(got)
		}
	}
}

// Regression test for https://golang.org/issue/13341:
// within a map literal, if a key expression is a composite literal,
// Go 1.5 allows its type to be omitted.  An & operation may be implied.
func init() {
	type S struct{ x int }
	// same as map[*S]bool{&S{x: 1}: true}
	m := map[*S]bool{{x: 1}: true}
	for s := range m {
		if s.x != 1 {
			panic(s) // wrong key
		}
		return
	}
	panic("map is empty")
}

func main() {
}
```

## File: go/ssa/interp/testdata/convert.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test conversion operations.

package main

func left(x int)  { _ = 1 << x }
func right(x int) { _ = 1 >> x }

func main() {
	wantPanic(
		func() {
			left(-1)
		},
		"runtime error: negative shift amount",
	)
	wantPanic(
		func() {
			right(-1)
		},
		"runtime error: negative shift amount",
	)
	wantPanic(
		func() {
			const maxInt32 = 1<<31 - 1
			var idx int64 = maxInt32*2 + 8
			x := make([]int, 16)
			_ = x[idx]
		},
		"runtime error: runtime error: index out of range [4294967302] with length 16",
	)
}

func wantPanic(fn func(), s string) {
	defer func() {
		err := recover()
		if err == nil {
			panic("expected panic")
		}
		if got := err.(error).Error(); got != s {
			panic("expected panic " + s + " got " + got)
		}
	}()
	fn()
}
```

## File: go/ssa/interp/testdata/coverage.go
```go
// This interpreter test is designed to run very quickly yet provide
// some coverage of a broad selection of constructs.
//
// Validate this file with 'go run' after editing.
// TODO(adonovan): break this into small files organized by theme.

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func init() {
	// Call of variadic function with (implicit) empty slice.
	if x := fmt.Sprint(); x != "" {
		panic(x)
	}
}

type empty interface{}

type I interface {
	f() int
}

type T struct{ z int }

func (t T) f() int { return t.z }

func use(interface{}) {}

var counter = 2

// Test initialization, including init blocks containing 'return'.
// Assertion is in main.
func init() {
	counter *= 3
	return
	counter *= 3
}

func init() {
	counter *= 5
	return
	counter *= 5
}

// Recursion.
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fibgen(ch chan int) {
	for x := 0; x < 10; x++ {
		ch <- fib(x)
	}
	close(ch)
}

// Goroutines and channels.
func init() {
	ch := make(chan int)
	go fibgen(ch)
	var fibs []int
	for v := range ch {
		fibs = append(fibs, v)
		if len(fibs) == 10 {
			break
		}
	}
	if x := fmt.Sprint(fibs); x != "[0 1 1 2 3 5 8 13 21 34]" {
		panic(x)
	}
}

// Test of aliasing.
func init() {
	type S struct {
		a, b string
	}

	s1 := []string{"foo", "bar"}
	s2 := s1 // creates an alias
	s2[0] = "wiz"
	if x := fmt.Sprint(s1, s2); x != "[wiz bar] [wiz bar]" {
		panic(x)
	}

	pa1 := &[2]string{"foo", "bar"}
	pa2 := pa1 // creates an alias
	pa2[0] = "wiz"
	if x := fmt.Sprint(*pa1, *pa2); x != "[wiz bar] [wiz bar]" {
		panic(x)
	}

	a1 := [2]string{"foo", "bar"}
	a2 := a1 // creates a copy
	a2[0] = "wiz"
	if x := fmt.Sprint(a1, a2); x != "[foo bar] [wiz bar]" {
		panic(x)
	}

	t1 := S{"foo", "bar"}
	t2 := t1 // copy
	t2.a = "wiz"
	if x := fmt.Sprint(t1, t2); x != "{foo bar} {wiz bar}" {
		panic(x)
	}
}

func main() {
	print() // legal

	if counter != 2*3*5 {
		panic(counter)
	}

	// Test builtins (e.g. complex) preserve named argument types.
	type N complex128
	var n N
	n = complex(1.0, 2.0)
	if n != complex(1.0, 2.0) {
		panic(n)
	}
	if x := reflect.TypeOf(n).String(); x != "main.N" {
		panic(x)
	}
	if real(n) != 1.0 || imag(n) != 2.0 {
		panic(n)
	}

	// Channel + select.
	ch := make(chan int, 1)
	select {
	case ch <- 1:
		// ok
	default:
		panic("couldn't send")
	}
	if <-ch != 1 {
		panic("couldn't receive")
	}
	// A "receive" select-case that doesn't declare its vars.  (regression test)
	anint := 0
	ok := false
	select {
	case anint, ok = <-ch:
	case anint = <-ch:
	default:
	}
	_ = anint
	_ = ok

	// Anon structs with methods.
	anon := struct{ T }{T: T{z: 1}}
	if x := anon.f(); x != 1 {
		panic(x)
	}
	var i I = anon
	if x := i.f(); x != 1 {
		panic(x)
	}
	// NB. precise output of reflect.Type.String is undefined.
	if x := reflect.TypeOf(i).String(); x != "struct { main.T }" && x != "struct{main.T}" {
		panic(x)
	}

	// fmt.
	const message = "Hello, World!"
	if fmt.Sprint("Hello", ", ", "World", "!") != message {
		panic("oops")
	}

	// Type assertion.
	type S struct {
		f int
	}
	var e empty = S{f: 42}
	switch v := e.(type) {
	case S:
		if v.f != 42 {
			panic(v.f)
		}
	default:
		panic(reflect.TypeOf(v))
	}
	if i, ok := e.(I); ok {
		panic(i)
	}

	// Switch.
	var x int
	switch x {
	case 1:
		panic(x)
		fallthrough
	case 2, 3:
		panic(x)
	default:
		// ok
	}
	// empty switch
	switch {
	}
	// empty switch
	switch {
	default:
	}
	// empty switch
	switch {
	default:
		fallthrough
	case false:
	}

	// string -> []rune conversion.
	use([]rune("foo"))

	// Calls of form x.f().
	type S2 struct {
		f func() int
	}
	S2{f: func() int { return 1 }}.f() // field is a func value
	T{}.f()                            // method call
	i.f()                              // interface method invocation
	(interface {
		f() int
	}(T{})).f() // anon interface method invocation

	// Map lookup.
	if v, ok := map[string]string{}["foo5"]; v != "" || ok {
		panic("oops")
	}

	// Regression test: implicit address-taken struct literal
	// inside literal map element.
	_ = map[int]*struct{}{0: {}}
}

type mybool bool

func (mybool) f() {}

func init() {
	type mybool bool
	var b mybool
	var i interface{} = b || b // result preserves types of operands
	_ = i.(mybool)

	i = false && b // result preserves type of "typed" operand
	_ = i.(mybool)

	i = b || true // result preserves type of "typed" operand
	_ = i.(mybool)
}

func init() {
	var x, y int
	var b mybool = x == y // x==y is an untyped bool
	b.f()
}

// Simple closures.
func init() {
	b := 3
	f := func(a int) int {
		return a + b
	}
	b++
	if x := f(1); x != 5 { // 1+4 == 5
		panic(x)
	}
	b++
	if x := f(2); x != 7 { // 2+5 == 7
		panic(x)
	}
	if b := f(1) < 16 || f(2) < 17; !b {
		panic("oops")
	}
}

// Shifts.
func init() {
	var i int64 = 1
	var u uint64 = 1 << 32
	if x := i << uint32(u); x != 1 {
		panic(x)
	}
	if x := i << uint64(u); x != 0 {
		panic(x)
	}
}

// Implicit conversion of delete() key operand.
func init() {
	type I interface{}
	m := make(map[I]bool)
	m[1] = true
	m[I(2)] = true
	if len(m) != 2 {
		panic(m)
	}
	delete(m, I(1))
	delete(m, 2)
	if len(m) != 0 {
		panic(m)
	}
}

// An I->I conversion always succeeds.
func init() {
	var x I
	if I(x) != I(nil) {
		panic("I->I conversion failed")
	}
}

// An I->I type-assert fails iff the value is nil.
func init() {
	defer func() {
		r := fmt.Sprint(recover())
		// Exact error varies by toolchain.
		if r != "runtime error: interface conversion: interface is nil, not main.I" &&
			r != "interface conversion: interface is nil, not main.I" {
			panic("I->I type assertion succeeded for nil value")
		}
	}()
	var x I
	_ = x.(I)
}

//////////////////////////////////////////////////////////////////////
// Variadic bridge methods and interface thunks.

type VT int

var vcount = 0

func (VT) f(x int, y ...string) {
	vcount++
	if x != 1 {
		panic(x)
	}
	if len(y) != 2 || y[0] != "foo" || y[1] != "bar" {
		panic(y)
	}
}

type VS struct {
	VT
}

type VI interface {
	f(x int, y ...string)
}

func init() {
	foobar := []string{"foo", "bar"}
	var s VS
	s.f(1, "foo", "bar")
	s.f(1, foobar...)
	if vcount != 2 {
		panic("s.f not called twice")
	}

	fn := VI.f
	fn(s, 1, "foo", "bar")
	fn(s, 1, foobar...)
	if vcount != 4 {
		panic("I.f not called twice")
	}
}

// Multiple labels on same statement.
func multipleLabels() {
	var trace []int
	i := 0
one:
two:
	for ; i < 3; i++ {
		trace = append(trace, i)
		switch i {
		case 0:
			continue two
		case 1:
			i++
			goto one
		case 2:
			break two
		}
	}
	if x := fmt.Sprint(trace); x != "[0 1 2]" {
		panic(x)
	}
}

func init() {
	multipleLabels()
}

func init() {
	// Struct equivalence ignores blank fields.
	type s struct{ x, _, z int }
	s1 := s{x: 1, z: 3}
	s2 := s{x: 1, z: 3}
	if s1 != s2 {
		panic("not equal")
	}
}

func init() {
	// A slice var can be compared to const []T nil.
	var i interface{} = []string{"foo"}
	var j interface{} = []string(nil)
	if i.([]string) == nil {
		panic("expected i non-nil")
	}
	if j.([]string) != nil {
		panic("expected j nil")
	}
	// But two slices cannot be compared, even if one is nil.
	defer func() {
		r := fmt.Sprint(recover())
		if !(strings.Contains(r, "compar") && strings.Contains(r, "[]string")) {
			panic("want panic from slice comparison, got " + r)
		}
	}()
	_ = i == j // interface comparison recurses on types
}

func init() {
	// Regression test for SSA renaming bug.
	var ints []int
	for range "foo" {
		var x int
		x++
		ints = append(ints, x)
	}
	if fmt.Sprint(ints) != "[1 1 1]" {
		panic(ints)
	}
}

// Regression test for issue 6949:
// []byte("foo") is not a constant since it allocates memory.
func init() {
	var r string
	for i, b := range "ABC" {
		x := []byte("abc")
		x[i] = byte(b)
		r += string(x)
	}
	if r != "AbcaBcabC" {
		panic(r)
	}
}

// Test of 3-operand x[lo:hi:max] slice.
func init() {
	s := []int{0, 1, 2, 3}
	lenCapLoHi := func(x []int) [4]int { return [4]int{len(x), cap(x), x[0], x[len(x)-1]} }
	if got := lenCapLoHi(s[1:3]); got != [4]int{2, 3, 1, 2} {
		panic(got)
	}
	if got := lenCapLoHi(s[1:3:3]); got != [4]int{2, 2, 1, 2} {
		panic(got)
	}
	max := 3
	if "a"[0] == 'a' {
		max = 2 // max is non-constant, even in SSA form
	}
	if got := lenCapLoHi(s[1:2:max]); got != [4]int{1, 1, 1, 1} {
		panic(got)
	}
}

var one = 1 // not a constant

// Test makeslice.
func init() {
	check := func(s []string, wantLen, wantCap int) {
		if len(s) != wantLen {
			panic(len(s))
		}
		if cap(s) != wantCap {
			panic(cap(s))
		}
	}
	//                                       SSA form:
	check(make([]string, 10), 10, 10)     // new([10]string)[:10]
	check(make([]string, one), 1, 1)      // make([]string, one, one)
	check(make([]string, 0, 10), 0, 10)   // new([10]string)[:0]
	check(make([]string, 0, one), 0, 1)   // make([]string, 0, one)
	check(make([]string, one, 10), 1, 10) // new([10]string)[:one]
	check(make([]string, one, one), 1, 1) // make([]string, one, one)
}

// Test that a nice error is issued by indirection wrappers.
func init() {
	var ptr *T
	var i I = ptr

	defer func() {
		r := fmt.Sprint(recover())
		// Exact error varies by toolchain:
		if r != "runtime error: value method (main.T).f called using nil *main.T pointer" &&
			r != "value method (main.T).f called using nil *main.T pointer" {
			panic("want panic from call with nil receiver, got " + r)
		}
	}()
	i.f()
	panic("unreachable")
}

// Regression test for a subtle bug in which copying values would causes
// subcomponents of aggregate variables to change address, breaking
// aliases.
func init() {
	type T struct{ f int }
	var x T
	p := &x.f
	x = T{}
	*p = 1
	if x.f != 1 {
		panic("lost store")
	}
	if p != &x.f {
		panic("unstable address")
	}
}
```

## File: go/ssa/interp/testdata/deepequal.go
```go
// This interpreter test is designed to test the test copy of DeepEqual.
//
// Validate this file with 'go run' after editing.

package main

import "reflect"

func assert(cond bool) {
	if !cond {
		panic("failed")
	}
}

type X int
type Y struct {
	y *Y
	z [3]int
}

var (
	a = []int{0, 1, 2, 3}
	b = []X{0, 1, 2, 3}
	c = map[int]string{0: "zero", 1: "one"}
	d = map[X]string{0: "zero", 1: "one"}
	e = &Y{}
	f = (*Y)(nil)
	g = &Y{y: e}
	h *Y
)

func init() {
	h = &Y{} // h->h
	h.y = h
}

func main() {
	assert(reflect.DeepEqual(nil, nil))
	assert(reflect.DeepEqual((*int)(nil), (*int)(nil)))
	assert(!reflect.DeepEqual(nil, (*int)(nil)))

	assert(reflect.DeepEqual(0, 0))
	assert(!reflect.DeepEqual(0, int64(0)))

	assert(!reflect.DeepEqual("", 0))

	assert(reflect.DeepEqual(a, []int{0, 1, 2, 3}))
	assert(!reflect.DeepEqual(a, []int{0, 1, 2}))
	assert(!reflect.DeepEqual(a, []int{0, 1, 0, 3}))

	assert(reflect.DeepEqual(b, []X{0, 1, 2, 3}))
	assert(!reflect.DeepEqual(b, []X{0, 1, 0, 3}))

	assert(reflect.DeepEqual(c, map[int]string{0: "zero", 1: "one"}))
	assert(!reflect.DeepEqual(c, map[int]string{0: "zero", 1: "one", 2: "two"}))
	assert(!reflect.DeepEqual(c, map[int]string{1: "one", 2: "two"}))
	assert(!reflect.DeepEqual(c, map[int]string{1: "one"}))

	assert(reflect.DeepEqual(d, map[X]string{0: "zero", 1: "one"}))
	assert(!reflect.DeepEqual(d, map[int]string{0: "zero", 1: "one"}))

	assert(reflect.DeepEqual(e, &Y{}))
	assert(reflect.DeepEqual(e, &Y{z: [3]int{0, 0, 0}}))
	assert(!reflect.DeepEqual(e, &Y{z: [3]int{0, 1, 0}}))

	assert(reflect.DeepEqual(f, (*Y)(nil)))
	assert(!reflect.DeepEqual(f, nil))

	// eq_h -> eq_h. Pointer structure and elements are equal so DeepEqual.
	eq_h := &Y{}
	eq_h.y = eq_h
	assert(reflect.DeepEqual(h, eq_h))

	// deepeq_h->h->h. Pointed to elem of (deepeq_h, h) are (h,h). (h,h) are deep equal so h and deepeq_h are DeepEqual.
	deepeq_h := &Y{}
	deepeq_h.y = h
	assert(reflect.DeepEqual(h, deepeq_h))

	distinct := []interface{}{a, b, c, d, e, f, g, h}
	for x := range distinct {
		for y := range distinct {
			assert((x == y) == reflect.DeepEqual(distinct[x], distinct[y]))
		}
	}

	// anonymous struct types.
	assert(reflect.DeepEqual(struct{}{}, struct{}{}))
	assert(reflect.DeepEqual(struct{ x int }{1}, struct{ x int }{1}))
	assert(!reflect.DeepEqual(struct{ x int }{}, struct{ x int }{5}))
	assert(!reflect.DeepEqual(struct{ x, y int }{0, 1}, struct{ x int }{0}))
	assert(reflect.DeepEqual(struct{ x, y int }{2, 3}, struct{ x, y int }{2, 3}))
	assert(!reflect.DeepEqual(struct{ x, y int }{4, 5}, struct{ x, y int }{4, 6}))
}
```

## File: go/ssa/interp/testdata/defer.go
```go
package main

// Tests of defer.  (Deferred recover() belongs is recover.go.)

import "fmt"

func deferMutatesResults(noArgReturn bool) (a, b int) {
	defer func() {
		if a != 1 || b != 2 {
			panic(fmt.Sprint(a, b))
		}
		a, b = 3, 4
	}()
	if noArgReturn {
		a, b = 1, 2
		return
	}
	return 1, 2
}

func init() {
	a, b := deferMutatesResults(true)
	if a != 3 || b != 4 {
		panic(fmt.Sprint(a, b))
	}
	a, b = deferMutatesResults(false)
	if a != 3 || b != 4 {
		panic(fmt.Sprint(a, b))
	}
}

// We concatenate init blocks to make a single function, but we must
// run defers at the end of each block, not the combined function.
var deferCount = 0

func init() {
	deferCount = 1
	defer func() {
		deferCount++
	}()
	// defer runs HERE
}

func init() {
	// Strictly speaking the spec says deferCount may be 0 or 2
	// since the relative order of init blocks is unspecified.
	if deferCount != 2 {
		panic(deferCount) // defer call has not run!
	}
}

func main() {
}
```

## File: go/ssa/interp/testdata/fieldprom.go
```go
package main

// Tests of field promotion logic.

type A struct {
	x int
	y *int
}

type B struct {
	p int
	q *int
}

type C struct {
	A
	*B
}

type D struct {
	a int
	C
}

func assert(cond bool) {
	if !cond {
		panic("failed")
	}
}

func f1(c C) {
	assert(c.x == c.A.x)
	assert(c.y == c.A.y)
	assert(&c.x == &c.A.x)
	assert(&c.y == &c.A.y)

	assert(c.p == c.B.p)
	assert(c.q == c.B.q)
	assert(&c.p == &c.B.p)
	assert(&c.q == &c.B.q)

	c.x = 1
	*c.y = 1
	c.p = 1
	*c.q = 1
}

func f2(c *C) {
	assert(c.x == c.A.x)
	assert(c.y == c.A.y)
	assert(&c.x == &c.A.x)
	assert(&c.y == &c.A.y)

	assert(c.p == c.B.p)
	assert(c.q == c.B.q)
	assert(&c.p == &c.B.p)
	assert(&c.q == &c.B.q)

	c.x = 1
	*c.y = 1
	c.p = 1
	*c.q = 1
}

func f3(d D) {
	assert(d.x == d.C.A.x)
	assert(d.y == d.C.A.y)
	assert(&d.x == &d.C.A.x)
	assert(&d.y == &d.C.A.y)

	assert(d.p == d.C.B.p)
	assert(d.q == d.C.B.q)
	assert(&d.p == &d.C.B.p)
	assert(&d.q == &d.C.B.q)

	d.x = 1
	*d.y = 1
	d.p = 1
	*d.q = 1
}

func f4(d *D) {
	assert(d.x == d.C.A.x)
	assert(d.y == d.C.A.y)
	assert(&d.x == &d.C.A.x)
	assert(&d.y == &d.C.A.y)

	assert(d.p == d.C.B.p)
	assert(d.q == d.C.B.q)
	assert(&d.p == &d.C.B.p)
	assert(&d.q == &d.C.B.q)

	d.x = 1
	*d.y = 1
	d.p = 1
	*d.q = 1
}

func main() {
	y := 123
	c := C{
		A{x: 42, y: &y},
		&B{p: 42, q: &y},
	}

	assert(&c.x == &c.A.x)

	f1(c)
	f2(&c)

	d := D{C: c}
	f3(d)
	f4(&d)
}
```

## File: go/ssa/interp/testdata/forvarlifetime_go122.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
)

func main() {
	test_init()
	bound()
	manyvars()
	nocond()
	nopost()
	address_sequences()
	post_escapes()

	// Clones from cmd/compile/internal/loopvar/testdata .
	for_complicated_esc_address()
	for_esc_address()
	for_esc_closure()
	for_esc_method()
}

// After go1.22, each i will have a distinct address and value.
var distinct = func(m, n int) []*int {
	var r []*int
	for i := m; i <= n; i++ {
		r = append(r, &i)
	}
	return r
}(3, 5)

func test_init() {
	if len(distinct) != 3 {
		panic(distinct)
	}
	for i, v := range []int{3, 4, 5} {
		if v != *(distinct[i]) {
			panic(distinct)
		}
	}
}

func bound() {
	b := func(k int) func() int {
		var f func() int
		for i := 0; i < k; i++ {
			f = func() int { return i } // address before post updates i. So last value in the body.
		}
		return f
	}

	if got := b(0); got != nil {
		panic(got)
	}
	if got := b(5); got() != 4 {
		panic(got())
	}
}

func manyvars() {
	// Tests declaring many variables and having one in the middle escape.
	var f func() int
	for i, j, k, l, m, n, o, p := 7, 6, 5, 4, 3, 2, 1, 0; p < 6; l, p = l+1, p+1 {
		_, _, _, _, _, _, _, _ = i, j, k, l, m, n, o, p
		f = func() int { return l } // address *before* post updates l
	}
	if f() != 9 { // l == p+4
		panic(f())
	}
}

func nocond() {
	var c, b, e *int
	for p := 0; ; p++ {
		if p%7 == 0 {
			c = &p
			continue
		} else if p == 20 {
			b = &p
			break
		}
		e = &p
	}

	if *c != 14 {
		panic(c)
	}
	if *b != 20 {
		panic(b)
	}
	if *e != 19 {
		panic(e)
	}
}

func nopost() {
	var first, last *int
	for p := 0; p < 20; {
		if first == nil {
			first = &p
		}
		last = &p

		p++
	}

	if *first != 1 {
		panic(first)
	}
	if *last != 20 {
		panic(last)
	}
}

func address_sequences() {
	var c, b, p []*int

	cond := func(x *int) bool {
		c = append(c, x)
		return *x < 5
	}
	body := func(x *int) {
		b = append(b, x)
	}
	post := func(x *int) {
		p = append(p, x)
		(*x)++
	}
	for i := 0; cond(&i); post(&i) {
		body(&i)
	}

	if c[0] == c[1] {
		panic(c)
	}

	if !reflect.DeepEqual(c[:5], b) {
		panic(c)
	}

	if !reflect.DeepEqual(c[1:], p) {
		panic(c)
	}

	if !reflect.DeepEqual(b[1:], p[:4]) {
		panic(b)
	}
}

func post_escapes() {
	var p []*int
	post := func(x *int) {
		p = append(p, x)
		(*x)++
	}

	for i := 0; i < 5; post(&i) {
	}

	var got []int
	for _, x := range p {
		got = append(got, *x)
	}
	if want := []int{1, 2, 3, 4, 5}; !reflect.DeepEqual(got, want) {
		panic(got)
	}
}

func for_complicated_esc_address() {
	// Clone of for_complicated_esc_adress.go
	ss, sa := shared(23)
	ps, pa := private(23)
	es, ea := experiment(23)

	if ss != ps || ss != es || ea != pa || sa == pa {
		println("shared s, a", ss, sa, "; private, s, a", ps, pa, "; experiment s, a", es, ea)
		panic("for_complicated_esc_address")
	}
}

func experiment(x int) (int, int) {
	sum := 0
	var is []*int
	for i := x; i != 1; i = i / 2 {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}

	return sum, asum
}

func private(x int) (int, int) {
	sum := 0
	var is []*int
	I := x
	for ; I != 1; I = I / 2 {
		i := I
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				I = i
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
		I = i
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}

	return sum, asum
}

func shared(x int) (int, int) {
	sum := 0
	var is []*int
	i := x
	for ; i != 1; i = i / 2 {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}
	return sum, asum
}

func for_esc_address() {
	// Clone of for_esc_address.go
	sum := 0
	var is []*int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, &i)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, pi := range is {
		sum += *pi
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected ", 20, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_address")
	}
}

func for_esc_closure() {
	var is []func() int

	// Clone of for_esc_closure.go
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected ", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected ", 20, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_closure")
	}
}

type I int

func (x *I) method() int {
	return int(*x)
}

func for_esc_method() {
	// Clone of for_esc_method.go
	var is []func() int
	sum := 0
	for i := I(0); int(i) < 10; i++ {
		for j := 0; j < 10; j++ {
			if int(i) == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, i.method)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected ", 90, ", saw ", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, m := range is {
		sum += m()
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected ", 20, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_method")
	}
}
```

## File: go/ssa/interp/testdata/forvarlifetime_old.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21

// goversion can be pinned to anything strictly before 1.22.

package main

import (
	"reflect"
)

func main() {
	test_init()
	bound()
	manyvars()
	nocond()
	nopost()
	address_sequences()
	post_escapes()

	// Clones from cmd/compile/internal/loopvar/testdata .
	for_complicated_esc_address()
	for_esc_address()
	for_esc_closure()
	for_esc_method()
}

// pre-go1.22 all of i will have the same address and the value of 6.
var same = func(m, n int) []*int {
	var r []*int
	for i := m; i <= n; i++ {
		r = append(r, &i)
	}
	return r
}(3, 5)

func test_init() {
	if len(same) != 3 {
		panic(same)
	}
	for i := range same {
		for j := range same {
			if !(same[i] == same[j]) {
				panic(same)
			}
		}
	}
	for i := range same {
		if *(same[i]) != 6 {
			panic(same)
		}
	}
}

func bound() {
	b := func(k int) func() int {
		var f func() int
		for i := 0; i < k; i++ {
			f = func() int { return i } // shared address will equal k.
		}
		return f
	}

	if got := b(0); got != nil {
		panic(got)
	}
	if got := b(5); got() != 5 {
		panic(got())
	}
}

func manyvars() {
	// Tests declaring many variables and having one in the middle escape.
	var f func() int
	for i, j, k, l, m, n, o, p := 7, 6, 5, 4, 3, 2, 1, 0; p < 6; l, p = l+1, p+1 {
		_, _, _, _, _, _, _, _ = i, j, k, l, m, n, o, p
		f = func() int { return l } // address *after* post updates l
	}
	if f() != 10 { // l == p+4
		panic(f())
	}
}

func nocond() {
	var c, b, e *int
	for p := 0; ; p++ {
		if p%7 == 0 {
			c = &p
			continue
		} else if p == 20 {
			b = &p
			break
		}
		e = &p
	}

	if *c != 20 {
		panic(c)
	}
	if *b != 20 {
		panic(b)
	}
	if *e != 20 {
		panic(e)
	}
}

func nopost() {
	var first, last *int
	for p := 0; p < 20; {
		if first == nil {
			first = &p
		}
		last = &p

		p++
	}

	if *first != 20 {
		panic(first)
	}
	if *last != 20 {
		panic(last)
	}
}

func address_sequences() {
	var c, b, p []*int

	cond := func(x *int) bool {
		c = append(c, x)
		return *x < 5
	}
	body := func(x *int) {
		b = append(b, x)
	}
	post := func(x *int) {
		p = append(p, x)
		(*x)++
	}
	for i := 0; cond(&i); post(&i) {
		body(&i)
	}

	if c[0] != c[1] {
		panic(c)
	}

	if !reflect.DeepEqual(c[:5], b) {
		panic(c)
	}

	if !reflect.DeepEqual(c[1:], p) {
		panic(c)
	}

	if !reflect.DeepEqual(b[1:], p[:4]) {
		panic(b)
	}
}

func post_escapes() {
	var p []*int
	post := func(x *int) {
		p = append(p, x)
		(*x)++
	}

	for i := 0; i < 5; post(&i) {
	}

	var got []int
	for _, x := range p {
		got = append(got, *x)
	}
	if want := []int{5, 5, 5, 5, 5}; !reflect.DeepEqual(got, want) {
		panic(got)
	}
}

func for_complicated_esc_address() {
	// Clone of for_complicated_esc_adress.go
	ss, sa := shared(23)
	ps, pa := private(23)
	es, ea := experiment(23)

	if ss != ps || ss != es || sa != ea || pa != 188 {
		println("shared s, a", ss, sa, "; private, s, a", ps, pa, "; experiment s, a", es, ea)
		panic("for_complicated_esc_address")
	}
}

func experiment(x int) (int, int) {
	sum := 0
	var is []*int
	for i := x; i != 1; i = i / 2 {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}

	return sum, asum
}

func private(x int) (int, int) {
	sum := 0
	var is []*int
	I := x
	for ; I != 1; I = I / 2 {
		i := I
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				I = i
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
		I = i
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}

	return sum, asum
}

func shared(x int) (int, int) {
	sum := 0
	var is []*int
	i := x
	for ; i != 1; i = i / 2 {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		i = i*3 + 1
		if i&1 == 0 {
			is = append(is, &i)
			for i&2 == 0 {
				i = i >> 1
			}
		} else {
			i = i + i
		}
	}

	asum := 0
	for _, pi := range is {
		asum += *pi
	}
	return sum, asum
}

func for_esc_address() {
	// Clone of for_esc_address.go
	sum := 0
	var is []*int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, &i)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, pi := range is {
		sum += *pi
	}
	if sum != 10+10+10+10+10 {
		println("wrong sum, expected ", 50, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_address")
	}
}

func for_esc_closure() {
	// Clone of for_esc_closure.go
	var is []func() int
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected ", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 10+10+10+10+10 {
		println("wrong sum, expected ", 50, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_closure")
	}
}

type I int

func (x *I) method() int {
	return int(*x)
}

func for_esc_method() {
	// Clone of for_esc_method.go
	sum := 0
	var is []func() int
	for i := I(0); int(i) < 10; i++ {
		for j := 0; j < 10; j++ {
			if int(i) == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, i.method)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected ", 90, ", saw ", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, m := range is {
		sum += m()
	}
	if sum != 10+10+10+10+10 {
		println("wrong sum, expected ", 50, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("for_esc_method")
	}
}
```

## File: go/ssa/interp/testdata/ifaceconv.go
```go
package main

// Tests of interface conversions and type assertions.

type I0 interface {
}
type I1 interface {
	f()
}
type I2 interface {
	f()
	g()
}

type C0 struct{}
type C1 struct{}

func (C1) f() {}

type C2 struct{}

func (C2) f() {}
func (C2) g() {}

func main() {
	var i0 I0
	var i1 I1
	var i2 I2

	// Nil always causes a type assertion to fail, even to the
	// same type.
	if _, ok := i0.(I0); ok {
		panic("nil i0.(I0) succeeded")
	}
	if _, ok := i1.(I1); ok {
		panic("nil i1.(I1) succeeded")
	}
	if _, ok := i2.(I2); ok {
		panic("nil i2.(I2) succeeded")
	}

	// Conversions can't fail, even with nil.
	_ = I0(i0)

	_ = I0(i1)
	_ = I1(i1)

	_ = I0(i2)
	_ = I1(i2)
	_ = I2(i2)

	// Non-nil type assertions pass or fail based on the concrete type.
	i1 = C1{}
	if _, ok := i1.(I0); !ok {
		panic("C1 i1.(I0) failed")
	}
	if _, ok := i1.(I1); !ok {
		panic("C1 i1.(I1) failed")
	}
	if _, ok := i1.(I2); ok {
		panic("C1 i1.(I2) succeeded")
	}

	i1 = C2{}
	if _, ok := i1.(I0); !ok {
		panic("C2 i1.(I0) failed")
	}
	if _, ok := i1.(I1); !ok {
		panic("C2 i1.(I1) failed")
	}
	if _, ok := i1.(I2); !ok {
		panic("C2 i1.(I2) failed")
	}

	// Conversions can't fail.
	i1 = C1{}
	if I0(i1) == nil {
		panic("C1 I0(i1) was nil")
	}
	if I1(i1) == nil {
		panic("C1 I1(i1) was nil")
	}
}
```

## File: go/ssa/interp/testdata/ifaceprom.go
```go
package main

// Test of promotion of methods of an interface embedded within a
// struct.  In particular, this test exercises that the correct
// method is called.

type I interface {
	one() int
	two() string
}

type S struct {
	I
}

type impl struct{}

func (impl) one() int {
	return 1
}

func (impl) two() string {
	return "two"
}

func main() {
	var s S
	s.I = impl{}
	if one := s.I.one(); one != 1 {
		panic(one)
	}
	if one := s.one(); one != 1 {
		panic(one)
	}
	closOne := s.I.one
	if one := closOne(); one != 1 {
		panic(one)
	}
	closOne = s.one
	if one := closOne(); one != 1 {
		panic(one)
	}

	if two := s.I.two(); two != "two" {
		panic(two)
	}
	if two := s.two(); two != "two" {
		panic(two)
	}
	closTwo := s.I.two
	if two := closTwo(); two != "two" {
		panic(two)
	}
	closTwo = s.two
	if two := closTwo(); two != "two" {
		panic(two)
	}
}
```

## File: go/ssa/interp/testdata/initorder.go
```go
package main

import "fmt"

// Test of initialization order of package-level vars.

var counter int

func next() int {
	c := counter
	counter++
	return c
}

func next2() (x int, y int) {
	x = next()
	y = next()
	return
}

func makeOrder() int {
	_, _, _, _ = f, b, d, e
	return 0
}

func main() {
	// Initialization constraints:
	// - {f,b,c/d,e} < order  (ref graph traversal)
	// - order < {a}          (lexical order)
	// - b < c/d < e < f      (lexical order)
	// Solution: a b c/d e f
	abcdef := [6]int{a, b, c, d, e, f}
	if abcdef != [6]int{0, 1, 2, 3, 4, 5} {
		panic(abcdef)
	}

	// Initializers of even blank globals are evaluated.
	if g != 1 {
		panic(g)
	}
}

var order = makeOrder()

var a, b = next(), next()
var c, d = next2()
var e, f = next(), next()

var (
	g int
	_ = func() int { g = 1; return 0 }()
)

// ------------------------------------------------------------------------

var order2 []string

func create(x int, name string) int {
	order2 = append(order2, name)
	return x
}

var C = create(B+1, "C")
var A, B = create(1, "A"), create(2, "B")

// Initialization order of package-level value specs.
func init() {
	x := fmt.Sprint(order2)
	// Result varies by toolchain.  This is a spec bug.
	if x != "[B C A]" && // gc
		x != "[A B C]" { // go/types
		panic(x)
	}
	if C != 3 {
		panic(c)
	}
}
```

## File: go/ssa/interp/testdata/methprom.go
```go
package main

// Tests of method promotion logic.

type A struct{ magic int }

func (a A) x() {
	if a.magic != 1 {
		panic(a.magic)
	}
}
func (a *A) y() *A {
	return a
}

type B struct{ magic int }

func (b B) p() {
	if b.magic != 2 {
		panic(b.magic)
	}
}
func (b *B) q() {
	if b != theC.B {
		panic("oops")
	}
}

type I interface {
	f()
}

type impl struct{ magic int }

func (i impl) f() {
	if i.magic != 3 {
		panic("oops")
	}
}

type C struct {
	A
	*B
	I
}

func assert(cond bool) {
	if !cond {
		panic("failed")
	}
}

var theC = C{
	A: A{1},
	B: &B{2},
	I: impl{3},
}

func addr() *C {
	return &theC
}

func value() C {
	return theC
}

func main() {
	// address
	addr().x()
	if addr().y() != &theC.A {
		panic("oops")
	}
	addr().p()
	addr().q()
	addr().f()

	// addressable value
	var c C = value()
	c.x()
	if c.y() != &c.A {
		panic("oops")
	}
	c.p()
	c.q()
	c.f()

	// non-addressable value
	value().x()
	// value().y() // not in method set
	value().p()
	value().q()
	value().f()
}
```

## File: go/ssa/interp/testdata/minmax.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
)

func main() {
	TestMinFloat()
	TestMaxFloat()
	TestMinMaxInt()
	TestMinMaxUint8()
	TestMinMaxString()
}

func errorf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
func fatalf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }

// derived from $GOROOT/src/runtime/minmax_test.go

var (
	zero    = math.Copysign(0, +1)
	negZero = math.Copysign(0, -1)
	inf     = math.Inf(+1)
	negInf  = math.Inf(-1)
	nan     = math.NaN()
)

var tests = []struct{ min, max float64 }{
	{1, 2},
	{-2, 1},
	{negZero, zero},
	{zero, inf},
	{negInf, zero},
	{negInf, inf},
	{1, inf},
	{negInf, 1},
}

var all = []float64{1, 2, -1, -2, zero, negZero, inf, negInf, nan}

func eq(x, y float64) bool {
	return x == y && math.Signbit(x) == math.Signbit(y)
}

func TestMinFloat() {
	for _, tt := range tests {
		if z := min(tt.min, tt.max); !eq(z, tt.min) {
			errorf("min(%v, %v) = %v, want %v", tt.min, tt.max, z, tt.min)
		}
		if z := min(tt.max, tt.min); !eq(z, tt.min) {
			errorf("min(%v, %v) = %v, want %v", tt.max, tt.min, z, tt.min)
		}
	}
	for _, x := range all {
		if z := min(nan, x); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
		if z := min(x, nan); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
	}
}

func TestMaxFloat() {
	for _, tt := range tests {
		if z := max(tt.min, tt.max); !eq(z, tt.max) {
			errorf("max(%v, %v) = %v, want %v", tt.min, tt.max, z, tt.max)
		}
		if z := max(tt.max, tt.min); !eq(z, tt.max) {
			errorf("max(%v, %v) = %v, want %v", tt.max, tt.min, z, tt.max)
		}
	}
	for _, x := range all {
		if z := max(nan, x); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
		if z := max(x, nan); !math.IsNaN(z) {
			errorf("min(%v, %v) = %v, want %v", nan, x, z, nan)
		}
	}
}

// testMinMax tests that min/max behave correctly on every pair of
// values in vals.
//
// vals should be a sequence of values in strictly ascending order.
func testMinMax[T int | uint8 | string](vals ...T) {
	for i, x := range vals {
		for _, y := range vals[i+1:] {
			if !(x < y) {
				fatalf("values out of order: !(%v < %v)", x, y)
			}

			if z := min(x, y); z != x {
				errorf("min(%v, %v) = %v, want %v", x, y, z, x)
			}
			if z := min(y, x); z != x {
				errorf("min(%v, %v) = %v, want %v", y, x, z, x)
			}

			if z := max(x, y); z != y {
				errorf("max(%v, %v) = %v, want %v", x, y, z, y)
			}
			if z := max(y, x); z != y {
				errorf("max(%v, %v) = %v, want %v", y, x, z, y)
			}
		}
	}
}

func TestMinMaxInt()    { testMinMax[int](-7, 0, 9) }
func TestMinMaxUint8()  { testMinMax[uint8](0, 1, 2, 4, 7) }
func TestMinMaxString() { testMinMax[string]("a", "b", "c") }
```

## File: go/ssa/interp/testdata/mrvchain.go
```go
// Tests of call chaining f(g()) when g has multiple return values (MRVs).
// See https://code.google.com/p/go/issues/detail?id=4573.

package main

func assert(actual, expected int) {
	if actual != expected {
		panic(actual)
	}
}

func g() (int, int) {
	return 5, 7
}

func g2() (float64, float64) {
	return 5, 7
}

func f1v(x int, v ...int) {
	assert(x, 5)
	assert(v[0], 7)
}

func f2(x, y int) {
	assert(x, 5)
	assert(y, 7)
}

func f2v(x, y int, v ...int) {
	assert(x, 5)
	assert(y, 7)
	assert(len(v), 0)
}

func complexArgs() (float64, float64) {
	return 5, 7
}

func appendArgs() ([]string, string) {
	return []string{"foo"}, "bar"
}

func h() (i interface{}, ok bool) {
	m := map[int]string{1: "hi"}
	i, ok = m[1] // string->interface{} conversion within multi-valued expression
	return
}

func h2() (i interface{}, ok bool) {
	ch := make(chan string, 1)
	ch <- "hi"
	i, ok = <-ch // string->interface{} conversion within multi-valued expression
	return
}

func main() {
	f1v(g())
	f2(g())
	f2v(g())
	if c := complex(complexArgs()); c != 5+7i {
		panic(c)
	}
	if s := append(appendArgs()); len(s) != 2 || s[0] != "foo" || s[1] != "bar" {
		panic(s)
	}
	i, ok := h()
	if !ok || i.(string) != "hi" {
		panic(i)
	}
	i, ok = h2()
	if !ok || i.(string) != "hi" {
		panic(i)
	}
}
```

## File: go/ssa/interp/testdata/range.go
```go
package main

// Tests of range loops.

import "fmt"

// Range over string.
func init() {
	if x := len("Hello, 世界"); x != 13 { // bytes
		panic(x)
	}
	var indices []int
	var runes []rune
	for i, r := range "Hello, 世界" {
		runes = append(runes, r)
		indices = append(indices, i)
	}
	if x := fmt.Sprint(runes); x != "[72 101 108 108 111 44 32 19990 30028]" {
		panic(x)
	}
	if x := fmt.Sprint(indices); x != "[0 1 2 3 4 5 6 7 10]" {
		panic(x)
	}
	s := ""
	for _, r := range runes {
		s += string(r)
	}
	if s != "Hello, 世界" {
		panic(s)
	}

	var x int
	for range "Hello, 世界" {
		x++
	}
	if x != len(indices) {
		panic(x)
	}
}

// Regression test for range of pointer to named array type.
func init() {
	type intarr [3]int
	ia := intarr{1, 2, 3}
	var count int
	for _, x := range &ia {
		count += x
	}
	if count != 6 {
		panic(count)
	}
}

func main() {
}
```

## File: go/ssa/interp/testdata/rangefunc.go
```go
// Range over functions.

// Currently requires 1.22 and GOEXPERIMENT=rangefunc.

// Fork of src/cmd/compile/internal/rangefunc/rangefunc_test.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	TestCheck("TestCheck")
	TestCooperativeBadOfSliceIndex("TestCooperativeBadOfSliceIndex")
	TestCooperativeBadOfSliceIndexCheck("TestCooperativeBadOfSliceIndexCheck")
	TestTrickyIterAll("TestTrickyIterAll")
	TestTrickyIterOne("TestTrickyIterOne")
	TestTrickyIterZero("TestTrickyIterZero")
	TestTrickyIterZeroCheck("TestTrickyIterZeroCheck")
	TestTrickyIterEcho("TestTrickyIterEcho")
	TestTrickyIterEcho2("TestTrickyIterEcho2")
	TestBreak1("TestBreak1")
	TestBreak2("TestBreak2")
	TestContinue("TestContinue")
	TestBreak3("TestBreak3")
	TestBreak1BadA("TestBreak1BadA")
	TestBreak1BadB("TestBreak1BadB")
	TestMultiCont0("TestMultiCont0")
	TestMultiCont1("TestMultiCont1")
	TestMultiCont2("TestMultiCont2")
	TestMultiCont3("TestMultiCont3")
	TestMultiBreak0("TestMultiBreak0")
	TestMultiBreak1("TestMultiBreak1")
	TestMultiBreak2("TestMultiBreak2")
	TestMultiBreak3("TestMultiBreak3")
	TestPanickyIterator1("TestPanickyIterator1")
	TestPanickyIterator1Check("TestPanickyIterator1Check")
	TestPanickyIterator2("TestPanickyIterator2")
	TestPanickyIterator2Check("TestPanickyIterator2Check")
	TestPanickyIterator3("TestPanickyIterator3")
	TestPanickyIterator3Check("TestPanickyIterator3Check")
	TestPanickyIterator4("TestPanickyIterator4")
	TestPanickyIterator4Check("TestPanickyIterator4Check")
	TestVeryBad1("TestVeryBad1")
	TestVeryBad2("TestVeryBad2")
	TestVeryBadCheck("TestVeryBadCheck")
	TestOk("TestOk")
	TestBreak1BadDefer("TestBreak1BadDefer")
	TestReturns("TestReturns")
	TestGotoA("TestGotoA")
	TestGotoB("TestGotoB")
	TestPanicReturns("TestPanicReturns")
}

type testingT string

func (t testingT) Log(args ...any) {
	s := fmt.Sprint(args...)
	println(t, "\t", s)
}

func (t testingT) Error(args ...any) {
	s := string(t) + "\terror: " + fmt.Sprint(args...)
	panic(s)
}

// slicesEqual is a clone of slices.Equal
func slicesEqual[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

type Seq[T any] func(yield func(T) bool)
type Seq2[T1, T2 any] func(yield func(T1, T2) bool)

// OfSliceIndex returns a Seq2 over the elements of s. It is equivalent
// to range s.
func OfSliceIndex[T any, S ~[]T](s S) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
		return
	}
}

// BadOfSliceIndex is "bad" because it ignores the return value from yield
// and just keeps on iterating.
func BadOfSliceIndex[T any, S ~[]T](s S) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			yield(i, v)
		}
		return
	}
}

// VeryBadOfSliceIndex is "very bad" because it ignores the return value from yield
// and just keeps on iterating, and also wraps that call in a defer-recover so it can
// keep on trying after the first panic.
func VeryBadOfSliceIndex[T any, S ~[]T](s S) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			func() {
				defer func() {
					recover()
				}()
				yield(i, v)
			}()
		}
		return
	}
}

// SwallowPanicOfSliceIndex hides panics and converts them to normal return
func SwallowPanicOfSliceIndex[T any, S ~[]T](s S) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			done := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						done = true
					}
				}()
				done = !yield(i, v)
			}()
			if done {
				return
			}
		}
		return
	}
}

// PanickyOfSliceIndex iterates the slice but panics if it exits the loop early
func PanickyOfSliceIndex[T any, S ~[]T](s S) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				panic("Panicky iterator panicking")
			}
		}
		return
	}
}

// CooperativeBadOfSliceIndex calls the loop body from a goroutine after
// a ping on a channel, and returns recover()on that same channel.
func CooperativeBadOfSliceIndex[T any, S ~[]T](s S, proceed chan any) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				// if the body breaks, call yield just once in a goroutine
				go func() {
					<-proceed
					defer func() {
						proceed <- recover()
					}()
					yield(0, s[0])
				}()
				return
			}
		}
		return
	}
}

// TrickyIterator is a type intended to test whether an iterator that
// calls a yield function after loop exit must inevitably escape the
// closure; this might be relevant to future checking/optimization.
type TrickyIterator struct {
	yield func(int, int) bool
}

func (ti *TrickyIterator) iterEcho(s []int) Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, v := range s {
			if !yield(i, v) {
				ti.yield = yield
				return
			}
			if ti.yield != nil && !ti.yield(i, v) {
				return
			}
		}
		ti.yield = yield
		return
	}
}

func (ti *TrickyIterator) iterAll(s []int) Seq2[int, int] {
	return func(yield func(int, int) bool) {
		ti.yield = yield // Save yield for future abuse
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
		return
	}
}
func (ti *TrickyIterator) iterOne(s []int) Seq2[int, int] {
	return func(yield func(int, int) bool) {
		ti.yield = yield // Save yield for future abuse
		if len(s) > 0 {  // Not in a loop might escape differently
			yield(0, s[0])
		}
		return
	}
}
func (ti *TrickyIterator) iterZero(s []int) Seq2[int, int] {
	return func(yield func(int, int) bool) {
		ti.yield = yield // Save yield for future abuse
		// Don't call it at all, maybe it won't escape
		return
	}
}
func (ti *TrickyIterator) fail() {
	if ti.yield != nil {
		ti.yield(1, 1)
	}
}

func matchError(r any, x string) bool {
	if r == nil {
		return false
	}
	if x == "" {
		return true
	}
	switch p := r.(type) {
	case string:
		return p == x
	case errorString:
		return p.Error() == x
	case error:
		return strings.Contains(p.Error(), x)
	}
	return false
}

func matchErrorHelper(t testingT, r any, x string) {
	if matchError(r, x) {
		t.Log("Saw expected panic: ", r)
	} else {
		t.Error("Saw wrong panic: '", r, "' . expected '", x, "'")
	}
}

const DONE = 0          // body of loop has exited in a non-panic way
const READY = 1         // body of loop has not exited yet, is not running
const PANIC = 2         // body of loop is either currently running, or has panicked
const EXHAUSTED = 3     // iterator function return, i.e., sequence is "exhausted"
const MISSING_PANIC = 4 // overload "READY" for panic call

// An errorString represents a runtime error described by a single string.
type errorString string

func (e errorString) Error() string {
	return string(e)
}

const (
	// RERR_ is for runtime error, and may be regexps/substrings, to simplify use of tests with tools
	RERR_DONE      = "yield function called after range loop exit"
	RERR_PANIC     = "range function continued iteration after loop body panic"
	RERR_EXHAUSTED = "yield function called after range loop exit" // ssa does not distinguish DONE and EXHAUSTED
	RERR_MISSING   = "iterator call did not preserve panic"

	// CERR_ is for checked errors in the Check combinator defined above, and should be literal strings
	CERR_PFX       = "checked rangefunc error: "
	CERR_DONE      = CERR_PFX + "loop iteration after body done"
	CERR_PANIC     = CERR_PFX + "loop iteration after panic"
	CERR_EXHAUSTED = CERR_PFX + "loop iteration after iterator exit"
	CERR_MISSING   = CERR_PFX + "loop iterator swallowed panic"
)

var fail []error = []error{
	errorString(CERR_DONE),
	errorString(CERR_PFX + "loop iterator, unexpected error"),
	errorString(CERR_PANIC),
	errorString(CERR_EXHAUSTED),
	errorString(CERR_MISSING),
}

// Check wraps the function body passed to iterator forall
// in code that ensures that it cannot (successfully) be called
// either after body return false (control flow out of loop) or
// forall itself returns (the iteration is now done).
//
// Note that this can catch errors before the inserted checks.
func Check[U, V any](forall Seq2[U, V]) Seq2[U, V] {
	return func(body func(U, V) bool) {
		state := READY
		forall(func(u U, v V) bool {
			if state != READY {
				panic(fail[state])
			}
			state = PANIC
			ret := body(u, v)
			if ret {
				state = READY
			} else {
				state = DONE
			}
			return ret
		})
		if state == PANIC {
			panic(fail[MISSING_PANIC])
		}
		state = EXHAUSTED
	}
}

func TestCheck(t testingT) {
	i := 0
	defer func() {
		t.Log("i = ", i) // 45
		matchErrorHelper(t, recover(), CERR_DONE)
	}()
	for _, x := range Check(BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) {
		i += x
		if i > 4*9 {
			break
		}
	}
}

func TestCooperativeBadOfSliceIndex(t testingT) {
	i := 0
	proceed := make(chan any)
	for _, x := range CooperativeBadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, proceed) {
		i += x
		if i >= 36 {
			break
		}
	}
	proceed <- true
	r := <-proceed
	matchErrorHelper(t, r, RERR_EXHAUSTED)
	if i != 36 {
		t.Error("Expected i == 36, saw ", i, "instead")
	} else {
		t.Log("i = ", i)
	}
}

func TestCooperativeBadOfSliceIndexCheck(t testingT) {
	i := 0
	proceed := make(chan any)
	for _, x := range Check(CooperativeBadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, proceed)) {
		i += x
		if i >= 36 {
			break
		}
	}
	proceed <- true
	r := <-proceed
	matchErrorHelper(t, r, CERR_EXHAUSTED)

	if i != 36 {
		t.Error("Expected i == 36, saw ", i, "instead")
	} else {
		t.Log("i = ", i)
	}
}

func TestTrickyIterAll(t testingT) {
	trickItAll := TrickyIterator{}
	i := 0
	for _, x := range trickItAll.iterAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		i += x
		if i >= 36 {
			break
		}
	}
	if i != 36 {
		t.Error("Expected i == 36, saw ", i, " instead")
	} else {
		t.Log("i = ", i)
	}
	defer func() {
		matchErrorHelper(t, recover(), RERR_EXHAUSTED)
	}()
	trickItAll.fail()
}

func TestTrickyIterOne(t testingT) {
	trickItOne := TrickyIterator{}
	i := 0
	for _, x := range trickItOne.iterOne([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		i += x
		if i >= 36 {
			break
		}
	}
	if i != 1 {
		t.Error("Expected i == 1, saw ", i, " instead")
	} else {
		t.Log("i = ", i)
	}
	defer func() {
		matchErrorHelper(t, recover(), RERR_EXHAUSTED)
	}()
	trickItOne.fail()
}

func TestTrickyIterZero(t testingT) {
	trickItZero := TrickyIterator{}
	i := 0
	for _, x := range trickItZero.iterZero([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		i += x
		if i >= 36 {
			break
		}
	}
	// Don't care about value, ought to be 0 anyhow.
	t.Log("i = ", i)
	defer func() {
		matchErrorHelper(t, recover(), RERR_EXHAUSTED)
	}()
	trickItZero.fail()
}

func TestTrickyIterZeroCheck(t testingT) {
	trickItZero := TrickyIterator{}
	i := 0
	for _, x := range Check(trickItZero.iterZero([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) {
		i += x
		if i >= 36 {
			break
		}
	}
	// Don't care about value, ought to be 0 anyhow.
	t.Log("i = ", i)
	defer func() {
		matchErrorHelper(t, recover(), CERR_EXHAUSTED)
	}()
	trickItZero.fail()
}

func TestTrickyIterEcho(t testingT) {
	trickItAll := TrickyIterator{}
	i := 0
	for _, x := range trickItAll.iterAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("first loop i=", i)
		i += x
		if i >= 10 {
			break
		}
	}

	if i != 10 {
		t.Error("Expected i == 10, saw", i, "instead")
	} else {
		t.Log("i = ", i)
	}

	defer func() {
		matchErrorHelper(t, recover(), RERR_EXHAUSTED)
		t.Log("end i=", i)
	}()

	i = 0
	for _, x := range trickItAll.iterEcho([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("second loop i=", i)
		if x >= 5 {
			break
		}
	}

}

func TestTrickyIterEcho2(t testingT) {
	trickItAll := TrickyIterator{}
	var i int

	defer func() {
		matchErrorHelper(t, recover(), RERR_EXHAUSTED)
		t.Log("end i=", i)
	}()

	for k := range 2 {
		i = 0
		for _, x := range trickItAll.iterEcho([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			t.Log("k=", k, ",x=", x, ",i=", i)
			i += x
			if i >= 10 {
				break
			}
		}
		t.Log("i = ", i)

		if i != 10 {
			t.Error("Expected i == 10, saw ", i, "instead")
		}
	}
}

// TestBreak1 should just work, with well-behaved iterators.
// (The misbehaving iterator detector should not trigger.)
func TestBreak1(t testingT) {
	var result []int
	var expect = []int{1, 2, -1, 1, 2, -2, 1, 2, -3}
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4}) {
		if x == -4 {
			break
		}
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				break
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, " got ", result)
	}
}

// TestBreak2 should just work, with well-behaved iterators.
// (The misbehaving iterator detector should not trigger.)
func TestBreak2(t testingT) {
	var result []int
	var expect = []int{1, 2, -1, 1, 2, -2, 1, 2, -3}
outer:
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4}) {
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				break
			}
			if x == -4 {
				break outer
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got ", result)
	}
}

// TestContinue should just work, with well-behaved iterators.
// (The misbehaving iterator detector should not trigger.)
func TestContinue(t testingT) {
	var result []int
	var expect = []int{-1, 1, 2, -2, 1, 2, -3, 1, 2, -4}
outer:
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4}) {
		result = append(result, x)
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				continue outer
			}
			if x == -4 {
				break outer
			}
			result = append(result, y)
		}
		result = append(result, x-10)
	}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got ", result)
	}
}

// TestBreak3 should just work, with well-behaved iterators.
// (The misbehaving iterator detector should not trigger.)
func TestBreak3(t testingT) {
	var result []int
	var expect = []int{100, 10, 2, 4, 200, 10, 2, 4, 20, 2, 4, 300, 10, 2, 4, 20, 2, 4, 30}
X:
	for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
	Y:
		for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
			if 10*y >= x {
				break
			}
			result = append(result, y)
			if y == 30 {
				continue X
			}
		Z:
			for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
				if z&1 == 1 {
					continue Z
				}
				result = append(result, z)
				if z >= 4 {
					continue Y
				}
			}
			result = append(result, -y) // should never be executed
		}
		result = append(result, x)
	}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got ", result)
	}
}

// TestBreak1BadA should end in a panic when the outer-loop's
// single-level break is ignore by BadOfSliceIndex
func TestBreak1BadA(t testingT) {
	var result []int
	var expect = []int{1, 2, -1, 1, 2, -2, 1, 2, -3}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range BadOfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		if x == -4 {
			break
		}
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				break
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
}

// TestBreak1BadB should end in a panic, sooner, when the inner-loop's
// (nested) single-level break is ignored by BadOfSliceIndex
func TestBreak1BadB(t testingT) {
	var result []int
	var expect = []int{1, 2} // inner breaks, panics, after before outer appends
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		if x == -4 {
			break
		}
		for _, y := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				break
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
}

// TestMultiCont0 tests multilevel continue with no bad iterators
// (it should just work)
func TestMultiCont0(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4, 2000}
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						continue W // modified to be multilevel
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got %v", expect, result)
	}
}

// TestMultiCont1 tests multilevel continue with a bad iterator
// in the outermost loop exited by the continue.
func TestMultiCont1(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range BadOfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						continue W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiCont2 tests multilevel continue with a bad iterator
// in a middle loop exited by the continue.
func TestMultiCont2(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range BadOfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						continue W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiCont3 tests multilevel continue with a bad iterator
// in the innermost loop exited by the continue.
func TestMultiCont3(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						continue W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiBreak0 tests multilevel break with a bad iterator
// in the outermost loop exited by the break (the outermost loop).
func TestMultiBreak0(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range BadOfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						break W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiBreak1 tests multilevel break with a bad iterator
// in an intermediate loop exited by the break.
func TestMultiBreak1(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range BadOfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						break W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiBreak2 tests multilevel break with two bad iterators
// in intermediate loops exited by the break.
func TestMultiBreak2(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range BadOfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range BadOfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						break W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestMultiBreak3 tests multilevel break with the bad iterator
// in the innermost loop exited by the break.
func TestMultiBreak3(t testingT) {
	var result []int
	var expect = []int{1000, 10, 2, 4}
	defer func() {
		t.Log(result)
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got", result)
		}
	}()
W:
	for _, w := range OfSliceIndex([]int{1000, 2000}) {
		result = append(result, w)
		if w == 2000 {
			break
		}
		for _, x := range OfSliceIndex([]int{100, 200, 300, 400}) {
			for _, y := range OfSliceIndex([]int{10, 20, 30, 40}) {
				result = append(result, y)
				for _, z := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
					if z&1 == 1 {
						continue
					}
					result = append(result, z)
					if z >= 4 {
						break W
					}
				}
				result = append(result, -y) // should never be executed
			}
			result = append(result, x)
		}
	}
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

func TestPanickyIterator1(t testingT) {
	var result []int
	var expect = []int{1, 2, 3, 4}
	defer func() {
		matchErrorHelper(t, recover(), "Panicky iterator panicking")
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, z := range PanickyOfSliceIndex([]int{1, 2, 3, 4}) {
		result = append(result, z)
		if z == 4 {
			break
		}
	}
}

func TestPanickyIterator1Check(t testingT) {
	var result []int
	var expect = []int{1, 2, 3, 4}
	defer func() {
		matchErrorHelper(t, recover(), "Panicky iterator panicking")
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, z := range Check(PanickyOfSliceIndex([]int{1, 2, 3, 4})) {
		result = append(result, z)
		if z == 4 {
			break
		}
	}
}

func TestPanickyIterator2(t testingT) {
	var result []int
	var expect = []int{100, 10, 1, 2}
	defer func() {
		matchErrorHelper(t, recover(), RERR_MISSING)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range OfSliceIndex([]int{100, 200}) {
		result = append(result, x)
	Y:
		// swallows panics and iterates to end BUT `break Y` disables the body, so--> 10, 1, 2
		for _, y := range VeryBadOfSliceIndex([]int{10, 20}) {
			result = append(result, y)

			// converts early exit into a panic --> 1, 2
			for k, z := range PanickyOfSliceIndex([]int{1, 2}) { // iterator panics
				result = append(result, z)
				if k == 1 {
					break Y
				}
			}
		}
	}
}

func TestPanickyIterator2Check(t testingT) {
	var result []int
	var expect = []int{100, 10, 1, 2}
	defer func() {
		matchErrorHelper(t, recover(), CERR_MISSING)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range Check(OfSliceIndex([]int{100, 200})) {
		result = append(result, x)
	Y:
		// swallows panics and iterates to end BUT `break Y` disables the body, so--> 10, 1, 2
		for _, y := range Check(VeryBadOfSliceIndex([]int{10, 20})) {
			result = append(result, y)

			// converts early exit into a panic --> 1, 2
			for k, z := range Check(PanickyOfSliceIndex([]int{1, 2})) { // iterator panics
				result = append(result, z)
				if k == 1 {
					break Y
				}
			}
		}
	}
}

func TestPanickyIterator3(t testingT) {
	var result []int
	var expect = []int{100, 10, 1, 2, 200, 10, 1, 2}
	defer func() {
		if r := recover(); r != nil {
			t.Error("Unexpected panic ", r)
		}
		t.Log(result)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range OfSliceIndex([]int{100, 200}) {
		result = append(result, x)
	Y:
		// swallows panics and iterates to end BUT `break Y` disables the body, so--> 10, 1, 2
		// This is cross-checked against the checked iterator below; the combinator should behave the same.
		for _, y := range VeryBadOfSliceIndex([]int{10, 20}) {
			result = append(result, y)

			for k, z := range OfSliceIndex([]int{1, 2}) { // iterator does not panic
				result = append(result, z)
				if k == 1 {
					break Y
				}
			}
		}
	}
}
func TestPanickyIterator3Check(t testingT) {
	var result []int
	var expect = []int{100, 10, 1, 2, 200, 10, 1, 2}
	defer func() {
		if r := recover(); r != nil {
			t.Error("Unexpected panic ", r)
		}
		t.Log(result)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range Check(OfSliceIndex([]int{100, 200})) {
		result = append(result, x)
	Y:
		// swallows panics and iterates to end BUT `break Y` disables the body, so--> 10, 1, 2
		for _, y := range Check(VeryBadOfSliceIndex([]int{10, 20})) {
			result = append(result, y)

			for k, z := range Check(OfSliceIndex([]int{1, 2})) { // iterator does not panic
				result = append(result, z)
				if k == 1 {
					break Y
				}
			}
		}
	}
}

func TestPanickyIterator4(t testingT) {
	var result []int
	var expect = []int{1, 2, 3}
	defer func() {
		matchErrorHelper(t, recover(), RERR_MISSING)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range SwallowPanicOfSliceIndex([]int{1, 2, 3, 4}) {
		result = append(result, x)
		if x == 3 {
			panic("x is 3")
		}
	}

}

func TestPanickyIterator4Check(t testingT) {
	var result []int
	var expect = []int{1, 2, 3}
	defer func() {
		matchErrorHelper(t, recover(), CERR_MISSING)
		if !slicesEqual(expect, result) {
			t.Error("Expected ", expect, ", got ", result)
		}
	}()
	for _, x := range Check(SwallowPanicOfSliceIndex([]int{1, 2, 3, 4})) {
		result = append(result, x)
		if x == 3 {
			panic("x is 3")
		}
	}

}

// veryBad tests that a loop nest behaves sensibly in the face of a
// "very bad" iterator.  In this case, "sensibly" means that the
// break out of X still occurs after the very bad iterator finally
// quits running (the control flow bread crumbs remain.)
func veryBad(s []int) []int {
	var result []int
X:
	for _, x := range OfSliceIndex([]int{1, 2, 3}) {
		result = append(result, x)
		for _, y := range VeryBadOfSliceIndex(s) {
			result = append(result, y)
			break X
		}
		for _, z := range OfSliceIndex([]int{100, 200, 300}) {
			result = append(result, z)
			if z == 100 {
				break
			}
		}
	}
	return result
}

// veryBadCheck wraps a "very bad" iterator with Check,
// demonstrating that the very bad iterator also hides panics
// thrown by Check.
func veryBadCheck(s []int) []int {
	var result []int
X:
	for _, x := range OfSliceIndex([]int{1, 2, 3}) {
		result = append(result, x)
		for _, y := range Check(VeryBadOfSliceIndex(s)) {
			result = append(result, y)
			break X
		}
		for _, z := range OfSliceIndex([]int{100, 200, 300}) {
			result = append(result, z)
			if z == 100 {
				break
			}
		}
	}
	return result
}

// okay is the not-bad version of veryBad.
// They should behave the same.
func okay(s []int) []int {
	var result []int
X:
	for _, x := range OfSliceIndex([]int{1, 2, 3}) {
		result = append(result, x)
		for _, y := range OfSliceIndex(s) {
			result = append(result, y)
			break X
		}
		for _, z := range OfSliceIndex([]int{100, 200, 300}) {
			result = append(result, z)
			if z == 100 {
				break
			}
		}
	}
	return result
}

// TestVeryBad1 checks the behavior of an extremely poorly behaved iterator.
func TestVeryBad1(t testingT) {
	result := veryBad([]int{10, 20, 30, 40, 50}) // odd length
	expect := []int{1, 10}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestVeryBad2 checks the behavior of an extremely poorly behaved iterator.
func TestVeryBad2(t testingT) {
	result := veryBad([]int{10, 20, 30, 40}) // even length
	expect := []int{1, 10}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestVeryBadCheck checks the behavior of an extremely poorly behaved iterator,
// which also suppresses the exceptions from "Check"
func TestVeryBadCheck(t testingT) {
	result := veryBadCheck([]int{10, 20, 30, 40}) // even length
	expect := []int{1, 10}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// TestOk is the nice version of the very bad iterator.
func TestOk(t testingT) {
	result := okay([]int{10, 20, 30, 40, 50}) // odd length
	expect := []int{1, 10}
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
}

// testBreak1BadDefer checks that defer behaves properly even in
// the presence of loop bodies panicking out of bad iterators.
// (i.e., the instrumentation did not break defer in these loops)
func testBreak1BadDefer(t testingT) (result []int) {
	var expect = []int{1, 2, -1, 1, 2, -2, 1, 2, -3, -30, -20, -10}
	defer func() {
		matchErrorHelper(t, recover(), RERR_DONE)
		if !slicesEqual(expect, result) {
			t.Error("(Inner) Expected ", expect, ", got", result)
		}
	}()
	for _, x := range BadOfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				break
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	return
}

func TestBreak1BadDefer(t testingT) {
	var result []int
	var expect = []int{1, 2, -1, 1, 2, -2, 1, 2, -3, -30, -20, -10}
	result = testBreak1BadDefer(t)
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("(Outer) Expected ", expect, ", got ", result)
	}
}

// testReturn1 has no bad iterators.
func testReturn1() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				return
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	return
}

// testReturn2 has an outermost bad iterator
func testReturn2() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range BadOfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				return
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	return
}

// testReturn3 has an innermost bad iterator
func testReturn3() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				return
			}
			result = append(result, y)
		}
	}
	return
}

// testReturn4 has no bad iterators, but exercises  return variable rewriting
// differs from testReturn1 because deferred append to "result" does not change
// the return value in this case.
func testReturn4(t testingT) (_ []int, _ []int, err any) {
	var result []int
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				return result, result, nil
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
	return
}

// TestReturns checks that returns through bad iterators behave properly,
// for inner and outer bad iterators.
func TestReturns(t testingT) {
	var result []int
	var result2 []int
	var expect = []int{-1, 1, 2, -10}
	var expect2 = []int{-1, 1, 2}
	var err any
	result, err = testReturn1()
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	result, err = testReturn2()
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}
	result, err = testReturn3()
	t.Log(result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}

	result, result2, err = testReturn4(t)
	if !slicesEqual(expect2, result) {
		t.Error("Expected ", expect2, "got", result)
	}
	if !slicesEqual(expect2, result2) {
		t.Error("Expected ", expect2, "got", result2)
	}
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

// testGotoA1 tests loop-nest-internal goto, no bad iterators.
func testGotoA1() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto A
			}
			result = append(result, y)
		}
		result = append(result, x)
	A:
	}
	return
}

// testGotoA2 tests loop-nest-internal goto, outer bad iterator.
func testGotoA2() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range BadOfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto A
			}
			result = append(result, y)
		}
		result = append(result, x)
	A:
	}
	return
}

// testGotoA3 tests loop-nest-internal goto, inner bad iterator.
func testGotoA3() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto A
			}
			result = append(result, y)
		}
		result = append(result, x)
	A:
	}
	return
}
func TestGotoA(t testingT) {
	var result []int
	var expect = []int{-1, 1, 2, -2, 1, 2, -3, 1, 2, -4, -30, -20, -10}
	var expect3 = []int{-1, 1, 2, -10} // first goto becomes a panic
	var err any
	result, err = testGotoA1()
	t.Log("testGotoA1", result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	result, err = testGotoA2()
	t.Log("testGotoA2", result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}
	result, err = testGotoA3()
	t.Log("testGotoA3", result)
	if !slicesEqual(expect3, result) {
		t.Error("Expected %v, got %v", expect3, result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}
}

// testGotoB1 tests loop-nest-exiting goto, no bad iterators.
func testGotoB1() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto B
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
B:
	result = append(result, 999)
	return
}

// testGotoB2 tests loop-nest-exiting goto, outer bad iterator.
func testGotoB2() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range BadOfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range OfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto B
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
B:
	result = append(result, 999)
	return
}

// testGotoB3 tests loop-nest-exiting goto, inner bad iterator.
func testGotoB3() (result []int, err any) {
	defer func() {
		err = recover()
	}()
	for _, x := range OfSliceIndex([]int{-1, -2, -3, -4, -5}) {
		result = append(result, x)
		if x == -4 {
			break
		}
		defer func() {
			result = append(result, x*10)
		}()
		for _, y := range BadOfSliceIndex([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			if y == 3 {
				goto B
			}
			result = append(result, y)
		}
		result = append(result, x)
	}
B:
	result = append(result, 999)
	return
}

func TestGotoB(t testingT) {
	var result []int
	var expect = []int{-1, 1, 2, 999, -10}
	var expectX = []int{-1, 1, 2, -10}
	var err any
	result, err = testGotoB1()
	t.Log("testGotoB1", result)
	if !slicesEqual(expect, result) {
		t.Error("Expected ", expect, ", got", result)
	}
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	result, err = testGotoB2()
	t.Log("testGotoB2", result)
	if !slicesEqual(expectX, result) {
		t.Error("Expected %v, got %v", expectX, result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}

	result, err = testGotoB3()
	t.Log("testGotoB3", result)
	if !slicesEqual(expectX, result) {
		t.Error("Expected %v, got %v", expectX, result)
	}
	if err == nil {
		t.Error("Missing expected error")
	} else {
		matchErrorHelper(t, err, RERR_DONE)
	}
}

// once returns an iterator that runs its loop body once with the supplied value
func once[T any](x T) Seq[T] {
	return func(yield func(T) bool) {
		yield(x)
	}
}

// terrify converts an iterator into one that panics with the supplied string
// if/when the loop body terminates early (returns false, for break, goto, outer
// continue, or return).
func terrify[T any](s string, forall Seq[T]) Seq[T] {
	return func(yield func(T) bool) {
		forall(func(v T) bool {
			if !yield(v) {
				panic(s)
			}
			return true
		})
	}
}

func use[T any](T) {
}

// f runs a not-rangefunc iterator that recovers from a panic that follows execution of a return.
// what does f return?
func f() string {
	defer func() { recover() }()
	defer panic("f panic")
	for _, s := range []string{"f return"} {
		return s
	}
	return "f not reached"
}

// g runs a rangefunc iterator that recovers from a panic that follows execution of a return.
// what does g return?
func g() string {
	defer func() { recover() }()
	for s := range terrify("g panic", once("g return")) {
		return s
	}
	return "g not reached"
}

// h runs a rangefunc iterator that recovers from a panic that follows execution of a return.
// the panic occurs in the rangefunc iterator itself.
// what does h return?
func h() (hashS string) {
	defer func() { recover() }()
	for s := range terrify("h panic", once("h return")) {
		hashS := s
		use(hashS)
		return s
	}
	return "h not reached"
}

func j() (hashS string) {
	defer func() { recover() }()
	for s := range terrify("j panic", once("j return")) {
		hashS = s
		return
	}
	return "j not reached"
}

// k runs a rangefunc iterator that recovers from a panic that follows execution of a return.
// the panic occurs in the rangefunc iterator itself.
// k includes an additional mechanism to for making the return happen
// what does k return?
func k() (hashS string) {
	_return := func(s string) { hashS = s }

	defer func() { recover() }()
	for s := range terrify("k panic", once("k return")) {
		_return(s)
		return
	}
	return "k not reached"
}

func m() (hashS string) {
	_return := func(s string) { hashS = s }

	defer func() { recover() }()
	for s := range terrify("m panic", once("m return")) {
		defer _return(s)
		return s + ", but should be replaced in a defer"
	}
	return "m not reached"
}

func n() string {
	defer func() { recover() }()
	for s := range terrify("n panic", once("n return")) {
		return s + func(s string) string {
			defer func() { recover() }()
			for s := range terrify("n closure panic", once(s)) {
				return s
			}
			return "n closure not reached"
		}(" and n closure return")
	}
	return "n not reached"
}

type terrifyTestCase struct {
	f func() string
	e string
}

func TestPanicReturns(t testingT) {
	tcs := []terrifyTestCase{
		{f, "f return"},
		{g, "g return"},
		{h, "h return"},
		{k, "k return"},
		{j, "j return"},
		{m, "m return"},
		{n, "n return and n closure return"},
	}

	for _, tc := range tcs {
		got := tc.f()
		if got != tc.e {
			t.Error("Got '", got, "' expected ", tc.e)
		} else {
			t.Log("Got expected '", got, "'")
		}
	}
}
```

## File: go/ssa/interp/testdata/rangeoverint.go
```go
package main

// Range over integers (Go 1.22).

import "fmt"

func f() {
	s := "AB"
	for range 5 {
		s += s
	}
	if s != "ABABABABABABABABABABABABABABABABABABABABABABABABABABABABABABABAB" {
		panic(s)
	}

	var t []int
	for i := range 10 {
		t = append(t, i)
	}
	if got, want := fmt.Sprint(t), "[0 1 2 3 4 5 6 7 8 9]"; got != want {
		panic(got)
	}

	var u []uint
	for i := range uint(3) {
		u = append(u, i)
	}
	if got, want := fmt.Sprint(u), "[0 1 2]"; got != want {
		panic(got)
	}

	for i := range 0 {
		panic(i)
	}

	for i := range int(-1) {
		panic(i)
	}

	for _, test := range []struct {
		x    int
		b, c bool
		want string
	}{
		{-1, false, false, "[-123 -123]"},
		{0, false, false, "[-123 -123]"},
		{1, false, false, "[-123 0 333 333]"},
		{2, false, false, "[-123 0 333 1 333 333]"},
		{2, false, true, "[-123 0 222 1 222 222]"},
		{2, true, false, "[-123 0 111 111]"},
		{3, false, false, "[-123 0 333 1 333 2 333 333]"},
	} {
		got := fmt.Sprint(valueSequence(test.x, test.b, test.c))
		if got != test.want {
			panic(fmt.Sprint(test, got))
		}
	}
}

// valueSequence returns a sequence of the values of i.
// b causes an early break and c causes a continue.
func valueSequence(x int, b, c bool) []int {
	var vals []int
	var i int = -123
	vals = append(vals, i)
	for i = range x {
		vals = append(vals, i)
		if b {
			i = 111
			vals = append(vals, i)
			break
		} else if c {
			i = 222
			vals = append(vals, i)
			continue
		}
		i = 333
		vals = append(vals, i)
	}
	vals = append(vals, i)
	return vals
}

func main() { f() }
```

## File: go/ssa/interp/testdata/rangevarlifetime_go122.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.22

package main

func main() {
	test_init()

	// Clones from cmd/compile/internal/loopvar/testdata .
	range_esc_address()
	range_esc_closure()
	range_esc_method()
}

// After go1.22, each i will have a distinct address.
var distinct = func(a [3]int) []*int {
	var r []*int
	for i := range a {
		r = append(r, &i)
	}
	return r
}([3]int{})

func test_init() {
	if len(distinct) != 3 {
		panic(distinct)
	}
	for i := 0; i < 3; i++ {
		if i != *(distinct[i]) {
			panic(distinct)
		}
	}
}

func range_esc_address() {
	// Clone of range_esc_address.go
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []*int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, &i)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw ", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, pi := range is {
		sum += *pi
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected", 20, ", saw", sum)
		bug = true
	}
	if bug {
		panic("range_esc_address")
	}
}

func range_esc_closure() {
	// Clone of range_esc_closure.go
	var ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var is []func() int

	sum := 0
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected ", 20, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_closure")
	}
}

type I int

func (x *I) method() int {
	return int(*x)
}

func range_esc_method() {
	// Clone of range_esc_method.go
	var ints = []I{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []func() int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if int(i) == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, i.method)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, m := range is {
		sum += m()
	}
	if sum != 0+2+4+6+8 {
		println("wrong sum, expected ", 20, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_method")
	}
}
```

## File: go/ssa/interp/testdata/rangevarlifetime_old.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21

// goversion can be pinned to anything strictly before 1.22.

package main

func main() {
	test_init()

	// Clones from cmd/compile/internal/loopvar/testdata .
	range_esc_address()
	range_esc_closure()
	range_esc_method()
}

// pre-go1.22 all of i will have the same address.
var same = func(a [3]int) []*int {
	var r []*int
	for i := range a {
		r = append(r, &i)
	}
	return r
}([3]int{})

func test_init() {
	if len(same) != 3 {
		panic(same)
	}
	for i := range same {
		for j := range same {
			if !(same[i] == same[j]) {
				panic(same)
			}
		}
	}
	for i := range same {
		if *(same[i]) != 2 {
			panic(same)
		}
	}
}

func range_esc_address() {
	// Clone of range_esc_address.go
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []*int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, &i)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw ", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, pi := range is {
		sum += *pi
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected", 45, ", saw", sum)
		bug = true
	}
	if bug {
		panic("range_esc_address")
	}
}

func range_esc_closure() {
	// Clone of range_esc_closure.go
	var ints = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var is []func() int

	sum := 0
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected ", 45, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_closure")
	}
}

type I int

func (x *I) method() int {
	return int(*x)
}

func range_esc_method() {
	// Clone of range_esc_method.go
	var ints = []I{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0
	var is []func() int
	for _, i := range ints {
		for j := 0; j < 10; j++ {
			if int(i) == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, i.method)
		}
	}

	bug := false
	if sum != 100-10 {
		println("wrong sum, expected", 90, ", saw", sum)
		bug = true
	}
	if len(is) != 5 {
		println("wrong iterations, expected ", 5, ", saw", len(is))
		bug = true
	}
	sum = 0
	for _, m := range is {
		sum += m()
	}
	if sum != 9+9+9+9+9 {
		println("wrong sum, expected ", 45, ", saw ", sum)
		bug = true
	}
	if bug {
		panic("range_esc_method")
	}
}
```

## File: go/ssa/interp/testdata/recover.go
```go
package main

// Tests of panic/recover.

import "fmt"

func fortyTwo() (r int) {
	r = 42
	// The next two statements simulate a 'return' statement.
	defer func() { recover() }()
	panic(nil)
}

func zero() int {
	defer func() { recover() }()
	panic(1)
}

func zeroEmpty() (int, string) {
	defer func() { recover() }()
	panic(1)
}

func main() {
	if r := fortyTwo(); r != 42 {
		panic(r)
	}
	if r := zero(); r != 0 {
		panic(r)
	}
	if r, s := zeroEmpty(); r != 0 || s != "" {
		panic(fmt.Sprint(r, s))
	}
}
```

## File: go/ssa/interp/testdata/reflect.go
```go
package main

import "reflect"

func main() {
	// Regression test for issue 9462.
	got := reflect.SliceOf(reflect.TypeOf(byte(0))).String()
	if got != "[]uint8" && got != "[]byte" { // result varies by toolchain
		println("BUG: " + got)
	}
}
```

## File: go/ssa/interp/testdata/slice2array.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test for slice to array conversion introduced in go1.20
// See: https://tip.golang.org/ref/spec#Conversions_from_slice_to_array_pointer

package main

func main() {
	s := make([]byte, 3, 4)
	s[0], s[1], s[2] = 2, 3, 5
	a := ([2]byte)(s)
	s[0] = 7

	if a != [2]byte{2, 3} {
		panic("converted from non-nil slice to array")
	}

	{
		var s []int
		a := ([0]int)(s)
		if a != [0]int{} {
			panic("zero len array is not equal")
		}
	}

	if emptyToEmptyDoesNotPanic() {
		panic("no panic expected from emptyToEmptyDoesNotPanic()")
	}
	if !threeToFourDoesPanic() {
		panic("panic expected from threeToFourDoesPanic()")
	}

	if !fourPanicsWhileOneDoesNot[[4]int]() {
		panic("panic expected from fourPanicsWhileOneDoesNot[[4]int]()")
	}
	if fourPanicsWhileOneDoesNot[[1]int]() {
		panic("no panic expected from fourPanicsWhileOneDoesNot[[1]int]()")
	}

	if !fourPanicsWhileZeroDoesNot[[4]int]() {
		panic("panic expected from fourPanicsWhileZeroDoesNot[[4]int]()")
	}
	if fourPanicsWhileZeroDoesNot[[0]int]() {
		panic("no panic expected from fourPanicsWhileZeroDoesNot[[0]int]()")
	}
}

func emptyToEmptyDoesNotPanic() (raised bool) {
	defer func() {
		if e := recover(); e != nil {
			raised = true
		}
	}()
	var s []int
	_ = ([0]int)(s)
	return false
}

func threeToFourDoesPanic() (raised bool) {
	defer func() {
		if e := recover(); e != nil {
			raised = true
		}
	}()
	s := make([]int, 3, 5)
	_ = ([4]int)(s)
	return false
}

func fourPanicsWhileOneDoesNot[T [1]int | [4]int]() (raised bool) {
	defer func() {
		if e := recover(); e != nil {
			raised = true
		}
	}()
	s := make([]int, 3, 5)
	_ = T(s)
	return false
}

func fourPanicsWhileZeroDoesNot[T [0]int | [4]int]() (raised bool) {
	defer func() {
		if e := recover(); e != nil {
			raised = true
		}
	}()
	var s []int
	_ = T(s)
	return false
}
```

## File: go/ssa/interp/testdata/slice2arrayptr.go
```go
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test for slice to array pointer conversion introduced in go1.17
// See: https://tip.golang.org/ref/spec#Conversions_from_slice_to_array_pointer

package main

func main() {
	s := make([]byte, 2, 4)
	if s0 := (*[0]byte)(s); s0 == nil {
		panic("converted from non-nil slice result in nil array pointer")
	}
	if s2 := (*[2]byte)(s); &s2[0] != &s[0] {
		panic("the converted array is not slice underlying array")
	}
	wantPanic(
		func() {
			_ = (*[4]byte)(s) // panics: len([4]byte) > len(s)
		},
		"runtime error: array length is greater than slice length",
	)

	var t []string
	if t0 := (*[0]string)(t); t0 != nil {
		panic("nil slice converted to *[0]byte should be nil")
	}
	wantPanic(
		func() {
			_ = (*[1]string)(t) // panics: len([1]string) > len(t)
		},
		"runtime error: array length is greater than slice length",
	)

	f()
}

type arr [2]int

func f() {
	s := []int{1, 2, 3, 4}
	_ = *(*arr)(s)
}

func wantPanic(fn func(), s string) {
	defer func() {
		err := recover()
		if err == nil {
			panic("expected panic")
		}
		if got := err.(error).Error(); got != s {
			panic("expected panic " + s + " got " + got)
		}
	}()
	fn()
}
```

## File: go/ssa/interp/testdata/static.go
```go
package main

// Static tests of SSA builder (via the sanity checker).
// Dynamic semantics are not exercised.

func init() {
	// Regression test for issue 6806.
	ch := make(chan int)
	select {
	case n, _ := <-ch:
		_ = n
	default:
		// The default case disables the simplification of
		// select to a simple receive statement.
	}

	// value,ok-form receive where TypeOf(ok) is a named boolean.
	type mybool bool
	var x int
	var y mybool
	select {
	case x, y = <-ch:
	default:
		// The default case disables the simplification of
		// select to a simple receive statement.
	}
	_ = x
	_ = y
}

var a int

// Regression test for issue 7840 (covered by SSA sanity checker).
func bug7840() bool {
	// This creates a single-predecessor block with a φ-node.
	return false && a == 0 && a == 0
}

// A blocking select (sans "default:") cannot fall through.
// Regression test for issue 7022.
func bug7022() int {
	var c1, c2 chan int
	select {
	case <-c1:
		return 123
	case <-c2:
		return 456
	}
}

// Parens should not prevent intrinsic treatment of built-ins.
// (Regression test for a crash.)
func init() {
	_ = (new)(int)
	_ = (make)([]int, 0)
}

func main() {}
```

## File: go/ssa/interp/testdata/typeassert.go
```go
// Tests of type asserts.
// Requires type parameters.
package typeassert

type fooer interface{ foo() string }

type X int

func (_ X) foo() string { return "x" }

func f[T fooer](x T) func() string {
	return x.foo
}

func main() {
	if f[X](0)() != "x" {
		panic("f[X]() != 'x'")
	}

	p := false
	func() {
		defer func() {
			if recover() != nil {
				p = true
			}
		}()
		f[fooer](nil) // panics on x.foo when T is an interface and nil.
	}()
	if !p {
		panic("f[fooer] did not panic")
	}
}
```

## File: go/ssa/interp/testdata/width32.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test interpretation on 32 bit widths.

package main

func main() {
	mapSize()
}

func mapSize() {
	// Tests for the size argument of make on a map type.
	const tooBigFor32 = 1<<33 - 1
	wantPanic(
		func() {
			_ = make(map[int]int, int64(tooBigFor32))
		},
		"runtime error: ssa.MakeMap.Reserve value 8589934591 does not fit in int",
	)

	// TODO: Enable the following if sizeof(int) can be different for host and target.
	// _ = make(map[int]int, tooBigFor32)
	//
	// Second arg to make in `make(map[int]int, tooBigFor32)` is an untyped int and
	// is converted into an int explicitly in ssa.
	// This has a different value on 32 and 64 bit systems.
}

func wantPanic(fn func(), s string) {
	defer func() {
		err := recover()
		if err == nil {
			panic("expected panic")
		}
		if got := err.(error).Error(); got != s {
			panic("expected panic " + s + " got " + got)
		}
	}()
	fn()
}
```

## File: go/ssa/interp/testdata/zeros.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test interpretation on zero values with type params.
package zeros

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func tp0[T int | string | float64]() T { return T(0) }

func tpFalse[T ~bool]() T { return T(false) }

func tpEmptyString[T string | []byte]() T { return T("") }

func tpNil[T *int | []byte]() T { return T(nil) }

func main() {
	// zero values
	var zi int
	var zf float64
	var zs string

	assert(zi == int(0), "zero value of int is int(0)")
	assert(zf == float64(0), "zero value of float64 is float64(0)")
	assert(zs != string(0), "zero value of string is not string(0)")

	assert(zi == tp0[int](), "zero value of int is int(0)")
	assert(zf == tp0[float64](), "zero value of float64 is float64(0)")
	assert(zs != tp0[string](), "zero value of string is not string(0)")

	assert(zf == -0.0, "constant -0.0 is converted to 0.0")

	assert(!tpFalse[bool](), "zero value of bool is false")

	assert(tpEmptyString[string]() == zs, `zero value of string is string("")`)
	assert(len(tpEmptyString[[]byte]()) == 0, `[]byte("") is empty`)

	assert(tpNil[*int]() == nil, "nil is nil")
	assert(tpNil[[]byte]() == nil, "nil is nil")
}
```

## File: go/ssa/interp/external.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

// Emulated functions that we cannot interpret because they are
// external or because they use "unsafe" or "reflect" operations.

import (
	"bytes"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type externalFn func(fr *frame, args []value) value

// TODO(adonovan): fix: reflect.Value abstracts an lvalue or an
// rvalue; Set() causes mutations that can be observed via aliases.
// We have not captured that correctly here.

// Key strings are from Function.String().
var externals = make(map[string]externalFn)

func init() {
	// That little dot ۰ is an Arabic zero numeral (U+06F0), categories [Nd].
	for k, v := range map[string]externalFn{
		"(reflect.Value).Bool":            ext۰reflect۰Value۰Bool,
		"(reflect.Value).CanAddr":         ext۰reflect۰Value۰CanAddr,
		"(reflect.Value).CanInterface":    ext۰reflect۰Value۰CanInterface,
		"(reflect.Value).Elem":            ext۰reflect۰Value۰Elem,
		"(reflect.Value).Field":           ext۰reflect۰Value۰Field,
		"(reflect.Value).Float":           ext۰reflect۰Value۰Float,
		"(reflect.Value).Index":           ext۰reflect۰Value۰Index,
		"(reflect.Value).Int":             ext۰reflect۰Value۰Int,
		"(reflect.Value).Interface":       ext۰reflect۰Value۰Interface,
		"(reflect.Value).IsNil":           ext۰reflect۰Value۰IsNil,
		"(reflect.Value).IsValid":         ext۰reflect۰Value۰IsValid,
		"(reflect.Value).Kind":            ext۰reflect۰Value۰Kind,
		"(reflect.Value).Len":             ext۰reflect۰Value۰Len,
		"(reflect.Value).MapIndex":        ext۰reflect۰Value۰MapIndex,
		"(reflect.Value).MapKeys":         ext۰reflect۰Value۰MapKeys,
		"(reflect.Value).NumField":        ext۰reflect۰Value۰NumField,
		"(reflect.Value).NumMethod":       ext۰reflect۰Value۰NumMethod,
		"(reflect.Value).Pointer":         ext۰reflect۰Value۰Pointer,
		"(reflect.Value).Set":             ext۰reflect۰Value۰Set,
		"(reflect.Value).String":          ext۰reflect۰Value۰String,
		"(reflect.Value).Type":            ext۰reflect۰Value۰Type,
		"(reflect.Value).Uint":            ext۰reflect۰Value۰Uint,
		"(reflect.error).Error":           ext۰reflect۰error۰Error,
		"(reflect.rtype).Bits":            ext۰reflect۰rtype۰Bits,
		"(reflect.rtype).Elem":            ext۰reflect۰rtype۰Elem,
		"(reflect.rtype).Field":           ext۰reflect۰rtype۰Field,
		"(reflect.rtype).In":              ext۰reflect۰rtype۰In,
		"(reflect.rtype).Kind":            ext۰reflect۰rtype۰Kind,
		"(reflect.rtype).NumField":        ext۰reflect۰rtype۰NumField,
		"(reflect.rtype).NumIn":           ext۰reflect۰rtype۰NumIn,
		"(reflect.rtype).NumMethod":       ext۰reflect۰rtype۰NumMethod,
		"(reflect.rtype).NumOut":          ext۰reflect۰rtype۰NumOut,
		"(reflect.rtype).Out":             ext۰reflect۰rtype۰Out,
		"(reflect.rtype).Size":            ext۰reflect۰rtype۰Size,
		"(reflect.rtype).String":          ext۰reflect۰rtype۰String,
		"bytes.Equal":                     ext۰bytes۰Equal,
		"bytes.IndexByte":                 ext۰bytes۰IndexByte,
		"fmt.Sprint":                      ext۰fmt۰Sprint,
		"math.Abs":                        ext۰math۰Abs,
		"math.Copysign":                   ext۰math۰Copysign,
		"math.Exp":                        ext۰math۰Exp,
		"math.Float32bits":                ext۰math۰Float32bits,
		"math.Float32frombits":            ext۰math۰Float32frombits,
		"math.Float64bits":                ext۰math۰Float64bits,
		"math.Float64frombits":            ext۰math۰Float64frombits,
		"math.Inf":                        ext۰math۰Inf,
		"math.IsNaN":                      ext۰math۰IsNaN,
		"math.Ldexp":                      ext۰math۰Ldexp,
		"math.Log":                        ext۰math۰Log,
		"math.Min":                        ext۰math۰Min,
		"math.NaN":                        ext۰math۰NaN,
		"math.Sqrt":                       ext۰math۰Sqrt,
		"os.Exit":                         ext۰os۰Exit,
		"os.Getenv":                       ext۰os۰Getenv,
		"reflect.New":                     ext۰reflect۰New,
		"reflect.SliceOf":                 ext۰reflect۰SliceOf,
		"reflect.TypeOf":                  ext۰reflect۰TypeOf,
		"reflect.ValueOf":                 ext۰reflect۰ValueOf,
		"reflect.Zero":                    ext۰reflect۰Zero,
		"runtime.Breakpoint":              ext۰runtime۰Breakpoint,
		"runtime.GC":                      ext۰runtime۰GC,
		"runtime.GOMAXPROCS":              ext۰runtime۰GOMAXPROCS,
		"runtime.GOROOT":                  ext۰runtime۰GOROOT,
		"runtime.Goexit":                  ext۰runtime۰Goexit,
		"runtime.Gosched":                 ext۰runtime۰Gosched,
		"runtime.NumCPU":                  ext۰runtime۰NumCPU,
		"sort.Float64s":                   ext۰sort۰Float64s,
		"sort.Ints":                       ext۰sort۰Ints,
		"sort.Strings":                    ext۰sort۰Strings,
		"strconv.Atoi":                    ext۰strconv۰Atoi,
		"strconv.Itoa":                    ext۰strconv۰Itoa,
		"strconv.FormatFloat":             ext۰strconv۰FormatFloat,
		"strings.Count":                   ext۰strings۰Count,
		"strings.EqualFold":               ext۰strings۰EqualFold,
		"strings.Index":                   ext۰strings۰Index,
		"strings.IndexByte":               ext۰strings۰IndexByte,
		"strings.Replace":                 ext۰strings۰Replace,
		"strings.ToLower":                 ext۰strings۰ToLower,
		"time.Sleep":                      ext۰time۰Sleep,
		"unicode/utf8.DecodeRuneInString": ext۰unicode۰utf8۰DecodeRuneInString,
	} {
		externals[k] = v
	}
}

func ext۰bytes۰Equal(fr *frame, args []value) value {
	// func Equal(a, b []byte) bool
	a := args[0].([]value)
	b := args[1].([]value)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ext۰bytes۰IndexByte(fr *frame, args []value) value {
	// func IndexByte(s []byte, c byte) int
	s := args[0].([]value)
	c := args[1].(byte)
	for i, b := range s {
		if b.(byte) == c {
			return i
		}
	}
	return -1
}

func ext۰math۰Float64frombits(fr *frame, args []value) value {
	return math.Float64frombits(args[0].(uint64))
}

func ext۰math۰Float64bits(fr *frame, args []value) value {
	return math.Float64bits(args[0].(float64))
}

func ext۰math۰Float32frombits(fr *frame, args []value) value {
	return math.Float32frombits(args[0].(uint32))
}

func ext۰math۰Abs(fr *frame, args []value) value {
	return math.Abs(args[0].(float64))
}

func ext۰math۰Copysign(fr *frame, args []value) value {
	return math.Copysign(args[0].(float64), args[1].(float64))
}

func ext۰math۰Exp(fr *frame, args []value) value {
	return math.Exp(args[0].(float64))
}

func ext۰math۰Float32bits(fr *frame, args []value) value {
	return math.Float32bits(args[0].(float32))
}

func ext۰math۰Min(fr *frame, args []value) value {
	return math.Min(args[0].(float64), args[1].(float64))
}

func ext۰math۰NaN(fr *frame, args []value) value {
	return math.NaN()
}

func ext۰math۰IsNaN(fr *frame, args []value) value {
	return math.IsNaN(args[0].(float64))
}

func ext۰math۰Inf(fr *frame, args []value) value {
	return math.Inf(args[0].(int))
}

func ext۰math۰Ldexp(fr *frame, args []value) value {
	return math.Ldexp(args[0].(float64), args[1].(int))
}

func ext۰math۰Log(fr *frame, args []value) value {
	return math.Log(args[0].(float64))
}

func ext۰math۰Sqrt(fr *frame, args []value) value {
	return math.Sqrt(args[0].(float64))
}

func ext۰runtime۰Breakpoint(fr *frame, args []value) value {
	runtime.Breakpoint()
	return nil
}

func ext۰sort۰Ints(fr *frame, args []value) value {
	x := args[0].([]value)
	sort.Slice(x, func(i, j int) bool {
		return x[i].(int) < x[j].(int)
	})
	return nil
}
func ext۰sort۰Strings(fr *frame, args []value) value {
	x := args[0].([]value)
	sort.Slice(x, func(i, j int) bool {
		return x[i].(string) < x[j].(string)
	})
	return nil
}
func ext۰sort۰Float64s(fr *frame, args []value) value {
	x := args[0].([]value)
	sort.Slice(x, func(i, j int) bool {
		return x[i].(float64) < x[j].(float64)
	})
	return nil
}

func ext۰strconv۰Atoi(fr *frame, args []value) value {
	i, e := strconv.Atoi(args[0].(string))
	if e != nil {
		return tuple{i, iface{fr.i.runtimeErrorString, e.Error()}}
	}
	return tuple{i, iface{}}
}
func ext۰strconv۰Itoa(fr *frame, args []value) value {
	return strconv.Itoa(args[0].(int))
}
func ext۰strconv۰FormatFloat(fr *frame, args []value) value {
	return strconv.FormatFloat(args[0].(float64), args[1].(byte), args[2].(int), args[3].(int))
}

func ext۰strings۰Count(fr *frame, args []value) value {
	return strings.Count(args[0].(string), args[1].(string))
}

func ext۰strings۰EqualFold(fr *frame, args []value) value {
	return strings.EqualFold(args[0].(string), args[1].(string))
}
func ext۰strings۰IndexByte(fr *frame, args []value) value {
	return strings.IndexByte(args[0].(string), args[1].(byte))
}

func ext۰strings۰Index(fr *frame, args []value) value {
	return strings.Index(args[0].(string), args[1].(string))
}

func ext۰strings۰Replace(fr *frame, args []value) value {
	// func Replace(s, old, new string, n int) string
	s := args[0].(string)
	new := args[1].(string)
	old := args[2].(string)
	n := args[3].(int)
	return strings.Replace(s, old, new, n)
}

func ext۰strings۰ToLower(fr *frame, args []value) value {
	return strings.ToLower(args[0].(string))
}

func ext۰runtime۰GOMAXPROCS(fr *frame, args []value) value {
	// Ignore args[0]; don't let the interpreted program
	// set the interpreter's GOMAXPROCS!
	return runtime.GOMAXPROCS(0)
}

func ext۰runtime۰Goexit(fr *frame, args []value) value {
	// TODO(adonovan): don't kill the interpreter's main goroutine.
	runtime.Goexit()
	return nil
}

func ext۰runtime۰GOROOT(fr *frame, args []value) value {
	return runtime.GOROOT()
}

func ext۰runtime۰GC(fr *frame, args []value) value {
	runtime.GC()
	return nil
}

func ext۰runtime۰Gosched(fr *frame, args []value) value {
	runtime.Gosched()
	return nil
}

func ext۰runtime۰NumCPU(fr *frame, args []value) value {
	return runtime.NumCPU()
}

func ext۰time۰Sleep(fr *frame, args []value) value {
	time.Sleep(time.Duration(args[0].(int64)))
	return nil
}

func ext۰os۰Getenv(fr *frame, args []value) value {
	name := args[0].(string)
	switch name {
	case "GOSSAINTERP":
		return "1"
	}
	return os.Getenv(name)
}

func ext۰os۰Exit(fr *frame, args []value) value {
	panic(exitPanic(args[0].(int)))
}

func ext۰unicode۰utf8۰DecodeRuneInString(fr *frame, args []value) value {
	r, n := utf8.DecodeRuneInString(args[0].(string))
	return tuple{r, n}
}

// A fake function for turning an arbitrary value into a string.
// Handles only the cases needed by the tests.
// Uses same logic as 'print' built-in.
func ext۰fmt۰Sprint(fr *frame, args []value) value {
	buf := new(bytes.Buffer)
	wasStr := false
	for i, arg := range args[0].([]value) {
		x := arg.(iface).v
		_, isStr := x.(string)
		if i > 0 && !wasStr && !isStr {
			buf.WriteByte(' ')
		}
		wasStr = isStr
		buf.WriteString(toString(x))
	}
	return buf.String()
}
```

## File: go/ssa/interp/interp_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp_test

// This test runs the SSA interpreter over sample Go programs.
// Because the interpreter requires intrinsics for assembly
// functions and many low-level runtime routines, it is inherently
// not robust to evolutionary change in the standard library.
// Therefore the test cases are restricted to programs that
// use a fake standard library in testdata/src containing a tiny
// subset of simple functions useful for writing assertions.
//
// We no longer attempt to interpret any real standard packages such as
// fmt or testing, as it proved too fragile.

import (
	"bytes"
	"fmt"
	"go/build"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
	"unsafe"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/interp"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testenv"
)

// Each line contains a space-separated list of $GOROOT/test/
// filenames comprising the main package of a program.
// They are ordered quickest-first, roughly.
//
// If a test in this list fails spuriously, remove it.
var gorootTestTests = []string{
	"235.go",
	"alias1.go",
	"func5.go",
	"func6.go",
	"func7.go",
	"func8.go",
	"helloworld.go",
	"varinit.go",
	"escape3.go",
	"initcomma.go",
	"cmp.go",
	"compos.go",
	"turing.go",
	"indirect.go",
	"complit.go",
	"for.go",
	"struct0.go",
	"intcvt.go",
	"printbig.go",
	"deferprint.go",
	"escape.go",
	"range.go",
	"const4.go",
	"float_lit.go",
	"bigalg.go",
	"decl.go",
	"if.go",
	"named.go",
	"bigmap.go",
	"func.go",
	"reorder2.go",
	"gc.go",
	"simassign.go",
	"iota.go",
	"nilptr2.go",
	"utf.go",
	"method.go",
	"char_lit.go",
	"env.go",
	"int_lit.go",
	"string_lit.go",
	"defer.go",
	"typeswitch.go",
	"stringrange.go",
	"reorder.go",
	"method3.go",
	"literal.go",
	"nul1.go", // doesn't actually assert anything (errorcheckoutput)
	"zerodivide.go",
	"convert.go",
	"convT2X.go",
	"switch.go",
	"ddd.go",
	"blank.go", // partly disabled
	"closedchan.go",
	"divide.go",
	"rename.go",
	"nil.go",
	"recover1.go",
	"recover2.go",
	"recover3.go",
	"typeswitch1.go",
	"floatcmp.go",
	"crlf.go", // doesn't actually assert anything (runoutput)
}

// These are files in go.tools/go/ssa/interp/testdata/.
var testdataTests = []string{
	"boundmeth.go",
	"complit.go",
	"convert.go",
	"coverage.go",
	"deepequal.go",
	"defer.go",
	"fieldprom.go",
	"forvarlifetime_old.go",
	"ifaceconv.go",
	"ifaceprom.go",
	"initorder.go",
	"methprom.go",
	"mrvchain.go",
	"range.go",
	"rangeoverint.go",
	"recover.go",
	"reflect.go",
	"slice2arrayptr.go",
	"static.go",
	"width32.go",
	"rangevarlifetime_old.go",
	"fixedbugs/issue52342.go",
	"fixedbugs/issue55115.go",
	"fixedbugs/issue52835.go",
	"fixedbugs/issue55086.go",
	"fixedbugs/issue66783.go",
	"fixedbugs/issue69929.go",
	"typeassert.go",
	"zeros.go",
	"slice2array.go",
	"minmax.go",
	"rangevarlifetime_go122.go",
	"forvarlifetime_go122.go",
}

func init() {
	// GOROOT/test used to assume that GOOS and GOARCH were explicitly set in the
	// environment, so do that here for TestGorootTest.
	os.Setenv("GOOS", runtime.GOOS)
	os.Setenv("GOARCH", runtime.GOARCH)
}

// run runs a single test. On success it returns the captured std{out,err}.
func run(t *testing.T, input string, goroot string) string {
	testenv.NeedsExec(t) // really we just need os.Pipe, but os/exec uses pipes

	t.Logf("Input: %s\n", input)

	start := time.Now()

	ctx := build.Default // copy
	ctx.GOROOT = goroot
	ctx.GOOS = runtime.GOOS
	ctx.GOARCH = runtime.GOARCH
	if filepath.Base(input) == "width32.go" && unsafe.Sizeof(int(0)) > 4 {
		t.Skipf("skipping: width32.go checks behavior for a 32-bit int")
	}

	gover := ""
	if p := testenv.Go1Point(); p > 0 {
		gover = fmt.Sprintf("go1.%d", p)
	}

	conf := loader.Config{Build: &ctx, TypeChecker: types.Config{GoVersion: gover}}
	if _, err := conf.FromArgs([]string{input}, true); err != nil {
		t.Fatalf("FromArgs(%s) failed: %s", input, err)
	}

	conf.Import("runtime")

	// Print a helpful hint if we don't make it to the end.
	var hint string
	defer func() {
		t.Logf("Duration: %v", time.Since(start))
		if hint != "" {
			t.Log("FAIL")
			t.Log(hint)
		} else {
			t.Log("PASS")
		}
	}()

	hint = fmt.Sprintf("To dump SSA representation, run:\n%% go build golang.org/x/tools/cmd/ssadump && ./ssadump -test -build=CFP %s\n", input)

	iprog, err := conf.Load()
	if err != nil {
		t.Fatalf("conf.Load(%s) failed: %s", input, err)
	}

	bmode := ssa.InstantiateGenerics | ssa.SanityCheckFunctions
	// bmode |= ssa.PrintFunctions // enable for debugging
	prog := ssautil.CreateProgram(iprog, bmode)
	prog.Build()

	mainPkg := prog.Package(iprog.Created[0].Pkg)
	if mainPkg == nil {
		t.Fatalf("not a main package: %s", input)
	}

	sizes := types.SizesFor("gc", ctx.GOARCH)
	if sizes.Sizeof(types.Typ[types.Int]) < 4 {
		panic("bogus SizesFor")
	}
	hint = fmt.Sprintf("To trace execution, run:\n%% go build golang.org/x/tools/cmd/ssadump && ./ssadump -build=C -test -run --interp=T %s\n", input)

	// Capture anything written by the interpreter to os.Std{out,err}
	// by temporarily redirecting them to a buffer via a pipe.
	//
	// While capturing is in effect, we must not write any
	// test-related stuff to stderr (including log.Print, t.Log, etc).
	var restore func() string // restore files and log+return the mixed out/err.
	{
		// Connect std{out,err} to pipe.
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatalf("can't create pipe for stderr: %v", err)
		}
		savedStdout := os.Stdout
		savedStderr := os.Stderr
		os.Stdout = w
		os.Stderr = w

		// Buffer what is written.
		var buf strings.Builder
		done := make(chan struct{})
		go func() {
			if _, err := io.Copy(&buf, r); err != nil {
				fmt.Fprintf(savedStderr, "io.Copy: %v", err)
			}
			close(done)
		}()

		// Finally, restore the files and log what was captured.
		restore = func() string {
			os.Stdout = savedStdout
			os.Stderr = savedStderr
			w.Close()
			<-done
			captured := buf.String()
			t.Logf("Interpreter's stdout+stderr:\n%s", captured)
			return captured
		}
	}

	var imode interp.Mode // default mode
	// imode |= interp.DisableRecover // enable for debugging
	// imode |= interp.EnableTracing // enable for debugging
	exitCode := interp.Interpret(mainPkg, imode, sizes, input, []string{})
	capturedOutput := restore()
	if exitCode != 0 {
		t.Fatalf("interpreting %s: exit code was %d", input, exitCode)
	}
	// $GOROOT/test tests use this convention:
	if strings.Contains(capturedOutput, "BUG") {
		t.Fatalf("interpreting %s: exited zero but output contained 'BUG'", input)
	}

	hint = "" // call off the hounds

	return capturedOutput
}

// makeGoroot copies testdata/src into the "src" directory of a temporary
// location to mimic GOROOT/src, and adds a file "runtime/consts.go" containing
// declarations for GOOS and GOARCH that match the GOOS and GOARCH of this test.
//
// It returns the directory that should be used for GOROOT.
func makeGoroot(t *testing.T) string {
	goroot := t.TempDir()
	src := filepath.Join(goroot, "src")

	err := filepath.Walk("testdata/src", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel("testdata/src", path)
		if err != nil {
			return err
		}
		targ := filepath.Join(src, rel)

		if info.IsDir() {
			return os.Mkdir(targ, info.Mode().Perm()|0700)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targ, b, info.Mode().Perm())
	})
	if err != nil {
		t.Fatal(err)
	}

	constsGo := fmt.Sprintf(`package runtime
const GOOS = %q
const GOARCH = %q
`, runtime.GOOS, runtime.GOARCH)
	err = os.WriteFile(filepath.Join(src, "runtime/consts.go"), []byte(constsGo), 0644)
	if err != nil {
		t.Fatal(err)
	}

	return goroot
}

// TestTestdataFiles runs the interpreter on testdata/*.go.
func TestTestdataFiles(t *testing.T) {
	goroot := makeGoroot(t)
	for _, input := range testdataTests {
		t.Run(input, func(t *testing.T) {
			run(t, filepath.Join("testdata", input), goroot)
		})
	}
}

// TestGorootTest runs the interpreter on $GOROOT/test/*.go.
func TestGorootTest(t *testing.T) {
	testenv.NeedsGOROOTDir(t, "test")

	goroot := makeGoroot(t)
	for _, input := range gorootTestTests {
		t.Run(input, func(t *testing.T) {
			run(t, filepath.Join(build.Default.GOROOT, "test", input), goroot)
		})
	}
}

// TestTypeparamTest runs the interpreter on runnable examples
// in $GOROOT/test/typeparam/*.go.

func TestTypeparamTest(t *testing.T) {
	testenv.NeedsGOROOTDir(t, "test")

	if runtime.GOARCH == "wasm" {
		// See ssa/TestTypeparamTest.
		t.Skip("Consistent flakes on wasm (e.g. https://go.dev/issues/64726)")
	}

	goroot := makeGoroot(t)

	// Skip known failures for the given reason.
	// TODO(taking): Address these.
	skip := map[string]string{
		"chans.go":      "interp tests do not support runtime.SetFinalizer",
		"issue23536.go": "unknown reason",
		"issue48042.go": "interp tests do not handle reflect.Value.SetInt",
		"issue47716.go": "interp tests do not handle unsafe.Sizeof",
		"issue50419.go": "interp tests do not handle dispatch to String() correctly",
		"issue51733.go": "interp does not handle unsafe casts",
		"ordered.go":    "math.NaN() comparisons not being handled correctly",
		"orderedmap.go": "interp tests do not support runtime.SetFinalizer",
		"stringer.go":   "unknown reason",
		"issue48317.go": "interp tests do not support encoding/json",
		"issue48318.go": "interp tests do not support encoding/json",
		"issue58513.go": "interp tests do not support runtime.Caller",
	}
	// Collect all of the .go files in dir that are runnable.
	dir := filepath.Join(build.Default.GOROOT, "test", "typeparam")
	list, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range list {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue // Consider standalone go files.
		}
		t.Run(entry.Name(), func(t *testing.T) {
			input := filepath.Join(dir, entry.Name())
			src, err := os.ReadFile(input)
			if err != nil {
				t.Fatal(err)
			}

			// Only build test files that can be compiled, or compiled and run.
			if !bytes.HasPrefix(src, []byte("// run")) || bytes.HasPrefix(src, []byte("// rundir")) {
				t.Logf("Not a `// run` file: %s", entry.Name())
				return
			}

			if reason := skip[entry.Name()]; reason != "" {
				t.Skipf("skipping: %s", reason)
			}

			run(t, input, goroot)
		})
	}
}
```

## File: go/ssa/interp/interp.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ssa/interp defines an interpreter for the SSA
// representation of Go programs.
//
// This interpreter is provided as an adjunct for testing the SSA
// construction algorithm.  Its purpose is to provide a minimal
// metacircular implementation of the dynamic semantics of each SSA
// instruction.  It is not, and will never be, a production-quality Go
// interpreter.
//
// The following is a partial list of Go features that are currently
// unsupported or incomplete in the interpreter.
//
// * Unsafe operations, including all uses of unsafe.Pointer, are
// impossible to support given the "boxed" value representation we
// have chosen.
//
// * The reflect package is only partially implemented.
//
// * The "testing" package is no longer supported because it
// depends on low-level details that change too often.
//
// * "sync/atomic" operations are not atomic due to the "boxed" value
// representation: it is not possible to read, modify and write an
// interface value atomically. As a consequence, Mutexes are currently
// broken.
//
// * recover is only partially implemented.  Also, the interpreter
// makes no attempt to distinguish target panics from interpreter
// crashes.
//
// * the sizes of the int, uint and uintptr types in the target
// program are assumed to be the same as those of the interpreter
// itself.
//
// * all values occupy space, even those of types defined by the spec
// to have zero size, e.g. struct{}.  This can cause asymptotic
// performance degradation.
//
// * os.Exit is implemented using panic, causing deferred functions to
// run.
package interp // import "golang.org/x/tools/go/ssa/interp"

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"os"
	"reflect"
	"runtime"
	"slices"
	"sync/atomic"
	_ "unsafe"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/internal/typeparams"
)

type continuation int

const (
	kNext continuation = iota
	kReturn
	kJump
)

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover Mode = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	EnableTracing                   // Print a trace of all instructions as they are interpreted.
)

type methodSet map[string]*ssa.Function

// State shared between all interpreted goroutines.
type interpreter struct {
	osArgs             []value                // the value of os.Args
	prog               *ssa.Program           // the SSA program
	globals            map[*ssa.Global]*value // addresses of global variables (immutable)
	mode               Mode                   // interpreter options
	reflectPackage     *ssa.Package           // the fake reflect package
	errorMethods       methodSet              // the method set of reflect.error, which implements the error interface.
	rtypeMethods       methodSet              // the method set of rtype, which implements the reflect.Type interface.
	runtimeErrorString types.Type             // the runtime.errorString type
	sizes              types.Sizes            // the effective type-sizing function
	goroutines         int32                  // atomically updated
}

type deferred struct {
	fn    value
	args  []value
	instr *ssa.Defer
	tail  *deferred
}

type frame struct {
	i                *interpreter
	caller           *frame
	fn               *ssa.Function
	block, prevBlock *ssa.BasicBlock
	env              map[ssa.Value]value // dynamic values of SSA variables
	locals           []value
	defers           *deferred
	result           value
	panicking        bool
	panic            any
	phitemps         []value // temporaries for parallel phi assignment
}

func (fr *frame) get(key ssa.Value) value {
	switch key := key.(type) {
	case nil:
		// Hack; simplifies handling of optional attributes
		// such as ssa.Slice.{Low,High}.
		return nil
	case *ssa.Function, *ssa.Builtin:
		return key
	case *ssa.Const:
		return constValue(key)
	case *ssa.Global:
		if r, ok := fr.i.globals[key]; ok {
			return r
		}
	}
	if r, ok := fr.env[key]; ok {
		return r
	}
	panic(fmt.Sprintf("get: no value for %T: %v", key, key.Name()))
}

// runDefer runs a deferred call d.
// It always returns normally, but may set or clear fr.panic.
func (fr *frame) runDefer(d *deferred) {
	if fr.i.mode&EnableTracing != 0 {
		fmt.Fprintf(os.Stderr, "%s: invoking deferred function call\n",
			fr.i.prog.Fset.Position(d.instr.Pos()))
	}
	var ok bool
	defer func() {
		if !ok {
			// Deferred call created a new state of panic.
			fr.panicking = true
			fr.panic = recover()
		}
	}()
	call(fr.i, fr, d.instr.Pos(), d.fn, d.args)
	ok = true
}

// runDefers executes fr's deferred function calls in LIFO order.
//
// On entry, fr.panicking indicates a state of panic; if
// true, fr.panic contains the panic value.
//
// On completion, if a deferred call started a panic, or if no
// deferred call recovered from a previous state of panic, then
// runDefers itself panics after the last deferred call has run.
//
// If there was no initial state of panic, or it was recovered from,
// runDefers returns normally.
func (fr *frame) runDefers() {
	for d := fr.defers; d != nil; d = d.tail {
		fr.runDefer(d)
	}
	fr.defers = nil
	if fr.panicking {
		panic(fr.panic) // new panic, or still panicking
	}
}

// lookupMethod returns the method set for type typ, which may be one
// of the interpreter's fake types.
func lookupMethod(i *interpreter, typ types.Type, meth *types.Func) *ssa.Function {
	switch typ {
	case rtypeType:
		return i.rtypeMethods[meth.Id()]
	case errorType:
		return i.errorMethods[meth.Id()]
	}
	return i.prog.LookupMethod(typ, meth.Pkg(), meth.Name())
}

// visitInstr interprets a single ssa.Instruction within the activation
// record frame.  It returns a continuation value indicating where to
// read the next instruction from.
func visitInstr(fr *frame, instr ssa.Instruction) continuation {
	switch instr := instr.(type) {
	case *ssa.DebugRef:
		// no-op

	case *ssa.UnOp:
		fr.env[instr] = unop(instr, fr.get(instr.X))

	case *ssa.BinOp:
		fr.env[instr] = binop(instr.Op, instr.X.Type(), fr.get(instr.X), fr.get(instr.Y))

	case *ssa.Call:
		fn, args := prepareCall(fr, &instr.Call)
		fr.env[instr] = call(fr.i, fr, instr.Pos(), fn, args)

	case *ssa.ChangeInterface:
		fr.env[instr] = fr.get(instr.X)

	case *ssa.ChangeType:
		fr.env[instr] = fr.get(instr.X) // (can't fail)

	case *ssa.Convert:
		fr.env[instr] = conv(instr.Type(), instr.X.Type(), fr.get(instr.X))

	case *ssa.SliceToArrayPointer:
		fr.env[instr] = sliceToArrayPointer(instr.Type(), instr.X.Type(), fr.get(instr.X))

	case *ssa.MakeInterface:
		fr.env[instr] = iface{t: instr.X.Type(), v: fr.get(instr.X)}

	case *ssa.Extract:
		fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]

	case *ssa.Slice:
		fr.env[instr] = slice(fr.get(instr.X), fr.get(instr.Low), fr.get(instr.High), fr.get(instr.Max))

	case *ssa.Return:
		switch len(instr.Results) {
		case 0:
		case 1:
			fr.result = fr.get(instr.Results[0])
		default:
			var res []value
			for _, r := range instr.Results {
				res = append(res, fr.get(r))
			}
			fr.result = tuple(res)
		}
		fr.block = nil
		return kReturn

	case *ssa.RunDefers:
		fr.runDefers()

	case *ssa.Panic:
		panic(targetPanic{fr.get(instr.X)})

	case *ssa.Send:
		fr.get(instr.Chan).(chan value) <- fr.get(instr.X)

	case *ssa.Store:
		store(typeparams.MustDeref(instr.Addr.Type()), fr.get(instr.Addr).(*value), fr.get(instr.Val))

	case *ssa.If:
		succ := 1
		if fr.get(instr.Cond).(bool) {
			succ = 0
		}
		fr.prevBlock, fr.block = fr.block, fr.block.Succs[succ]
		return kJump

	case *ssa.Jump:
		fr.prevBlock, fr.block = fr.block, fr.block.Succs[0]
		return kJump

	case *ssa.Defer:
		fn, args := prepareCall(fr, &instr.Call)
		defers := &fr.defers
		if into := fr.get(instr.DeferStack); into != nil {
			defers = into.(**deferred)
		}
		*defers = &deferred{
			fn:    fn,
			args:  args,
			instr: instr,
			tail:  *defers,
		}

	case *ssa.Go:
		fn, args := prepareCall(fr, &instr.Call)
		atomic.AddInt32(&fr.i.goroutines, 1)
		go func() {
			call(fr.i, nil, instr.Pos(), fn, args)
			atomic.AddInt32(&fr.i.goroutines, -1)
		}()

	case *ssa.MakeChan:
		fr.env[instr] = make(chan value, asInt64(fr.get(instr.Size)))

	case *ssa.Alloc:
		var addr *value
		if instr.Heap {
			// new
			addr = new(value)
			fr.env[instr] = addr
		} else {
			// local
			addr = fr.env[instr].(*value)
		}
		*addr = zero(typeparams.MustDeref(instr.Type()))

	case *ssa.MakeSlice:
		slice := make([]value, asInt64(fr.get(instr.Cap)))
		tElt := instr.Type().Underlying().(*types.Slice).Elem()
		for i := range slice {
			slice[i] = zero(tElt)
		}
		fr.env[instr] = slice[:asInt64(fr.get(instr.Len))]

	case *ssa.MakeMap:
		var reserve int64
		if instr.Reserve != nil {
			reserve = asInt64(fr.get(instr.Reserve))
		}
		if !fitsInt(reserve, fr.i.sizes) {
			panic(fmt.Sprintf("ssa.MakeMap.Reserve value %d does not fit in int", reserve))
		}
		fr.env[instr] = makeMap(instr.Type().Underlying().(*types.Map).Key(), reserve)

	case *ssa.Range:
		fr.env[instr] = rangeIter(fr.get(instr.X), instr.X.Type())

	case *ssa.Next:
		fr.env[instr] = fr.get(instr.Iter).(iter).next()

	case *ssa.FieldAddr:
		fr.env[instr] = &(*fr.get(instr.X).(*value)).(structure)[instr.Field]

	case *ssa.Field:
		fr.env[instr] = fr.get(instr.X).(structure)[instr.Field]

	case *ssa.IndexAddr:
		x := fr.get(instr.X)
		idx := fr.get(instr.Index)
		switch x := x.(type) {
		case []value:
			fr.env[instr] = &x[asInt64(idx)]
		case *value: // *array
			fr.env[instr] = &(*x).(array)[asInt64(idx)]
		default:
			panic(fmt.Sprintf("unexpected x type in IndexAddr: %T", x))
		}

	case *ssa.Index:
		x := fr.get(instr.X)
		idx := fr.get(instr.Index)

		switch x := x.(type) {
		case array:
			fr.env[instr] = x[asInt64(idx)]
		case string:
			fr.env[instr] = x[asInt64(idx)]
		default:
			panic(fmt.Sprintf("unexpected x type in Index: %T", x))
		}

	case *ssa.Lookup:
		fr.env[instr] = lookup(instr, fr.get(instr.X), fr.get(instr.Index))

	case *ssa.MapUpdate:
		m := fr.get(instr.Map)
		key := fr.get(instr.Key)
		v := fr.get(instr.Value)
		switch m := m.(type) {
		case map[value]value:
			m[key] = v
		case *hashmap:
			m.insert(key.(hashable), v)
		default:
			panic(fmt.Sprintf("illegal map type: %T", m))
		}

	case *ssa.TypeAssert:
		fr.env[instr] = typeAssert(fr.i, instr, fr.get(instr.X).(iface))

	case *ssa.MakeClosure:
		var bindings []value
		for _, binding := range instr.Bindings {
			bindings = append(bindings, fr.get(binding))
		}
		fr.env[instr] = &closure{instr.Fn.(*ssa.Function), bindings}

	case *ssa.Phi:
		log.Fatal("unreachable") // phis are processed at block entry

	case *ssa.Select:
		var cases []reflect.SelectCase
		if !instr.Blocking {
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectDefault,
			})
		}
		for _, state := range instr.States {
			var dir reflect.SelectDir
			if state.Dir == types.RecvOnly {
				dir = reflect.SelectRecv
			} else {
				dir = reflect.SelectSend
			}
			var send reflect.Value
			if state.Send != nil {
				send = reflect.ValueOf(fr.get(state.Send))
			}
			cases = append(cases, reflect.SelectCase{
				Dir:  dir,
				Chan: reflect.ValueOf(fr.get(state.Chan)),
				Send: send,
			})
		}
		chosen, recv, recvOk := reflect.Select(cases)
		if !instr.Blocking {
			chosen-- // default case should have index -1.
		}
		r := tuple{chosen, recvOk}
		for i, st := range instr.States {
			if st.Dir == types.RecvOnly {
				var v value
				if i == chosen && recvOk {
					// No need to copy since send makes an unaliased copy.
					v = recv.Interface().(value)
				} else {
					v = zero(st.Chan.Type().Underlying().(*types.Chan).Elem())
				}
				r = append(r, v)
			}
		}
		fr.env[instr] = r

	default:
		panic(fmt.Sprintf("unexpected instruction: %T", instr))
	}

	// if val, ok := instr.(ssa.Value); ok {
	// 	fmt.Println(toString(fr.env[val])) // debugging
	// }

	return kNext
}

// prepareCall determines the function value and argument values for a
// function call in a Call, Go or Defer instruction, performing
// interface method lookup if needed.
func prepareCall(fr *frame, call *ssa.CallCommon) (fn value, args []value) {
	v := fr.get(call.Value)
	if call.Method == nil {
		// Function call.
		fn = v
	} else {
		// Interface method invocation.
		recv := v.(iface)
		if recv.t == nil {
			panic("method invoked on nil interface")
		}
		if f := lookupMethod(fr.i, recv.t, call.Method); f == nil {
			// Unreachable in well-typed programs.
			panic(fmt.Sprintf("method set for dynamic type %v does not contain %s", recv.t, call.Method))
		} else {
			fn = f
		}
		args = append(args, recv.v)
	}
	for _, arg := range call.Args {
		args = append(args, fr.get(arg))
	}
	return
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
func call(i *interpreter, caller *frame, callpos token.Pos, fn value, args []value) value {
	switch fn := fn.(type) {
	case *ssa.Function:
		if fn == nil {
			panic("call of nil function") // nil of func type
		}
		return callSSA(i, caller, callpos, fn, args, nil)
	case *closure:
		return callSSA(i, caller, callpos, fn.Fn, args, fn.Env)
	case *ssa.Builtin:
		return callBuiltin(caller, callpos, fn, args)
	}
	panic(fmt.Sprintf("cannot call %T", fn))
}

func loc(fset *token.FileSet, pos token.Pos) string {
	if pos == token.NoPos {
		return ""
	}
	return " at " + fset.Position(pos).String()
}

// callSSA interprets a call to function fn with arguments args,
// and lexical environment env, returning its result.
// callpos is the position of the callsite.
func callSSA(i *interpreter, caller *frame, callpos token.Pos, fn *ssa.Function, args []value, env []value) value {
	if i.mode&EnableTracing != 0 {
		fset := fn.Prog.Fset
		// TODO(adonovan): fix: loc() lies for external functions.
		fmt.Fprintf(os.Stderr, "Entering %s%s.\n", fn, loc(fset, fn.Pos()))
		suffix := ""
		if caller != nil {
			suffix = ", resuming " + caller.fn.String() + loc(fset, callpos)
		}
		defer fmt.Fprintf(os.Stderr, "Leaving %s%s.\n", fn, suffix)
	}
	fr := &frame{
		i:      i,
		caller: caller, // for panic/recover
		fn:     fn,
	}
	if fn.Parent() == nil {
		name := fn.String()
		if ext := externals[name]; ext != nil {
			if i.mode&EnableTracing != 0 {
				fmt.Fprintln(os.Stderr, "\t(external)")
			}
			return ext(fr, args)
		}
		if fn.Blocks == nil {
			panic("no code for function: " + name)
		}
	}

	// generic function body?
	if fn.TypeParams().Len() > 0 && len(fn.TypeArgs()) == 0 {
		panic("interp requires ssa.BuilderMode to include InstantiateGenerics to execute generics")
	}

	fr.env = make(map[ssa.Value]value)
	fr.block = fn.Blocks[0]
	fr.locals = make([]value, len(fn.Locals))
	for i, l := range fn.Locals {
		fr.locals[i] = zero(typeparams.MustDeref(l.Type()))
		fr.env[l] = &fr.locals[i]
	}
	for i, p := range fn.Params {
		fr.env[p] = args[i]
	}
	for i, fv := range fn.FreeVars {
		fr.env[fv] = env[i]
	}
	for fr.block != nil {
		runFrame(fr)
	}
	// Destroy the locals to avoid accidental use after return.
	for i := range fn.Locals {
		fr.locals[i] = bad{}
	}
	return fr.result
}

// runFrame executes SSA instructions starting at fr.block and
// continuing until a return, a panic, or a recovered panic.
//
// After a panic, runFrame panics.
//
// After a normal return, fr.result contains the result of the call
// and fr.block is nil.
//
// A recovered panic in a function without named return parameters
// (NRPs) becomes a normal return of the zero value of the function's
// result type.
//
// After a recovered panic in a function with NRPs, fr.result is
// undefined and fr.block contains the block at which to resume
// control.
func runFrame(fr *frame) {
	defer func() {
		if fr.block == nil {
			return // normal return
		}
		if fr.i.mode&DisableRecover != 0 {
			return // let interpreter crash
		}
		fr.panicking = true
		fr.panic = recover()
		if fr.i.mode&EnableTracing != 0 {
			fmt.Fprintf(os.Stderr, "Panicking: %T %v.\n", fr.panic, fr.panic)
		}
		fr.runDefers()
		fr.block = fr.fn.Recover
	}()

	for {
		if fr.i.mode&EnableTracing != 0 {
			fmt.Fprintf(os.Stderr, ".%s:\n", fr.block)
		}

		nonPhis := executePhis(fr)
		for _, instr := range nonPhis {
			if fr.i.mode&EnableTracing != 0 {
				if v, ok := instr.(ssa.Value); ok {
					fmt.Fprintln(os.Stderr, "\t", v.Name(), "=", instr)
				} else {
					fmt.Fprintln(os.Stderr, "\t", instr)
				}
			}
			if visitInstr(fr, instr) == kReturn {
				return
			}
			// Inv: kNext (continue) or kJump (last instr)
		}
	}
}

// executePhis executes the phi-nodes at the start of the current
// block and returns the non-phi instructions.
func executePhis(fr *frame) []ssa.Instruction {
	firstNonPhi := -1
	for i, instr := range fr.block.Instrs {
		if _, ok := instr.(*ssa.Phi); !ok {
			firstNonPhi = i
			break
		}
	}
	// Inv: 0 <= firstNonPhi; every block contains a non-phi.

	nonPhis := fr.block.Instrs[firstNonPhi:]
	if firstNonPhi > 0 {
		phis := fr.block.Instrs[:firstNonPhi]
		// Execute parallel assignment of phis.
		//
		// See "the swap problem" in Briggs et al's "Practical Improvements
		// to the Construction and Destruction of SSA Form" for discussion.
		predIndex := slices.Index(fr.block.Preds, fr.prevBlock)
		fr.phitemps = fr.phitemps[:0]
		for _, phi := range phis {
			phi := phi.(*ssa.Phi)
			if fr.i.mode&EnableTracing != 0 {
				fmt.Fprintln(os.Stderr, "\t", phi.Name(), "=", phi)
			}
			fr.phitemps = append(fr.phitemps, fr.get(phi.Edges[predIndex]))
		}
		for i, phi := range phis {
			fr.env[phi.(*ssa.Phi)] = fr.phitemps[i]
		}
	}
	return nonPhis
}

// doRecover implements the recover() built-in.
func doRecover(caller *frame) value {
	// recover() must be exactly one level beneath the deferred
	// function (two levels beneath the panicking function) to
	// have any effect.  Thus we ignore both "defer recover()" and
	// "defer f() -> g() -> recover()".
	if caller.i.mode&DisableRecover == 0 &&
		caller != nil && !caller.panicking &&
		caller.caller != nil && caller.caller.panicking {
		caller.caller.panicking = false
		p := caller.caller.panic
		caller.caller.panic = nil

		// TODO(adonovan): support runtime.Goexit.
		switch p := p.(type) {
		case targetPanic:
			// The target program explicitly called panic().
			return p.v
		case runtime.Error:
			// The interpreter encountered a runtime error.
			return iface{caller.i.runtimeErrorString, p.Error()}
		case string:
			// The interpreter explicitly called panic().
			return iface{caller.i.runtimeErrorString, p}
		default:
			panic(fmt.Sprintf("unexpected panic type %T in target call to recover()", p))
		}
	}
	return iface{}
}

// Interpret interprets the Go program whose main package is mainpkg.
// mode specifies various interpreter options.  filename and args are
// the initial values of os.Args for the target program.  sizes is the
// effective type-sizing function for this program.
//
// Interpret returns the exit code of the program: 2 for panic (like
// gc does), or the argument to os.Exit for normal termination.
//
// The SSA program must include the "runtime" package.
//
// Type parameterized functions must have been built with
// InstantiateGenerics in the ssa.BuilderMode to be interpreted.
func Interpret(mainpkg *ssa.Package, mode Mode, sizes types.Sizes, filename string, args []string) (exitCode int) {
	i := &interpreter{
		prog:       mainpkg.Prog,
		globals:    make(map[*ssa.Global]*value),
		mode:       mode,
		sizes:      sizes,
		goroutines: 1,
	}
	runtimePkg := i.prog.ImportedPackage("runtime")
	if runtimePkg == nil {
		panic("ssa.Program doesn't include runtime package")
	}
	i.runtimeErrorString = runtimePkg.Type("errorString").Object().Type()

	initReflect(i)

	i.osArgs = append(i.osArgs, filename)
	for _, arg := range args {
		i.osArgs = append(i.osArgs, arg)
	}

	for _, pkg := range i.prog.AllPackages() {
		// Initialize global storage.
		for _, m := range pkg.Members {
			switch v := m.(type) {
			case *ssa.Global:
				cell := zero(typeparams.MustDeref(v.Type()))
				i.globals[v] = &cell
			}
		}
	}

	// Top-level error handler.
	exitCode = 2
	defer func() {
		if exitCode != 2 || i.mode&DisableRecover != 0 {
			return
		}
		switch p := recover().(type) {
		case exitPanic:
			exitCode = int(p)
			return
		case targetPanic:
			fmt.Fprintln(os.Stderr, "panic:", toString(p.v))
		case runtime.Error:
			fmt.Fprintln(os.Stderr, "panic:", p.Error())
		case string:
			fmt.Fprintln(os.Stderr, "panic:", p)
		default:
			fmt.Fprintf(os.Stderr, "panic: unexpected type: %T: %v\n", p, p)
		}

		// TODO(adonovan): dump panicking interpreter goroutine?
		// buf := make([]byte, 0x10000)
		// runtime.Stack(buf, false)
		// fmt.Fprintln(os.Stderr, string(buf))
		// (Or dump panicking target goroutine?)
	}()

	// Run!
	call(i, nil, token.NoPos, mainpkg.Func("init"), nil)
	if mainFn := mainpkg.Func("main"); mainFn != nil {
		call(i, nil, token.NoPos, mainFn, nil)
		exitCode = 0
	} else {
		fmt.Fprintln(os.Stderr, "No main function.")
		exitCode = 1
	}
	return
}
```

## File: go/ssa/interp/map.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

// Custom hashtable atop map.
// For use when the key's equivalence relation is not consistent with ==.

// The Go specification doesn't address the atomicity of map operations.
// The FAQ states that an implementation is permitted to crash on
// concurrent map access.

import (
	"go/types"
)

type hashable interface {
	hash(t types.Type) int
	eq(t types.Type, x any) bool
}

type entry struct {
	key   hashable
	value value
	next  *entry
}

// A hashtable atop the built-in map.  Since each bucket contains
// exactly one hash value, there's no need to perform hash-equality
// tests when walking the linked list.  Rehashing is done by the
// underlying map.
type hashmap struct {
	keyType types.Type
	table   map[int]*entry
	length  int // number of entries in map
}

// makeMap returns an empty initialized map of key type kt,
// preallocating space for reserve elements.
func makeMap(kt types.Type, reserve int64) value {
	if usesBuiltinMap(kt) {
		return make(map[value]value, reserve)
	}
	return &hashmap{keyType: kt, table: make(map[int]*entry, reserve)}
}

// delete removes the association for key k, if any.
func (m *hashmap) delete(k hashable) {
	if m != nil {
		hash := k.hash(m.keyType)
		head := m.table[hash]
		if head != nil {
			if k.eq(m.keyType, head.key) {
				m.table[hash] = head.next
				m.length--
				return
			}
			prev := head
			for e := head.next; e != nil; e = e.next {
				if k.eq(m.keyType, e.key) {
					prev.next = e.next
					m.length--
					return
				}
				prev = e
			}
		}
	}
}

// lookup returns the value associated with key k, if present, or
// value(nil) otherwise.
func (m *hashmap) lookup(k hashable) value {
	if m != nil {
		hash := k.hash(m.keyType)
		for e := m.table[hash]; e != nil; e = e.next {
			if k.eq(m.keyType, e.key) {
				return e.value
			}
		}
	}
	return nil
}

// insert updates the map to associate key k with value v.  If there
// was already an association for an eq() (though not necessarily ==)
// k, the previous key remains in the map and its associated value is
// updated.
func (m *hashmap) insert(k hashable, v value) {
	hash := k.hash(m.keyType)
	head := m.table[hash]
	for e := head; e != nil; e = e.next {
		if k.eq(m.keyType, e.key) {
			e.value = v
			return
		}
	}
	m.table[hash] = &entry{
		key:   k,
		value: v,
		next:  head,
	}
	m.length++
}

// len returns the number of key/value associations in the map.
func (m *hashmap) len() int {
	if m != nil {
		return m.length
	}
	return 0
}

// entries returns a rangeable map of entries.
func (m *hashmap) entries() map[int]*entry {
	if m != nil {
		return m.table
	}
	return nil
}
```

## File: go/ssa/interp/ops.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

import (
	"bytes"
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/internal/typeparams"
)

// If the target program panics, the interpreter panics with this type.
type targetPanic struct {
	v value
}

func (p targetPanic) String() string {
	return toString(p.v)
}

// If the target program calls exit, the interpreter panics with this type.
type exitPanic int

// constValue returns the value of the constant with the
// dynamic type tag appropriate for c.Type().
func constValue(c *ssa.Const) value {
	if c.Value == nil {
		return zero(c.Type()) // typed zero
	}
	// c is not a type parameter so it's underlying type is basic.

	if t, ok := c.Type().Underlying().(*types.Basic); ok {
		// TODO(adonovan): eliminate untyped constants from SSA form.
		switch t.Kind() {
		case types.Bool, types.UntypedBool:
			return constant.BoolVal(c.Value)
		case types.Int, types.UntypedInt:
			// Assume sizeof(int) is same on host and target.
			return int(c.Int64())
		case types.Int8:
			return int8(c.Int64())
		case types.Int16:
			return int16(c.Int64())
		case types.Int32, types.UntypedRune:
			return int32(c.Int64())
		case types.Int64:
			return c.Int64()
		case types.Uint:
			// Assume sizeof(uint) is same on host and target.
			return uint(c.Uint64())
		case types.Uint8:
			return uint8(c.Uint64())
		case types.Uint16:
			return uint16(c.Uint64())
		case types.Uint32:
			return uint32(c.Uint64())
		case types.Uint64:
			return c.Uint64()
		case types.Uintptr:
			// Assume sizeof(uintptr) is same on host and target.
			return uintptr(c.Uint64())
		case types.Float32:
			return float32(c.Float64())
		case types.Float64, types.UntypedFloat:
			return c.Float64()
		case types.Complex64:
			return complex64(c.Complex128())
		case types.Complex128, types.UntypedComplex:
			return c.Complex128()
		case types.String, types.UntypedString:
			if c.Value.Kind() == constant.String {
				return constant.StringVal(c.Value)
			}
			return string(rune(c.Int64()))
		}
	}

	panic(fmt.Sprintf("constValue: %s", c))
}

// fitsInt returns true if x fits in type int according to sizes.
func fitsInt(x int64, sizes types.Sizes) bool {
	intSize := sizes.Sizeof(types.Typ[types.Int])
	if intSize < sizes.Sizeof(types.Typ[types.Int64]) {
		maxInt := int64(1)<<((intSize*8)-1) - 1
		minInt := -int64(1) << ((intSize * 8) - 1)
		return minInt <= x && x <= maxInt
	}
	return true
}

// asInt64 converts x, which must be an integer, to an int64.
//
// Callers that need a value directly usable as an int should combine this with fitsInt().
func asInt64(x value) int64 {
	switch x := x.(type) {
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x)
	case uintptr:
		return int64(x)
	}
	panic(fmt.Sprintf("cannot convert %T to int64", x))
}

// asUint64 converts x, which must be an unsigned integer, to a uint64
// suitable for use as a bitwise shift count.
func asUint64(x value) uint64 {
	switch x := x.(type) {
	case uint:
		return uint64(x)
	case uint8:
		return uint64(x)
	case uint16:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint64:
		return x
	case uintptr:
		return uint64(x)
	}
	panic(fmt.Sprintf("cannot convert %T to uint64", x))
}

// asUnsigned returns the value of x, which must be an integer type, as its equivalent unsigned type,
// and returns true if x is non-negative.
func asUnsigned(x value) (value, bool) {
	switch x := x.(type) {
	case int:
		return uint(x), x >= 0
	case int8:
		return uint8(x), x >= 0
	case int16:
		return uint16(x), x >= 0
	case int32:
		return uint32(x), x >= 0
	case int64:
		return uint64(x), x >= 0
	case uint, uint8, uint32, uint64, uintptr:
		return x, true
	}
	panic(fmt.Sprintf("cannot convert %T to unsigned", x))
}

// zero returns a new "zero" value of the specified type.
func zero(t types.Type) value {
	switch t := t.(type) {
	case *types.Basic:
		if t.Kind() == types.UntypedNil {
			panic("untyped nil has no zero value")
		}
		if t.Info()&types.IsUntyped != 0 {
			// TODO(adonovan): make it an invariant that
			// this is unreachable.  Currently some
			// constants have 'untyped' types when they
			// should be defaulted by the typechecker.
			t = types.Default(t).(*types.Basic)
		}
		switch t.Kind() {
		case types.Bool:
			return false
		case types.Int:
			return int(0)
		case types.Int8:
			return int8(0)
		case types.Int16:
			return int16(0)
		case types.Int32:
			return int32(0)
		case types.Int64:
			return int64(0)
		case types.Uint:
			return uint(0)
		case types.Uint8:
			return uint8(0)
		case types.Uint16:
			return uint16(0)
		case types.Uint32:
			return uint32(0)
		case types.Uint64:
			return uint64(0)
		case types.Uintptr:
			return uintptr(0)
		case types.Float32:
			return float32(0)
		case types.Float64:
			return float64(0)
		case types.Complex64:
			return complex64(0)
		case types.Complex128:
			return complex128(0)
		case types.String:
			return ""
		case types.UnsafePointer:
			return unsafe.Pointer(nil)
		default:
			panic(fmt.Sprint("zero for unexpected type:", t))
		}
	case *types.Pointer:
		return (*value)(nil)
	case *types.Array:
		a := make(array, t.Len())
		for i := range a {
			a[i] = zero(t.Elem())
		}
		return a
	case *types.Named:
		return zero(t.Underlying())
	case *types.Alias:
		return zero(types.Unalias(t))
	case *types.Interface:
		return iface{} // nil type, methodset and value
	case *types.Slice:
		return []value(nil)
	case *types.Struct:
		s := make(structure, t.NumFields())
		for i := range s {
			s[i] = zero(t.Field(i).Type())
		}
		return s
	case *types.Tuple:
		if t.Len() == 1 {
			return zero(t.At(0).Type())
		}
		s := make(tuple, t.Len())
		for i := range s {
			s[i] = zero(t.At(i).Type())
		}
		return s
	case *types.Chan:
		return chan value(nil)
	case *types.Map:
		if usesBuiltinMap(t.Key()) {
			return map[value]value(nil)
		}
		return (*hashmap)(nil)
	case *types.Signature:
		return (*ssa.Function)(nil)
	}
	panic(fmt.Sprint("zero: unexpected ", t))
}

// slice returns x[lo:hi:max].  Any of lo, hi and max may be nil.
func slice(x, lo, hi, max value) value {
	var Len, Cap int
	switch x := x.(type) {
	case string:
		Len = len(x)
	case []value:
		Len = len(x)
		Cap = cap(x)
	case *value: // *array
		a := (*x).(array)
		Len = len(a)
		Cap = cap(a)
	}

	l := int64(0)
	if lo != nil {
		l = asInt64(lo)
	}

	h := int64(Len)
	if hi != nil {
		h = asInt64(hi)
	}

	m := int64(Cap)
	if max != nil {
		m = asInt64(max)
	}

	switch x := x.(type) {
	case string:
		return x[l:h]
	case []value:
		return x[l:h:m]
	case *value: // *array
		a := (*x).(array)
		return []value(a)[l:h:m]
	}
	panic(fmt.Sprintf("slice: unexpected X type: %T", x))
}

// lookup returns x[idx] where x is a map.
func lookup(instr *ssa.Lookup, x, idx value) value {
	switch x := x.(type) { // map or string
	case map[value]value, *hashmap:
		var v value
		var ok bool
		switch x := x.(type) {
		case map[value]value:
			v, ok = x[idx]
		case *hashmap:
			v = x.lookup(idx.(hashable))
			ok = v != nil
		}
		if !ok {
			v = zero(instr.X.Type().Underlying().(*types.Map).Elem())
		}
		if instr.CommaOk {
			v = tuple{v, ok}
		}
		return v
	}
	panic(fmt.Sprintf("unexpected x type in Lookup: %T", x))
}

// binop implements all arithmetic and logical binary operators for
// numeric datatypes and strings.  Both operands must have identical
// dynamic type.
func binop(op token.Token, t types.Type, x, y value) value {
	switch op {
	case token.ADD:
		switch x.(type) {
		case int:
			return x.(int) + y.(int)
		case int8:
			return x.(int8) + y.(int8)
		case int16:
			return x.(int16) + y.(int16)
		case int32:
			return x.(int32) + y.(int32)
		case int64:
			return x.(int64) + y.(int64)
		case uint:
			return x.(uint) + y.(uint)
		case uint8:
			return x.(uint8) + y.(uint8)
		case uint16:
			return x.(uint16) + y.(uint16)
		case uint32:
			return x.(uint32) + y.(uint32)
		case uint64:
			return x.(uint64) + y.(uint64)
		case uintptr:
			return x.(uintptr) + y.(uintptr)
		case float32:
			return x.(float32) + y.(float32)
		case float64:
			return x.(float64) + y.(float64)
		case complex64:
			return x.(complex64) + y.(complex64)
		case complex128:
			return x.(complex128) + y.(complex128)
		case string:
			return x.(string) + y.(string)
		}

	case token.SUB:
		switch x.(type) {
		case int:
			return x.(int) - y.(int)
		case int8:
			return x.(int8) - y.(int8)
		case int16:
			return x.(int16) - y.(int16)
		case int32:
			return x.(int32) - y.(int32)
		case int64:
			return x.(int64) - y.(int64)
		case uint:
			return x.(uint) - y.(uint)
		case uint8:
			return x.(uint8) - y.(uint8)
		case uint16:
			return x.(uint16) - y.(uint16)
		case uint32:
			return x.(uint32) - y.(uint32)
		case uint64:
			return x.(uint64) - y.(uint64)
		case uintptr:
			return x.(uintptr) - y.(uintptr)
		case float32:
			return x.(float32) - y.(float32)
		case float64:
			return x.(float64) - y.(float64)
		case complex64:
			return x.(complex64) - y.(complex64)
		case complex128:
			return x.(complex128) - y.(complex128)
		}

	case token.MUL:
		switch x.(type) {
		case int:
			return x.(int) * y.(int)
		case int8:
			return x.(int8) * y.(int8)
		case int16:
			return x.(int16) * y.(int16)
		case int32:
			return x.(int32) * y.(int32)
		case int64:
			return x.(int64) * y.(int64)
		case uint:
			return x.(uint) * y.(uint)
		case uint8:
			return x.(uint8) * y.(uint8)
		case uint16:
			return x.(uint16) * y.(uint16)
		case uint32:
			return x.(uint32) * y.(uint32)
		case uint64:
			return x.(uint64) * y.(uint64)
		case uintptr:
			return x.(uintptr) * y.(uintptr)
		case float32:
			return x.(float32) * y.(float32)
		case float64:
			return x.(float64) * y.(float64)
		case complex64:
			return x.(complex64) * y.(complex64)
		case complex128:
			return x.(complex128) * y.(complex128)
		}

	case token.QUO:
		switch x.(type) {
		case int:
			return x.(int) / y.(int)
		case int8:
			return x.(int8) / y.(int8)
		case int16:
			return x.(int16) / y.(int16)
		case int32:
			return x.(int32) / y.(int32)
		case int64:
			return x.(int64) / y.(int64)
		case uint:
			return x.(uint) / y.(uint)
		case uint8:
			return x.(uint8) / y.(uint8)
		case uint16:
			return x.(uint16) / y.(uint16)
		case uint32:
			return x.(uint32) / y.(uint32)
		case uint64:
			return x.(uint64) / y.(uint64)
		case uintptr:
			return x.(uintptr) / y.(uintptr)
		case float32:
			return x.(float32) / y.(float32)
		case float64:
			return x.(float64) / y.(float64)
		case complex64:
			return x.(complex64) / y.(complex64)
		case complex128:
			return x.(complex128) / y.(complex128)
		}

	case token.REM:
		switch x.(type) {
		case int:
			return x.(int) % y.(int)
		case int8:
			return x.(int8) % y.(int8)
		case int16:
			return x.(int16) % y.(int16)
		case int32:
			return x.(int32) % y.(int32)
		case int64:
			return x.(int64) % y.(int64)
		case uint:
			return x.(uint) % y.(uint)
		case uint8:
			return x.(uint8) % y.(uint8)
		case uint16:
			return x.(uint16) % y.(uint16)
		case uint32:
			return x.(uint32) % y.(uint32)
		case uint64:
			return x.(uint64) % y.(uint64)
		case uintptr:
			return x.(uintptr) % y.(uintptr)
		}

	case token.AND:
		switch x.(type) {
		case int:
			return x.(int) & y.(int)
		case int8:
			return x.(int8) & y.(int8)
		case int16:
			return x.(int16) & y.(int16)
		case int32:
			return x.(int32) & y.(int32)
		case int64:
			return x.(int64) & y.(int64)
		case uint:
			return x.(uint) & y.(uint)
		case uint8:
			return x.(uint8) & y.(uint8)
		case uint16:
			return x.(uint16) & y.(uint16)
		case uint32:
			return x.(uint32) & y.(uint32)
		case uint64:
			return x.(uint64) & y.(uint64)
		case uintptr:
			return x.(uintptr) & y.(uintptr)
		}

	case token.OR:
		switch x.(type) {
		case int:
			return x.(int) | y.(int)
		case int8:
			return x.(int8) | y.(int8)
		case int16:
			return x.(int16) | y.(int16)
		case int32:
			return x.(int32) | y.(int32)
		case int64:
			return x.(int64) | y.(int64)
		case uint:
			return x.(uint) | y.(uint)
		case uint8:
			return x.(uint8) | y.(uint8)
		case uint16:
			return x.(uint16) | y.(uint16)
		case uint32:
			return x.(uint32) | y.(uint32)
		case uint64:
			return x.(uint64) | y.(uint64)
		case uintptr:
			return x.(uintptr) | y.(uintptr)
		}

	case token.XOR:
		switch x.(type) {
		case int:
			return x.(int) ^ y.(int)
		case int8:
			return x.(int8) ^ y.(int8)
		case int16:
			return x.(int16) ^ y.(int16)
		case int32:
			return x.(int32) ^ y.(int32)
		case int64:
			return x.(int64) ^ y.(int64)
		case uint:
			return x.(uint) ^ y.(uint)
		case uint8:
			return x.(uint8) ^ y.(uint8)
		case uint16:
			return x.(uint16) ^ y.(uint16)
		case uint32:
			return x.(uint32) ^ y.(uint32)
		case uint64:
			return x.(uint64) ^ y.(uint64)
		case uintptr:
			return x.(uintptr) ^ y.(uintptr)
		}

	case token.AND_NOT:
		switch x.(type) {
		case int:
			return x.(int) &^ y.(int)
		case int8:
			return x.(int8) &^ y.(int8)
		case int16:
			return x.(int16) &^ y.(int16)
		case int32:
			return x.(int32) &^ y.(int32)
		case int64:
			return x.(int64) &^ y.(int64)
		case uint:
			return x.(uint) &^ y.(uint)
		case uint8:
			return x.(uint8) &^ y.(uint8)
		case uint16:
			return x.(uint16) &^ y.(uint16)
		case uint32:
			return x.(uint32) &^ y.(uint32)
		case uint64:
			return x.(uint64) &^ y.(uint64)
		case uintptr:
			return x.(uintptr) &^ y.(uintptr)
		}

	case token.SHL:
		u, ok := asUnsigned(y)
		if !ok {
			panic("negative shift amount")
		}
		y := asUint64(u)
		switch x.(type) {
		case int:
			return x.(int) << y
		case int8:
			return x.(int8) << y
		case int16:
			return x.(int16) << y
		case int32:
			return x.(int32) << y
		case int64:
			return x.(int64) << y
		case uint:
			return x.(uint) << y
		case uint8:
			return x.(uint8) << y
		case uint16:
			return x.(uint16) << y
		case uint32:
			return x.(uint32) << y
		case uint64:
			return x.(uint64) << y
		case uintptr:
			return x.(uintptr) << y
		}

	case token.SHR:
		u, ok := asUnsigned(y)
		if !ok {
			panic("negative shift amount")
		}
		y := asUint64(u)
		switch x.(type) {
		case int:
			return x.(int) >> y
		case int8:
			return x.(int8) >> y
		case int16:
			return x.(int16) >> y
		case int32:
			return x.(int32) >> y
		case int64:
			return x.(int64) >> y
		case uint:
			return x.(uint) >> y
		case uint8:
			return x.(uint8) >> y
		case uint16:
			return x.(uint16) >> y
		case uint32:
			return x.(uint32) >> y
		case uint64:
			return x.(uint64) >> y
		case uintptr:
			return x.(uintptr) >> y
		}

	case token.LSS:
		switch x.(type) {
		case int:
			return x.(int) < y.(int)
		case int8:
			return x.(int8) < y.(int8)
		case int16:
			return x.(int16) < y.(int16)
		case int32:
			return x.(int32) < y.(int32)
		case int64:
			return x.(int64) < y.(int64)
		case uint:
			return x.(uint) < y.(uint)
		case uint8:
			return x.(uint8) < y.(uint8)
		case uint16:
			return x.(uint16) < y.(uint16)
		case uint32:
			return x.(uint32) < y.(uint32)
		case uint64:
			return x.(uint64) < y.(uint64)
		case uintptr:
			return x.(uintptr) < y.(uintptr)
		case float32:
			return x.(float32) < y.(float32)
		case float64:
			return x.(float64) < y.(float64)
		case string:
			return x.(string) < y.(string)
		}

	case token.LEQ:
		switch x.(type) {
		case int:
			return x.(int) <= y.(int)
		case int8:
			return x.(int8) <= y.(int8)
		case int16:
			return x.(int16) <= y.(int16)
		case int32:
			return x.(int32) <= y.(int32)
		case int64:
			return x.(int64) <= y.(int64)
		case uint:
			return x.(uint) <= y.(uint)
		case uint8:
			return x.(uint8) <= y.(uint8)
		case uint16:
			return x.(uint16) <= y.(uint16)
		case uint32:
			return x.(uint32) <= y.(uint32)
		case uint64:
			return x.(uint64) <= y.(uint64)
		case uintptr:
			return x.(uintptr) <= y.(uintptr)
		case float32:
			return x.(float32) <= y.(float32)
		case float64:
			return x.(float64) <= y.(float64)
		case string:
			return x.(string) <= y.(string)
		}

	case token.EQL:
		return eqnil(t, x, y)

	case token.NEQ:
		return !eqnil(t, x, y)

	case token.GTR:
		switch x.(type) {
		case int:
			return x.(int) > y.(int)
		case int8:
			return x.(int8) > y.(int8)
		case int16:
			return x.(int16) > y.(int16)
		case int32:
			return x.(int32) > y.(int32)
		case int64:
			return x.(int64) > y.(int64)
		case uint:
			return x.(uint) > y.(uint)
		case uint8:
			return x.(uint8) > y.(uint8)
		case uint16:
			return x.(uint16) > y.(uint16)
		case uint32:
			return x.(uint32) > y.(uint32)
		case uint64:
			return x.(uint64) > y.(uint64)
		case uintptr:
			return x.(uintptr) > y.(uintptr)
		case float32:
			return x.(float32) > y.(float32)
		case float64:
			return x.(float64) > y.(float64)
		case string:
			return x.(string) > y.(string)
		}

	case token.GEQ:
		switch x.(type) {
		case int:
			return x.(int) >= y.(int)
		case int8:
			return x.(int8) >= y.(int8)
		case int16:
			return x.(int16) >= y.(int16)
		case int32:
			return x.(int32) >= y.(int32)
		case int64:
			return x.(int64) >= y.(int64)
		case uint:
			return x.(uint) >= y.(uint)
		case uint8:
			return x.(uint8) >= y.(uint8)
		case uint16:
			return x.(uint16) >= y.(uint16)
		case uint32:
			return x.(uint32) >= y.(uint32)
		case uint64:
			return x.(uint64) >= y.(uint64)
		case uintptr:
			return x.(uintptr) >= y.(uintptr)
		case float32:
			return x.(float32) >= y.(float32)
		case float64:
			return x.(float64) >= y.(float64)
		case string:
			return x.(string) >= y.(string)
		}
	}
	panic(fmt.Sprintf("invalid binary op: %T %s %T", x, op, y))
}

// eqnil returns the comparison x == y using the equivalence relation
// appropriate for type t.
// If t is a reference type, at most one of x or y may be a nil value
// of that type.
func eqnil(t types.Type, x, y value) bool {
	switch t.Underlying().(type) {
	case *types.Map, *types.Signature, *types.Slice:
		// Since these types don't support comparison,
		// one of the operands must be a literal nil.
		switch x := x.(type) {
		case *hashmap:
			return (x != nil) == (y.(*hashmap) != nil)
		case map[value]value:
			return (x != nil) == (y.(map[value]value) != nil)
		case *ssa.Function:
			switch y := y.(type) {
			case *ssa.Function:
				return (x != nil) == (y != nil)
			case *closure:
				return true
			}
		case *closure:
			return (x != nil) == (y.(*ssa.Function) != nil)
		case []value:
			return (x != nil) == (y.([]value) != nil)
		}
		panic(fmt.Sprintf("eqnil(%s): illegal dynamic type: %T", t, x))
	}

	return equals(t, x, y)
}

func unop(instr *ssa.UnOp, x value) value {
	switch instr.Op {
	case token.ARROW: // receive
		v, ok := <-x.(chan value)
		if !ok {
			v = zero(instr.X.Type().Underlying().(*types.Chan).Elem())
		}
		if instr.CommaOk {
			v = tuple{v, ok}
		}
		return v
	case token.SUB:
		switch x := x.(type) {
		case int:
			return -x
		case int8:
			return -x
		case int16:
			return -x
		case int32:
			return -x
		case int64:
			return -x
		case uint:
			return -x
		case uint8:
			return -x
		case uint16:
			return -x
		case uint32:
			return -x
		case uint64:
			return -x
		case uintptr:
			return -x
		case float32:
			return -x
		case float64:
			return -x
		case complex64:
			return -x
		case complex128:
			return -x
		}
	case token.MUL:
		return load(typeparams.MustDeref(instr.X.Type()), x.(*value))
	case token.NOT:
		return !x.(bool)
	case token.XOR:
		switch x := x.(type) {
		case int:
			return ^x
		case int8:
			return ^x
		case int16:
			return ^x
		case int32:
			return ^x
		case int64:
			return ^x
		case uint:
			return ^x
		case uint8:
			return ^x
		case uint16:
			return ^x
		case uint32:
			return ^x
		case uint64:
			return ^x
		case uintptr:
			return ^x
		}
	}
	panic(fmt.Sprintf("invalid unary op %s %T", instr.Op, x))
}

// typeAssert checks whether dynamic type of itf is instr.AssertedType.
// It returns the extracted value on success, and panics on failure,
// unless instr.CommaOk, in which case it always returns a "value,ok" tuple.
func typeAssert(i *interpreter, instr *ssa.TypeAssert, itf iface) value {
	var v value
	err := ""
	if itf.t == nil {
		err = fmt.Sprintf("interface conversion: interface is nil, not %s", instr.AssertedType)

	} else if idst, ok := instr.AssertedType.Underlying().(*types.Interface); ok {
		v = itf
		err = checkInterface(i, idst, itf)

	} else if types.Identical(itf.t, instr.AssertedType) {
		v = itf.v // extract value

	} else {
		err = fmt.Sprintf("interface conversion: interface is %s, not %s", itf.t, instr.AssertedType)
	}
	// Note: if instr.Underlying==true ever becomes reachable from interp check that
	// types.Identical(itf.t.Underlying(), instr.AssertedType)

	if err != "" {
		if !instr.CommaOk {
			panic(err)
		}
		return tuple{zero(instr.AssertedType), false}
	}
	if instr.CommaOk {
		return tuple{v, true}
	}
	return v
}

// This variable is no longer used but remains to prevent build breakage.
var CapturedOutput *bytes.Buffer

// callBuiltin interprets a call to builtin fn with arguments args,
// returning its result.
func callBuiltin(caller *frame, callpos token.Pos, fn *ssa.Builtin, args []value) value {
	switch fn.Name() {
	case "append":
		if len(args) == 1 {
			return args[0]
		}
		if s, ok := args[1].(string); ok {
			// append([]byte, ...string) []byte
			arg0 := args[0].([]value)
			for i := 0; i < len(s); i++ {
				arg0 = append(arg0, s[i])
			}
			return arg0
		}
		// append([]T, ...[]T) []T
		return append(args[0].([]value), args[1].([]value)...)

	case "copy": // copy([]T, []T) int or copy([]byte, string) int
		src := args[1]
		if _, ok := src.(string); ok {
			params := fn.Type().(*types.Signature).Params()
			src = conv(params.At(0).Type(), params.At(1).Type(), src)
		}
		return copy(args[0].([]value), src.([]value))

	case "close": // close(chan T)
		close(args[0].(chan value))
		return nil

	case "delete": // delete(map[K]value, K)
		switch m := args[0].(type) {
		case map[value]value:
			delete(m, args[1])
		case *hashmap:
			m.delete(args[1].(hashable))
		default:
			panic(fmt.Sprintf("illegal map type: %T", m))
		}
		return nil

	case "print", "println": // print(any, ...)
		ln := fn.Name() == "println"
		var buf bytes.Buffer
		for i, arg := range args {
			if i > 0 && ln {
				buf.WriteRune(' ')
			}
			buf.WriteString(toString(arg))
		}
		if ln {
			buf.WriteRune('\n')
		}
		os.Stderr.Write(buf.Bytes())
		return nil

	case "len":
		switch x := args[0].(type) {
		case string:
			return len(x)
		case array:
			return len(x)
		case *value:
			return len((*x).(array))
		case []value:
			return len(x)
		case map[value]value:
			return len(x)
		case *hashmap:
			return x.len()
		case chan value:
			return len(x)
		default:
			panic(fmt.Sprintf("len: illegal operand: %T", x))
		}

	case "cap":
		switch x := args[0].(type) {
		case array:
			return cap(x)
		case *value:
			return cap((*x).(array))
		case []value:
			return cap(x)
		case chan value:
			return cap(x)
		default:
			panic(fmt.Sprintf("cap: illegal operand: %T", x))
		}

	case "min":
		return foldLeft(min, args)
	case "max":
		return foldLeft(max, args)

	case "real":
		switch c := args[0].(type) {
		case complex64:
			return real(c)
		case complex128:
			return real(c)
		default:
			panic(fmt.Sprintf("real: illegal operand: %T", c))
		}

	case "imag":
		switch c := args[0].(type) {
		case complex64:
			return imag(c)
		case complex128:
			return imag(c)
		default:
			panic(fmt.Sprintf("imag: illegal operand: %T", c))
		}

	case "complex":
		switch f := args[0].(type) {
		case float32:
			return complex(f, args[1].(float32))
		case float64:
			return complex(f, args[1].(float64))
		default:
			panic(fmt.Sprintf("complex: illegal operand: %T", f))
		}

	case "panic":
		// ssa.Panic handles most cases; this is only for "go
		// panic" or "defer panic".
		panic(targetPanic{args[0]})

	case "recover":
		return doRecover(caller)

	case "ssa:wrapnilchk":
		recv := args[0]
		if recv.(*value) == nil {
			recvType := args[1]
			methodName := args[2]
			panic(fmt.Sprintf("value method (%s).%s called using nil *%s pointer",
				recvType, methodName, recvType))
		}
		return recv

	case "ssa:deferstack":
		return &caller.defers
	}

	panic("unknown built-in: " + fn.Name())
}

func rangeIter(x value, t types.Type) iter {
	switch x := x.(type) {
	case map[value]value:
		return &mapIter{iter: reflect.ValueOf(x).MapRange()}
	case *hashmap:
		return &hashmapIter{iter: reflect.ValueOf(x.entries()).MapRange()}
	case string:
		return &stringIter{Reader: strings.NewReader(x)}
	}
	panic(fmt.Sprintf("cannot range over %T", x))
}

// widen widens a basic typed value x to the widest type of its
// category, one of:
//
//	bool, int64, uint64, float64, complex128, string.
//
// This is inefficient but reduces the size of the cross-product of
// cases we have to consider.
func widen(x value) value {
	switch y := x.(type) {
	case bool, int64, uint64, float64, complex128, string, unsafe.Pointer:
		return x
	case int:
		return int64(y)
	case int8:
		return int64(y)
	case int16:
		return int64(y)
	case int32:
		return int64(y)
	case uint:
		return uint64(y)
	case uint8:
		return uint64(y)
	case uint16:
		return uint64(y)
	case uint32:
		return uint64(y)
	case uintptr:
		return uint64(y)
	case float32:
		return float64(y)
	case complex64:
		return complex128(y)
	}
	panic(fmt.Sprintf("cannot widen %T", x))
}

// conv converts the value x of type t_src to type t_dst and returns
// the result.
// Possible cases are described with the ssa.Convert operator.
func conv(t_dst, t_src types.Type, x value) value {
	ut_src := t_src.Underlying()
	ut_dst := t_dst.Underlying()

	// Destination type is not an "untyped" type.
	if b, ok := ut_dst.(*types.Basic); ok && b.Info()&types.IsUntyped != 0 {
		panic("oops: conversion to 'untyped' type: " + b.String())
	}

	// Nor is it an interface type.
	if _, ok := ut_dst.(*types.Interface); ok {
		if _, ok := ut_src.(*types.Interface); ok {
			panic("oops: Convert should be ChangeInterface")
		} else {
			panic("oops: Convert should be MakeInterface")
		}
	}

	// Remaining conversions:
	//    + untyped string/number/bool constant to a specific
	//      representation.
	//    + conversions between non-complex numeric types.
	//    + conversions between complex numeric types.
	//    + integer/[]byte/[]rune -> string.
	//    + string -> []byte/[]rune.
	//
	// All are treated the same: first we extract the value to the
	// widest representation (int64, uint64, float64, complex128,
	// or string), then we convert it to the desired type.

	switch ut_src := ut_src.(type) {
	case *types.Pointer:
		switch ut_dst := ut_dst.(type) {
		case *types.Basic:
			// *value to unsafe.Pointer?
			if ut_dst.Kind() == types.UnsafePointer {
				return unsafe.Pointer(x.(*value))
			}
		}

	case *types.Slice:
		// []byte or []rune -> string
		switch ut_src.Elem().Underlying().(*types.Basic).Kind() {
		case types.Byte:
			x := x.([]value)
			b := make([]byte, 0, len(x))
			for i := range x {
				b = append(b, x[i].(byte))
			}
			return string(b)

		case types.Rune:
			x := x.([]value)
			r := make([]rune, 0, len(x))
			for i := range x {
				r = append(r, x[i].(rune))
			}
			return string(r)
		}

	case *types.Basic:
		x = widen(x)

		// integer -> string?
		if ut_src.Info()&types.IsInteger != 0 {
			if ut_dst, ok := ut_dst.(*types.Basic); ok && ut_dst.Kind() == types.String {
				return fmt.Sprintf("%c", x)
			}
		}

		// string -> []rune, []byte or string?
		if s, ok := x.(string); ok {
			switch ut_dst := ut_dst.(type) {
			case *types.Slice:
				var res []value
				switch ut_dst.Elem().Underlying().(*types.Basic).Kind() {
				case types.Rune:
					for _, r := range []rune(s) {
						res = append(res, r)
					}
					return res
				case types.Byte:
					for _, b := range []byte(s) {
						res = append(res, b)
					}
					return res
				}
			case *types.Basic:
				if ut_dst.Kind() == types.String {
					return x.(string)
				}
			}
			break // fail: no other conversions for string
		}

		// unsafe.Pointer -> *value
		if ut_src.Kind() == types.UnsafePointer {
			// TODO(adonovan): this is wrong and cannot
			// really be fixed with the current design.
			//
			// return (*value)(x.(unsafe.Pointer))
			// creates a new pointer of a different
			// type but the underlying interface value
			// knows its "true" type and so cannot be
			// meaningfully used through the new pointer.
			//
			// To make this work, the interpreter needs to
			// simulate the memory layout of a real
			// compiled implementation.
			//
			// To at least preserve type-safety, we'll
			// just return the zero value of the
			// destination type.
			return zero(t_dst)
		}

		// Conversions between complex numeric types?
		if ut_src.Info()&types.IsComplex != 0 {
			switch ut_dst.(*types.Basic).Kind() {
			case types.Complex64:
				return complex64(x.(complex128))
			case types.Complex128:
				return x.(complex128)
			}
			break // fail: no other conversions for complex
		}

		// Conversions between non-complex numeric types?
		if ut_src.Info()&types.IsNumeric != 0 {
			kind := ut_dst.(*types.Basic).Kind()
			switch x := x.(type) {
			case int64: // signed integer -> numeric?
				switch kind {
				case types.Int:
					return int(x)
				case types.Int8:
					return int8(x)
				case types.Int16:
					return int16(x)
				case types.Int32:
					return int32(x)
				case types.Int64:
					return int64(x)
				case types.Uint:
					return uint(x)
				case types.Uint8:
					return uint8(x)
				case types.Uint16:
					return uint16(x)
				case types.Uint32:
					return uint32(x)
				case types.Uint64:
					return uint64(x)
				case types.Uintptr:
					return uintptr(x)
				case types.Float32:
					return float32(x)
				case types.Float64:
					return float64(x)
				}

			case uint64: // unsigned integer -> numeric?
				switch kind {
				case types.Int:
					return int(x)
				case types.Int8:
					return int8(x)
				case types.Int16:
					return int16(x)
				case types.Int32:
					return int32(x)
				case types.Int64:
					return int64(x)
				case types.Uint:
					return uint(x)
				case types.Uint8:
					return uint8(x)
				case types.Uint16:
					return uint16(x)
				case types.Uint32:
					return uint32(x)
				case types.Uint64:
					return uint64(x)
				case types.Uintptr:
					return uintptr(x)
				case types.Float32:
					return float32(x)
				case types.Float64:
					return float64(x)
				}

			case float64: // floating point -> numeric?
				switch kind {
				case types.Int:
					return int(x)
				case types.Int8:
					return int8(x)
				case types.Int16:
					return int16(x)
				case types.Int32:
					return int32(x)
				case types.Int64:
					return int64(x)
				case types.Uint:
					return uint(x)
				case types.Uint8:
					return uint8(x)
				case types.Uint16:
					return uint16(x)
				case types.Uint32:
					return uint32(x)
				case types.Uint64:
					return uint64(x)
				case types.Uintptr:
					return uintptr(x)
				case types.Float32:
					return float32(x)
				case types.Float64:
					return float64(x)
				}
			}
		}
	}

	panic(fmt.Sprintf("unsupported conversion: %s  -> %s, dynamic type %T", t_src, t_dst, x))
}

// sliceToArrayPointer converts the value x of type slice to type t_dst
// a pointer to array and returns the result.
func sliceToArrayPointer(t_dst, t_src types.Type, x value) value {
	if _, ok := t_src.Underlying().(*types.Slice); ok {
		if ptr, ok := t_dst.Underlying().(*types.Pointer); ok {
			if arr, ok := ptr.Elem().Underlying().(*types.Array); ok {
				x := x.([]value)
				if arr.Len() > int64(len(x)) {
					panic("array length is greater than slice length")
				}
				if x == nil {
					return zero(t_dst)
				}
				v := value(array(x[:arr.Len()]))
				return &v
			}
		}
	}

	panic(fmt.Sprintf("unsupported conversion: %s  -> %s, dynamic type %T", t_src, t_dst, x))
}

// checkInterface checks that the method set of x implements the
// interface itype.
// On success it returns "", on failure, an error message.
func checkInterface(i *interpreter, itype *types.Interface, x iface) string {
	if meth, _ := types.MissingMethod(x.t, itype, true); meth != nil {
		return fmt.Sprintf("interface conversion: %v is not %v: missing method %s",
			x.t, itype, meth.Name())
	}
	return "" // ok
}

func foldLeft(op func(value, value) value, args []value) value {
	x := args[0]
	for _, arg := range args[1:] {
		x = op(x, arg)
	}
	return x
}

func min(x, y value) value {
	switch x := x.(type) {
	case float32:
		return fmin(x, y.(float32))
	case float64:
		return fmin(x, y.(float64))
	}

	// return (y < x) ? y : x
	if binop(token.LSS, nil, y, x).(bool) {
		return y
	}
	return x
}

func max(x, y value) value {
	switch x := x.(type) {
	case float32:
		return fmax(x, y.(float32))
	case float64:
		return fmax(x, y.(float64))
	}

	// return (y > x) ? y : x
	if binop(token.GTR, nil, y, x).(bool) {
		return y
	}
	return x
}

// copied from $GOROOT/src/runtime/minmax.go

type floaty interface{ ~float32 | ~float64 }

func fmin[F floaty](x, y F) F {
	if y != y || y < x {
		return y
	}
	if x != x || x < y || x != 0 {
		return x
	}
	// x and y are both ±0
	// if either is -0, return -0; else return +0
	return forbits(x, y)
}

func fmax[F floaty](x, y F) F {
	if y != y || y > x {
		return y
	}
	if x != x || x > y || x != 0 {
		return x
	}
	// x and y are both ±0
	// if both are -0, return -0; else return +0
	return fandbits(x, y)
}

func forbits[F floaty](x, y F) F {
	switch unsafe.Sizeof(x) {
	case 4:
		*(*uint32)(unsafe.Pointer(&x)) |= *(*uint32)(unsafe.Pointer(&y))
	case 8:
		*(*uint64)(unsafe.Pointer(&x)) |= *(*uint64)(unsafe.Pointer(&y))
	}
	return x
}

func fandbits[F floaty](x, y F) F {
	switch unsafe.Sizeof(x) {
	case 4:
		*(*uint32)(unsafe.Pointer(&x)) &= *(*uint32)(unsafe.Pointer(&y))
	case 8:
		*(*uint64)(unsafe.Pointer(&x)) &= *(*uint64)(unsafe.Pointer(&y))
	}
	return x
}
```

## File: go/ssa/interp/rangefunc_test.go
```go
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp_test

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/tools/internal/testenv"
)

func TestIssue69298(t *testing.T) {
	testenv.NeedsGo1Point(t, 23)

	goroot := makeGoroot(t)
	run(t, filepath.Join("testdata", "fixedbugs", "issue69298.go"), goroot)
}

func TestRangeFunc(t *testing.T) {
	testenv.NeedsGo1Point(t, 23)

	goroot := makeGoroot(t)
	out := run(t, filepath.Join("testdata", "rangefunc.go"), goroot)

	// Check the output of the tests.
	const (
		RERR_DONE      = "Saw expected panic: yield function called after range loop exit"
		RERR_MISSING   = "Saw expected panic: iterator call did not preserve panic"
		RERR_EXHAUSTED = RERR_DONE // ssa does not distinguish. Same message as RERR_DONE.

		CERR_DONE      = "Saw expected panic: checked rangefunc error: loop iteration after body done"
		CERR_EXHAUSTED = "Saw expected panic: checked rangefunc error: loop iteration after iterator exit"
		CERR_MISSING   = "Saw expected panic: checked rangefunc error: loop iterator swallowed panic"

		panickyIterMsg = "Saw expected panic: Panicky iterator panicking"
	)
	expected := map[string][]string{
		// rangefunc.go
		"TestCheck":                           {"i = 45", CERR_DONE},
		"TestCooperativeBadOfSliceIndex":      {RERR_EXHAUSTED, "i = 36"},
		"TestCooperativeBadOfSliceIndexCheck": {CERR_EXHAUSTED, "i = 36"},
		"TestTrickyIterAll":                   {"i = 36", RERR_EXHAUSTED},
		"TestTrickyIterOne":                   {"i = 1", RERR_EXHAUSTED},
		"TestTrickyIterZero":                  {"i = 0", RERR_EXHAUSTED},
		"TestTrickyIterZeroCheck":             {"i = 0", CERR_EXHAUSTED},
		"TestTrickyIterEcho": {
			"first loop i=0",
			"first loop i=1",
			"first loop i=3",
			"first loop i=6",
			"i = 10",
			"second loop i=0",
			RERR_EXHAUSTED,
			"end i=0",
		},
		"TestTrickyIterEcho2": {
			"k=0,x=1,i=0",
			"k=0,x=2,i=1",
			"k=0,x=3,i=3",
			"k=0,x=4,i=6",
			"i = 10",
			"k=1,x=1,i=0",
			RERR_EXHAUSTED,
			"end i=1",
		},
		"TestBreak1":                {"[1 2 -1 1 2 -2 1 2 -3]"},
		"TestBreak2":                {"[1 2 -1 1 2 -2 1 2 -3]"},
		"TestContinue":              {"[-1 1 2 -2 1 2 -3 1 2 -4]"},
		"TestBreak3":                {"[100 10 2 4 200 10 2 4 20 2 4 300 10 2 4 20 2 4 30]"},
		"TestBreak1BadA":            {"[1 2 -1 1 2 -2 1 2 -3]", RERR_DONE},
		"TestBreak1BadB":            {"[1 2]", RERR_DONE},
		"TestMultiCont0":            {"[1000 10 2 4 2000]"},
		"TestMultiCont1":            {"[1000 10 2 4]", RERR_DONE},
		"TestMultiCont2":            {"[1000 10 2 4]", RERR_DONE},
		"TestMultiCont3":            {"[1000 10 2 4]", RERR_DONE},
		"TestMultiBreak0":           {"[1000 10 2 4]", RERR_DONE},
		"TestMultiBreak1":           {"[1000 10 2 4]", RERR_DONE},
		"TestMultiBreak2":           {"[1000 10 2 4]", RERR_DONE},
		"TestMultiBreak3":           {"[1000 10 2 4]", RERR_DONE},
		"TestPanickyIterator1":      {panickyIterMsg},
		"TestPanickyIterator1Check": {panickyIterMsg},
		"TestPanickyIterator2":      {RERR_MISSING},
		"TestPanickyIterator2Check": {CERR_MISSING},
		"TestPanickyIterator3":      {"[100 10 1 2 200 10 1 2]"},
		"TestPanickyIterator3Check": {"[100 10 1 2 200 10 1 2]"},
		"TestPanickyIterator4":      {RERR_MISSING},
		"TestPanickyIterator4Check": {CERR_MISSING},
		"TestVeryBad1":              {"[1 10]"},
		"TestVeryBad2":              {"[1 10]"},
		"TestVeryBadCheck":          {"[1 10]"},
		"TestOk":                    {"[1 10]"},
		"TestBreak1BadDefer":        {RERR_DONE, "[1 2 -1 1 2 -2 1 2 -3 -30 -20 -10]"},
		"TestReturns":               {"[-1 1 2 -10]", "[-1 1 2 -10]", RERR_DONE, "[-1 1 2 -10]", RERR_DONE},
		"TestGotoA":                 {"testGotoA1[-1 1 2 -2 1 2 -3 1 2 -4 -30 -20 -10]", "testGotoA2[-1 1 2 -2 1 2 -3 1 2 -4 -30 -20 -10]", RERR_DONE, "testGotoA3[-1 1 2 -10]", RERR_DONE},
		"TestGotoB":                 {"testGotoB1[-1 1 2 999 -10]", "testGotoB2[-1 1 2 -10]", RERR_DONE, "testGotoB3[-1 1 2 -10]", RERR_DONE},
		"TestPanicReturns": {
			"Got expected 'f return'",
			"Got expected 'g return'",
			"Got expected 'h return'",
			"Got expected 'k return'",
			"Got expected 'j return'",
			"Got expected 'm return'",
			"Got expected 'n return and n closure return'",
		},
	}
	got := make(map[string][]string)
	for _, ln := range strings.Split(out, "\n") {
		if ind := strings.Index(ln, " \t "); ind >= 0 {
			n, m := ln[:ind], ln[ind+3:]
			got[n] = append(got[n], m)
		}
	}
	for n, es := range expected {
		if gs := got[n]; !reflect.DeepEqual(es, gs) {
			t.Errorf("Output of test %s did not match expected output %v. got %v", n, es, gs)
		}
	}
	for n, gs := range got {
		if expected[n] == nil {
			t.Errorf("No expected output for test %s. got %v", n, gs)
		}
	}
}
```

## File: go/ssa/interp/reflect.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

// Emulated "reflect" package.
//
// We completely replace the built-in "reflect" package.
// The only thing clients can depend upon are that reflect.Type is an
// interface and reflect.Value is an (opaque) struct.

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

type opaqueType struct {
	types.Type
	name string
}

func (t *opaqueType) String() string { return t.name }

// A bogus "reflect" type-checker package.  Shared across interpreters.
var reflectTypesPackage = types.NewPackage("reflect", "reflect")

// rtype is the concrete type the interpreter uses to implement the
// reflect.Type interface.
//
// type rtype <opaque>
var rtypeType = makeNamedType("rtype", &opaqueType{nil, "rtype"})

// error is an (interpreted) named type whose underlying type is string.
// The interpreter uses it for all implementations of the built-in error
// interface that it creates.
// We put it in the "reflect" package for expedience.
//
// type error string
var errorType = makeNamedType("error", &opaqueType{nil, "error"})

func makeNamedType(name string, underlying types.Type) *types.Named {
	obj := types.NewTypeName(token.NoPos, reflectTypesPackage, name, nil)
	return types.NewNamed(obj, underlying, nil)
}

func makeReflectValue(t types.Type, v value) value {
	return structure{rtype{t}, v}
}

// Given a reflect.Value, returns its rtype.
func rV2T(v value) rtype {
	return v.(structure)[0].(rtype)
}

// Given a reflect.Value, returns the underlying interpreter value.
func rV2V(v value) value {
	return v.(structure)[1]
}

// makeReflectType boxes up an rtype in a reflect.Type interface.
func makeReflectType(rt rtype) value {
	return iface{rtypeType, rt}
}

func ext۰reflect۰rtype۰Bits(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) int
	rt := args[0].(rtype).t
	basic, ok := rt.Underlying().(*types.Basic)
	if !ok {
		panic(fmt.Sprintf("reflect.Type.Bits(%T): non-basic type", rt))
	}
	return int(fr.i.sizes.Sizeof(basic)) * 8
}

func ext۰reflect۰rtype۰Elem(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) reflect.Type
	return makeReflectType(rtype{args[0].(rtype).t.Underlying().(interface {
		Elem() types.Type
	}).Elem()})
}

func ext۰reflect۰rtype۰Field(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype, i int) reflect.StructField
	st := args[0].(rtype).t.Underlying().(*types.Struct)
	i := args[1].(int)
	f := st.Field(i)
	return structure{
		f.Name(),
		f.Pkg().Path(),
		makeReflectType(rtype{f.Type()}),
		st.Tag(i),
		0,         // TODO(adonovan): offset
		[]value{}, // TODO(adonovan): indices
		f.Anonymous(),
	}
}

func ext۰reflect۰rtype۰In(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype, i int) int
	i := args[1].(int)
	return makeReflectType(rtype{args[0].(rtype).t.(*types.Signature).Params().At(i).Type()})
}

func ext۰reflect۰rtype۰Kind(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) uint
	return uint(reflectKind(args[0].(rtype).t))
}

func ext۰reflect۰rtype۰NumField(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) int
	return args[0].(rtype).t.Underlying().(*types.Struct).NumFields()
}

func ext۰reflect۰rtype۰NumIn(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) int
	return args[0].(rtype).t.Underlying().(*types.Signature).Params().Len()
}

func ext۰reflect۰rtype۰NumMethod(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) int
	return fr.i.prog.MethodSets.MethodSet(args[0].(rtype).t).Len()
}

func ext۰reflect۰rtype۰NumOut(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) int
	return args[0].(rtype).t.Underlying().(*types.Signature).Results().Len()
}

func ext۰reflect۰rtype۰Out(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype, i int) int
	i := args[1].(int)
	return makeReflectType(rtype{args[0].(rtype).t.Underlying().(*types.Signature).Results().At(i).Type()})
}

func ext۰reflect۰rtype۰Size(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) uintptr
	return uintptr(fr.i.sizes.Sizeof(args[0].(rtype).t))
}

func ext۰reflect۰rtype۰String(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) string
	return args[0].(rtype).t.String()
}

func ext۰reflect۰New(fr *frame, args []value) value {
	// Signature: func (t reflect.Type) reflect.Value
	t := args[0].(iface).v.(rtype).t
	alloc := zero(t)
	return makeReflectValue(types.NewPointer(t), &alloc)
}

func ext۰reflect۰SliceOf(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) Type
	return makeReflectType(rtype{types.NewSlice(args[0].(iface).v.(rtype).t)})
}

func ext۰reflect۰TypeOf(fr *frame, args []value) value {
	// Signature: func (t reflect.rtype) Type
	return makeReflectType(rtype{args[0].(iface).t})
}

func ext۰reflect۰ValueOf(fr *frame, args []value) value {
	// Signature: func (interface{}) reflect.Value
	itf := args[0].(iface)
	return makeReflectValue(itf.t, itf.v)
}

func ext۰reflect۰Zero(fr *frame, args []value) value {
	// Signature: func (t reflect.Type) reflect.Value
	t := args[0].(iface).v.(rtype).t
	return makeReflectValue(t, zero(t))
}

func reflectKind(t types.Type) reflect.Kind {
	switch t := t.(type) {
	case *types.Named, *types.Alias:
		return reflectKind(t.Underlying())
	case *types.Basic:
		switch t.Kind() {
		case types.Bool:
			return reflect.Bool
		case types.Int:
			return reflect.Int
		case types.Int8:
			return reflect.Int8
		case types.Int16:
			return reflect.Int16
		case types.Int32:
			return reflect.Int32
		case types.Int64:
			return reflect.Int64
		case types.Uint:
			return reflect.Uint
		case types.Uint8:
			return reflect.Uint8
		case types.Uint16:
			return reflect.Uint16
		case types.Uint32:
			return reflect.Uint32
		case types.Uint64:
			return reflect.Uint64
		case types.Uintptr:
			return reflect.Uintptr
		case types.Float32:
			return reflect.Float32
		case types.Float64:
			return reflect.Float64
		case types.Complex64:
			return reflect.Complex64
		case types.Complex128:
			return reflect.Complex128
		case types.String:
			return reflect.String
		case types.UnsafePointer:
			return reflect.UnsafePointer
		}
	case *types.Array:
		return reflect.Array
	case *types.Chan:
		return reflect.Chan
	case *types.Signature:
		return reflect.Func
	case *types.Interface:
		return reflect.Interface
	case *types.Map:
		return reflect.Map
	case *types.Pointer:
		return reflect.Pointer
	case *types.Slice:
		return reflect.Slice
	case *types.Struct:
		return reflect.Struct
	}
	panic(fmt.Sprint("unexpected type: ", t))
}

func ext۰reflect۰Value۰Kind(fr *frame, args []value) value {
	// Signature: func (reflect.Value) uint
	return uint(reflectKind(rV2T(args[0]).t))
}

func ext۰reflect۰Value۰String(fr *frame, args []value) value {
	// Signature: func (reflect.Value) string
	return toString(rV2V(args[0]))
}

func ext۰reflect۰Value۰Type(fr *frame, args []value) value {
	// Signature: func (reflect.Value) reflect.Type
	return makeReflectType(rV2T(args[0]))
}

func ext۰reflect۰Value۰Uint(fr *frame, args []value) value {
	// Signature: func (reflect.Value) uint64
	switch v := rV2V(args[0]).(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	case uintptr:
		return uint64(v)
	}
	panic("reflect.Value.Uint")
}

func ext۰reflect۰Value۰Len(fr *frame, args []value) value {
	// Signature: func (reflect.Value) int
	switch v := rV2V(args[0]).(type) {
	case string:
		return len(v)
	case array:
		return len(v)
	case chan value:
		return cap(v)
	case []value:
		return len(v)
	case *hashmap:
		return v.len()
	case map[value]value:
		return len(v)
	default:
		panic(fmt.Sprintf("reflect.(Value).Len(%v)", v))
	}
}

func ext۰reflect۰Value۰MapIndex(fr *frame, args []value) value {
	// Signature: func (reflect.Value) Value
	tValue := rV2T(args[0]).t.Underlying().(*types.Map).Key()
	k := rV2V(args[1])
	switch m := rV2V(args[0]).(type) {
	case map[value]value:
		if v, ok := m[k]; ok {
			return makeReflectValue(tValue, v)
		}

	case *hashmap:
		if v := m.lookup(k.(hashable)); v != nil {
			return makeReflectValue(tValue, v)
		}

	default:
		panic(fmt.Sprintf("(reflect.Value).MapIndex(%T, %T)", m, k))
	}
	return makeReflectValue(nil, nil)
}

func ext۰reflect۰Value۰MapKeys(fr *frame, args []value) value {
	// Signature: func (reflect.Value) []Value
	var keys []value
	tKey := rV2T(args[0]).t.Underlying().(*types.Map).Key()
	switch v := rV2V(args[0]).(type) {
	case map[value]value:
		for k := range v {
			keys = append(keys, makeReflectValue(tKey, k))
		}

	case *hashmap:
		for _, e := range v.entries() {
			for ; e != nil; e = e.next {
				keys = append(keys, makeReflectValue(tKey, e.key))
			}
		}

	default:
		panic(fmt.Sprintf("(reflect.Value).MapKeys(%T)", v))
	}
	return keys
}

func ext۰reflect۰Value۰NumField(fr *frame, args []value) value {
	// Signature: func (reflect.Value) int
	return len(rV2V(args[0]).(structure))
}

func ext۰reflect۰Value۰NumMethod(fr *frame, args []value) value {
	// Signature: func (reflect.Value) int
	return fr.i.prog.MethodSets.MethodSet(rV2T(args[0]).t).Len()
}

func ext۰reflect۰Value۰Pointer(fr *frame, args []value) value {
	// Signature: func (v reflect.Value) uintptr
	switch v := rV2V(args[0]).(type) {
	case *value:
		return uintptr(unsafe.Pointer(v))
	case chan value:
		return reflect.ValueOf(v).Pointer()
	case []value:
		return reflect.ValueOf(v).Pointer()
	case *hashmap:
		return reflect.ValueOf(v.entries()).Pointer()
	case map[value]value:
		return reflect.ValueOf(v).Pointer()
	case *ssa.Function:
		return uintptr(unsafe.Pointer(v))
	case *closure:
		return uintptr(unsafe.Pointer(v))
	default:
		panic(fmt.Sprintf("reflect.(Value).Pointer(%T)", v))
	}
}

func ext۰reflect۰Value۰Index(fr *frame, args []value) value {
	// Signature: func (v reflect.Value, i int) Value
	i := args[1].(int)
	t := rV2T(args[0]).t.Underlying()
	switch v := rV2V(args[0]).(type) {
	case array:
		return makeReflectValue(t.(*types.Array).Elem(), v[i])
	case []value:
		return makeReflectValue(t.(*types.Slice).Elem(), v[i])
	default:
		panic(fmt.Sprintf("reflect.(Value).Index(%T)", v))
	}
}

func ext۰reflect۰Value۰Bool(fr *frame, args []value) value {
	// Signature: func (reflect.Value) bool
	return rV2V(args[0]).(bool)
}

func ext۰reflect۰Value۰CanAddr(fr *frame, args []value) value {
	// Signature: func (v reflect.Value) bool
	// Always false for our representation.
	return false
}

func ext۰reflect۰Value۰CanInterface(fr *frame, args []value) value {
	// Signature: func (v reflect.Value) bool
	// Always true for our representation.
	return true
}

func ext۰reflect۰Value۰Elem(fr *frame, args []value) value {
	// Signature: func (v reflect.Value) reflect.Value
	switch x := rV2V(args[0]).(type) {
	case iface:
		return makeReflectValue(x.t, x.v)
	case *value:
		var v value
		if x != nil {
			v = *x
		}
		return makeReflectValue(rV2T(args[0]).t.Underlying().(*types.Pointer).Elem(), v)
	default:
		panic(fmt.Sprintf("reflect.(Value).Elem(%T)", x))
	}
}

func ext۰reflect۰Value۰Field(fr *frame, args []value) value {
	// Signature: func (v reflect.Value, i int) reflect.Value
	v := args[0]
	i := args[1].(int)
	return makeReflectValue(rV2T(v).t.Underlying().(*types.Struct).Field(i).Type(), rV2V(v).(structure)[i])
}

func ext۰reflect۰Value۰Float(fr *frame, args []value) value {
	// Signature: func (reflect.Value) float64
	switch v := rV2V(args[0]).(type) {
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	}
	panic("reflect.Value.Float")
}

func ext۰reflect۰Value۰Interface(fr *frame, args []value) value {
	// Signature: func (v reflect.Value) interface{}
	return ext۰reflect۰valueInterface(fr, args)
}

func ext۰reflect۰Value۰Int(fr *frame, args []value) value {
	// Signature: func (reflect.Value) int64
	switch x := rV2V(args[0]).(type) {
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	default:
		panic(fmt.Sprintf("reflect.(Value).Int(%T)", x))
	}
}

func ext۰reflect۰Value۰IsNil(fr *frame, args []value) value {
	// Signature: func (reflect.Value) bool
	switch x := rV2V(args[0]).(type) {
	case *value:
		return x == nil
	case chan value:
		return x == nil
	case map[value]value:
		return x == nil
	case *hashmap:
		return x == nil
	case iface:
		return x.t == nil
	case []value:
		return x == nil
	case *ssa.Function:
		return x == nil
	case *ssa.Builtin:
		return x == nil
	case *closure:
		return x == nil
	default:
		panic(fmt.Sprintf("reflect.(Value).IsNil(%T)", x))
	}
}

func ext۰reflect۰Value۰IsValid(fr *frame, args []value) value {
	// Signature: func (reflect.Value) bool
	return rV2V(args[0]) != nil
}

func ext۰reflect۰Value۰Set(fr *frame, args []value) value {
	// TODO(adonovan): implement.
	return nil
}

func ext۰reflect۰valueInterface(fr *frame, args []value) value {
	// Signature: func (v reflect.Value, safe bool) interface{}
	v := args[0].(structure)
	return iface{rV2T(v).t, rV2V(v)}
}

func ext۰reflect۰error۰Error(fr *frame, args []value) value {
	return args[0]
}

// newMethod creates a new method of the specified name, package and receiver type.
func newMethod(pkg *ssa.Package, recvType types.Type, name string) *ssa.Function {
	// TODO(adonovan): fix: hack: currently the only part of Signature
	// that is needed is the "pointerness" of Recv.Type, and for
	// now, we'll set it to always be false since we're only
	// concerned with rtype.  Encapsulate this better.
	sig := types.NewSignatureType(types.NewParam(token.NoPos, nil, "recv", recvType), nil, nil, nil, nil, false)
	fn := pkg.Prog.NewFunction(name, sig, "fake reflect method")
	fn.Pkg = pkg
	return fn
}

func initReflect(i *interpreter) {
	i.reflectPackage = &ssa.Package{
		Prog:    i.prog,
		Pkg:     reflectTypesPackage,
		Members: make(map[string]ssa.Member),
	}

	// Clobber the type-checker's notion of reflect.Value's
	// underlying type so that it more closely matches the fake one
	// (at least in the number of fields---we lie about the type of
	// the rtype field).
	//
	// We must ensure that calls to (ssa.Value).Type() return the
	// fake type so that correct "shape" is used when allocating
	// variables, making zero values, loading, and storing.
	//
	// TODO(adonovan): obviously this is a hack.  We need a cleaner
	// way to fake the reflect package (almost---DeepEqual is fine).
	// One approach would be not to even load its source code, but
	// provide fake source files.  This would guarantee that no bad
	// information leaks into other packages.
	if r := i.prog.ImportedPackage("reflect"); r != nil {
		rV := r.Pkg.Scope().Lookup("Value").Type().(*types.Named)

		// delete bodies of the old methods
		mset := i.prog.MethodSets.MethodSet(rV)
		for j := 0; j < mset.Len(); j++ {
			i.prog.MethodValue(mset.At(j)).Blocks = nil
		}

		tEface := types.NewInterface(nil, nil).Complete()
		rV.SetUnderlying(types.NewStruct([]*types.Var{
			types.NewField(token.NoPos, r.Pkg, "t", tEface, false), // a lie
			types.NewField(token.NoPos, r.Pkg, "v", tEface, false),
		}, nil))
	}

	i.rtypeMethods = methodSet{
		"Bits":      newMethod(i.reflectPackage, rtypeType, "Bits"),
		"Elem":      newMethod(i.reflectPackage, rtypeType, "Elem"),
		"Field":     newMethod(i.reflectPackage, rtypeType, "Field"),
		"In":        newMethod(i.reflectPackage, rtypeType, "In"),
		"Kind":      newMethod(i.reflectPackage, rtypeType, "Kind"),
		"NumField":  newMethod(i.reflectPackage, rtypeType, "NumField"),
		"NumIn":     newMethod(i.reflectPackage, rtypeType, "NumIn"),
		"NumMethod": newMethod(i.reflectPackage, rtypeType, "NumMethod"),
		"NumOut":    newMethod(i.reflectPackage, rtypeType, "NumOut"),
		"Out":       newMethod(i.reflectPackage, rtypeType, "Out"),
		"Size":      newMethod(i.reflectPackage, rtypeType, "Size"),
		"String":    newMethod(i.reflectPackage, rtypeType, "String"),
	}
	i.errorMethods = methodSet{
		"Error": newMethod(i.reflectPackage, errorType, "Error"),
	}
}
```

## File: go/ssa/interp/value.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

// Values
//
// All interpreter values are "boxed" in the empty interface, value.
// The range of possible dynamic types within value are:
//
// - bool
// - numbers (all built-in int/float/complex types are distinguished)
// - string
// - map[value]value --- maps for which  usesBuiltinMap(keyType)
//   *hashmap        --- maps for which !usesBuiltinMap(keyType)
// - chan value
// - []value --- slices
// - iface --- interfaces.
// - structure --- structs.  Fields are ordered and accessed by numeric indices.
// - array --- arrays.
// - *value --- pointers.  Careful: *value is a distinct type from *array etc.
// - *ssa.Function \
//   *ssa.Builtin   } --- functions.  A nil 'func' is always of type *ssa.Function.
//   *closure      /
// - tuple --- as returned by Return, Next, "value,ok" modes, etc.
// - iter --- iterators from 'range' over map or string.
// - bad --- a poison pill for locals that have gone out of scope.
// - rtype -- the interpreter's concrete implementation of reflect.Type
// - **deferred -- the address of a frame's defer stack for a Defer._Stack.
//
// Note that nil is not on this list.
//
// Pay close attention to whether or not the dynamic type is a pointer.
// The compiler cannot help you since value is an empty interface.

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

type value any

type tuple []value

type array []value

type iface struct {
	t types.Type // never an "untyped" type
	v value
}

type structure []value

// For map, array, *array, slice, string or channel.
type iter interface {
	// next returns a Tuple (key, value, ok).
	// key and value are unaliased, e.g. copies of the sequence element.
	next() tuple
}

type closure struct {
	Fn  *ssa.Function
	Env []value
}

type bad struct{}

type rtype struct {
	t types.Type
}

// Hash functions and equivalence relation:

// hashString computes the FNV hash of s.
func hashString(s string) int {
	var h uint32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return int(h)
}

var (
	mu     sync.Mutex
	hasher = typeutil.MakeHasher()
)

// hashType returns a hash for t such that
// types.Identical(x, y) => hashType(x) == hashType(y).
func hashType(t types.Type) int {
	return int(hasher.Hash(t))
}

// usesBuiltinMap returns true if the built-in hash function and
// equivalence relation for type t are consistent with those of the
// interpreter's representation of type t.  Such types are: all basic
// types (bool, numbers, string), pointers and channels.
//
// usesBuiltinMap returns false for types that require a custom map
// implementation: interfaces, arrays and structs.
//
// Panic ensues if t is an invalid map key type: function, map or slice.
func usesBuiltinMap(t types.Type) bool {
	switch t := t.(type) {
	case *types.Basic, *types.Chan, *types.Pointer:
		return true
	case *types.Named, *types.Alias:
		return usesBuiltinMap(t.Underlying())
	case *types.Interface, *types.Array, *types.Struct:
		return false
	}
	panic(fmt.Sprintf("invalid map key type: %T", t))
}

func (x array) eq(t types.Type, _y any) bool {
	y := _y.(array)
	tElt := t.Underlying().(*types.Array).Elem()
	for i, xi := range x {
		if !equals(tElt, xi, y[i]) {
			return false
		}
	}
	return true
}

func (x array) hash(t types.Type) int {
	h := 0
	tElt := t.Underlying().(*types.Array).Elem()
	for _, xi := range x {
		h += hash(t, tElt, xi)
	}
	return h
}

func (x structure) eq(t types.Type, _y any) bool {
	y := _y.(structure)
	tStruct := t.Underlying().(*types.Struct)
	for i, n := 0, tStruct.NumFields(); i < n; i++ {
		if f := tStruct.Field(i); !f.Anonymous() {
			if !equals(f.Type(), x[i], y[i]) {
				return false
			}
		}
	}
	return true
}

func (x structure) hash(t types.Type) int {
	tStruct := t.Underlying().(*types.Struct)
	h := 0
	for i, n := 0, tStruct.NumFields(); i < n; i++ {
		if f := tStruct.Field(i); !f.Anonymous() {
			h += hash(t, f.Type(), x[i])
		}
	}
	return h
}

// nil-tolerant variant of types.Identical.
func sameType(x, y types.Type) bool {
	if x == nil {
		return y == nil
	}
	return y != nil && types.Identical(x, y)
}

func (x iface) eq(t types.Type, _y any) bool {
	y := _y.(iface)
	return sameType(x.t, y.t) && (x.t == nil || equals(x.t, x.v, y.v))
}

func (x iface) hash(outer types.Type) int {
	return hashType(x.t)*8581 + hash(outer, x.t, x.v)
}

func (x rtype) hash(_ types.Type) int {
	return hashType(x.t)
}

func (x rtype) eq(_ types.Type, y any) bool {
	return types.Identical(x.t, y.(rtype).t)
}

// equals returns true iff x and y are equal according to Go's
// linguistic equivalence relation for type t.
// In a well-typed program, the dynamic types of x and y are
// guaranteed equal.
func equals(t types.Type, x, y value) bool {
	switch x := x.(type) {
	case bool:
		return x == y.(bool)
	case int:
		return x == y.(int)
	case int8:
		return x == y.(int8)
	case int16:
		return x == y.(int16)
	case int32:
		return x == y.(int32)
	case int64:
		return x == y.(int64)
	case uint:
		return x == y.(uint)
	case uint8:
		return x == y.(uint8)
	case uint16:
		return x == y.(uint16)
	case uint32:
		return x == y.(uint32)
	case uint64:
		return x == y.(uint64)
	case uintptr:
		return x == y.(uintptr)
	case float32:
		return x == y.(float32)
	case float64:
		return x == y.(float64)
	case complex64:
		return x == y.(complex64)
	case complex128:
		return x == y.(complex128)
	case string:
		return x == y.(string)
	case *value:
		return x == y.(*value)
	case chan value:
		return x == y.(chan value)
	case structure:
		return x.eq(t, y)
	case array:
		return x.eq(t, y)
	case iface:
		return x.eq(t, y)
	case rtype:
		return x.eq(t, y)
	}

	// Since map, func and slice don't support comparison, this
	// case is only reachable if one of x or y is literally nil
	// (handled in eqnil) or via interface{} values.
	panic(fmt.Sprintf("comparing uncomparable type %s", t))
}

// Returns an integer hash of x such that equals(x, y) => hash(x) == hash(y).
// The outer type is used only for the "unhashable" panic message.
func hash(outer, t types.Type, x value) int {
	switch x := x.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case uintptr:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case complex64:
		return int(real(x))
	case complex128:
		return int(real(x))
	case string:
		return hashString(x)
	case *value:
		return int(uintptr(unsafe.Pointer(x)))
	case chan value:
		return int(uintptr(reflect.ValueOf(x).Pointer()))
	case structure:
		return x.hash(t)
	case array:
		return x.hash(t)
	case iface:
		return x.hash(t)
	case rtype:
		return x.hash(t)
	}
	panic(fmt.Sprintf("unhashable type %v", outer))
}

// reflect.Value struct values don't have a fixed shape, since the
// payload can be a scalar or an aggregate depending on the instance.
// So store (and load) can't simply use recursion over the shape of the
// rhs value, or the lhs, to copy the value; we need the static type
// information.  (We can't make reflect.Value a new basic data type
// because its "structness" is exposed to Go programs.)

// load returns the value of type T in *addr.
func load(T types.Type, addr *value) value {
	switch T := T.Underlying().(type) {
	case *types.Struct:
		v := (*addr).(structure)
		a := make(structure, len(v))
		for i := range a {
			a[i] = load(T.Field(i).Type(), &v[i])
		}
		return a
	case *types.Array:
		v := (*addr).(array)
		a := make(array, len(v))
		for i := range a {
			a[i] = load(T.Elem(), &v[i])
		}
		return a
	default:
		return *addr
	}
}

// store stores value v of type T into *addr.
func store(T types.Type, addr *value, v value) {
	switch T := T.Underlying().(type) {
	case *types.Struct:
		lhs := (*addr).(structure)
		rhs := v.(structure)
		for i := range lhs {
			store(T.Field(i).Type(), &lhs[i], rhs[i])
		}
	case *types.Array:
		lhs := (*addr).(array)
		rhs := v.(array)
		for i := range lhs {
			store(T.Elem(), &lhs[i], rhs[i])
		}
	default:
		*addr = v
	}
}

// Prints in the style of built-in println.
// (More or less; in gc println is actually a compiler intrinsic and
// can distinguish println(1) from println(interface{}(1)).)
func writeValue(buf *bytes.Buffer, v value) {
	switch v := v.(type) {
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string:
		fmt.Fprintf(buf, "%v", v)

	case map[value]value:
		buf.WriteString("map[")
		sep := ""
		for k, e := range v {
			buf.WriteString(sep)
			sep = " "
			writeValue(buf, k)
			buf.WriteString(":")
			writeValue(buf, e)
		}
		buf.WriteString("]")

	case *hashmap:
		buf.WriteString("map[")
		sep := " "
		for _, e := range v.entries() {
			for e != nil {
				buf.WriteString(sep)
				sep = " "
				writeValue(buf, e.key)
				buf.WriteString(":")
				writeValue(buf, e.value)
				e = e.next
			}
		}
		buf.WriteString("]")

	case chan value:
		fmt.Fprintf(buf, "%v", v) // (an address)

	case *value:
		if v == nil {
			buf.WriteString("<nil>")
		} else {
			fmt.Fprintf(buf, "%p", v)
		}

	case iface:
		fmt.Fprintf(buf, "(%s, ", v.t)
		writeValue(buf, v.v)
		buf.WriteString(")")

	case structure:
		buf.WriteString("{")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(" ")
			}
			writeValue(buf, e)
		}
		buf.WriteString("}")

	case array:
		buf.WriteString("[")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(" ")
			}
			writeValue(buf, e)
		}
		buf.WriteString("]")

	case []value:
		buf.WriteString("[")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(" ")
			}
			writeValue(buf, e)
		}
		buf.WriteString("]")

	case *ssa.Function, *ssa.Builtin, *closure:
		fmt.Fprintf(buf, "%p", v) // (an address)

	case rtype:
		buf.WriteString(v.t.String())

	case tuple:
		// Unreachable in well-formed Go programs
		buf.WriteString("(")
		for i, e := range v {
			if i > 0 {
				buf.WriteString(", ")
			}
			writeValue(buf, e)
		}
		buf.WriteString(")")

	default:
		fmt.Fprintf(buf, "<%T>", v)
	}
}

// Implements printing of Go values in the style of built-in println.
func toString(v value) string {
	var b bytes.Buffer
	writeValue(&b, v)
	return b.String()
}

// ------------------------------------------------------------------------
// Iterators

type stringIter struct {
	*strings.Reader
	i int
}

func (it *stringIter) next() tuple {
	okv := make(tuple, 3)
	ch, n, err := it.ReadRune()
	ok := err != io.EOF
	okv[0] = ok
	if ok {
		okv[1] = it.i
		okv[2] = ch
	}
	it.i += n
	return okv
}

type mapIter struct {
	iter *reflect.MapIter
	ok   bool
}

func (it *mapIter) next() tuple {
	it.ok = it.iter.Next()
	if !it.ok {
		return []value{false, nil, nil}
	}
	k, v := it.iter.Key().Interface(), it.iter.Value().Interface()
	return []value{true, k, v}
}

type hashmapIter struct {
	iter *reflect.MapIter
	ok   bool
	cur  *entry
}

func (it *hashmapIter) next() tuple {
	for {
		if it.cur != nil {
			k, v := it.cur.key, it.cur.value
			it.cur = it.cur.next
			return []value{true, k, v}
		}
		it.ok = it.iter.Next()
		if !it.ok {
			return []value{false, nil, nil}
		}
		it.cur = it.iter.Value().Interface().(*entry)
	}
}
```

## File: go/ssa/ssautil/testdata/switches.txtar
```
-- go.mod --
module example.com
go 1.22

-- switches.go --
package main

// This file is the input to TestSwitches in switch_test.go.
// Each multiway conditional with constant or type cases (Switch)
// discovered by Switches is printed, and compared with the
// comments.
//
// The body of each case is printed as the value of its first
// instruction.

// -------- Value switches --------

func SimpleSwitch(x, y int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(23:int)
	// case 3:int: print(23:int)
	// case 4:int: print(3:int)
	// default: x == y
	// }
	switch x {
	case 1:
		print(1)
	case 2, 3:
		print(23)
		fallthrough
	case 4:
		print(3)
	default:
		print(4)
	case y:
		print(5)
	}
	print(6)
}

func four() int { return 4 }

// A non-constant case makes a switch "impure", but its pure
// cases form two separate switches.
func SwitchWithNonConstantCase(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(23:int)
	// case 3:int: print(23:int)
	// default: four()
	// }

	// switch x {
	// case 5:int: print(5:int)
	// case 6:int: print(6:int)
	// default: print("done":string)
	// }
	switch x {
	case 1:
		print(1)
	case 2, 3:
		print(23)
	case four():
		print(3)
	case 5:
		print(5)
	case 6:
		print(6)
	}
	print("done")
}

// Switches may be found even where the source
// program doesn't have a switch statement.

func ImplicitSwitches(x, y int) {
	// switch x {
	// case 1:int: print(12:int)
	// case 2:int: print(12:int)
	// default: x < 5:int
	// }
	if x == 1 || 2 == x || x < 5 {
		print(12)
	}

	// switch x {
	// case 3:int: print(34:int)
	// case 4:int: print(34:int)
	// default: x == y
	// }
	if x == 3 || 4 == x || x == y {
		print(34)
	}

	// Not a switch: no consistent variable.
	if x == 5 || y == 6 {
		print(56)
	}

	// Not a switch: only one constant comparison.
	if x == 7 || x == y {
		print(78)
	}
}

func IfElseBasedSwitch(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(2:int)
	// default: print("else":string)
	// }
	if x == 1 {
		print(1)
	} else if x == 2 {
		print(2)
	} else {
		print("else")
	}
}

func GotoBasedSwitch(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(2:int)
	// default: print("else":string)
	// }
	if x == 1 {
		goto L1
	}
	if x == 2 {
		goto L2
	}
	print("else")
L1:
	print(1)
	goto end
L2:
	print(2)
end:
}

func SwitchInAForLoop(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(2:int)
	// default: print("head":string)
	// }
loop:
	for {
		print("head")
		switch x {
		case 1:
			print(1)
			break loop
		case 2:
			print(2)
			break loop
		}
	}
}

// This case is a switch in a for-loop, both constructed using goto.
// As before, the default case points back to the block containing the
// switch, but that's ok.
func SwitchInAForLoopUsingGoto(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: print(2:int)
	// default: print("head":string)
	// }
loop:
	print("head")
	if x == 1 {
		goto L1
	}
	if x == 2 {
		goto L2
	}
	goto loop
L1:
	print(1)
	goto end
L2:
	print(2)
end:
}

func UnstructuredSwitchInAForLoop(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 2:int: x == 1:int
	// default: print("end":string)
	// }
	for {
		if x == 1 {
			print(1)
			return
		}
		if x == 2 {
			continue
		}
		break
	}
	print("end")
}

func CaseWithMultiplePreds(x int) {
	for {
		if x == 1 {
			print(1)
			return
		}
	loop:
		// This block has multiple predecessors,
		// so can't be treated as a switch case.
		if x == 2 {
			goto loop
		}
		break
	}
	print("end")
}

func DuplicateConstantsAreNotEliminated(x int) {
	// switch x {
	// case 1:int: print(1:int)
	// case 1:int: print("1a":string)
	// case 2:int: print(2:int)
	// default: return
	// }
	if x == 1 {
		print(1)
	} else if x == 1 { // duplicate => unreachable
		print("1a")
	} else if x == 2 {
		print(2)
	}
}

// Interface values (created by comparisons) are not constants,
// so ConstSwitch.X is never of interface type.
func MakeInterfaceIsNotAConstant(x interface{}) {
	if x == "foo" {
		print("foo")
	} else if x == 1 {
		print(1)
	}
}

func ZeroInitializedVarsAreConstants(x int) {
	// switch x {
	// case 0:int: print(1:int)
	// case 2:int: print(2:int)
	// default: print("end":string)
	// }
	var zero int // SSA construction replaces zero with 0
	if x == zero {
		print(1)
	} else if x == 2 {
		print(2)
	}
	print("end")
}

// -------- Select --------

// NB, potentially fragile reliance on register number.
func SelectDesugarsToSwitch(ch chan int) {
	// switch t1 {
	// case 0:int: extract t0 #2
	// case 1:int: println(0:int)
	// case 2:int: println(1:int)
	// default: println("default":string)
	// }
	select {
	case x := <-ch:
		println(x)
	case <-ch:
		println(0)
	case ch <- 1:
		println(1)
	default:
		println("default")
	}
}

// NB, potentially fragile reliance on register number.
func NonblockingSelectDefaultCasePanics(ch chan int) {
	// switch t1 {
	// case 0:int: extract t0 #2
	// case 1:int: println(0:int)
	// case 2:int: println(1:int)
	// default: make interface{} <- string ("blocking select m...":string)
	// }
	select {
	case x := <-ch:
		println(x)
	case <-ch:
		println(0)
	case ch <- 1:
		println(1)
	}
}

// -------- Type switches --------

// NB, reliance on fragile register numbering.
func SimpleTypeSwitch(x interface{}) {
	// switch x.(type) {
	// case t3 int: println(x)
	// case t7 bool: println(x)
	// case t10 string: println(t10)
	// default: println(x)
	// }
	switch y := x.(type) {
	case nil:
		println(y)
	case int, bool:
		println(y)
	case string:
		println(y)
	default:
		println(y)
	}
}

// NB, potentially fragile reliance on register number.
func DuplicateTypesAreNotEliminated(x interface{}) {
	// switch x.(type) {
	// case t1 string: println(1:int)
	// case t5 interface{}: println(t5)
	// case t9 int: println(3:int)
	// default: return
	// }
	switch y := x.(type) {
	case string:
		println(1)
	case interface{}:
		println(y)
	case int:
		println(3) // unreachable!
	}
}

// NB, potentially fragile reliance on register number.
func AdHocTypeSwitch(x interface{}) {
	// switch x.(type) {
	// case t1 int: println(t1)
	// case t5 string: println(t5)
	// default: print("default":string)
	// }
	if i, ok := x.(int); ok {
		println(i)
	} else if s, ok := x.(string); ok {
		println(s)
	} else {
		print("default")
	}
}
```

## File: go/ssa/ssautil/deprecated_test.go
```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil_test

// Tests of deprecated public APIs.
// We are keeping some tests around to have some test of the public API.

import (
	"go/parser"
	"os"
	"testing"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testenv"
)

// TestCreateProgram tests CreateProgram which has an x/tools/go/loader.Program.
func TestCreateProgram(t *testing.T) {
	testenv.NeedsGoBuild(t) // for importer.Default()

	conf := loader.Config{ParserMode: parser.ParseComments}
	f, err := conf.ParseFile("hello.go", hello)
	if err != nil {
		t.Fatal(err)
	}

	conf.CreateFromFiles("main", f)
	iprog, err := conf.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(iprog.Created) != 1 {
		t.Fatalf("Expected 1 Created package. got %d", len(iprog.Created))
	}
	pkg := iprog.Created[0].Pkg

	prog := ssautil.CreateProgram(iprog, ssa.BuilderMode(0))
	ssapkg := prog.Package(pkg)
	ssapkg.Build()

	if pkg.Name() != "main" {
		t.Errorf("pkg.Name() = %s, want main", pkg.Name())
	}
	if ssapkg.Func("main") == nil {
		ssapkg.WriteTo(os.Stderr)
		t.Errorf("ssapkg has no main function")
	}
}
```

## File: go/ssa/ssautil/deprecated.go
```go
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil

// This file contains deprecated public APIs.
// We discourage their use.

import (
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

// CreateProgram returns a new program in SSA form, given a program
// loaded from source.  An SSA package is created for each transitively
// error-free package of lprog.
//
// Code for bodies of functions is not built until Build is called
// on the result.
//
// The mode parameter controls diagnostics and checking during SSA construction.
//
// Deprecated: Use [golang.org/x/tools/go/packages] and the [Packages]
// function instead; see ssa.Example_loadPackages.
func CreateProgram(lprog *loader.Program, mode ssa.BuilderMode) *ssa.Program {
	prog := ssa.NewProgram(lprog.Fset, mode)

	for _, info := range lprog.AllPackages {
		if info.TransitivelyErrorFree {
			prog.CreatePackage(info.Pkg, info.Files, &info.Info, info.Importable)
		}
	}

	return prog
}
```

## File: go/ssa/ssautil/load_test.go
```go
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil_test

import (
	"bytes"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/packagestest"
	"golang.org/x/tools/internal/testenv"
)

const hello = `package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
`

func TestBuildPackage(t *testing.T) {
	testenv.NeedsGoBuild(t) // for importer.Default()

	// There is a more substantial test of BuildPackage and the
	// SSA program it builds in ../ssa/builder_test.go.

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, mode := range []ssa.BuilderMode{
		ssa.SanityCheckFunctions,
		ssa.InstantiateGenerics | ssa.SanityCheckFunctions,
	} {
		pkg := types.NewPackage("hello", "")
		ssapkg, _, err := ssautil.BuildPackage(&types.Config{Importer: importer.Default()}, fset, pkg, []*ast.File{f}, mode)
		if err != nil {
			t.Fatal(err)
		}
		if pkg.Name() != "main" {
			t.Errorf("pkg.Name() = %s, want main", pkg.Name())
		}
		if ssapkg.Func("main") == nil {
			ssapkg.WriteTo(os.Stderr)
			t.Errorf("ssapkg has no main function")
		}

	}
}

func TestPackages(t *testing.T) {
	testenv.NeedsGoPackages(t)

	cfg := &packages.Config{Mode: packages.LoadSyntax}
	initial, err := packages.Load(cfg, "bytes")
	if err != nil {
		t.Fatal(err)
	}
	if packages.PrintErrors(initial) > 0 {
		t.Fatal("there were errors")
	}

	for _, mode := range []ssa.BuilderMode{
		ssa.SanityCheckFunctions,
		ssa.SanityCheckFunctions | ssa.InstantiateGenerics,
	} {
		prog, pkgs := ssautil.Packages(initial, mode)
		bytesNewBuffer := pkgs[0].Func("NewBuffer")
		bytesNewBuffer.Pkg.Build()

		// We'll dump the SSA of bytes.NewBuffer because it is small and stable.
		out := new(bytes.Buffer)
		bytesNewBuffer.WriteTo(out)

		// For determinism, sanitize the location.
		location := prog.Fset.Position(bytesNewBuffer.Pos()).String()
		got := strings.Replace(out.String(), location, "$GOROOT/src/bytes/buffer.go:1", -1)

		want := `
# Name: bytes.NewBuffer
# Package: bytes
# Location: $GOROOT/src/bytes/buffer.go:1
func NewBuffer(buf []byte) *Buffer:
0:                                                                entry P:0 S:0
	t0 = new Buffer (complit)                                       *Buffer
	t1 = &t0.buf [#0]                                               *[]byte
	*t1 = buf
	return t0

`[1:]
		if got != want {
			t.Errorf("bytes.NewBuffer SSA = <<%s>>, want <<%s>>", got, want)
		}
	}
}

func TestBuildPackage_MissingImport(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "bad.go", `package bad; import "missing"`, 0)
	if err != nil {
		t.Fatal(err)
	}

	pkg := types.NewPackage("bad", "")
	ssapkg, _, err := ssautil.BuildPackage(new(types.Config), fset, pkg, []*ast.File{f}, ssa.BuilderMode(0))
	if err == nil || ssapkg != nil {
		t.Fatal("BuildPackage succeeded unexpectedly")
	}
}

func TestIssue28106(t *testing.T) {
	testenv.NeedsGoPackages(t)

	// In go1.10, go/packages loads all packages from source, not
	// export data, but does not type check function bodies of
	// imported packages. This test ensures that we do not attempt
	// to run the SSA builder on functions without type information.
	cfg := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(cfg, "runtime")
	if err != nil {
		t.Fatal(err)
	}
	prog, _ := ssautil.Packages(pkgs, ssa.BuilderMode(0))
	prog.Build() // no crash
}

func TestIssue53604(t *testing.T) {
	// Tests that variable initializers are not added to init() when syntax
	// is not present but types.Info is available.
	//
	// Packages x, y, z are loaded with mode `packages.LoadSyntax`.
	// Package x imports y, and y imports z.
	// Packages are built using ssautil.Packages() with x and z as roots.
	// This setup creates y using CreatePackage(pkg, files, info, ...)
	// where len(files) == 0 but info != nil.
	//
	// Tests that globals from y are not initialized.
	e := packagestest.Export(t, packagestest.Modules, []packagestest.Module{
		{
			Name: "golang.org/fake",
			Files: map[string]any{
				"x/x.go": `package x; import "golang.org/fake/y"; var V = y.F()`,
				"y/y.go": `package y; import "golang.org/fake/z"; var F = func () *int { return &z.Z } `,
				"z/z.go": `package z; var Z int`,
			},
		},
	})
	defer e.Cleanup()

	// Load x and z as entry packages using packages.LoadSyntax
	e.Config.Mode = packages.LoadSyntax
	pkgs, err := packages.Load(e.Config, path.Join(e.Temp(), "fake/x"), path.Join(e.Temp(), "fake/z"))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range pkgs {
		if len(p.Errors) > 0 {
			t.Fatalf("%v", p.Errors)
		}
	}

	prog, _ := ssautil.Packages(pkgs, ssa.BuilderMode(0))
	prog.Build()

	// y does not initialize F.
	y := prog.ImportedPackage("golang.org/fake/y")
	if y == nil {
		t.Fatal("Failed to load intermediate package y")
	}
	yinit := y.Members["init"].(*ssa.Function)
	for _, bb := range yinit.Blocks {
		for _, i := range bb.Instrs {
			if store, ok := i.(*ssa.Store); ok && store.Addr == y.Var("F") {
				t.Errorf("y.init() stores to F %v", store)
			}
		}
	}

}
```

## File: go/ssa/ssautil/load.go
```go
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil

// This file defines utility functions for constructing programs in SSA form.

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
)

// Packages creates an SSA program for a set of packages.
//
// The packages must have been loaded from source syntax using the
// [packages.Load] function in [packages.LoadSyntax] or
// [packages.LoadAllSyntax] mode.
//
// Packages creates an SSA package for each well-typed package in the
// initial list, plus all their dependencies. The resulting list of
// packages corresponds to the list of initial packages, and may contain
// a nil if SSA code could not be constructed for the corresponding initial
// package due to type errors.
//
// Code for bodies of functions is not built until [Program.Build] is
// called on the resulting Program. SSA code is constructed only for
// the initial packages with well-typed syntax trees.
//
// The mode parameter controls diagnostics and checking during SSA construction.
func Packages(initial []*packages.Package, mode ssa.BuilderMode) (*ssa.Program, []*ssa.Package) {
	// TODO(adonovan): opt: this calls CreatePackage far more than
	// necessary: for all dependencies, not just the (non-initial)
	// direct dependencies of the initial packages.
	//
	// But can it reasonably be changed without breaking the
	// spirit and/or letter of the law above? Clients may notice
	// if we call CreatePackage less, as methods like
	// Program.FuncValue will return nil. Or must we provide a new
	// function (and perhaps deprecate this one)? Is it worth it?
	//
	// Tim King makes the interesting point that it would be
	// possible to entirely alleviate the client from the burden
	// of calling CreatePackage for non-syntax packages, if we
	// were to treat vars and funcs lazily in the same way we now
	// treat methods. (In essence, try to move away from the
	// notion of ssa.Packages, and make the Program answer
	// all reasonable questions about any types.Object.)

	return doPackages(initial, mode, false)
}

// AllPackages creates an SSA program for a set of packages plus all
// their dependencies.
//
// The packages must have been loaded from source syntax using the
// [packages.Load] function in [packages.LoadAllSyntax] mode.
//
// AllPackages creates an SSA package for each well-typed package in the
// initial list, plus all their dependencies. The resulting list of
// packages corresponds to the list of initial packages, and may contain
// a nil if SSA code could not be constructed for the corresponding
// initial package due to type errors.
//
// Code for bodies of functions is not built until Build is called on
// the resulting Program. SSA code is constructed for all packages with
// well-typed syntax trees.
//
// The mode parameter controls diagnostics and checking during SSA construction.
func AllPackages(initial []*packages.Package, mode ssa.BuilderMode) (*ssa.Program, []*ssa.Package) {
	return doPackages(initial, mode, true)
}

func doPackages(initial []*packages.Package, mode ssa.BuilderMode, deps bool) (*ssa.Program, []*ssa.Package) {

	var fset *token.FileSet
	if len(initial) > 0 {
		fset = initial[0].Fset
	}

	prog := ssa.NewProgram(fset, mode)

	isInitial := make(map[*packages.Package]bool, len(initial))
	for _, p := range initial {
		isInitial[p] = true
	}

	ssamap := make(map[*packages.Package]*ssa.Package)
	packages.Visit(initial, nil, func(p *packages.Package) {
		if p.Types != nil && !p.IllTyped {
			var files []*ast.File
			var info *types.Info
			if deps || isInitial[p] {
				files = p.Syntax
				info = p.TypesInfo
			}
			ssamap[p] = prog.CreatePackage(p.Types, files, info, true)
		}
	})

	var ssapkgs []*ssa.Package
	for _, p := range initial {
		ssapkgs = append(ssapkgs, ssamap[p]) // may be nil
	}
	return prog, ssapkgs
}

// BuildPackage builds an SSA program with SSA intermediate
// representation (IR) for all functions of a single package.
//
// It populates pkg by type-checking the specified file syntax trees.  All
// dependencies are loaded using the importer specified by tc, which
// typically loads compiler export data; SSA code cannot be built for
// those packages.  BuildPackage then constructs an [ssa.Program] with all
// dependency packages created, and builds and returns the SSA package
// corresponding to pkg.
//
// The caller must have set pkg.Path to the import path.
//
// The operation fails if there were any type-checking or import errors.
//
// See ../example_test.go for an example.
func BuildPackage(tc *types.Config, fset *token.FileSet, pkg *types.Package, files []*ast.File, mode ssa.BuilderMode) (*ssa.Package, *types.Info, error) {
	if fset == nil {
		panic("no token.FileSet")
	}
	if pkg.Path() == "" {
		panic("package has no import path")
	}

	info := &types.Info{
		Types:        make(map[ast.Expr]types.TypeAndValue),
		Defs:         make(map[*ast.Ident]types.Object),
		Uses:         make(map[*ast.Ident]types.Object),
		Implicits:    make(map[ast.Node]types.Object),
		Instances:    make(map[*ast.Ident]types.Instance),
		Scopes:       make(map[ast.Node]*types.Scope),
		Selections:   make(map[*ast.SelectorExpr]*types.Selection),
		FileVersions: make(map[*ast.File]string),
	}
	if err := types.NewChecker(tc, fset, pkg, info).Files(files); err != nil {
		return nil, nil, err
	}

	prog := ssa.NewProgram(fset, mode)

	// Create SSA packages for all imports.
	// Order is not significant.
	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				prog.CreatePackage(p, nil, nil, true)
				createAll(p.Imports())
			}
		}
	}
	createAll(pkg.Imports())

	// TODO(adonovan): we could replace createAll with just:
	//
	// // Create SSA packages for all imports.
	// for _, p := range pkg.Imports() {
	// 	prog.CreatePackage(p, nil, nil, true)
	// }
	//
	// (with minor changes to changes to ../builder_test.go as
	// shown in CL 511715 PS 10.) But this would strictly violate
	// the letter of the doc comment above, which says "all
	// dependencies created".
	//
	// Tim makes the good point with some extra work we could
	// remove the need for any CreatePackage calls except the
	// ones with syntax (i.e. primary packages). Of course
	// You wouldn't have ssa.Packages and Members for as
	// many things but no-one really uses that anyway.
	// I wish I had done this from the outset.

	// Create and build the primary package.
	ssapkg := prog.CreatePackage(pkg, files, info, false)
	ssapkg.Build()
	return ssapkg, info, nil
}
```

## File: go/ssa/ssautil/switch_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// No testdata on Android.

//go:build !android

package ssautil_test

import (
	"strings"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testfiles"
	"golang.org/x/tools/txtar"
)

func TestSwitches(t *testing.T) {
	archive, err := txtar.ParseFile("testdata/switches.txtar")
	if err != nil {
		t.Fatal(err)
	}
	ppkgs := testfiles.LoadPackages(t, archive, ".")
	if len(ppkgs) != 1 {
		t.Fatalf("Expected to load one package but got %d", len(ppkgs))
	}
	f := ppkgs[0].Syntax[0]

	prog, _ := ssautil.Packages(ppkgs, ssa.BuilderMode(0))
	mainPkg := prog.Package(ppkgs[0].Types)
	mainPkg.Build()

	for _, mem := range mainPkg.Members {
		if fn, ok := mem.(*ssa.Function); ok {
			if fn.Synthetic != "" {
				continue // e.g. init()
			}
			// Each (multi-line) "switch" comment within
			// this function must match the printed form
			// of a ConstSwitch.
			var wantSwitches []string
			for _, c := range f.Comments {
				if fn.Syntax().Pos() <= c.Pos() && c.Pos() < fn.Syntax().End() {
					text := strings.TrimSpace(c.Text())
					if strings.HasPrefix(text, "switch ") {
						wantSwitches = append(wantSwitches, text)
					}
				}
			}

			switches := ssautil.Switches(fn)
			if len(switches) != len(wantSwitches) {
				t.Errorf("in %s, found %d switches, want %d", fn, len(switches), len(wantSwitches))
			}
			for i, sw := range switches {
				got := sw.String()
				if i >= len(wantSwitches) {
					continue
				}
				want := wantSwitches[i]
				if got != want {
					t.Errorf("in %s, found switch %d: got <<%s>>, want <<%s>>", fn, i, got, want)
				}
			}
		}
	}
}
```

## File: go/ssa/ssautil/switch.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil

// This file implements discovery of switch and type-switch constructs
// from low-level control flow.
//
// Many techniques exist for compiling a high-level switch with
// constant cases to efficient machine code.  The optimal choice will
// depend on the data type, the specific case values, the code in the
// body of each case, and the hardware.
// Some examples:
// - a lookup table (for a switch that maps constants to constants)
// - a computed goto
// - a binary tree
// - a perfect hash
// - a two-level switch (to partition constant strings by their first byte).

import (
	"bytes"
	"fmt"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// A ConstCase represents a single constant comparison.
// It is part of a Switch.
type ConstCase struct {
	Block *ssa.BasicBlock // block performing the comparison
	Body  *ssa.BasicBlock // body of the case
	Value *ssa.Const      // case comparand
}

// A TypeCase represents a single type assertion.
// It is part of a Switch.
type TypeCase struct {
	Block   *ssa.BasicBlock // block performing the type assert
	Body    *ssa.BasicBlock // body of the case
	Type    types.Type      // case type
	Binding ssa.Value       // value bound by this case
}

// A Switch is a logical high-level control flow operation
// (a multiway branch) discovered by analysis of a CFG containing
// only if/else chains.  It is not part of the ssa.Instruction set.
//
// One of ConstCases and TypeCases has length >= 2;
// the other is nil.
//
// In a value switch, the list of cases may contain duplicate constants.
// A type switch may contain duplicate types, or types assignable
// to an interface type also in the list.
// TODO(adonovan): eliminate such duplicates.
type Switch struct {
	Start      *ssa.BasicBlock // block containing start of if/else chain
	X          ssa.Value       // the switch operand
	ConstCases []ConstCase     // ordered list of constant comparisons
	TypeCases  []TypeCase      // ordered list of type assertions
	Default    *ssa.BasicBlock // successor if all comparisons fail
}

func (sw *Switch) String() string {
	// We represent each block by the String() of its
	// first Instruction, e.g. "print(42:int)".
	var buf bytes.Buffer
	if sw.ConstCases != nil {
		fmt.Fprintf(&buf, "switch %s {\n", sw.X.Name())
		for _, c := range sw.ConstCases {
			fmt.Fprintf(&buf, "case %s: %s\n", c.Value, c.Body.Instrs[0])
		}
	} else {
		fmt.Fprintf(&buf, "switch %s.(type) {\n", sw.X.Name())
		for _, c := range sw.TypeCases {
			fmt.Fprintf(&buf, "case %s %s: %s\n",
				c.Binding.Name(), c.Type, c.Body.Instrs[0])
		}
	}
	if sw.Default != nil {
		fmt.Fprintf(&buf, "default: %s\n", sw.Default.Instrs[0])
	}
	fmt.Fprintf(&buf, "}")
	return buf.String()
}

// Switches examines the control-flow graph of fn and returns the
// set of inferred value and type switches.  A value switch tests an
// ssa.Value for equality against two or more compile-time constant
// values.  Switches involving link-time constants (addresses) are
// ignored.  A type switch type-asserts an ssa.Value against two or
// more types.
//
// The switches are returned in dominance order.
//
// The resulting switches do not necessarily correspond to uses of the
// 'switch' keyword in the source: for example, a single source-level
// switch statement with non-constant cases may result in zero, one or
// many Switches, one per plural sequence of constant cases.
// Switches may even be inferred from if/else- or goto-based control flow.
// (In general, the control flow constructs of the source program
// cannot be faithfully reproduced from the SSA representation.)
func Switches(fn *ssa.Function) []Switch {
	// Traverse the CFG in dominance order, so we don't
	// enter an if/else-chain in the middle.
	var switches []Switch
	seen := make(map[*ssa.BasicBlock]bool) // TODO(adonovan): opt: use ssa.blockSet
	for _, b := range fn.DomPreorder() {
		if x, k := isComparisonBlock(b); x != nil {
			// Block b starts a switch.
			sw := Switch{Start: b, X: x}
			valueSwitch(&sw, k, seen)
			if len(sw.ConstCases) > 1 {
				switches = append(switches, sw)
			}
		}

		if y, x, T := isTypeAssertBlock(b); y != nil {
			// Block b starts a type switch.
			sw := Switch{Start: b, X: x}
			typeSwitch(&sw, y, T, seen)
			if len(sw.TypeCases) > 1 {
				switches = append(switches, sw)
			}
		}
	}
	return switches
}

func valueSwitch(sw *Switch, k *ssa.Const, seen map[*ssa.BasicBlock]bool) {
	b := sw.Start
	x := sw.X
	for x == sw.X {
		if seen[b] {
			break
		}
		seen[b] = true

		sw.ConstCases = append(sw.ConstCases, ConstCase{
			Block: b,
			Body:  b.Succs[0],
			Value: k,
		})
		b = b.Succs[1]
		if len(b.Instrs) > 2 {
			// Block b contains not just 'if x == k',
			// so it may have side effects that
			// make it unsafe to elide.
			break
		}
		if len(b.Preds) != 1 {
			// Block b has multiple predecessors,
			// so it cannot be treated as a case.
			break
		}
		x, k = isComparisonBlock(b)
	}
	sw.Default = b
}

func typeSwitch(sw *Switch, y ssa.Value, T types.Type, seen map[*ssa.BasicBlock]bool) {
	b := sw.Start
	x := sw.X
	for x == sw.X {
		if seen[b] {
			break
		}
		seen[b] = true

		sw.TypeCases = append(sw.TypeCases, TypeCase{
			Block:   b,
			Body:    b.Succs[0],
			Type:    T,
			Binding: y,
		})
		b = b.Succs[1]
		if len(b.Instrs) > 4 {
			// Block b contains not just
			//  {TypeAssert; Extract #0; Extract #1; If}
			// so it may have side effects that
			// make it unsafe to elide.
			break
		}
		if len(b.Preds) != 1 {
			// Block b has multiple predecessors,
			// so it cannot be treated as a case.
			break
		}
		y, x, T = isTypeAssertBlock(b)
	}
	sw.Default = b
}

// isComparisonBlock returns the operands (v, k) if a block ends with
// a comparison v==k, where k is a compile-time constant.
func isComparisonBlock(b *ssa.BasicBlock) (v ssa.Value, k *ssa.Const) {
	if n := len(b.Instrs); n >= 2 {
		if i, ok := b.Instrs[n-1].(*ssa.If); ok {
			if binop, ok := i.Cond.(*ssa.BinOp); ok && binop.Block() == b && binop.Op == token.EQL {
				if k, ok := binop.Y.(*ssa.Const); ok {
					return binop.X, k
				}
				if k, ok := binop.X.(*ssa.Const); ok {
					return binop.Y, k
				}
			}
		}
	}
	return
}

// isTypeAssertBlock returns the operands (y, x, T) if a block ends with
// a type assertion "if y, ok := x.(T); ok {".
func isTypeAssertBlock(b *ssa.BasicBlock) (y, x ssa.Value, T types.Type) {
	if n := len(b.Instrs); n >= 4 {
		if i, ok := b.Instrs[n-1].(*ssa.If); ok {
			if ext1, ok := i.Cond.(*ssa.Extract); ok && ext1.Block() == b && ext1.Index == 1 {
				if ta, ok := ext1.Tuple.(*ssa.TypeAssert); ok && ta.Block() == b {
					// hack: relies upon instruction ordering.
					if ext0, ok := b.Instrs[n-3].(*ssa.Extract); ok {
						return ext0, ta.X, ta.AssertedType
					}
				}
			}
		}
	}
	return
}
```

## File: go/ssa/ssautil/visit.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil // import "golang.org/x/tools/go/ssa/ssautil"

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/ssa"

	_ "unsafe" // for linkname hack
)

// This file defines utilities for visiting the SSA representation of
// a Program.
//
// TODO(adonovan): test coverage.

// AllFunctions finds and returns the set of functions potentially
// needed by program prog, as determined by a simple linker-style
// reachability algorithm starting from the members and method-sets of
// each package.  The result may include anonymous functions and
// synthetic wrappers.
//
// Precondition: all packages are built.
//
// TODO(adonovan): this function is underspecified. It doesn't
// actually work like a linker, which computes reachability from main
// using something like go/callgraph/cha (without materializing the
// call graph). In fact, it treats all public functions and all
// methods of public non-parameterized types as roots, even though
// they may be unreachable--but only in packages created from syntax.
//
// I think we should deprecate AllFunctions function in favor of two
// clearly defined ones:
//
//  1. The first would efficiently compute CHA reachability from a set
//     of main packages, making it suitable for a whole-program
//     analysis context with InstantiateGenerics, in conjunction with
//     Program.Build.
//
//  2. The second would return only the set of functions corresponding
//     to source Func{Decl,Lit} syntax, like SrcFunctions in
//     go/analysis/passes/buildssa; this is suitable for
//     package-at-a-time (or handful of packages) context.
//     ssa.Package could easily expose it as a field.
//
// We could add them unexported for now and use them via the linkname hack.
func AllFunctions(prog *ssa.Program) map[*ssa.Function]bool {
	seen := make(map[*ssa.Function]bool)

	var function func(fn *ssa.Function)
	function = func(fn *ssa.Function) {
		if !seen[fn] {
			seen[fn] = true
			var buf [10]*ssa.Value // avoid alloc in common case
			for _, b := range fn.Blocks {
				for _, instr := range b.Instrs {
					for _, op := range instr.Operands(buf[:0]) {
						if fn, ok := (*op).(*ssa.Function); ok {
							function(fn)
						}
					}
				}
			}
		}
	}

	// TODO(adonovan): opt: provide a way to share a builder
	// across a sequence of MethodValue calls.

	methodsOf := func(T types.Type) {
		if !types.IsInterface(T) {
			mset := prog.MethodSets.MethodSet(T)
			for i := 0; i < mset.Len(); i++ {
				function(prog.MethodValue(mset.At(i)))
			}
		}
	}

	// Historically, Program.RuntimeTypes used to include the type
	// of any exported member of a package loaded from syntax that
	// has a non-parameterized type, plus all types
	// reachable from that type using reflection, even though
	// these runtime types may not be required for them.
	//
	// Rather than break existing programs that rely on
	// AllFunctions visiting extra methods that are unreferenced
	// by IR and unreachable via reflection, we moved the logic
	// here, unprincipled though it is.
	// (See doc comment for better ideas.)
	//
	// Nonetheless, after the move, we no longer visit every
	// method of any type recursively reachable from T, only the
	// methods of T and *T themselves, and we only apply this to
	// named types T, and not to the type of every exported
	// package member.
	exportedTypeHack := func(t *ssa.Type) {
		if isSyntactic(t.Package()) &&
			ast.IsExported(t.Name()) &&
			!types.IsInterface(t.Type()) {
			// Consider only named types.
			// (Ignore aliases and unsafe.Pointer.)
			if named, ok := t.Type().(*types.Named); ok {
				if named.TypeParams() == nil {
					methodsOf(named)                   //  T
					methodsOf(types.NewPointer(named)) // *T
				}
			}
		}
	}

	for _, pkg := range prog.AllPackages() {
		for _, mem := range pkg.Members {
			switch mem := mem.(type) {
			case *ssa.Function:
				// Visit all package-level declared functions.
				function(mem)

			case *ssa.Type:
				exportedTypeHack(mem)
			}
		}
	}

	// Visit all methods of types for which runtime types were
	// materialized, as they are reachable through reflection.
	for _, T := range prog.RuntimeTypes() {
		methodsOf(T)
	}

	return seen
}

// MainPackages returns the subset of the specified packages
// named "main" that define a main function.
// The result may include synthetic "testmain" packages.
func MainPackages(pkgs []*ssa.Package) []*ssa.Package {
	var mains []*ssa.Package
	for _, pkg := range pkgs {
		if pkg.Pkg.Name() == "main" && pkg.Func("main") != nil {
			mains = append(mains, pkg)
		}
	}
	return mains
}

// TODO(adonovan): propose a principled API for this. One possibility
// is a new field, Package.SrcFunctions []*Function, which would
// contain the list of SrcFunctions described in point 2 of the
// AllFunctions doc comment, or nil if the package is not from syntax.
// But perhaps overloading nil vs empty slice is too subtle.
//
//go:linkname isSyntactic golang.org/x/tools/go/ssa.isSyntactic
func isSyntactic(pkg *ssa.Package) bool
```

## File: go/ssa/testdata/fixedbugs/issue66783a.go
```go
//go:build ignore
// +build ignore

package issue66783a

type S[T any] struct {
	a T
}

func (s S[T]) M() {
	type A S[T]
	type B[U any] A
	_ = B[rune](s)
}

// M[int]

// panic: in (issue66783a.S[int]).M[int]:
// cannot convert term *t0 (issue66783a.S[int] [within struct{a int}])
// to type issue66783a.B[rune] [within struct{a T}] [recovered]

func M() {
	S[int]{}.M()
}
```

## File: go/ssa/testdata/fixedbugs/issue66783b.go
```go
//go:build ignore
// +build ignore

package issue66783b

type I1[T any] interface {
	M(T)
}

type I2[T any] I1[T]

func foo[T any](i I2[T]) {
	_ = i.M
}

type S[T any] struct{}

func (s S[T]) M(t T) {}

func M2() {
	foo[int](I2[int](S[int]{}))
}
```

## File: go/ssa/testdata/src/bytes/bytes.go
```go
package bytes

func Compare(a, b []byte) int
```

## File: go/ssa/testdata/src/context/context.go
```go
package context

type Context interface {
	Done() <-chan struct{}
}

func Background() Context
```

## File: go/ssa/testdata/src/encoding/json/json.go
```go
package json

func Marshal(v any) ([]byte, error)
func Unmarshal(data []byte, v any) error
```

## File: go/ssa/testdata/src/encoding/xml/xml.go
```go
package xml

func Marshal(v any) ([]byte, error)
func Unmarshal(data []byte, v any) error
```

## File: go/ssa/testdata/src/encoding/encoding.go
```go
package encoding

type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}

type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}
```

## File: go/ssa/testdata/src/errors/errors.go
```go
package errors

func New(text string) error
```

## File: go/ssa/testdata/src/fmt/fmt.go
```go
package fmt

func Sprint(args ...interface{}) string
func Sprintln(args ...interface{}) string
func Sprintf(format string, args ...interface{}) string

func Print(args ...interface{}) (int, error)
func Println(args ...interface{})
func Printf(format string, args ...interface{}) (int, error)

func Errorf(format string, args ...interface{}) error
```

## File: go/ssa/testdata/src/io/io.go
```go
package io

import "errors"

var EOF = errors.New("EOF")
```

## File: go/ssa/testdata/src/log/log.go
```go
package log

func Println(v ...interface{})
func Fatalln(v ...interface{})
func Fatalf(format string, v ...any)
```

## File: go/ssa/testdata/src/math/math.go
```go
package math

func NaN() float64

func Inf(int) float64

func IsNaN(float64) bool

func Float64bits(float64) uint64

func Signbit(x float64) bool

func Sqrt(x float64) float64

func Sin(x float64) float64
```

## File: go/ssa/testdata/src/os/os.go
```go
package os

func Getenv(string) string

func Exit(int)
```

## File: go/ssa/testdata/src/reflect/reflect.go
```go
package reflect

type Type interface {
	Elem() Type
	Kind() Kind
	String() string
}

type Value struct{}

func (Value) String() string
func (Value) Elem() Value
func (Value) Field(int) Value
func (Value) Index(i int) Value
func (Value) Int() int64
func (Value) Interface() interface{}
func (Value) IsNil() bool
func (Value) IsValid() bool
func (Value) Kind() Kind
func (Value) Len() int
func (Value) MapIndex(Value) Value
func (Value) MapKeys() []Value
func (Value) NumField() int
func (Value) Pointer() uintptr
func (Value) SetInt(int64)
func (Value) Type() Type

func SliceOf(Type) Type
func TypeOf(interface{}) Type
func ValueOf(interface{}) Value

type Kind uint

const (
	Invalid Kind = iota
	Int
	Pointer
)

func DeepEqual(x, y interface{}) bool
```

## File: go/ssa/testdata/src/runtime/runtime.go
```go
package runtime

func GC()

func SetFinalizer(obj, finalizer any)

func Caller(skip int) (pc uintptr, file string, line int, ok bool)
```

## File: go/ssa/testdata/src/sort/sort.go
```go
package sort

func Strings(x []string)
func Ints(x []int)
func Float64s(x []float64)

func Sort(data Interface)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

## File: go/ssa/testdata/src/strconv/strconv.go
```go
package strconv

func Itoa(i int) string
func Atoi(s string) (int, error)

func FormatFloat(float64, byte, int, int) string
```

## File: go/ssa/testdata/src/strings/strings.go
```go
package strings

func Replace(s, old, new string, n int) string
func Index(haystack, needle string) int
func Contains(haystack, needle string) bool
func HasPrefix(s, prefix string) bool
func EqualFold(s, t string) bool
func ToLower(s string) string

type Builder struct{}

func (b *Builder) WriteString(s string) (int, error)
func (b *Builder) String() string
```

## File: go/ssa/testdata/src/sync/atomic/atomic.go
```go
package atomic

import "unsafe"

func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
```

## File: go/ssa/testdata/src/sync/sync.go
```go
package sync

type Mutex struct{}

func (m *Mutex) Lock()
func (m *Mutex) Unlock()

type WaitGroup struct{}

func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
```

## File: go/ssa/testdata/src/time/time.go
```go
package time

type Duration int64

func Sleep(Duration)

func NewTimer(d Duration) *Timer

type Timer struct {
	C <-chan Time
}

func (t *Timer) Stop() bool

type Time struct{}

func After(d Duration) <-chan Time

const (
	Nanosecond Duration = iota // Specific values do not matter here.
	Second
	Minute
	Hour
)
```

## File: go/ssa/testdata/src/unsafe/unsafe.go
```go
package unsafe

// Empty unsafe package helps other packages load.
// TODO(taking): determine why.
```

## File: go/ssa/testdata/src/README.txt
```
These files are present to test building ssa on go files that use signatures from standard library packages.

Only the exported members used by the tests are needed.

Providing these decreases testing time ~10x (90s -> 8s) compared to building the standard library packages form source during tests.
```

## File: go/ssa/testdata/indirect.txtar
```
-- go.mod --
module testdata
go 1.18

-- a/a.go --
package a

import "testdata/b"

func A() {
	var x b.B
	x.F()
}

-- b/b.go --
package b

import "testdata/c"

type B struct { c.C }

-- c/c.go --
package c

type C int
func (C) F() {}
```

## File: go/ssa/testdata/objlookup.go
```go
package main

// This file is the input to TestObjValueLookup in source_test.go,
// which ensures that each occurrence of an ident defining or
// referring to a func, var or const object can be mapped to its
// corresponding SSA Value.
//
// For every reference to a var object, we use annotations in comments
// to denote both the expected SSA Value kind, and whether to expect
// its value (x) or its address (&x).
//
// For const and func objects, the results don't vary by reference and
// are always values not addresses, so no annotations are needed.  The
// declaration is enough.

import (
	"fmt"
	"os"
)

type J int

func (*J) method() {}

const globalConst = 0

var globalVar int //@ ssa(globalVar,"&Global")

func globalFunc() {}

type I interface {
	interfaceMethod()
}

type S struct {
	x int //@ ssa(x,"nil")
}

func main() {
	print(globalVar) //@ ssa(globalVar,"UnOp")
	globalVar = 1    //@ ssa(globalVar,"Const")

	var v0 int = 1 //@ ssa(v0,"Const") // simple local value spec
	if v0 > 0 {    //@ ssa(v0,"Const")
		v0 = 2 //@ ssa(v0,"Const")
	}
	print(v0) //@ ssa(v0,"Phi")

	// v1 is captured and thus implicitly address-taken.
	var v1 int = 1         //@ ssa(v1,"Const")
	v1 = 2                 //@ ssa(v1,"Const")
	fmt.Println(v1)        //@ ssa(v1,"UnOp") // load
	f := func(param int) { //@ ssa(f,"MakeClosure"), ssa(param,"Parameter")
		if y := 1; y > 0 { //@ ssa(y,"Const")
			print(v1, param) //@ ssa(v1,"UnOp") /*load*/, ssa(param,"Parameter")
		}
		param = 2      //@ ssa(param,"Const")
		println(param) //@ ssa(param,"Const")
	}

	f(0) //@ ssa(f,"MakeClosure")

	var v2 int //@ ssa(v2,"Const") // implicitly zero-initialized local value spec
	print(v2)  //@ ssa(v2,"Const")

	m := make(map[string]int) //@ ssa(m,"MakeMap")

	// Local value spec with multi-valued RHS:
	var v3, v4 = m[""] //@ ssa(v3,"Extract"), ssa(v4,"Extract"), ssa(m,"MakeMap")
	print(v3)          //@ ssa(v3,"Extract")
	print(v4)          //@ ssa(v4,"Extract")

	v3++    //@ ssa(v3,"BinOp") // assign with op
	v3 += 2 //@ ssa(v3,"BinOp") // assign with op

	v5, v6 := false, "" //@ ssa(v5,"Const"), ssa(v6,"Const") // defining assignment
	print(v5)           //@ ssa(v5,"Const")
	print(v6)           //@ ssa(v6,"Const")

	var v7 S    //@ ssa(v7,"&Alloc")
	v7.x = 1    //@ ssa(v7,"&Alloc"), ssa(x,"&FieldAddr")
	print(v7.x) //@ ssa(v7,"&Alloc"), ssa(x,"&FieldAddr")

	var v8 [1]int //@ ssa(v8,"&Alloc")
	v8[0] = 0     //@ ssa(v8,"&Alloc")
	print(v8[:])  //@ ssa(v8,"&Alloc")
	_ = v8[0]     //@ ssa(v8,"&Alloc")
	_ = v8[:][0]  //@ ssa(v8,"&Alloc")
	v8ptr := &v8  //@ ssa(v8ptr,"Alloc"), ssa(v8,"&Alloc")
	_ = v8ptr[0]  //@ ssa(v8ptr,"Alloc")
	_ = *v8ptr    //@ ssa(v8ptr,"Alloc")

	v8a := make([]int, 1) //@ ssa(v8a,"Slice")
	v8a[0] = 0            //@ ssa(v8a,"Slice")
	print(v8a[:])         //@ ssa(v8a,"Slice")

	v9 := S{} //@ ssa(v9,"&Alloc")

	v10 := &v9 //@ ssa(v10,"Alloc"), ssa(v9,"&Alloc")
	_ = v10    //@ ssa(v10,"Alloc")

	var v11 *J = nil //@ ssa(v11,"Const")
	v11.method()     //@ ssa(v11,"Const")

	var v12 J    //@ ssa(v12,"&Alloc")
	v12.method() //@ ssa(v12,"&Alloc") // implicitly address-taken

	// NB, in the following, 'method' resolves to the *types.Func
	// of (*J).method, so it doesn't help us locate the specific
	// ssa.Values here: a bound-method closure and a promotion
	// wrapper.
	_ = v11.method            //@ ssa(v11,"Const")
	_ = (*struct{ J }).method //@ ssa(J,"nil")

	// These vars are not optimised away.
	if false {
		v13 := 0     //@ ssa(v13,"Const")
		println(v13) //@ ssa(v13,"Const")
	}

	switch x := 1; x { //@ ssa(x,"Const")
	case v0: //@ ssa(v0,"Phi")
	}

	for k, v := range m { //@ ssa(k,"Extract"), ssa(v,"Extract"), ssa(m,"MakeMap")
		_ = k //@ ssa(k,"Extract")
		v++   //@ ssa(v,"BinOp")
	}

	if y := 0; y > 1 { //@ ssa(y,"Const"), ssa(y,"Const")
	}

	var i interface{}      //@ ssa(i,"Const") // nil interface
	i = 1                  //@ ssa(i,"MakeInterface")
	switch i := i.(type) { //@ ssa(i,"MakeInterface"), ssa(i,"MakeInterface")
	case int:
		println(i) //@ ssa(i,"Extract")
	}

	ch := make(chan int) //@ ssa(ch,"MakeChan")
	select {
	case x := <-ch: //@ ssa(x,"UnOp") /*receive*/, ssa(ch,"MakeChan")
		_ = x //@ ssa(x,"UnOp")
	}

	// .Op is an inter-package FieldVal-selection.
	var err os.PathError //@ ssa(err,"&Alloc")
	_ = err.Op           //@ ssa(err,"&Alloc"), ssa(Op,"&FieldAddr")
	_ = &err.Op          //@ ssa(err,"&Alloc"), ssa(Op,"&FieldAddr")

	// Exercise corner-cases of lvalues vs rvalues.
	// (Guessing IsAddr from the 'pointerness' won't cut it here.)
	type N *N
	var n N    //@ ssa(n,"Const")
	n1 := n    //@ ssa(n1,"Const"), ssa(n,"Const")
	n2 := &n1  //@ ssa(n2,"Alloc"), ssa(n1,"&Alloc")
	n3 := *n2  //@ ssa(n3,"UnOp"), ssa(n2,"Alloc")
	n4 := **n3 //@ ssa(n4,"UnOp"), ssa(n3,"UnOp")
	_ = n4     //@ ssa(n4,"UnOp")
}
```

## File: go/ssa/testdata/structconv.go
```go
// This file is the input to TestValueForExprStructConv in identical_test.go,
// which uses the same framework as TestValueForExpr does in source_test.go.
//
// In Go 1.8, struct conversions are permitted even when the struct types have
// different tags. This wasn't permitted in earlier versions of Go, so this file
// exists separately from valueforexpr.go to just test this behavior in Go 1.8
// and later.

package main

type t1 struct {
	x int
}
type t2 struct {
	x int `tag`
}

func main() {
	var tv1 t1
	var tv2 t2 = /*@ChangeType*/ (t2(tv1))
	_ = tv2
}
```

## File: go/ssa/testdata/valueforexpr.go
```go
package main

// This file is the input to TestValueForExpr in source_test.go, which
// ensures that each expression e immediately following a /*@kind*/(x)
// annotation, when passed to Function.ValueForExpr(e), returns a
// non-nil Value of the same type as e and of kind 'kind'.

func f(spilled, unspilled int) {
	_ = /*@UnOp*/ (spilled)
	_ = /*@Parameter*/ (unspilled)
	_ = /*@nil*/ (1 + 2) // (constant)
	i := 0

	f := func() (int, int) { return 0, 0 }

	/*@Call*/
	(print( /*@BinOp*/ (i + 1)))
	_, _ = /*@Call*/ (f())
	ch := /*@MakeChan*/ (make(chan int))
	/*@UnOp*/ (<-ch)
	x := /*@UnOp*/ (<-ch)
	_ = x
	select {
	case /*@Extract*/ (<-ch):
	case x := /*@Extract*/ (<-ch):
		_ = x
	}
	defer /*@Function*/ (func() {
	})()
	go /*@Function*/ (func() {
	})()
	y := 0
	if true && /*@BinOp*/ (bool(y > 0)) {
		y = 1
	}
	_ = /*@Phi*/ (y)
	map1 := /*@MakeMap*/ (make(map[string]string))
	_ = map1
	_ = /*@Slice*/ (make([]int, 0))
	_ = /*@MakeClosure*/ (func() { print(spilled) })

	sl := []int{}
	_ = /*@Slice*/ (sl[:0])

	_ = /*@nil*/ (new(int)) // optimized away
	tmp := /*@Alloc*/ (new(int))
	_ = tmp
	var iface interface{}
	_ = /*@TypeAssert*/ (iface.(int))
	_ = /*@UnOp*/ (sl[0])
	_ = /*@IndexAddr*/ (&sl[0])
	_ = /*@Index*/ ([2]int{}[0])
	var p *int
	_ = /*@UnOp*/ (*p)

	_ = /*@UnOp*/ (global)
	/*@UnOp*/ (global)[""] = ""
	/*@Global*/ (global) = map[string]string{}

	var local t
	/*UnOp*/ (local.x) = 1

	// Exercise corner-cases of lvalues vs rvalues.
	type N *N
	var n N
	/*@UnOp*/ (n) = /*@UnOp*/ (n)
	/*@ChangeType*/ (n) = /*@Alloc*/ (&n)
	/*@UnOp*/ (n) = /*@UnOp*/ (*n)
	/*@UnOp*/ (n) = /*@UnOp*/ (**n)
}

func complit() {
	// Composite literals.
	// We get different results for
	// - composite literal as value (e.g. operand to print)
	// - composite literal initializer for addressable value
	// - composite literal value assigned to blank var

	// 1. Slices
	print( /*@Slice*/ ([]int{}))
	print( /*@Alloc*/ (&[]int{}))
	print(& /*@Slice*/ ([]int{}))

	sl1 := /*@Slice*/ ([]int{})
	sl2 := /*@Alloc*/ (&[]int{})
	sl3 := & /*@Slice*/ ([]int{})
	_, _, _ = sl1, sl2, sl3

	_ = /*@Slice*/ ([]int{})
	_ = /*@nil*/ (& /*@Slice*/ ([]int{})) // & optimized away
	_ = & /*@Slice*/ ([]int{})

	// 2. Arrays
	print( /*@Const*/ ([1]int{}))
	print( /*@Alloc*/ (&[1]int{}))
	print(& /*@Alloc*/ ([1]int{}))

	arr1 := /*@Const*/ ([1]int{})
	arr2 := /*@Alloc*/ (&[1]int{})
	arr3 := & /*@Alloc*/ ([1]int{})
	_, _, _ = arr1, arr2, arr3

	_ = /*@Const*/ ([1]int{})
	_ = /*@nil*/ (& /*@Const*/ ([1]int{})) // & optimized away
	_ = & /*@Const*/ ([1]int{})

	// 3. Maps
	type M map[int]int
	print( /*@MakeMap*/ (M{}))
	print( /*@Alloc*/ (&M{}))
	print(& /*@MakeMap*/ (M{}))

	m1 := /*@MakeMap*/ (M{})
	m2 := /*@Alloc*/ (&M{})
	m3 := & /*@MakeMap*/ (M{})
	_, _, _ = m1, m2, m3

	_ = /*@MakeMap*/ (M{})
	_ = /*@nil*/ (& /*@MakeMap*/ (M{})) // & optimized away
	_ = & /*@MakeMap*/ (M{})

	// 4. Structs
	print( /*@Const*/ (struct{}{}))
	print( /*@Alloc*/ (&struct{}{}))
	print(& /*@Alloc*/ (struct{}{}))

	s1 := /*@Const*/ (struct{}{})
	s2 := /*@Alloc*/ (&struct{}{})
	s3 := & /*@Alloc*/ (struct{}{})
	_, _, _ = s1, s2, s3

	_ = /*@Const*/ (struct{}{})
	_ = /*@nil*/ (& /*@Const*/ (struct{}{})) // & optimized away
	_ = & /*@Const*/ (struct{}{})
}

type t struct{ x int }

// Ensure we can locate methods of named types.
func (t) f(param int) {
	_ = /*@Parameter*/ (param)
}

// Ensure we can locate init functions.
func init() {
	m := /*@MakeMap*/ (make(map[string]string))
	_ = m
}

// Ensure we can locate variables in initializer expressions.
var global = /*@MakeMap*/ (make(map[string]string))
```

## File: go/ssa/block.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "fmt"

// This file implements the BasicBlock type.

// addEdge adds a control-flow graph edge from from to to.
func addEdge(from, to *BasicBlock) {
	from.Succs = append(from.Succs, to)
	to.Preds = append(to.Preds, from)
}

// Parent returns the function that contains block b.
func (b *BasicBlock) Parent() *Function { return b.parent }

// String returns a human-readable label of this block.
// It is not guaranteed unique within the function.
func (b *BasicBlock) String() string {
	return fmt.Sprintf("%d", b.Index)
}

// emit appends an instruction to the current basic block.
// If the instruction defines a Value, it is returned.
func (b *BasicBlock) emit(i Instruction) Value {
	i.setBlock(b)
	b.Instrs = append(b.Instrs, i)
	v, _ := i.(Value)
	return v
}

// predIndex returns the i such that b.Preds[i] == c or panics if
// there is none.
func (b *BasicBlock) predIndex(c *BasicBlock) int {
	for i, pred := range b.Preds {
		if pred == c {
			return i
		}
	}
	panic(fmt.Sprintf("no edge %s -> %s", c, b))
}

// hasPhi returns true if b.Instrs contains φ-nodes.
func (b *BasicBlock) hasPhi() bool {
	_, ok := b.Instrs[0].(*Phi)
	return ok
}

// phis returns the prefix of b.Instrs containing all the block's φ-nodes.
func (b *BasicBlock) phis() []Instruction {
	for i, instr := range b.Instrs {
		if _, ok := instr.(*Phi); !ok {
			return b.Instrs[:i]
		}
	}
	return nil // unreachable in well-formed blocks
}

// replacePred replaces all occurrences of p in b's predecessor list with q.
// Ordinarily there should be at most one.
func (b *BasicBlock) replacePred(p, q *BasicBlock) {
	for i, pred := range b.Preds {
		if pred == p {
			b.Preds[i] = q
		}
	}
}

// replaceSucc replaces all occurrences of p in b's successor list with q.
// Ordinarily there should be at most one.
func (b *BasicBlock) replaceSucc(p, q *BasicBlock) {
	for i, succ := range b.Succs {
		if succ == p {
			b.Succs[i] = q
		}
	}
}

// removePred removes all occurrences of p in b's
// predecessor list and φ-nodes.
// Ordinarily there should be at most one.
func (b *BasicBlock) removePred(p *BasicBlock) {
	phis := b.phis()

	// We must preserve edge order for φ-nodes.
	j := 0
	for i, pred := range b.Preds {
		if pred != p {
			b.Preds[j] = b.Preds[i]
			// Strike out φ-edge too.
			for _, instr := range phis {
				phi := instr.(*Phi)
				phi.Edges[j] = phi.Edges[i]
			}
			j++
		}
	}
	// Nil out b.Preds[j:] and φ-edges[j:] to aid GC.
	for i := j; i < len(b.Preds); i++ {
		b.Preds[i] = nil
		for _, instr := range phis {
			instr.(*Phi).Edges[i] = nil
		}
	}
	b.Preds = b.Preds[:j]
	for _, instr := range phis {
		phi := instr.(*Phi)
		phi.Edges = phi.Edges[:j]
	}
}
```

## File: go/ssa/blockopt.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// Simple block optimizations to simplify the control flow graph.

// TODO(adonovan): opt: instead of creating several "unreachable" blocks
// per function in the Builder, reuse a single one (e.g. at Blocks[1])
// to reduce garbage.

import (
	"fmt"
	"os"
)

// If true, perform sanity checking and show progress at each
// successive iteration of optimizeBlocks.  Very verbose.
const debugBlockOpt = false

// markReachable sets Index=-1 for all blocks reachable from b.
func markReachable(b *BasicBlock) {
	b.Index = -1
	for _, succ := range b.Succs {
		if succ.Index == 0 {
			markReachable(succ)
		}
	}
}

// deleteUnreachableBlocks marks all reachable blocks of f and
// eliminates (nils) all others, including possibly cyclic subgraphs.
func deleteUnreachableBlocks(f *Function) {
	const white, black = 0, -1
	// We borrow b.Index temporarily as the mark bit.
	for _, b := range f.Blocks {
		b.Index = white
	}
	markReachable(f.Blocks[0])
	if f.Recover != nil {
		markReachable(f.Recover)
	}
	for i, b := range f.Blocks {
		if b.Index == white {
			for _, c := range b.Succs {
				if c.Index == black {
					c.removePred(b) // delete white->black edge
				}
			}
			if debugBlockOpt {
				fmt.Fprintln(os.Stderr, "unreachable", b)
			}
			f.Blocks[i] = nil // delete b
		}
	}
	f.removeNilBlocks()
}

// jumpThreading attempts to apply simple jump-threading to block b,
// in which a->b->c become a->c if b is just a Jump.
// The result is true if the optimization was applied.
func jumpThreading(f *Function, b *BasicBlock) bool {
	if b.Index == 0 {
		return false // don't apply to entry block
	}
	if b.Instrs == nil {
		return false
	}
	if _, ok := b.Instrs[0].(*Jump); !ok {
		return false // not just a jump
	}
	c := b.Succs[0]
	if c == b {
		return false // don't apply to degenerate jump-to-self.
	}
	if c.hasPhi() {
		return false // not sound without more effort
	}
	for j, a := range b.Preds {
		a.replaceSucc(b, c)

		// If a now has two edges to c, replace its degenerate If by Jump.
		if len(a.Succs) == 2 && a.Succs[0] == c && a.Succs[1] == c {
			jump := new(Jump)
			jump.setBlock(a)
			a.Instrs[len(a.Instrs)-1] = jump
			a.Succs = a.Succs[:1]
			c.removePred(b)
		} else {
			if j == 0 {
				c.replacePred(b, a)
			} else {
				c.Preds = append(c.Preds, a)
			}
		}

		if debugBlockOpt {
			fmt.Fprintln(os.Stderr, "jumpThreading", a, b, c)
		}
	}
	f.Blocks[b.Index] = nil // delete b
	return true
}

// fuseBlocks attempts to apply the block fusion optimization to block
// a, in which a->b becomes ab if len(a.Succs)==len(b.Preds)==1.
// The result is true if the optimization was applied.
func fuseBlocks(f *Function, a *BasicBlock) bool {
	if len(a.Succs) != 1 {
		return false
	}
	b := a.Succs[0]
	if len(b.Preds) != 1 {
		return false
	}

	// Degenerate &&/|| ops may result in a straight-line CFG
	// containing φ-nodes. (Ideally we'd replace such them with
	// their sole operand but that requires Referrers, built later.)
	if b.hasPhi() {
		return false // not sound without further effort
	}

	// Eliminate jump at end of A, then copy all of B across.
	a.Instrs = append(a.Instrs[:len(a.Instrs)-1], b.Instrs...)
	for _, instr := range b.Instrs {
		instr.setBlock(a)
	}

	// A inherits B's successors
	a.Succs = append(a.succs2[:0], b.Succs...)

	// Fix up Preds links of all successors of B.
	for _, c := range b.Succs {
		c.replacePred(b, a)
	}

	if debugBlockOpt {
		fmt.Fprintln(os.Stderr, "fuseBlocks", a, b)
	}

	f.Blocks[b.Index] = nil // delete b
	return true
}

// optimizeBlocks() performs some simple block optimizations on a
// completed function: dead block elimination, block fusion, jump
// threading.
func optimizeBlocks(f *Function) {
	deleteUnreachableBlocks(f)

	// Loop until no further progress.
	changed := true
	for changed {
		changed = false

		if debugBlockOpt {
			f.WriteTo(os.Stderr)
			mustSanityCheck(f, nil)
		}

		for _, b := range f.Blocks {
			// f.Blocks will temporarily contain nils to indicate
			// deleted blocks; we remove them at the end.
			if b == nil {
				continue
			}

			// Fuse blocks.  b->c becomes bc.
			if fuseBlocks(f, b) {
				changed = true
			}

			// a->b->c becomes a->c if b contains only a Jump.
			if jumpThreading(f, b) {
				changed = true
				continue // (b was disconnected)
			}
		}
	}
	f.removeNilBlocks()
}
```

## File: go/ssa/builder_generic_test.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"reflect"
	"sort"
	"testing"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/internal/expect"
)

// TestGenericBodies tests that bodies of generic functions and methods containing
// different constructs can be built in BuilderMode(0).
//
// Each test specifies the contents of package containing a single go file.
// Each call print(arg0, arg1, ...) to the builtin print function
// in ssa is correlated a comment at the end of the line of the form:
//
//	//@ types(a, b, c)
//
// where a, b and c are the types of the arguments to the print call
// serialized using go/types.Type.String().
// See x/tools/internal/expect for details on the syntax.
func TestGenericBodies(t *testing.T) {
	for _, content := range []string{
		`
		package p00

		func f(x int) {
			var i interface{}
			print(i, 0) //@ types("interface{}", int)
			print()     //@ types()
			print(x)    //@ types(int)
		}
		`,
		`
		package p01

		func f[T any](x T) {
			print(x) //@ types(T)
		}
		`,
		`
		package p02

		func f[T ~int]() {
			var x T
			print(x) //@ types(T)
		}
		`,
		`
		package p03

		func a[T ~[4]byte](x T) {
			for k, v := range x {
				print(x, k, v) //@ types(T, int, byte)
			}
		}
		func b[T ~*[4]byte](x T) {
			for k, v := range x {
				print(x, k, v) //@ types(T, int, byte)
			}
		}
		func c[T ~[]byte](x T) {
			for k, v := range x {
				print(x, k, v) //@ types(T, int, byte)
			}
		}
		func d[T ~string](x T) {
			for k, v := range x {
				print(x, k, v) //@ types(T, int, rune)
			}
		}
		func e[T ~map[int]string](x T) {
			for k, v := range x {
				print(x, k, v) //@ types(T, int, string)
			}
		}
		func f[T ~chan string](x T) {
			for v := range x {
				print(x, v) //@ types(T, string)
			}
		}

		func From() {
			type A [4]byte
			print(a[A]) //@ types("func(x p03.A)")

			type B *[4]byte
			print(b[B]) //@ types("func(x p03.B)")

			type C []byte
			print(c[C]) //@ types("func(x p03.C)")

			type D string
			print(d[D]) //@ types("func(x p03.D)")

			type E map[int]string
			print(e[E]) //@ types("func(x p03.E)")

			type F chan string
			print(f[F]) //@ types("func(x p03.F)")
		}
		`,
		`
		package p05

		func f[S any, T ~chan S](x T) {
			for v := range x {
				print(x, v) //@ types(T, S)
			}
		}

		func From() {
			type F chan string
			print(f[string, F]) //@ types("func(x p05.F)")
		}
		`,
		`
		package p06

		func fibonacci[T ~chan int](c, quit T) {
			x, y := 0, 1
			for {
				select {
				case c <- x:
					x, y = y, x+y
				case <-quit:
					print(c, quit, x, y) //@ types(T, T, int, int)
					return
				}
			}
		}
		func start[T ~chan int](c, quit T) {
			go func() {
				for i := 0; i < 10; i++ {
					print(<-c) //@ types(int)
				}
				quit <- 0
			}()
		}
		func From() {
			type F chan int
			c := make(F)
			quit := make(F)
			print(start[F], c, quit)     //@ types("func(c p06.F, quit p06.F)", "p06.F", "p06.F")
			print(fibonacci[F], c, quit) //@ types("func(c p06.F, quit p06.F)", "p06.F", "p06.F")
		}
		`,
		`
		package p07

		func f[T ~struct{ x int; y string }](i int) T {
			u := []T{ T{0, "lorem"},  T{1, "ipsum"}}
			return u[i]
		}
		func From() {
			type S struct{ x int; y string }
			print(f[S])     //@ types("func(i int) p07.S")
		}
		`,
		`
		package p08

		func f[T ~[4]int8](x T, l, h int) []int8 {
			return x[l:h]
		}
		func g[T ~*[4]int16](x T, l, h int) []int16 {
			return x[l:h]
		}
		func h[T ~[]int32](x T, l, h int) T {
			return x[l:h]
		}
		func From() {
			type F [4]int8
			type G *[4]int16
			type H []int32
			print(f[F](F{}, 0, 0))  //@ types("[]int8")
			print(g[G](nil, 0, 0)) //@ types("[]int16")
			print(h[H](nil, 0, 0)) //@ types("p08.H")
		}
		`,
		`
		package p09

		func h[E any, T ~[]E](x T, l, h int) []E {
			s := x[l:h]
			print(s) //@ types("T")
			return s
		}
		func From() {
			type H []int32
			print(h[int32, H](nil, 0, 0)) //@ types("[]int32")
		}
		`,
		`
		package p10

		// Test "make" builtin with different forms on core types and
		// when capacities are constants or variable.
		func h[E any, T ~[]E](m, n int) {
			print(make(T, 3))    //@ types(T)
			print(make(T, 3, 5)) //@ types(T)
			print(make(T, m))    //@ types(T)
			print(make(T, m, n)) //@ types(T)
		}
		func i[K comparable, E any, T ~map[K]E](m int) {
			print(make(T))    //@ types(T)
			print(make(T, 5)) //@ types(T)
			print(make(T, m)) //@ types(T)
		}
		func j[E any, T ~chan E](m int) {
			print(make(T))    //@ types(T)
			print(make(T, 6)) //@ types(T)
			print(make(T, m)) //@ types(T)
		}
		func From() {
			type H []int32
			h[int32, H](3, 4)
			type I map[int8]H
			i[int8, H, I](5)
			type J chan I
			j[I, J](6)
		}
		`,
		`
		package p11

		func h[T ~[4]int](x T) {
			print(len(x), cap(x)) //@ types(int, int)
		}
		func i[T ~[4]byte | []int | ~chan uint8](x T) {
			print(len(x), cap(x)) //@ types(int, int)
		}
		func j[T ~[4]int | any | map[string]int]() {
			print(new(T)) //@ types("*T")
		}
		func k[T ~[4]int | any | map[string]int](x T) {
			print(x) //@ types(T)
			panic(x)
		}
		`,
		`
		package p12

		func f[E any, F ~func() E](x F) {
			print(x, x()) //@ types(F, E)
		}
		func From() {
			type T func() int
			f[int, T](func() int { return 0 })
			f[int, func() int](func() int { return 1 })
		}
		`,
		`
		package p13

		func f[E any, M ~map[string]E](m M) {
			y, ok := m["lorem"]
			print(m, y, ok) //@ types(M, E, bool)
		}
		func From() {
			type O map[string][]int
			f(O{"lorem": []int{0, 1, 2, 3}})
		}
		`,
		`
		package p14

		func a[T interface{ []int64 | [5]int64 }](x T) int64 {
			print(x, x[2], x[3]) //@ types(T, int64, int64)
			x[2] = 5
			return x[3]
		}
		func b[T interface{ []byte | string }](x T) byte {
			print(x, x[3]) //@ types(T, byte)
			return x[3]
		}
		func c[T interface{ []byte }](x T) byte {
			print(x, x[2], x[3]) //@ types(T, byte, byte)
			x[2] = 'b'
			return x[3]
		}
		func d[T interface{ map[int]int64 }](x T) int64 {
			print(x, x[2], x[3]) //@ types(T, int64, int64)
			x[2] = 43
			return x[3]
		}
		func e[T ~string](t T) {
			print(t, t[0]) //@ types(T, uint8)
		}
		func f[T ~string|[]byte](t T) {
			print(t, t[0]) //@ types(T, uint8)
		}
		func g[T []byte](t T) {
			print(t, t[0]) //@ types(T, byte)
		}
		func h[T ~[4]int|[]int](t T) {
			print(t, t[0]) //@ types(T, int)
		}
		func i[T ~[4]int|*[4]int|[]int](t T) {
			print(t, t[0]) //@ types(T, int)
		}
		func j[T ~[4]int|*[4]int|[]int](t T) {
			print(t, &t[0]) //@ types(T, "*int")
		}
		`,
		`
		package p15

		type MyInt int
		type Other int
		type MyInterface interface{ foo() }

		// ChangeType tests
		func ct0(x int) { v := MyInt(x);  print(x, v) /*@ types(int, "p15.MyInt")*/ }
		func ct1[T MyInt | Other, S int ](x S) { v := T(x);  print(x, v) /*@ types(S, T)*/ }
		func ct2[T int, S MyInt | int ](x S) { v := T(x); print(x, v) /*@ types(S, T)*/ }
		func ct3[T MyInt | Other, S MyInt | int ](x S) { v := T(x) ; print(x, v) /*@ types(S, T)*/ }

		// Convert tests
		func co0[T int | int8](x MyInt) { v := T(x); print(x, v) /*@ types("p15.MyInt", T)*/}
		func co1[T int | int8](x T) { v := MyInt(x); print(x, v) /*@ types(T, "p15.MyInt")*/ }
		func co2[S, T int | int8](x T) { v := S(x); print(x, v) /*@ types(T, S)*/ }

		// MakeInterface tests
		func mi0[T MyInterface](x T) { v := MyInterface(x); print(x, v) /*@ types(T, "p15.MyInterface")*/ }

		// NewConst tests
		func nc0[T any]() { v := (*T)(nil); print(v) /*@ types("*T")*/}

		// SliceToArrayPointer
		func sl0[T *[4]int | *[2]int](x []int) { v := T(x); print(x, v) /*@ types("[]int", T)*/ }
		func sl1[T *[4]int | *[2]int, S []int](x S) { v := T(x); print(x, v) /*@ types(S, T)*/ }
		`,
		`
		package p16

		func c[T interface{ foo() string }](x T) {
			print(x, x.foo, x.foo())  /*@ types(T, "func() string", string)*/
		}
		`,
		`
		package p17

		func eq[T comparable](t T, i interface{}) bool {
			return t == i
		}
		`,
		// TODO(59983): investigate why writing g.c panics in (*FieldAddr).String.
		`
		package p18

		type S struct{ f int }
		func c[P *S]() []P { return []P{{f: 1}} }
		`,
		`
		package p19

		func sign[bytes []byte | string](s bytes) (bool, bool) {
			neg := false
			if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
				neg = s[0] == '-'
				s = s[1:]
			}
			return !neg, len(s) > 0
		}
		`,
		`package p20

		func digits[bytes []byte | string](s bytes) bool {
			for _, c := range []byte(s) {
				if c < '0' || '9' < c {
					return false
				}
			}
			return true
		}
		`,
		`
		package p21

		type E interface{}

		func Foo[T E, PT interface{ *T }]() T {
			pt := PT(new(T))
			x := *pt
			print(x)  /*@ types(T)*/
			return x
		}
		`,
		`
		package p22

		func f[M any, PM *M](p PM) {
			var m M
			*p = m
			print(m)  /*@ types(M)*/
			print(p)  /*@ types(PM)*/
		}
		`,
		`
		package p23

		type A struct{int}
		func (*A) Marker() {}

		type B struct{string}
		func (*B) Marker() {}

		type C struct{float32}
		func (*C) Marker() {}

		func process[T interface {
			*A
			*B
			*C
			Marker()
		}](v T) {
			v.Marker()
			a := *(any(v).(*A)); print(a)  /*@ types("p23.A")*/
			b := *(any(v).(*B)); print(b)  /*@ types("p23.B")*/
			c := *(any(v).(*C)); print(c)  /*@ types("p23.C")*/
		}
		`,
		`
		package p24

		func a[T any](f func() [4]T) {
			x := len(f())
			print(x) /*@ types("int")*/
		}

		func b[T [4]any](f func() T) {
			x := len(f())
			print(x) /*@ types("int")*/
		}

		func c[T any](f func() *[4]T) {
			x := len(f())
			print(x) /*@ types("int")*/
		}

		func d[T *[4]any](f func() T) {
			x := len(f())
			print(x) /*@ types("int")*/
		}
		`,
		`
		package p25

		func a[T any]() {
			var f func() [4]T
			for i, v := range f() {
				print(i, v) /*@ types("int", "T")*/
			}
		}

		func b[T [4]any](f func() T) {
			for i, v := range f() {
				print(i, v) /*@ types("int", "any")*/
			}
		}

		func c[T any](f func() *[4]T) {
			for i, v := range f() {
				print(i, v) /*@ types("int", "T")*/
			}
		}

		func d[T *[4]any](f func() T) {
			for i, v := range f() {
				print(i, v) /*@ types("int", "any")*/
			}
		}
		`,
		`
		package issue64324

		type bar[T any] interface {
			Bar(int) T
		}
		type foo[T any] interface {
			bar[[]T]
			*T
		}
		func Foo[T any, F foo[T]](d int) {
			m := new(T)
			f := F(m)
			print(f.Bar(d)) /*@ types("[]T")*/
		}
		`, `
		package issue64324b

		type bar[T any] interface {
			Bar(int) T
		}
		type baz[T any] interface {
			bar[*int]
			*int
		}

		func Baz[I baz[string]](d int) {
			m := new(int)
			f := I(m)
			print(f.Bar(d)) /*@ types("*int")*/
		}
		`,
	} {
		pkgname := parsePackageClause(t, content)
		t.Run(pkgname, func(t *testing.T) {
			t.Parallel()
			ssapkg, ppkg := buildPackage(t, content, ssa.SanityCheckFunctions)
			fset := ssapkg.Prog.Fset

			// Collect all notes in f, i.e. comments starting with "//@ types".
			notes, err := expect.ExtractGo(fset, ppkg.Syntax[0])
			if err != nil {
				t.Errorf("expect.ExtractGo: %v", err)
			}

			// Collect calls to the builtin print function.
			fns := make(map[*ssa.Function]bool)
			for _, mem := range ssapkg.Members {
				if fn, ok := mem.(*ssa.Function); ok {
					fns[fn] = true
				}
			}
			probes := callsTo(fns, "print")
			expectations := matchNotes(fset, notes, probes)

			for call := range probes {
				if expectations[call] == nil {
					t.Errorf("Unmatched call: %v", call)
				}
			}

			// Check each expectation.
			for call, note := range expectations {
				var args []string
				for _, a := range call.Args {
					args = append(args, a.Type().String())
				}
				if got, want := fmt.Sprint(args), fmt.Sprint(note.Args); got != want {
					t.Errorf("Arguments to print() were expected to be %q. got %q", want, got)
					logFunction(t, probes[call])
				}
			}
		})
	}
}

// callsTo finds all calls to an SSA value named fname,
// and returns a map from each call site to its enclosing function.
func callsTo(fns map[*ssa.Function]bool, fname string) map[*ssa.CallCommon]*ssa.Function {
	callsites := make(map[*ssa.CallCommon]*ssa.Function)
	for fn := range fns {
		for _, bb := range fn.Blocks {
			for _, i := range bb.Instrs {
				if i, ok := i.(ssa.CallInstruction); ok {
					call := i.Common()
					if call.Value.Name() == fname {
						callsites[call] = fn
					}
				}
			}
		}
	}
	return callsites
}

// matchNotes returns a mapping from call sites (found by callsTo)
// to the first "//@ note" comment on the same line.
func matchNotes(fset *token.FileSet, notes []*expect.Note, calls map[*ssa.CallCommon]*ssa.Function) map[*ssa.CallCommon]*expect.Note {
	// Matches each probe with a note that has the same line.
	sameLine := func(x, y token.Pos) bool {
		xp := fset.Position(x)
		yp := fset.Position(y)
		return xp.Filename == yp.Filename && xp.Line == yp.Line
	}
	expectations := make(map[*ssa.CallCommon]*expect.Note)
	for call := range calls {
		for _, note := range notes {
			if sameLine(call.Pos(), note.Pos) {
				expectations[call] = note
				break // first match is good enough.
			}
		}
	}
	return expectations
}

// TestInstructionString tests serializing instructions via Instruction.String().
func TestInstructionString(t *testing.T) {
	// Tests (ssa.Instruction).String(). Instructions are from a single go file.
	// The Instructions tested are those that match a comment of the form:
	//
	//	//@ instrs(f, kind, strs...)
	//
	// where f is the name of the function, kind is the type of the instructions matched
	// within the function, and tests that the String() value for all of the instructions
	// matched of String() is strs (in some order).
	// See x/tools/go/expect for details on the syntax.

	const contents = `
	package p

	//@ instrs("f0", "*ssa.TypeAssert")
	//@ instrs("f0", "*ssa.Call", "print(nil:interface{}, 0:int)")
	func f0(x int) { // non-generic smoke test.
		var i interface{}
		print(i, 0)
	}

	//@ instrs("f1", "*ssa.Alloc", "local T (u)")
	//@ instrs("f1", "*ssa.FieldAddr", "&t0.x [#0]")
	func f1[T ~struct{ x string }]() T {
		u := T{"lorem"}
		return u
	}

	//@ instrs("f1b", "*ssa.Alloc", "new T (complit)")
	//@ instrs("f1b", "*ssa.FieldAddr", "&t0.x [#0]")
	func f1b[T ~struct{ x string }]() *T {
		u := &T{"lorem"}
		return u
	}

	//@ instrs("f2", "*ssa.TypeAssert", "typeassert t0.(interface{})")
	//@ instrs("f2", "*ssa.Call", "invoke x.foo()")
	func f2[T interface{ foo() string }](x T) {
		_ = x.foo
		_ = x.foo()
	}

	//@ instrs("f3", "*ssa.TypeAssert", "typeassert t0.(interface{})")
	//@ instrs("f3", "*ssa.Call", "invoke x.foo()")
	func f3[T interface{ foo() string; comparable }](x T) {
		_ = x.foo
		_ = x.foo()
	}

	//@ instrs("f4", "*ssa.BinOp", "t1 + 1:int", "t2 < 4:int")
	//@ instrs("f4", "*ssa.Call", "f()", "print(t2, t4)")
	func f4[T [4]string](f func() T) {
		for i, v := range f() {
			print(i, v)
		}
	}

	//@ instrs("f5", "*ssa.Call", "nil:func()()")
	func f5() {
		var f func()
		f()
	}

	type S struct{ f int }

	//@ instrs("f6", "*ssa.Alloc", "new [1]P (slicelit)", "new S (complit)")
	//@ instrs("f6", "*ssa.IndexAddr", "&t0[0:int]")
	//@ instrs("f6", "*ssa.FieldAddr", "&t2.f [#0]")
	func f6[P *S]() []P { return []P{{f: 1}} }

	//@ instrs("f7", "*ssa.Alloc", "local S (complit)")
	//@ instrs("f7", "*ssa.FieldAddr", "&t0.f [#0]")
	func f7[T any, S struct{f T}](x T) S { return S{f: x} }

	//@ instrs("f8", "*ssa.Alloc", "new [1]P (slicelit)", "new struct{f T} (complit)")
	//@ instrs("f8", "*ssa.IndexAddr", "&t0[0:int]")
	//@ instrs("f8", "*ssa.FieldAddr", "&t2.f [#0]")
	func f8[T any, P *struct{f T}](x T) []P { return []P{{f: x}} }

	//@ instrs("f9", "*ssa.Alloc", "new [1]PS (slicelit)", "new S (complit)")
	//@ instrs("f9", "*ssa.IndexAddr", "&t0[0:int]")
	//@ instrs("f9", "*ssa.FieldAddr", "&t2.f [#0]")
	func f9[T any, S struct{f T}, PS *S](x T) {
		_ = []PS{{f: x}}
	}

	//@ instrs("f10", "*ssa.FieldAddr", "&t0.x [#0]")
	//@ instrs("f10", "*ssa.Store", "*t0 = *new(T):T", "*t1 = 4:int")
	func f10[T ~struct{ x, y int }]() T {
		var u T
		u = T{x: 4}
		return u
	}

	//@ instrs("f11", "*ssa.FieldAddr", "&t1.y [#1]")
	//@ instrs("f11", "*ssa.Store", "*t1 = *new(T):T", "*t2 = 5:int")
	func f11[T ~struct{ x, y int }, PT *T]() PT {
		var u PT = new(T)
		*u = T{y: 5}
		return u
	}

	//@ instrs("f12", "*ssa.Alloc", "new struct{f T} (complit)")
	//@ instrs("f12", "*ssa.MakeMap", "make map[P]bool 1:int")
	func f12[T any, P *struct{f T}](x T) map[P]bool { return map[P]bool{{}: true} }

	//@ instrs("f13", "*ssa.IndexAddr", "&v[0:int]")
	//@ instrs("f13", "*ssa.Store", "*t0 = 7:int", "*v = *new(A):A")
	func f13[A [3]int, PA *A](v PA) {
		*v = A{7}
	}

	//@ instrs("f14", "*ssa.Call", "invoke t1.Set(0:int)")
	func f14[T any, PT interface {
		Set(int)
		*T
	}]() {
		var t T
		p := PT(&t)
		p.Set(0)
	}

	//@ instrs("f15", "*ssa.MakeClosure", "make closure (interface{Set(int); *T}).Set$bound [t1]")
	func f15[T any, PT interface {
		Set(int)
		*T
	}]() func(int) {
		var t T
		p := PT(&t)
		return p.Set
	}
	`

	// Parse
	conf := loader.Config{ParserMode: parser.ParseComments}
	const fname = "p.go"
	f, err := conf.ParseFile(fname, contents)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	conf.CreateFromFiles("p", f)

	// Load
	lprog, err := conf.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	// Create and build SSA
	prog := ssa.NewProgram(lprog.Fset, ssa.SanityCheckFunctions)
	for _, info := range lprog.AllPackages {
		if info.TransitivelyErrorFree {
			prog.CreatePackage(info.Pkg, info.Files, &info.Info, info.Importable)
		}
	}
	p := prog.Package(lprog.Package("p").Pkg)
	p.Build()

	// Collect all notes in f, i.e. comments starting with "//@ instr".
	notes, err := expect.ExtractGo(prog.Fset, f)
	if err != nil {
		t.Errorf("expect.ExtractGo: %v", err)
	}

	// Expectation is a {function, type string} -> {want, matches}
	// where matches is all Instructions.String() that match the key.
	// Each expecation is that some permutation of matches is wants.
	type expKey struct {
		function string
		kind     string
	}
	type expValue struct {
		wants   []string
		matches []string
	}
	expectations := make(map[expKey]*expValue)
	for _, note := range notes {
		if note.Name == "instrs" {
			if len(note.Args) < 2 {
				t.Error("Had @instrs annotation without at least 2 arguments")
				continue
			}
			fn, kind := fmt.Sprint(note.Args[0]), fmt.Sprint(note.Args[1])
			var wants []string
			for _, arg := range note.Args[2:] {
				wants = append(wants, fmt.Sprint(arg))
			}
			expectations[expKey{fn, kind}] = &expValue{wants, nil}
		}
	}

	// Collect all Instructions that match the expectations.
	for _, mem := range p.Members {
		if fn, ok := mem.(*ssa.Function); ok {
			for _, bb := range fn.Blocks {
				for _, i := range bb.Instrs {
					kind := fmt.Sprintf("%T", i)
					if e := expectations[expKey{fn.Name(), kind}]; e != nil {
						e.matches = append(e.matches, i.String())
					}
				}
			}
		}
	}

	// Check each expectation.
	for key, value := range expectations {
		fn, ok := p.Members[key.function].(*ssa.Function)
		if !ok {
			t.Errorf("Expectation on %s does not match a member in %s", key.function, p.Pkg.Name())
		}
		got, want := value.matches, value.wants
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Within %s wanted instructions of kind %s: %q. got %q", key.function, key.kind, want, got)
			logFunction(t, fn)
		}
	}
}

func logFunction(t testing.TB, fn *ssa.Function) {
	// TODO: Consider adding a ssa.Function.GoString() so this can be logged to t via '%#v'.
	var buf bytes.Buffer
	ssa.WriteFunction(&buf, fn)
	t.Log(buf.String())
}
```

## File: go/ssa/builder_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"

	"golang.org/x/sync/errgroup"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/expect"
	"golang.org/x/tools/internal/testenv"
)

func isEmpty(f *ssa.Function) bool { return f.Blocks == nil }

// Tests that programs partially loaded from gc object files contain
// functions with no code for the external portions, but are otherwise ok.
func TestBuildPackage(t *testing.T) {
	testenv.NeedsGoBuild(t) // for importer.Default()

	input := `
package main

import (
	"bytes"
	"io"
	"testing"
)

func main() {
        var t testing.T
	    t.Parallel()    // static call to external declared method
        t.Fail()        // static call to promoted external declared method
        testing.Short() // static call to external package-level function

        var w io.Writer = new(bytes.Buffer)
        w.Write(nil)    // interface invoke of external declared method
}
`

	// Parse the file.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "input.go", input, 0)
	if err != nil {
		t.Fatal(err)
		return
	}

	// Build an SSA program from the parsed file.
	// Load its dependencies from gc binary export data.
	mode := ssa.SanityCheckFunctions
	mainPkg, _, err := ssautil.BuildPackage(&types.Config{Importer: importer.Default()}, fset,
		types.NewPackage("main", ""), []*ast.File{f}, mode)
	if err != nil {
		t.Fatal(err)
		return
	}

	// The main package, its direct and indirect dependencies are loaded.
	deps := []string{
		// directly imported dependencies:
		"bytes", "io", "testing",
		// indirect dependencies mentioned by
		// the direct imports' export data
		"sync", "unicode", "time",
	}

	prog := mainPkg.Prog
	all := prog.AllPackages()
	if len(all) <= len(deps) {
		t.Errorf("unexpected set of loaded packages: %q", all)
	}
	for _, path := range deps {
		pkg := prog.ImportedPackage(path)
		if pkg == nil {
			t.Errorf("package not loaded: %q", path)
			continue
		}

		// External packages should have no function bodies (except for wrappers).
		isExt := pkg != mainPkg

		// init()
		if isExt && !isEmpty(pkg.Func("init")) {
			t.Errorf("external package %s has non-empty init", pkg)
		} else if !isExt && isEmpty(pkg.Func("init")) {
			t.Errorf("main package %s has empty init", pkg)
		}

		for _, mem := range pkg.Members {
			switch mem := mem.(type) {
			case *ssa.Function:
				// Functions at package level.
				if isExt && !isEmpty(mem) {
					t.Errorf("external function %s is non-empty", mem)
				} else if !isExt && isEmpty(mem) {
					t.Errorf("function %s is empty", mem)
				}

			case *ssa.Type:
				// Methods of named types T.
				// (In this test, all exported methods belong to *T not T.)
				if !isExt {
					t.Fatalf("unexpected name type in main package: %s", mem)
				}
				mset := prog.MethodSets.MethodSet(types.NewPointer(mem.Type()))
				for i, n := 0, mset.Len(); i < n; i++ {
					m := prog.MethodValue(mset.At(i))
					// For external types, only synthetic wrappers have code.
					expExt := !strings.Contains(m.Synthetic, "wrapper")
					if expExt && !isEmpty(m) {
						t.Errorf("external method %s is non-empty: %s",
							m, m.Synthetic)
					} else if !expExt && isEmpty(m) {
						t.Errorf("method function %s is empty: %s",
							m, m.Synthetic)
					}
				}
			}
		}
	}

	expectedCallee := []string{
		"(*testing.T).Parallel",
		"(*testing.common).Fail",
		"testing.Short",
		"N/A",
	}
	callNum := 0
	for _, b := range mainPkg.Func("main").Blocks {
		for _, instr := range b.Instrs {
			switch instr := instr.(type) {
			case ssa.CallInstruction:
				call := instr.Common()
				if want := expectedCallee[callNum]; want != "N/A" {
					got := call.StaticCallee().String()
					if want != got {
						t.Errorf("call #%d from main.main: got callee %s, want %s",
							callNum, got, want)
					}
				}
				callNum++
			}
		}
	}
	if callNum != 4 {
		t.Errorf("in main.main: got %d calls, want %d", callNum, 4)
	}
}

// Tests that methods from indirect dependencies not subject to
// CreatePackage are created as needed.
func TestNoIndirectCreatePackage(t *testing.T) {
	testenv.NeedsGoBuild(t) // for go/packages

	fs := openTxtar(t, filepath.Join(analysistest.TestData(), "indirect.txtar"))
	pkgs := loadPackages(t, fs, "testdata/a")
	a := pkgs[0]

	// Create a from syntax, its direct deps b from types, but not indirect deps c.
	prog := ssa.NewProgram(a.Fset, ssa.SanityCheckFunctions|ssa.PrintFunctions)
	aSSA := prog.CreatePackage(a.Types, a.Syntax, a.TypesInfo, false)
	for _, p := range a.Types.Imports() {
		prog.CreatePackage(p, nil, nil, true)
	}

	// Build SSA for package a.
	aSSA.Build()

	// Find the function in the sole call in the sole block of function a.A.
	var got string
	for _, instr := range aSSA.Members["A"].(*ssa.Function).Blocks[0].Instrs {
		if call, ok := instr.(*ssa.Call); ok {
			f := call.Call.Value.(*ssa.Function)
			got = fmt.Sprintf("%v # %s", f, f.Synthetic)
			break
		}
	}
	want := "(testdata/c.C).F # from type information (on demand)"
	if got != want {
		t.Errorf("for sole call in a.A, got: <<%s>>, want <<%s>>", got, want)
	}
}

// TestRuntimeTypes tests that (*Program).RuntimeTypes() includes all necessary types.
func TestRuntimeTypes(t *testing.T) {
	testenv.NeedsGoBuild(t) // for importer.Default()

	// TODO(adonovan): these test cases don't really make logical
	// sense any more. Rethink.

	tests := []struct {
		input string
		want  []string
	}{
		// A package-level type is needed.
		{`package A; type T struct{}; func (T) f() {}; var x any = T{}`,
			[]string{"*p.T", "p.T"},
		},
		// An unexported package-level type is not needed.
		{`package B; type t struct{}; func (t) f() {}`,
			nil,
		},
		// Subcomponents of type of exported package-level var are needed.
		{`package C; import "bytes"; var V struct {*bytes.Buffer}; var x any = &V`,
			[]string{"*bytes.Buffer", "*struct{*bytes.Buffer}", "struct{*bytes.Buffer}"},
		},
		// Subcomponents of type of unexported package-level var are not needed.
		{`package D; import "bytes"; var v struct {*bytes.Buffer}; var x any = v`,
			[]string{"*bytes.Buffer", "struct{*bytes.Buffer}"},
		},
		// Subcomponents of type of exported package-level function are needed.
		{`package E; import "bytes"; func F(struct {*bytes.Buffer}) {}; var v any = F`,
			[]string{"*bytes.Buffer", "struct{*bytes.Buffer}"},
		},
		// Subcomponents of type of unexported package-level function are not needed.
		{`package F; import "bytes"; func f(struct {*bytes.Buffer}) {}; var v any = f`,
			[]string{"*bytes.Buffer", "struct{*bytes.Buffer}"},
		},
		// Subcomponents of type of exported method of uninstantiated unexported type are not needed.
		{`package G; import "bytes"; type x struct{}; func (x) G(struct {*bytes.Buffer}) {}; var v x`,
			nil,
		},
		// ...unless used by MakeInterface.
		{`package G2; import "bytes"; type x struct{}; func (x) G(struct {*bytes.Buffer}) {}; var v interface{} = x{}`,
			[]string{"*bytes.Buffer", "*p.x", "p.x", "struct{*bytes.Buffer}"},
		},
		// Subcomponents of type of unexported method are not needed.
		{`package I; import "bytes"; type X struct{}; func (X) G(struct {*bytes.Buffer}) {}; var x any = X{}`,
			[]string{"*bytes.Buffer", "*p.X", "p.X", "struct{*bytes.Buffer}"},
		},
		// Local types aren't needed.
		{`package J; import "bytes"; func f() { type T struct {*bytes.Buffer}; var t T; _ = t }`,
			nil,
		},
		// ...unless used by MakeInterface.
		{`package K; import "bytes"; func f() { type T struct {*bytes.Buffer}; _ = interface{}(T{}) }`,
			[]string{"*bytes.Buffer", "*p.T", "p.T"},
		},
		// Types used as operand of MakeInterface are needed.
		{`package L; import "bytes"; func f() { _ = interface{}(struct{*bytes.Buffer}{}) }`,
			[]string{"*bytes.Buffer", "struct{*bytes.Buffer}"},
		},
		// MakeInterface is optimized away when storing to a blank.
		{`package M; import "bytes"; var _ interface{} = struct{*bytes.Buffer}{}`,
			nil,
		},
		// MakeInterface does not create runtime type for parameterized types.
		{`package N; var g interface{}; func f[S any]() { var v []S; g = v }; `,
			nil,
		},
	}
	for _, test := range tests {
		// Parse the file.
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "input.go", test.input, 0)
		if err != nil {
			t.Errorf("test %q: %s", test.input[:15], err)
			continue
		}

		// Create a single-file main package.
		// Load dependencies from gc binary export data.
		mode := ssa.SanityCheckFunctions
		ssapkg, _, err := ssautil.BuildPackage(&types.Config{Importer: importer.Default()}, fset,
			types.NewPackage("p", ""), []*ast.File{f}, mode)
		if err != nil {
			t.Errorf("test %q: %s", test.input[:15], err)
			continue
		}

		var typstrs []string
		for _, T := range ssapkg.Prog.RuntimeTypes() {
			if types.IsInterface(T) || types.NewMethodSet(T).Len() == 0 {
				continue // skip interfaces and types without methods
			}
			typstrs = append(typstrs, T.String())
		}
		sort.Strings(typstrs)

		if !reflect.DeepEqual(typstrs, test.want) {
			t.Errorf("test 'package %s': got %q, want %q",
				f.Name.Name, typstrs, test.want)
		}
	}
}

// TestInit tests that synthesized init functions are correctly formed.
// Bare init functions omit calls to dependent init functions and the use of
// an init guard. They are useful in cases where the client uses a different
// calling convention for init functions, or cases where it is easier for a
// client to analyze bare init functions. Both of these aspects are used by
// the llgo compiler for simpler integration with gccgo's runtime library,
// and to simplify the analysis whereby it deduces which stores to globals
// can be lowered to global initializers.
func TestInit(t *testing.T) {
	tests := []struct {
		mode        ssa.BuilderMode
		input, want string
	}{
		{0, `package A; import _ "errors"; var i int = 42`,
			`# Name: A.init
# Package: A
# Synthetic: package initializer
func init():
0:                                                                entry P:0 S:2
	t0 = *init$guard                                                   bool
	if t0 goto 2 else 1
1:                                                           init.start P:1 S:1
	*init$guard = true:bool
	t1 = errors.init()                                                   ()
	*i = 42:int
	jump 2
2:                                                            init.done P:2 S:0
	return

`},
		{ssa.BareInits, `package B; import _ "errors"; var i int = 42`,
			`# Name: B.init
# Package: B
# Synthetic: package initializer
func init():
0:                                                                entry P:0 S:0
	*i = 42:int
	return

`},
	}
	for _, test := range tests {
		// Create a single-file main package.
		mainPkg, _ := buildPackage(t, test.input, test.mode)
		name := mainPkg.Pkg.Name()
		initFunc := mainPkg.Func("init")
		if initFunc == nil {
			t.Errorf("test 'package %s': no init function", name)
			continue
		}

		var initbuf bytes.Buffer
		_, err := initFunc.WriteTo(&initbuf)
		if err != nil {
			t.Errorf("test 'package %s': WriteTo: %s", name, err)
			continue
		}

		if initbuf.String() != test.want {
			t.Errorf("test 'package %s': got %s, want %s", name, initbuf.String(), test.want)
		}
	}
}

// TestSyntheticFuncs checks that the expected synthetic functions are
// created, reachable, and not duplicated.
func TestSyntheticFuncs(t *testing.T) {
	const input = `package P
type T int
func (T) f() int
func (*T) g() int
var (
	// thunks
	a = T.f
	b = T.f
	c = (struct{T}).f
	d = (struct{T}).f
	e = (*T).g
	f = (*T).g
	g = (struct{*T}).g
	h = (struct{*T}).g

	// bounds
	i = T(0).f
	j = T(0).f
	k = new(T).g
	l = new(T).g

	// wrappers
	m interface{} = struct{T}{}
	n interface{} = struct{T}{}
	o interface{} = struct{*T}{}
	p interface{} = struct{*T}{}
	q interface{} = new(struct{T})
	r interface{} = new(struct{T})
	s interface{} = new(struct{*T})
	t interface{} = new(struct{*T})
)
`
	pkg, _ := buildPackage(t, input, ssa.BuilderMode(0))

	// Enumerate reachable synthetic functions
	want := map[string]string{
		"(*P.T).g$bound": "bound method wrapper for func (*P.T).g() int",
		"(P.T).f$bound":  "bound method wrapper for func (P.T).f() int",

		"(*P.T).g$thunk":         "thunk for func (*P.T).g() int",
		"(P.T).f$thunk":          "thunk for func (P.T).f() int",
		"(struct{*P.T}).g$thunk": "thunk for func (*P.T).g() int",
		"(struct{P.T}).f$thunk":  "thunk for func (P.T).f() int",

		"(*P.T).f":          "wrapper for func (P.T).f() int",
		"(*struct{*P.T}).f": "wrapper for func (P.T).f() int",
		"(*struct{*P.T}).g": "wrapper for func (*P.T).g() int",
		"(*struct{P.T}).f":  "wrapper for func (P.T).f() int",
		"(*struct{P.T}).g":  "wrapper for func (*P.T).g() int",
		"(struct{*P.T}).f":  "wrapper for func (P.T).f() int",
		"(struct{*P.T}).g":  "wrapper for func (*P.T).g() int",
		"(struct{P.T}).f":   "wrapper for func (P.T).f() int",

		"P.init": "package initializer",
	}
	var seen []string // may contain dups
	for fn := range ssautil.AllFunctions(pkg.Prog) {
		if fn.Synthetic == "" {
			continue
		}
		name := fn.String()
		wantDescr, ok := want[name]
		if !ok {
			t.Errorf("got unexpected/duplicate func: %q: %q", name, fn.Synthetic)
			continue
		}
		seen = append(seen, name)

		if wantDescr != fn.Synthetic {
			t.Errorf("(%s).Synthetic = %q, want %q", name, fn.Synthetic, wantDescr)
		}
	}

	for _, name := range seen {
		delete(want, name)
	}
	for fn, descr := range want {
		t.Errorf("want func: %q: %q", fn, descr)
	}
}

// TestPhiElimination ensures that dead phis, including those that
// participate in a cycle, are properly eliminated.
func TestPhiElimination(t *testing.T) {
	const input = `
package p

func f() error

func g(slice []int) {
	for {
		for range slice {
			// e should not be lifted to a dead φ-node.
			e := f()
			h(e)
		}
	}
}

func h(error)
`
	// The SSA code for this function should look something like this:
	// 0:
	//         jump 1
	// 1:
	//         t0 = len(slice)
	//         jump 2
	// 2:
	//         t1 = phi [1: -1:int, 3: t2]
	//         t2 = t1 + 1:int
	//         t3 = t2 < t0
	//         if t3 goto 3 else 1
	// 3:
	//         t4 = f()
	//         t5 = h(t4)
	//         jump 2
	//
	// But earlier versions of the SSA construction algorithm would
	// additionally generate this cycle of dead phis:
	//
	// 1:
	//         t7 = phi [0: nil:error, 2: t8] #e
	//         ...
	// 2:
	//         t8 = phi [1: t7, 3: t4] #e
	//         ...

	p, _ := buildPackage(t, input, ssa.BuilderMode(0))
	g := p.Func("g")

	phis := 0
	for _, b := range g.Blocks {
		for _, instr := range b.Instrs {
			if _, ok := instr.(*ssa.Phi); ok {
				phis++
			}
		}
	}
	if phis != 1 {
		g.WriteTo(os.Stderr)
		t.Errorf("expected a single Phi (for the range index), got %d", phis)
	}
}

// TestGenericDecls ensures that *unused* generic types, methods and functions
// signatures can be built.
//
// TODO(taking): Add calls from non-generic functions to instantiations of generic functions.
// TODO(taking): Add globals with types that are instantiations of generic functions.
func TestGenericDecls(t *testing.T) {
	const input = `
package p

import "unsafe"

type Pointer[T any] struct {
	v unsafe.Pointer
}

func (x *Pointer[T]) Load() *T {
	return (*T)(LoadPointer(&x.v))
}

func Load[T any](x *Pointer[T]) *T {
	return x.Load()
}

func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
`
	// The SSA members for this package should look something like this:
	//          func  LoadPointer func(addr *unsafe.Pointer) (val unsafe.Pointer)
	//      type  Pointer     struct{v unsafe.Pointer}
	//        method (*Pointer[T any]) Load() *T
	//      func  init        func()
	//      var   init$guard  bool

	p, _ := buildPackage(t, input, ssa.BuilderMode(0))

	if load := p.Func("Load"); load.Signature.TypeParams().Len() != 1 {
		t.Errorf("expected a single type param T for Load got %q", load.Signature)
	}
	if ptr := p.Type("Pointer"); ptr.Type().(*types.Named).TypeParams().Len() != 1 {
		t.Errorf("expected a single type param T for Pointer got %q", ptr.Type())
	}
}

func TestGenericWrappers(t *testing.T) {
	const input = `
package p

type S[T any] struct {
	t *T
}

func (x S[T]) M() T {
	return *(x.t)
}

var thunk = S[int].M

var g S[int]
var bound = g.M

type R[T any] struct{ S[T] }

var indirect = R[int].M
`
	// The relevant SSA members for this package should look something like this:
	// var   bound      func() int
	// var   thunk      func(S[int]) int
	// var   wrapper    func(R[int]) int

	for _, mode := range []ssa.BuilderMode{ssa.BuilderMode(0), ssa.InstantiateGenerics} {
		p, _ := buildPackage(t, input, mode)

		for _, entry := range []struct {
			name    string // name of the package variable
			typ     string // type of the package variable
			wrapper string // wrapper function to which the package variable is set
			callee  string // callee within the wrapper function
		}{
			{
				"bound",
				"*func() int",
				"(p.S[int]).M$bound",
				"(p.S[int]).M[int]",
			},
			{
				"thunk",
				"*func(p.S[int]) int",
				"(p.S[int]).M$thunk",
				"(p.S[int]).M[int]",
			},
			{
				"indirect",
				"*func(p.R[int]) int",
				"(p.R[int]).M$thunk",
				"(p.S[int]).M[int]",
			},
		} {
			entry := entry
			t.Run(entry.name, func(t *testing.T) {
				v := p.Var(entry.name)
				if v == nil {
					t.Fatalf("Did not find variable for %q in %s", entry.name, p.String())
				}
				if v.Type().String() != entry.typ {
					t.Errorf("Expected type for variable %s: %q. got %q", v, entry.typ, v.Type())
				}

				// Find the wrapper for v. This is stored exactly once in init.
				var wrapper *ssa.Function
				for _, bb := range p.Func("init").Blocks {
					for _, i := range bb.Instrs {
						if store, ok := i.(*ssa.Store); ok && v == store.Addr {
							switch val := store.Val.(type) {
							case *ssa.Function:
								wrapper = val
							case *ssa.MakeClosure:
								wrapper = val.Fn.(*ssa.Function)
							}
						}
					}
				}
				if wrapper == nil {
					t.Fatalf("failed to find wrapper function for %s", entry.name)
				}
				if wrapper.String() != entry.wrapper {
					t.Errorf("Expected wrapper function %q. got %q", wrapper, entry.wrapper)
				}

				// Find the callee within the wrapper. There should be exactly one call.
				var callee *ssa.Function
				for _, bb := range wrapper.Blocks {
					for _, i := range bb.Instrs {
						if call, ok := i.(*ssa.Call); ok {
							callee = call.Call.StaticCallee()
						}
					}
				}
				if callee == nil {
					t.Fatalf("failed to find callee within wrapper %s", wrapper)
				}
				if callee.String() != entry.callee {
					t.Errorf("Expected callee in wrapper %q is %q. got %q", v, entry.callee, callee)
				}
			})
		}
	}
}

// TestTypeparamTest builds SSA over compilable examples in $GOROOT/test/typeparam/*.go.

func TestTypeparamTest(t *testing.T) {
	testenv.NeedsGOROOTDir(t, "test")

	// Tests use a fake goroot to stub out standard libraries with declarations in
	// testdata/src. Decreases runtime from ~80s to ~1s.

	if runtime.GOARCH == "wasm" {
		// Consistent flakes on wasm (#64726, #69409, #69410).
		// Needs more investigation, but more likely a wasm issue
		// Disabling for now.
		t.Skip("Consistent flakes on wasm (e.g. https://go.dev/issues/64726)")
	}

	// located GOROOT based on the relative path of errors in $GOROOT/src/errors
	stdPkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedFiles,
	}, "errors")
	if err != nil {
		t.Fatalf("Failed to load errors package from std: %s", err)
	}
	goroot := filepath.Dir(filepath.Dir(filepath.Dir(stdPkgs[0].GoFiles[0])))
	dir := filepath.Join(goroot, "test", "typeparam")
	if _, err = os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		t.Skipf("test/typeparam doesn't exist under GOROOT %s", goroot)
	}

	// Collect all of the .go files in
	fsys := os.DirFS(dir)
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		t.Fatal(err)
	}

	// Each call to buildPackage calls package.Load, which invokes "go list",
	// and with over 300 subtests this can be very slow (minutes, or tens
	// on some platforms). So, we use an overlay to map each test file to a
	// distinct single-file package and load them all at once.
	overlay := map[string][]byte{
		"go.mod": goMod("example.com", -1),
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue // Consider standalone go files.
		}
		src, err := fs.ReadFile(fsys, entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		// Only build test files that can be compiled, or compiled and run.
		if !bytes.HasPrefix(src, []byte("// run")) && !bytes.HasPrefix(src, []byte("// compile")) {
			t.Logf("%s: not detected as a run test", entry.Name())
			continue
		}

		filename := fmt.Sprintf("%s/main.go", entry.Name())
		overlay[filename] = src
	}

	// load all packages inside the overlay so 'go list' will be triggered only once.
	pkgs := loadPackages(t, overlayFS(overlay), "./...")
	for _, p := range pkgs {
		originFilename := filepath.Base(filepath.Dir(p.GoFiles[0]))
		t.Run(originFilename, func(t *testing.T) {
			t.Parallel()
			prog, _ := ssautil.Packages([]*packages.Package{p}, ssa.SanityCheckFunctions|ssa.InstantiateGenerics)
			prog.Package(p.Types).Build()
		})
	}
}

// TestOrderOfOperations ensures order of operations are as intended.
func TestOrderOfOperations(t *testing.T) {
	// Testing for the order of operations within an expression is done
	// by collecting the sequence of direct function calls within a *Function.
	// Callees are all external functions so they cannot be safely re-ordered by ssa.
	const input = `
package p

func a() int
func b() int
func c() int

func slice(s []int) []int { return s[a():b()] }
func sliceMax(s []int) []int { return s[a():b():c()] }

`

	p, _ := buildPackage(t, input, ssa.BuilderMode(0))

	for _, item := range []struct {
		fn   string
		want string // sequence of calls within the function.
	}{
		{"sliceMax", "[a() b() c()]"},
		{"slice", "[a() b()]"},
	} {
		fn := p.Func(item.fn)
		want := item.want
		t.Run(item.fn, func(t *testing.T) {
			t.Parallel()

			var calls []string
			for _, b := range fn.Blocks {
				for _, instr := range b.Instrs {
					if call, ok := instr.(ssa.CallInstruction); ok {
						calls = append(calls, call.String())
					}
				}
			}
			if got := fmt.Sprint(calls); got != want {
				fn.WriteTo(os.Stderr)
				t.Errorf("Expected sequence of function calls in %s was %s. got %s", fn, want, got)
			}
		})
	}
}

// TestGenericFunctionSelector ensures generic functions from other packages can be selected.
func TestGenericFunctionSelector(t *testing.T) {
	fsys := overlayFS(map[string][]byte{
		"go.mod":  goMod("example.com", -1),
		"main.go": []byte(`package main; import "example.com/a"; func main() { a.F[int](); a.G[int,string](); a.H(0) }`),
		"a/a.go":  []byte(`package a; func F[T any](){}; func G[S, T any](){}; func H[T any](a T){} `),
	})

	for _, mode := range []ssa.BuilderMode{
		ssa.SanityCheckFunctions,
		ssa.SanityCheckFunctions | ssa.InstantiateGenerics,
	} {

		pkgs := loadPackages(t, fsys, "example.com") // package main
		if len(pkgs) != 1 {
			t.Fatalf("Expected 1 root package but got %d", len(pkgs))
		}
		prog, _ := ssautil.Packages(pkgs, mode)
		p := prog.Package(pkgs[0].Types)
		p.Build()

		if p.Pkg.Name() != "main" {
			t.Fatalf("Expected the second package is main but got %s", p.Pkg.Name())
		}
		p.Build()

		var callees []string // callees of the CallInstruction.String() in main().
		for _, b := range p.Func("main").Blocks {
			for _, i := range b.Instrs {
				if call, ok := i.(ssa.CallInstruction); ok {
					if callee := call.Common().StaticCallee(); call != nil {
						callees = append(callees, callee.String())
					} else {
						t.Errorf("CallInstruction without StaticCallee() %q", call)
					}
				}
			}
		}
		sort.Strings(callees) // ignore the order in the code.

		want := "[example.com/a.F[int] example.com/a.G[int string] example.com/a.H[int]]"
		if got := fmt.Sprint(callees); got != want {
			t.Errorf("Expected main() to contain calls %v. got %v", want, got)
		}
	}
}

func TestIssue58491(t *testing.T) {
	// Test that a local type reaches type param in instantiation.
	src := `
		package p

		func foo[T any](blocking func() (T, error)) error {
			type result struct {
				res T
				error // ensure the method set of result is non-empty
			}

			res := make(chan result, 1)
			go func() {
				var r result
				r.res, r.error = blocking()
				res <- r
			}()
			r := <-res
			err := r // require the rtype for result when instantiated
			return err
		}
		var Inst = foo[int]
	`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	files := []*ast.File{f}

	pkg := types.NewPackage("p", "")
	conf := &types.Config{}
	p, _, err := ssautil.BuildPackage(conf, fset, pkg, files, ssa.SanityCheckFunctions|ssa.InstantiateGenerics)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Find the local type result instantiated with int.
	var found bool
	for _, rt := range p.Prog.RuntimeTypes() {
		if n, ok := rt.(*types.Named); ok {
			if u, ok := n.Underlying().(*types.Struct); ok {
				found = true
				if got, want := n.String(), "p.result"; got != want {
					t.Errorf("Expected the name %s got: %s", want, got)
				}
				if got, want := u.String(), "struct{res int; error}"; got != want {
					t.Errorf("Expected the underlying type of %s to be %s. got %s", n, want, got)
				}
			}
		}
	}
	if !found {
		t.Error("Failed to find any Named to struct types")
	}
}

func TestIssue58491Rec(t *testing.T) {
	// Roughly the same as TestIssue58491 but with a recursive type.
	src := `
		package p

		func foo[T any]() error {
			type result struct {
				res T
				next *result
				error // ensure the method set of result is non-empty
			}

			r := &result{}
			err := r // require the rtype for result when instantiated
			return err
		}
		var Inst = foo[int]
	`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	files := []*ast.File{f}

	pkg := types.NewPackage("p", "")
	conf := &types.Config{}
	p, _, err := ssautil.BuildPackage(conf, fset, pkg, files, ssa.SanityCheckFunctions|ssa.InstantiateGenerics)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Find the local type result instantiated with int.
	var found bool
	for _, rt := range p.Prog.RuntimeTypes() {
		if n, ok := types.Unalias(rt).(*types.Named); ok {
			if u, ok := n.Underlying().(*types.Struct); ok {
				found = true
				if got, want := n.String(), "p.result"; got != want {
					t.Errorf("Expected the name %s got: %s", want, got)
				}
				if got, want := u.String(), "struct{res int; next *p.result; error}"; got != want {
					t.Errorf("Expected the underlying type of %s to be %s. got %s", n, want, got)
				}
			}
		}
	}
	if !found {
		t.Error("Failed to find any Named to struct types")
	}
}

// TestSyntax ensures that a function's Syntax is available.
func TestSyntax(t *testing.T) {
	const input = `package p

	type P int
	func (x *P) g() *P { return x }

	func F[T ~int]() *T {
		type S1 *T
		type S2 *T
		type S3 *T
		f1 := func() S1 {
			f2 := func() S2 {
				return S2(nil)
			}
			return S1(f2())
		}
		f3 := func() S3 {
			return S3(f1())
		}
		return (*T)(f3())
	}
	var g = F[int]
	var _ = F[P] // unreferenced => not instantiated
	`

	p, _ := buildPackage(t, input, ssa.InstantiateGenerics)
	prog := p.Prog

	// Collect syntax information for all of the functions.
	got := make(map[string]string)
	for fn := range ssautil.AllFunctions(prog) {
		if fn.Name() == "init" {
			continue
		}
		syntax := fn.Syntax()
		if got[fn.Name()] != "" {
			t.Error("dup")
		}
		got[fn.Name()] = fmt.Sprintf("%T : %s @ %d", syntax, fn.Signature, prog.Fset.Position(syntax.Pos()).Line)
	}

	want := map[string]string{
		"g":          "*ast.FuncDecl : func() *p.P @ 4",
		"F":          "*ast.FuncDecl : func[T ~int]() *T @ 6",
		"F$1":        "*ast.FuncLit : func() p.S1 @ 10",
		"F$1$1":      "*ast.FuncLit : func() p.S2 @ 11",
		"F$2":        "*ast.FuncLit : func() p.S3 @ 16",
		"F[int]":     "*ast.FuncDecl : func() *int @ 6",
		"F[int]$1":   "*ast.FuncLit : func() p.S1 @ 10",
		"F[int]$1$1": "*ast.FuncLit : func() p.S2 @ 11",
		"F[int]$2":   "*ast.FuncLit : func() p.S3 @ 16",
		// ...but no F[P] etc as they are unreferenced.
		// (NB: GlobalDebug mode would cause them to be referenced.)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected the functions with signature to be:\n\t%#v.\n Got:\n\t%#v", want, got)
	}
}

func TestGo117Builtins(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		importer types.Importer
	}{
		{"slice to array pointer", "package p; var s []byte; var _ = (*[4]byte)(s)", nil},
		{"unsafe slice", `package p; import "unsafe"; var _ = unsafe.Add(nil, 0)`, importer.Default()},
		{"unsafe add", `package p; import "unsafe"; var _ = unsafe.Slice((*int)(nil), 0)`, importer.Default()},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "p.go", tc.src, parser.ParseComments)
			if err != nil {
				t.Error(err)
			}
			files := []*ast.File{f}

			pkg := types.NewPackage("p", "")
			conf := &types.Config{Importer: tc.importer}
			if _, _, err := ssautil.BuildPackage(conf, fset, pkg, files, ssa.SanityCheckFunctions); err != nil {
				t.Error(err)
			}
		})
	}
}

// TestLabels just tests that anonymous labels are handled.
func TestLabels(t *testing.T) {
	tests := []string{
		`package main
		  func main() { _:println(1) }`,
		`package main
		  func main() { _:println(1); _:println(2)}`,
	}
	for _, test := range tests {
		buildPackage(t, test, ssa.BuilderMode(0))
	}
}

func TestFixedBugs(t *testing.T) {
	for _, name := range []string{
		"issue66783a",
		"issue66783b",
	} {

		t.Run(name, func(t *testing.T) {
			base := name + ".go"
			path := filepath.Join(analysistest.TestData(), "fixedbugs", base)
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				t.Fatal(err)
			}
			files := []*ast.File{f}
			pkg := types.NewPackage(name, name)
			mode := ssa.SanityCheckFunctions | ssa.InstantiateGenerics
			// mode |= ssa.PrintFunctions // debug mode
			if _, _, err := ssautil.BuildPackage(&types.Config{}, fset, pkg, files, mode); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestIssue67079(t *testing.T) {
	// This test reproduced a race in the SSA builder nearly 100% of the time.

	// Load the package.
	const src = `package p; type T int; func (T) f() {}; var _ = (*T).f`
	spkg, ppkg := buildPackage(t, src, ssa.BuilderMode(0))
	prog := spkg.Prog
	var g errgroup.Group

	// Access bodies of all functions.
	g.Go(func() error {
		for fn := range ssautil.AllFunctions(prog) {
			for _, b := range fn.Blocks {
				for _, instr := range b.Instrs {
					if call, ok := instr.(*ssa.Call); ok {
						call.Common().StaticCallee() // access call.Value
					}
				}
			}
		}
		return nil
	})

	// Force building of wrappers.
	g.Go(func() error {
		ptrT := types.NewPointer(ppkg.Types.Scope().Lookup("T").Type())
		ptrTf := types.NewMethodSet(ptrT).At(0) // (*T).f symbol
		prog.MethodValue(ptrTf)
		return nil
	})

	g.Wait() // ignore error
}

func TestGenericAliases(t *testing.T) {
	testenv.NeedsGo1Point(t, 23)

	if os.Getenv("GENERICALIASTEST_CHILD") == "1" {
		testGenericAliases(t)
		return
	}

	testenv.NeedsExec(t)
	testenv.NeedsTool(t, "go")

	cmd := exec.Command(os.Args[0], "-test.run=TestGenericAliases")
	cmd.Env = append(os.Environ(),
		"GENERICALIASTEST_CHILD=1",
		"GODEBUG=gotypesalias=1",
		"GOEXPERIMENT=aliastypeparams",
	)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		t.Logf("out=<<%s>>", out)
	}
	var exitcode int
	if err, ok := err.(*exec.ExitError); ok {
		exitcode = err.ExitCode()
	}
	const want = 0
	if exitcode != want {
		t.Errorf("exited %d, want %d", exitcode, want)
	}
}

func testGenericAliases(t *testing.T) {
	testenv.NeedsGoExperiment(t, "aliastypeparams")

	const source = `
package p

type A = uint8
type B[T any] = [4]T

var F = f[string]

func f[S any]() {
	// Two copies of f are made: p.f[S] and p.f[string]

	var v A // application of A that is declared outside of f without no type arguments
	print("p.f", "String", "p.A", v)
	print("p.f", "==", v, uint8(0))
	print("p.f[string]", "String", "p.A", v)
	print("p.f[string]", "==", v, uint8(0))


	var u B[S] // application of B that is declared outside declared outside of f with type arguments
	print("p.f", "String", "p.B[S]", u)
	print("p.f", "==", u, [4]S{})
	print("p.f[string]", "String", "p.B[string]", u)
	print("p.f[string]", "==", u, [4]string{})

	type C[T any] = struct{ s S; ap *B[T]} // declaration within f with type params
	var w C[int] // application of C with type arguments
	print("p.f", "String", "p.C[int]", w)
	print("p.f", "==", w, struct{ s S; ap *[4]int}{})
	print("p.f[string]", "String", "p.C[int]", w)
	print("p.f[string]", "==", w, struct{ s string; ap *[4]int}{})
}
`

	p, _ := buildPackage(t, source, ssa.InstantiateGenerics)

	probes := callsTo(ssautil.AllFunctions(p.Prog), "print")
	if got, want := len(probes), 3*4*2; got != want {
		t.Errorf("Found %v probes, expected %v", got, want)
	}

	const debug = false // enable to debug skips
	skipped := 0
	for probe, fn := range probes {
		// Each probe is of the form:
		// 		print("within", "test", head, tail)
		// The probe only matches within a function whose fn.String() is within.
		// This allows for different instantiations of fn to match different probes.
		// On a match, it applies the test named "test" to head::tail.
		if len(probe.Args) < 3 {
			t.Fatalf("probe %v did not have enough arguments", probe)
		}
		within, test, head, tail := constString(probe.Args[0]), probe.Args[1], probe.Args[2], probe.Args[3:]
		if within != fn.String() {
			skipped++
			if debug {
				t.Logf("Skipping %q within %q", within, fn.String())
			}
			continue // does not match function
		}

		switch test := constString(test); test {
		case "==": // All of the values are types.Identical.
			for _, v := range tail {
				if !types.Identical(head.Type(), v.Type()) {
					t.Errorf("Expected %v and %v to have identical types", head, v)
				}
			}
		case "String": // head is a string constant that all values in tail must match Type().String()
			want := constString(head)
			for _, v := range tail {
				if got := v.Type().String(); got != want {
					t.Errorf("%s: %v had the Type().String()=%q. expected %q", within, v, got, want)
				}
			}
		default:
			t.Errorf("%q is not a test subcommand", test)
		}
	}
	if want := 3 * 4; skipped != want {
		t.Errorf("Skipped %d probes, expected to skip %d", skipped, want)
	}
}

// constString returns the value of a string constant
// or "<not a constant string>" if the value is not a string constant.
func constString(v ssa.Value) string {
	if c, ok := v.(*ssa.Const); ok {
		str := c.Value.String()
		return strings.Trim(str, `"`)
	}
	return "<not a constant string>"
}

// TestMultipleGoversions tests that globals initialized to equivalent
// function literals are compiled based on the different GoVersion in each file.
func TestMultipleGoversions(t *testing.T) {
	var contents = map[string]string{
		"post.go": `
	//go:build go1.22
	package p

	var distinct = func(l []int) {
		for i := range l {
			print(&i)
		}
	}
	`,
		"pre.go": `
	package p

	var same = func(l []int) {
		for i := range l {
			print(&i)
		}
	}
	`,
	}

	fset := token.NewFileSet()
	var files []*ast.File
	for _, fname := range []string{"post.go", "pre.go"} {
		file, err := parser.ParseFile(fset, fname, contents[fname], 0)
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, file)
	}

	pkg := types.NewPackage("p", "")
	conf := &types.Config{Importer: nil, GoVersion: "go1.21"}
	p, _, err := ssautil.BuildPackage(conf, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		t.Fatal(err)
	}

	// Test that global is initialized to a function literal that was
	// compiled to have the expected for loop range variable lifetime for i.
	for _, test := range []struct {
		global *ssa.Global
		want   string // basic block to []*ssa.Alloc.
	}{
		{p.Var("same"), "map[entry:[new int (i)]]"},               // i is allocated in the entry block.
		{p.Var("distinct"), "map[rangeindex.body:[new int (i)]]"}, // i is allocated in the body block.
	} {
		// Find the function the test.name global is initialized to.
		var fn *ssa.Function
		for _, b := range p.Func("init").Blocks {
			for _, instr := range b.Instrs {
				if s, ok := instr.(*ssa.Store); ok && s.Addr == test.global {
					fn, _ = s.Val.(*ssa.Function)
				}
			}
		}
		if fn == nil {
			t.Fatalf("Failed to find *ssa.Function for initial value of global %s", test.global)
		}

		allocs := make(map[string][]string) // block comments -> []Alloc
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				if a, ok := instr.(*ssa.Alloc); ok {
					allocs[b.Comment] = append(allocs[b.Comment], a.String())
				}
			}
		}
		if got := fmt.Sprint(allocs); got != test.want {
			t.Errorf("[%s:=%s] expected the allocations to be in the basic blocks %q, got %q", test.global, fn, test.want, got)
		}
	}
}

// TestRangeOverInt tests that, in a range-over-int (#61405),
// the type of each range var v (identified by print(v) calls)
// has the expected type.
func TestRangeOverInt(t *testing.T) {
	const rangeOverIntSrc = `
		package p

		type I uint8

		func noKey(x int) {
			for range x {
				// does not crash
			}
		}

		func untypedConstantOperand() {
			for i := range 10 {
				print(i) /*@ types("int")*/
			}
		}

		func unsignedOperand(x uint64) {
			for i := range x {
				print(i) /*@ types("uint64")*/
			}
		}

		func namedOperand(x I) {
			for i := range x {
				print(i)  /*@ types("p.I")*/
			}
		}

		func typeparamOperand[T int](x T) {
			for i := range x {
				print(i)  /*@ types("T")*/
			}
		}

		func assignment(x I) {
			var k I
			for k = range x {
				print(k) /*@ types("p.I")*/
			}
		}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", rangeOverIntSrc, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	pkg := types.NewPackage("p", "")
	conf := &types.Config{}
	p, _, err := ssautil.BuildPackage(conf, fset, pkg, []*ast.File{f}, ssa.SanityCheckFunctions)
	if err != nil {
		t.Fatal(err)
	}

	// Collect all notes in f, i.e. comments starting with "//@ types".
	notes, err := expect.ExtractGo(fset, f)
	if err != nil {
		t.Fatal(err)
	}

	// Collect calls to the built-in print function.
	fns := make(map[*ssa.Function]bool)
	for _, mem := range p.Members {
		if fn, ok := mem.(*ssa.Function); ok {
			fns[fn] = true
		}
	}
	probes := callsTo(fns, "print")
	expectations := matchNotes(fset, notes, probes)

	for call := range probes {
		if expectations[call] == nil {
			t.Errorf("Unmatched call: %v @ %s", call, fset.Position(call.Pos()))
		}
	}

	// Check each expectation.
	for call, note := range expectations {
		var args []string
		for _, a := range call.Args {
			args = append(args, a.Type().String())
		}
		if got, want := fmt.Sprint(args), fmt.Sprint(note.Args); got != want {
			at := fset.Position(call.Pos())
			t.Errorf("%s: arguments to print had types %s, want %s", at, got, want)
			logFunction(t, probes[call])
		}
	}
}

func TestBuildPackageGo120(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		importer types.Importer
	}{
		{"slice to array", "package p; var s []byte; var _ = ([4]byte)(s)", nil},
		{"slice to zero length array", "package p; var s []byte; var _ = ([0]byte)(s)", nil},
		{"slice to zero length array type parameter", "package p; var s []byte; func f[T ~[0]byte]() { tmp := (T)(s); var z T; _ = tmp == z}", nil},
		{"slice to non-zero length array type parameter", "package p; var s []byte; func h[T ~[1]byte | [4]byte]() { tmp := T(s); var z T; _ = tmp == z}", nil},
		{"slice to maybe-zero length array type parameter", "package p; var s []byte; func g[T ~[0]byte | [4]byte]() { tmp := T(s); var z T; _ = tmp == z}", nil},
		{
			"rune sequence to sequence cast patterns", `
			package p
			// Each of fXX functions describes a 1.20 legal cast between sequences of runes
			// as []rune, pointers to rune arrays, rune arrays, or strings.
			//
			// Comments listed given the current emitted instructions [approximately].
			// If multiple conversions are needed, these are separated by |.
			// rune was selected as it leads to string casts (byte is similar).
			// The length 2 is not significant.
			// Multiple array lengths may occur in a cast in practice (including 0).
			func f00[S string, D string](s S)                               { _ = D(s) } // ChangeType
			func f01[S string, D []rune](s S)                               { _ = D(s) } // Convert
			func f02[S string, D []rune | string](s S)                      { _ = D(s) } // ChangeType | Convert
			func f03[S [2]rune, D [2]rune](s S)                             { _ = D(s) } // ChangeType
			func f04[S *[2]rune, D *[2]rune](s S)                           { _ = D(s) } // ChangeType
			func f05[S []rune, D string](s S)                               { _ = D(s) } // Convert
			func f06[S []rune, D [2]rune](s S)                              { _ = D(s) } // SliceToArrayPointer; Deref
			func f07[S []rune, D [2]rune | string](s S)                     { _ = D(s) } // SliceToArrayPointer; Deref | Convert
			func f08[S []rune, D *[2]rune](s S)                             { _ = D(s) } // SliceToArrayPointer
			func f09[S []rune, D *[2]rune | string](s S)                    { _ = D(s) } // SliceToArrayPointer; Deref | Convert
			func f10[S []rune, D *[2]rune | [2]rune](s S)                   { _ = D(s) } // SliceToArrayPointer | SliceToArrayPointer; Deref
			func f11[S []rune, D *[2]rune | [2]rune | string](s S)          { _ = D(s) } // SliceToArrayPointer | SliceToArrayPointer; Deref | Convert
			func f12[S []rune, D []rune](s S)                               { _ = D(s) } // ChangeType
			func f13[S []rune, D []rune | string](s S)                      { _ = D(s) } // Convert | ChangeType
			func f14[S []rune, D []rune | [2]rune](s S)                     { _ = D(s) } // ChangeType | SliceToArrayPointer; Deref
			func f15[S []rune, D []rune | [2]rune | string](s S)            { _ = D(s) } // ChangeType | SliceToArrayPointer; Deref | Convert
			func f16[S []rune, D []rune | *[2]rune](s S)                    { _ = D(s) } // ChangeType | SliceToArrayPointer
			func f17[S []rune, D []rune | *[2]rune | string](s S)           { _ = D(s) } // ChangeType | SliceToArrayPointer | Convert
			func f18[S []rune, D []rune | *[2]rune | [2]rune](s S)          { _ = D(s) } // ChangeType | SliceToArrayPointer | SliceToArrayPointer; Deref
			func f19[S []rune, D []rune | *[2]rune | [2]rune | string](s S) { _ = D(s) } // ChangeType | SliceToArrayPointer | SliceToArrayPointer; Deref | Convert
			func f20[S []rune | string, D string](s S)                      { _ = D(s) } // Convert | ChangeType
			func f21[S []rune | string, D []rune](s S)                      { _ = D(s) } // Convert | ChangeType
			func f22[S []rune | string, D []rune | string](s S)             { _ = D(s) } // ChangeType | Convert | Convert | ChangeType
			func f23[S []rune | [2]rune, D [2]rune](s S)                    { _ = D(s) } // SliceToArrayPointer; Deref | ChangeType
			func f24[S []rune | *[2]rune, D *[2]rune](s S)                  { _ = D(s) } // SliceToArrayPointer | ChangeType
			`, nil,
		},
		{
			"matching named and underlying types", `
			package p
			type a string
			type b string
			func g0[S []rune | a | b, D []rune | a | b](s S)      { _ = D(s) }
			func g1[S []rune | ~string, D []rune | a | b](s S)    { _ = D(s) }
			func g2[S []rune | a | b, D []rune | ~string](s S)    { _ = D(s) }
			func g3[S []rune | ~string, D []rune |~string](s S)   { _ = D(s) }
			`, nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "p.go", tc.src, 0)
			if err != nil {
				t.Error(err)
			}
			files := []*ast.File{f}

			pkg := types.NewPackage("p", "")
			conf := &types.Config{Importer: tc.importer}
			_, _, err = ssautil.BuildPackage(conf, fset, pkg, files, ssa.SanityCheckFunctions)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
```

## File: go/ssa/builder.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines the builder, which builds SSA-form IR for function bodies.
//
// SSA construction has two phases, "create" and "build". First, one
// or more packages are created in any order by a sequence of calls to
// CreatePackage, either from syntax or from mere type information.
// Each created package has a complete set of Members (const, var,
// type, func) that can be accessed through methods like
// Program.FuncValue.
//
// It is not necessary to call CreatePackage for all dependencies of
// each syntax package, only for its direct imports. (In future
// perhaps even this restriction may be lifted.)
//
// Second, packages created from syntax are built, by one or more
// calls to Package.Build, which may be concurrent; or by a call to
// Program.Build, which builds all packages in parallel. Building
// traverses the type-annotated syntax tree of each function body and
// creates SSA-form IR, a control-flow graph of instructions,
// populating fields such as Function.Body, .Params, and others.
//
// Building may create additional methods, including:
// - wrapper methods (e.g. for embeddding, or implicit &recv)
// - bound method closures (e.g. for use(recv.f))
// - thunks (e.g. for use(I.f) or use(T.f))
// - generic instances (e.g. to produce f[int] from f[any]).
// As these methods are created, they are added to the build queue,
// and then processed in turn, until a fixed point is reached,
// Since these methods might belong to packages that were not
// created (by a call to CreatePackage), their Pkg field is unset.
//
// Instances of generic functions may be either instantiated (f[int]
// is a copy of f[T] with substitutions) or wrapped (f[int] delegates
// to f[T]), depending on the availability of generic syntax and the
// InstantiateGenerics mode flag.
//
// Each package has an initializer function named "init" that calls
// the initializer functions of each direct import, computes and
// assigns the initial value of each global variable, and calls each
// source-level function named "init". (These generate SSA functions
// named "init#1", "init#2", etc.)
//
// Runtime types
//
// Each MakeInterface operation is a conversion from a non-interface
// type to an interface type. The semantics of this operation requires
// a runtime type descriptor, which is the type portion of an
// interface, and the value abstracted by reflect.Type.
//
// The program accumulates all non-parameterized types that are
// encountered as MakeInterface operands, along with all types that
// may be derived from them using reflection. This set is available as
// Program.RuntimeTypes, and the methods of these types may be
// reachable via interface calls or reflection even if they are never
// referenced from the SSA IR. (In practice, algorithms such as RTA
// that compute reachability from package main perform their own
// tracking of runtime types at a finer grain, so this feature is not
// very useful.)
//
// Function literals
//
// Anonymous functions must be built as soon as they are encountered,
// as it may affect locals of the enclosing function, but they are not
// marked 'built' until the end of the outermost enclosing function.
// (Among other things, this causes them to be logged in top-down order.)
//
// The Function.build fields determines the algorithm for building the
// function body. It is cleared to mark that building is complete.

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"os"
	"runtime"
	"sync"

	"slices"

	"golang.org/x/tools/internal/typeparams"
	"golang.org/x/tools/internal/versions"
)

type opaqueType struct{ name string }

func (t *opaqueType) String() string         { return t.name }
func (t *opaqueType) Underlying() types.Type { return t }

var (
	varOk    = newVar("ok", tBool)
	varIndex = newVar("index", tInt)

	// Type constants.
	tBool       = types.Typ[types.Bool]
	tByte       = types.Typ[types.Byte]
	tInt        = types.Typ[types.Int]
	tInvalid    = types.Typ[types.Invalid]
	tString     = types.Typ[types.String]
	tUntypedNil = types.Typ[types.UntypedNil]

	tRangeIter  = &opaqueType{"iter"}                         // the type of all "range" iterators
	tDeferStack = types.NewPointer(&opaqueType{"deferStack"}) // the type of a "deferStack" from ssa:deferstack()
	tEface      = types.NewInterfaceType(nil, nil).Complete()

	// SSA Value constants.
	vZero  = intConst(0)
	vOne   = intConst(1)
	vTrue  = NewConst(constant.MakeBool(true), tBool)
	vFalse = NewConst(constant.MakeBool(false), tBool)

	jReady = intConst(0)  // range-over-func jump is READY
	jBusy  = intConst(-1) // range-over-func jump is BUSY
	jDone  = intConst(-2) // range-over-func jump is DONE

	// The ssa:deferstack intrinsic returns the current function's defer stack.
	vDeferStack = &Builtin{
		name: "ssa:deferstack",
		sig:  types.NewSignatureType(nil, nil, nil, nil, types.NewTuple(anonVar(tDeferStack)), false),
	}
)

// builder holds state associated with the package currently being built.
// Its methods contain all the logic for AST-to-SSA conversion.
//
// All Functions belong to the same Program.
//
// builders are not thread-safe.
type builder struct {
	fns []*Function // Functions that have finished their CREATE phases.

	finished int // finished is the length of the prefix of fns containing built functions.

	// The task of building shared functions within the builder.
	// Shared functions are ones the the builder may either create or lookup.
	// These may be built by other builders in parallel.
	// The task is done when the builder has finished iterating, and it
	// waits for all shared functions to finish building.
	// nil implies there are no hared functions to wait on.
	buildshared *task
}

// shared is done when the builder has built all of the
// enqueued functions to a fixed-point.
func (b *builder) shared() *task {
	if b.buildshared == nil { // lazily-initialize
		b.buildshared = &task{done: make(chan unit)}
	}
	return b.buildshared
}

// enqueue fn to be built by the builder.
func (b *builder) enqueue(fn *Function) {
	b.fns = append(b.fns, fn)
}

// waitForSharedFunction indicates that the builder should wait until
// the potentially shared function fn has finished building.
//
// This should include any functions that may be built by other
// builders.
func (b *builder) waitForSharedFunction(fn *Function) {
	if fn.buildshared != nil { // maybe need to wait?
		s := b.shared()
		s.addEdge(fn.buildshared)
	}
}

// cond emits to fn code to evaluate boolean condition e and jump
// to t or f depending on its value, performing various simplifications.
//
// Postcondition: fn.currentBlock is nil.
func (b *builder) cond(fn *Function, e ast.Expr, t, f *BasicBlock) {
	switch e := e.(type) {
	case *ast.ParenExpr:
		b.cond(fn, e.X, t, f)
		return

	case *ast.BinaryExpr:
		switch e.Op {
		case token.LAND:
			ltrue := fn.newBasicBlock("cond.true")
			b.cond(fn, e.X, ltrue, f)
			fn.currentBlock = ltrue
			b.cond(fn, e.Y, t, f)
			return

		case token.LOR:
			lfalse := fn.newBasicBlock("cond.false")
			b.cond(fn, e.X, t, lfalse)
			fn.currentBlock = lfalse
			b.cond(fn, e.Y, t, f)
			return
		}

	case *ast.UnaryExpr:
		if e.Op == token.NOT {
			b.cond(fn, e.X, f, t)
			return
		}
	}

	// A traditional compiler would simplify "if false" (etc) here
	// but we do not, for better fidelity to the source code.
	//
	// The value of a constant condition may be platform-specific,
	// and may cause blocks that are reachable in some configuration
	// to be hidden from subsequent analyses such as bug-finding tools.
	emitIf(fn, b.expr(fn, e), t, f)
}

// logicalBinop emits code to fn to evaluate e, a &&- or
// ||-expression whose reified boolean value is wanted.
// The value is returned.
func (b *builder) logicalBinop(fn *Function, e *ast.BinaryExpr) Value {
	rhs := fn.newBasicBlock("binop.rhs")
	done := fn.newBasicBlock("binop.done")

	// T(e) = T(e.X) = T(e.Y) after untyped constants have been
	// eliminated.
	// TODO(adonovan): not true; MyBool==MyBool yields UntypedBool.
	t := fn.typeOf(e)

	var short Value // value of the short-circuit path
	switch e.Op {
	case token.LAND:
		b.cond(fn, e.X, rhs, done)
		short = NewConst(constant.MakeBool(false), t)

	case token.LOR:
		b.cond(fn, e.X, done, rhs)
		short = NewConst(constant.MakeBool(true), t)
	}

	// Is rhs unreachable?
	if rhs.Preds == nil {
		// Simplify false&&y to false, true||y to true.
		fn.currentBlock = done
		return short
	}

	// Is done unreachable?
	if done.Preds == nil {
		// Simplify true&&y (or false||y) to y.
		fn.currentBlock = rhs
		return b.expr(fn, e.Y)
	}

	// All edges from e.X to done carry the short-circuit value.
	var edges []Value
	for range done.Preds {
		edges = append(edges, short)
	}

	// The edge from e.Y to done carries the value of e.Y.
	fn.currentBlock = rhs
	edges = append(edges, b.expr(fn, e.Y))
	emitJump(fn, done)
	fn.currentBlock = done

	phi := &Phi{Edges: edges, Comment: e.Op.String()}
	phi.pos = e.OpPos
	phi.typ = t
	return done.emit(phi)
}

// exprN lowers a multi-result expression e to SSA form, emitting code
// to fn and returning a single Value whose type is a *types.Tuple.
// The caller must access the components via Extract.
//
// Multi-result expressions include CallExprs in a multi-value
// assignment or return statement, and "value,ok" uses of
// TypeAssertExpr, IndexExpr (when X is a map), and UnaryExpr (when Op
// is token.ARROW).
func (b *builder) exprN(fn *Function, e ast.Expr) Value {
	typ := fn.typeOf(e).(*types.Tuple)
	switch e := e.(type) {
	case *ast.ParenExpr:
		return b.exprN(fn, e.X)

	case *ast.CallExpr:
		// Currently, no built-in function nor type conversion
		// has multiple results, so we can avoid some of the
		// cases for single-valued CallExpr.
		var c Call
		b.setCall(fn, e, &c.Call)
		c.typ = typ
		return fn.emit(&c)

	case *ast.IndexExpr:
		mapt := typeparams.CoreType(fn.typeOf(e.X)).(*types.Map) // ,ok must be a map.
		lookup := &Lookup{
			X:       b.expr(fn, e.X),
			Index:   emitConv(fn, b.expr(fn, e.Index), mapt.Key()),
			CommaOk: true,
		}
		lookup.setType(typ)
		lookup.setPos(e.Lbrack)
		return fn.emit(lookup)

	case *ast.TypeAssertExpr:
		return emitTypeTest(fn, b.expr(fn, e.X), typ.At(0).Type(), e.Lparen)

	case *ast.UnaryExpr: // must be receive <-
		unop := &UnOp{
			Op:      token.ARROW,
			X:       b.expr(fn, e.X),
			CommaOk: true,
		}
		unop.setType(typ)
		unop.setPos(e.OpPos)
		return fn.emit(unop)
	}
	panic(fmt.Sprintf("exprN(%T) in %s", e, fn))
}

// builtin emits to fn SSA instructions to implement a call to the
// built-in function obj with the specified arguments
// and return type.  It returns the value defined by the result.
//
// The result is nil if no special handling was required; in this case
// the caller should treat this like an ordinary library function
// call.
func (b *builder) builtin(fn *Function, obj *types.Builtin, args []ast.Expr, typ types.Type, pos token.Pos) Value {
	typ = fn.typ(typ)
	switch obj.Name() {
	case "make":
		switch ct := typeparams.CoreType(typ).(type) {
		case *types.Slice:
			n := b.expr(fn, args[1])
			m := n
			if len(args) == 3 {
				m = b.expr(fn, args[2])
			}
			if m, ok := m.(*Const); ok {
				// treat make([]T, n, m) as new([m]T)[:n]
				cap := m.Int64()
				at := types.NewArray(ct.Elem(), cap)
				v := &Slice{
					X:    emitNew(fn, at, pos, "makeslice"),
					High: n,
				}
				v.setPos(pos)
				v.setType(typ)
				return fn.emit(v)
			}
			v := &MakeSlice{
				Len: n,
				Cap: m,
			}
			v.setPos(pos)
			v.setType(typ)
			return fn.emit(v)

		case *types.Map:
			var res Value
			if len(args) == 2 {
				res = b.expr(fn, args[1])
			}
			v := &MakeMap{Reserve: res}
			v.setPos(pos)
			v.setType(typ)
			return fn.emit(v)

		case *types.Chan:
			var sz Value = vZero
			if len(args) == 2 {
				sz = b.expr(fn, args[1])
			}
			v := &MakeChan{Size: sz}
			v.setPos(pos)
			v.setType(typ)
			return fn.emit(v)
		}

	case "new":
		return emitNew(fn, typeparams.MustDeref(typ), pos, "new")

	case "len", "cap":
		// Special case: len or cap of an array or *array is
		// based on the type, not the value which may be nil.
		// We must still evaluate the value, though.  (If it
		// was side-effect free, the whole call would have
		// been constant-folded.)
		t := typeparams.Deref(fn.typeOf(args[0]))
		if at, ok := typeparams.CoreType(t).(*types.Array); ok {
			b.expr(fn, args[0]) // for effects only
			return intConst(at.Len())
		}
		// Otherwise treat as normal.

	case "panic":
		fn.emit(&Panic{
			X:   emitConv(fn, b.expr(fn, args[0]), tEface),
			pos: pos,
		})
		fn.currentBlock = fn.newBasicBlock("unreachable")
		return vTrue // any non-nil Value will do
	}
	return nil // treat all others as a regular function call
}

// addr lowers a single-result addressable expression e to SSA form,
// emitting code to fn and returning the location (an lvalue) defined
// by the expression.
//
// If escaping is true, addr marks the base variable of the
// addressable expression e as being a potentially escaping pointer
// value.  For example, in this code:
//
//	a := A{
//	  b: [1]B{B{c: 1}}
//	}
//	return &a.b[0].c
//
// the application of & causes a.b[0].c to have its address taken,
// which means that ultimately the local variable a must be
// heap-allocated.  This is a simple but very conservative escape
// analysis.
//
// Operations forming potentially escaping pointers include:
// - &x, including when implicit in method call or composite literals.
// - a[:] iff a is an array (not *array)
// - references to variables in lexically enclosing functions.
func (b *builder) addr(fn *Function, e ast.Expr, escaping bool) lvalue {
	switch e := e.(type) {
	case *ast.Ident:
		if isBlankIdent(e) {
			return blank{}
		}
		obj := fn.objectOf(e).(*types.Var)
		var v Value
		if g := fn.Prog.packageLevelMember(obj); g != nil {
			v = g.(*Global) // var (address)
		} else {
			v = fn.lookup(obj, escaping)
		}
		return &address{addr: v, pos: e.Pos(), expr: e}

	case *ast.CompositeLit:
		typ := typeparams.Deref(fn.typeOf(e))
		var v *Alloc
		if escaping {
			v = emitNew(fn, typ, e.Lbrace, "complit")
		} else {
			v = emitLocal(fn, typ, e.Lbrace, "complit")
		}
		var sb storebuf
		b.compLit(fn, v, e, true, &sb)
		sb.emit(fn)
		return &address{addr: v, pos: e.Lbrace, expr: e}

	case *ast.ParenExpr:
		return b.addr(fn, e.X, escaping)

	case *ast.SelectorExpr:
		sel := fn.selection(e)
		if sel == nil {
			// qualified identifier
			return b.addr(fn, e.Sel, escaping)
		}
		if sel.kind != types.FieldVal {
			panic(sel)
		}
		wantAddr := true
		v := b.receiver(fn, e.X, wantAddr, escaping, sel)
		index := sel.index[len(sel.index)-1]
		fld := fieldOf(typeparams.MustDeref(v.Type()), index) // v is an addr.

		// Due to the two phases of resolving AssignStmt, a panic from x.f = p()
		// when x is nil is required to come after the side-effects of
		// evaluating x and p().
		emit := func(fn *Function) Value {
			return emitFieldSelection(fn, v, index, true, e.Sel)
		}
		return &lazyAddress{addr: emit, t: fld.Type(), pos: e.Sel.Pos(), expr: e.Sel}

	case *ast.IndexExpr:
		xt := fn.typeOf(e.X)
		elem, mode := indexType(xt)
		var x Value
		var et types.Type
		switch mode {
		case ixArrVar: // array, array|slice, array|*array, or array|*array|slice.
			x = b.addr(fn, e.X, escaping).address(fn)
			et = types.NewPointer(elem)
		case ixVar: // *array, slice, *array|slice
			x = b.expr(fn, e.X)
			et = types.NewPointer(elem)
		case ixMap:
			mt := typeparams.CoreType(xt).(*types.Map)
			return &element{
				m:   b.expr(fn, e.X),
				k:   emitConv(fn, b.expr(fn, e.Index), mt.Key()),
				t:   mt.Elem(),
				pos: e.Lbrack,
			}
		default:
			panic("unexpected container type in IndexExpr: " + xt.String())
		}
		index := b.expr(fn, e.Index)
		if isUntyped(index.Type()) {
			index = emitConv(fn, index, tInt)
		}
		// Due to the two phases of resolving AssignStmt, a panic from x[i] = p()
		// when x is nil or i is out-of-bounds is required to come after the
		// side-effects of evaluating x, i and p().
		emit := func(fn *Function) Value {
			v := &IndexAddr{
				X:     x,
				Index: index,
			}
			v.setPos(e.Lbrack)
			v.setType(et)
			return fn.emit(v)
		}
		return &lazyAddress{addr: emit, t: typeparams.MustDeref(et), pos: e.Lbrack, expr: e}

	case *ast.StarExpr:
		return &address{addr: b.expr(fn, e.X), pos: e.Star, expr: e}
	}

	panic(fmt.Sprintf("unexpected address expression: %T", e))
}

type store struct {
	lhs lvalue
	rhs Value
}

type storebuf struct{ stores []store }

func (sb *storebuf) store(lhs lvalue, rhs Value) {
	sb.stores = append(sb.stores, store{lhs, rhs})
}

func (sb *storebuf) emit(fn *Function) {
	for _, s := range sb.stores {
		s.lhs.store(fn, s.rhs)
	}
}

// assign emits to fn code to initialize the lvalue loc with the value
// of expression e.  If isZero is true, assign assumes that loc holds
// the zero value for its type.
//
// This is equivalent to loc.store(fn, b.expr(fn, e)), but may generate
// better code in some cases, e.g., for composite literals in an
// addressable location.
//
// If sb is not nil, assign generates code to evaluate expression e, but
// not to update loc.  Instead, the necessary stores are appended to the
// storebuf sb so that they can be executed later.  This allows correct
// in-place update of existing variables when the RHS is a composite
// literal that may reference parts of the LHS.
func (b *builder) assign(fn *Function, loc lvalue, e ast.Expr, isZero bool, sb *storebuf) {
	// Can we initialize it in place?
	if e, ok := ast.Unparen(e).(*ast.CompositeLit); ok {
		// A CompositeLit never evaluates to a pointer,
		// so if the type of the location is a pointer,
		// an &-operation is implied.
		if !is[blank](loc) && isPointerCore(loc.typ()) { // avoid calling blank.typ()
			ptr := b.addr(fn, e, true).address(fn)
			// copy address
			if sb != nil {
				sb.store(loc, ptr)
			} else {
				loc.store(fn, ptr)
			}
			return
		}

		if _, ok := loc.(*address); ok {
			if isNonTypeParamInterface(loc.typ()) {
				// e.g. var x interface{} = T{...}
				// Can't in-place initialize an interface value.
				// Fall back to copying.
			} else {
				// x = T{...} or x := T{...}
				addr := loc.address(fn)
				if sb != nil {
					b.compLit(fn, addr, e, isZero, sb)
				} else {
					var sb storebuf
					b.compLit(fn, addr, e, isZero, &sb)
					sb.emit(fn)
				}

				// Subtle: emit debug ref for aggregate types only;
				// slice and map are handled by store ops in compLit.
				switch typeparams.CoreType(loc.typ()).(type) {
				case *types.Struct, *types.Array:
					emitDebugRef(fn, e, addr, true)
				}

				return
			}
		}
	}

	// simple case: just copy
	rhs := b.expr(fn, e)
	if sb != nil {
		sb.store(loc, rhs)
	} else {
		loc.store(fn, rhs)
	}
}

// expr lowers a single-result expression e to SSA form, emitting code
// to fn and returning the Value defined by the expression.
func (b *builder) expr(fn *Function, e ast.Expr) Value {
	e = ast.Unparen(e)

	tv := fn.info.Types[e]

	// Is expression a constant?
	if tv.Value != nil {
		return NewConst(tv.Value, fn.typ(tv.Type))
	}

	var v Value
	if tv.Addressable() {
		// Prefer pointer arithmetic ({Index,Field}Addr) followed
		// by Load over subelement extraction (e.g. Index, Field),
		// to avoid large copies.
		v = b.addr(fn, e, false).load(fn)
	} else {
		v = b.expr0(fn, e, tv)
	}
	if fn.debugInfo() {
		emitDebugRef(fn, e, v, false)
	}
	return v
}

func (b *builder) expr0(fn *Function, e ast.Expr, tv types.TypeAndValue) Value {
	switch e := e.(type) {
	case *ast.BasicLit:
		panic("non-constant BasicLit") // unreachable

	case *ast.FuncLit:
		/* function literal */
		anon := &Function{
			name:           fmt.Sprintf("%s$%d", fn.Name(), 1+len(fn.AnonFuncs)),
			Signature:      fn.typeOf(e.Type).(*types.Signature),
			pos:            e.Type.Func,
			parent:         fn,
			anonIdx:        int32(len(fn.AnonFuncs)),
			Pkg:            fn.Pkg,
			Prog:           fn.Prog,
			syntax:         e,
			info:           fn.info,
			goversion:      fn.goversion,
			build:          (*builder).buildFromSyntax,
			topLevelOrigin: nil,           // use anonIdx to lookup an anon instance's origin.
			typeparams:     fn.typeparams, // share the parent's type parameters.
			typeargs:       fn.typeargs,   // share the parent's type arguments.
			subst:          fn.subst,      // share the parent's type substitutions.
			uniq:           fn.uniq,       // start from parent's unique values
		}
		fn.AnonFuncs = append(fn.AnonFuncs, anon)
		// Build anon immediately, as it may cause fn's locals to escape.
		// (It is not marked 'built' until the end of the enclosing FuncDecl.)
		anon.build(b, anon)
		fn.uniq = anon.uniq // resume after anon's unique values
		if anon.FreeVars == nil {
			return anon
		}
		v := &MakeClosure{Fn: anon}
		v.setType(fn.typ(tv.Type))
		for _, fv := range anon.FreeVars {
			v.Bindings = append(v.Bindings, fv.outer)
			fv.outer = nil
		}
		return fn.emit(v)

	case *ast.TypeAssertExpr: // single-result form only
		return emitTypeAssert(fn, b.expr(fn, e.X), fn.typ(tv.Type), e.Lparen)

	case *ast.CallExpr:
		if fn.info.Types[e.Fun].IsType() {
			// Explicit type conversion, e.g. string(x) or big.Int(x)
			x := b.expr(fn, e.Args[0])
			y := emitConv(fn, x, fn.typ(tv.Type))
			if y != x {
				switch y := y.(type) {
				case *Convert:
					y.pos = e.Lparen
				case *ChangeType:
					y.pos = e.Lparen
				case *MakeInterface:
					y.pos = e.Lparen
				case *SliceToArrayPointer:
					y.pos = e.Lparen
				case *UnOp: // conversion from slice to array.
					y.pos = e.Lparen
				}
			}
			return y
		}
		// Call to "intrinsic" built-ins, e.g. new, make, panic.
		if id, ok := ast.Unparen(e.Fun).(*ast.Ident); ok {
			if obj, ok := fn.info.Uses[id].(*types.Builtin); ok {
				if v := b.builtin(fn, obj, e.Args, fn.typ(tv.Type), e.Lparen); v != nil {
					return v
				}
			}
		}
		// Regular function call.
		var v Call
		b.setCall(fn, e, &v.Call)
		v.setType(fn.typ(tv.Type))
		return fn.emit(&v)

	case *ast.UnaryExpr:
		switch e.Op {
		case token.AND: // &X --- potentially escaping.
			addr := b.addr(fn, e.X, true)
			if _, ok := ast.Unparen(e.X).(*ast.StarExpr); ok {
				// &*p must panic if p is nil (http://golang.org/s/go12nil).
				// For simplicity, we'll just (suboptimally) rely
				// on the side effects of a load.
				// TODO(adonovan): emit dedicated nilcheck.
				addr.load(fn)
			}
			return addr.address(fn)
		case token.ADD:
			return b.expr(fn, e.X)
		case token.NOT, token.ARROW, token.SUB, token.XOR: // ! <- - ^
			v := &UnOp{
				Op: e.Op,
				X:  b.expr(fn, e.X),
			}
			v.setPos(e.OpPos)
			v.setType(fn.typ(tv.Type))
			return fn.emit(v)
		default:
			panic(e.Op)
		}

	case *ast.BinaryExpr:
		switch e.Op {
		case token.LAND, token.LOR:
			return b.logicalBinop(fn, e)
		case token.SHL, token.SHR:
			fallthrough
		case token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.AND, token.OR, token.XOR, token.AND_NOT:
			return emitArith(fn, e.Op, b.expr(fn, e.X), b.expr(fn, e.Y), fn.typ(tv.Type), e.OpPos)

		case token.EQL, token.NEQ, token.GTR, token.LSS, token.LEQ, token.GEQ:
			cmp := emitCompare(fn, e.Op, b.expr(fn, e.X), b.expr(fn, e.Y), e.OpPos)
			// The type of x==y may be UntypedBool.
			return emitConv(fn, cmp, types.Default(fn.typ(tv.Type)))
		default:
			panic("illegal op in BinaryExpr: " + e.Op.String())
		}

	case *ast.SliceExpr:
		var low, high, max Value
		var x Value
		xtyp := fn.typeOf(e.X)
		switch typeparams.CoreType(xtyp).(type) {
		case *types.Array:
			// Potentially escaping.
			x = b.addr(fn, e.X, true).address(fn)
		case *types.Basic, *types.Slice, *types.Pointer: // *array
			x = b.expr(fn, e.X)
		default:
			// core type exception?
			if isBytestring(xtyp) {
				x = b.expr(fn, e.X) // bytestring is handled as string and []byte.
			} else {
				panic("unexpected sequence type in SliceExpr")
			}
		}
		if e.Low != nil {
			low = b.expr(fn, e.Low)
		}
		if e.High != nil {
			high = b.expr(fn, e.High)
		}
		if e.Slice3 {
			max = b.expr(fn, e.Max)
		}
		v := &Slice{
			X:    x,
			Low:  low,
			High: high,
			Max:  max,
		}
		v.setPos(e.Lbrack)
		v.setType(fn.typ(tv.Type))
		return fn.emit(v)

	case *ast.Ident:
		obj := fn.info.Uses[e]
		// Universal built-in or nil?
		switch obj := obj.(type) {
		case *types.Builtin:
			return &Builtin{name: obj.Name(), sig: fn.instanceType(e).(*types.Signature)}
		case *types.Nil:
			return zeroConst(fn.instanceType(e))
		}

		// Package-level func or var?
		// (obj must belong to same package or a direct import.)
		if v := fn.Prog.packageLevelMember(obj); v != nil {
			if g, ok := v.(*Global); ok {
				return emitLoad(fn, g) // var (address)
			}
			callee := v.(*Function) // (func)
			if callee.typeparams.Len() > 0 {
				targs := fn.subst.types(instanceArgs(fn.info, e))
				callee = callee.instance(targs, b)
			}
			return callee
		}
		// Local var.
		return emitLoad(fn, fn.lookup(obj.(*types.Var), false)) // var (address)

	case *ast.SelectorExpr:
		sel := fn.selection(e)
		if sel == nil {
			// builtin unsafe.{Add,Slice}
			if obj, ok := fn.info.Uses[e.Sel].(*types.Builtin); ok {
				return &Builtin{name: obj.Name(), sig: fn.typ(tv.Type).(*types.Signature)}
			}
			// qualified identifier
			return b.expr(fn, e.Sel)
		}
		switch sel.kind {
		case types.MethodExpr:
			// (*T).f or T.f, the method f from the method-set of type T.
			// The result is a "thunk".
			thunk := createThunk(fn.Prog, sel)
			b.enqueue(thunk)
			return emitConv(fn, thunk, fn.typ(tv.Type))

		case types.MethodVal:
			// e.f where e is an expression and f is a method.
			// The result is a "bound".
			obj := sel.obj.(*types.Func)
			rt := fn.typ(recvType(obj))
			wantAddr := isPointer(rt)
			escaping := true
			v := b.receiver(fn, e.X, wantAddr, escaping, sel)

			if types.IsInterface(rt) {
				// If v may be an interface type I (after instantiating),
				// we must emit a check that v is non-nil.
				if recv, ok := types.Unalias(sel.recv).(*types.TypeParam); ok {
					// Emit a nil check if any possible instantiation of the
					// type parameter is an interface type.
					if !typeSetIsEmpty(recv) {
						// recv has a concrete term its typeset.
						// So it cannot be instantiated as an interface.
						//
						// Example:
						// func _[T interface{~int; Foo()}] () {
						//    var v T
						//    _ = v.Foo // <-- MethodVal
						// }
					} else {
						// rt may be instantiated as an interface.
						// Emit nil check: typeassert (any(v)).(any).
						emitTypeAssert(fn, emitConv(fn, v, tEface), tEface, token.NoPos)
					}
				} else {
					// non-type param interface
					// Emit nil check: typeassert v.(I).
					emitTypeAssert(fn, v, rt, e.Sel.Pos())
				}
			}
			if targs := receiverTypeArgs(obj); len(targs) > 0 {
				// obj is generic.
				obj = fn.Prog.canon.instantiateMethod(obj, fn.subst.types(targs), fn.Prog.ctxt)
			}
			bound := createBound(fn.Prog, obj)
			b.enqueue(bound)

			c := &MakeClosure{
				Fn:       bound,
				Bindings: []Value{v},
			}
			c.setPos(e.Sel.Pos())
			c.setType(fn.typ(tv.Type))
			return fn.emit(c)

		case types.FieldVal:
			indices := sel.index
			last := len(indices) - 1
			v := b.expr(fn, e.X)
			v = emitImplicitSelections(fn, v, indices[:last], e.Pos())
			v = emitFieldSelection(fn, v, indices[last], false, e.Sel)
			return v
		}

		panic("unexpected expression-relative selector")

	case *ast.IndexListExpr:
		// f[X, Y] must be a generic function
		if !instance(fn.info, e.X) {
			panic("unexpected expression-could not match index list to instantiation")
		}
		return b.expr(fn, e.X) // Handle instantiation within the *Ident or *SelectorExpr cases.

	case *ast.IndexExpr:
		if instance(fn.info, e.X) {
			return b.expr(fn, e.X) // Handle instantiation within the *Ident or *SelectorExpr cases.
		}
		// not a generic instantiation.
		xt := fn.typeOf(e.X)
		switch et, mode := indexType(xt); mode {
		case ixVar:
			// Addressable slice/array; use IndexAddr and Load.
			return b.addr(fn, e, false).load(fn)

		case ixArrVar, ixValue:
			// An array in a register, a string or a combined type that contains
			// either an [_]array (ixArrVar) or string (ixValue).

			// Note: for ixArrVar and CoreType(xt)==nil can be IndexAddr and Load.
			index := b.expr(fn, e.Index)
			if isUntyped(index.Type()) {
				index = emitConv(fn, index, tInt)
			}
			v := &Index{
				X:     b.expr(fn, e.X),
				Index: index,
			}
			v.setPos(e.Lbrack)
			v.setType(et)
			return fn.emit(v)

		case ixMap:
			ct := typeparams.CoreType(xt).(*types.Map)
			v := &Lookup{
				X:     b.expr(fn, e.X),
				Index: emitConv(fn, b.expr(fn, e.Index), ct.Key()),
			}
			v.setPos(e.Lbrack)
			v.setType(ct.Elem())
			return fn.emit(v)
		default:
			panic("unexpected container type in IndexExpr: " + xt.String())
		}

	case *ast.CompositeLit, *ast.StarExpr:
		// Addressable types (lvalues)
		return b.addr(fn, e, false).load(fn)
	}

	panic(fmt.Sprintf("unexpected expr: %T", e))
}

// stmtList emits to fn code for all statements in list.
func (b *builder) stmtList(fn *Function, list []ast.Stmt) {
	for _, s := range list {
		b.stmt(fn, s)
	}
}

// receiver emits to fn code for expression e in the "receiver"
// position of selection e.f (where f may be a field or a method) and
// returns the effective receiver after applying the implicit field
// selections of sel.
//
// wantAddr requests that the result is an address.  If
// !sel.indirect, this may require that e be built in addr() mode; it
// must thus be addressable.
//
// escaping is defined as per builder.addr().
func (b *builder) receiver(fn *Function, e ast.Expr, wantAddr, escaping bool, sel *selection) Value {
	var v Value
	if wantAddr && !sel.indirect && !isPointerCore(fn.typeOf(e)) {
		v = b.addr(fn, e, escaping).address(fn)
	} else {
		v = b.expr(fn, e)
	}

	last := len(sel.index) - 1
	// The position of implicit selection is the position of the inducing receiver expression.
	v = emitImplicitSelections(fn, v, sel.index[:last], e.Pos())
	if types.IsInterface(v.Type()) {
		// When v is an interface, sel.Kind()==MethodValue and v.f is invoked.
		// So v is not loaded, even if v has a pointer core type.
	} else if !wantAddr && isPointerCore(v.Type()) {
		v = emitLoad(fn, v)
	}
	return v
}

// setCallFunc populates the function parts of a CallCommon structure
// (Func, Method, Recv, Args[0]) based on the kind of invocation
// occurring in e.
func (b *builder) setCallFunc(fn *Function, e *ast.CallExpr, c *CallCommon) {
	c.pos = e.Lparen

	// Is this a method call?
	if selector, ok := ast.Unparen(e.Fun).(*ast.SelectorExpr); ok {
		sel := fn.selection(selector)
		if sel != nil && sel.kind == types.MethodVal {
			obj := sel.obj.(*types.Func)
			recv := recvType(obj)

			wantAddr := isPointer(recv)
			escaping := true
			v := b.receiver(fn, selector.X, wantAddr, escaping, sel)
			if types.IsInterface(recv) {
				// Invoke-mode call.
				c.Value = v // possibly type param
				c.Method = obj
			} else {
				// "Call"-mode call.
				c.Value = fn.Prog.objectMethod(obj, b)
				c.Args = append(c.Args, v)
			}
			return
		}

		// sel.kind==MethodExpr indicates T.f() or (*T).f():
		// a statically dispatched call to the method f in the
		// method-set of T or *T.  T may be an interface.
		//
		// e.Fun would evaluate to a concrete method, interface
		// wrapper function, or promotion wrapper.
		//
		// For now, we evaluate it in the usual way.
		//
		// TODO(adonovan): opt: inline expr() here, to make the
		// call static and to avoid generation of wrappers.
		// It's somewhat tricky as it may consume the first
		// actual parameter if the call is "invoke" mode.
		//
		// Examples:
		//  type T struct{}; func (T) f() {}   // "call" mode
		//  type T interface { f() }           // "invoke" mode
		//
		//  type S struct{ T }
		//
		//  var s S
		//  S.f(s)
		//  (*S).f(&s)
		//
		// Suggested approach:
		// - consume the first actual parameter expression
		//   and build it with b.expr().
		// - apply implicit field selections.
		// - use MethodVal logic to populate fields of c.
	}

	// Evaluate the function operand in the usual way.
	c.Value = b.expr(fn, e.Fun)
}

// emitCallArgs emits to f code for the actual parameters of call e to
// a (possibly built-in) function of effective type sig.
// The argument values are appended to args, which is then returned.
func (b *builder) emitCallArgs(fn *Function, sig *types.Signature, e *ast.CallExpr, args []Value) []Value {
	// f(x, y, z...): pass slice z straight through.
	if e.Ellipsis != 0 {
		for i, arg := range e.Args {
			v := emitConv(fn, b.expr(fn, arg), sig.Params().At(i).Type())
			args = append(args, v)
		}
		return args
	}

	offset := len(args) // 1 if call has receiver, 0 otherwise

	// Evaluate actual parameter expressions.
	//
	// If this is a chained call of the form f(g()) where g has
	// multiple return values (MRV), they are flattened out into
	// args; a suffix of them may end up in a varargs slice.
	for _, arg := range e.Args {
		v := b.expr(fn, arg)
		if ttuple, ok := v.Type().(*types.Tuple); ok { // MRV chain
			for i, n := 0, ttuple.Len(); i < n; i++ {
				args = append(args, emitExtract(fn, v, i))
			}
		} else {
			args = append(args, v)
		}
	}

	// Actual->formal assignability conversions for normal parameters.
	np := sig.Params().Len() // number of normal parameters
	if sig.Variadic() {
		np--
	}
	for i := 0; i < np; i++ {
		args[offset+i] = emitConv(fn, args[offset+i], sig.Params().At(i).Type())
	}

	// Actual->formal assignability conversions for variadic parameter,
	// and construction of slice.
	if sig.Variadic() {
		varargs := args[offset+np:]
		st := sig.Params().At(np).Type().(*types.Slice)
		vt := st.Elem()
		if len(varargs) == 0 {
			args = append(args, zeroConst(st))
		} else {
			// Replace a suffix of args with a slice containing it.
			at := types.NewArray(vt, int64(len(varargs)))
			a := emitNew(fn, at, token.NoPos, "varargs")
			a.setPos(e.Rparen)
			for i, arg := range varargs {
				iaddr := &IndexAddr{
					X:     a,
					Index: intConst(int64(i)),
				}
				iaddr.setType(types.NewPointer(vt))
				fn.emit(iaddr)
				emitStore(fn, iaddr, arg, arg.Pos())
			}
			s := &Slice{X: a}
			s.setType(st)
			args[offset+np] = fn.emit(s)
			args = args[:offset+np+1]
		}
	}
	return args
}

// setCall emits to fn code to evaluate all the parameters of a function
// call e, and populates *c with those values.
func (b *builder) setCall(fn *Function, e *ast.CallExpr, c *CallCommon) {
	// First deal with the f(...) part and optional receiver.
	b.setCallFunc(fn, e, c)

	// Then append the other actual parameters.
	sig, _ := typeparams.CoreType(fn.typeOf(e.Fun)).(*types.Signature)
	if sig == nil {
		panic(fmt.Sprintf("no signature for call of %s", e.Fun))
	}
	c.Args = b.emitCallArgs(fn, sig, e, c.Args)
}

// assignOp emits to fn code to perform loc <op>= val.
func (b *builder) assignOp(fn *Function, loc lvalue, val Value, op token.Token, pos token.Pos) {
	loc.store(fn, emitArith(fn, op, loc.load(fn), val, loc.typ(), pos))
}

// localValueSpec emits to fn code to define all of the vars in the
// function-local ValueSpec, spec.
func (b *builder) localValueSpec(fn *Function, spec *ast.ValueSpec) {
	switch {
	case len(spec.Values) == len(spec.Names):
		// e.g. var x, y = 0, 1
		// 1:1 assignment
		for i, id := range spec.Names {
			if !isBlankIdent(id) {
				emitLocalVar(fn, identVar(fn, id))
			}
			lval := b.addr(fn, id, false) // non-escaping
			b.assign(fn, lval, spec.Values[i], true, nil)
		}

	case len(spec.Values) == 0:
		// e.g. var x, y int
		// Locals are implicitly zero-initialized.
		for _, id := range spec.Names {
			if !isBlankIdent(id) {
				lhs := emitLocalVar(fn, identVar(fn, id))
				if fn.debugInfo() {
					emitDebugRef(fn, id, lhs, true)
				}
			}
		}

	default:
		// e.g. var x, y = pos()
		tuple := b.exprN(fn, spec.Values[0])
		for i, id := range spec.Names {
			if !isBlankIdent(id) {
				emitLocalVar(fn, identVar(fn, id))
				lhs := b.addr(fn, id, false) // non-escaping
				lhs.store(fn, emitExtract(fn, tuple, i))
			}
		}
	}
}

// assignStmt emits code to fn for a parallel assignment of rhss to lhss.
// isDef is true if this is a short variable declaration (:=).
//
// Note the similarity with localValueSpec.
func (b *builder) assignStmt(fn *Function, lhss, rhss []ast.Expr, isDef bool) {
	// Side effects of all LHSs and RHSs must occur in left-to-right order.
	lvals := make([]lvalue, len(lhss))
	isZero := make([]bool, len(lhss))
	for i, lhs := range lhss {
		var lval lvalue = blank{}
		if !isBlankIdent(lhs) {
			if isDef {
				if obj, ok := fn.info.Defs[lhs.(*ast.Ident)].(*types.Var); ok {
					emitLocalVar(fn, obj)
					isZero[i] = true
				}
			}
			lval = b.addr(fn, lhs, false) // non-escaping
		}
		lvals[i] = lval
	}
	if len(lhss) == len(rhss) {
		// Simple assignment:   x     = f()        (!isDef)
		// Parallel assignment: x, y  = f(), g()   (!isDef)
		// or short var decl:   x, y := f(), g()   (isDef)
		//
		// In all cases, the RHSs may refer to the LHSs,
		// so we need a storebuf.
		var sb storebuf
		for i := range rhss {
			b.assign(fn, lvals[i], rhss[i], isZero[i], &sb)
		}
		sb.emit(fn)
	} else {
		// e.g. x, y = pos()
		tuple := b.exprN(fn, rhss[0])
		emitDebugRef(fn, rhss[0], tuple, false)
		for i, lval := range lvals {
			lval.store(fn, emitExtract(fn, tuple, i))
		}
	}
}

// arrayLen returns the length of the array whose composite literal elements are elts.
func (b *builder) arrayLen(fn *Function, elts []ast.Expr) int64 {
	var max int64 = -1
	var i int64 = -1
	for _, e := range elts {
		if kv, ok := e.(*ast.KeyValueExpr); ok {
			i = b.expr(fn, kv.Key).(*Const).Int64()
		} else {
			i++
		}
		if i > max {
			max = i
		}
	}
	return max + 1
}

// compLit emits to fn code to initialize a composite literal e at
// address addr with type typ.
//
// Nested composite literals are recursively initialized in place
// where possible. If isZero is true, compLit assumes that addr
// holds the zero value for typ.
//
// Because the elements of a composite literal may refer to the
// variables being updated, as in the second line below,
//
//	x := T{a: 1}
//	x = T{a: x.a}
//
// all the reads must occur before all the writes.  Thus all stores to
// loc are emitted to the storebuf sb for later execution.
//
// A CompositeLit may have pointer type only in the recursive (nested)
// case when the type name is implicit.  e.g. in []*T{{}}, the inner
// literal has type *T behaves like &T{}.
// In that case, addr must hold a T, not a *T.
func (b *builder) compLit(fn *Function, addr Value, e *ast.CompositeLit, isZero bool, sb *storebuf) {
	typ := typeparams.Deref(fn.typeOf(e)) // retain the named/alias/param type, if any
	switch t := typeparams.CoreType(typ).(type) {
	case *types.Struct:
		if !isZero && len(e.Elts) != t.NumFields() {
			// memclear
			zt := typeparams.MustDeref(addr.Type())
			sb.store(&address{addr, e.Lbrace, nil}, zeroConst(zt))
			isZero = true
		}
		for i, e := range e.Elts {
			fieldIndex := i
			pos := e.Pos()
			if kv, ok := e.(*ast.KeyValueExpr); ok {
				fname := kv.Key.(*ast.Ident).Name
				for i, n := 0, t.NumFields(); i < n; i++ {
					sf := t.Field(i)
					if sf.Name() == fname {
						fieldIndex = i
						pos = kv.Colon
						e = kv.Value
						break
					}
				}
			}
			sf := t.Field(fieldIndex)
			faddr := &FieldAddr{
				X:     addr,
				Field: fieldIndex,
			}
			faddr.setPos(pos)
			faddr.setType(types.NewPointer(sf.Type()))
			fn.emit(faddr)
			b.assign(fn, &address{addr: faddr, pos: pos, expr: e}, e, isZero, sb)
		}

	case *types.Array, *types.Slice:
		var at *types.Array
		var array Value
		switch t := t.(type) {
		case *types.Slice:
			at = types.NewArray(t.Elem(), b.arrayLen(fn, e.Elts))
			array = emitNew(fn, at, e.Lbrace, "slicelit")
		case *types.Array:
			at = t
			array = addr

			if !isZero && int64(len(e.Elts)) != at.Len() {
				// memclear
				zt := typeparams.MustDeref(array.Type())
				sb.store(&address{array, e.Lbrace, nil}, zeroConst(zt))
			}
		}

		var idx *Const
		for _, e := range e.Elts {
			pos := e.Pos()
			if kv, ok := e.(*ast.KeyValueExpr); ok {
				idx = b.expr(fn, kv.Key).(*Const)
				pos = kv.Colon
				e = kv.Value
			} else {
				var idxval int64
				if idx != nil {
					idxval = idx.Int64() + 1
				}
				idx = intConst(idxval)
			}
			iaddr := &IndexAddr{
				X:     array,
				Index: idx,
			}
			iaddr.setType(types.NewPointer(at.Elem()))
			fn.emit(iaddr)
			if t != at { // slice
				// backing array is unaliased => storebuf not needed.
				b.assign(fn, &address{addr: iaddr, pos: pos, expr: e}, e, true, nil)
			} else {
				b.assign(fn, &address{addr: iaddr, pos: pos, expr: e}, e, true, sb)
			}
		}

		if t != at { // slice
			s := &Slice{X: array}
			s.setPos(e.Lbrace)
			s.setType(typ)
			sb.store(&address{addr: addr, pos: e.Lbrace, expr: e}, fn.emit(s))
		}

	case *types.Map:
		m := &MakeMap{Reserve: intConst(int64(len(e.Elts)))}
		m.setPos(e.Lbrace)
		m.setType(typ)
		fn.emit(m)
		for _, e := range e.Elts {
			e := e.(*ast.KeyValueExpr)

			// If a key expression in a map literal is itself a
			// composite literal, the type may be omitted.
			// For example:
			//	map[*struct{}]bool{{}: true}
			// An &-operation may be implied:
			//	map[*struct{}]bool{&struct{}{}: true}
			wantAddr := false
			if _, ok := ast.Unparen(e.Key).(*ast.CompositeLit); ok {
				wantAddr = isPointerCore(t.Key())
			}

			var key Value
			if wantAddr {
				// A CompositeLit never evaluates to a pointer,
				// so if the type of the location is a pointer,
				// an &-operation is implied.
				key = b.addr(fn, e.Key, true).address(fn)
			} else {
				key = b.expr(fn, e.Key)
			}

			loc := element{
				m:   m,
				k:   emitConv(fn, key, t.Key()),
				t:   t.Elem(),
				pos: e.Colon,
			}

			// We call assign() only because it takes care
			// of any &-operation required in the recursive
			// case, e.g.,
			// map[int]*struct{}{0: {}} implies &struct{}{}.
			// In-place update is of course impossible,
			// and no storebuf is needed.
			b.assign(fn, &loc, e.Value, true, nil)
		}
		sb.store(&address{addr: addr, pos: e.Lbrace, expr: e}, m)

	default:
		panic("unexpected CompositeLit type: " + typ.String())
	}
}

// switchStmt emits to fn code for the switch statement s, optionally
// labelled by label.
func (b *builder) switchStmt(fn *Function, s *ast.SwitchStmt, label *lblock) {
	// We treat SwitchStmt like a sequential if-else chain.
	// Multiway dispatch can be recovered later by ssautil.Switches()
	// to those cases that are free of side effects.
	if s.Init != nil {
		b.stmt(fn, s.Init)
	}
	var tag Value = vTrue
	if s.Tag != nil {
		tag = b.expr(fn, s.Tag)
	}
	done := fn.newBasicBlock("switch.done")
	if label != nil {
		label._break = done
	}
	// We pull the default case (if present) down to the end.
	// But each fallthrough label must point to the next
	// body block in source order, so we preallocate a
	// body block (fallthru) for the next case.
	// Unfortunately this makes for a confusing block order.
	var dfltBody *[]ast.Stmt
	var dfltFallthrough *BasicBlock
	var fallthru, dfltBlock *BasicBlock
	ncases := len(s.Body.List)
	for i, clause := range s.Body.List {
		body := fallthru
		if body == nil {
			body = fn.newBasicBlock("switch.body") // first case only
		}

		// Preallocate body block for the next case.
		fallthru = done
		if i+1 < ncases {
			fallthru = fn.newBasicBlock("switch.body")
		}

		cc := clause.(*ast.CaseClause)
		if cc.List == nil {
			// Default case.
			dfltBody = &cc.Body
			dfltFallthrough = fallthru
			dfltBlock = body
			continue
		}

		var nextCond *BasicBlock
		for _, cond := range cc.List {
			nextCond = fn.newBasicBlock("switch.next")
			// TODO(adonovan): opt: when tag==vTrue, we'd
			// get better code if we use b.cond(cond)
			// instead of BinOp(EQL, tag, b.expr(cond))
			// followed by If.  Don't forget conversions
			// though.
			cond := emitCompare(fn, token.EQL, tag, b.expr(fn, cond), cond.Pos())
			emitIf(fn, cond, body, nextCond)
			fn.currentBlock = nextCond
		}
		fn.currentBlock = body
		fn.targets = &targets{
			tail:         fn.targets,
			_break:       done,
			_fallthrough: fallthru,
		}
		b.stmtList(fn, cc.Body)
		fn.targets = fn.targets.tail
		emitJump(fn, done)
		fn.currentBlock = nextCond
	}
	if dfltBlock != nil {
		emitJump(fn, dfltBlock)
		fn.currentBlock = dfltBlock
		fn.targets = &targets{
			tail:         fn.targets,
			_break:       done,
			_fallthrough: dfltFallthrough,
		}
		b.stmtList(fn, *dfltBody)
		fn.targets = fn.targets.tail
	}
	emitJump(fn, done)
	fn.currentBlock = done
}

// typeSwitchStmt emits to fn code for the type switch statement s, optionally
// labelled by label.
func (b *builder) typeSwitchStmt(fn *Function, s *ast.TypeSwitchStmt, label *lblock) {
	// We treat TypeSwitchStmt like a sequential if-else chain.
	// Multiway dispatch can be recovered later by ssautil.Switches().

	// Typeswitch lowering:
	//
	// var x X
	// switch y := x.(type) {
	// case T1, T2: S1                  // >1 	(y := x)
	// case nil:    SN                  // nil 	(y := x)
	// default:     SD                  // 0 types 	(y := x)
	// case T3:     S3                  // 1 type 	(y := x.(T3))
	// }
	//
	//      ...s.Init...
	// 	x := eval x
	// .caseT1:
	// 	t1, ok1 := typeswitch,ok x <T1>
	// 	if ok1 then goto S1 else goto .caseT2
	// .caseT2:
	// 	t2, ok2 := typeswitch,ok x <T2>
	// 	if ok2 then goto S1 else goto .caseNil
	// .S1:
	//      y := x
	// 	...S1...
	// 	goto done
	// .caseNil:
	// 	if t2, ok2 := typeswitch,ok x <T2>
	// 	if x == nil then goto SN else goto .caseT3
	// .SN:
	//      y := x
	// 	...SN...
	// 	goto done
	// .caseT3:
	// 	t3, ok3 := typeswitch,ok x <T3>
	// 	if ok3 then goto S3 else goto default
	// .S3:
	//      y := t3
	// 	...S3...
	// 	goto done
	// .default:
	//      y := x
	// 	...SD...
	// 	goto done
	// .done:
	if s.Init != nil {
		b.stmt(fn, s.Init)
	}

	var x Value
	switch ass := s.Assign.(type) {
	case *ast.ExprStmt: // x.(type)
		x = b.expr(fn, ast.Unparen(ass.X).(*ast.TypeAssertExpr).X)
	case *ast.AssignStmt: // y := x.(type)
		x = b.expr(fn, ast.Unparen(ass.Rhs[0]).(*ast.TypeAssertExpr).X)
	}

	done := fn.newBasicBlock("typeswitch.done")
	if label != nil {
		label._break = done
	}
	var default_ *ast.CaseClause
	for _, clause := range s.Body.List {
		cc := clause.(*ast.CaseClause)
		if cc.List == nil {
			default_ = cc
			continue
		}
		body := fn.newBasicBlock("typeswitch.body")
		var next *BasicBlock
		var casetype types.Type
		var ti Value // ti, ok := typeassert,ok x <Ti>
		for _, cond := range cc.List {
			next = fn.newBasicBlock("typeswitch.next")
			casetype = fn.typeOf(cond)
			var condv Value
			if casetype == tUntypedNil {
				condv = emitCompare(fn, token.EQL, x, zeroConst(x.Type()), cond.Pos())
				ti = x
			} else {
				yok := emitTypeTest(fn, x, casetype, cc.Case)
				ti = emitExtract(fn, yok, 0)
				condv = emitExtract(fn, yok, 1)
			}
			emitIf(fn, condv, body, next)
			fn.currentBlock = next
		}
		if len(cc.List) != 1 {
			ti = x
		}
		fn.currentBlock = body
		b.typeCaseBody(fn, cc, ti, done)
		fn.currentBlock = next
	}
	if default_ != nil {
		b.typeCaseBody(fn, default_, x, done)
	} else {
		emitJump(fn, done)
	}
	fn.currentBlock = done
}

func (b *builder) typeCaseBody(fn *Function, cc *ast.CaseClause, x Value, done *BasicBlock) {
	if obj, ok := fn.info.Implicits[cc].(*types.Var); ok {
		// In a switch y := x.(type), each case clause
		// implicitly declares a distinct object y.
		// In a single-type case, y has that type.
		// In multi-type cases, 'case nil' and default,
		// y has the same type as the interface operand.
		emitStore(fn, emitLocalVar(fn, obj), x, obj.Pos())
	}
	fn.targets = &targets{
		tail:   fn.targets,
		_break: done,
	}
	b.stmtList(fn, cc.Body)
	fn.targets = fn.targets.tail
	emitJump(fn, done)
}

// selectStmt emits to fn code for the select statement s, optionally
// labelled by label.
func (b *builder) selectStmt(fn *Function, s *ast.SelectStmt, label *lblock) {
	// A blocking select of a single case degenerates to a
	// simple send or receive.
	// TODO(adonovan): opt: is this optimization worth its weight?
	if len(s.Body.List) == 1 {
		clause := s.Body.List[0].(*ast.CommClause)
		if clause.Comm != nil {
			b.stmt(fn, clause.Comm)
			done := fn.newBasicBlock("select.done")
			if label != nil {
				label._break = done
			}
			fn.targets = &targets{
				tail:   fn.targets,
				_break: done,
			}
			b.stmtList(fn, clause.Body)
			fn.targets = fn.targets.tail
			emitJump(fn, done)
			fn.currentBlock = done
			return
		}
	}

	// First evaluate all channels in all cases, and find
	// the directions of each state.
	var states []*SelectState
	blocking := true
	debugInfo := fn.debugInfo()
	for _, clause := range s.Body.List {
		var st *SelectState
		switch comm := clause.(*ast.CommClause).Comm.(type) {
		case nil: // default case
			blocking = false
			continue

		case *ast.SendStmt: // ch<- i
			ch := b.expr(fn, comm.Chan)
			chtyp := typeparams.CoreType(fn.typ(ch.Type())).(*types.Chan)
			st = &SelectState{
				Dir:  types.SendOnly,
				Chan: ch,
				Send: emitConv(fn, b.expr(fn, comm.Value), chtyp.Elem()),
				Pos:  comm.Arrow,
			}
			if debugInfo {
				st.DebugNode = comm
			}

		case *ast.AssignStmt: // x := <-ch
			recv := ast.Unparen(comm.Rhs[0]).(*ast.UnaryExpr)
			st = &SelectState{
				Dir:  types.RecvOnly,
				Chan: b.expr(fn, recv.X),
				Pos:  recv.OpPos,
			}
			if debugInfo {
				st.DebugNode = recv
			}

		case *ast.ExprStmt: // <-ch
			recv := ast.Unparen(comm.X).(*ast.UnaryExpr)
			st = &SelectState{
				Dir:  types.RecvOnly,
				Chan: b.expr(fn, recv.X),
				Pos:  recv.OpPos,
			}
			if debugInfo {
				st.DebugNode = recv
			}
		}
		states = append(states, st)
	}

	// We dispatch on the (fair) result of Select using a
	// sequential if-else chain, in effect:
	//
	// idx, recvOk, r0...r_n-1 := select(...)
	// if idx == 0 {  // receive on channel 0  (first receive => r0)
	//     x, ok := r0, recvOk
	//     ...state0...
	// } else if v == 1 {   // send on channel 1
	//     ...state1...
	// } else {
	//     ...default...
	// }
	sel := &Select{
		States:   states,
		Blocking: blocking,
	}
	sel.setPos(s.Select)
	var vars []*types.Var
	vars = append(vars, varIndex, varOk)
	for _, st := range states {
		if st.Dir == types.RecvOnly {
			chtyp := typeparams.CoreType(fn.typ(st.Chan.Type())).(*types.Chan)
			vars = append(vars, anonVar(chtyp.Elem()))
		}
	}
	sel.setType(types.NewTuple(vars...))

	fn.emit(sel)
	idx := emitExtract(fn, sel, 0)

	done := fn.newBasicBlock("select.done")
	if label != nil {
		label._break = done
	}

	var defaultBody *[]ast.Stmt
	state := 0
	r := 2 // index in 'sel' tuple of value; increments if st.Dir==RECV
	for _, cc := range s.Body.List {
		clause := cc.(*ast.CommClause)
		if clause.Comm == nil {
			defaultBody = &clause.Body
			continue
		}
		body := fn.newBasicBlock("select.body")
		next := fn.newBasicBlock("select.next")
		emitIf(fn, emitCompare(fn, token.EQL, idx, intConst(int64(state)), token.NoPos), body, next)
		fn.currentBlock = body
		fn.targets = &targets{
			tail:   fn.targets,
			_break: done,
		}
		switch comm := clause.Comm.(type) {
		case *ast.ExprStmt: // <-ch
			if debugInfo {
				v := emitExtract(fn, sel, r)
				emitDebugRef(fn, states[state].DebugNode.(ast.Expr), v, false)
			}
			r++

		case *ast.AssignStmt: // x := <-states[state].Chan
			if comm.Tok == token.DEFINE {
				emitLocalVar(fn, identVar(fn, comm.Lhs[0].(*ast.Ident)))
			}
			x := b.addr(fn, comm.Lhs[0], false) // non-escaping
			v := emitExtract(fn, sel, r)
			if debugInfo {
				emitDebugRef(fn, states[state].DebugNode.(ast.Expr), v, false)
			}
			x.store(fn, v)

			if len(comm.Lhs) == 2 { // x, ok := ...
				if comm.Tok == token.DEFINE {
					emitLocalVar(fn, identVar(fn, comm.Lhs[1].(*ast.Ident)))
				}
				ok := b.addr(fn, comm.Lhs[1], false) // non-escaping
				ok.store(fn, emitExtract(fn, sel, 1))
			}
			r++
		}
		b.stmtList(fn, clause.Body)
		fn.targets = fn.targets.tail
		emitJump(fn, done)
		fn.currentBlock = next
		state++
	}
	if defaultBody != nil {
		fn.targets = &targets{
			tail:   fn.targets,
			_break: done,
		}
		b.stmtList(fn, *defaultBody)
		fn.targets = fn.targets.tail
	} else {
		// A blocking select must match some case.
		// (This should really be a runtime.errorString, not a string.)
		fn.emit(&Panic{
			X: emitConv(fn, stringConst("blocking select matched no case"), tEface),
		})
		fn.currentBlock = fn.newBasicBlock("unreachable")
	}
	emitJump(fn, done)
	fn.currentBlock = done
}

// forStmt emits to fn code for the for statement s, optionally
// labelled by label.
func (b *builder) forStmt(fn *Function, s *ast.ForStmt, label *lblock) {
	// Use forStmtGo122 instead if it applies.
	if s.Init != nil {
		if assign, ok := s.Init.(*ast.AssignStmt); ok && assign.Tok == token.DEFINE {
			if versions.AtLeast(fn.goversion, versions.Go1_22) {
				b.forStmtGo122(fn, s, label)
				return
			}
		}
	}

	//     ...init...
	//     jump loop
	// loop:
	//     if cond goto body else done
	// body:
	//     ...body...
	//     jump post
	// post:                                 (target of continue)
	//     ...post...
	//     jump loop
	// done:                                 (target of break)
	if s.Init != nil {
		b.stmt(fn, s.Init)
	}

	body := fn.newBasicBlock("for.body")
	done := fn.newBasicBlock("for.done") // target of 'break'
	loop := body                         // target of back-edge
	if s.Cond != nil {
		loop = fn.newBasicBlock("for.loop")
	}
	cont := loop // target of 'continue'
	if s.Post != nil {
		cont = fn.newBasicBlock("for.post")
	}
	if label != nil {
		label._break = done
		label._continue = cont
	}
	emitJump(fn, loop)
	fn.currentBlock = loop
	if loop != body {
		b.cond(fn, s.Cond, body, done)
		fn.currentBlock = body
	}
	fn.targets = &targets{
		tail:      fn.targets,
		_break:    done,
		_continue: cont,
	}
	b.stmt(fn, s.Body)
	fn.targets = fn.targets.tail
	emitJump(fn, cont)

	if s.Post != nil {
		fn.currentBlock = cont
		b.stmt(fn, s.Post)
		emitJump(fn, loop) // back-edge
	}
	fn.currentBlock = done
}

// forStmtGo122 emits to fn code for the for statement s, optionally
// labelled by label. s must define its variables.
//
// This allocates once per loop iteration. This is only correct in
// GoVersions >= go1.22.
func (b *builder) forStmtGo122(fn *Function, s *ast.ForStmt, label *lblock) {
	//     i_outer = alloc[T]
	//     *i_outer = ...init...        // under objects[i] = i_outer
	//     jump loop
	// loop:
	//     i = phi [head: i_outer, loop: i_next]
	//     ...cond...                   // under objects[i] = i
	//     if cond goto body else done
	// body:
	//     ...body...                   // under objects[i] = i (same as loop)
	//     jump post
	// post:
	//     tmp = *i
	//     i_next = alloc[T]
	//     *i_next = tmp
	//     ...post...                   // under objects[i] = i_next
	//     goto loop
	// done:

	init := s.Init.(*ast.AssignStmt)
	startingBlocks := len(fn.Blocks)

	pre := fn.currentBlock               // current block before starting
	loop := fn.newBasicBlock("for.loop") // target of back-edge
	body := fn.newBasicBlock("for.body")
	post := fn.newBasicBlock("for.post") // target of 'continue'
	done := fn.newBasicBlock("for.done") // target of 'break'

	// For each of the n loop variables, we create five SSA values,
	// outer, phi, next, load, and store in pre, loop, and post.
	// There is no limit on n.
	type loopVar struct {
		obj   *types.Var
		outer *Alloc
		phi   *Phi
		load  *UnOp
		next  *Alloc
		store *Store
	}
	vars := make([]loopVar, len(init.Lhs))
	for i, lhs := range init.Lhs {
		v := identVar(fn, lhs.(*ast.Ident))
		typ := fn.typ(v.Type())

		fn.currentBlock = pre
		outer := emitLocal(fn, typ, v.Pos(), v.Name())

		fn.currentBlock = loop
		phi := &Phi{Comment: v.Name()}
		phi.pos = v.Pos()
		phi.typ = outer.Type()
		fn.emit(phi)

		fn.currentBlock = post
		// If next is local, it reuses the address and zeroes the old value so
		// load before allocating next.
		load := emitLoad(fn, phi)
		next := emitLocal(fn, typ, v.Pos(), v.Name())
		store := emitStore(fn, next, load, token.NoPos)

		phi.Edges = []Value{outer, next} // pre edge is emitted before post edge.

		vars[i] = loopVar{v, outer, phi, load, next, store}
	}

	// ...init... under fn.objects[v] = i_outer
	fn.currentBlock = pre
	for _, v := range vars {
		fn.vars[v.obj] = v.outer
	}
	const isDef = false // assign to already-allocated outers
	b.assignStmt(fn, init.Lhs, init.Rhs, isDef)
	if label != nil {
		label._break = done
		label._continue = post
	}
	emitJump(fn, loop)

	// ...cond... under fn.objects[v] = i
	fn.currentBlock = loop
	for _, v := range vars {
		fn.vars[v.obj] = v.phi
	}
	if s.Cond != nil {
		b.cond(fn, s.Cond, body, done)
	} else {
		emitJump(fn, body)
	}

	// ...body... under fn.objects[v] = i
	fn.currentBlock = body
	fn.targets = &targets{
		tail:      fn.targets,
		_break:    done,
		_continue: post,
	}
	b.stmt(fn, s.Body)
	fn.targets = fn.targets.tail
	emitJump(fn, post)

	// ...post... under fn.objects[v] = i_next
	for _, v := range vars {
		fn.vars[v.obj] = v.next
	}
	fn.currentBlock = post
	if s.Post != nil {
		b.stmt(fn, s.Post)
	}
	emitJump(fn, loop) // back-edge
	fn.currentBlock = done

	// For each loop variable that does not escape,
	// (the common case), fuse its next cells into its
	// (local) outer cell as they have disjoint live ranges.
	//
	// It is sufficient to test whether i_next escapes,
	// because its Heap flag will be marked true if either
	// the cond or post expression causes i to escape
	// (because escape distributes over phi).
	var nlocals int
	for _, v := range vars {
		if !v.next.Heap {
			nlocals++
		}
	}
	if nlocals > 0 {
		replace := make(map[Value]Value, 2*nlocals)
		dead := make(map[Instruction]bool, 4*nlocals)
		for _, v := range vars {
			if !v.next.Heap {
				replace[v.next] = v.outer
				replace[v.phi] = v.outer
				dead[v.phi], dead[v.next], dead[v.load], dead[v.store] = true, true, true, true
			}
		}

		// Replace all uses of i_next and phi with i_outer.
		// Referrers have not been built for fn yet so only update Instruction operands.
		// We need only look within the blocks added by the loop.
		var operands []*Value // recycle storage
		for _, b := range fn.Blocks[startingBlocks:] {
			for _, instr := range b.Instrs {
				operands = instr.Operands(operands[:0])
				for _, ptr := range operands {
					k := *ptr
					if v := replace[k]; v != nil {
						*ptr = v
					}
				}
			}
		}

		// Remove instructions for phi, load, and store.
		// lift() will remove the unused i_next *Alloc.
		isDead := func(i Instruction) bool { return dead[i] }
		loop.Instrs = slices.DeleteFunc(loop.Instrs, isDead)
		post.Instrs = slices.DeleteFunc(post.Instrs, isDead)
	}
}

// rangeIndexed emits to fn the header for an integer-indexed loop
// over array, *array or slice value x.
// The v result is defined only if tv is non-nil.
// forPos is the position of the "for" token.
func (b *builder) rangeIndexed(fn *Function, x Value, tv types.Type, pos token.Pos) (k, v Value, loop, done *BasicBlock) {
	//
	//     length = len(x)
	//     index = -1
	// loop:                                     (target of continue)
	//     index++
	//     if index < length goto body else done
	// body:
	//     k = index
	//     v = x[index]
	//     ...body...
	//     jump loop
	// done:                                     (target of break)

	// Determine number of iterations.
	var length Value
	dt := typeparams.Deref(x.Type())
	if arr, ok := typeparams.CoreType(dt).(*types.Array); ok {
		// For array or *array, the number of iterations is
		// known statically thanks to the type.  We avoid a
		// data dependence upon x, permitting later dead-code
		// elimination if x is pure, static unrolling, etc.
		// Ranging over a nil *array may have >0 iterations.
		// We still generate code for x, in case it has effects.
		length = intConst(arr.Len())
	} else {
		// length = len(x).
		var c Call
		c.Call.Value = makeLen(x.Type())
		c.Call.Args = []Value{x}
		c.setType(tInt)
		length = fn.emit(&c)
	}

	index := emitLocal(fn, tInt, token.NoPos, "rangeindex")
	emitStore(fn, index, intConst(-1), pos)

	loop = fn.newBasicBlock("rangeindex.loop")
	emitJump(fn, loop)
	fn.currentBlock = loop

	incr := &BinOp{
		Op: token.ADD,
		X:  emitLoad(fn, index),
		Y:  vOne,
	}
	incr.setType(tInt)
	emitStore(fn, index, fn.emit(incr), pos)

	body := fn.newBasicBlock("rangeindex.body")
	done = fn.newBasicBlock("rangeindex.done")
	emitIf(fn, emitCompare(fn, token.LSS, incr, length, token.NoPos), body, done)
	fn.currentBlock = body

	k = emitLoad(fn, index)
	if tv != nil {
		switch t := typeparams.CoreType(x.Type()).(type) {
		case *types.Array:
			instr := &Index{
				X:     x,
				Index: k,
			}
			instr.setType(t.Elem())
			instr.setPos(x.Pos())
			v = fn.emit(instr)

		case *types.Pointer: // *array
			instr := &IndexAddr{
				X:     x,
				Index: k,
			}
			instr.setType(types.NewPointer(t.Elem().Underlying().(*types.Array).Elem()))
			instr.setPos(x.Pos())
			v = emitLoad(fn, fn.emit(instr))

		case *types.Slice:
			instr := &IndexAddr{
				X:     x,
				Index: k,
			}
			instr.setType(types.NewPointer(t.Elem()))
			instr.setPos(x.Pos())
			v = emitLoad(fn, fn.emit(instr))

		default:
			panic("rangeIndexed x:" + t.String())
		}
	}
	return
}

// rangeIter emits to fn the header for a loop using
// Range/Next/Extract to iterate over map or string value x.
// tk and tv are the types of the key/value results k and v, or nil
// if the respective component is not wanted.
func (b *builder) rangeIter(fn *Function, x Value, tk, tv types.Type, pos token.Pos) (k, v Value, loop, done *BasicBlock) {
	//
	//     it = range x
	// loop:                                   (target of continue)
	//     okv = next it                       (ok, key, value)
	//     ok = extract okv #0
	//     if ok goto body else done
	// body:
	//     k = extract okv #1
	//     v = extract okv #2
	//     ...body...
	//     jump loop
	// done:                                   (target of break)
	//

	if tk == nil {
		tk = tInvalid
	}
	if tv == nil {
		tv = tInvalid
	}

	rng := &Range{X: x}
	rng.setPos(pos)
	rng.setType(tRangeIter)
	it := fn.emit(rng)

	loop = fn.newBasicBlock("rangeiter.loop")
	emitJump(fn, loop)
	fn.currentBlock = loop

	okv := &Next{
		Iter:     it,
		IsString: isBasic(typeparams.CoreType(x.Type())),
	}
	okv.setType(types.NewTuple(
		varOk,
		newVar("k", tk),
		newVar("v", tv),
	))
	fn.emit(okv)

	body := fn.newBasicBlock("rangeiter.body")
	done = fn.newBasicBlock("rangeiter.done")
	emitIf(fn, emitExtract(fn, okv, 0), body, done)
	fn.currentBlock = body

	if tk != tInvalid {
		k = emitExtract(fn, okv, 1)
	}
	if tv != tInvalid {
		v = emitExtract(fn, okv, 2)
	}
	return
}

// rangeChan emits to fn the header for a loop that receives from
// channel x until it fails.
// tk is the channel's element type, or nil if the k result is
// not wanted
// pos is the position of the '=' or ':=' token.
func (b *builder) rangeChan(fn *Function, x Value, tk types.Type, pos token.Pos) (k Value, loop, done *BasicBlock) {
	//
	// loop:                                   (target of continue)
	//     ko = <-x                            (key, ok)
	//     ok = extract ko #1
	//     if ok goto body else done
	// body:
	//     k = extract ko #0
	//     ...body...
	//     goto loop
	// done:                                   (target of break)

	loop = fn.newBasicBlock("rangechan.loop")
	emitJump(fn, loop)
	fn.currentBlock = loop
	recv := &UnOp{
		Op:      token.ARROW,
		X:       x,
		CommaOk: true,
	}
	recv.setPos(pos)
	recv.setType(types.NewTuple(
		newVar("k", typeparams.CoreType(x.Type()).(*types.Chan).Elem()),
		varOk,
	))
	ko := fn.emit(recv)
	body := fn.newBasicBlock("rangechan.body")
	done = fn.newBasicBlock("rangechan.done")
	emitIf(fn, emitExtract(fn, ko, 1), body, done)
	fn.currentBlock = body
	if tk != nil {
		k = emitExtract(fn, ko, 0)
	}
	return
}

// rangeInt emits to fn the header for a range loop with an integer operand.
// tk is the key value's type, or nil if the k result is not wanted.
// pos is the position of the "for" token.
func (b *builder) rangeInt(fn *Function, x Value, tk types.Type, pos token.Pos) (k Value, loop, done *BasicBlock) {
	//
	//     iter = 0
	//     if 0 < x goto body else done
	// loop:                                   (target of continue)
	//     iter++
	//     if iter < x goto body else done
	// body:
	//     k = x
	//     ...body...
	//     jump loop
	// done:                                   (target of break)

	if isUntyped(x.Type()) {
		x = emitConv(fn, x, tInt)
	}

	T := x.Type()
	iter := emitLocal(fn, T, token.NoPos, "rangeint.iter")
	// x may be unsigned. Avoid initializing x to -1.

	body := fn.newBasicBlock("rangeint.body")
	done = fn.newBasicBlock("rangeint.done")
	emitIf(fn, emitCompare(fn, token.LSS, zeroConst(T), x, token.NoPos), body, done)

	loop = fn.newBasicBlock("rangeint.loop")
	fn.currentBlock = loop

	incr := &BinOp{
		Op: token.ADD,
		X:  emitLoad(fn, iter),
		Y:  emitConv(fn, vOne, T),
	}
	incr.setType(T)
	emitStore(fn, iter, fn.emit(incr), pos)
	emitIf(fn, emitCompare(fn, token.LSS, incr, x, token.NoPos), body, done)
	fn.currentBlock = body

	if tk != nil {
		// Integer types (int, uint8, etc.) are named and
		// we know that k is assignable to x when tk != nil.
		// This implies tk and T are identical so no conversion is needed.
		k = emitLoad(fn, iter)
	}

	return
}

// rangeStmt emits to fn code for the range statement s, optionally
// labelled by label.
func (b *builder) rangeStmt(fn *Function, s *ast.RangeStmt, label *lblock) {
	var tk, tv types.Type
	if s.Key != nil && !isBlankIdent(s.Key) {
		tk = fn.typeOf(s.Key)
	}
	if s.Value != nil && !isBlankIdent(s.Value) {
		tv = fn.typeOf(s.Value)
	}

	// create locals for s.Key and s.Value.
	createVars := func() {
		// Unlike a short variable declaration, a RangeStmt
		// using := never redeclares an existing variable; it
		// always creates a new one.
		if tk != nil {
			emitLocalVar(fn, identVar(fn, s.Key.(*ast.Ident)))
		}
		if tv != nil {
			emitLocalVar(fn, identVar(fn, s.Value.(*ast.Ident)))
		}
	}

	afterGo122 := versions.AtLeast(fn.goversion, versions.Go1_22)
	if s.Tok == token.DEFINE && !afterGo122 {
		// pre-go1.22: If iteration variables are defined (:=), this
		// occurs once outside the loop.
		createVars()
	}

	x := b.expr(fn, s.X)

	var k, v Value
	var loop, done *BasicBlock
	switch rt := typeparams.CoreType(x.Type()).(type) {
	case *types.Slice, *types.Array, *types.Pointer: // *array
		k, v, loop, done = b.rangeIndexed(fn, x, tv, s.For)

	case *types.Chan:
		k, loop, done = b.rangeChan(fn, x, tk, s.For)

	case *types.Map:
		k, v, loop, done = b.rangeIter(fn, x, tk, tv, s.For)

	case *types.Basic:
		switch {
		case rt.Info()&types.IsString != 0:
			k, v, loop, done = b.rangeIter(fn, x, tk, tv, s.For)

		case rt.Info()&types.IsInteger != 0:
			k, loop, done = b.rangeInt(fn, x, tk, s.For)

		default:
			panic("Cannot range over basic type: " + rt.String())
		}

	case *types.Signature:
		// Special case rewrite (fn.goversion >= go1.23):
		// 	for x := range f { ... }
		// into
		// 	f(func(x T) bool { ... })
		b.rangeFunc(fn, x, tk, tv, s, label)
		return

	default:
		panic("Cannot range over: " + rt.String())
	}

	if s.Tok == token.DEFINE && afterGo122 {
		// go1.22: If iteration variables are defined (:=), this occurs inside the loop.
		createVars()
	}

	// Evaluate both LHS expressions before we update either.
	var kl, vl lvalue
	if tk != nil {
		kl = b.addr(fn, s.Key, false) // non-escaping
	}
	if tv != nil {
		vl = b.addr(fn, s.Value, false) // non-escaping
	}
	if tk != nil {
		kl.store(fn, k)
	}
	if tv != nil {
		vl.store(fn, v)
	}

	if label != nil {
		label._break = done
		label._continue = loop
	}

	fn.targets = &targets{
		tail:      fn.targets,
		_break:    done,
		_continue: loop,
	}
	b.stmt(fn, s.Body)
	fn.targets = fn.targets.tail
	emitJump(fn, loop) // back-edge
	fn.currentBlock = done
}

// rangeFunc emits to fn code for the range-over-func rng.Body of the iterator
// function x, optionally labelled by label. It creates a new anonymous function
// yield for rng and builds the function.
func (b *builder) rangeFunc(fn *Function, x Value, tk, tv types.Type, rng *ast.RangeStmt, label *lblock) {
	// Consider the SSA code for the outermost range-over-func in fn:
	//
	//   func fn(...) (ret R) {
	//     ...
	//     for k, v = range x {
	// 	     ...
	//     }
	//     ...
	//   }
	//
	// The code emitted into fn will look something like this.
	//
	// loop:
	//     jump := READY
	//     y := make closure yield [ret, deferstack, jump, k, v]
	//     x(y)
	//     switch jump {
	//        [see resuming execution]
	//     }
	//     goto done
	// done:
	//     ...
	//
	// where yield is a new synthetic yield function:
	//
	// func yield(_k tk, _v tv) bool
	//   free variables: [ret, stack, jump, k, v]
	// {
	//    entry:
	//      if jump != READY then goto invalid else valid
	//    invalid:
	//      panic("iterator called when it is not in a ready state")
	//    valid:
	//      jump = BUSY
	//      k = _k
	//      v = _v
	//    ...
	//    cont:
	//      jump = READY
	//      return true
	// }
	//
	// Yield state:
	//
	// Each range loop has an associated jump variable that records
	// the state of the iterator. A yield function is initially
	// in a READY (0) and callable state.  If the yield function is called
	// and is not in READY state, it panics. When it is called in a callable
	// state, it becomes BUSY. When execution reaches the end of the body
	// of the loop (or a continue statement targeting the loop is executed),
	// the yield function returns true and resumes being in a READY state.
	// After the iterator function x(y) returns, then if the yield function
	// is in a READY state, the yield enters the DONE state.
	//
	// Each lowered control statement (break X, continue X, goto Z, or return)
	// that exits the loop sets the variable to a unique positive EXIT value,
	// before returning false from the yield function.
	//
	// If the yield function returns abruptly due to a panic or GoExit,
	// it remains in a BUSY state. The generated code asserts that, after
	// the iterator call x(y) returns normally, the jump variable state
	// is DONE.
	//
	// Resuming execution:
	//
	// The code generated for the range statement checks the jump
	// variable to determine how to resume execution.
	//
	//    switch jump {
	//    case BUSY:  panic("...")
	//    case DONE:  goto done
	//    case READY: state = DONE; goto done
	//    case 123:   ... // action for exit 123.
	//    case 456:   ... // action for exit 456.
	//    ...
	//    }
	//
	// Forward goto statements within a yield are jumps to labels that
	// have not yet been traversed in fn. They may be in the Body of the
	// function. What we emit for these is:
	//
	//    goto target
	//  target:
	//    ...
	//
	// We leave an unresolved exit in yield.exits to check at the end
	// of building yield if it encountered target in the body. If it
	// encountered target, no additional work is required. Otherwise,
	// the yield emits a new early exit in the basic block for target.
	// We expect that blockopt will fuse the early exit into the case
	// block later. The unresolved exit is then added to yield.parent.exits.

	loop := fn.newBasicBlock("rangefunc.loop")
	done := fn.newBasicBlock("rangefunc.done")

	// These are targets within y.
	fn.targets = &targets{
		tail:   fn.targets,
		_break: done,
		// _continue is within y.
	}
	if label != nil {
		label._break = done
		// _continue is within y
	}

	emitJump(fn, loop)
	fn.currentBlock = loop

	// loop:
	//     jump := READY

	anonIdx := len(fn.AnonFuncs)

	jump := newVar(fmt.Sprintf("jump$%d", anonIdx+1), tInt)
	emitLocalVar(fn, jump) // zero value is READY

	xsig := typeparams.CoreType(x.Type()).(*types.Signature)
	ysig := typeparams.CoreType(xsig.Params().At(0).Type()).(*types.Signature)

	/* synthetic yield function for body of range-over-func loop */
	y := &Function{
		name:           fmt.Sprintf("%s$%d", fn.Name(), anonIdx+1),
		Signature:      ysig,
		Synthetic:      "range-over-func yield",
		pos:            rng.Range,
		parent:         fn,
		anonIdx:        int32(len(fn.AnonFuncs)),
		Pkg:            fn.Pkg,
		Prog:           fn.Prog,
		syntax:         rng,
		info:           fn.info,
		goversion:      fn.goversion,
		build:          (*builder).buildYieldFunc,
		topLevelOrigin: nil,
		typeparams:     fn.typeparams,
		typeargs:       fn.typeargs,
		subst:          fn.subst,
		jump:           jump,
		deferstack:     fn.deferstack,
		returnVars:     fn.returnVars, // use the parent's return variables
		uniq:           fn.uniq,       // start from parent's unique values
	}

	// If the RangeStmt has a label, this is how it is passed to buildYieldFunc.
	if label != nil {
		y.lblocks = map[*types.Label]*lblock{label.label: nil}
	}
	fn.AnonFuncs = append(fn.AnonFuncs, y)

	// Build y immediately. It may:
	// * cause fn's locals to escape, and
	// * create new exit nodes in exits.
	// (y is not marked 'built' until the end of the enclosing FuncDecl.)
	unresolved := len(fn.exits)
	y.build(b, y)
	fn.uniq = y.uniq // resume after y's unique values

	// Emit the call of y.
	//   c := MakeClosure y
	//   x(c)
	c := &MakeClosure{Fn: y}
	c.setType(ysig)
	for _, fv := range y.FreeVars {
		c.Bindings = append(c.Bindings, fv.outer)
		fv.outer = nil
	}
	fn.emit(c)
	call := Call{
		Call: CallCommon{
			Value: x,
			Args:  []Value{c},
			pos:   token.NoPos,
		},
	}
	call.setType(xsig.Results())
	fn.emit(&call)

	exits := fn.exits[unresolved:]
	b.buildYieldResume(fn, jump, exits, done)

	emitJump(fn, done)
	fn.currentBlock = done
	// pop the stack for the range-over-func
	fn.targets = fn.targets.tail
}

// buildYieldResume emits to fn code for how to resume execution once a call to
// the iterator function over the yield function returns x(y). It does this by building
// a switch over the value of jump for when it is READY, BUSY, or EXIT(id).
func (b *builder) buildYieldResume(fn *Function, jump *types.Var, exits []*exit, done *BasicBlock) {
	//    v := *jump
	//    switch v {
	//    case BUSY:    panic("...")
	//    case READY:   jump = DONE; goto done
	//    case EXIT(a): ...
	//    case EXIT(b): ...
	//    ...
	//    }
	v := emitLoad(fn, fn.lookup(jump, false))

	// case BUSY: panic("...")
	isbusy := fn.newBasicBlock("rangefunc.resume.busy")
	ifready := fn.newBasicBlock("rangefunc.resume.ready.check")
	emitIf(fn, emitCompare(fn, token.EQL, v, jBusy, token.NoPos), isbusy, ifready)
	fn.currentBlock = isbusy
	fn.emit(&Panic{
		X: emitConv(fn, stringConst("iterator call did not preserve panic"), tEface),
	})
	fn.currentBlock = ifready

	// case READY: jump = DONE; goto done
	isready := fn.newBasicBlock("rangefunc.resume.ready")
	ifexit := fn.newBasicBlock("rangefunc.resume.exits")
	emitIf(fn, emitCompare(fn, token.EQL, v, jReady, token.NoPos), isready, ifexit)
	fn.currentBlock = isready
	storeVar(fn, jump, jDone, token.NoPos)
	emitJump(fn, done)
	fn.currentBlock = ifexit

	for _, e := range exits {
		id := intConst(e.id)

		//  case EXIT(id): { /* do e */ }
		cond := emitCompare(fn, token.EQL, v, id, e.pos)
		matchb := fn.newBasicBlock("rangefunc.resume.match")
		cndb := fn.newBasicBlock("rangefunc.resume.cnd")
		emitIf(fn, cond, matchb, cndb)
		fn.currentBlock = matchb

		// Cases to fill in the { /* do e */ } bit.
		switch {
		case e.label != nil: // forward goto?
			// case EXIT(id): goto lb // label
			lb := fn.lblockOf(e.label)
			// Do not mark lb as resolved.
			// If fn does not contain label, lb remains unresolved and
			// fn must itself be a range-over-func function. lb will be:
			//   lb:
			//     fn.jump = id
			//     return false
			emitJump(fn, lb._goto)

		case e.to != fn: // e jumps to an ancestor of fn?
			// case EXIT(id): { fn.jump = id; return false }
			// fn is a range-over-func function.
			storeVar(fn, fn.jump, id, token.NoPos)
			fn.emit(&Return{Results: []Value{vFalse}, pos: e.pos})

		case e.block == nil && e.label == nil: // return from fn?
			// case EXIT(id): { return ... }
			fn.emit(new(RunDefers))
			results := make([]Value, len(fn.results))
			for i, r := range fn.results {
				results[i] = emitLoad(fn, r)
			}
			fn.emit(&Return{Results: results, pos: e.pos})

		case e.block != nil:
			// case EXIT(id): goto block
			emitJump(fn, e.block)

		default:
			panic("unreachable")
		}
		fn.currentBlock = cndb
	}
}

// stmt lowers statement s to SSA form, emitting code to fn.
func (b *builder) stmt(fn *Function, _s ast.Stmt) {
	// The label of the current statement.  If non-nil, its _goto
	// target is always set; its _break and _continue are set only
	// within the body of switch/typeswitch/select/for/range.
	// It is effectively an additional default-nil parameter of stmt().
	var label *lblock
start:
	switch s := _s.(type) {
	case *ast.EmptyStmt:
		// ignore.  (Usually removed by gofmt.)

	case *ast.DeclStmt: // Con, Var or Typ
		d := s.Decl.(*ast.GenDecl)
		if d.Tok == token.VAR {
			for _, spec := range d.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					b.localValueSpec(fn, vs)
				}
			}
		}

	case *ast.LabeledStmt:
		if s.Label.Name == "_" {
			// Blank labels can't be the target of a goto, break,
			// or continue statement, so we don't need a new block.
			_s = s.Stmt
			goto start
		}
		label = fn.lblockOf(fn.label(s.Label))
		label.resolved = true
		emitJump(fn, label._goto)
		fn.currentBlock = label._goto
		_s = s.Stmt
		goto start // effectively: tailcall stmt(fn, s.Stmt, label)

	case *ast.ExprStmt:
		b.expr(fn, s.X)

	case *ast.SendStmt:
		chtyp := typeparams.CoreType(fn.typeOf(s.Chan)).(*types.Chan)
		fn.emit(&Send{
			Chan: b.expr(fn, s.Chan),
			X:    emitConv(fn, b.expr(fn, s.Value), chtyp.Elem()),
			pos:  s.Arrow,
		})

	case *ast.IncDecStmt:
		op := token.ADD
		if s.Tok == token.DEC {
			op = token.SUB
		}
		loc := b.addr(fn, s.X, false)
		b.assignOp(fn, loc, NewConst(constant.MakeInt64(1), loc.typ()), op, s.Pos())

	case *ast.AssignStmt:
		switch s.Tok {
		case token.ASSIGN, token.DEFINE:
			b.assignStmt(fn, s.Lhs, s.Rhs, s.Tok == token.DEFINE)

		default: // +=, etc.
			op := s.Tok + token.ADD - token.ADD_ASSIGN
			b.assignOp(fn, b.addr(fn, s.Lhs[0], false), b.expr(fn, s.Rhs[0]), op, s.Pos())
		}

	case *ast.GoStmt:
		// The "intrinsics" new/make/len/cap are forbidden here.
		// panic is treated like an ordinary function call.
		v := Go{pos: s.Go}
		b.setCall(fn, s.Call, &v.Call)
		fn.emit(&v)

	case *ast.DeferStmt:
		// The "intrinsics" new/make/len/cap are forbidden here.
		// panic is treated like an ordinary function call.
		deferstack := emitLoad(fn, fn.lookup(fn.deferstack, false))
		v := Defer{pos: s.Defer, DeferStack: deferstack}
		b.setCall(fn, s.Call, &v.Call)
		fn.emit(&v)

		// A deferred call can cause recovery from panic,
		// and control resumes at the Recover block.
		createRecoverBlock(fn.source)

	case *ast.ReturnStmt:
		b.returnStmt(fn, s)

	case *ast.BranchStmt:
		b.branchStmt(fn, s)

	case *ast.BlockStmt:
		b.stmtList(fn, s.List)

	case *ast.IfStmt:
		if s.Init != nil {
			b.stmt(fn, s.Init)
		}
		then := fn.newBasicBlock("if.then")
		done := fn.newBasicBlock("if.done")
		els := done
		if s.Else != nil {
			els = fn.newBasicBlock("if.else")
		}
		b.cond(fn, s.Cond, then, els)
		fn.currentBlock = then
		b.stmt(fn, s.Body)
		emitJump(fn, done)

		if s.Else != nil {
			fn.currentBlock = els
			b.stmt(fn, s.Else)
			emitJump(fn, done)
		}

		fn.currentBlock = done

	case *ast.SwitchStmt:
		b.switchStmt(fn, s, label)

	case *ast.TypeSwitchStmt:
		b.typeSwitchStmt(fn, s, label)

	case *ast.SelectStmt:
		b.selectStmt(fn, s, label)

	case *ast.ForStmt:
		b.forStmt(fn, s, label)

	case *ast.RangeStmt:
		b.rangeStmt(fn, s, label)

	default:
		panic(fmt.Sprintf("unexpected statement kind: %T", s))
	}
}

func (b *builder) branchStmt(fn *Function, s *ast.BranchStmt) {
	var block *BasicBlock
	if s.Label == nil {
		block = targetedBlock(fn, s.Tok)
	} else {
		target := fn.label(s.Label)
		block = labelledBlock(fn, target, s.Tok)
		if block == nil { // forward goto
			lb := fn.lblockOf(target)
			block = lb._goto // jump to lb._goto
			if fn.jump != nil {
				// fn is a range-over-func and the goto may exit fn.
				// Create an exit and resolve it at the end of
				// builder.buildYieldFunc.
				labelExit(fn, target, s.Pos())
			}
		}
	}
	to := block.parent

	if to == fn {
		emitJump(fn, block)
	} else { // break outside of fn.
		// fn must be a range-over-func
		e := blockExit(fn, block, s.Pos())
		storeVar(fn, fn.jump, intConst(e.id), e.pos)
		fn.emit(&Return{Results: []Value{vFalse}, pos: e.pos})
	}
	fn.currentBlock = fn.newBasicBlock("unreachable")
}

func (b *builder) returnStmt(fn *Function, s *ast.ReturnStmt) {
	var results []Value

	sig := fn.source.Signature // signature of the enclosing source function

	// Convert return operands to result type.
	if len(s.Results) == 1 && sig.Results().Len() > 1 {
		// Return of one expression in a multi-valued function.
		tuple := b.exprN(fn, s.Results[0])
		ttuple := tuple.Type().(*types.Tuple)
		for i, n := 0, ttuple.Len(); i < n; i++ {
			results = append(results,
				emitConv(fn, emitExtract(fn, tuple, i),
					sig.Results().At(i).Type()))
		}
	} else {
		// 1:1 return, or no-arg return in non-void function.
		for i, r := range s.Results {
			v := emitConv(fn, b.expr(fn, r), sig.Results().At(i).Type())
			results = append(results, v)
		}
	}

	// Store the results.
	for i, r := range results {
		var result Value // fn.source.result[i] conceptually
		if fn == fn.source {
			result = fn.results[i]
		} else { // lookup needed?
			result = fn.lookup(fn.returnVars[i], false)
		}
		emitStore(fn, result, r, s.Return)
	}

	if fn.jump != nil {
		// Return from body of a range-over-func.
		// The return statement is syntactically within the loop,
		// but the generated code is in the 'switch jump {...}' after it.
		e := returnExit(fn, s.Pos())
		storeVar(fn, fn.jump, intConst(e.id), e.pos)
		fn.emit(&Return{Results: []Value{vFalse}, pos: e.pos})
		fn.currentBlock = fn.newBasicBlock("unreachable")
		return
	}

	// Run function calls deferred in this
	// function when explicitly returning from it.
	fn.emit(new(RunDefers))
	// Reload (potentially) named result variables to form the result tuple.
	results = results[:0]
	for _, nr := range fn.results {
		results = append(results, emitLoad(fn, nr))
	}
	fn.emit(&Return{Results: results, pos: s.Return})
	fn.currentBlock = fn.newBasicBlock("unreachable")
}

// A buildFunc is a strategy for building the SSA body for a function.
type buildFunc = func(*builder, *Function)

// iterate causes all created but unbuilt functions to be built. As
// this may create new methods, the process is iterated until it
// converges.
//
// Waits for any dependencies to finish building.
func (b *builder) iterate() {
	for ; b.finished < len(b.fns); b.finished++ {
		fn := b.fns[b.finished]
		b.buildFunction(fn)
	}

	b.buildshared.markDone()
	b.buildshared.wait()
}

// buildFunction builds SSA code for the body of function fn.  Idempotent.
func (b *builder) buildFunction(fn *Function) {
	if fn.build != nil {
		assert(fn.parent == nil, "anonymous functions should not be built by buildFunction()")

		if fn.Prog.mode&LogSource != 0 {
			defer logStack("build %s @ %s", fn, fn.Prog.Fset.Position(fn.pos))()
		}
		fn.build(b, fn)
		fn.done()
	}
}

// buildParamsOnly builds fn.Params from fn.Signature, but does not build fn.Body.
func (b *builder) buildParamsOnly(fn *Function) {
	// For external (C, asm) functions or functions loaded from
	// export data, we must set fn.Params even though there is no
	// body code to reference them.
	if recv := fn.Signature.Recv(); recv != nil {
		fn.addParamVar(recv)
	}
	params := fn.Signature.Params()
	for i, n := 0, params.Len(); i < n; i++ {
		fn.addParamVar(params.At(i))
	}
}

// buildFromSyntax builds fn.Body from fn.syntax, which must be non-nil.
func (b *builder) buildFromSyntax(fn *Function) {
	var (
		recvField *ast.FieldList
		body      *ast.BlockStmt
		functype  *ast.FuncType
	)
	switch syntax := fn.syntax.(type) {
	case *ast.FuncDecl:
		functype = syntax.Type
		recvField = syntax.Recv
		body = syntax.Body
		if body == nil {
			b.buildParamsOnly(fn) // no body (non-Go function)
			return
		}
	case *ast.FuncLit:
		functype = syntax.Type
		body = syntax.Body
	case nil:
		panic("no syntax")
	default:
		panic(syntax) // unexpected syntax
	}
	fn.source = fn
	fn.startBody()
	fn.createSyntacticParams(recvField, functype)
	fn.createDeferStack()
	b.stmt(fn, body)
	if cb := fn.currentBlock; cb != nil && (cb == fn.Blocks[0] || cb == fn.Recover || cb.Preds != nil) {
		// Control fell off the end of the function's body block.
		//
		// Block optimizations eliminate the current block, if
		// unreachable.  It is a builder invariant that
		// if this no-arg return is ill-typed for
		// fn.Signature.Results, this block must be
		// unreachable.  The sanity checker checks this.
		fn.emit(new(RunDefers))
		fn.emit(new(Return))
	}
	fn.finishBody()
}

// buildYieldFunc builds the body of the yield function created
// from a range-over-func *ast.RangeStmt.
func (b *builder) buildYieldFunc(fn *Function) {
	// See builder.rangeFunc for detailed documentation on how fn is set up.
	//
	// In pseudo-Go this roughly builds:
	// func yield(_k tk, _v tv) bool {
	// 	   if jump != READY { panic("yield function called after range loop exit") }
	//     jump = BUSY
	//     k, v = _k, _v // assign the iterator variable (if needed)
	//     ... // rng.Body
	//   continue:
	//     jump = READY
	//     return true
	// }
	s := fn.syntax.(*ast.RangeStmt)
	fn.source = fn.parent.source
	fn.startBody()
	params := fn.Signature.Params()
	for i := 0; i < params.Len(); i++ {
		fn.addParamVar(params.At(i))
	}

	// Initial targets
	ycont := fn.newBasicBlock("yield-continue")
	// lblocks is either {} or is {label: nil} where label is the label of syntax.
	for label := range fn.lblocks {
		fn.lblocks[label] = &lblock{
			label:     label,
			resolved:  true,
			_goto:     ycont,
			_continue: ycont,
			// `break label` statement targets fn.parent.targets._break
		}
	}
	fn.targets = &targets{
		tail:      fn.targets,
		_continue: ycont,
		// `break` statement targets fn.parent.targets._break.
	}

	// continue:
	//   jump = READY
	//   return true
	saved := fn.currentBlock
	fn.currentBlock = ycont
	storeVar(fn, fn.jump, jReady, s.Body.Rbrace)
	// A yield function's own deferstack is always empty, so rundefers is not needed.
	fn.emit(&Return{Results: []Value{vTrue}, pos: token.NoPos})

	// Emit header:
	//
	//   if jump != READY { panic("yield iterator accessed after exit") }
	//   jump = BUSY
	//   k, v = _k, _v
	fn.currentBlock = saved
	yloop := fn.newBasicBlock("yield-loop")
	invalid := fn.newBasicBlock("yield-invalid")

	jumpVal := emitLoad(fn, fn.lookup(fn.jump, true))
	emitIf(fn, emitCompare(fn, token.EQL, jumpVal, jReady, token.NoPos), yloop, invalid)
	fn.currentBlock = invalid
	fn.emit(&Panic{
		X: emitConv(fn, stringConst("yield function called after range loop exit"), tEface),
	})

	fn.currentBlock = yloop
	storeVar(fn, fn.jump, jBusy, s.Body.Rbrace)

	// Initialize k and v from params.
	var tk, tv types.Type
	if s.Key != nil && !isBlankIdent(s.Key) {
		tk = fn.typeOf(s.Key) // fn.parent.typeOf is identical
	}
	if s.Value != nil && !isBlankIdent(s.Value) {
		tv = fn.typeOf(s.Value)
	}
	if s.Tok == token.DEFINE {
		if tk != nil {
			emitLocalVar(fn, identVar(fn, s.Key.(*ast.Ident)))
		}
		if tv != nil {
			emitLocalVar(fn, identVar(fn, s.Value.(*ast.Ident)))
		}
	}
	var k, v Value
	if len(fn.Params) > 0 {
		k = fn.Params[0]
	}
	if len(fn.Params) > 1 {
		v = fn.Params[1]
	}
	var kl, vl lvalue
	if tk != nil {
		kl = b.addr(fn, s.Key, false) // non-escaping
	}
	if tv != nil {
		vl = b.addr(fn, s.Value, false) // non-escaping
	}
	if tk != nil {
		kl.store(fn, k)
	}
	if tv != nil {
		vl.store(fn, v)
	}

	// Build the body of the range loop.
	b.stmt(fn, s.Body)
	if cb := fn.currentBlock; cb != nil && (cb == fn.Blocks[0] || cb == fn.Recover || cb.Preds != nil) {
		// Control fell off the end of the function's body block.
		// Block optimizations eliminate the current block, if
		// unreachable.
		emitJump(fn, ycont)
	}
	// pop the stack for the yield function
	fn.targets = fn.targets.tail

	// Clean up exits and promote any unresolved exits to fn.parent.
	for _, e := range fn.exits {
		if e.label != nil {
			lb := fn.lblocks[e.label]
			if lb.resolved {
				// label was resolved. Do not turn lb into an exit.
				// e does not need to be handled by the parent.
				continue
			}

			// _goto becomes an exit.
			//   _goto:
			//     jump = id
			//     return false
			fn.currentBlock = lb._goto
			id := intConst(e.id)
			storeVar(fn, fn.jump, id, e.pos)
			fn.emit(&Return{Results: []Value{vFalse}, pos: e.pos})
		}

		if e.to != fn { // e needs to be handled by the parent too.
			fn.parent.exits = append(fn.parent.exits, e)
		}
	}

	fn.finishBody()
}

// addMakeInterfaceType records non-interface type t as the type of
// the operand a MakeInterface operation, for [Program.RuntimeTypes].
//
// Acquires prog.makeInterfaceTypesMu.
func addMakeInterfaceType(prog *Program, t types.Type) {
	prog.makeInterfaceTypesMu.Lock()
	defer prog.makeInterfaceTypesMu.Unlock()
	if prog.makeInterfaceTypes == nil {
		prog.makeInterfaceTypes = make(map[types.Type]unit)
	}
	prog.makeInterfaceTypes[t] = unit{}
}

// Build calls Package.Build for each package in prog.
// Building occurs in parallel unless the BuildSerially mode flag was set.
//
// Build is intended for whole-program analysis; a typical compiler
// need only build a single package.
//
// Build is idempotent and thread-safe.
func (prog *Program) Build() {
	var wg sync.WaitGroup
	for _, p := range prog.packages {
		if prog.mode&BuildSerially != 0 {
			p.Build()
		} else {
			wg.Add(1)
			cpuLimit <- unit{} // acquire a token
			go func(p *Package) {
				p.Build()
				wg.Done()
				<-cpuLimit // release a token
			}(p)
		}
	}
	wg.Wait()
}

// cpuLimit is a counting semaphore to limit CPU parallelism.
var cpuLimit = make(chan unit, runtime.GOMAXPROCS(0))

// Build builds SSA code for all functions and vars in package p.
//
// CreatePackage must have been called for all of p's direct imports
// (and hence its direct imports must have been error-free). It is not
// necessary to call CreatePackage for indirect dependencies.
// Functions will be created for all necessary methods in those
// packages on demand.
//
// Build is idempotent and thread-safe.
func (p *Package) Build() { p.buildOnce.Do(p.build) }

func (p *Package) build() {
	if p.info == nil {
		return // synthetic package, e.g. "testmain"
	}
	if p.Prog.mode&LogSource != 0 {
		defer logStack("build %s", p)()
	}

	b := builder{fns: p.created}
	b.iterate()

	// We no longer need transient information: ASTs or go/types deductions.
	p.info = nil
	p.created = nil
	p.files = nil
	p.initVersion = nil

	if p.Prog.mode&SanityCheckFunctions != 0 {
		sanityCheckPackage(p)
	}
}

// buildPackageInit builds fn.Body for the synthetic package initializer.
func (b *builder) buildPackageInit(fn *Function) {
	p := fn.Pkg
	fn.startBody()

	var done *BasicBlock

	if p.Prog.mode&BareInits == 0 {
		// Make init() skip if package is already initialized.
		initguard := p.Var("init$guard")
		doinit := fn.newBasicBlock("init.start")
		done = fn.newBasicBlock("init.done")
		emitIf(fn, emitLoad(fn, initguard), done, doinit)
		fn.currentBlock = doinit
		emitStore(fn, initguard, vTrue, token.NoPos)

		// Call the init() function of each package we import.
		for _, pkg := range p.Pkg.Imports() {
			prereq := p.Prog.packages[pkg]
			if prereq == nil {
				panic(fmt.Sprintf("Package(%q).Build(): unsatisfied import: Program.CreatePackage(%q) was not called", p.Pkg.Path(), pkg.Path()))
			}
			var v Call
			v.Call.Value = prereq.init
			v.Call.pos = fn.pos
			v.setType(types.NewTuple())
			fn.emit(&v)
		}
	}

	// Initialize package-level vars in correct order.
	if len(p.info.InitOrder) > 0 && len(p.files) == 0 {
		panic("no source files provided for package. cannot initialize globals")
	}

	for _, varinit := range p.info.InitOrder {
		if fn.Prog.mode&LogSource != 0 {
			fmt.Fprintf(os.Stderr, "build global initializer %v @ %s\n",
				varinit.Lhs, p.Prog.Fset.Position(varinit.Rhs.Pos()))
		}
		// Initializers for global vars are evaluated in dependency
		// order, but may come from arbitrary files of the package
		// with different versions, so we transiently update
		// fn.goversion for each one. (Since init is a synthetic
		// function it has no syntax of its own that needs a version.)
		fn.goversion = p.initVersion[varinit.Rhs]
		if len(varinit.Lhs) == 1 {
			// 1:1 initialization: var x, y = a(), b()
			var lval lvalue
			if v := varinit.Lhs[0]; v.Name() != "_" {
				lval = &address{addr: p.objects[v].(*Global), pos: v.Pos()}
			} else {
				lval = blank{}
			}
			b.assign(fn, lval, varinit.Rhs, true, nil)
		} else {
			// n:1 initialization: var x, y :=  f()
			tuple := b.exprN(fn, varinit.Rhs)
			for i, v := range varinit.Lhs {
				if v.Name() == "_" {
					continue
				}
				emitStore(fn, p.objects[v].(*Global), emitExtract(fn, tuple, i), v.Pos())
			}
		}
	}

	// The rest of the init function is synthetic:
	// no syntax, info, goversion.
	fn.info = nil
	fn.goversion = ""

	// Call all of the declared init() functions in source order.
	for _, file := range p.files {
		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				id := decl.Name
				if !isBlankIdent(id) && id.Name == "init" && decl.Recv == nil {
					declaredInit := p.objects[p.info.Defs[id]].(*Function)
					var v Call
					v.Call.Value = declaredInit
					v.setType(types.NewTuple())
					p.init.emit(&v)
				}
			}
		}
	}

	// Finish up init().
	if p.Prog.mode&BareInits == 0 {
		emitJump(fn, done)
		fn.currentBlock = done
	}
	fn.emit(new(Return))
	fn.finishBody()
}
```

## File: go/ssa/const_test.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"math/big"
	"strings"
	"testing"

	"golang.org/x/tools/go/ssa"
)

func TestConstString(t *testing.T) {
	const source = `
	package P

	type Named string

	func fn() (int, bool, string) 
	func gen[T int]() {}
	`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}

	var conf types.Config
	pkg, err := conf.Check("P", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range []struct {
		expr     string // type expression
		constant any    // constant value
		want     string // expected String() value
	}{
		{"int", int64(0), "0:int"},
		{"int64", int64(0), "0:int64"},
		{"float32", int64(0), "0:float32"},
		{"float32", big.NewFloat(1.5), "1.5:float32"},
		{"bool", false, "false:bool"},
		{"string", "", `"":string`},
		{"Named", "", `"":P.Named`},
		{"struct{x string}", nil, "struct{x string}{}:struct{x string}"},
		{"[]int", nil, "nil:[]int"},
		{"[3]int", nil, "[3]int{}:[3]int"},
		{"*int", nil, "nil:*int"},
		{"interface{}", nil, "nil:interface{}"},
		{"interface{string}", nil, `"":interface{string}`},
		{"interface{int|int64}", nil, "0:interface{int|int64}"},
		{"interface{bool}", nil, "false:interface{bool}"},
		{"interface{bool|int}", nil, "invalid:interface{bool|int}"},
		{"interface{int|string}", nil, "invalid:interface{int|string}"},
		{"interface{bool|string}", nil, "invalid:interface{bool|string}"},
		{"interface{struct{x string}}", nil, "invalid:interface{struct{x string}}"},
		{"interface{int|int64}", int64(1), "1:interface{int|int64}"},
		{"interface{~bool}", true, "true:interface{~bool}"},
		{"interface{Named}", "lorem ipsum", `"lorem ipsum":interface{P.Named}`},
		{"func() (int, bool, string)", nil, "nil:func() (int, bool, string)"},
	} {
		// Eval() expr for its type.
		tv, err := types.Eval(fset, pkg, 0, test.expr)
		if err != nil {
			t.Fatalf("Eval(%s) failed: %v", test.expr, err)
		}
		var val constant.Value
		if test.constant != nil {
			val = constant.Make(test.constant)
		}
		c := ssa.NewConst(val, tv.Type)
		got := strings.ReplaceAll(c.String(), " | ", "|") // Accept both interface{a | b} and interface{a|b}.
		if got != test.want {
			t.Errorf("ssa.NewConst(%v, %s).String() = %v, want %v", val, tv.Type, got, test.want)
		}
	}

	// Test tuples
	fn := pkg.Scope().Lookup("fn")
	tup := fn.Type().(*types.Signature).Results()
	if got, want := ssa.NewConst(nil, tup).String(), `(0, false, ""):(int, bool, string)`; got != want {
		t.Errorf("ssa.NewConst(%v, %s).String() = %v, want %v", nil, tup, got, want)
	}

	// Test type-param
	gen := pkg.Scope().Lookup("gen")
	tp := gen.Type().(*types.Signature).TypeParams().At(0)
	if got, want := ssa.NewConst(nil, tp).String(), "0:T"; got != want {
		t.Errorf("ssa.NewConst(%v, %s).String() = %v, want %v", nil, tup, got, want)
	}
}
```

## File: go/ssa/const.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines the Const SSA value type.

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/internal/typeparams"
	"golang.org/x/tools/internal/typesinternal"
)

// NewConst returns a new constant of the specified value and type.
// val must be valid according to the specification of Const.Value.
func NewConst(val constant.Value, typ types.Type) *Const {
	if val == nil {
		switch soleTypeKind(typ) {
		case types.IsBoolean:
			val = constant.MakeBool(false)
		case types.IsInteger:
			val = constant.MakeInt64(0)
		case types.IsString:
			val = constant.MakeString("")
		}
	}
	return &Const{typ, val}
}

// soleTypeKind returns a BasicInfo for which constant.Value can
// represent all zero values for the types in the type set.
//
//	types.IsBoolean for false is a representative.
//	types.IsInteger for 0
//	types.IsString for ""
//	0 otherwise.
func soleTypeKind(typ types.Type) types.BasicInfo {
	// State records the set of possible zero values (false, 0, "").
	// Candidates (perhaps all) are eliminated during the type-set
	// iteration, which executes at least once.
	state := types.IsBoolean | types.IsInteger | types.IsString
	underIs(typ, func(ut types.Type) bool {
		var c types.BasicInfo
		if t, ok := ut.(*types.Basic); ok {
			c = t.Info()
		}
		if c&types.IsNumeric != 0 { // int/float/complex
			c = types.IsInteger
		}
		state = state & c
		return state != 0
	})
	return state
}

// intConst returns an 'int' constant that evaluates to i.
// (i is an int64 in case the host is narrower than the target.)
func intConst(i int64) *Const {
	return NewConst(constant.MakeInt64(i), tInt)
}

// stringConst returns a 'string' constant that evaluates to s.
func stringConst(s string) *Const {
	return NewConst(constant.MakeString(s), tString)
}

// zeroConst returns a new "zero" constant of the specified type.
func zeroConst(t types.Type) *Const {
	return NewConst(nil, t)
}

func (c *Const) RelString(from *types.Package) string {
	var s string
	if c.Value == nil {
		s, _ = typesinternal.ZeroString(c.typ, types.RelativeTo(from))
	} else if c.Value.Kind() == constant.String {
		s = constant.StringVal(c.Value)
		const max = 20
		// TODO(adonovan): don't cut a rune in half.
		if len(s) > max {
			s = s[:max-3] + "..." // abbreviate
		}
		s = strconv.Quote(s)
	} else {
		s = c.Value.String()
	}
	return s + ":" + relType(c.Type(), from)
}

func (c *Const) Name() string {
	return c.RelString(nil)
}

func (c *Const) String() string {
	return c.Name()
}

func (c *Const) Type() types.Type {
	return c.typ
}

func (c *Const) Referrers() *[]Instruction {
	return nil
}

func (c *Const) Parent() *Function { return nil }

func (c *Const) Pos() token.Pos {
	return token.NoPos
}

// IsNil returns true if this constant is a nil value of
// a nillable reference type (pointer, slice, channel, map, or function),
// a basic interface type, or
// a type parameter all of whose possible instantiations are themselves nillable.
func (c *Const) IsNil() bool {
	return c.Value == nil && nillable(c.typ)
}

// nillable reports whether *new(T) == nil is legal for type T.
func nillable(t types.Type) bool {
	if typeparams.IsTypeParam(t) {
		return underIs(t, func(u types.Type) bool {
			// empty type set (u==nil) => any underlying types => not nillable
			return u != nil && nillable(u)
		})
	}
	switch t.Underlying().(type) {
	case *types.Pointer, *types.Slice, *types.Chan, *types.Map, *types.Signature:
		return true
	case *types.Interface:
		return true // basic interface.
	default:
		return false
	}
}

// TODO(adonovan): move everything below into golang.org/x/tools/go/ssa/interp.

// Int64 returns the numeric value of this constant truncated to fit
// a signed 64-bit integer.
func (c *Const) Int64() int64 {
	switch x := constant.ToInt(c.Value); x.Kind() {
	case constant.Int:
		if i, ok := constant.Int64Val(x); ok {
			return i
		}
		return 0
	case constant.Float:
		f, _ := constant.Float64Val(x)
		return int64(f)
	}
	panic(fmt.Sprintf("unexpected constant value: %T", c.Value))
}

// Uint64 returns the numeric value of this constant truncated to fit
// an unsigned 64-bit integer.
func (c *Const) Uint64() uint64 {
	switch x := constant.ToInt(c.Value); x.Kind() {
	case constant.Int:
		if u, ok := constant.Uint64Val(x); ok {
			return u
		}
		return 0
	case constant.Float:
		f, _ := constant.Float64Val(x)
		return uint64(f)
	}
	panic(fmt.Sprintf("unexpected constant value: %T", c.Value))
}

// Float64 returns the numeric value of this constant truncated to fit
// a float64.
func (c *Const) Float64() float64 {
	x := constant.ToFloat(c.Value) // (c.Value == nil) => x.Kind() == Unknown
	f, _ := constant.Float64Val(x)
	return f
}

// Complex128 returns the complex value of this constant truncated to
// fit a complex128.
func (c *Const) Complex128() complex128 {
	x := constant.ToComplex(c.Value) // (c.Value == nil) => x.Kind() == Unknown
	re, _ := constant.Float64Val(constant.Real(x))
	im, _ := constant.Float64Val(constant.Imag(x))
	return complex(re, im)
}
```

## File: go/ssa/create.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file implements the CREATE phase of SSA construction.
// See builder.go for explanation.

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"sync"

	"golang.org/x/tools/internal/versions"
)

// NewProgram returns a new SSA Program.
//
// mode controls diagnostics and checking during SSA construction.
//
// To construct an SSA program:
//
//   - Call NewProgram to create an empty Program.
//   - Call CreatePackage providing typed syntax for each package
//     you want to build, and call it with types but not
//     syntax for each of those package's direct dependencies.
//   - Call [Package.Build] on each syntax package you wish to build,
//     or [Program.Build] to build all of them.
//
// See the Example tests for simple examples.
func NewProgram(fset *token.FileSet, mode BuilderMode) *Program {
	return &Program{
		Fset:     fset,
		imported: make(map[string]*Package),
		packages: make(map[*types.Package]*Package),
		mode:     mode,
		canon:    newCanonizer(),
		ctxt:     types.NewContext(),
	}
}

// memberFromObject populates package pkg with a member for the
// typechecker object obj.
//
// For objects from Go source code, syntax is the associated syntax
// tree (for funcs and vars only) and goversion defines the
// appropriate interpretation; they will be used during the build
// phase.
func memberFromObject(pkg *Package, obj types.Object, syntax ast.Node, goversion string) {
	name := obj.Name()
	switch obj := obj.(type) {
	case *types.Builtin:
		if pkg.Pkg != types.Unsafe {
			panic("unexpected builtin object: " + obj.String())
		}

	case *types.TypeName:
		if name != "_" {
			pkg.Members[name] = &Type{
				object: obj,
				pkg:    pkg,
			}
		}

	case *types.Const:
		c := &NamedConst{
			object: obj,
			Value:  NewConst(obj.Val(), obj.Type()),
			pkg:    pkg,
		}
		pkg.objects[obj] = c
		if name != "_" {
			pkg.Members[name] = c
		}

	case *types.Var:
		g := &Global{
			Pkg:    pkg,
			name:   name,
			object: obj,
			typ:    types.NewPointer(obj.Type()), // address
			pos:    obj.Pos(),
		}
		pkg.objects[obj] = g
		if name != "_" {
			pkg.Members[name] = g
		}

	case *types.Func:
		sig := obj.Type().(*types.Signature)
		if sig.Recv() == nil && name == "init" {
			pkg.ninit++
			name = fmt.Sprintf("init#%d", pkg.ninit)
		}
		fn := createFunction(pkg.Prog, obj, name, syntax, pkg.info, goversion)
		fn.Pkg = pkg
		pkg.created = append(pkg.created, fn)
		pkg.objects[obj] = fn
		if name != "_" && sig.Recv() == nil {
			pkg.Members[name] = fn // package-level function
		}

	default: // (incl. *types.Package)
		panic("unexpected Object type: " + obj.String())
	}
}

// createFunction creates a function or method. It supports both
// CreatePackage (with or without syntax) and the on-demand creation
// of methods in non-created packages based on their types.Func.
func createFunction(prog *Program, obj *types.Func, name string, syntax ast.Node, info *types.Info, goversion string) *Function {
	sig := obj.Type().(*types.Signature)

	// Collect type parameters.
	var tparams *types.TypeParamList
	if rtparams := sig.RecvTypeParams(); rtparams.Len() > 0 {
		tparams = rtparams // method of generic type
	} else if sigparams := sig.TypeParams(); sigparams.Len() > 0 {
		tparams = sigparams // generic function
	}

	/* declared function/method (from syntax or export data) */
	fn := &Function{
		name:       name,
		object:     obj,
		Signature:  sig,
		build:      (*builder).buildFromSyntax,
		syntax:     syntax,
		info:       info,
		goversion:  goversion,
		pos:        obj.Pos(),
		Pkg:        nil, // may be set by caller
		Prog:       prog,
		typeparams: tparams,
	}
	if fn.syntax == nil {
		fn.Synthetic = "from type information"
		fn.build = (*builder).buildParamsOnly
	}
	if tparams.Len() > 0 {
		fn.generic = new(generic)
	}
	return fn
}

// membersFromDecl populates package pkg with members for each
// typechecker object (var, func, const or type) associated with the
// specified decl.
func membersFromDecl(pkg *Package, decl ast.Decl, goversion string) {
	switch decl := decl.(type) {
	case *ast.GenDecl: // import, const, type or var
		switch decl.Tok {
		case token.CONST:
			for _, spec := range decl.Specs {
				for _, id := range spec.(*ast.ValueSpec).Names {
					memberFromObject(pkg, pkg.info.Defs[id], nil, "")
				}
			}

		case token.VAR:
			for _, spec := range decl.Specs {
				for _, rhs := range spec.(*ast.ValueSpec).Values {
					pkg.initVersion[rhs] = goversion
				}
				for _, id := range spec.(*ast.ValueSpec).Names {
					memberFromObject(pkg, pkg.info.Defs[id], spec, goversion)
				}
			}

		case token.TYPE:
			for _, spec := range decl.Specs {
				id := spec.(*ast.TypeSpec).Name
				memberFromObject(pkg, pkg.info.Defs[id], nil, "")
			}
		}

	case *ast.FuncDecl:
		id := decl.Name
		memberFromObject(pkg, pkg.info.Defs[id], decl, goversion)
	}
}

// CreatePackage creates and returns an SSA Package from the
// specified type-checked, error-free file ASTs, and populates its
// Members mapping.
//
// importable determines whether this package should be returned by a
// subsequent call to ImportedPackage(pkg.Path()).
//
// The real work of building SSA form for each function is not done
// until a subsequent call to Package.Build.
func (prog *Program) CreatePackage(pkg *types.Package, files []*ast.File, info *types.Info, importable bool) *Package {
	if pkg == nil {
		panic("nil pkg") // otherwise pkg.Scope below returns types.Universe!
	}
	p := &Package{
		Prog:    prog,
		Members: make(map[string]Member),
		objects: make(map[types.Object]Member),
		Pkg:     pkg,
		syntax:  info != nil,
		// transient values (cleared after Package.Build)
		info:        info,
		files:       files,
		initVersion: make(map[ast.Expr]string),
	}

	/* synthesized package initializer */
	p.init = &Function{
		name:      "init",
		Signature: new(types.Signature),
		Synthetic: "package initializer",
		Pkg:       p,
		Prog:      prog,
		build:     (*builder).buildPackageInit,
		info:      p.info,
		goversion: "", // See Package.build for details.
	}
	p.Members[p.init.name] = p.init
	p.created = append(p.created, p.init)

	// Allocate all package members: vars, funcs, consts and types.
	if len(files) > 0 {
		// Go source package.
		for _, file := range files {
			goversion := versions.Lang(versions.FileVersion(p.info, file))
			for _, decl := range file.Decls {
				membersFromDecl(p, decl, goversion)
			}
		}
	} else {
		// GC-compiled binary package (or "unsafe")
		// No code.
		// No position information.
		scope := p.Pkg.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			memberFromObject(p, obj, nil, "")
			if obj, ok := obj.(*types.TypeName); ok {
				// No Unalias: aliases should not duplicate methods.
				if named, ok := obj.Type().(*types.Named); ok {
					for i, n := 0, named.NumMethods(); i < n; i++ {
						memberFromObject(p, named.Method(i), nil, "")
					}
				}
			}
		}
	}

	if prog.mode&BareInits == 0 {
		// Add initializer guard variable.
		initguard := &Global{
			Pkg:  p,
			name: "init$guard",
			typ:  types.NewPointer(tBool),
		}
		p.Members[initguard.Name()] = initguard
	}

	if prog.mode&GlobalDebug != 0 {
		p.SetDebugMode(true)
	}

	if prog.mode&PrintPackages != 0 {
		printMu.Lock()
		p.WriteTo(os.Stdout)
		printMu.Unlock()
	}

	if importable {
		prog.imported[p.Pkg.Path()] = p
	}
	prog.packages[p.Pkg] = p

	return p
}

// printMu serializes printing of Packages/Functions to stdout.
var printMu sync.Mutex

// AllPackages returns a new slice containing all packages created by
// prog.CreatePackage in unspecified order.
func (prog *Program) AllPackages() []*Package {
	pkgs := make([]*Package, 0, len(prog.packages))
	for _, pkg := range prog.packages {
		pkgs = append(pkgs, pkg)
	}
	return pkgs
}

// ImportedPackage returns the importable Package whose PkgPath
// is path, or nil if no such Package has been created.
//
// A parameter to CreatePackage determines whether a package should be
// considered importable. For example, no import declaration can resolve
// to the ad-hoc main package created by 'go build foo.go'.
//
// TODO(adonovan): rethink this function and the "importable" concept;
// most packages are importable. This function assumes that all
// types.Package.Path values are unique within the ssa.Program, which is
// false---yet this function remains very convenient.
// Clients should use (*Program).Package instead where possible.
// SSA doesn't really need a string-keyed map of packages.
//
// Furthermore, the graph of packages may contain multiple variants
// (e.g. "p" vs "p as compiled for q.test"), and each has a different
// view of its dependencies.
func (prog *Program) ImportedPackage(path string) *Package {
	return prog.imported[path]
}
```

## File: go/ssa/doc.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ssa defines a representation of the elements of Go programs
// (packages, types, functions, variables and constants) using a
// static single-assignment (SSA) form intermediate representation
// (IR) for the bodies of functions.
//
// For an introduction to SSA form, see
// http://en.wikipedia.org/wiki/Static_single_assignment_form.
// This page provides a broader reading list:
// http://www.dcs.gla.ac.uk/~jsinger/ssa.html.
//
// The level of abstraction of the SSA form is intentionally close to
// the source language to facilitate construction of source analysis
// tools.  It is not intended for machine code generation.
//
// All looping, branching and switching constructs are replaced with
// unstructured control flow.  Higher-level control flow constructs
// such as multi-way branch can be reconstructed as needed; see
// [golang.org/x/tools/go/ssa/ssautil.Switches] for an example.
//
// The simplest way to create the SSA representation of a package is
// to load typed syntax trees using [golang.org/x/tools/go/packages], then
// invoke the [golang.org/x/tools/go/ssa/ssautil.Packages] helper function.
// (See the package-level Examples named LoadPackages and LoadWholeProgram.)
// The resulting [ssa.Program] contains all the packages and their
// members, but SSA code is not created for function bodies until a
// subsequent call to [Package.Build] or [Program.Build].
//
// The builder initially builds a naive SSA form in which all local
// variables are addresses of stack locations with explicit loads and
// stores.  Registerisation of eligible locals and φ-node insertion
// using dominance and dataflow are then performed as a second pass
// called "lifting" to improve the accuracy and performance of
// subsequent analyses; this pass can be skipped by setting the
// NaiveForm builder flag.
//
// The primary interfaces of this package are:
//
//   - [Member]: a named member of a Go package.
//   - [Value]: an expression that yields a value.
//   - [Instruction]: a statement that consumes values and performs computation.
//   - [Node]: a [Value] or [Instruction] (emphasizing its membership in the SSA value graph)
//
// A computation that yields a result implements both the [Value] and
// [Instruction] interfaces.  The following table shows for each
// concrete type which of these interfaces it implements.
//
//	                   Value?          Instruction?      Member?
//	*Alloc                ✔               ✔
//	*BinOp                ✔               ✔
//	*Builtin              ✔
//	*Call                 ✔               ✔
//	*ChangeInterface      ✔               ✔
//	*ChangeType           ✔               ✔
//	*Const                ✔
//	*Convert              ✔               ✔
//	*DebugRef                             ✔
//	*Defer                                ✔
//	*Extract              ✔               ✔
//	*Field                ✔               ✔
//	*FieldAddr            ✔               ✔
//	*FreeVar              ✔
//	*Function             ✔                               ✔ (func)
//	*Global               ✔                               ✔ (var)
//	*Go                                   ✔
//	*If                                   ✔
//	*Index                ✔               ✔
//	*IndexAddr            ✔               ✔
//	*Jump                                 ✔
//	*Lookup               ✔               ✔
//	*MakeChan             ✔               ✔
//	*MakeClosure          ✔               ✔
//	*MakeInterface        ✔               ✔
//	*MakeMap              ✔               ✔
//	*MakeSlice            ✔               ✔
//	*MapUpdate                            ✔
//	*MultiConvert         ✔               ✔
//	*NamedConst                                           ✔ (const)
//	*Next                 ✔               ✔
//	*Panic                                ✔
//	*Parameter            ✔
//	*Phi                  ✔               ✔
//	*Range                ✔               ✔
//	*Return                               ✔
//	*RunDefers                            ✔
//	*Select               ✔               ✔
//	*Send                                 ✔
//	*Slice                ✔               ✔
//	*SliceToArrayPointer  ✔               ✔
//	*Store                                ✔
//	*Type                                                 ✔ (type)
//	*TypeAssert           ✔               ✔
//	*UnOp                 ✔               ✔
//
// Other key types in this package include: [Program], [Package], [Function]
// and [BasicBlock].
//
// The program representation constructed by this package is fully
// resolved internally, i.e. it does not rely on the names of Values,
// Packages, Functions, Types or BasicBlocks for the correct
// interpretation of the program.  Only the identities of objects and
// the topology of the SSA and type graphs are semantically
// significant.  (There is one exception: [types.Id] values, which identify field
// and method names, contain strings.)  Avoidance of name-based
// operations simplifies the implementation of subsequent passes and
// can make them very efficient.  Many objects are nonetheless named
// to aid in debugging, but it is not essential that the names be
// either accurate or unambiguous.  The public API exposes a number of
// name-based maps for client convenience.
//
// The [golang.org/x/tools/go/ssa/ssautil] package provides various
// helper functions, for example to simplify loading a Go program into
// SSA form.
//
// TODO(adonovan): write a how-to document for all the various cases
// of trying to determine corresponding elements across the four
// domains of source locations, ast.Nodes, types.Objects,
// ssa.Values/Instructions.
package ssa // import "golang.org/x/tools/go/ssa"
```

## File: go/ssa/dom_test.go
```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testenv"
)

func TestDominatorOrder(t *testing.T) {
	testenv.NeedsGoBuild(t) // for go/packages

	const src = `package p

func f(cond bool) {
	// (Print operands match BasicBlock IDs.)
	print(0)
	if cond {
		print(1)
	} else {
		print(2)
	}
	print(3)
}
`
	dir := t.TempDir()
	cfg := &packages.Config{
		Dir:  dir,
		Mode: packages.LoadSyntax,
		Overlay: map[string][]byte{
			filepath.Join(dir, "p.go"): []byte(src),
		},
	}
	initial, err := packages.Load(cfg, "./p.go")
	if err != nil {
		t.Fatal(err)
	}
	if packages.PrintErrors(initial) > 0 {
		t.Fatal("packages contain errors")
	}
	_, pkgs := ssautil.Packages(initial, 0)
	p := pkgs[0]
	p.Build()
	f := p.Func("f")

	if got, want := fmt.Sprint(f.DomPreorder()), "[0 1 2 3]"; got != want {
		t.Errorf("DomPreorder: got %v, want %s", got, want)
	}
	if got, want := fmt.Sprint(f.DomPostorder()), "[1 2 3 0]"; got != want {
		t.Errorf("DomPostorder: got %v, want %s", got, want)
	}
}
```

## File: go/ssa/dom.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines algorithms related to dominance.

// Dominator tree construction ----------------------------------------
//
// We use the algorithm described in Lengauer & Tarjan. 1979.  A fast
// algorithm for finding dominators in a flowgraph.
// http://doi.acm.org/10.1145/357062.357071
//
// We also apply the optimizations to SLT described in Georgiadis et
// al, Finding Dominators in Practice, JGAA 2006,
// http://jgaa.info/accepted/2006/GeorgiadisTarjanWerneck2006.10.1.pdf
// to avoid the need for buckets of size > 1.

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"sort"
)

// Idom returns the block that immediately dominates b:
// its parent in the dominator tree, if any.
// Neither the entry node (b.Index==0) nor recover node
// (b==b.Parent().Recover()) have a parent.
func (b *BasicBlock) Idom() *BasicBlock { return b.dom.idom }

// Dominees returns the list of blocks that b immediately dominates:
// its children in the dominator tree.
func (b *BasicBlock) Dominees() []*BasicBlock { return b.dom.children }

// Dominates reports whether b dominates c.
func (b *BasicBlock) Dominates(c *BasicBlock) bool {
	return b.dom.pre <= c.dom.pre && c.dom.post <= b.dom.post
}

// DomPreorder returns a new slice containing the blocks of f
// in a preorder traversal of the dominator tree.
func (f *Function) DomPreorder() []*BasicBlock {
	slice := append([]*BasicBlock(nil), f.Blocks...)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].dom.pre < slice[j].dom.pre
	})
	return slice
}

// DomPostorder returns a new slice containing the blocks of f
// in a postorder traversal of the dominator tree.
// (This is not the same as a postdominance order.)
func (f *Function) DomPostorder() []*BasicBlock {
	slice := append([]*BasicBlock(nil), f.Blocks...)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].dom.post < slice[j].dom.post
	})
	return slice
}

// domInfo contains a BasicBlock's dominance information.
type domInfo struct {
	idom      *BasicBlock   // immediate dominator (parent in domtree)
	children  []*BasicBlock // nodes immediately dominated by this one
	pre, post int32         // pre- and post-order numbering within domtree
}

// ltState holds the working state for Lengauer-Tarjan algorithm
// (during which domInfo.pre is repurposed for CFG DFS preorder number).
type ltState struct {
	// Each slice is indexed by b.Index.
	sdom     []*BasicBlock // b's semidominator
	parent   []*BasicBlock // b's parent in DFS traversal of CFG
	ancestor []*BasicBlock // b's ancestor with least sdom
}

// dfs implements the depth-first search part of the LT algorithm.
func (lt *ltState) dfs(v *BasicBlock, i int32, preorder []*BasicBlock) int32 {
	preorder[i] = v
	v.dom.pre = i // For now: DFS preorder of spanning tree of CFG
	i++
	lt.sdom[v.Index] = v
	lt.link(nil, v)
	for _, w := range v.Succs {
		if lt.sdom[w.Index] == nil {
			lt.parent[w.Index] = v
			i = lt.dfs(w, i, preorder)
		}
	}
	return i
}

// eval implements the EVAL part of the LT algorithm.
func (lt *ltState) eval(v *BasicBlock) *BasicBlock {
	// TODO(adonovan): opt: do path compression per simple LT.
	u := v
	for ; lt.ancestor[v.Index] != nil; v = lt.ancestor[v.Index] {
		if lt.sdom[v.Index].dom.pre < lt.sdom[u.Index].dom.pre {
			u = v
		}
	}
	return u
}

// link implements the LINK part of the LT algorithm.
func (lt *ltState) link(v, w *BasicBlock) {
	lt.ancestor[w.Index] = v
}

// buildDomTree computes the dominator tree of f using the LT algorithm.
// Precondition: all blocks are reachable (e.g. optimizeBlocks has been run).
func buildDomTree(f *Function) {
	// The step numbers refer to the original LT paper; the
	// reordering is due to Georgiadis.

	// Clear any previous domInfo.
	for _, b := range f.Blocks {
		b.dom = domInfo{}
	}

	n := len(f.Blocks)
	// Allocate space for 5 contiguous [n]*BasicBlock arrays:
	// sdom, parent, ancestor, preorder, buckets.
	space := make([]*BasicBlock, 5*n)
	lt := ltState{
		sdom:     space[0:n],
		parent:   space[n : 2*n],
		ancestor: space[2*n : 3*n],
	}

	// Step 1.  Number vertices by depth-first preorder.
	preorder := space[3*n : 4*n]
	root := f.Blocks[0]
	prenum := lt.dfs(root, 0, preorder)
	recover := f.Recover
	if recover != nil {
		lt.dfs(recover, prenum, preorder)
	}

	buckets := space[4*n : 5*n]
	copy(buckets, preorder)

	// In reverse preorder...
	for i := int32(n) - 1; i > 0; i-- {
		w := preorder[i]

		// Step 3. Implicitly define the immediate dominator of each node.
		for v := buckets[i]; v != w; v = buckets[v.dom.pre] {
			u := lt.eval(v)
			if lt.sdom[u.Index].dom.pre < i {
				v.dom.idom = u
			} else {
				v.dom.idom = w
			}
		}

		// Step 2. Compute the semidominators of all nodes.
		lt.sdom[w.Index] = lt.parent[w.Index]
		for _, v := range w.Preds {
			u := lt.eval(v)
			if lt.sdom[u.Index].dom.pre < lt.sdom[w.Index].dom.pre {
				lt.sdom[w.Index] = lt.sdom[u.Index]
			}
		}

		lt.link(lt.parent[w.Index], w)

		if lt.parent[w.Index] == lt.sdom[w.Index] {
			w.dom.idom = lt.parent[w.Index]
		} else {
			buckets[i] = buckets[lt.sdom[w.Index].dom.pre]
			buckets[lt.sdom[w.Index].dom.pre] = w
		}
	}

	// The final 'Step 3' is now outside the loop.
	for v := buckets[0]; v != root; v = buckets[v.dom.pre] {
		v.dom.idom = root
	}

	// Step 4. Explicitly define the immediate dominator of each
	// node, in preorder.
	for _, w := range preorder[1:] {
		if w == root || w == recover {
			w.dom.idom = nil
		} else {
			if w.dom.idom != lt.sdom[w.Index] {
				w.dom.idom = w.dom.idom.dom.idom
			}
			// Calculate Children relation as inverse of Idom.
			w.dom.idom.dom.children = append(w.dom.idom.dom.children, w)
		}
	}

	pre, post := numberDomTree(root, 0, 0)
	if recover != nil {
		numberDomTree(recover, pre, post)
	}

	// printDomTreeDot(os.Stderr, f)        // debugging
	// printDomTreeText(os.Stderr, root, 0) // debugging

	if f.Prog.mode&SanityCheckFunctions != 0 {
		sanityCheckDomTree(f)
	}
}

// numberDomTree sets the pre- and post-order numbers of a depth-first
// traversal of the dominator tree rooted at v.  These are used to
// answer dominance queries in constant time.
func numberDomTree(v *BasicBlock, pre, post int32) (int32, int32) {
	v.dom.pre = pre
	pre++
	for _, child := range v.dom.children {
		pre, post = numberDomTree(child, pre, post)
	}
	v.dom.post = post
	post++
	return pre, post
}

// Testing utilities ----------------------------------------

// sanityCheckDomTree checks the correctness of the dominator tree
// computed by the LT algorithm by comparing against the dominance
// relation computed by a naive Kildall-style forward dataflow
// analysis (Algorithm 10.16 from the "Dragon" book).
func sanityCheckDomTree(f *Function) {
	n := len(f.Blocks)

	// D[i] is the set of blocks that dominate f.Blocks[i],
	// represented as a bit-set of block indices.
	D := make([]big.Int, n)

	one := big.NewInt(1)

	// all is the set of all blocks; constant.
	var all big.Int
	all.Set(one).Lsh(&all, uint(n)).Sub(&all, one)

	// Initialization.
	for i, b := range f.Blocks {
		if i == 0 || b == f.Recover {
			// A root is dominated only by itself.
			D[i].SetBit(&D[0], 0, 1)
		} else {
			// All other blocks are (initially) dominated
			// by every block.
			D[i].Set(&all)
		}
	}

	// Iteration until fixed point.
	for changed := true; changed; {
		changed = false
		for i, b := range f.Blocks {
			if i == 0 || b == f.Recover {
				continue
			}
			// Compute intersection across predecessors.
			var x big.Int
			x.Set(&all)
			for _, pred := range b.Preds {
				x.And(&x, &D[pred.Index])
			}
			x.SetBit(&x, i, 1) // a block always dominates itself.
			if D[i].Cmp(&x) != 0 {
				D[i].Set(&x)
				changed = true
			}
		}
	}

	// Check the entire relation.  O(n^2).
	// The Recover block (if any) must be treated specially so we skip it.
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b, c := f.Blocks[i], f.Blocks[j]
			if c == f.Recover {
				continue
			}
			actual := b.Dominates(c)
			expected := D[j].Bit(i) == 1
			if actual != expected {
				fmt.Fprintf(os.Stderr, "dominates(%s, %s)==%t, want %t\n", b, c, actual, expected)
				ok = false
			}
		}
	}

	preorder := f.DomPreorder()
	for _, b := range f.Blocks {
		if got := preorder[b.dom.pre]; got != b {
			fmt.Fprintf(os.Stderr, "preorder[%d]==%s, want %s\n", b.dom.pre, got, b)
			ok = false
		}
	}

	if !ok {
		panic("sanityCheckDomTree failed for " + f.String())
	}

}

// Printing functions ----------------------------------------

// printDomTreeText prints the dominator tree as text, using indentation.
func printDomTreeText(buf *bytes.Buffer, v *BasicBlock, indent int) {
	fmt.Fprintf(buf, "%*s%s\n", 4*indent, "", v)
	for _, child := range v.dom.children {
		printDomTreeText(buf, child, indent+1)
	}
}

// printDomTreeDot prints the dominator tree of f in AT&T GraphViz
// (.dot) format.
// (unused; retained for debugging)
func printDomTreeDot(buf *bytes.Buffer, f *Function) {
	fmt.Fprintln(buf, "//", f)
	fmt.Fprintln(buf, "digraph domtree {")
	for i, b := range f.Blocks {
		v := b.dom
		fmt.Fprintf(buf, "\tn%d [label=\"%s (%d, %d)\",shape=\"rectangle\"];\n", v.pre, b, v.pre, v.post)
		// TODO(adonovan): improve appearance of edges
		// belonging to both dominator tree and CFG.

		// Dominator tree edge.
		if i != 0 {
			fmt.Fprintf(buf, "\tn%d -> n%d [style=\"solid\",weight=100];\n", v.idom.dom.pre, v.pre)
		}
		// CFG edges.
		for _, pred := range b.Preds {
			fmt.Fprintf(buf, "\tn%d -> n%d [style=\"dotted\",weight=0];\n", pred.dom.pre, v.pre)
		}
	}
	fmt.Fprintln(buf, "}")
}
```

## File: go/ssa/emit.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// Helpers for emitting SSA instructions.

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/internal/typeparams"
)

// emitAlloc emits to f a new Alloc instruction allocating a variable
// of type typ.
//
// The caller must set Alloc.Heap=true (for a heap-allocated variable)
// or add the Alloc to f.Locals (for a frame-allocated variable).
//
// During building, a variable in f.Locals may have its Heap flag
// set when it is discovered that its address is taken.
// These Allocs are removed from f.Locals at the end.
//
// The builder should generally call one of the emit{New,Local,LocalVar} wrappers instead.
func emitAlloc(f *Function, typ types.Type, pos token.Pos, comment string) *Alloc {
	v := &Alloc{Comment: comment}
	v.setType(types.NewPointer(typ))
	v.setPos(pos)
	f.emit(v)
	return v
}

// emitNew emits to f a new Alloc instruction heap-allocating a
// variable of type typ. pos is the optional source location.
func emitNew(f *Function, typ types.Type, pos token.Pos, comment string) *Alloc {
	alloc := emitAlloc(f, typ, pos, comment)
	alloc.Heap = true
	return alloc
}

// emitLocal creates a local var for (t, pos, comment) and
// emits an Alloc instruction for it.
//
// (Use this function or emitNew for synthetic variables;
// for source-level variables in the same function, use emitLocalVar.)
func emitLocal(f *Function, t types.Type, pos token.Pos, comment string) *Alloc {
	local := emitAlloc(f, t, pos, comment)
	f.Locals = append(f.Locals, local)
	return local
}

// emitLocalVar creates a local var for v and emits an Alloc instruction for it.
// Subsequent calls to f.lookup(v) return it.
// It applies the appropriate generic instantiation to the type.
func emitLocalVar(f *Function, v *types.Var) *Alloc {
	alloc := emitLocal(f, f.typ(v.Type()), v.Pos(), v.Name())
	f.vars[v] = alloc
	return alloc
}

// emitLoad emits to f an instruction to load the address addr into a
// new temporary, and returns the value so defined.
func emitLoad(f *Function, addr Value) *UnOp {
	v := &UnOp{Op: token.MUL, X: addr}
	v.setType(typeparams.MustDeref(addr.Type()))
	f.emit(v)
	return v
}

// emitDebugRef emits to f a DebugRef pseudo-instruction associating
// expression e with value v.
func emitDebugRef(f *Function, e ast.Expr, v Value, isAddr bool) {
	if !f.debugInfo() {
		return // debugging not enabled
	}
	if v == nil || e == nil {
		panic("nil")
	}
	var obj types.Object
	e = ast.Unparen(e)
	if id, ok := e.(*ast.Ident); ok {
		if isBlankIdent(id) {
			return
		}
		obj = f.objectOf(id)
		switch obj.(type) {
		case *types.Nil, *types.Const, *types.Builtin:
			return
		}
	}
	f.emit(&DebugRef{
		X:      v,
		Expr:   e,
		IsAddr: isAddr,
		object: obj,
	})
}

// emitArith emits to f code to compute the binary operation op(x, y)
// where op is an eager shift, logical or arithmetic operation.
// (Use emitCompare() for comparisons and Builder.logicalBinop() for
// non-eager operations.)
func emitArith(f *Function, op token.Token, x, y Value, t types.Type, pos token.Pos) Value {
	switch op {
	case token.SHL, token.SHR:
		x = emitConv(f, x, t)
		// y may be signed or an 'untyped' constant.

		// There is a runtime panic if y is signed and <0. Instead of inserting a check for y<0
		// and converting to an unsigned value (like the compiler) leave y as is.

		if isUntyped(y.Type().Underlying()) {
			// Untyped conversion:
			// Spec https://go.dev/ref/spec#Operators:
			// The right operand in a shift expression must have integer type or be an untyped constant
			// representable by a value of type uint.
			y = emitConv(f, y, types.Typ[types.Uint])
		}

	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.AND, token.OR, token.XOR, token.AND_NOT:
		x = emitConv(f, x, t)
		y = emitConv(f, y, t)

	default:
		panic("illegal op in emitArith: " + op.String())

	}
	v := &BinOp{
		Op: op,
		X:  x,
		Y:  y,
	}
	v.setPos(pos)
	v.setType(t)
	return f.emit(v)
}

// emitCompare emits to f code compute the boolean result of
// comparison 'x op y'.
func emitCompare(f *Function, op token.Token, x, y Value, pos token.Pos) Value {
	xt := x.Type().Underlying()
	yt := y.Type().Underlying()

	// Special case to optimise a tagless SwitchStmt so that
	// these are equivalent
	//   switch { case e: ...}
	//   switch true { case e: ... }
	//   if e==true { ... }
	// even in the case when e's type is an interface.
	// TODO(adonovan): opt: generalise to x==true, false!=y, etc.
	if x == vTrue && op == token.EQL {
		if yt, ok := yt.(*types.Basic); ok && yt.Info()&types.IsBoolean != 0 {
			return y
		}
	}

	if types.Identical(xt, yt) {
		// no conversion necessary
	} else if isNonTypeParamInterface(x.Type()) {
		y = emitConv(f, y, x.Type())
	} else if isNonTypeParamInterface(y.Type()) {
		x = emitConv(f, x, y.Type())
	} else if _, ok := x.(*Const); ok {
		x = emitConv(f, x, y.Type())
	} else if _, ok := y.(*Const); ok {
		y = emitConv(f, y, x.Type())
	} else {
		// other cases, e.g. channels.  No-op.
	}

	v := &BinOp{
		Op: op,
		X:  x,
		Y:  y,
	}
	v.setPos(pos)
	v.setType(tBool)
	return f.emit(v)
}

// isValuePreserving returns true if a conversion from ut_src to
// ut_dst is value-preserving, i.e. just a change of type.
// Precondition: neither argument is a named or alias type.
func isValuePreserving(ut_src, ut_dst types.Type) bool {
	// Identical underlying types?
	if types.IdenticalIgnoreTags(ut_dst, ut_src) {
		return true
	}

	switch ut_dst.(type) {
	case *types.Chan:
		// Conversion between channel types?
		_, ok := ut_src.(*types.Chan)
		return ok

	case *types.Pointer:
		// Conversion between pointers with identical base types?
		_, ok := ut_src.(*types.Pointer)
		return ok
	}
	return false
}

// emitConv emits to f code to convert Value val to exactly type typ,
// and returns the converted value.  Implicit conversions are required
// by language assignability rules in assignments, parameter passing,
// etc.
func emitConv(f *Function, val Value, typ types.Type) Value {
	t_src := val.Type()

	// Identical types?  Conversion is a no-op.
	if types.Identical(t_src, typ) {
		return val
	}
	ut_dst := typ.Underlying()
	ut_src := t_src.Underlying()

	// Conversion to, or construction of a value of, an interface type?
	if isNonTypeParamInterface(typ) {
		// Interface name change?
		if isValuePreserving(ut_src, ut_dst) {
			c := &ChangeType{X: val}
			c.setType(typ)
			return f.emit(c)
		}

		// Assignment from one interface type to another?
		if isNonTypeParamInterface(t_src) {
			c := &ChangeInterface{X: val}
			c.setType(typ)
			return f.emit(c)
		}

		// Untyped nil constant?  Return interface-typed nil constant.
		if ut_src == tUntypedNil {
			return zeroConst(typ)
		}

		// Convert (non-nil) "untyped" literals to their default type.
		if t, ok := ut_src.(*types.Basic); ok && t.Info()&types.IsUntyped != 0 {
			val = emitConv(f, val, types.Default(ut_src))
		}

		// Record the types of operands to MakeInterface, if
		// non-parameterized, as they are the set of runtime types.
		t := val.Type()
		if f.typeparams.Len() == 0 || !f.Prog.isParameterized(t) {
			addMakeInterfaceType(f.Prog, t)
		}

		mi := &MakeInterface{X: val}
		mi.setType(typ)
		return f.emit(mi)
	}

	// conversionCase describes an instruction pattern that maybe emitted to
	// model d <- s for d in dst_terms and s in src_terms.
	// Multiple conversions can match the same pattern.
	type conversionCase uint8
	const (
		changeType conversionCase = 1 << iota
		sliceToArray
		sliceToArrayPtr
		sliceTo0Array
		sliceTo0ArrayPtr
		convert
	)
	// classify the conversion case of a source type us to a destination type ud.
	// us and ud are underlying types (not *Named or *Alias)
	classify := func(us, ud types.Type) conversionCase {
		// Just a change of type, but not value or representation?
		if isValuePreserving(us, ud) {
			return changeType
		}

		// Conversion from slice to array or slice to array pointer?
		if slice, ok := us.(*types.Slice); ok {
			var arr *types.Array
			var ptr bool
			// Conversion from slice to array pointer?
			switch d := ud.(type) {
			case *types.Array:
				arr = d
			case *types.Pointer:
				arr, _ = d.Elem().Underlying().(*types.Array)
				ptr = true
			}
			if arr != nil && types.Identical(slice.Elem(), arr.Elem()) {
				if arr.Len() == 0 {
					if ptr {
						return sliceTo0ArrayPtr
					} else {
						return sliceTo0Array
					}
				}
				if ptr {
					return sliceToArrayPtr
				} else {
					return sliceToArray
				}
			}
		}

		// The only remaining case in well-typed code is a representation-
		// changing conversion of basic types (possibly with []byte/[]rune).
		if !isBasic(us) && !isBasic(ud) {
			panic(fmt.Sprintf("in %s: cannot convert term %s (%s [within %s]) to type %s [within %s]", f, val, val.Type(), us, typ, ud))
		}
		return convert
	}

	var classifications conversionCase
	underIs(ut_src, func(us types.Type) bool {
		return underIs(ut_dst, func(ud types.Type) bool {
			if us != nil && ud != nil {
				classifications |= classify(us, ud)
			}
			return classifications != 0
		})
	})
	if classifications == 0 {
		panic(fmt.Sprintf("in %s: cannot convert %s (%s) to %s", f, val, val.Type(), typ))
	}

	// Conversion of a compile-time constant value?
	if c, ok := val.(*Const); ok {
		// Conversion to a basic type?
		if isBasic(ut_dst) {
			// Conversion of a compile-time constant to
			// another constant type results in a new
			// constant of the destination type and
			// (initially) the same abstract value.
			// We don't truncate the value yet.
			return NewConst(c.Value, typ)
		}
		// Can we always convert from zero value without panicking?
		const mayPanic = sliceToArray | sliceToArrayPtr
		if c.Value == nil && classifications&mayPanic == 0 {
			return NewConst(nil, typ)
		}

		// We're converting from constant to non-constant type,
		// e.g. string -> []byte/[]rune.
	}

	switch classifications {
	case changeType: // representation-preserving change
		c := &ChangeType{X: val}
		c.setType(typ)
		return f.emit(c)

	case sliceToArrayPtr, sliceTo0ArrayPtr: // slice to array pointer
		c := &SliceToArrayPointer{X: val}
		c.setType(typ)
		return f.emit(c)

	case sliceToArray: // slice to arrays (not zero-length)
		ptype := types.NewPointer(typ)
		p := &SliceToArrayPointer{X: val}
		p.setType(ptype)
		x := f.emit(p)
		unOp := &UnOp{Op: token.MUL, X: x}
		unOp.setType(typ)
		return f.emit(unOp)

	case sliceTo0Array: // slice to zero-length arrays (constant)
		return zeroConst(typ)

	case convert: // representation-changing conversion
		c := &Convert{X: val}
		c.setType(typ)
		return f.emit(c)

	default: // The conversion represents a cross product.
		c := &MultiConvert{X: val, from: t_src, to: typ}
		c.setType(typ)
		return f.emit(c)
	}
}

// emitTypeCoercion emits to f code to coerce the type of a
// Value v to exactly type typ, and returns the coerced value.
//
// Requires that coercing v.Typ() to typ is a value preserving change.
//
// Currently used only when v.Type() is a type instance of typ or vice versa.
// A type v is a type instance of a type t if there exists a
// type parameter substitution σ s.t. σ(v) == t. Example:
//
//	σ(func(T) T) == func(int) int for σ == [T ↦ int]
//
// This happens in instantiation wrappers for conversion
// from an instantiation to a parameterized type (and vice versa)
// with σ substituting f.typeparams by f.typeargs.
func emitTypeCoercion(f *Function, v Value, typ types.Type) Value {
	if types.Identical(v.Type(), typ) {
		return v // no coercion needed
	}
	// TODO(taking): for instances should we record which side is the instance?
	c := &ChangeType{
		X: v,
	}
	c.setType(typ)
	f.emit(c)
	return c
}

// emitStore emits to f an instruction to store value val at location
// addr, applying implicit conversions as required by assignability rules.
func emitStore(f *Function, addr, val Value, pos token.Pos) *Store {
	typ := typeparams.MustDeref(addr.Type())
	s := &Store{
		Addr: addr,
		Val:  emitConv(f, val, typ),
		pos:  pos,
	}
	f.emit(s)
	return s
}

// emitJump emits to f a jump to target, and updates the control-flow graph.
// Postcondition: f.currentBlock is nil.
func emitJump(f *Function, target *BasicBlock) {
	b := f.currentBlock
	b.emit(new(Jump))
	addEdge(b, target)
	f.currentBlock = nil
}

// emitIf emits to f a conditional jump to tblock or fblock based on
// cond, and updates the control-flow graph.
// Postcondition: f.currentBlock is nil.
func emitIf(f *Function, cond Value, tblock, fblock *BasicBlock) {
	b := f.currentBlock
	b.emit(&If{Cond: cond})
	addEdge(b, tblock)
	addEdge(b, fblock)
	f.currentBlock = nil
}

// emitExtract emits to f an instruction to extract the index'th
// component of tuple.  It returns the extracted value.
func emitExtract(f *Function, tuple Value, index int) Value {
	e := &Extract{Tuple: tuple, Index: index}
	e.setType(tuple.Type().(*types.Tuple).At(index).Type())
	return f.emit(e)
}

// emitTypeAssert emits to f a type assertion value := x.(t) and
// returns the value.  x.Type() must be an interface.
func emitTypeAssert(f *Function, x Value, t types.Type, pos token.Pos) Value {
	a := &TypeAssert{X: x, AssertedType: t}
	a.setPos(pos)
	a.setType(t)
	return f.emit(a)
}

// emitTypeTest emits to f a type test value,ok := x.(t) and returns
// a (value, ok) tuple.  x.Type() must be an interface.
func emitTypeTest(f *Function, x Value, t types.Type, pos token.Pos) Value {
	a := &TypeAssert{
		X:            x,
		AssertedType: t,
		CommaOk:      true,
	}
	a.setPos(pos)
	a.setType(types.NewTuple(
		newVar("value", t),
		varOk,
	))
	return f.emit(a)
}

// emitTailCall emits to f a function call in tail position.  The
// caller is responsible for all fields of 'call' except its type.
// Intended for wrapper methods.
// Precondition: f does/will not use deferred procedure calls.
// Postcondition: f.currentBlock is nil.
func emitTailCall(f *Function, call *Call) {
	tresults := f.Signature.Results()
	nr := tresults.Len()
	if nr == 1 {
		call.typ = tresults.At(0).Type()
	} else {
		call.typ = tresults
	}
	tuple := f.emit(call)
	var ret Return
	switch nr {
	case 0:
		// no-op
	case 1:
		ret.Results = []Value{tuple}
	default:
		for i := 0; i < nr; i++ {
			v := emitExtract(f, tuple, i)
			// TODO(adonovan): in principle, this is required:
			//   v = emitConv(f, o.Type, f.Signature.Results[i].Type)
			// but in practice emitTailCall is only used when
			// the types exactly match.
			ret.Results = append(ret.Results, v)
		}
	}
	f.emit(&ret)
	f.currentBlock = nil
}

// emitImplicitSelections emits to f code to apply the sequence of
// implicit field selections specified by indices to base value v, and
// returns the selected value.
//
// If v is the address of a struct, the result will be the address of
// a field; if it is the value of a struct, the result will be the
// value of a field.
func emitImplicitSelections(f *Function, v Value, indices []int, pos token.Pos) Value {
	for _, index := range indices {
		if isPointerCore(v.Type()) {
			fld := fieldOf(typeparams.MustDeref(v.Type()), index)
			instr := &FieldAddr{
				X:     v,
				Field: index,
			}
			instr.setPos(pos)
			instr.setType(types.NewPointer(fld.Type()))
			v = f.emit(instr)
			// Load the field's value iff indirectly embedded.
			if isPointerCore(fld.Type()) {
				v = emitLoad(f, v)
			}
		} else {
			fld := fieldOf(v.Type(), index)
			instr := &Field{
				X:     v,
				Field: index,
			}
			instr.setPos(pos)
			instr.setType(fld.Type())
			v = f.emit(instr)
		}
	}
	return v
}

// emitFieldSelection emits to f code to select the index'th field of v.
//
// If wantAddr, the input must be a pointer-to-struct and the result
// will be the field's address; otherwise the result will be the
// field's value.
// Ident id is used for position and debug info.
func emitFieldSelection(f *Function, v Value, index int, wantAddr bool, id *ast.Ident) Value {
	if isPointerCore(v.Type()) {
		fld := fieldOf(typeparams.MustDeref(v.Type()), index)
		instr := &FieldAddr{
			X:     v,
			Field: index,
		}
		instr.setPos(id.Pos())
		instr.setType(types.NewPointer(fld.Type()))
		v = f.emit(instr)
		// Load the field's value iff we don't want its address.
		if !wantAddr {
			v = emitLoad(f, v)
		}
	} else {
		fld := fieldOf(v.Type(), index)
		instr := &Field{
			X:     v,
			Field: index,
		}
		instr.setPos(id.Pos())
		instr.setType(fld.Type())
		v = f.emit(instr)
	}
	emitDebugRef(f, id, v, wantAddr)
	return v
}

// createRecoverBlock emits to f a block of code to return after a
// recovered panic, and sets f.Recover to it.
//
// If f's result parameters are named, the code loads and returns
// their current values, otherwise it returns the zero values of their
// type.
//
// Idempotent.
func createRecoverBlock(f *Function) {
	if f.Recover != nil {
		return // already created
	}
	saved := f.currentBlock

	f.Recover = f.newBasicBlock("recover")
	f.currentBlock = f.Recover

	var results []Value
	// Reload NRPs to form value tuple.
	for _, nr := range f.results {
		results = append(results, emitLoad(f, nr))
	}

	f.emit(&Return{Results: results})

	f.currentBlock = saved
}
```

## File: go/ssa/example_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !android && !ios && (unix || aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || plan9 || windows)

package ssa_test

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const hello = `
package main

import "fmt"

const message = "Hello, World!"

func main() {
	fmt.Println(message)
}
`

// This program demonstrates how to run the SSA builder on a single
// package of one or more already-parsed files. Its dependencies are
// loaded from compiler export data. This is what you'd typically use
// for a compiler; it does not depend on the obsolete
// [golang.org/x/tools/go/loader].
//
// It shows the printed representation of packages, functions, and
// instructions.  Within the function listing, the name of each
// BasicBlock such as ".0.entry" is printed left-aligned, followed by
// the block's Instructions.
//
// For each instruction that defines an SSA virtual register
// (i.e. implements Value), the type of that value is shown in the
// right column.
//
// Build and run the ssadump.go program if you want a standalone tool
// with similar functionality. It is located at
// [golang.org/x/tools/cmd/ssadump].
//
// Use ssautil.BuildPackage only if you have parsed--but not
// type-checked--syntax trees. Typically, clients already have typed
// syntax, perhaps obtained from golang.org/x/tools/go/packages.
// In that case, see the other examples for simpler approaches.
func Example_buildPackage() {
	// Parse the source files.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
	if err != nil {
		fmt.Print(err) // parse error
		return
	}
	files := []*ast.File{f}

	// Create the type-checker's package.
	pkg := types.NewPackage("hello", "")

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	hello, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Print(err) // type error in some package
		return
	}

	// Print out the package.
	hello.WriteTo(os.Stdout)

	// Print out the package-level functions.
	hello.Func("init").WriteTo(os.Stdout)
	hello.Func("main").WriteTo(os.Stdout)

	// Output:
	//
	// package hello:
	//   func  init       func()
	//   var   init$guard bool
	//   func  main       func()
	//   const message    message = "Hello, World!":untyped string
	//
	// # Name: hello.init
	// # Package: hello
	// # Synthetic: package initializer
	// func init():
	// 0:                                                                entry P:0 S:2
	// 	t0 = *init$guard                                                   bool
	// 	if t0 goto 2 else 1
	// 1:                                                           init.start P:1 S:1
	// 	*init$guard = true:bool
	// 	t1 = fmt.init()                                                      ()
	// 	jump 2
	// 2:                                                            init.done P:2 S:0
	// 	return
	//
	// # Name: hello.main
	// # Package: hello
	// # Location: hello.go:8:6
	// func main():
	// 0:                                                                entry P:0 S:0
	// 	t0 = new [1]any (varargs)                                       *[1]any
	// 	t1 = &t0[0:int]                                                    *any
	// 	t2 = make any <- string ("Hello, World!":string)                    any
	// 	*t1 = t2
	// 	t3 = slice t0[:]                                                  []any
	// 	t4 = fmt.Println(t3...)                              (n int, err error)
	// 	return
}

// This example builds SSA code for a set of packages using the
// [golang.org/x/tools/go/packages] API. This is what you would typically use for a
// analysis capable of operating on a single package.
func Example_loadPackages() {
	// Load, parse, and type-check the initial packages.
	cfg := &packages.Config{Mode: packages.LoadSyntax}
	initial, err := packages.Load(cfg, "fmt", "net/http")
	if err != nil {
		log.Fatal(err)
	}

	// Stop if any package had errors.
	// This step is optional; without it, the next step
	// will create SSA for only a subset of packages.
	if packages.PrintErrors(initial) > 0 {
		log.Fatalf("packages contain errors")
	}

	// Create SSA packages for all well-typed packages.
	prog, pkgs := ssautil.Packages(initial, ssa.PrintPackages)
	_ = prog

	// Build SSA code for the well-typed initial packages.
	for _, p := range pkgs {
		if p != nil {
			p.Build()
		}
	}
}

// This example builds SSA code for a set of packages plus all their dependencies,
// using the [golang.org/x/tools/go/packages] API.
// This is what you'd typically use for a whole-program analysis.
func Example_loadWholeProgram() {
	// Load, parse, and type-check the whole program.
	cfg := packages.Config{Mode: packages.LoadAllSyntax}
	initial, err := packages.Load(&cfg, "fmt", "net/http")
	if err != nil {
		log.Fatal(err)
	}

	// Create SSA packages for well-typed packages and their dependencies.
	prog, pkgs := ssautil.AllPackages(initial, ssa.PrintPackages|ssa.InstantiateGenerics)
	_ = pkgs

	// Build SSA code for the whole program.
	prog.Build()
}
```

## File: go/ssa/func.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file implements the Function type.

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"iter"
	"os"
	"strings"

	"golang.org/x/tools/internal/typeparams"
)

// Like ObjectOf, but panics instead of returning nil.
// Only valid during f's create and build phases.
func (f *Function) objectOf(id *ast.Ident) types.Object {
	if o := f.info.ObjectOf(id); o != nil {
		return o
	}
	panic(fmt.Sprintf("no types.Object for ast.Ident %s @ %s",
		id.Name, f.Prog.Fset.Position(id.Pos())))
}

// Like TypeOf, but panics instead of returning nil.
// Only valid during f's create and build phases.
func (f *Function) typeOf(e ast.Expr) types.Type {
	if T := f.info.TypeOf(e); T != nil {
		return f.typ(T)
	}
	panic(fmt.Sprintf("no type for %T @ %s", e, f.Prog.Fset.Position(e.Pos())))
}

// typ is the locally instantiated type of T.
// If f is not an instantiation, then f.typ(T)==T.
func (f *Function) typ(T types.Type) types.Type {
	return f.subst.typ(T)
}

// If id is an Instance, returns info.Instances[id].Type.
// Otherwise returns f.typeOf(id).
func (f *Function) instanceType(id *ast.Ident) types.Type {
	if t, ok := f.info.Instances[id]; ok {
		return t.Type
	}
	return f.typeOf(id)
}

// selection returns a *selection corresponding to f.info.Selections[selector]
// with potential updates for type substitution.
func (f *Function) selection(selector *ast.SelectorExpr) *selection {
	sel := f.info.Selections[selector]
	if sel == nil {
		return nil
	}

	switch sel.Kind() {
	case types.MethodExpr, types.MethodVal:
		if recv := f.typ(sel.Recv()); recv != sel.Recv() {
			// recv changed during type substitution.
			pkg := f.declaredPackage().Pkg
			obj, index, indirect := types.LookupFieldOrMethod(recv, true, pkg, sel.Obj().Name())

			// sig replaces sel.Type(). See (types.Selection).Typ() for details.
			sig := obj.Type().(*types.Signature)
			sig = changeRecv(sig, newVar(sig.Recv().Name(), recv))
			if sel.Kind() == types.MethodExpr {
				sig = recvAsFirstArg(sig)
			}
			return &selection{
				kind:     sel.Kind(),
				recv:     recv,
				typ:      sig,
				obj:      obj,
				index:    index,
				indirect: indirect,
			}
		}
	}
	return toSelection(sel)
}

// Destinations associated with unlabelled for/switch/select stmts.
// We push/pop one of these as we enter/leave each construct and for
// each BranchStmt we scan for the innermost target of the right type.
type targets struct {
	tail         *targets // rest of stack
	_break       *BasicBlock
	_continue    *BasicBlock
	_fallthrough *BasicBlock
}

// Destinations associated with a labelled block.
// We populate these as labels are encountered in forward gotos or
// labelled statements.
// Forward gotos are resolved once it is known which statement they
// are associated with inside the Function.
type lblock struct {
	label     *types.Label // Label targeted by the blocks.
	resolved  bool         // _goto block encountered (back jump or resolved fwd jump)
	_goto     *BasicBlock
	_break    *BasicBlock
	_continue *BasicBlock
}

// label returns the symbol denoted by a label identifier.
//
// label should be a non-blank identifier (label.Name != "_").
func (f *Function) label(label *ast.Ident) *types.Label {
	return f.objectOf(label).(*types.Label)
}

// lblockOf returns the branch target associated with the
// specified label, creating it if needed.
func (f *Function) lblockOf(label *types.Label) *lblock {
	lb := f.lblocks[label]
	if lb == nil {
		lb = &lblock{
			label: label,
			_goto: f.newBasicBlock(label.Name()),
		}
		if f.lblocks == nil {
			f.lblocks = make(map[*types.Label]*lblock)
		}
		f.lblocks[label] = lb
	}
	return lb
}

// labelledBlock searches f for the block of the specified label.
//
// If f is a yield function, it additionally searches ancestor Functions
// corresponding to enclosing range-over-func statements within the
// same source function, so the returned block may belong to a different Function.
func labelledBlock(f *Function, label *types.Label, tok token.Token) *BasicBlock {
	if lb := f.lblocks[label]; lb != nil {
		var block *BasicBlock
		switch tok {
		case token.BREAK:
			block = lb._break
		case token.CONTINUE:
			block = lb._continue
		case token.GOTO:
			block = lb._goto
		}
		if block != nil {
			return block
		}
	}
	// Search ancestors if this is a yield function.
	if f.jump != nil {
		return labelledBlock(f.parent, label, tok)
	}
	return nil
}

// targetedBlock looks for the nearest block in f.targets
// (and f's ancestors) that matches tok's type, and returns
// the block and function it was found in.
func targetedBlock(f *Function, tok token.Token) *BasicBlock {
	if f == nil {
		return nil
	}
	for t := f.targets; t != nil; t = t.tail {
		var block *BasicBlock
		switch tok {
		case token.BREAK:
			block = t._break
		case token.CONTINUE:
			block = t._continue
		case token.FALLTHROUGH:
			block = t._fallthrough
		}
		if block != nil {
			return block
		}
	}
	// Search f's ancestors (in case f is a yield function).
	return targetedBlock(f.parent, tok)
}

// instrs returns an iterator that returns each reachable instruction of the SSA function.
func (f *Function) instrs() iter.Seq[Instruction] {
	return func(yield func(i Instruction) bool) {
		for _, block := range f.Blocks {
			for _, instr := range block.Instrs {
				if !yield(instr) {
					return
				}
			}
		}
	}
}

// addResultVar adds a result for a variable v to f.results and v to f.returnVars.
func (f *Function) addResultVar(v *types.Var) {
	result := emitLocalVar(f, v)
	f.results = append(f.results, result)
	f.returnVars = append(f.returnVars, v)
}

// addParamVar adds a parameter to f.Params.
func (f *Function) addParamVar(v *types.Var) *Parameter {
	name := v.Name()
	if name == "" {
		name = fmt.Sprintf("arg%d", len(f.Params))
	}
	param := &Parameter{
		name:   name,
		object: v,
		typ:    f.typ(v.Type()),
		parent: f,
	}
	f.Params = append(f.Params, param)
	return param
}

// addSpilledParam declares a parameter that is pre-spilled to the
// stack; the function body will load/store the spilled location.
// Subsequent lifting will eliminate spills where possible.
func (f *Function) addSpilledParam(obj *types.Var) {
	param := f.addParamVar(obj)
	spill := emitLocalVar(f, obj)
	f.emit(&Store{Addr: spill, Val: param})
}

// startBody initializes the function prior to generating SSA code for its body.
// Precondition: f.Type() already set.
func (f *Function) startBody() {
	f.currentBlock = f.newBasicBlock("entry")
	f.vars = make(map[*types.Var]Value) // needed for some synthetics, e.g. init
}

// createSyntacticParams populates f.Params and generates code (spills
// and named result locals) for all the parameters declared in the
// syntax.  In addition it populates the f.objects mapping.
//
// Preconditions:
// f.startBody() was called. f.info != nil.
// Postcondition:
// len(f.Params) == len(f.Signature.Params) + (f.Signature.Recv() ? 1 : 0)
func (f *Function) createSyntacticParams(recv *ast.FieldList, functype *ast.FuncType) {
	// Receiver (at most one inner iteration).
	if recv != nil {
		for _, field := range recv.List {
			for _, n := range field.Names {
				f.addSpilledParam(identVar(f, n))
			}
			// Anonymous receiver?  No need to spill.
			if field.Names == nil {
				f.addParamVar(f.Signature.Recv())
			}
		}
	}

	// Parameters.
	if functype.Params != nil {
		n := len(f.Params) // 1 if has recv, 0 otherwise
		for _, field := range functype.Params.List {
			for _, n := range field.Names {
				f.addSpilledParam(identVar(f, n))
			}
			// Anonymous parameter?  No need to spill.
			if field.Names == nil {
				f.addParamVar(f.Signature.Params().At(len(f.Params) - n))
			}
		}
	}

	// Results.
	if functype.Results != nil {
		for _, field := range functype.Results.List {
			// Implicit "var" decl of locals for named results.
			for _, n := range field.Names {
				v := identVar(f, n)
				f.addResultVar(v)
			}
			// Implicit "var" decl of local for an unnamed result.
			if field.Names == nil {
				v := f.Signature.Results().At(len(f.results))
				f.addResultVar(v)
			}
		}
	}
}

// createDeferStack initializes fn.deferstack to local variable
// initialized to a ssa:deferstack() call.
func (fn *Function) createDeferStack() {
	// Each syntactic function makes a call to ssa:deferstack,
	// which is spilled to a local. Unused ones are later removed.
	fn.deferstack = newVar("defer$stack", tDeferStack)
	call := &Call{Call: CallCommon{Value: vDeferStack}}
	call.setType(tDeferStack)
	deferstack := fn.emit(call)
	spill := emitLocalVar(fn, fn.deferstack)
	emitStore(fn, spill, deferstack, token.NoPos)
}

type setNumable interface {
	setNum(int)
}

// numberRegisters assigns numbers to all SSA registers
// (value-defining Instructions) in f, to aid debugging.
// (Non-Instruction Values are named at construction.)
func numberRegisters(f *Function) {
	v := 0
	for _, b := range f.Blocks {
		for _, instr := range b.Instrs {
			switch instr.(type) {
			case Value:
				instr.(setNumable).setNum(v)
				v++
			}
		}
	}
}

// buildReferrers populates the def/use information in all non-nil
// Value.Referrers slice.
// Precondition: all such slices are initially empty.
func buildReferrers(f *Function) {
	var rands []*Value
	for _, b := range f.Blocks {
		for _, instr := range b.Instrs {
			rands = instr.Operands(rands[:0]) // recycle storage
			for _, rand := range rands {
				if r := *rand; r != nil {
					if ref := r.Referrers(); ref != nil {
						*ref = append(*ref, instr)
					}
				}
			}
		}
	}
}

// finishBody() finalizes the contents of the function after SSA code generation of its body.
//
// The function is not done being built until done() is called.
func (f *Function) finishBody() {
	f.currentBlock = nil
	f.lblocks = nil
	f.returnVars = nil
	f.jump = nil
	f.source = nil
	f.exits = nil

	// Remove from f.Locals any Allocs that escape to the heap.
	j := 0
	for _, l := range f.Locals {
		if !l.Heap {
			f.Locals[j] = l
			j++
		}
	}
	// Nil out f.Locals[j:] to aid GC.
	for i := j; i < len(f.Locals); i++ {
		f.Locals[i] = nil
	}
	f.Locals = f.Locals[:j]

	optimizeBlocks(f)

	buildReferrers(f)

	buildDomTree(f)

	if f.Prog.mode&NaiveForm == 0 {
		// For debugging pre-state of lifting pass:
		// numberRegisters(f)
		// f.WriteTo(os.Stderr)
		lift(f)
	}

	// clear remaining builder state
	f.results = nil    // (used by lifting)
	f.deferstack = nil // (used by lifting)
	f.vars = nil       // (used by lifting)
	f.subst = nil

	numberRegisters(f) // uses f.namedRegisters
}

// done marks the building of f's SSA body complete,
// along with any nested functions, and optionally prints them.
func (f *Function) done() {
	assert(f.parent == nil, "done called on an anonymous function")

	var visit func(*Function)
	visit = func(f *Function) {
		for _, anon := range f.AnonFuncs {
			visit(anon) // anon is done building before f.
		}

		f.uniq = 0    // done with uniq
		f.build = nil // function is built

		if f.Prog.mode&PrintFunctions != 0 {
			printMu.Lock()
			f.WriteTo(os.Stdout)
			printMu.Unlock()
		}

		if f.Prog.mode&SanityCheckFunctions != 0 {
			mustSanityCheck(f, nil)
		}
	}
	visit(f)
}

// removeNilBlocks eliminates nils from f.Blocks and updates each
// BasicBlock.Index.  Use this after any pass that may delete blocks.
func (f *Function) removeNilBlocks() {
	j := 0
	for _, b := range f.Blocks {
		if b != nil {
			b.Index = j
			f.Blocks[j] = b
			j++
		}
	}
	// Nil out f.Blocks[j:] to aid GC.
	for i := j; i < len(f.Blocks); i++ {
		f.Blocks[i] = nil
	}
	f.Blocks = f.Blocks[:j]
}

// SetDebugMode sets the debug mode for package pkg.  If true, all its
// functions will include full debug info.  This greatly increases the
// size of the instruction stream, and causes Functions to depend upon
// the ASTs, potentially keeping them live in memory for longer.
func (pkg *Package) SetDebugMode(debug bool) {
	pkg.debug = debug
}

// debugInfo reports whether debug info is wanted for this function.
func (f *Function) debugInfo() bool {
	// debug info for instantiations follows the debug info of their origin.
	p := f.declaredPackage()
	return p != nil && p.debug
}

// lookup returns the address of the named variable identified by obj
// that is local to function f or one of its enclosing functions.
// If escaping, the reference comes from a potentially escaping pointer
// expression and the referent must be heap-allocated.
// We assume the referent is a *Alloc or *Phi.
// (The only Phis at this stage are those created directly by go1.22 "for" loops.)
func (f *Function) lookup(obj *types.Var, escaping bool) Value {
	if v, ok := f.vars[obj]; ok {
		if escaping {
			switch v := v.(type) {
			case *Alloc:
				v.Heap = true
			case *Phi:
				for _, edge := range v.Edges {
					if alloc, ok := edge.(*Alloc); ok {
						alloc.Heap = true
					}
				}
			}
		}
		return v // function-local var (address)
	}

	// Definition must be in an enclosing function;
	// plumb it through intervening closures.
	if f.parent == nil {
		panic("no ssa.Value for " + obj.String())
	}
	outer := f.parent.lookup(obj, true) // escaping
	v := &FreeVar{
		name:   obj.Name(),
		typ:    outer.Type(),
		pos:    outer.Pos(),
		outer:  outer,
		parent: f,
	}
	f.vars[obj] = v
	f.FreeVars = append(f.FreeVars, v)
	return v
}

// emit emits the specified instruction to function f.
func (f *Function) emit(instr Instruction) Value {
	return f.currentBlock.emit(instr)
}

// RelString returns the full name of this function, qualified by
// package name, receiver type, etc.
//
// The specific formatting rules are not guaranteed and may change.
//
// Examples:
//
//	"math.IsNaN"                  // a package-level function
//	"(*bytes.Buffer).Bytes"       // a declared method or a wrapper
//	"(*bytes.Buffer).Bytes$thunk" // thunk (func wrapping method; receiver is param 0)
//	"(*bytes.Buffer).Bytes$bound" // bound (func wrapping method; receiver supplied by closure)
//	"main.main$1"                 // an anonymous function in main
//	"main.init#1"                 // a declared init function
//	"main.init"                   // the synthesized package initializer
//
// When these functions are referred to from within the same package
// (i.e. from == f.Pkg.Object), they are rendered without the package path.
// For example: "IsNaN", "(*Buffer).Bytes", etc.
//
// All non-synthetic functions have distinct package-qualified names.
// (But two methods may have the same name "(T).f" if one is a synthetic
// wrapper promoting a non-exported method "f" from another package; in
// that case, the strings are equal but the identifiers "f" are distinct.)
func (f *Function) RelString(from *types.Package) string {
	// Anonymous?
	if f.parent != nil {
		// An anonymous function's Name() looks like "parentName$1",
		// but its String() should include the type/package/etc.
		parent := f.parent.RelString(from)
		for i, anon := range f.parent.AnonFuncs {
			if anon == f {
				return fmt.Sprintf("%s$%d", parent, 1+i)
			}
		}

		return f.name // should never happen
	}

	// Method (declared or wrapper)?
	if recv := f.Signature.Recv(); recv != nil {
		return f.relMethod(from, recv.Type())
	}

	// Thunk?
	if f.method != nil {
		return f.relMethod(from, f.method.recv)
	}

	// Bound?
	if len(f.FreeVars) == 1 && strings.HasSuffix(f.name, "$bound") {
		return f.relMethod(from, f.FreeVars[0].Type())
	}

	// Package-level function?
	// Prefix with package name for cross-package references only.
	if p := f.relPkg(); p != nil && p != from {
		return fmt.Sprintf("%s.%s", p.Path(), f.name)
	}

	// Unknown.
	return f.name
}

func (f *Function) relMethod(from *types.Package, recv types.Type) string {
	return fmt.Sprintf("(%s).%s", relType(recv, from), f.name)
}

// writeSignature writes to buf the signature sig in declaration syntax.
func writeSignature(buf *bytes.Buffer, from *types.Package, name string, sig *types.Signature) {
	buf.WriteString("func ")
	if recv := sig.Recv(); recv != nil {
		buf.WriteString("(")
		if name := recv.Name(); name != "" {
			buf.WriteString(name)
			buf.WriteString(" ")
		}
		types.WriteType(buf, recv.Type(), types.RelativeTo(from))
		buf.WriteString(") ")
	}
	buf.WriteString(name)
	types.WriteSignature(buf, sig, types.RelativeTo(from))
}

// declaredPackage returns the package fn is declared in or nil if the
// function is not declared in a package.
func (fn *Function) declaredPackage() *Package {
	switch {
	case fn.Pkg != nil:
		return fn.Pkg // non-generic function  (does that follow??)
	case fn.topLevelOrigin != nil:
		return fn.topLevelOrigin.Pkg // instance of a named generic function
	case fn.parent != nil:
		return fn.parent.declaredPackage() // instance of an anonymous [generic] function
	default:
		return nil // function is not declared in a package, e.g. a wrapper.
	}
}

// relPkg returns types.Package fn is printed in relationship to.
func (fn *Function) relPkg() *types.Package {
	if p := fn.declaredPackage(); p != nil {
		return p.Pkg
	}
	return nil
}

var _ io.WriterTo = (*Function)(nil) // *Function implements io.Writer

func (f *Function) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	WriteFunction(&buf, f)
	n, err := w.Write(buf.Bytes())
	return int64(n), err
}

// WriteFunction writes to buf a human-readable "disassembly" of f.
func WriteFunction(buf *bytes.Buffer, f *Function) {
	fmt.Fprintf(buf, "# Name: %s\n", f.String())
	if f.Pkg != nil {
		fmt.Fprintf(buf, "# Package: %s\n", f.Pkg.Pkg.Path())
	}
	if syn := f.Synthetic; syn != "" {
		fmt.Fprintln(buf, "# Synthetic:", syn)
	}
	if pos := f.Pos(); pos.IsValid() {
		fmt.Fprintf(buf, "# Location: %s\n", f.Prog.Fset.Position(pos))
	}

	if f.parent != nil {
		fmt.Fprintf(buf, "# Parent: %s\n", f.parent.Name())
	}

	if f.Recover != nil {
		fmt.Fprintf(buf, "# Recover: %s\n", f.Recover)
	}

	from := f.relPkg()

	if f.FreeVars != nil {
		buf.WriteString("# Free variables:\n")
		for i, fv := range f.FreeVars {
			fmt.Fprintf(buf, "# % 3d:\t%s %s\n", i, fv.Name(), relType(fv.Type(), from))
		}
	}

	if len(f.Locals) > 0 {
		buf.WriteString("# Locals:\n")
		for i, l := range f.Locals {
			fmt.Fprintf(buf, "# % 3d:\t%s %s\n", i, l.Name(), relType(typeparams.MustDeref(l.Type()), from))
		}
	}
	writeSignature(buf, from, f.Name(), f.Signature)
	buf.WriteString(":\n")

	if f.Blocks == nil {
		buf.WriteString("\t(external)\n")
	}

	// NB. column calculations are confused by non-ASCII
	// characters and assume 8-space tabs.
	const punchcard = 80 // for old time's sake.
	const tabwidth = 8
	for _, b := range f.Blocks {
		if b == nil {
			// Corrupt CFG.
			fmt.Fprintf(buf, ".nil:\n")
			continue
		}
		n, _ := fmt.Fprintf(buf, "%d:", b.Index)
		bmsg := fmt.Sprintf("%s P:%d S:%d", b.Comment, len(b.Preds), len(b.Succs))
		fmt.Fprintf(buf, "%*s%s\n", punchcard-1-n-len(bmsg), "", bmsg)

		if false { // CFG debugging
			fmt.Fprintf(buf, "\t# CFG: %s --> %s --> %s\n", b.Preds, b, b.Succs)
		}
		for _, instr := range b.Instrs {
			buf.WriteString("\t")
			switch v := instr.(type) {
			case Value:
				l := punchcard - tabwidth
				// Left-align the instruction.
				if name := v.Name(); name != "" {
					n, _ := fmt.Fprintf(buf, "%s = ", name)
					l -= n
				}
				n, _ := buf.WriteString(instr.String())
				l -= n
				// Right-align the type if there's space.
				if t := v.Type(); t != nil {
					buf.WriteByte(' ')
					ts := relType(t, from)
					l -= len(ts) + len("  ") // (spaces before and after type)
					if l > 0 {
						fmt.Fprintf(buf, "%*s", l, "")
					}
					buf.WriteString(ts)
				}
			case nil:
				// Be robust against bad transforms.
				buf.WriteString("<deleted>")
			default:
				buf.WriteString(instr.String())
			}
			// -mode=S: show line numbers
			if f.Prog.mode&LogSource != 0 {
				if pos := instr.Pos(); pos.IsValid() {
					fmt.Fprintf(buf, " L%d", f.Prog.Fset.Position(pos).Line)
				}
			}
			buf.WriteString("\n")
		}
	}
	fmt.Fprintf(buf, "\n")
}

// newBasicBlock adds to f a new basic block and returns it.  It does
// not automatically become the current block for subsequent calls to emit.
// comment is an optional string for more readable debugging output.
func (f *Function) newBasicBlock(comment string) *BasicBlock {
	b := &BasicBlock{
		Index:   len(f.Blocks),
		Comment: comment,
		parent:  f,
	}
	b.Succs = b.succs2[:0]
	f.Blocks = append(f.Blocks, b)
	return b
}

// NewFunction returns a new synthetic Function instance belonging to
// prog, with its name and signature fields set as specified.
//
// The caller is responsible for initializing the remaining fields of
// the function object, e.g. Pkg, Params, Blocks.
//
// It is practically impossible for clients to construct well-formed
// SSA functions/packages/programs directly, so we assume this is the
// job of the Builder alone.  NewFunction exists to provide clients a
// little flexibility.  For example, analysis tools may wish to
// construct fake Functions for the root of the callgraph, a fake
// "reflect" package, etc.
//
// TODO(adonovan): think harder about the API here.
func (prog *Program) NewFunction(name string, sig *types.Signature, provenance string) *Function {
	return &Function{Prog: prog, name: name, Signature: sig, Synthetic: provenance}
}

// Syntax returns the function's syntax (*ast.Func{Decl,Lit})
// if it was produced from syntax or an *ast.RangeStmt if
// it is a range-over-func yield function.
func (f *Function) Syntax() ast.Node { return f.syntax }

// identVar returns the variable defined by id.
func identVar(fn *Function, id *ast.Ident) *types.Var {
	return fn.info.Defs[id].(*types.Var)
}

// unique returns a unique positive int within the source tree of f.
// The source tree of f includes all of f's ancestors by parent and all
// of the AnonFuncs contained within these.
func unique(f *Function) int64 {
	f.uniq++
	return f.uniq
}

// exit is a change of control flow going from a range-over-func
// yield function to an ancestor function caused by a break, continue,
// goto, or return statement.
//
// There are 3 types of exits:
// * return from the source function (from ReturnStmt),
// * jump to a block (from break and continue statements [labelled/unlabelled]),
// * go to a label (from goto statements).
//
// As the builder does one pass over the ast, it is unclear whether
// a forward goto statement will leave a range-over-func body.
// The function being exited to is unresolved until the end
// of building the range-over-func body.
type exit struct {
	id   int64     // unique value for exit within from and to
	from *Function // the function the exit starts from
	to   *Function // the function being exited to (nil if unresolved)
	pos  token.Pos

	block *BasicBlock  // basic block within to being jumped to.
	label *types.Label // forward label being jumped to via goto.
	// block == nil && label == nil => return
}

// storeVar emits to function f code to store a value v to a *types.Var x.
func storeVar(f *Function, x *types.Var, v Value, pos token.Pos) {
	emitStore(f, f.lookup(x, true), v, pos)
}

// labelExit creates a new exit to a yield fn to exit the function using a label.
func labelExit(fn *Function, label *types.Label, pos token.Pos) *exit {
	e := &exit{
		id:    unique(fn),
		from:  fn,
		to:    nil,
		pos:   pos,
		label: label,
	}
	fn.exits = append(fn.exits, e)
	return e
}

// blockExit creates a new exit to a yield fn that jumps to a basic block.
func blockExit(fn *Function, block *BasicBlock, pos token.Pos) *exit {
	e := &exit{
		id:    unique(fn),
		from:  fn,
		to:    block.parent,
		pos:   pos,
		block: block,
	}
	fn.exits = append(fn.exits, e)
	return e
}

// blockExit creates a new exit to a yield fn that returns the source function.
func returnExit(fn *Function, pos token.Pos) *exit {
	e := &exit{
		id:   unique(fn),
		from: fn,
		to:   fn.source,
		pos:  pos,
	}
	fn.exits = append(fn.exits, e)
	return e
}
```

## File: go/ssa/instantiate_test.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"fmt"
	"go/types"
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// TestNeedsInstance ensures that new method instances can be created via MethodValue.
func TestNeedsInstance(t *testing.T) {
	const input = `
package p

import "unsafe"

type Pointer[T any] struct {
	v unsafe.Pointer
}

func (x *Pointer[T]) Load() *T {
	return (*T)(LoadPointer(&x.v))
}

func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
`
	// The SSA members for this package should look something like this:
	//      func  LoadPointer func(addr *unsafe.Pointer) (val unsafe.Pointer)
	//      type  Pointer     struct{v unsafe.Pointer}
	//        method (*Pointer[T any]) Load() *T
	//      func  init        func()
	//      var   init$guard  bool

	for _, mode := range []ssa.BuilderMode{
		ssa.SanityCheckFunctions,
		ssa.SanityCheckFunctions | ssa.InstantiateGenerics,
	} {
		p, _ := buildPackage(t, input, mode)
		prog := p.Prog

		ptr := p.Type("Pointer").Type().(*types.Named)
		if ptr.NumMethods() != 1 {
			t.Fatalf("Expected Pointer to have 1 method. got %d", ptr.NumMethods())
		}

		obj := ptr.Method(0)
		if obj.Name() != "Load" {
			t.Errorf("Expected Pointer to have method named 'Load'. got %q", obj.Name())
		}

		meth := prog.FuncValue(obj)

		// instantiateLoadMethod returns the first method (Load) of the instantiation *Pointer[T].
		instantiateLoadMethod := func(T types.Type) *ssa.Function {
			ptrT, err := types.Instantiate(nil, ptr, []types.Type{T}, false)
			if err != nil {
				t.Fatalf("Failed to Instantiate %q by %q", ptr, T)
			}
			methods := types.NewMethodSet(types.NewPointer(ptrT))
			if methods.Len() != 1 {
				t.Fatalf("Expected 1 method for %q. got %d", ptrT, methods.Len())
			}
			return prog.MethodValue(methods.At(0))
		}

		intSliceTyp := types.NewSlice(types.Typ[types.Int])
		instance := instantiateLoadMethod(intSliceTyp) // (*Pointer[[]int]).Load
		if instance.Origin() != meth {
			t.Errorf("Expected Origin of %s to be %s. got %s", instance, meth, instance.Origin())
		}
		if len(instance.TypeArgs()) != 1 || !types.Identical(instance.TypeArgs()[0], intSliceTyp) {
			t.Errorf("Expected TypeArgs of %s to be %v. got %v", instance, []types.Type{intSliceTyp}, instance.TypeArgs())
		}

		// A second request with an identical type returns the same Function.
		second := instantiateLoadMethod(types.NewSlice(types.Typ[types.Int]))
		if second != instance {
			t.Error("Expected second identical instantiation to be the same function")
		}

		// (*Pointer[[]uint]).Load
		inst2 := instantiateLoadMethod(types.NewSlice(types.Typ[types.Uint]))

		if instance.Name() >= inst2.Name() {
			t.Errorf("Expected name of instance %s to be before instance %v", instance, inst2)
		}
	}
}

// TestCallsToInstances checks that calles of calls to generic functions,
// without monomorphization, are wrappers around the origin generic function.
func TestCallsToInstances(t *testing.T) {
	const input = `
package p

type I interface {
	Foo()
}

type A int
func (a A) Foo() {}

type J[T any] interface{ Bar() T }
type K[T any] struct{ J[T] }

func Id[T any] (t T) T {
	return t
}

func Lambda[T I]() func() func(T) {
	return func() func(T) {
		return T.Foo
	}
}

func NoOp[T any]() {}

func Bar[T interface { Foo(); ~int | ~string }, U any] (t T, u U) {
	Id[U](u)
	Id[T](t)
}

func Make[T any]() interface{} {
	NoOp[K[T]]()
	return nil
}

func entry(i int, a A) int {
	Lambda[A]()()(a)

	x := Make[int]()
	if j, ok := x.(interface{ Bar() int }); ok {
		print(j)
	}

	Bar[A, int](a, i)

	return Id[int](i)
}
`
	p, _ := buildPackage(t, input, ssa.SanityCheckFunctions)
	all := ssautil.AllFunctions(p.Prog)

	for _, ti := range []struct {
		orig         string
		instance     string
		tparams      string
		targs        string
		chTypeInstrs int // number of ChangeType instructions in f's body
	}{
		{"Id", "Id[int]", "[T]", "[int]", 2},
		{"Lambda", "Lambda[p.A]", "[T]", "[p.A]", 1},
		{"Make", "Make[int]", "[T]", "[int]", 0},
		{"NoOp", "NoOp[p.K[T]]", "[T]", "[p.K[T]]", 0},
	} {
		test := ti
		t.Run(test.instance, func(t *testing.T) {
			f := p.Members[test.orig].(*ssa.Function)
			if f == nil {
				t.Fatalf("origin function not found")
			}

			var i *ssa.Function
			for _, fn := range instancesOf(all, f) {
				if fn.Name() == test.instance {
					i = fn
					break
				}
			}
			if i == nil {
				t.Fatalf("instance not found")
			}

			// for logging on failures
			var body strings.Builder
			i.WriteTo(&body)
			t.Log(body.String())

			if len(i.Blocks) != 1 {
				t.Fatalf("body has more than 1 block")
			}

			if instrs := changeTypeInstrs(i.Blocks[0]); instrs != test.chTypeInstrs {
				t.Errorf("want %v instructions; got %v", test.chTypeInstrs, instrs)
			}

			if test.tparams != tparams(i) {
				t.Errorf("want %v type params; got %v", test.tparams, tparams(i))
			}

			if test.targs != targs(i) {
				t.Errorf("want %v type arguments; got %v", test.targs, targs(i))
			}
		})
	}
}

func tparams(f *ssa.Function) string {
	tplist := f.TypeParams()
	var tps []string
	for i := 0; i < tplist.Len(); i++ {
		tps = append(tps, tplist.At(i).String())
	}
	return fmt.Sprint(tps)
}

func targs(f *ssa.Function) string {
	var tas []string
	for _, ta := range f.TypeArgs() {
		tas = append(tas, ta.String())
	}
	return fmt.Sprint(tas)
}

func changeTypeInstrs(b *ssa.BasicBlock) int {
	cnt := 0
	for _, i := range b.Instrs {
		if _, ok := i.(*ssa.ChangeType); ok {
			cnt++
		}
	}
	return cnt
}

func TestInstanceUniqueness(t *testing.T) {
	const input = `
package p

func H[T any](t T) {
	print(t)
}

func F[T any](t T) {
	H[T](t)
	H[T](t)
	H[T](t)
}

func G[T any](t T) {
	H[T](t)
	H[T](t)
}

func Foo[T any, S any](t T, s S) {
	Foo[S, T](s, t)
	Foo[T, S](t, s)
}
`
	p, _ := buildPackage(t, input, ssa.SanityCheckFunctions)

	all := ssautil.AllFunctions(p.Prog)
	for _, test := range []struct {
		orig      string
		instances string
	}{
		{"H", "[p.H[T] p.H[T]]"},
		{"Foo", "[p.Foo[S T] p.Foo[T S]]"},
	} {
		t.Run(test.orig, func(t *testing.T) {
			f := p.Members[test.orig].(*ssa.Function)
			if f == nil {
				t.Fatalf("origin function not found")
			}

			instances := instancesOf(all, f)
			sort.Slice(instances, func(i, j int) bool { return instances[i].Name() < instances[j].Name() })

			if got := fmt.Sprintf("%v", instances); !reflect.DeepEqual(got, test.instances) {
				t.Errorf("got %v instances, want %v", got, test.instances)
			}
		})
	}
}

// instancesOf returns a new unordered slice of all instances of the
// specified function g in fns.
func instancesOf(fns map[*ssa.Function]bool, g *ssa.Function) []*ssa.Function {
	var instances []*ssa.Function
	for fn := range fns {
		if fn != g && fn.Origin() == g {
			instances = append(instances, fn)
		}
	}
	return instances
}
```

## File: go/ssa/instantiate.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"fmt"
	"go/types"
	"sync"
)

// A generic records information about a generic origin function,
// including a cache of existing instantiations.
type generic struct {
	instancesMu sync.Mutex
	instances   map[*typeList]*Function // canonical type arguments to an instance.
}

// instance returns a Function that is the instantiation of generic
// origin function fn with the type arguments targs.
//
// Any created instance is added to cr.
//
// Acquires fn.generic.instancesMu.
func (fn *Function) instance(targs []types.Type, b *builder) *Function {
	key := fn.Prog.canon.List(targs)

	gen := fn.generic

	gen.instancesMu.Lock()
	defer gen.instancesMu.Unlock()
	inst, ok := gen.instances[key]
	if !ok {
		inst = createInstance(fn, targs)
		inst.buildshared = b.shared()
		b.enqueue(inst)

		if gen.instances == nil {
			gen.instances = make(map[*typeList]*Function)
		}
		gen.instances[key] = inst
	} else {
		b.waitForSharedFunction(inst)
	}
	return inst
}

// createInstance returns the instantiation of generic function fn using targs.
//
// Requires fn.generic.instancesMu.
func createInstance(fn *Function, targs []types.Type) *Function {
	prog := fn.Prog

	// Compute signature.
	var sig *types.Signature
	var obj *types.Func
	if recv := fn.Signature.Recv(); recv != nil {
		// method
		obj = prog.canon.instantiateMethod(fn.object, targs, prog.ctxt)
		sig = obj.Type().(*types.Signature)
	} else {
		// function
		instSig, err := types.Instantiate(prog.ctxt, fn.Signature, targs, false)
		if err != nil {
			panic(err)
		}
		instance, ok := instSig.(*types.Signature)
		if !ok {
			panic("Instantiate of a Signature returned a non-signature")
		}
		obj = fn.object // instantiation does not exist yet
		sig = prog.canon.Type(instance).(*types.Signature)
	}

	// Choose strategy (instance or wrapper).
	var (
		synthetic string
		subst     *subster
		build     buildFunc
	)
	if prog.mode&InstantiateGenerics != 0 && !prog.isParameterized(targs...) {
		synthetic = fmt.Sprintf("instance of %s", fn.Name())
		if fn.syntax != nil {
			subst = makeSubster(prog.ctxt, obj, fn.typeparams, targs, false)
			build = (*builder).buildFromSyntax
		} else {
			build = (*builder).buildParamsOnly
		}
	} else {
		synthetic = fmt.Sprintf("instantiation wrapper of %s", fn.Name())
		build = (*builder).buildInstantiationWrapper
	}

	/* generic instance or instantiation wrapper */
	return &Function{
		name:           fmt.Sprintf("%s%s", fn.Name(), targs), // may not be unique
		object:         obj,
		Signature:      sig,
		Synthetic:      synthetic,
		syntax:         fn.syntax,    // \
		info:           fn.info,      //  } empty for non-created packages
		goversion:      fn.goversion, // /
		build:          build,
		topLevelOrigin: fn,
		pos:            obj.Pos(),
		Pkg:            nil,
		Prog:           fn.Prog,
		typeparams:     fn.typeparams, // share with origin
		typeargs:       targs,
		subst:          subst,
	}
}

// isParameterized reports whether any of the specified types contains
// a free type parameter. It is safe to call concurrently.
func (prog *Program) isParameterized(ts ...types.Type) bool {
	prog.hasParamsMu.Lock()
	defer prog.hasParamsMu.Unlock()

	// TODO(adonovan): profile. If this operation is expensive,
	// handle the most common but shallow cases such as T, pkg.T,
	// *T without consulting the cache under the lock.

	for _, t := range ts {
		if prog.hasParams.Has(t) {
			return true
		}
	}
	return false
}
```

## File: go/ssa/lift.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines the lifting pass which tries to "lift" Alloc
// cells (new/local variables) into SSA registers, replacing loads
// with the dominating stored value, eliminating loads and stores, and
// inserting φ-nodes as needed.

// Cited papers and resources:
//
// Ron Cytron et al. 1991. Efficiently computing SSA form...
// http://doi.acm.org/10.1145/115372.115320
//
// Cooper, Harvey, Kennedy.  2001.  A Simple, Fast Dominance Algorithm.
// Software Practice and Experience 2001, 4:1-10.
// http://www.hipersoft.rice.edu/grads/publications/dom14.pdf
//
// Daniel Berlin, llvmdev mailing list, 2012.
// http://lists.cs.uiuc.edu/pipermail/llvmdev/2012-January/046638.html
// (Be sure to expand the whole thread.)

// TODO(adonovan): opt: there are many optimizations worth evaluating, and
// the conventional wisdom for SSA construction is that a simple
// algorithm well engineered often beats those of better asymptotic
// complexity on all but the most egregious inputs.
//
// Danny Berlin suggests that the Cooper et al. algorithm for
// computing the dominance frontier is superior to Cytron et al.
// Furthermore he recommends that rather than computing the DF for the
// whole function then renaming all alloc cells, it may be cheaper to
// compute the DF for each alloc cell separately and throw it away.
//
// Consider exploiting liveness information to avoid creating dead
// φ-nodes which we then immediately remove.
//
// Also see many other "TODO: opt" suggestions in the code.

import (
	"fmt"
	"go/token"
	"math/big"
	"os"
	"slices"

	"golang.org/x/tools/internal/typeparams"
)

// If true, show diagnostic information at each step of lifting.
// Very verbose.
const debugLifting = false

// domFrontier maps each block to the set of blocks in its dominance
// frontier.  The outer slice is conceptually a map keyed by
// Block.Index.  The inner slice is conceptually a set, possibly
// containing duplicates.
//
// TODO(adonovan): opt: measure impact of dups; consider a packed bit
// representation, e.g. big.Int, and bitwise parallel operations for
// the union step in the Children loop.
//
// domFrontier's methods mutate the slice's elements but not its
// length, so their receivers needn't be pointers.
type domFrontier [][]*BasicBlock

func (df domFrontier) add(u, v *BasicBlock) {
	p := &df[u.Index]
	*p = append(*p, v)
}

// build builds the dominance frontier df for the dominator (sub)tree
// rooted at u, using the Cytron et al. algorithm.
//
// TODO(adonovan): opt: consider Berlin approach, computing pruned SSA
// by pruning the entire IDF computation, rather than merely pruning
// the DF -> IDF step.
func (df domFrontier) build(u *BasicBlock) {
	// Encounter each node u in postorder of dom tree.
	for _, child := range u.dom.children {
		df.build(child)
	}
	for _, vb := range u.Succs {
		if v := vb.dom; v.idom != u {
			df.add(u, vb)
		}
	}
	for _, w := range u.dom.children {
		for _, vb := range df[w.Index] {
			// TODO(adonovan): opt: use word-parallel bitwise union.
			if v := vb.dom; v.idom != u {
				df.add(u, vb)
			}
		}
	}
}

func buildDomFrontier(fn *Function) domFrontier {
	df := make(domFrontier, len(fn.Blocks))
	df.build(fn.Blocks[0])
	if fn.Recover != nil {
		df.build(fn.Recover)
	}
	return df
}

func removeInstr(refs []Instruction, instr Instruction) []Instruction {
	return slices.DeleteFunc(refs, func(i Instruction) bool { return i == instr })
}

// lift replaces local and new Allocs accessed only with
// load/store by SSA registers, inserting φ-nodes where necessary.
// The result is a program in classical pruned SSA form.
//
// Preconditions:
// - fn has no dead blocks (blockopt has run).
// - Def/use info (Operands and Referrers) is up-to-date.
// - The dominator tree is up-to-date.
func lift(fn *Function) {
	// TODO(adonovan): opt: lots of little optimizations may be
	// worthwhile here, especially if they cause us to avoid
	// buildDomFrontier.  For example:
	//
	// - Alloc never loaded?  Eliminate.
	// - Alloc never stored?  Replace all loads with a zero constant.
	// - Alloc stored once?  Replace loads with dominating store;
	//   don't forget that an Alloc is itself an effective store
	//   of zero.
	// - Alloc used only within a single block?
	//   Use degenerate algorithm avoiding φ-nodes.
	// - Consider synergy with scalar replacement of aggregates (SRA).
	//   e.g. *(&x.f) where x is an Alloc.
	//   Perhaps we'd get better results if we generated this as x.f
	//   i.e. Field(x, .f) instead of Load(FieldIndex(x, .f)).
	//   Unclear.
	//
	// But we will start with the simplest correct code.
	df := buildDomFrontier(fn)

	if debugLifting {
		title := false
		for i, blocks := range df {
			if blocks != nil {
				if !title {
					fmt.Fprintf(os.Stderr, "Dominance frontier of %s:\n", fn)
					title = true
				}
				fmt.Fprintf(os.Stderr, "\t%s: %s\n", fn.Blocks[i], blocks)
			}
		}
	}

	newPhis := make(newPhiMap)

	// During this pass we will replace some BasicBlock.Instrs
	// (allocs, loads and stores) with nil, keeping a count in
	// BasicBlock.gaps.  At the end we will reset Instrs to the
	// concatenation of all non-dead newPhis and non-nil Instrs
	// for the block, reusing the original array if space permits.

	// While we're here, we also eliminate 'rundefers'
	// instructions and ssa:deferstack() in functions that contain no
	// 'defer' instructions. For now, we also eliminate
	// 's = ssa:deferstack()' calls if s doesn't escape, replacing s
	// with nil in Defer{DeferStack: s}. This has the same meaning,
	// but allows eliminating the intrinsic function `ssa:deferstack()`
	// (unless it is needed due to range-over-func instances). This gives
	// ssa users more time to support range-over-func.
	usesDefer := false
	deferstackAlloc, deferstackCall := deferstackPreamble(fn)
	eliminateDeferStack := deferstackAlloc != nil && !deferstackAlloc.Heap

	// A counter used to generate ~unique ids for Phi nodes, as an
	// aid to debugging.  We use large numbers to make them highly
	// visible.  All nodes are renumbered later.
	fresh := 1000

	// Determine which allocs we can lift and number them densely.
	// The renaming phase uses this numbering for compact maps.
	numAllocs := 0
	for _, b := range fn.Blocks {
		b.gaps = 0
		b.rundefers = 0
		for _, instr := range b.Instrs {
			switch instr := instr.(type) {
			case *Alloc:
				index := -1
				if liftAlloc(df, instr, newPhis, &fresh) {
					index = numAllocs
					numAllocs++
				}
				instr.index = index
			case *Defer:
				usesDefer = true
				if eliminateDeferStack {
					// Clear DeferStack and remove references to loads
					if instr.DeferStack != nil {
						if refs := instr.DeferStack.Referrers(); refs != nil {
							*refs = removeInstr(*refs, instr)
						}
						instr.DeferStack = nil
					}
				}
			case *RunDefers:
				b.rundefers++
			}
		}
	}

	// renaming maps an alloc (keyed by index) to its replacement
	// value.  Initially the renaming contains nil, signifying the
	// zero constant of the appropriate type; we construct the
	// Const lazily at most once on each path through the domtree.
	// TODO(adonovan): opt: cache per-function not per subtree.
	renaming := make([]Value, numAllocs)

	// Renaming.
	rename(fn.Blocks[0], renaming, newPhis)

	// Eliminate dead φ-nodes.
	removeDeadPhis(fn.Blocks, newPhis)

	// Eliminate ssa:deferstack() call.
	if eliminateDeferStack {
		b := deferstackCall.block
		for i, instr := range b.Instrs {
			if instr == deferstackCall {
				b.Instrs[i] = nil
				b.gaps++
				break
			}
		}
	}

	// Prepend remaining live φ-nodes to each block.
	for _, b := range fn.Blocks {
		nps := newPhis[b]
		j := len(nps)

		rundefersToKill := b.rundefers
		if usesDefer {
			rundefersToKill = 0
		}

		if j+b.gaps+rundefersToKill == 0 {
			continue // fast path: no new phis or gaps
		}

		// Compact nps + non-nil Instrs into a new slice.
		// TODO(adonovan): opt: compact in situ (rightwards)
		// if Instrs has sufficient space or slack.
		dst := make([]Instruction, len(b.Instrs)+j-b.gaps-rundefersToKill)
		for i, np := range nps {
			dst[i] = np.phi
		}
		for _, instr := range b.Instrs {
			if instr == nil {
				continue
			}
			if !usesDefer {
				if _, ok := instr.(*RunDefers); ok {
					continue
				}
			}
			dst[j] = instr
			j++
		}
		b.Instrs = dst
	}

	// Remove any fn.Locals that were lifted.
	j := 0
	for _, l := range fn.Locals {
		if l.index < 0 {
			fn.Locals[j] = l
			j++
		}
	}
	// Nil out fn.Locals[j:] to aid GC.
	for i := j; i < len(fn.Locals); i++ {
		fn.Locals[i] = nil
	}
	fn.Locals = fn.Locals[:j]
}

// removeDeadPhis removes φ-nodes not transitively needed by a
// non-Phi, non-DebugRef instruction.
func removeDeadPhis(blocks []*BasicBlock, newPhis newPhiMap) {
	// First pass: find the set of "live" φ-nodes: those reachable
	// from some non-Phi instruction.
	//
	// We compute reachability in reverse, starting from each φ,
	// rather than forwards, starting from each live non-Phi
	// instruction, because this way visits much less of the
	// Value graph.
	livePhis := make(map[*Phi]bool)
	for _, npList := range newPhis {
		for _, np := range npList {
			phi := np.phi
			if !livePhis[phi] && phiHasDirectReferrer(phi) {
				markLivePhi(livePhis, phi)
			}
		}
	}

	// Existing φ-nodes due to && and || operators
	// are all considered live (see Go issue 19622).
	for _, b := range blocks {
		for _, phi := range b.phis() {
			markLivePhi(livePhis, phi.(*Phi))
		}
	}

	// Second pass: eliminate unused phis from newPhis.
	for block, npList := range newPhis {
		j := 0
		for _, np := range npList {
			if livePhis[np.phi] {
				npList[j] = np
				j++
			} else {
				// discard it, first removing it from referrers
				for _, val := range np.phi.Edges {
					if refs := val.Referrers(); refs != nil {
						*refs = removeInstr(*refs, np.phi)
					}
				}
				np.phi.block = nil
			}
		}
		newPhis[block] = npList[:j]
	}
}

// markLivePhi marks phi, and all φ-nodes transitively reachable via
// its Operands, live.
func markLivePhi(livePhis map[*Phi]bool, phi *Phi) {
	livePhis[phi] = true
	for _, rand := range phi.Operands(nil) {
		if q, ok := (*rand).(*Phi); ok {
			if !livePhis[q] {
				markLivePhi(livePhis, q)
			}
		}
	}
}

// phiHasDirectReferrer reports whether phi is directly referred to by
// a non-Phi instruction.  Such instructions are the
// roots of the liveness traversal.
func phiHasDirectReferrer(phi *Phi) bool {
	for _, instr := range *phi.Referrers() {
		if _, ok := instr.(*Phi); !ok {
			return true
		}
	}
	return false
}

type blockSet struct{ big.Int } // (inherit methods from Int)

// add adds b to the set and returns true if the set changed.
func (s *blockSet) add(b *BasicBlock) bool {
	i := b.Index
	if s.Bit(i) != 0 {
		return false
	}
	s.SetBit(&s.Int, i, 1)
	return true
}

// take removes an arbitrary element from a set s and
// returns its index, or returns -1 if empty.
func (s *blockSet) take() int {
	l := s.BitLen()
	for i := 0; i < l; i++ {
		if s.Bit(i) == 1 {
			s.SetBit(&s.Int, i, 0)
			return i
		}
	}
	return -1
}

// newPhi is a pair of a newly introduced φ-node and the lifted Alloc
// it replaces.
type newPhi struct {
	phi   *Phi
	alloc *Alloc
}

// newPhiMap records for each basic block, the set of newPhis that
// must be prepended to the block.
type newPhiMap map[*BasicBlock][]newPhi

// liftAlloc determines whether alloc can be lifted into registers,
// and if so, it populates newPhis with all the φ-nodes it may require
// and returns true.
//
// fresh is a source of fresh ids for phi nodes.
func liftAlloc(df domFrontier, alloc *Alloc, newPhis newPhiMap, fresh *int) bool {
	// Don't lift result values in functions that defer
	// calls that may recover from panic.
	if fn := alloc.Parent(); fn.Recover != nil {
		for _, nr := range fn.results {
			if nr == alloc {
				return false
			}
		}
	}

	// Compute defblocks, the set of blocks containing a
	// definition of the alloc cell.
	var defblocks blockSet
	for _, instr := range *alloc.Referrers() {
		// Bail out if we discover the alloc is not liftable;
		// the only operations permitted to use the alloc are
		// loads/stores into the cell, and DebugRef.
		switch instr := instr.(type) {
		case *Store:
			if instr.Val == alloc {
				return false // address used as value
			}
			if instr.Addr != alloc {
				panic("Alloc.Referrers is inconsistent")
			}
			defblocks.add(instr.Block())
		case *UnOp:
			if instr.Op != token.MUL {
				return false // not a load
			}
			if instr.X != alloc {
				panic("Alloc.Referrers is inconsistent")
			}
		case *DebugRef:
			// ok
		default:
			return false // some other instruction
		}
	}
	// The Alloc itself counts as a (zero) definition of the cell.
	defblocks.add(alloc.Block())

	if debugLifting {
		fmt.Fprintln(os.Stderr, "\tlifting ", alloc, alloc.Name())
	}

	fn := alloc.Parent()

	// Φ-insertion.
	//
	// What follows is the body of the main loop of the insert-φ
	// function described by Cytron et al, but instead of using
	// counter tricks, we just reset the 'hasAlready' and 'work'
	// sets each iteration.  These are bitmaps so it's pretty cheap.
	//
	// TODO(adonovan): opt: recycle slice storage for W,
	// hasAlready, defBlocks across liftAlloc calls.
	var hasAlready blockSet

	// Initialize W and work to defblocks.
	var work blockSet = defblocks // blocks seen
	var W blockSet                // blocks to do
	W.Set(&defblocks.Int)

	// Traverse iterated dominance frontier, inserting φ-nodes.
	for i := W.take(); i != -1; i = W.take() {
		u := fn.Blocks[i]
		for _, v := range df[u.Index] {
			if hasAlready.add(v) {
				// Create φ-node.
				// It will be prepended to v.Instrs later, if needed.
				phi := &Phi{
					Edges:   make([]Value, len(v.Preds)),
					Comment: alloc.Comment,
				}
				// This is merely a debugging aid:
				phi.setNum(*fresh)
				*fresh++

				phi.pos = alloc.Pos()
				phi.setType(typeparams.MustDeref(alloc.Type()))
				phi.block = v
				if debugLifting {
					fmt.Fprintf(os.Stderr, "\tplace %s = %s at block %s\n", phi.Name(), phi, v)
				}
				newPhis[v] = append(newPhis[v], newPhi{phi, alloc})

				if work.add(v) {
					W.add(v)
				}
			}
		}
	}

	return true
}

// replaceAll replaces all intraprocedural uses of x with y,
// updating x.Referrers and y.Referrers.
// Precondition: x.Referrers() != nil, i.e. x must be local to some function.
func replaceAll(x, y Value) {
	var rands []*Value
	pxrefs := x.Referrers()
	pyrefs := y.Referrers()
	for _, instr := range *pxrefs {
		rands = instr.Operands(rands[:0]) // recycle storage
		for _, rand := range rands {
			if *rand != nil {
				if *rand == x {
					*rand = y
				}
			}
		}
		if pyrefs != nil {
			*pyrefs = append(*pyrefs, instr) // dups ok
		}
	}
	*pxrefs = nil // x is now unreferenced
}

// renamed returns the value to which alloc is being renamed,
// constructing it lazily if it's the implicit zero initialization.
func renamed(renaming []Value, alloc *Alloc) Value {
	v := renaming[alloc.index]
	if v == nil {
		v = zeroConst(typeparams.MustDeref(alloc.Type()))
		renaming[alloc.index] = v
	}
	return v
}

// rename implements the (Cytron et al) SSA renaming algorithm, a
// preorder traversal of the dominator tree replacing all loads of
// Alloc cells with the value stored to that cell by the dominating
// store instruction.  For lifting, we need only consider loads,
// stores and φ-nodes.
//
// renaming is a map from *Alloc (keyed by index number) to its
// dominating stored value; newPhis[x] is the set of new φ-nodes to be
// prepended to block x.
func rename(u *BasicBlock, renaming []Value, newPhis newPhiMap) {
	// Each φ-node becomes the new name for its associated Alloc.
	for _, np := range newPhis[u] {
		phi := np.phi
		alloc := np.alloc
		renaming[alloc.index] = phi
	}

	// Rename loads and stores of allocs.
	for i, instr := range u.Instrs {
		switch instr := instr.(type) {
		case *Alloc:
			if instr.index >= 0 { // store of zero to Alloc cell
				// Replace dominated loads by the zero value.
				renaming[instr.index] = nil
				if debugLifting {
					fmt.Fprintf(os.Stderr, "\tkill alloc %s\n", instr)
				}
				// Delete the Alloc.
				u.Instrs[i] = nil
				u.gaps++
			}

		case *Store:
			if alloc, ok := instr.Addr.(*Alloc); ok && alloc.index >= 0 { // store to Alloc cell
				// Replace dominated loads by the stored value.
				renaming[alloc.index] = instr.Val
				if debugLifting {
					fmt.Fprintf(os.Stderr, "\tkill store %s; new value: %s\n",
						instr, instr.Val.Name())
				}
				// Remove the store from the referrer list of the stored value.
				if refs := instr.Val.Referrers(); refs != nil {
					*refs = removeInstr(*refs, instr)
				}
				// Delete the Store.
				u.Instrs[i] = nil
				u.gaps++
			}

		case *UnOp:
			if instr.Op == token.MUL {
				if alloc, ok := instr.X.(*Alloc); ok && alloc.index >= 0 { // load of Alloc cell
					newval := renamed(renaming, alloc)
					if debugLifting {
						fmt.Fprintf(os.Stderr, "\tupdate load %s = %s with %s\n",
							instr.Name(), instr, newval.Name())
					}
					// Replace all references to
					// the loaded value by the
					// dominating stored value.
					replaceAll(instr, newval)
					// Delete the Load.
					u.Instrs[i] = nil
					u.gaps++
				}
			}

		case *DebugRef:
			if alloc, ok := instr.X.(*Alloc); ok && alloc.index >= 0 { // ref of Alloc cell
				if instr.IsAddr {
					instr.X = renamed(renaming, alloc)
					instr.IsAddr = false

					// Add DebugRef to instr.X's referrers.
					if refs := instr.X.Referrers(); refs != nil {
						*refs = append(*refs, instr)
					}
				} else {
					// A source expression denotes the address
					// of an Alloc that was optimized away.
					instr.X = nil

					// Delete the DebugRef.
					u.Instrs[i] = nil
					u.gaps++
				}
			}
		}
	}

	// For each φ-node in a CFG successor, rename the edge.
	for _, v := range u.Succs {
		phis := newPhis[v]
		if len(phis) == 0 {
			continue
		}
		i := v.predIndex(u)
		for _, np := range phis {
			phi := np.phi
			alloc := np.alloc
			newval := renamed(renaming, alloc)
			if debugLifting {
				fmt.Fprintf(os.Stderr, "\tsetphi %s edge %s -> %s (#%d) (alloc=%s) := %s\n",
					phi.Name(), u, v, i, alloc.Name(), newval.Name())
			}
			phi.Edges[i] = newval
			if prefs := newval.Referrers(); prefs != nil {
				*prefs = append(*prefs, phi)
			}
		}
	}

	// Continue depth-first recursion over domtree, pushing a
	// fresh copy of the renaming map for each subtree.
	for i, v := range u.dom.children {
		r := renaming
		if i < len(u.dom.children)-1 {
			// On all but the final iteration, we must make
			// a copy to avoid destructive update.
			r = make([]Value, len(renaming))
			copy(r, renaming)
		}
		rename(v, r, newPhis)
	}

}

// deferstackPreamble returns the *Alloc and ssa:deferstack() call for fn.deferstack.
func deferstackPreamble(fn *Function) (*Alloc, *Call) {
	if alloc, _ := fn.vars[fn.deferstack].(*Alloc); alloc != nil {
		for _, ref := range *alloc.Referrers() {
			if ref, _ := ref.(*Store); ref != nil && ref.Addr == alloc {
				if call, _ := ref.Val.(*Call); call != nil {
					return alloc, call
				}
			}
		}
	}
	return nil, nil
}
```

## File: go/ssa/lvalue.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// lvalues are the union of addressable expressions and map-index
// expressions.

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/internal/typeparams"
)

// An lvalue represents an assignable location that may appear on the
// left-hand side of an assignment.  This is a generalization of a
// pointer to permit updates to elements of maps.
type lvalue interface {
	store(fn *Function, v Value) // stores v into the location
	load(fn *Function) Value     // loads the contents of the location
	address(fn *Function) Value  // address of the location
	typ() types.Type             // returns the type of the location
}

// An address is an lvalue represented by a true pointer.
type address struct {
	addr Value     // must have a pointer core type.
	pos  token.Pos // source position
	expr ast.Expr  // source syntax of the value (not address) [debug mode]
}

func (a *address) load(fn *Function) Value {
	load := emitLoad(fn, a.addr)
	load.pos = a.pos
	return load
}

func (a *address) store(fn *Function, v Value) {
	store := emitStore(fn, a.addr, v, a.pos)
	if a.expr != nil {
		// store.Val is v, converted for assignability.
		emitDebugRef(fn, a.expr, store.Val, false)
	}
}

func (a *address) address(fn *Function) Value {
	if a.expr != nil {
		emitDebugRef(fn, a.expr, a.addr, true)
	}
	return a.addr
}

func (a *address) typ() types.Type {
	return typeparams.MustDeref(a.addr.Type())
}

// An element is an lvalue represented by m[k], the location of an
// element of a map.  These locations are not addressable
// since pointers cannot be formed from them, but they do support
// load() and store().
type element struct {
	m, k Value      // map
	t    types.Type // map element type
	pos  token.Pos  // source position of colon ({k:v}) or lbrack (m[k]=v)
}

func (e *element) load(fn *Function) Value {
	l := &Lookup{
		X:     e.m,
		Index: e.k,
	}
	l.setPos(e.pos)
	l.setType(e.t)
	return fn.emit(l)
}

func (e *element) store(fn *Function, v Value) {
	up := &MapUpdate{
		Map:   e.m,
		Key:   e.k,
		Value: emitConv(fn, v, e.t),
	}
	up.pos = e.pos
	fn.emit(up)
}

func (e *element) address(fn *Function) Value {
	panic("map elements are not addressable")
}

func (e *element) typ() types.Type {
	return e.t
}

// A lazyAddress is an lvalue whose address is the result of an instruction.
// These work like an *address except a new address.address() Value
// is created on each load, store and address call.
// A lazyAddress can be used to control when a side effect (nil pointer
// dereference, index out of bounds) of using a location happens.
type lazyAddress struct {
	addr func(fn *Function) Value // emit to fn the computation of the address
	t    types.Type               // type of the location
	pos  token.Pos                // source position
	expr ast.Expr                 // source syntax of the value (not address) [debug mode]
}

func (l *lazyAddress) load(fn *Function) Value {
	load := emitLoad(fn, l.addr(fn))
	load.pos = l.pos
	return load
}

func (l *lazyAddress) store(fn *Function, v Value) {
	store := emitStore(fn, l.addr(fn), v, l.pos)
	if l.expr != nil {
		// store.Val is v, converted for assignability.
		emitDebugRef(fn, l.expr, store.Val, false)
	}
}

func (l *lazyAddress) address(fn *Function) Value {
	addr := l.addr(fn)
	if l.expr != nil {
		emitDebugRef(fn, l.expr, addr, true)
	}
	return addr
}

func (l *lazyAddress) typ() types.Type { return l.t }

// A blank is a dummy variable whose name is "_".
// It is not reified: loads are illegal and stores are ignored.
type blank struct{}

func (bl blank) load(fn *Function) Value {
	panic("blank.load is illegal")
}

func (bl blank) store(fn *Function, v Value) {
	// no-op
}

func (bl blank) address(fn *Function) Value {
	panic("blank var is not addressable")
}

func (bl blank) typ() types.Type {
	// This should be the type of the blank Ident; the typechecker
	// doesn't provide this yet, but fortunately, we don't need it
	// yet either.
	panic("blank.typ is unimplemented")
}
```

## File: go/ssa/methods_test.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// Tests that MethodValue returns the expected method.
func TestMethodValue(t *testing.T) {
	input := `
package p

type I interface{ M() }

type S int
func (S) M() {}
type R[T any] struct{ S }

var i I
var s S
var r R[string]

func selections[T any]() {
	_ = i.M
	_ = s.M
	_ = r.M

	var v R[T]
	_ = v.M
}
`

	// Parse the file.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "input.go", input, 0)
	if err != nil {
		t.Error(err)
		return
	}

	// Build an SSA program from the parsed file.
	p, info, err := ssautil.BuildPackage(&types.Config{}, fset,
		types.NewPackage("p", ""), []*ast.File{f}, ssa.SanityCheckFunctions)
	if err != nil {
		t.Error(err)
		return
	}

	// Collect all of the *types.Selection in the function "selections".
	var selections []*types.Selection
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "selections" {
			for _, stmt := range fn.Body.List {
				if assign, ok := stmt.(*ast.AssignStmt); ok {
					sel := assign.Rhs[0].(*ast.SelectorExpr)
					selections = append(selections, info.Selections[sel])
				}
			}
		}
	}

	wants := map[string]string{
		"method (p.S) M()":         "(p.S).M",
		"method (p.R[string]) M()": "(p.R[string]).M",
		"method (p.I) M()":         "nil", // interface
		"method (p.R[T]) M()":      "nil", // parameterized
	}
	if len(wants) != len(selections) {
		t.Fatalf("Wanted %d selections. got %d", len(wants), len(selections))
	}
	for _, selection := range selections {
		var got string
		if m := p.Prog.MethodValue(selection); m != nil {
			got = m.String()
		} else {
			got = "nil"
		}
		if want := wants[selection.String()]; want != got {
			t.Errorf("p.Prog.MethodValue(%s) expected %q. got %q", selection, want, got)
		}
	}
}
```

## File: go/ssa/methods.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines utilities for population of method sets.

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/typesinternal"
)

// MethodValue returns the Function implementing method sel, building
// wrapper methods on demand. It returns nil if sel denotes an
// interface or generic method.
//
// Precondition: sel.Kind() == MethodVal.
//
// Thread-safe.
//
// Acquires prog.methodsMu.
func (prog *Program) MethodValue(sel *types.Selection) *Function {
	if sel.Kind() != types.MethodVal {
		panic(fmt.Sprintf("MethodValue(%s) kind != MethodVal", sel))
	}
	T := sel.Recv()
	if types.IsInterface(T) {
		return nil // interface method or type parameter
	}

	if prog.isParameterized(T) {
		return nil // generic method
	}

	if prog.mode&LogSource != 0 {
		defer logStack("MethodValue %s %v", T, sel)()
	}

	var b builder

	m := func() *Function {
		prog.methodsMu.Lock()
		defer prog.methodsMu.Unlock()

		// Get or create SSA method set.
		mset, ok := prog.methodSets.At(T).(*methodSet)
		if !ok {
			mset = &methodSet{mapping: make(map[string]*Function)}
			prog.methodSets.Set(T, mset)
		}

		// Get or create SSA method.
		id := sel.Obj().Id()
		fn, ok := mset.mapping[id]
		if !ok {
			obj := sel.Obj().(*types.Func)
			needsPromotion := len(sel.Index()) > 1
			needsIndirection := !isPointer(recvType(obj)) && isPointer(T)
			if needsPromotion || needsIndirection {
				fn = createWrapper(prog, toSelection(sel))
				fn.buildshared = b.shared()
				b.enqueue(fn)
			} else {
				fn = prog.objectMethod(obj, &b)
			}
			if fn.Signature.Recv() == nil {
				panic(fn)
			}
			mset.mapping[id] = fn
		} else {
			b.waitForSharedFunction(fn)
		}

		return fn
	}()

	b.iterate()

	return m
}

// objectMethod returns the Function for a given method symbol.
// The symbol may be an instance of a generic function. It need not
// belong to an existing SSA package created by a call to
// prog.CreatePackage.
//
// objectMethod panics if the function is not a method.
//
// Acquires prog.objectMethodsMu.
func (prog *Program) objectMethod(obj *types.Func, b *builder) *Function {
	sig := obj.Type().(*types.Signature)
	if sig.Recv() == nil {
		panic("not a method: " + obj.String())
	}

	// Belongs to a created package?
	if fn := prog.FuncValue(obj); fn != nil {
		return fn
	}

	// Instantiation of generic?
	if originObj := obj.Origin(); originObj != obj {
		origin := prog.objectMethod(originObj, b)
		assert(origin.typeparams.Len() > 0, "origin is not generic")
		targs := receiverTypeArgs(obj)
		return origin.instance(targs, b)
	}

	// Consult/update cache of methods created from types.Func.
	prog.objectMethodsMu.Lock()
	defer prog.objectMethodsMu.Unlock()
	fn, ok := prog.objectMethods[obj]
	if !ok {
		fn = createFunction(prog, obj, obj.Name(), nil, nil, "")
		fn.Synthetic = "from type information (on demand)"
		fn.buildshared = b.shared()
		b.enqueue(fn)

		if prog.objectMethods == nil {
			prog.objectMethods = make(map[*types.Func]*Function)
		}
		prog.objectMethods[obj] = fn
	} else {
		b.waitForSharedFunction(fn)
	}
	return fn
}

// LookupMethod returns the implementation of the method of type T
// identified by (pkg, name).  It returns nil if the method exists but
// is an interface method or generic method, and panics if T has no such method.
func (prog *Program) LookupMethod(T types.Type, pkg *types.Package, name string) *Function {
	sel := prog.MethodSets.MethodSet(T).Lookup(pkg, name)
	if sel == nil {
		panic(fmt.Sprintf("%s has no method %s", T, types.Id(pkg, name)))
	}
	return prog.MethodValue(sel)
}

// methodSet contains the (concrete) methods of a concrete type (non-interface, non-parameterized).
type methodSet struct {
	mapping map[string]*Function // populated lazily
}

// RuntimeTypes returns a new unordered slice containing all types in
// the program for which a runtime type is required.
//
// A runtime type is required for any non-parameterized, non-interface
// type that is converted to an interface, or for any type (including
// interface types) derivable from one through reflection.
//
// The methods of such types may be reachable through reflection or
// interface calls even if they are never called directly.
//
// Thread-safe.
//
// Acquires prog.makeInterfaceTypesMu.
func (prog *Program) RuntimeTypes() []types.Type {
	prog.makeInterfaceTypesMu.Lock()
	defer prog.makeInterfaceTypesMu.Unlock()

	// Compute the derived types on demand, since many SSA clients
	// never call RuntimeTypes, and those that do typically call
	// it once (often within ssautil.AllFunctions, which will
	// eventually not use it; see Go issue #69291.) This
	// eliminates the need to eagerly compute all the element
	// types during SSA building.
	var runtimeTypes []types.Type
	add := func(t types.Type) { runtimeTypes = append(runtimeTypes, t) }
	var set typeutil.Map // for de-duping identical types
	for t := range prog.makeInterfaceTypes {
		typesinternal.ForEachElement(&set, &prog.MethodSets, t, add)
	}

	return runtimeTypes
}
```

## File: go/ssa/mode.go
```go
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines the BuilderMode type and its command-line flag.

import (
	"bytes"
	"fmt"
)

// BuilderMode is a bitmask of options for diagnostics and checking.
//
// *BuilderMode satisfies the flag.Value interface.  Example:
//
//	var mode = ssa.BuilderMode(0)
//	func init() { flag.Var(&mode, "build", ssa.BuilderModeDoc) }
type BuilderMode uint

const (
	PrintPackages        BuilderMode = 1 << iota // Print package inventory to stdout
	PrintFunctions                               // Print function SSA code to stdout
	LogSource                                    // Log source locations as SSA builder progresses
	SanityCheckFunctions                         // Perform sanity checking of function bodies
	NaiveForm                                    // Build naïve SSA form: don't replace local loads/stores with registers
	BuildSerially                                // Build packages serially, not in parallel.
	GlobalDebug                                  // Enable debug info for all packages
	BareInits                                    // Build init functions without guards or calls to dependent inits
	InstantiateGenerics                          // Instantiate generics functions (monomorphize) while building
)

const BuilderModeDoc = `Options controlling the SSA builder.
The value is a sequence of zero or more of these letters:
C	perform sanity [C]hecking of the SSA form.
D	include [D]ebug info for every function.
P	print [P]ackage inventory.
F	print [F]unction SSA code.
S	log [S]ource locations as SSA builder progresses.
L	build distinct packages seria[L]ly instead of in parallel.
N	build [N]aive SSA form: don't replace local loads/stores with registers.
I	build bare [I]nit functions: no init guards or calls to dependent inits.
G   instantiate [G]eneric function bodies via monomorphization
`

func (m BuilderMode) String() string {
	var buf bytes.Buffer
	if m&GlobalDebug != 0 {
		buf.WriteByte('D')
	}
	if m&PrintPackages != 0 {
		buf.WriteByte('P')
	}
	if m&PrintFunctions != 0 {
		buf.WriteByte('F')
	}
	if m&LogSource != 0 {
		buf.WriteByte('S')
	}
	if m&SanityCheckFunctions != 0 {
		buf.WriteByte('C')
	}
	if m&NaiveForm != 0 {
		buf.WriteByte('N')
	}
	if m&BuildSerially != 0 {
		buf.WriteByte('L')
	}
	if m&BareInits != 0 {
		buf.WriteByte('I')
	}
	if m&InstantiateGenerics != 0 {
		buf.WriteByte('G')
	}
	return buf.String()
}

// Set parses the flag characters in s and updates *m.
func (m *BuilderMode) Set(s string) error {
	var mode BuilderMode
	for _, c := range s {
		switch c {
		case 'D':
			mode |= GlobalDebug
		case 'P':
			mode |= PrintPackages
		case 'F':
			mode |= PrintFunctions
		case 'S':
			mode |= LogSource | BuildSerially
		case 'C':
			mode |= SanityCheckFunctions
		case 'N':
			mode |= NaiveForm
		case 'L':
			mode |= BuildSerially
		case 'I':
			mode |= BareInits
		case 'G':
			mode |= InstantiateGenerics
		default:
			return fmt.Errorf("unknown BuilderMode option: %q", c)
		}
	}
	*m = mode
	return nil
}

// Get returns m.
func (m BuilderMode) Get() any { return m }
```

## File: go/ssa/print.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file implements the String() methods for all Value and
// Instruction types.

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/typeparams"
)

// relName returns the name of v relative to i.
// In most cases, this is identical to v.Name(), but references to
// Functions (including methods) and Globals use RelString and
// all types are displayed with relType, so that only cross-package
// references are package-qualified.
func relName(v Value, i Instruction) string {
	var from *types.Package
	if i != nil {
		from = i.Parent().relPkg()
	}
	switch v := v.(type) {
	case Member: // *Function or *Global
		return v.RelString(from)
	case *Const:
		return v.RelString(from)
	}
	return v.Name()
}

func relType(t types.Type, from *types.Package) string {
	return types.TypeString(t, types.RelativeTo(from))
}

func relTerm(term *types.Term, from *types.Package) string {
	s := relType(term.Type(), from)
	if term.Tilde() {
		return "~" + s
	}
	return s
}

func relString(m Member, from *types.Package) string {
	// NB: not all globals have an Object (e.g. init$guard),
	// so use Package().Object not Object.Package().
	if pkg := m.Package().Pkg; pkg != nil && pkg != from {
		return fmt.Sprintf("%s.%s", pkg.Path(), m.Name())
	}
	return m.Name()
}

// Value.String()
//
// This method is provided only for debugging.
// It never appears in disassembly, which uses Value.Name().

func (v *Parameter) String() string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("parameter %s : %s", v.Name(), relType(v.Type(), from))
}

func (v *FreeVar) String() string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("freevar %s : %s", v.Name(), relType(v.Type(), from))
}

func (v *Builtin) String() string {
	return fmt.Sprintf("builtin %s", v.Name())
}

// Instruction.String()

func (v *Alloc) String() string {
	op := "local"
	if v.Heap {
		op = "new"
	}
	from := v.Parent().relPkg()
	return fmt.Sprintf("%s %s (%s)", op, relType(typeparams.MustDeref(v.Type()), from), v.Comment)
}

func (v *Phi) String() string {
	var b bytes.Buffer
	b.WriteString("phi [")
	for i, edge := range v.Edges {
		if i > 0 {
			b.WriteString(", ")
		}
		// Be robust against malformed CFG.
		if v.block == nil {
			b.WriteString("??")
			continue
		}
		block := -1
		if i < len(v.block.Preds) {
			block = v.block.Preds[i].Index
		}
		fmt.Fprintf(&b, "%d: ", block)
		edgeVal := "<nil>" // be robust
		if edge != nil {
			edgeVal = relName(edge, v)
		}
		b.WriteString(edgeVal)
	}
	b.WriteString("]")
	if v.Comment != "" {
		b.WriteString(" #")
		b.WriteString(v.Comment)
	}
	return b.String()
}

func printCall(v *CallCommon, prefix string, instr Instruction) string {
	var b bytes.Buffer
	b.WriteString(prefix)
	if !v.IsInvoke() {
		b.WriteString(relName(v.Value, instr))
	} else {
		fmt.Fprintf(&b, "invoke %s.%s", relName(v.Value, instr), v.Method.Name())
	}
	b.WriteString("(")
	for i, arg := range v.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(relName(arg, instr))
	}
	if v.Signature().Variadic() {
		b.WriteString("...")
	}
	b.WriteString(")")
	return b.String()
}

func (c *CallCommon) String() string {
	return printCall(c, "", nil)
}

func (v *Call) String() string {
	return printCall(&v.Call, "", v)
}

func (v *BinOp) String() string {
	return fmt.Sprintf("%s %s %s", relName(v.X, v), v.Op.String(), relName(v.Y, v))
}

func (v *UnOp) String() string {
	return fmt.Sprintf("%s%s%s", v.Op, relName(v.X, v), commaOk(v.CommaOk))
}

func printConv(prefix string, v, x Value) string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("%s %s <- %s (%s)",
		prefix,
		relType(v.Type(), from),
		relType(x.Type(), from),
		relName(x, v.(Instruction)))
}

func (v *ChangeType) String() string          { return printConv("changetype", v, v.X) }
func (v *Convert) String() string             { return printConv("convert", v, v.X) }
func (v *ChangeInterface) String() string     { return printConv("change interface", v, v.X) }
func (v *SliceToArrayPointer) String() string { return printConv("slice to array pointer", v, v.X) }
func (v *MakeInterface) String() string       { return printConv("make", v, v.X) }

func (v *MultiConvert) String() string {
	from := v.Parent().relPkg()

	var b strings.Builder
	b.WriteString(printConv("multiconvert", v, v.X))
	b.WriteString(" [")
	for i, s := range termListOf(v.from) {
		for j, d := range termListOf(v.to) {
			if i != 0 || j != 0 {
				b.WriteString(" | ")
			}
			fmt.Fprintf(&b, "%s <- %s", relTerm(d, from), relTerm(s, from))
		}
	}
	b.WriteString("]")
	return b.String()
}

func (v *MakeClosure) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "make closure %s", relName(v.Fn, v))
	if v.Bindings != nil {
		b.WriteString(" [")
		for i, c := range v.Bindings {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(relName(c, v))
		}
		b.WriteString("]")
	}
	return b.String()
}

func (v *MakeSlice) String() string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("make %s %s %s",
		relType(v.Type(), from),
		relName(v.Len, v),
		relName(v.Cap, v))
}

func (v *Slice) String() string {
	var b bytes.Buffer
	b.WriteString("slice ")
	b.WriteString(relName(v.X, v))
	b.WriteString("[")
	if v.Low != nil {
		b.WriteString(relName(v.Low, v))
	}
	b.WriteString(":")
	if v.High != nil {
		b.WriteString(relName(v.High, v))
	}
	if v.Max != nil {
		b.WriteString(":")
		b.WriteString(relName(v.Max, v))
	}
	b.WriteString("]")
	return b.String()
}

func (v *MakeMap) String() string {
	res := ""
	if v.Reserve != nil {
		res = relName(v.Reserve, v)
	}
	from := v.Parent().relPkg()
	return fmt.Sprintf("make %s %s", relType(v.Type(), from), res)
}

func (v *MakeChan) String() string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("make %s %s", relType(v.Type(), from), relName(v.Size, v))
}

func (v *FieldAddr) String() string {
	// Be robust against a bad index.
	name := "?"
	if fld := fieldOf(typeparams.MustDeref(v.X.Type()), v.Field); fld != nil {
		name = fld.Name()
	}
	return fmt.Sprintf("&%s.%s [#%d]", relName(v.X, v), name, v.Field)
}

func (v *Field) String() string {
	// Be robust against a bad index.
	name := "?"
	if fld := fieldOf(v.X.Type(), v.Field); fld != nil {
		name = fld.Name()
	}
	return fmt.Sprintf("%s.%s [#%d]", relName(v.X, v), name, v.Field)
}

func (v *IndexAddr) String() string {
	return fmt.Sprintf("&%s[%s]", relName(v.X, v), relName(v.Index, v))
}

func (v *Index) String() string {
	return fmt.Sprintf("%s[%s]", relName(v.X, v), relName(v.Index, v))
}

func (v *Lookup) String() string {
	return fmt.Sprintf("%s[%s]%s", relName(v.X, v), relName(v.Index, v), commaOk(v.CommaOk))
}

func (v *Range) String() string {
	return "range " + relName(v.X, v)
}

func (v *Next) String() string {
	return "next " + relName(v.Iter, v)
}

func (v *TypeAssert) String() string {
	from := v.Parent().relPkg()
	return fmt.Sprintf("typeassert%s %s.(%s)", commaOk(v.CommaOk), relName(v.X, v), relType(v.AssertedType, from))
}

func (v *Extract) String() string {
	return fmt.Sprintf("extract %s #%d", relName(v.Tuple, v), v.Index)
}

func (s *Jump) String() string {
	// Be robust against malformed CFG.
	block := -1
	if s.block != nil && len(s.block.Succs) == 1 {
		block = s.block.Succs[0].Index
	}
	return fmt.Sprintf("jump %d", block)
}

func (s *If) String() string {
	// Be robust against malformed CFG.
	tblock, fblock := -1, -1
	if s.block != nil && len(s.block.Succs) == 2 {
		tblock = s.block.Succs[0].Index
		fblock = s.block.Succs[1].Index
	}
	return fmt.Sprintf("if %s goto %d else %d", relName(s.Cond, s), tblock, fblock)
}

func (s *Go) String() string {
	return printCall(&s.Call, "go ", s)
}

func (s *Panic) String() string {
	return "panic " + relName(s.X, s)
}

func (s *Return) String() string {
	var b bytes.Buffer
	b.WriteString("return")
	for i, r := range s.Results {
		if i == 0 {
			b.WriteString(" ")
		} else {
			b.WriteString(", ")
		}
		b.WriteString(relName(r, s))
	}
	return b.String()
}

func (*RunDefers) String() string {
	return "rundefers"
}

func (s *Send) String() string {
	return fmt.Sprintf("send %s <- %s", relName(s.Chan, s), relName(s.X, s))
}

func (s *Defer) String() string {
	prefix := "defer "
	if s.DeferStack != nil {
		prefix += "[" + relName(s.DeferStack, s) + "] "
	}
	c := printCall(&s.Call, prefix, s)
	return c
}

func (s *Select) String() string {
	var b bytes.Buffer
	for i, st := range s.States {
		if i > 0 {
			b.WriteString(", ")
		}
		if st.Dir == types.RecvOnly {
			b.WriteString("<-")
			b.WriteString(relName(st.Chan, s))
		} else {
			b.WriteString(relName(st.Chan, s))
			b.WriteString("<-")
			b.WriteString(relName(st.Send, s))
		}
	}
	non := ""
	if !s.Blocking {
		non = "non"
	}
	return fmt.Sprintf("select %sblocking [%s]", non, b.String())
}

func (s *Store) String() string {
	return fmt.Sprintf("*%s = %s", relName(s.Addr, s), relName(s.Val, s))
}

func (s *MapUpdate) String() string {
	return fmt.Sprintf("%s[%s] = %s", relName(s.Map, s), relName(s.Key, s), relName(s.Value, s))
}

func (s *DebugRef) String() string {
	p := s.Parent().Prog.Fset.Position(s.Pos())
	var descr any
	if s.object != nil {
		descr = s.object // e.g. "var x int"
	} else {
		descr = reflect.TypeOf(s.Expr) // e.g. "*ast.CallExpr"
	}
	var addr string
	if s.IsAddr {
		addr = "address of "
	}
	return fmt.Sprintf("; %s%s @ %d:%d is %s", addr, descr, p.Line, p.Column, s.X.Name())
}

func (p *Package) String() string {
	return "package " + p.Pkg.Path()
}

var _ io.WriterTo = (*Package)(nil) // *Package implements io.Writer

func (p *Package) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	WritePackage(&buf, p)
	n, err := w.Write(buf.Bytes())
	return int64(n), err
}

// WritePackage writes to buf a human-readable summary of p.
func WritePackage(buf *bytes.Buffer, p *Package) {
	fmt.Fprintf(buf, "%s:\n", p)

	var names []string
	maxname := 0
	for name := range p.Members {
		if l := len(name); l > maxname {
			maxname = l
		}
		names = append(names, name)
	}

	from := p.Pkg
	sort.Strings(names)
	for _, name := range names {
		switch mem := p.Members[name].(type) {
		case *NamedConst:
			fmt.Fprintf(buf, "  const %-*s %s = %s\n",
				maxname, name, mem.Name(), mem.Value.RelString(from))

		case *Function:
			fmt.Fprintf(buf, "  func  %-*s %s\n",
				maxname, name, relType(mem.Type(), from))

		case *Type:
			fmt.Fprintf(buf, "  type  %-*s %s\n",
				maxname, name, relType(mem.Type().Underlying(), from))
			for _, meth := range typeutil.IntuitiveMethodSet(mem.Type(), &p.Prog.MethodSets) {
				fmt.Fprintf(buf, "    %s\n", types.SelectionString(meth, types.RelativeTo(from)))
			}

		case *Global:
			fmt.Fprintf(buf, "  var   %-*s %s\n",
				maxname, name, relType(typeparams.MustDeref(mem.Type()), from))
		}
	}

	fmt.Fprintf(buf, "\n")
}

func commaOk(x bool) string {
	if x {
		return ",ok"
	}
	return ""
}
```

## File: go/ssa/sanity.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// An optional pass for sanity-checking invariants of the SSA representation.
// Currently it checks CFG invariants but little at the instruction level.

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"os"
	"strings"
)

type sanity struct {
	reporter io.Writer
	fn       *Function
	block    *BasicBlock
	instrs   map[Instruction]unit
	insane   bool
}

// sanityCheck performs integrity checking of the SSA representation
// of the function fn and returns true if it was valid.  Diagnostics
// are written to reporter if non-nil, os.Stderr otherwise.  Some
// diagnostics are only warnings and do not imply a negative result.
//
// Sanity-checking is intended to facilitate the debugging of code
// transformation passes.
func sanityCheck(fn *Function, reporter io.Writer) bool {
	if reporter == nil {
		reporter = os.Stderr
	}
	return (&sanity{reporter: reporter}).checkFunction(fn)
}

// mustSanityCheck is like sanityCheck but panics instead of returning
// a negative result.
func mustSanityCheck(fn *Function, reporter io.Writer) {
	if !sanityCheck(fn, reporter) {
		fn.WriteTo(os.Stderr)
		panic("SanityCheck failed")
	}
}

func (s *sanity) diagnostic(prefix, format string, args ...any) {
	fmt.Fprintf(s.reporter, "%s: function %s", prefix, s.fn)
	if s.block != nil {
		fmt.Fprintf(s.reporter, ", block %s", s.block)
	}
	io.WriteString(s.reporter, ": ")
	fmt.Fprintf(s.reporter, format, args...)
	io.WriteString(s.reporter, "\n")
}

func (s *sanity) errorf(format string, args ...any) {
	s.insane = true
	s.diagnostic("Error", format, args...)
}

func (s *sanity) warnf(format string, args ...any) {
	s.diagnostic("Warning", format, args...)
}

// findDuplicate returns an arbitrary basic block that appeared more
// than once in blocks, or nil if all were unique.
func findDuplicate(blocks []*BasicBlock) *BasicBlock {
	if len(blocks) < 2 {
		return nil
	}
	if blocks[0] == blocks[1] {
		return blocks[0]
	}
	// Slow path:
	m := make(map[*BasicBlock]bool)
	for _, b := range blocks {
		if m[b] {
			return b
		}
		m[b] = true
	}
	return nil
}

func (s *sanity) checkInstr(idx int, instr Instruction) {
	switch instr := instr.(type) {
	case *If, *Jump, *Return, *Panic:
		s.errorf("control flow instruction not at end of block")
	case *Phi:
		if idx == 0 {
			// It suffices to apply this check to just the first phi node.
			if dup := findDuplicate(s.block.Preds); dup != nil {
				s.errorf("phi node in block with duplicate predecessor %s", dup)
			}
		} else {
			prev := s.block.Instrs[idx-1]
			if _, ok := prev.(*Phi); !ok {
				s.errorf("Phi instruction follows a non-Phi: %T", prev)
			}
		}
		if ne, np := len(instr.Edges), len(s.block.Preds); ne != np {
			s.errorf("phi node has %d edges but %d predecessors", ne, np)

		} else {
			for i, e := range instr.Edges {
				if e == nil {
					s.errorf("phi node '%s' has no value for edge #%d from %s", instr.Comment, i, s.block.Preds[i])
				} else if !types.Identical(instr.typ, e.Type()) {
					s.errorf("phi node '%s' has a different type (%s) for edge #%d from %s (%s)",
						instr.Comment, instr.Type(), i, s.block.Preds[i], e.Type())
				}
			}
		}

	case *Alloc:
		if !instr.Heap {
			found := false
			for _, l := range s.fn.Locals {
				if l == instr {
					found = true
					break
				}
			}
			if !found {
				s.errorf("local alloc %s = %s does not appear in Function.Locals", instr.Name(), instr)
			}
		}

	case *BinOp:
	case *Call:
		if common := instr.Call; common.IsInvoke() {
			if !types.IsInterface(common.Value.Type()) {
				s.errorf("invoke on %s (%s) which is not an interface type (or type param)", common.Value, common.Value.Type())
			}
		}
	case *ChangeInterface:
	case *ChangeType:
	case *SliceToArrayPointer:
	case *Convert:
		if from := instr.X.Type(); !isBasicConvTypes(from) {
			if to := instr.Type(); !isBasicConvTypes(to) {
				s.errorf("convert %s -> %s: at least one type must be basic (or all basic, []byte, or []rune)", from, to)
			}
		}
	case *MultiConvert:
	case *Defer:
	case *Extract:
	case *Field:
	case *FieldAddr:
	case *Go:
	case *Index:
	case *IndexAddr:
	case *Lookup:
	case *MakeChan:
	case *MakeClosure:
		numFree := len(instr.Fn.(*Function).FreeVars)
		numBind := len(instr.Bindings)
		if numFree != numBind {
			s.errorf("MakeClosure has %d Bindings for function %s with %d free vars",
				numBind, instr.Fn, numFree)

		}
		if recv := instr.Type().(*types.Signature).Recv(); recv != nil {
			s.errorf("MakeClosure's type includes receiver %s", recv.Type())
		}

	case *MakeInterface:
	case *MakeMap:
	case *MakeSlice:
	case *MapUpdate:
	case *Next:
	case *Range:
	case *RunDefers:
	case *Select:
	case *Send:
	case *Slice:
	case *Store:
	case *TypeAssert:
	case *UnOp:
	case *DebugRef:
		// TODO(adonovan): implement checks.
	default:
		panic(fmt.Sprintf("Unknown instruction type: %T", instr))
	}

	if call, ok := instr.(CallInstruction); ok {
		if call.Common().Signature() == nil {
			s.errorf("nil signature: %s", call)
		}
	}

	// Check that value-defining instructions have valid types
	// and a valid referrer list.
	if v, ok := instr.(Value); ok {
		t := v.Type()
		if t == nil {
			s.errorf("no type: %s = %s", v.Name(), v)
		} else if t == tRangeIter || t == tDeferStack {
			// not a proper type; ignore.
		} else if b, ok := t.Underlying().(*types.Basic); ok && b.Info()&types.IsUntyped != 0 {
			s.errorf("instruction has 'untyped' result: %s = %s : %s", v.Name(), v, t)
		}
		s.checkReferrerList(v)
	}

	// Untyped constants are legal as instruction Operands(),
	// for example:
	//   _ = "foo"[0]
	// or:
	//   if wordsize==64 {...}

	// All other non-Instruction Values can be found via their
	// enclosing Function or Package.
}

func (s *sanity) checkFinalInstr(instr Instruction) {
	switch instr := instr.(type) {
	case *If:
		if nsuccs := len(s.block.Succs); nsuccs != 2 {
			s.errorf("If-terminated block has %d successors; expected 2", nsuccs)
			return
		}
		if s.block.Succs[0] == s.block.Succs[1] {
			s.errorf("If-instruction has same True, False target blocks: %s", s.block.Succs[0])
			return
		}

	case *Jump:
		if nsuccs := len(s.block.Succs); nsuccs != 1 {
			s.errorf("Jump-terminated block has %d successors; expected 1", nsuccs)
			return
		}

	case *Return:
		if nsuccs := len(s.block.Succs); nsuccs != 0 {
			s.errorf("Return-terminated block has %d successors; expected none", nsuccs)
			return
		}
		if na, nf := len(instr.Results), s.fn.Signature.Results().Len(); nf != na {
			s.errorf("%d-ary return in %d-ary function", na, nf)
		}

	case *Panic:
		if nsuccs := len(s.block.Succs); nsuccs != 0 {
			s.errorf("Panic-terminated block has %d successors; expected none", nsuccs)
			return
		}

	default:
		s.errorf("non-control flow instruction at end of block")
	}
}

func (s *sanity) checkBlock(b *BasicBlock, index int) {
	s.block = b

	if b.Index != index {
		s.errorf("block has incorrect Index %d", b.Index)
	}
	if b.parent != s.fn {
		s.errorf("block has incorrect parent %s", b.parent)
	}

	// Check all blocks are reachable.
	// (The entry block is always implicitly reachable,
	// as is the Recover block, if any.)
	if (index > 0 && b != b.parent.Recover) && len(b.Preds) == 0 {
		s.warnf("unreachable block")
		if b.Instrs == nil {
			// Since this block is about to be pruned,
			// tolerating transient problems in it
			// simplifies other optimizations.
			return
		}
	}

	// Check predecessor and successor relations are dual,
	// and that all blocks in CFG belong to same function.
	for _, a := range b.Preds {
		found := false
		for _, bb := range a.Succs {
			if bb == b {
				found = true
				break
			}
		}
		if !found {
			s.errorf("expected successor edge in predecessor %s; found only: %s", a, a.Succs)
		}
		if a.parent != s.fn {
			s.errorf("predecessor %s belongs to different function %s", a, a.parent)
		}
	}
	for _, c := range b.Succs {
		found := false
		for _, bb := range c.Preds {
			if bb == b {
				found = true
				break
			}
		}
		if !found {
			s.errorf("expected predecessor edge in successor %s; found only: %s", c, c.Preds)
		}
		if c.parent != s.fn {
			s.errorf("successor %s belongs to different function %s", c, c.parent)
		}
	}

	// Check each instruction is sane.
	n := len(b.Instrs)
	if n == 0 {
		s.errorf("basic block contains no instructions")
	}
	var rands [10]*Value // reuse storage
	for j, instr := range b.Instrs {
		if instr == nil {
			s.errorf("nil instruction at index %d", j)
			continue
		}
		if b2 := instr.Block(); b2 == nil {
			s.errorf("nil Block() for instruction at index %d", j)
			continue
		} else if b2 != b {
			s.errorf("wrong Block() (%s) for instruction at index %d ", b2, j)
			continue
		}
		if j < n-1 {
			s.checkInstr(j, instr)
		} else {
			s.checkFinalInstr(instr)
		}

		// Check Instruction.Operands.
	operands:
		for i, op := range instr.Operands(rands[:0]) {
			if op == nil {
				s.errorf("nil operand pointer %d of %s", i, instr)
				continue
			}
			val := *op
			if val == nil {
				continue // a nil operand is ok
			}

			// Check that "untyped" types only appear on constant operands.
			if _, ok := (*op).(*Const); !ok {
				if basic, ok := (*op).Type().Underlying().(*types.Basic); ok {
					if basic.Info()&types.IsUntyped != 0 {
						s.errorf("operand #%d of %s is untyped: %s", i, instr, basic)
					}
				}
			}

			// Check that Operands that are also Instructions belong to same function.
			// TODO(adonovan): also check their block dominates block b.
			if val, ok := val.(Instruction); ok {
				if val.Block() == nil {
					s.errorf("operand %d of %s is an instruction (%s) that belongs to no block", i, instr, val)
				} else if val.Parent() != s.fn {
					s.errorf("operand %d of %s is an instruction (%s) from function %s", i, instr, val, val.Parent())
				}
			}

			// Check that each function-local operand of
			// instr refers back to instr.  (NB: quadratic)
			switch val := val.(type) {
			case *Const, *Global, *Builtin:
				continue // not local
			case *Function:
				if val.parent == nil {
					continue // only anon functions are local
				}
			}

			// TODO(adonovan): check val.Parent() != nil <=> val.Referrers() is defined.

			if refs := val.Referrers(); refs != nil {
				for _, ref := range *refs {
					if ref == instr {
						continue operands
					}
				}
				s.errorf("operand %d of %s (%s) does not refer to us", i, instr, val)
			} else {
				s.errorf("operand %d of %s (%s) has no referrers", i, instr, val)
			}
		}
	}
}

func (s *sanity) checkReferrerList(v Value) {
	refs := v.Referrers()
	if refs == nil {
		s.errorf("%s has missing referrer list", v.Name())
		return
	}
	for i, ref := range *refs {
		if _, ok := s.instrs[ref]; !ok {
			s.errorf("%s.Referrers()[%d] = %s is not an instruction belonging to this function", v.Name(), i, ref)
		}
	}
}

func (s *sanity) checkFunctionParams() {
	signature := s.fn.Signature
	params := s.fn.Params

	// startSigParams is the start of signature.Params() within params.
	startSigParams := 0
	if signature.Recv() != nil {
		startSigParams = 1
	}

	if startSigParams+signature.Params().Len() != len(params) {
		s.errorf("function has %d parameters in signature but has %d after building",
			startSigParams+signature.Params().Len(), len(params))
		return
	}

	for i, param := range params {
		var sigType types.Type
		si := i - startSigParams
		if si < 0 {
			sigType = signature.Recv().Type()
		} else {
			sigType = signature.Params().At(si).Type()
		}

		if !types.Identical(sigType, param.Type()) {
			s.errorf("expect type %s in signature but got type %s in param %d", param.Type(), sigType, i)
		}
	}
}

// checkTransientFields checks whether all transient fields of Function are cleared.
func (s *sanity) checkTransientFields() {
	fn := s.fn
	if fn.build != nil {
		s.errorf("function transient field 'build' is not nil")
	}
	if fn.currentBlock != nil {
		s.errorf("function transient field 'currentBlock' is not nil")
	}
	if fn.vars != nil {
		s.errorf("function transient field 'vars' is not nil")
	}
	if fn.results != nil {
		s.errorf("function transient field 'results' is not nil")
	}
	if fn.returnVars != nil {
		s.errorf("function transient field 'returnVars' is not nil")
	}
	if fn.targets != nil {
		s.errorf("function transient field 'targets' is not nil")
	}
	if fn.lblocks != nil {
		s.errorf("function transient field 'lblocks' is not nil")
	}
	if fn.subst != nil {
		s.errorf("function transient field 'subst' is not nil")
	}
	if fn.jump != nil {
		s.errorf("function transient field 'jump' is not nil")
	}
	if fn.deferstack != nil {
		s.errorf("function transient field 'deferstack' is not nil")
	}
	if fn.source != nil {
		s.errorf("function transient field 'source' is not nil")
	}
	if fn.exits != nil {
		s.errorf("function transient field 'exits' is not nil")
	}
	if fn.uniq != 0 {
		s.errorf("function transient field 'uniq' is not zero")
	}
}

func (s *sanity) checkFunction(fn *Function) bool {
	s.fn = fn
	s.checkFunctionParams()
	s.checkTransientFields()

	// TODO(taking): Sanity check origin, typeparams, and typeargs.
	if fn.Prog == nil {
		s.errorf("nil Prog")
	}

	var buf bytes.Buffer
	_ = fn.String()               // must not crash
	_ = fn.RelString(fn.relPkg()) // must not crash
	WriteFunction(&buf, fn)       // must not crash

	// All functions have a package, except delegates (which are
	// shared across packages, or duplicated as weak symbols in a
	// separate-compilation model), and error.Error.
	if fn.Pkg == nil {
		if strings.HasPrefix(fn.Synthetic, "from type information (on demand)") ||
			strings.HasPrefix(fn.Synthetic, "wrapper ") ||
			strings.HasPrefix(fn.Synthetic, "bound ") ||
			strings.HasPrefix(fn.Synthetic, "thunk ") ||
			strings.HasSuffix(fn.name, "Error") ||
			strings.HasPrefix(fn.Synthetic, "instance ") ||
			strings.HasPrefix(fn.Synthetic, "instantiation ") ||
			(fn.parent != nil && len(fn.typeargs) > 0) /* anon fun in instance */ {
			// ok
		} else {
			s.errorf("nil Pkg")
		}
	}
	if src, syn := fn.Synthetic == "", fn.Syntax() != nil; src != syn {
		if len(fn.typeargs) > 0 && fn.Prog.mode&InstantiateGenerics != 0 {
			// ok (instantiation with InstantiateGenerics on)
		} else if fn.topLevelOrigin != nil && len(fn.typeargs) > 0 {
			// ok (we always have the syntax set for instantiation)
		} else if _, rng := fn.syntax.(*ast.RangeStmt); rng && fn.Synthetic == "range-over-func yield" {
			// ok (range-func-yields are both synthetic and keep syntax)
		} else {
			s.errorf("got fromSource=%t, hasSyntax=%t; want same values", src, syn)
		}
	}

	// Build the set of valid referrers.
	s.instrs = make(map[Instruction]unit)

	// instrs are the instructions that are present in the function.
	for instr := range fn.instrs() {
		s.instrs[instr] = unit{}
	}

	// Check all Locals allocations appear in the function instruction.
	for i, l := range fn.Locals {
		if _, present := s.instrs[l]; !present {
			s.warnf("function doesn't contain Local alloc %s", l.Name())
		}

		if l.Parent() != fn {
			s.errorf("Local %s at index %d has wrong parent", l.Name(), i)
		}
		if l.Heap {
			s.errorf("Local %s at index %d has Heap flag set", l.Name(), i)
		}
	}
	for i, p := range fn.Params {
		if p.Parent() != fn {
			s.errorf("Param %s at index %d has wrong parent", p.Name(), i)
		}
		// Check common suffix of Signature and Params match type.
		if sig := fn.Signature; sig != nil {
			j := i - len(fn.Params) + sig.Params().Len() // index within sig.Params
			if j < 0 {
				continue
			}
			if !types.Identical(p.Type(), sig.Params().At(j).Type()) {
				s.errorf("Param %s at index %d has wrong type (%s, versus %s in Signature)", p.Name(), i, p.Type(), sig.Params().At(j).Type())

			}
		}
		s.checkReferrerList(p)
	}
	for i, fv := range fn.FreeVars {
		if fv.Parent() != fn {
			s.errorf("FreeVar %s at index %d has wrong parent", fv.Name(), i)
		}
		s.checkReferrerList(fv)
	}

	if fn.Blocks != nil && len(fn.Blocks) == 0 {
		// Function _had_ blocks (so it's not external) but
		// they were "optimized" away, even the entry block.
		s.errorf("Blocks slice is non-nil but empty")
	}
	for i, b := range fn.Blocks {
		if b == nil {
			s.warnf("nil *BasicBlock at f.Blocks[%d]", i)
			continue
		}
		s.checkBlock(b, i)
	}
	if fn.Recover != nil && fn.Blocks[fn.Recover.Index] != fn.Recover {
		s.errorf("Recover block is not in Blocks slice")
	}

	s.block = nil
	for i, anon := range fn.AnonFuncs {
		if anon.Parent() != fn {
			s.errorf("AnonFuncs[%d]=%s but %s.Parent()=%s", i, anon, anon, anon.Parent())
		}
		if i != int(anon.anonIdx) {
			s.errorf("AnonFuncs[%d]=%s but %s.anonIdx=%d", i, anon, anon, anon.anonIdx)
		}
	}
	s.fn = nil
	return !s.insane
}

// sanityCheckPackage checks invariants of packages upon creation.
// It does not require that the package is built.
// Unlike sanityCheck (for functions), it just panics at the first error.
func sanityCheckPackage(pkg *Package) {
	if pkg.Pkg == nil {
		panic(fmt.Sprintf("Package %s has no Object", pkg))
	}
	if pkg.info != nil {
		panic(fmt.Sprintf("package %s field 'info' is not cleared", pkg))
	}
	if pkg.files != nil {
		panic(fmt.Sprintf("package %s field 'files' is not cleared", pkg))
	}
	if pkg.created != nil {
		panic(fmt.Sprintf("package %s field 'created' is not cleared", pkg))
	}
	if pkg.initVersion != nil {
		panic(fmt.Sprintf("package %s field 'initVersion' is not cleared", pkg))
	}

	_ = pkg.String() // must not crash

	for name, mem := range pkg.Members {
		if name != mem.Name() {
			panic(fmt.Sprintf("%s: %T.Name() = %s, want %s",
				pkg.Pkg.Path(), mem, mem.Name(), name))
		}
		obj := mem.Object()
		if obj == nil {
			// This check is sound because fields
			// {Global,Function}.object have type
			// types.Object.  (If they were declared as
			// *types.{Var,Func}, we'd have a non-empty
			// interface containing a nil pointer.)

			continue // not all members have typechecker objects
		}
		if obj.Name() != name {
			if obj.Name() == "init" && strings.HasPrefix(mem.Name(), "init#") {
				// Ok.  The name of a declared init function varies between
				// its types.Func ("init") and its ssa.Function ("init#%d").
			} else {
				panic(fmt.Sprintf("%s: %T.Object().Name() = %s, want %s",
					pkg.Pkg.Path(), mem, obj.Name(), name))
			}
		}
		if obj.Pos() != mem.Pos() {
			panic(fmt.Sprintf("%s Pos=%d obj.Pos=%d", mem, mem.Pos(), obj.Pos()))
		}
	}
}
```

## File: go/ssa/source_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa_test

// This file defines tests of source-level debugging utilities.

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"os"
	"runtime"
	"strings"
	"testing"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/internal/expect"
)

func TestObjValueLookup(t *testing.T) {
	if runtime.GOOS == "android" {
		t.Skipf("no testdata directory on %s", runtime.GOOS)
	}

	src, err := os.ReadFile("testdata/objlookup.go")
	if err != nil {
		t.Fatal(err)
	}
	readFile := func(filename string) ([]byte, error) { return src, nil }

	mode := ssa.GlobalDebug /*|ssa.PrintFunctions*/
	mainPkg, ppkg := buildPackage(t, string(src), mode)
	fset := ppkg.Fset

	// Maps each var Ident (represented "name:linenum") to the
	// kind of ssa.Value we expect (represented "Constant", "&Alloc").
	expectations := make(map[string]string)

	// Each note of the form @ssa(x, "BinOp") in testdata/objlookup.go
	// specifies an expectation that an object named x declared on the
	// same line is associated with an ssa.Value of type *ssa.BinOp.
	notes, err := expect.ExtractGo(fset, ppkg.Syntax[0])
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range notes {
		if n.Name != "ssa" {
			t.Errorf("%v: unexpected note type %q, want \"ssa\"", fset.Position(n.Pos), n.Name)
			continue
		}
		if len(n.Args) != 2 {
			t.Errorf("%v: ssa has %d args, want 2", fset.Position(n.Pos), len(n.Args))
			continue
		}
		ident, ok := n.Args[0].(expect.Identifier)
		if !ok {
			t.Errorf("%v: got %v for arg 1, want identifier", fset.Position(n.Pos), n.Args[0])
			continue
		}
		exp, ok := n.Args[1].(string)
		if !ok {
			t.Errorf("%v: got %v for arg 2, want string", fset.Position(n.Pos), n.Args[1])
			continue
		}
		p, _, err := expect.MatchBefore(fset, readFile, n.Pos, string(ident))
		if err != nil {
			t.Error(err)
			continue
		}
		pos := fset.Position(p)
		key := fmt.Sprintf("%s:%d", ident, pos.Line)
		expectations[key] = exp
	}

	var varIds []*ast.Ident
	var varObjs []*types.Var
	for id, obj := range ppkg.TypesInfo.Defs {
		// Check invariants for func and const objects.
		switch obj := obj.(type) {
		case *types.Func:
			checkFuncValue(t, mainPkg.Prog, obj)

		case *types.Const:
			checkConstValue(t, mainPkg.Prog, obj)

		case *types.Var:
			if id.Name == "_" {
				continue
			}
			varIds = append(varIds, id)
			varObjs = append(varObjs, obj)
		}
	}
	for id, obj := range ppkg.TypesInfo.Uses {
		if obj, ok := obj.(*types.Var); ok {
			varIds = append(varIds, id)
			varObjs = append(varObjs, obj)
		}
	}

	// Check invariants for var objects.
	// The result varies based on the specific Ident.
	for i, id := range varIds {
		obj := varObjs[i]
		ref, _ := astutil.PathEnclosingInterval(ppkg.Syntax[0], id.Pos(), id.Pos())
		pos := fset.Position(id.Pos())
		exp := expectations[fmt.Sprintf("%s:%d", id.Name, pos.Line)]
		if exp == "" {
			t.Errorf("%s: no expectation for var ident %s ", pos, id.Name)
			continue
		}
		wantAddr := false
		if exp[0] == '&' {
			wantAddr = true
			exp = exp[1:]
		}
		checkVarValue(t, mainPkg, ref, obj, exp, wantAddr)
	}
}

func checkFuncValue(t *testing.T, prog *ssa.Program, obj *types.Func) {
	fn := prog.FuncValue(obj)
	// fmt.Printf("FuncValue(%s) = %s\n", obj, fn) // debugging
	if fn == nil {
		if obj.Name() != "interfaceMethod" {
			t.Errorf("FuncValue(%s) == nil", obj)
		}
		return
	}
	if fnobj := fn.Object(); fnobj != obj {
		t.Errorf("FuncValue(%s).Object() == %s; value was %s",
			obj, fnobj, fn.Name())
		return
	}
	if !types.Identical(fn.Type(), obj.Type()) {
		t.Errorf("FuncValue(%s).Type() == %s", obj, fn.Type())
		return
	}
}

func checkConstValue(t *testing.T, prog *ssa.Program, obj *types.Const) {
	c := prog.ConstValue(obj)
	// fmt.Printf("ConstValue(%s) = %s\n", obj, c) // debugging
	if c == nil {
		t.Errorf("ConstValue(%s) == nil", obj)
		return
	}
	if !types.Identical(c.Type(), obj.Type()) {
		t.Errorf("ConstValue(%s).Type() == %s", obj, c.Type())
		return
	}
	if obj.Name() != "nil" {
		if !constant.Compare(c.Value, token.EQL, obj.Val()) {
			t.Errorf("ConstValue(%s).Value (%s) != %s",
				obj, c.Value, obj.Val())
			return
		}
	}
}

func checkVarValue(t *testing.T, pkg *ssa.Package, ref []ast.Node, obj *types.Var, expKind string, wantAddr bool) {
	// The prefix of all assertions messages.
	prefix := fmt.Sprintf("VarValue(%s @ L%d)",
		obj, pkg.Prog.Fset.Position(ref[0].Pos()).Line)

	v, gotAddr := pkg.Prog.VarValue(obj, pkg, ref)

	// Kind is the concrete type of the ssa Value.
	gotKind := "nil"
	if v != nil {
		gotKind = fmt.Sprintf("%T", v)[len("*ssa."):]
	}

	// fmt.Printf("%s = %v (kind %q; expect %q) wantAddr=%t gotAddr=%t\n", prefix, v, gotKind, expKind, wantAddr, gotAddr) // debugging

	// Check the kinds match.
	// "nil" indicates expected failure (e.g. optimized away).
	if expKind != gotKind {
		t.Errorf("%s concrete type == %s, want %s", prefix, gotKind, expKind)
	}

	// Check the types match.
	// If wantAddr, the expected type is the object's address.
	if v != nil {
		expType := obj.Type()
		if wantAddr {
			expType = types.NewPointer(expType)
			if !gotAddr {
				t.Errorf("%s: got value, want address", prefix)
			}
		} else if gotAddr {
			t.Errorf("%s: got address, want value", prefix)
		}
		if !types.Identical(v.Type(), expType) {
			t.Errorf("%s.Type() == %s, want %s", prefix, v.Type(), expType)
		}
	}
}

// Ensure that, in debug mode, we can determine the ssa.Value
// corresponding to every ast.Expr.
func TestValueForExpr(t *testing.T) {
	testValueForExpr(t, "testdata/valueforexpr.go")
}

func TestValueForExprStructConv(t *testing.T) {
	testValueForExpr(t, "testdata/structconv.go")
}

func testValueForExpr(t *testing.T, testfile string) {
	if runtime.GOOS == "android" {
		t.Skipf("no testdata dir on %s", runtime.GOOS)
	}

	src, err := os.ReadFile(testfile)
	if err != nil {
		t.Fatal(err)
	}

	mode := ssa.GlobalDebug /*|ssa.PrintFunctions*/
	mainPkg, ppkg := buildPackage(t, string(src), mode)
	fset, file := ppkg.Fset, ppkg.Syntax[0]

	if false {
		// debugging
		for _, mem := range mainPkg.Members {
			if fn, ok := mem.(*ssa.Function); ok {
				fn.WriteTo(os.Stderr)
			}
		}
	}

	var parenExprs []*ast.ParenExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if n != nil {
			if e, ok := n.(*ast.ParenExpr); ok {
				parenExprs = append(parenExprs, e)
			}
		}
		return true
	})

	notes, err := expect.ExtractGo(fset, file)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range notes {
		want := n.Name
		if want == "nil" {
			want = "<nil>"
		}
		position := fset.Position(n.Pos)
		var e ast.Expr
		for _, paren := range parenExprs {
			if paren.Pos() > n.Pos {
				e = paren.X
				break
			}
		}
		if e == nil {
			t.Errorf("%s: note doesn't precede ParenExpr: %q", position, want)
			continue
		}

		path, _ := astutil.PathEnclosingInterval(file, n.Pos, n.Pos)
		if path == nil {
			t.Errorf("%s: can't find AST path from root to comment: %s", position, want)
			continue
		}

		fn := ssa.EnclosingFunction(mainPkg, path)
		if fn == nil {
			t.Errorf("%s: can't find enclosing function", position)
			continue
		}

		v, gotAddr := fn.ValueForExpr(e) // (may be nil)
		got := strings.TrimPrefix(fmt.Sprintf("%T", v), "*ssa.")
		if got != want {
			t.Errorf("%s: got value %q, want %q", position, got, want)
		}
		if v != nil {
			T := v.Type()
			if gotAddr {
				T = T.Underlying().(*types.Pointer).Elem() // deref
			}
			if etyp := ppkg.TypesInfo.TypeOf(e); !types.Identical(T, etyp) {
				t.Errorf("%s: got type %s, want %s", position, etyp, T)
			}
		}
	}
}

func TestEnclosingFunction(t *testing.T) {
	tests := []struct {
		desc   string
		input  string // the input file
		substr string // first occurrence of this string denotes interval
		fn     string // name of expected containing function
	}{
		// We use distinctive numbers as syntactic landmarks.
		{"Ordinary function", `
		  package main
		  func f() { println(1003) }`,
			"100", "main.f"},
		{"Methods", `
		  package main
          type T int
		  func (t T) f() { println(200) }`,
			"200", "(main.T).f"},
		{"Function literal", `
		  package main
		  func f() { println(func() { print(300) }) }`,
			"300", "main.f$1"},
		{"Doubly nested", `
		  package main
		  func f() { println(func() { print(func() { print(350) })})}`,
			"350", "main.f$1$1"},
		{"Implicit init for package-level var initializer", `
		  package main; var a = 400`,
			"400", "main.init"},
		{"No code for constants", "package main; const a = 500", "500", "(none)"},
		{" Explicit init", "package main; func init() { println(600) }", "600", "main.init#1"},
		{"Multiple explicit init functions", `
		  package main
		  func init() { println("foo") }
		  func init() { println(800) }`,
			"800", "main.init#2"},
		{"init containing FuncLit", `
		  package main
		  func init() { println(func(){print(900)}) }`,
			"900", "main.init#1$1"},
		{"generic", `
		    package main
			type S[T any] struct{}
			func (*S[T]) Foo() { println(1000) }
			type P[T any] struct{ *S[T] }`,
			"1000", "(*main.S[T]).Foo",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			pkg, ppkg := buildPackage(t, test.input, ssa.BuilderMode(0))
			fset, file := ppkg.Fset, ppkg.Syntax[0]

			// Find [start,end) positions of the first occurrence of substr in file.
			index := strings.Index(test.input, test.substr)
			if index < 0 {
				t.Fatalf("%q is not a substring of input", test.substr)
			}
			filePos := fset.File(file.Package)
			start, end := filePos.Pos(index), filePos.Pos(index+len(test.substr))

			path, exact := astutil.PathEnclosingInterval(file, start, end)
			if !exact {
				t.Fatalf("PathEnclosingInterval(%q) not exact", test.substr)
			}

			name := "(none)"
			fn := ssa.EnclosingFunction(pkg, path)
			if fn != nil {
				name = fn.String()
			}

			if name != test.fn {
				t.Errorf("EnclosingFunction(%q in %q) got %s, want %s",
					test.substr, test.input, name, test.fn)
			}

			// While we're here: test HasEnclosingFunction.
			if has := ssa.HasEnclosingFunction(pkg, path); has != (fn != nil) {
				t.Errorf("HasEnclosingFunction(%q in %q) got %v, want %v",
					test.substr, test.input, has, fn != nil)
			}
		})
	}
}
```

## File: go/ssa/source.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines utilities for working with source positions
// or source-level named entities ("objects").

// TODO(adonovan): test that {Value,Instruction}.Pos() positions match
// the originating syntax, as specified.

import (
	"go/ast"
	"go/token"
	"go/types"
)

// EnclosingFunction returns the function that contains the syntax
// node denoted by path.
//
// Syntax associated with package-level variable specifications is
// enclosed by the package's init() function.
//
// Returns nil if not found; reasons might include:
//   - the node is not enclosed by any function.
//   - the node is within an anonymous function (FuncLit) and
//     its SSA function has not been created yet
//     (pkg.Build() has not yet been called).
func EnclosingFunction(pkg *Package, path []ast.Node) *Function {
	// Start with package-level function...
	fn := findEnclosingPackageLevelFunction(pkg, path)
	if fn == nil {
		return nil // not in any function
	}

	// ...then walk down the nested anonymous functions.
	n := len(path)
outer:
	for i := range path {
		if lit, ok := path[n-1-i].(*ast.FuncLit); ok {
			for _, anon := range fn.AnonFuncs {
				if anon.Pos() == lit.Type.Func {
					fn = anon
					continue outer
				}
			}
			// SSA function not found:
			// - package not yet built, or maybe
			// - builder skipped FuncLit in dead block
			//   (in principle; but currently the Builder
			//   generates even dead FuncLits).
			return nil
		}
	}
	return fn
}

// HasEnclosingFunction returns true if the AST node denoted by path
// is contained within the declaration of some function or
// package-level variable.
//
// Unlike EnclosingFunction, the behaviour of this function does not
// depend on whether SSA code for pkg has been built, so it can be
// used to quickly reject check inputs that will cause
// EnclosingFunction to fail, prior to SSA building.
func HasEnclosingFunction(pkg *Package, path []ast.Node) bool {
	return findEnclosingPackageLevelFunction(pkg, path) != nil
}

// findEnclosingPackageLevelFunction returns the Function
// corresponding to the package-level function enclosing path.
func findEnclosingPackageLevelFunction(pkg *Package, path []ast.Node) *Function {
	if n := len(path); n >= 2 { // [... {Gen,Func}Decl File]
		switch decl := path[n-2].(type) {
		case *ast.GenDecl:
			if decl.Tok == token.VAR && n >= 3 {
				// Package-level 'var' initializer.
				return pkg.init
			}

		case *ast.FuncDecl:
			if decl.Recv == nil && decl.Name.Name == "init" {
				// Explicit init() function.
				for _, b := range pkg.init.Blocks {
					for _, instr := range b.Instrs {
						if instr, ok := instr.(*Call); ok {
							if callee, ok := instr.Call.Value.(*Function); ok && callee.Pkg == pkg && callee.Pos() == decl.Name.NamePos {
								return callee
							}
						}
					}
				}
				// Hack: return non-nil when SSA is not yet
				// built so that HasEnclosingFunction works.
				return pkg.init
			}
			// Declared function/method.
			return findNamedFunc(pkg, decl.Name.NamePos)
		}
	}
	return nil // not in any function
}

// findNamedFunc returns the named function whose FuncDecl.Ident is at
// position pos.
func findNamedFunc(pkg *Package, pos token.Pos) *Function {
	// Look at all package members and method sets of named types.
	// Not very efficient.
	for _, mem := range pkg.Members {
		switch mem := mem.(type) {
		case *Function:
			if mem.Pos() == pos {
				return mem
			}
		case *Type:
			mset := pkg.Prog.MethodSets.MethodSet(types.NewPointer(mem.Type()))
			for i, n := 0, mset.Len(); i < n; i++ {
				// Don't call Program.Method: avoid creating wrappers.
				obj := mset.At(i).Obj().(*types.Func)
				if obj.Pos() == pos {
					// obj from MethodSet may not be the origin type.
					m := obj.Origin()
					return pkg.objects[m].(*Function)
				}
			}
		}
	}
	return nil
}

// ValueForExpr returns the SSA Value that corresponds to non-constant
// expression e.
//
// It returns nil if no value was found, e.g.
//   - the expression is not lexically contained within f;
//   - f was not built with debug information; or
//   - e is a constant expression.  (For efficiency, no debug
//     information is stored for constants. Use
//     go/types.Info.Types[e].Value instead.)
//   - e is a reference to nil or a built-in function.
//   - the value was optimised away.
//
// If e is an addressable expression used in an lvalue context,
// value is the address denoted by e, and isAddr is true.
//
// The types of e (or &e, if isAddr) and the result are equal
// (modulo "untyped" bools resulting from comparisons).
//
// (Tip: to find the ssa.Value given a source position, use
// astutil.PathEnclosingInterval to locate the ast.Node, then
// EnclosingFunction to locate the Function, then ValueForExpr to find
// the ssa.Value.)
func (f *Function) ValueForExpr(e ast.Expr) (value Value, isAddr bool) {
	if f.debugInfo() { // (opt)
		e = ast.Unparen(e)
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				if ref, ok := instr.(*DebugRef); ok {
					if ref.Expr == e {
						return ref.X, ref.IsAddr
					}
				}
			}
		}
	}
	return
}

// --- Lookup functions for source-level named entities (types.Objects) ---

// Package returns the SSA Package corresponding to the specified
// type-checker package. It returns nil if no such Package was
// created by a prior call to prog.CreatePackage.
func (prog *Program) Package(pkg *types.Package) *Package {
	return prog.packages[pkg]
}

// packageLevelMember returns the package-level member corresponding
// to the specified symbol, which may be a package-level const
// (*NamedConst), var (*Global) or func/method (*Function) of some
// package in prog.
//
// It returns nil if the object belongs to a package that has not been
// created by prog.CreatePackage.
func (prog *Program) packageLevelMember(obj types.Object) Member {
	if pkg, ok := prog.packages[obj.Pkg()]; ok {
		return pkg.objects[obj]
	}
	return nil
}

// FuncValue returns the SSA function or (non-interface) method
// denoted by the specified func symbol. It returns nil if the symbol
// denotes an interface method, or belongs to a package that was not
// created by prog.CreatePackage.
func (prog *Program) FuncValue(obj *types.Func) *Function {
	fn, _ := prog.packageLevelMember(obj).(*Function)
	return fn
}

// ConstValue returns the SSA constant denoted by the specified const symbol.
func (prog *Program) ConstValue(obj *types.Const) *Const {
	// TODO(adonovan): opt: share (don't reallocate)
	// Consts for const objects and constant ast.Exprs.

	// Universal constant? {true,false,nil}
	if obj.Parent() == types.Universe {
		return NewConst(obj.Val(), obj.Type())
	}
	// Package-level named constant?
	if v := prog.packageLevelMember(obj); v != nil {
		return v.(*NamedConst).Value
	}
	return NewConst(obj.Val(), obj.Type())
}

// VarValue returns the SSA Value that corresponds to a specific
// identifier denoting the specified var symbol.
//
// VarValue returns nil if a local variable was not found, perhaps
// because its package was not built, the debug information was not
// requested during SSA construction, or the value was optimized away.
//
// ref is the path to an ast.Ident (e.g. from PathEnclosingInterval),
// and that ident must resolve to obj.
//
// pkg is the package enclosing the reference.  (A reference to a var
// always occurs within a function, so we need to know where to find it.)
//
// If the identifier is a field selector and its base expression is
// non-addressable, then VarValue returns the value of that field.
// For example:
//
//	func f() struct {x int}
//	f().x  // VarValue(x) returns a *Field instruction of type int
//
// All other identifiers denote addressable locations (variables).
// For them, VarValue may return either the variable's address or its
// value, even when the expression is evaluated only for its value; the
// situation is reported by isAddr, the second component of the result.
//
// If !isAddr, the returned value is the one associated with the
// specific identifier.  For example,
//
//	var x int    // VarValue(x) returns Const 0 here
//	x = 1        // VarValue(x) returns Const 1 here
//
// It is not specified whether the value or the address is returned in
// any particular case, as it may depend upon optimizations performed
// during SSA code generation, such as registerization, constant
// folding, avoidance of materialization of subexpressions, etc.
func (prog *Program) VarValue(obj *types.Var, pkg *Package, ref []ast.Node) (value Value, isAddr bool) {
	// All references to a var are local to some function, possibly init.
	fn := EnclosingFunction(pkg, ref)
	if fn == nil {
		return // e.g. def of struct field; SSA not built?
	}

	id := ref[0].(*ast.Ident)

	// Defining ident of a parameter?
	if id.Pos() == obj.Pos() {
		for _, param := range fn.Params {
			if param.Object() == obj {
				return param, false
			}
		}
	}

	// Other ident?
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			if dr, ok := instr.(*DebugRef); ok {
				if dr.Pos() == id.Pos() {
					return dr.X, dr.IsAddr
				}
			}
		}
	}

	// Defining ident of package-level var?
	if v := prog.packageLevelMember(obj); v != nil {
		return v.(*Global), true
	}

	return // e.g. debug info not requested, or var optimized away
}
```

## File: go/ssa/ssa.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This package defines a high-level intermediate representation for
// Go programs using static single-assignment (SSA) form.

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"sync"

	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/typeparams"
)

// A Program is a partial or complete Go program converted to SSA form.
type Program struct {
	Fset       *token.FileSet              // position information for the files of this Program
	imported   map[string]*Package         // all importable Packages, keyed by import path
	packages   map[*types.Package]*Package // all created Packages
	mode       BuilderMode                 // set of mode bits for SSA construction
	MethodSets typeutil.MethodSetCache     // cache of type-checker's method-sets

	canon *canonizer     // type canonicalization map
	ctxt  *types.Context // cache for type checking instantiations

	methodsMu  sync.Mutex
	methodSets typeutil.Map // maps type to its concrete *methodSet

	// memoization of whether a type refers to type parameters
	hasParamsMu sync.Mutex
	hasParams   typeparams.Free

	// set of concrete types used as MakeInterface operands
	makeInterfaceTypesMu sync.Mutex
	makeInterfaceTypes   map[types.Type]unit // (may contain redundant identical types)

	// objectMethods is a memoization of objectMethod
	// to avoid creation of duplicate methods from type information.
	objectMethodsMu sync.Mutex
	objectMethods   map[*types.Func]*Function
}

// A Package is a single analyzed Go package containing Members for
// all package-level functions, variables, constants and types it
// declares.  These may be accessed directly via Members, or via the
// type-specific accessor methods Func, Type, Var and Const.
//
// Members also contains entries for "init" (the synthetic package
// initializer) and "init#%d", the nth declared init function,
// and unspecified other things too.
type Package struct {
	Prog    *Program                // the owning program
	Pkg     *types.Package          // the corresponding go/types.Package
	Members map[string]Member       // all package members keyed by name (incl. init and init#%d)
	objects map[types.Object]Member // mapping of package objects to members (incl. methods). Contains *NamedConst, *Global, *Function (values but not types)
	init    *Function               // Func("init"); the package's init function
	debug   bool                    // include full debug info in this package
	syntax  bool                    // package was loaded from syntax

	// The following fields are set transiently, then cleared
	// after building.
	buildOnce   sync.Once           // ensures package building occurs once
	ninit       int32               // number of init functions
	info        *types.Info         // package type information
	files       []*ast.File         // package ASTs
	created     []*Function         // members created as a result of building this package (includes declared functions, wrappers)
	initVersion map[ast.Expr]string // goversion to use for each global var init expr
}

// A Member is a member of a Go package, implemented by *NamedConst,
// *Global, *Function, or *Type; they are created by package-level
// const, var, func and type declarations respectively.
type Member interface {
	Name() string                    // declared name of the package member
	String() string                  // package-qualified name of the package member
	RelString(*types.Package) string // like String, but relative refs are unqualified
	Object() types.Object            // typechecker's object for this member, if any
	Pos() token.Pos                  // position of member's declaration, if known
	Type() types.Type                // type of the package member
	Token() token.Token              // token.{VAR,FUNC,CONST,TYPE}
	Package() *Package               // the containing package
}

// A Type is a Member of a Package representing a package-level named type.
type Type struct {
	object *types.TypeName
	pkg    *Package
}

// A NamedConst is a Member of a Package representing a package-level
// named constant.
//
// Pos() returns the position of the declaring ast.ValueSpec.Names[*]
// identifier.
//
// NB: a NamedConst is not a Value; it contains a constant Value, which
// it augments with the name and position of its 'const' declaration.
type NamedConst struct {
	object *types.Const
	Value  *Const
	pkg    *Package
}

// A Value is an SSA value that can be referenced by an instruction.
type Value interface {
	// Name returns the name of this value, and determines how
	// this Value appears when used as an operand of an
	// Instruction.
	//
	// This is the same as the source name for Parameters,
	// Builtins, Functions, FreeVars, Globals.
	// For constants, it is a representation of the constant's value
	// and type.  For all other Values this is the name of the
	// virtual register defined by the instruction.
	//
	// The name of an SSA Value is not semantically significant,
	// and may not even be unique within a function.
	Name() string

	// If this value is an Instruction, String returns its
	// disassembled form; otherwise it returns unspecified
	// human-readable information about the Value, such as its
	// kind, name and type.
	String() string

	// Type returns the type of this value.  Many instructions
	// (e.g. IndexAddr) change their behaviour depending on the
	// types of their operands.
	Type() types.Type

	// Parent returns the function to which this Value belongs.
	// It returns nil for named Functions, Builtin, Const and Global.
	Parent() *Function

	// Referrers returns the list of instructions that have this
	// value as one of their operands; it may contain duplicates
	// if an instruction has a repeated operand.
	//
	// Referrers actually returns a pointer through which the
	// caller may perform mutations to the object's state.
	//
	// Referrers is currently only defined if Parent()!=nil,
	// i.e. for the function-local values FreeVar, Parameter,
	// Functions (iff anonymous) and all value-defining instructions.
	// It returns nil for named Functions, Builtin, Const and Global.
	//
	// Instruction.Operands contains the inverse of this relation.
	Referrers() *[]Instruction

	// Pos returns the location of the AST token most closely
	// associated with the operation that gave rise to this value,
	// or token.NoPos if it was not explicit in the source.
	//
	// For each ast.Node type, a particular token is designated as
	// the closest location for the expression, e.g. the Lparen
	// for an *ast.CallExpr.  This permits a compact but
	// approximate mapping from Values to source positions for use
	// in diagnostic messages, for example.
	//
	// (Do not use this position to determine which Value
	// corresponds to an ast.Expr; use Function.ValueForExpr
	// instead.  NB: it requires that the function was built with
	// debug information.)
	Pos() token.Pos
}

// An Instruction is an SSA instruction that computes a new Value or
// has some effect.
//
// An Instruction that defines a value (e.g. BinOp) also implements
// the Value interface; an Instruction that only has an effect (e.g. Store)
// does not.
type Instruction interface {
	// String returns the disassembled form of this value.
	//
	// Examples of Instructions that are Values:
	//       "x + y"     (BinOp)
	//       "len([])"   (Call)
	// Note that the name of the Value is not printed.
	//
	// Examples of Instructions that are not Values:
	//       "return x"  (Return)
	//       "*y = x"    (Store)
	//
	// (The separation Value.Name() from Value.String() is useful
	// for some analyses which distinguish the operation from the
	// value it defines, e.g., 'y = local int' is both an allocation
	// of memory 'local int' and a definition of a pointer y.)
	String() string

	// Parent returns the function to which this instruction
	// belongs.
	Parent() *Function

	// Block returns the basic block to which this instruction
	// belongs.
	Block() *BasicBlock

	// setBlock sets the basic block to which this instruction belongs.
	setBlock(*BasicBlock)

	// Operands returns the operands of this instruction: the
	// set of Values it references.
	//
	// Specifically, it appends their addresses to rands, a
	// user-provided slice, and returns the resulting slice,
	// permitting avoidance of memory allocation.
	//
	// The operands are appended in undefined order, but the order
	// is consistent for a given Instruction; the addresses are
	// always non-nil but may point to a nil Value.  Clients may
	// store through the pointers, e.g. to effect a value
	// renaming.
	//
	// Value.Referrers is a subset of the inverse of this
	// relation.  (Referrers are not tracked for all types of
	// Values.)
	Operands(rands []*Value) []*Value

	// Pos returns the location of the AST token most closely
	// associated with the operation that gave rise to this
	// instruction, or token.NoPos if it was not explicit in the
	// source.
	//
	// For each ast.Node type, a particular token is designated as
	// the closest location for the expression, e.g. the Go token
	// for an *ast.GoStmt.  This permits a compact but approximate
	// mapping from Instructions to source positions for use in
	// diagnostic messages, for example.
	//
	// (Do not use this position to determine which Instruction
	// corresponds to an ast.Expr; see the notes for Value.Pos.
	// This position may be used to determine which non-Value
	// Instruction corresponds to some ast.Stmts, but not all: If
	// and Jump instructions have no Pos(), for example.)
	Pos() token.Pos
}

// A Node is a node in the SSA value graph.  Every concrete type that
// implements Node is also either a Value, an Instruction, or both.
//
// Node contains the methods common to Value and Instruction, plus the
// Operands and Referrers methods generalized to return nil for
// non-Instructions and non-Values, respectively.
//
// Node is provided to simplify SSA graph algorithms.  Clients should
// use the more specific and informative Value or Instruction
// interfaces where appropriate.
type Node interface {
	// Common methods:
	String() string
	Pos() token.Pos
	Parent() *Function

	// Partial methods:
	Operands(rands []*Value) []*Value // nil for non-Instructions
	Referrers() *[]Instruction        // nil for non-Values
}

// Function represents the parameters, results, and code of a function
// or method.
//
// If Blocks is nil, this indicates an external function for which no
// Go source code is available.  In this case, FreeVars, Locals, and
// Params are nil too.  Clients performing whole-program analysis must
// handle external functions specially.
//
// Blocks contains the function's control-flow graph (CFG).
// Blocks[0] is the function entry point; block order is not otherwise
// semantically significant, though it may affect the readability of
// the disassembly.
// To iterate over the blocks in dominance order, use DomPreorder().
//
// Recover is an optional second entry point to which control resumes
// after a recovered panic.  The Recover block may contain only a return
// statement, preceded by a load of the function's named return
// parameters, if any.
//
// A nested function (Parent()!=nil) that refers to one or more
// lexically enclosing local variables ("free variables") has FreeVars.
// Such functions cannot be called directly but require a
// value created by MakeClosure which, via its Bindings, supplies
// values for these parameters.
//
// If the function is a method (Signature.Recv() != nil) then the first
// element of Params is the receiver parameter.
//
// A Go package may declare many functions called "init".
// For each one, Object().Name() returns "init" but Name() returns
// "init#1", etc, in declaration order.
//
// Pos() returns the declaring ast.FuncLit.Type.Func or the position
// of the ast.FuncDecl.Name, if the function was explicit in the
// source. Synthetic wrappers, for which Synthetic != "", may share
// the same position as the function they wrap.
// Syntax.Pos() always returns the position of the declaring "func" token.
//
// When the operand of a range statement is an iterator function,
// the loop body is transformed into a synthetic anonymous function
// that is passed as the yield argument in a call to the iterator.
// In that case, Function.Pos is the position of the "range" token,
// and Function.Syntax is the ast.RangeStmt.
//
// Synthetic functions, for which Synthetic != "", are functions
// that do not appear in the source AST. These include:
//   - method wrappers,
//   - thunks,
//   - bound functions,
//   - empty functions built from loaded type information,
//   - yield functions created from range-over-func loops,
//   - package init functions, and
//   - instantiations of generic functions.
//
// Synthetic wrapper functions may share the same position
// as the function they wrap.
//
// Type() returns the function's Signature.
//
// A generic function is a function or method that has uninstantiated type
// parameters (TypeParams() != nil). Consider a hypothetical generic
// method, (*Map[K,V]).Get. It may be instantiated with all
// non-parameterized types as (*Map[string,int]).Get or with
// parameterized types as (*Map[string,U]).Get, where U is a type parameter.
// In both instantiations, Origin() refers to the instantiated generic
// method, (*Map[K,V]).Get, TypeParams() refers to the parameters [K,V] of
// the generic method. TypeArgs() refers to [string,U] or [string,int],
// respectively, and is nil in the generic method.
type Function struct {
	name      string
	object    *types.Func // symbol for declared function (nil for FuncLit or synthetic init)
	method    *selection  // info about provenance of synthetic methods; thunk => non-nil
	Signature *types.Signature
	pos       token.Pos

	// source information
	Synthetic string      // provenance of synthetic function; "" for true source functions
	syntax    ast.Node    // *ast.Func{Decl,Lit}, if from syntax (incl. generic instances) or (*ast.RangeStmt if a yield function)
	info      *types.Info // type annotations (if syntax != nil)
	goversion string      // Go version of syntax (NB: init is special)

	parent *Function // enclosing function if anon; nil if global
	Pkg    *Package  // enclosing package; nil for shared funcs (wrappers and error.Error)
	Prog   *Program  // enclosing program

	buildshared *task // wait for a shared function to be done building (may be nil if <=1 builder ever needs to wait)

	// These fields are populated only when the function body is built:

	Params    []*Parameter  // function parameters; for methods, includes receiver
	FreeVars  []*FreeVar    // free variables whose values must be supplied by closure
	Locals    []*Alloc      // frame-allocated variables of this function
	Blocks    []*BasicBlock // basic blocks of the function; nil => external
	Recover   *BasicBlock   // optional; control transfers here after recovered panic
	AnonFuncs []*Function   // anonymous functions (from FuncLit,RangeStmt) directly beneath this one
	referrers []Instruction // referring instructions (iff Parent() != nil)
	anonIdx   int32         // position of a nested function in parent's AnonFuncs. fn.Parent()!=nil => fn.Parent().AnonFunc[fn.anonIdx] == fn.

	typeparams     *types.TypeParamList // type parameters of this function. typeparams.Len() > 0 => generic or instance of generic function
	typeargs       []types.Type         // type arguments that instantiated typeparams. len(typeargs) > 0 => instance of generic function
	topLevelOrigin *Function            // the origin function if this is an instance of a source function. nil if Parent()!=nil.
	generic        *generic             // instances of this function, if generic

	// The following fields are cleared after building.
	build        buildFunc                // algorithm to build function body (nil => built)
	currentBlock *BasicBlock              // where to emit code
	vars         map[*types.Var]Value     // addresses of local variables
	results      []*Alloc                 // result allocations of the current function
	returnVars   []*types.Var             // variables for a return statement. Either results or for range-over-func a parent's results
	targets      *targets                 // linked stack of branch targets
	lblocks      map[*types.Label]*lblock // labelled blocks
	subst        *subster                 // type parameter substitutions (if non-nil)
	jump         *types.Var               // synthetic variable for the yield state (non-nil => range-over-func)
	deferstack   *types.Var               // synthetic variable holding enclosing ssa:deferstack()
	source       *Function                // nearest enclosing source function
	exits        []*exit                  // exits of the function that need to be resolved
	uniq         int64                    // source of unique ints within the source tree while building
}

// BasicBlock represents an SSA basic block.
//
// The final element of Instrs is always an explicit transfer of
// control (If, Jump, Return, or Panic).
//
// A block may contain no Instructions only if it is unreachable,
// i.e., Preds is nil.  Empty blocks are typically pruned.
//
// BasicBlocks and their Preds/Succs relation form a (possibly cyclic)
// graph independent of the SSA Value graph: the control-flow graph or
// CFG.  It is illegal for multiple edges to exist between the same
// pair of blocks.
//
// Each BasicBlock is also a node in the dominator tree of the CFG.
// The tree may be navigated using Idom()/Dominees() and queried using
// Dominates().
//
// The order of Preds and Succs is significant (to Phi and If
// instructions, respectively).
type BasicBlock struct {
	Index        int            // index of this block within Parent().Blocks
	Comment      string         // optional label; no semantic significance
	parent       *Function      // parent function
	Instrs       []Instruction  // instructions in order
	Preds, Succs []*BasicBlock  // predecessors and successors
	succs2       [2]*BasicBlock // initial space for Succs
	dom          domInfo        // dominator tree info
	gaps         int            // number of nil Instrs (transient)
	rundefers    int            // number of rundefers (transient)
}

// Pure values ----------------------------------------

// A FreeVar represents a free variable of the function to which it
// belongs.
//
// FreeVars are used to implement anonymous functions, whose free
// variables are lexically captured in a closure formed by
// MakeClosure.  The value of such a free var is an Alloc or another
// FreeVar and is considered a potentially escaping heap address, with
// pointer type.
//
// FreeVars are also used to implement bound method closures.  Such a
// free var represents the receiver value and may be of any type that
// has concrete methods.
//
// Pos() returns the position of the value that was captured, which
// belongs to an enclosing function.
type FreeVar struct {
	name      string
	typ       types.Type
	pos       token.Pos
	parent    *Function
	referrers []Instruction

	// Transiently needed during building.
	outer Value // the Value captured from the enclosing context.
}

// A Parameter represents an input parameter of a function.
type Parameter struct {
	name      string
	object    *types.Var // non-nil
	typ       types.Type
	parent    *Function
	referrers []Instruction
}

// A Const represents a value known at build time.
//
// Consts include true constants of boolean, numeric, and string types, as
// defined by the Go spec; these are represented by a non-nil Value field.
//
// Consts also include the "zero" value of any type, of which the nil values
// of various pointer-like types are a special case; these are represented
// by a nil Value field.
//
// Pos() returns token.NoPos.
//
// Example printed forms:
//
//		42:int
//		"hello":untyped string
//		3+4i:MyComplex
//		nil:*int
//		nil:[]string
//		[3]int{}:[3]int
//		struct{x string}{}:struct{x string}
//	    0:interface{int|int64}
//	    nil:interface{bool|int} // no go/constant representation
type Const struct {
	typ   types.Type
	Value constant.Value
}

// A Global is a named Value holding the address of a package-level
// variable.
//
// Pos() returns the position of the ast.ValueSpec.Names[*]
// identifier.
type Global struct {
	name   string
	object types.Object // a *types.Var; may be nil for synthetics e.g. init$guard
	typ    types.Type
	pos    token.Pos

	Pkg *Package
}

// A Builtin represents a specific use of a built-in function, e.g. len.
//
// Builtins are immutable values.  Builtins do not have addresses.
// Builtins can only appear in CallCommon.Value.
//
// Name() indicates the function: one of the built-in functions from the
// Go spec (excluding "make" and "new") or one of these ssa-defined
// intrinsics:
//
//	// wrapnilchk returns ptr if non-nil, panics otherwise.
//	// (For use in indirection wrappers.)
//	func ssa:wrapnilchk(ptr *T, recvType, methodName string) *T
//
// Object() returns a *types.Builtin for built-ins defined by the spec,
// nil for others.
//
// Type() returns a *types.Signature representing the effective
// signature of the built-in for this call.
type Builtin struct {
	name string
	sig  *types.Signature
}

// Value-defining instructions  ----------------------------------------

// The Alloc instruction reserves space for a variable of the given type,
// zero-initializes it, and yields its address.
//
// Alloc values are always addresses, and have pointer types, so the
// type of the allocated variable is actually
// Type().Underlying().(*types.Pointer).Elem().
//
// If Heap is false, Alloc zero-initializes the same local variable in
// the call frame and returns its address; in this case the Alloc must
// be present in Function.Locals. We call this a "local" alloc.
//
// If Heap is true, Alloc allocates a new zero-initialized variable
// each time the instruction is executed. We call this a "new" alloc.
//
// When Alloc is applied to a channel, map or slice type, it returns
// the address of an uninitialized (nil) reference of that kind; store
// the result of MakeSlice, MakeMap or MakeChan in that location to
// instantiate these types.
//
// Pos() returns the ast.CompositeLit.Lbrace for a composite literal,
// or the ast.CallExpr.Rparen for a call to new() or for a call that
// allocates a varargs slice.
//
// Example printed form:
//
//	t0 = local int
//	t1 = new int
type Alloc struct {
	register
	Comment string
	Heap    bool
	index   int // dense numbering; for lifting
}

// The Phi instruction represents an SSA φ-node, which combines values
// that differ across incoming control-flow edges and yields a new
// value.  Within a block, all φ-nodes must appear before all non-φ
// nodes.
//
// Pos() returns the position of the && or || for short-circuit
// control-flow joins, or that of the *Alloc for φ-nodes inserted
// during SSA renaming.
//
// Example printed form:
//
//	t2 = phi [0: t0, 1: t1]
type Phi struct {
	register
	Comment string  // a hint as to its purpose
	Edges   []Value // Edges[i] is value for Block().Preds[i]
}

// The Call instruction represents a function or method call.
//
// The Call instruction yields the function result if there is exactly
// one.  Otherwise it returns a tuple, the components of which are
// accessed via Extract.
//
// See CallCommon for generic function call documentation.
//
// Pos() returns the ast.CallExpr.Lparen, if explicit in the source.
//
// Example printed form:
//
//	t2 = println(t0, t1)
//	t4 = t3()
//	t7 = invoke t5.Println(...t6)
type Call struct {
	register
	Call CallCommon
}

// The BinOp instruction yields the result of binary operation X Op Y.
//
// Pos() returns the ast.BinaryExpr.OpPos, if explicit in the source.
//
// Example printed form:
//
//	t1 = t0 + 1:int
type BinOp struct {
	register
	// One of:
	// ADD SUB MUL QUO REM          + - * / %
	// AND OR XOR SHL SHR AND_NOT   & | ^ << >> &^
	// EQL NEQ LSS LEQ GTR GEQ      == != < <= < >=
	Op   token.Token
	X, Y Value
}

// The UnOp instruction yields the result of Op X.
// ARROW is channel receive.
// MUL is pointer indirection (load).
// XOR is bitwise complement.
// SUB is negation.
// NOT is logical negation.
//
// If CommaOk and Op=ARROW, the result is a 2-tuple of the value above
// and a boolean indicating the success of the receive.  The
// components of the tuple are accessed using Extract.
//
// Pos() returns the ast.UnaryExpr.OpPos, if explicit in the source.
// For receive operations (ARROW) implicit in ranging over a channel,
// Pos() returns the ast.RangeStmt.For.
// For implicit memory loads (STAR), Pos() returns the position of the
// most closely associated source-level construct; the details are not
// specified.
//
// Example printed form:
//
//	t0 = *x
//	t2 = <-t1,ok
type UnOp struct {
	register
	Op      token.Token // One of: NOT SUB ARROW MUL XOR ! - <- * ^
	X       Value
	CommaOk bool
}

// The ChangeType instruction applies to X a value-preserving type
// change to Type().
//
// Type changes are permitted:
//   - between a named type and its underlying type.
//   - between two named types of the same underlying type.
//   - between (possibly named) pointers to identical base types.
//   - from a bidirectional channel to a read- or write-channel,
//     optionally adding/removing a name.
//   - between a type (t) and an instance of the type (tσ), i.e.
//     Type() == σ(X.Type()) (or X.Type()== σ(Type())) where
//     σ is the type substitution of Parent().TypeParams by
//     Parent().TypeArgs.
//
// This operation cannot fail dynamically.
//
// Type changes may to be to or from a type parameter (or both). All
// types in the type set of X.Type() have a value-preserving type
// change to all types in the type set of Type().
//
// Pos() returns the ast.CallExpr.Lparen, if the instruction arose
// from an explicit conversion in the source.
//
// Example printed form:
//
//	t1 = changetype *int <- IntPtr (t0)
type ChangeType struct {
	register
	X Value
}

// The Convert instruction yields the conversion of value X to type
// Type().  One or both of those types is basic (but possibly named).
//
// A conversion may change the value and representation of its operand.
// Conversions are permitted:
//   - between real numeric types.
//   - between complex numeric types.
//   - between string and []byte or []rune.
//   - between pointers and unsafe.Pointer.
//   - between unsafe.Pointer and uintptr.
//   - from (Unicode) integer to (UTF-8) string.
//
// A conversion may imply a type name change also.
//
// Conversions may to be to or from a type parameter. All types in
// the type set of X.Type() can be converted to all types in the type
// set of Type().
//
// This operation cannot fail dynamically.
//
// Conversions of untyped string/number/bool constants to a specific
// representation are eliminated during SSA construction.
//
// Pos() returns the ast.CallExpr.Lparen, if the instruction arose
// from an explicit conversion in the source.
//
// Example printed form:
//
//	t1 = convert []byte <- string (t0)
type Convert struct {
	register
	X Value
}

// The MultiConvert instruction yields the conversion of value X to type
// Type(). Either X.Type() or Type() must be a type parameter. Each
// type in the type set of X.Type() can be converted to each type in the
// type set of Type().
//
// See the documentation for Convert, ChangeType, and SliceToArrayPointer
// for the conversions that are permitted. Additionally conversions of
// slices to arrays are permitted.
//
// This operation can fail dynamically (see SliceToArrayPointer).
//
// Pos() returns the ast.CallExpr.Lparen, if the instruction arose
// from an explicit conversion in the source.
//
// Example printed form:
//
//	t1 = multiconvert D <- S (t0) [*[2]rune <- []rune | string <- []rune]
type MultiConvert struct {
	register
	X        Value
	from, to types.Type
}

// ChangeInterface constructs a value of one interface type from a
// value of another interface type known to be assignable to it.
// This operation cannot fail.
//
// Pos() returns the ast.CallExpr.Lparen if the instruction arose from
// an explicit T(e) conversion; the ast.TypeAssertExpr.Lparen if the
// instruction arose from an explicit e.(T) operation; or token.NoPos
// otherwise.
//
// Example printed form:
//
//	t1 = change interface interface{} <- I (t0)
type ChangeInterface struct {
	register
	X Value
}

// The SliceToArrayPointer instruction yields the conversion of slice X to
// array pointer.
//
// Pos() returns the ast.CallExpr.Lparen, if the instruction arose
// from an explicit conversion in the source.
//
// Conversion may to be to or from a type parameter. All types in
// the type set of X.Type() must be a slice types that can be converted to
// all types in the type set of Type() which must all be pointer to array
// types.
//
// This operation can fail dynamically if the length of the slice is less
// than the length of the array.
//
// Example printed form:
//
//	t1 = slice to array pointer *[4]byte <- []byte (t0)
type SliceToArrayPointer struct {
	register
	X Value
}

// MakeInterface constructs an instance of an interface type from a
// value of a concrete type.
//
// Use Program.MethodSets.MethodSet(X.Type()) to find the method-set
// of X, and Program.MethodValue(m) to find the implementation of a method.
//
// To construct the zero value of an interface type T, use:
//
//	NewConst(constant.MakeNil(), T, pos)
//
// Pos() returns the ast.CallExpr.Lparen, if the instruction arose
// from an explicit conversion in the source.
//
// Example printed form:
//
//	t1 = make interface{} <- int (42:int)
//	t2 = make Stringer <- t0
type MakeInterface struct {
	register
	X Value
}

// The MakeClosure instruction yields a closure value whose code is
// Fn and whose free variables' values are supplied by Bindings.
//
// Type() returns a (possibly named) *types.Signature.
//
// Pos() returns the ast.FuncLit.Type.Func for a function literal
// closure or the ast.SelectorExpr.Sel for a bound method closure.
//
// Example printed form:
//
//	t0 = make closure anon@1.2 [x y z]
//	t1 = make closure bound$(main.I).add [i]
type MakeClosure struct {
	register
	Fn       Value   // always a *Function
	Bindings []Value // values for each free variable in Fn.FreeVars
}

// The MakeMap instruction creates a new hash-table-based map object
// and yields a value of kind map.
//
// Type() returns a (possibly named) *types.Map.
//
// Pos() returns the ast.CallExpr.Lparen, if created by make(map), or
// the ast.CompositeLit.Lbrack if created by a literal.
//
// Example printed form:
//
//	t1 = make map[string]int t0
//	t1 = make StringIntMap t0
type MakeMap struct {
	register
	Reserve Value // initial space reservation; nil => default
}

// The MakeChan instruction creates a new channel object and yields a
// value of kind chan.
//
// Type() returns a (possibly named) *types.Chan.
//
// Pos() returns the ast.CallExpr.Lparen for the make(chan) that
// created it.
//
// Example printed form:
//
//	t0 = make chan int 0
//	t0 = make IntChan 0
type MakeChan struct {
	register
	Size Value // int; size of buffer; zero => synchronous.
}

// The MakeSlice instruction yields a slice of length Len backed by a
// newly allocated array of length Cap.
//
// Both Len and Cap must be non-nil Values of integer type.
//
// (Alloc(types.Array) followed by Slice will not suffice because
// Alloc can only create arrays of constant length.)
//
// Type() returns a (possibly named) *types.Slice.
//
// Pos() returns the ast.CallExpr.Lparen for the make([]T) that
// created it.
//
// Example printed form:
//
//	t1 = make []string 1:int t0
//	t1 = make StringSlice 1:int t0
type MakeSlice struct {
	register
	Len Value
	Cap Value
}

// The Slice instruction yields a slice of an existing string, slice
// or *array X between optional integer bounds Low and High.
//
// Dynamically, this instruction panics if X evaluates to a nil *array
// pointer.
//
// Type() returns string if the type of X was string, otherwise a
// *types.Slice with the same element type as X.
//
// Pos() returns the ast.SliceExpr.Lbrack if created by a x[:] slice
// operation, the ast.CompositeLit.Lbrace if created by a literal, or
// NoPos if not explicit in the source (e.g. a variadic argument slice).
//
// Example printed form:
//
//	t1 = slice t0[1:]
type Slice struct {
	register
	X              Value // slice, string, or *array
	Low, High, Max Value // each may be nil
}

// The FieldAddr instruction yields the address of Field of *struct X.
//
// The field is identified by its index within the field list of the
// struct type of X.
//
// Dynamically, this instruction panics if X evaluates to a nil
// pointer.
//
// Type() returns a (possibly named) *types.Pointer.
//
// Pos() returns the position of the ast.SelectorExpr.Sel for the
// field, if explicit in the source. For implicit selections, returns
// the position of the inducing explicit selection. If produced for a
// struct literal S{f: e}, it returns the position of the colon; for
// S{e} it returns the start of expression e.
//
// Example printed form:
//
//	t1 = &t0.name [#1]
type FieldAddr struct {
	register
	X     Value // *struct
	Field int   // index into CoreType(CoreType(X.Type()).(*types.Pointer).Elem()).(*types.Struct).Fields
}

// The Field instruction yields the Field of struct X.
//
// The field is identified by its index within the field list of the
// struct type of X; by using numeric indices we avoid ambiguity of
// package-local identifiers and permit compact representations.
//
// Pos() returns the position of the ast.SelectorExpr.Sel for the
// field, if explicit in the source. For implicit selections, returns
// the position of the inducing explicit selection.

// Example printed form:
//
//	t1 = t0.name [#1]
type Field struct {
	register
	X     Value // struct
	Field int   // index into CoreType(X.Type()).(*types.Struct).Fields
}

// The IndexAddr instruction yields the address of the element at
// index Index of collection X.  Index is an integer expression.
//
// The elements of maps and strings are not addressable; use Lookup (map),
// Index (string), or MapUpdate instead.
//
// Dynamically, this instruction panics if X evaluates to a nil *array
// pointer.
//
// Type() returns a (possibly named) *types.Pointer.
//
// Pos() returns the ast.IndexExpr.Lbrack for the index operation, if
// explicit in the source.
//
// Example printed form:
//
//	t2 = &t0[t1]
type IndexAddr struct {
	register
	X     Value // *array, slice or type parameter with types array, *array, or slice.
	Index Value // numeric index
}

// The Index instruction yields element Index of collection X, an array,
// string or type parameter containing an array, a string, a pointer to an,
// array or a slice.
//
// Pos() returns the ast.IndexExpr.Lbrack for the index operation, if
// explicit in the source.
//
// Example printed form:
//
//	t2 = t0[t1]
type Index struct {
	register
	X     Value // array, string or type parameter with types array, *array, slice, or string.
	Index Value // integer index
}

// The Lookup instruction yields element Index of collection map X.
// Index is the appropriate key type.
//
// If CommaOk, the result is a 2-tuple of the value above and a
// boolean indicating the result of a map membership test for the key.
// The components of the tuple are accessed using Extract.
//
// Pos() returns the ast.IndexExpr.Lbrack, if explicit in the source.
//
// Example printed form:
//
//	t2 = t0[t1]
//	t5 = t3[t4],ok
type Lookup struct {
	register
	X       Value // map
	Index   Value // key-typed index
	CommaOk bool  // return a value,ok pair
}

// SelectState is a helper for Select.
// It represents one goal state and its corresponding communication.
type SelectState struct {
	Dir       types.ChanDir // direction of case (SendOnly or RecvOnly)
	Chan      Value         // channel to use (for send or receive)
	Send      Value         // value to send (for send)
	Pos       token.Pos     // position of token.ARROW
	DebugNode ast.Node      // ast.SendStmt or ast.UnaryExpr(<-) [debug mode]
}

// The Select instruction tests whether (or blocks until) one
// of the specified sent or received states is entered.
//
// Let n be the number of States for which Dir==RECV and T_i (0<=i<n)
// be the element type of each such state's Chan.
// Select returns an n+2-tuple
//
//	(index int, recvOk bool, r_0 T_0, ... r_n-1 T_n-1)
//
// The tuple's components, described below, must be accessed via the
// Extract instruction.
//
// If Blocking, select waits until exactly one state holds, i.e. a
// channel becomes ready for the designated operation of sending or
// receiving; select chooses one among the ready states
// pseudorandomly, performs the send or receive operation, and sets
// 'index' to the index of the chosen channel.
//
// If !Blocking, select doesn't block if no states hold; instead it
// returns immediately with index equal to -1.
//
// If the chosen channel was used for a receive, the r_i component is
// set to the received value, where i is the index of that state among
// all n receive states; otherwise r_i has the zero value of type T_i.
// Note that the receive index i is not the same as the state
// index index.
//
// The second component of the triple, recvOk, is a boolean whose value
// is true iff the selected operation was a receive and the receive
// successfully yielded a value.
//
// Pos() returns the ast.SelectStmt.Select.
//
// Example printed form:
//
//	t3 = select nonblocking [<-t0, t1<-t2]
//	t4 = select blocking []
type Select struct {
	register
	States   []*SelectState
	Blocking bool
}

// The Range instruction yields an iterator over the domain and range
// of X, which must be a string or map.
//
// Elements are accessed via Next.
//
// Type() returns an opaque and degenerate "rangeIter" type.
//
// Pos() returns the ast.RangeStmt.For.
//
// Example printed form:
//
//	t0 = range "hello":string
type Range struct {
	register
	X Value // string or map
}

// The Next instruction reads and advances the (map or string)
// iterator Iter and returns a 3-tuple value (ok, k, v).  If the
// iterator is not exhausted, ok is true and k and v are the next
// elements of the domain and range, respectively.  Otherwise ok is
// false and k and v are undefined.
//
// Components of the tuple are accessed using Extract.
//
// The IsString field distinguishes iterators over strings from those
// over maps, as the Type() alone is insufficient: consider
// map[int]rune.
//
// Type() returns a *types.Tuple for the triple (ok, k, v).
// The types of k and/or v may be types.Invalid.
//
// Example printed form:
//
//	t1 = next t0
type Next struct {
	register
	Iter     Value
	IsString bool // true => string iterator; false => map iterator.
}

// The TypeAssert instruction tests whether interface value X has type
// AssertedType.
//
// If !CommaOk, on success it returns v, the result of the conversion
// (defined below); on failure it panics.
//
// If CommaOk: on success it returns a pair (v, true) where v is the
// result of the conversion; on failure it returns (z, false) where z
// is AssertedType's zero value.  The components of the pair must be
// accessed using the Extract instruction.
//
// If Underlying: tests whether interface value X has the underlying
// type AssertedType.
//
// If AssertedType is a concrete type, TypeAssert checks whether the
// dynamic type in interface X is equal to it, and if so, the result
// of the conversion is a copy of the value in the interface.
//
// If AssertedType is an interface, TypeAssert checks whether the
// dynamic type of the interface is assignable to it, and if so, the
// result of the conversion is a copy of the interface value X.
// If AssertedType is a superinterface of X.Type(), the operation will
// fail iff the operand is nil.  (Contrast with ChangeInterface, which
// performs no nil-check.)
//
// Type() reflects the actual type of the result, possibly a
// 2-types.Tuple; AssertedType is the asserted type.
//
// Depending on the TypeAssert's purpose, Pos may return:
//   - the ast.CallExpr.Lparen of an explicit T(e) conversion;
//   - the ast.TypeAssertExpr.Lparen of an explicit e.(T) operation;
//   - the ast.CaseClause.Case of a case of a type-switch statement;
//   - the Ident(m).NamePos of an interface method value i.m
//     (for which TypeAssert may be used to effect the nil check).
//
// Example printed form:
//
//	t1 = typeassert t0.(int)
//	t3 = typeassert,ok t2.(T)
type TypeAssert struct {
	register
	X            Value
	AssertedType types.Type
	CommaOk      bool
}

// The Extract instruction yields component Index of Tuple.
//
// This is used to access the results of instructions with multiple
// return values, such as Call, TypeAssert, Next, UnOp(ARROW) and
// IndexExpr(Map).
//
// Example printed form:
//
//	t1 = extract t0 #1
type Extract struct {
	register
	Tuple Value
	Index int
}

// Instructions executed for effect.  They do not yield a value. --------------------

// The Jump instruction transfers control to the sole successor of its
// owning block.
//
// A Jump must be the last instruction of its containing BasicBlock.
//
// Pos() returns NoPos.
//
// Example printed form:
//
//	jump done
type Jump struct {
	anInstruction
}

// The If instruction transfers control to one of the two successors
// of its owning block, depending on the boolean Cond: the first if
// true, the second if false.
//
// An If instruction must be the last instruction of its containing
// BasicBlock.
//
// Pos() returns NoPos.
//
// Example printed form:
//
//	if t0 goto done else body
type If struct {
	anInstruction
	Cond Value
}

// The Return instruction returns values and control back to the calling
// function.
//
// len(Results) is always equal to the number of results in the
// function's signature.
//
// If len(Results) > 1, Return returns a tuple value with the specified
// components which the caller must access using Extract instructions.
//
// There is no instruction to return a ready-made tuple like those
// returned by a "value,ok"-mode TypeAssert, Lookup or UnOp(ARROW) or
// a tail-call to a function with multiple result parameters.
//
// Return must be the last instruction of its containing BasicBlock.
// Such a block has no successors.
//
// Pos() returns the ast.ReturnStmt.Return, if explicit in the source.
//
// Example printed form:
//
//	return
//	return nil:I, 2:int
type Return struct {
	anInstruction
	Results []Value
	pos     token.Pos
}

// The RunDefers instruction pops and invokes the entire stack of
// procedure calls pushed by Defer instructions in this function.
//
// It is legal to encounter multiple 'rundefers' instructions in a
// single control-flow path through a function; this is useful in
// the combined init() function, for example.
//
// Pos() returns NoPos.
//
// Example printed form:
//
//	rundefers
type RunDefers struct {
	anInstruction
}

// The Panic instruction initiates a panic with value X.
//
// A Panic instruction must be the last instruction of its containing
// BasicBlock, which must have no successors.
//
// NB: 'go panic(x)' and 'defer panic(x)' do not use this instruction;
// they are treated as calls to a built-in function.
//
// Pos() returns the ast.CallExpr.Lparen if this panic was explicit
// in the source.
//
// Example printed form:
//
//	panic t0
type Panic struct {
	anInstruction
	X   Value // an interface{}
	pos token.Pos
}

// The Go instruction creates a new goroutine and calls the specified
// function within it.
//
// See CallCommon for generic function call documentation.
//
// Pos() returns the ast.GoStmt.Go.
//
// Example printed form:
//
//	go println(t0, t1)
//	go t3()
//	go invoke t5.Println(...t6)
type Go struct {
	anInstruction
	Call CallCommon
	pos  token.Pos
}

// The Defer instruction pushes the specified call onto a stack of
// functions to be called by a RunDefers instruction or by a panic.
//
// If DeferStack != nil, it indicates the defer list that the defer is
// added to. Defer list values come from the Builtin function
// ssa:deferstack. Calls to ssa:deferstack() produces the defer stack
// of the current function frame. DeferStack allows for deferring into an
// alternative function stack than the current function.
//
// See CallCommon for generic function call documentation.
//
// Pos() returns the ast.DeferStmt.Defer.
//
// Example printed form:
//
//	defer println(t0, t1)
//	defer t3()
//	defer invoke t5.Println(...t6)
type Defer struct {
	anInstruction
	Call       CallCommon
	DeferStack Value // stack of deferred functions (from ssa:deferstack() intrinsic) onto which this function is pushed
	pos        token.Pos
}

// The Send instruction sends X on channel Chan.
//
// Pos() returns the ast.SendStmt.Arrow, if explicit in the source.
//
// Example printed form:
//
//	send t0 <- t1
type Send struct {
	anInstruction
	Chan, X Value
	pos     token.Pos
}

// The Store instruction stores Val at address Addr.
// Stores can be of arbitrary types.
//
// Pos() returns the position of the source-level construct most closely
// associated with the memory store operation.
// Since implicit memory stores are numerous and varied and depend upon
// implementation choices, the details are not specified.
//
// Example printed form:
//
//	*x = y
type Store struct {
	anInstruction
	Addr Value
	Val  Value
	pos  token.Pos
}

// The MapUpdate instruction updates the association of Map[Key] to
// Value.
//
// Pos() returns the ast.KeyValueExpr.Colon or ast.IndexExpr.Lbrack,
// if explicit in the source.
//
// Example printed form:
//
//	t0[t1] = t2
type MapUpdate struct {
	anInstruction
	Map   Value
	Key   Value
	Value Value
	pos   token.Pos
}

// A DebugRef instruction maps a source-level expression Expr to the
// SSA value X that represents the value (!IsAddr) or address (IsAddr)
// of that expression.
//
// DebugRef is a pseudo-instruction: it has no dynamic effect.
//
// Pos() returns Expr.Pos(), the start position of the source-level
// expression.  This is not the same as the "designated" token as
// documented at Value.Pos(). e.g. CallExpr.Pos() does not return the
// position of the ("designated") Lparen token.
//
// If Expr is an *ast.Ident denoting a var or func, Object() returns
// the object; though this information can be obtained from the type
// checker, including it here greatly facilitates debugging.
// For non-Ident expressions, Object() returns nil.
//
// DebugRefs are generated only for functions built with debugging
// enabled; see Package.SetDebugMode() and the GlobalDebug builder
// mode flag.
//
// DebugRefs are not emitted for ast.Idents referring to constants or
// predeclared identifiers, since they are trivial and numerous.
// Nor are they emitted for ast.ParenExprs.
//
// (By representing these as instructions, rather than out-of-band,
// consistency is maintained during transformation passes by the
// ordinary SSA renaming machinery.)
//
// Example printed form:
//
//	; *ast.CallExpr @ 102:9 is t5
//	; var x float64 @ 109:72 is x
//	; address of *ast.CompositeLit @ 216:10 is t0
type DebugRef struct {
	// TODO(generics): Reconsider what DebugRefs are for generics.
	anInstruction
	Expr   ast.Expr     // the referring expression (never *ast.ParenExpr)
	object types.Object // the identity of the source var/func
	IsAddr bool         // Expr is addressable and X is the address it denotes
	X      Value        // the value or address of Expr
}

// Embeddable mix-ins and helpers for common parts of other structs. -----------

// register is a mix-in embedded by all SSA values that are also
// instructions, i.e. virtual registers, and provides a uniform
// implementation of most of the Value interface: Value.Name() is a
// numbered register (e.g. "t0"); the other methods are field accessors.
//
// Temporary names are automatically assigned to each register on
// completion of building a function in SSA form.
//
// Clients must not assume that the 'id' value (and the Name() derived
// from it) is unique within a function.  As always in this API,
// semantics are determined only by identity; names exist only to
// facilitate debugging.
type register struct {
	anInstruction
	num       int        // "name" of virtual register, e.g. "t0".  Not guaranteed unique.
	typ       types.Type // type of virtual register
	pos       token.Pos  // position of source expression, or NoPos
	referrers []Instruction
}

// anInstruction is a mix-in embedded by all Instructions.
// It provides the implementations of the Block and setBlock methods.
type anInstruction struct {
	block *BasicBlock // the basic block of this instruction
}

// CallCommon is contained by Go, Defer and Call to hold the
// common parts of a function or method call.
//
// Each CallCommon exists in one of two modes, function call and
// interface method invocation, or "call" and "invoke" for short.
//
// 1. "call" mode: when Method is nil (!IsInvoke), a CallCommon
// represents an ordinary function call of the value in Value,
// which may be a *Builtin, a *Function or any other value of kind
// 'func'.
//
// Value may be one of:
//
//	(a) a *Function, indicating a statically dispatched call
//	    to a package-level function, an anonymous function, or
//	    a method of a named type.
//	(b) a *MakeClosure, indicating an immediately applied
//	    function literal with free variables.
//	(c) a *Builtin, indicating a statically dispatched call
//	    to a built-in function.
//	(d) any other value, indicating a dynamically dispatched
//	    function call.
//
// StaticCallee returns the identity of the callee in cases
// (a) and (b), nil otherwise.
//
// Args contains the arguments to the call.  If Value is a method,
// Args[0] contains the receiver parameter.
//
// Example printed form:
//
//	t2 = println(t0, t1)
//	go t3()
//	defer t5(...t6)
//
// 2. "invoke" mode: when Method is non-nil (IsInvoke), a CallCommon
// represents a dynamically dispatched call to an interface method.
// In this mode, Value is the interface value and Method is the
// interface's abstract method. The interface value may be a type
// parameter. Note: an interface method may be shared by multiple
// interfaces due to embedding; Value.Type() provides the specific
// interface used for this call.
//
// Value is implicitly supplied to the concrete method implementation
// as the receiver parameter; in other words, Args[0] holds not the
// receiver but the first true argument.
//
// Example printed form:
//
//	t1 = invoke t0.String()
//	go invoke t3.Run(t2)
//	defer invoke t4.Handle(...t5)
//
// For all calls to variadic functions (Signature().Variadic()),
// the last element of Args is a slice.
type CallCommon struct {
	Value  Value       // receiver (invoke mode) or func value (call mode)
	Method *types.Func // interface method (invoke mode)
	Args   []Value     // actual parameters (in static method call, includes receiver)
	pos    token.Pos   // position of CallExpr.Lparen, iff explicit in source
}

// IsInvoke returns true if this call has "invoke" (not "call") mode.
func (c *CallCommon) IsInvoke() bool {
	return c.Method != nil
}

func (c *CallCommon) Pos() token.Pos { return c.pos }

// Signature returns the signature of the called function.
//
// For an "invoke"-mode call, the signature of the interface method is
// returned.
//
// In either "call" or "invoke" mode, if the callee is a method, its
// receiver is represented by sig.Recv, not sig.Params().At(0).
func (c *CallCommon) Signature() *types.Signature {
	if c.Method != nil {
		return c.Method.Type().(*types.Signature)
	}
	return typeparams.CoreType(c.Value.Type()).(*types.Signature)
}

// StaticCallee returns the callee if this is a trivially static
// "call"-mode call to a function.
func (c *CallCommon) StaticCallee() *Function {
	switch fn := c.Value.(type) {
	case *Function:
		return fn
	case *MakeClosure:
		return fn.Fn.(*Function)
	}
	return nil
}

// Description returns a description of the mode of this call suitable
// for a user interface, e.g., "static method call".
func (c *CallCommon) Description() string {
	switch fn := c.Value.(type) {
	case *Builtin:
		return "built-in function call"
	case *MakeClosure:
		return "static function closure call"
	case *Function:
		if fn.Signature.Recv() != nil {
			return "static method call"
		}
		return "static function call"
	}
	if c.IsInvoke() {
		return "dynamic method call" // ("invoke" mode)
	}
	return "dynamic function call"
}

// The CallInstruction interface, implemented by *Go, *Defer and *Call,
// exposes the common parts of function-calling instructions,
// yet provides a way back to the Value defined by *Call alone.
type CallInstruction interface {
	Instruction
	Common() *CallCommon // returns the common parts of the call
	Value() *Call        // returns the result value of the call (*Call) or nil (*Go, *Defer)
}

func (s *Call) Common() *CallCommon  { return &s.Call }
func (s *Defer) Common() *CallCommon { return &s.Call }
func (s *Go) Common() *CallCommon    { return &s.Call }

func (s *Call) Value() *Call  { return s }
func (s *Defer) Value() *Call { return nil }
func (s *Go) Value() *Call    { return nil }

func (v *Builtin) Type() types.Type        { return v.sig }
func (v *Builtin) Name() string            { return v.name }
func (*Builtin) Referrers() *[]Instruction { return nil }
func (v *Builtin) Pos() token.Pos          { return token.NoPos }
func (v *Builtin) Object() types.Object    { return types.Universe.Lookup(v.name) }
func (v *Builtin) Parent() *Function       { return nil }

func (v *FreeVar) Type() types.Type          { return v.typ }
func (v *FreeVar) Name() string              { return v.name }
func (v *FreeVar) Referrers() *[]Instruction { return &v.referrers }
func (v *FreeVar) Pos() token.Pos            { return v.pos }
func (v *FreeVar) Parent() *Function         { return v.parent }

func (v *Global) Type() types.Type                     { return v.typ }
func (v *Global) Name() string                         { return v.name }
func (v *Global) Parent() *Function                    { return nil }
func (v *Global) Pos() token.Pos                       { return v.pos }
func (v *Global) Referrers() *[]Instruction            { return nil }
func (v *Global) Token() token.Token                   { return token.VAR }
func (v *Global) Object() types.Object                 { return v.object }
func (v *Global) String() string                       { return v.RelString(nil) }
func (v *Global) Package() *Package                    { return v.Pkg }
func (v *Global) RelString(from *types.Package) string { return relString(v, from) }

func (v *Function) Name() string       { return v.name }
func (v *Function) Type() types.Type   { return v.Signature }
func (v *Function) Pos() token.Pos     { return v.pos }
func (v *Function) Token() token.Token { return token.FUNC }
func (v *Function) Object() types.Object {
	if v.object != nil {
		return types.Object(v.object)
	}
	return nil
}
func (v *Function) String() string    { return v.RelString(nil) }
func (v *Function) Package() *Package { return v.Pkg }
func (v *Function) Parent() *Function { return v.parent }
func (v *Function) Referrers() *[]Instruction {
	if v.parent != nil {
		return &v.referrers
	}
	return nil
}

// TypeParams are the function's type parameters if generic or the
// type parameters that were instantiated if fn is an instantiation.
func (fn *Function) TypeParams() *types.TypeParamList {
	return fn.typeparams
}

// TypeArgs are the types that TypeParams() were instantiated by to create fn
// from fn.Origin().
func (fn *Function) TypeArgs() []types.Type { return fn.typeargs }

// Origin returns the generic function from which fn was instantiated,
// or nil if fn is not an instantiation.
func (fn *Function) Origin() *Function {
	if fn.parent != nil && len(fn.typeargs) > 0 {
		// Nested functions are BUILT at a different time than their instances.
		// Build declared package if not yet BUILT. This is not an expected use
		// case, but is simple and robust.
		fn.declaredPackage().Build()
	}
	return origin(fn)
}

// origin is the function that fn is an instantiation of. Returns nil if fn is
// not an instantiation.
//
// Precondition: fn and the origin function are done building.
func origin(fn *Function) *Function {
	if fn.parent != nil && len(fn.typeargs) > 0 {
		return origin(fn.parent).AnonFuncs[fn.anonIdx]
	}
	return fn.topLevelOrigin
}

func (v *Parameter) Type() types.Type          { return v.typ }
func (v *Parameter) Name() string              { return v.name }
func (v *Parameter) Object() types.Object      { return v.object }
func (v *Parameter) Referrers() *[]Instruction { return &v.referrers }
func (v *Parameter) Pos() token.Pos            { return v.object.Pos() }
func (v *Parameter) Parent() *Function         { return v.parent }

func (v *Alloc) Type() types.Type          { return v.typ }
func (v *Alloc) Referrers() *[]Instruction { return &v.referrers }
func (v *Alloc) Pos() token.Pos            { return v.pos }

func (v *register) Type() types.Type          { return v.typ }
func (v *register) setType(typ types.Type)    { v.typ = typ }
func (v *register) Name() string              { return fmt.Sprintf("t%d", v.num) }
func (v *register) setNum(num int)            { v.num = num }
func (v *register) Referrers() *[]Instruction { return &v.referrers }
func (v *register) Pos() token.Pos            { return v.pos }
func (v *register) setPos(pos token.Pos)      { v.pos = pos }

func (v *anInstruction) Parent() *Function          { return v.block.parent }
func (v *anInstruction) Block() *BasicBlock         { return v.block }
func (v *anInstruction) setBlock(block *BasicBlock) { v.block = block }
func (v *anInstruction) Referrers() *[]Instruction  { return nil }

func (t *Type) Name() string                         { return t.object.Name() }
func (t *Type) Pos() token.Pos                       { return t.object.Pos() }
func (t *Type) Type() types.Type                     { return t.object.Type() }
func (t *Type) Token() token.Token                   { return token.TYPE }
func (t *Type) Object() types.Object                 { return t.object }
func (t *Type) String() string                       { return t.RelString(nil) }
func (t *Type) Package() *Package                    { return t.pkg }
func (t *Type) RelString(from *types.Package) string { return relString(t, from) }

func (c *NamedConst) Name() string                         { return c.object.Name() }
func (c *NamedConst) Pos() token.Pos                       { return c.object.Pos() }
func (c *NamedConst) String() string                       { return c.RelString(nil) }
func (c *NamedConst) Type() types.Type                     { return c.object.Type() }
func (c *NamedConst) Token() token.Token                   { return token.CONST }
func (c *NamedConst) Object() types.Object                 { return c.object }
func (c *NamedConst) Package() *Package                    { return c.pkg }
func (c *NamedConst) RelString(from *types.Package) string { return relString(c, from) }

func (d *DebugRef) Object() types.Object { return d.object }

// Func returns the package-level function of the specified name,
// or nil if not found.
func (p *Package) Func(name string) (f *Function) {
	f, _ = p.Members[name].(*Function)
	return
}

// Var returns the package-level variable of the specified name,
// or nil if not found.
func (p *Package) Var(name string) (g *Global) {
	g, _ = p.Members[name].(*Global)
	return
}

// Const returns the package-level constant of the specified name,
// or nil if not found.
func (p *Package) Const(name string) (c *NamedConst) {
	c, _ = p.Members[name].(*NamedConst)
	return
}

// Type returns the package-level type of the specified name,
// or nil if not found.
func (p *Package) Type(name string) (t *Type) {
	t, _ = p.Members[name].(*Type)
	return
}

func (v *Call) Pos() token.Pos      { return v.Call.pos }
func (s *Defer) Pos() token.Pos     { return s.pos }
func (s *Go) Pos() token.Pos        { return s.pos }
func (s *MapUpdate) Pos() token.Pos { return s.pos }
func (s *Panic) Pos() token.Pos     { return s.pos }
func (s *Return) Pos() token.Pos    { return s.pos }
func (s *Send) Pos() token.Pos      { return s.pos }
func (s *Store) Pos() token.Pos     { return s.pos }
func (s *If) Pos() token.Pos        { return token.NoPos }
func (s *Jump) Pos() token.Pos      { return token.NoPos }
func (s *RunDefers) Pos() token.Pos { return token.NoPos }
func (s *DebugRef) Pos() token.Pos  { return s.Expr.Pos() }

// Operands.

func (v *Alloc) Operands(rands []*Value) []*Value {
	return rands
}

func (v *BinOp) Operands(rands []*Value) []*Value {
	return append(rands, &v.X, &v.Y)
}

func (c *CallCommon) Operands(rands []*Value) []*Value {
	rands = append(rands, &c.Value)
	for i := range c.Args {
		rands = append(rands, &c.Args[i])
	}
	return rands
}

func (s *Go) Operands(rands []*Value) []*Value {
	return s.Call.Operands(rands)
}

func (s *Call) Operands(rands []*Value) []*Value {
	return s.Call.Operands(rands)
}

func (s *Defer) Operands(rands []*Value) []*Value {
	return append(s.Call.Operands(rands), &s.DeferStack)
}

func (v *ChangeInterface) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *ChangeType) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *Convert) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *MultiConvert) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *SliceToArrayPointer) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (s *DebugRef) Operands(rands []*Value) []*Value {
	return append(rands, &s.X)
}

func (v *Extract) Operands(rands []*Value) []*Value {
	return append(rands, &v.Tuple)
}

func (v *Field) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *FieldAddr) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (s *If) Operands(rands []*Value) []*Value {
	return append(rands, &s.Cond)
}

func (v *Index) Operands(rands []*Value) []*Value {
	return append(rands, &v.X, &v.Index)
}

func (v *IndexAddr) Operands(rands []*Value) []*Value {
	return append(rands, &v.X, &v.Index)
}

func (*Jump) Operands(rands []*Value) []*Value {
	return rands
}

func (v *Lookup) Operands(rands []*Value) []*Value {
	return append(rands, &v.X, &v.Index)
}

func (v *MakeChan) Operands(rands []*Value) []*Value {
	return append(rands, &v.Size)
}

func (v *MakeClosure) Operands(rands []*Value) []*Value {
	rands = append(rands, &v.Fn)
	for i := range v.Bindings {
		rands = append(rands, &v.Bindings[i])
	}
	return rands
}

func (v *MakeInterface) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *MakeMap) Operands(rands []*Value) []*Value {
	return append(rands, &v.Reserve)
}

func (v *MakeSlice) Operands(rands []*Value) []*Value {
	return append(rands, &v.Len, &v.Cap)
}

func (v *MapUpdate) Operands(rands []*Value) []*Value {
	return append(rands, &v.Map, &v.Key, &v.Value)
}

func (v *Next) Operands(rands []*Value) []*Value {
	return append(rands, &v.Iter)
}

func (s *Panic) Operands(rands []*Value) []*Value {
	return append(rands, &s.X)
}

func (v *Phi) Operands(rands []*Value) []*Value {
	for i := range v.Edges {
		rands = append(rands, &v.Edges[i])
	}
	return rands
}

func (v *Range) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (s *Return) Operands(rands []*Value) []*Value {
	for i := range s.Results {
		rands = append(rands, &s.Results[i])
	}
	return rands
}

func (*RunDefers) Operands(rands []*Value) []*Value {
	return rands
}

func (v *Select) Operands(rands []*Value) []*Value {
	for i := range v.States {
		rands = append(rands, &v.States[i].Chan, &v.States[i].Send)
	}
	return rands
}

func (s *Send) Operands(rands []*Value) []*Value {
	return append(rands, &s.Chan, &s.X)
}

func (v *Slice) Operands(rands []*Value) []*Value {
	return append(rands, &v.X, &v.Low, &v.High, &v.Max)
}

func (s *Store) Operands(rands []*Value) []*Value {
	return append(rands, &s.Addr, &s.Val)
}

func (v *TypeAssert) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

func (v *UnOp) Operands(rands []*Value) []*Value {
	return append(rands, &v.X)
}

// Non-Instruction Values:
func (v *Builtin) Operands(rands []*Value) []*Value   { return rands }
func (v *FreeVar) Operands(rands []*Value) []*Value   { return rands }
func (v *Const) Operands(rands []*Value) []*Value     { return rands }
func (v *Function) Operands(rands []*Value) []*Value  { return rands }
func (v *Global) Operands(rands []*Value) []*Value    { return rands }
func (v *Parameter) Operands(rands []*Value) []*Value { return rands }
```

## File: go/ssa/stdlib_test.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Incomplete source tree on Android.

//go:build !android

package ssa_test

// This file runs the SSA builder in sanity-checking mode on all
// packages beneath $GOROOT and prints some summary information.
//
// Run with "go test -cpu=8 to" set GOMAXPROCS.

import (
	"go/ast"
	"go/token"
	"go/types"
	"runtime"
	"testing"
	"time"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testenv"
)

func bytesAllocated() uint64 {
	runtime.GC()
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	return stats.Alloc
}

// TestStdlib loads the entire standard library and its tools and all
// their dependencies.
//
// (As of go1.23, std is transitively closed, so adding the -deps flag
// doesn't increase its result set. The cmd pseudomodule of course
// depends on a good chunk of std, but the std+cmd set is also
// transitively closed, so long as -pgo=off.)
//
// Apart from a small number of internal packages that are not
// returned by the 'std' query, the set is essentially transitively
// closed, so marginal per-dependency costs are invisible.
func TestStdlib(t *testing.T) {
	testLoad(t, 500, "std", "cmd")
}

// TestNetHTTP builds a single SSA package but not its dependencies.
// It may help reveal costs related to dependencies (e.g. unnecessary building).
func TestNetHTTP(t *testing.T) {
	testLoad(t, 120, "net/http")
}

// TestCycles loads two standard libraries that depend on the same
// generic instantiations.
// internal/trace/testtrace and net/http both depend on
// slices.Contains[[]string string] and slices.Index[[]string string]
// This can under some schedules create a cycle of dependencies
// where both need to wait on the other to finish building.
func TestCycles(t *testing.T) {
	testenv.NeedsGo1Point(t, 23) // internal/trace/testtrace was added in 1.23.
	testLoad(t, 120, "net/http", "internal/trace/testtrace")
}

func testLoad(t *testing.T, minPkgs int, patterns ...string) {
	// Note: most of the commentary below applies to TestStdlib.

	if testing.Short() {
		t.Skip("skipping in short mode; too slow (https://golang.org/issue/14113)") // ~5s
	}
	testenv.NeedsTool(t, "go")

	// Load, parse and type-check the program.
	t0 := time.Now()
	alloc0 := bytesAllocated()

	cfg := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		t.Fatal(err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		t.Fatal("there were errors loading the packages")
	}

	t1 := time.Now()
	alloc1 := bytesAllocated()

	// Create SSA packages.
	var mode ssa.BuilderMode
	// Comment out these lines during benchmarking.  Approx SSA build costs are noted.
	mode |= ssa.SanityCheckFunctions // + 2% space, + 4% time
	mode |= ssa.GlobalDebug          // +30% space, +18% time
	mode |= ssa.InstantiateGenerics  // + 0% space, + 2% time (unlikely to reproduce outside of stdlib)
	prog, _ := ssautil.Packages(pkgs, mode)

	t2 := time.Now()

	// Build SSA.
	prog.Build()

	t3 := time.Now()
	alloc3 := bytesAllocated()

	// Sanity check to ensure we haven't dropped large numbers of packages.
	numPkgs := len(prog.AllPackages())
	if numPkgs < minPkgs {
		t.Errorf("Loaded only %d packages, want at least %d", numPkgs, minPkgs)
	}

	// Keep pkgs reachable until after we've measured memory usage.
	if len(pkgs) == 0 {
		panic("unreachable")
	}

	srcFuncs := srcFunctions(prog, pkgs)
	allFuncs := ssautil.AllFunctions(prog)

	// The assertion below is not valid if the program contains
	// variants of the same package, such as the test variants
	// (e.g. package p as compiled for test executable x) obtained
	// when cfg.Tests=true. Profile-guided optimization may
	// lead to similar variation for non-test executables.
	//
	// Ideally, the test would assert that all functions within
	// each executable (more generally: within any singly rooted
	// transitively closed subgraph of the import graph) have
	// distinct names, but that isn't so easy to compute efficiently.
	// Disabling for now.
	if false {
		// Check that all non-synthetic functions have distinct names.
		// Synthetic wrappers for exported methods should be distinct too,
		// except for unexported ones (explained at (*Function).RelString).
		byName := make(map[string]*ssa.Function)
		for fn := range allFuncs {
			if fn.Synthetic == "" || ast.IsExported(fn.Name()) {
				str := fn.String()
				prev := byName[str]
				byName[str] = fn
				if prev != nil {
					t.Errorf("%s: duplicate function named %s",
						prog.Fset.Position(fn.Pos()), str)
					t.Errorf("%s:   (previously defined here)",
						prog.Fset.Position(prev.Pos()))
				}
			}
		}
	}

	// Dump some statistics.
	var numInstrs int
	for fn := range allFuncs {
		for _, b := range fn.Blocks {
			numInstrs += len(b.Instrs)
		}
	}

	// determine line count
	var lineCount int
	prog.Fset.Iterate(func(f *token.File) bool {
		lineCount += f.LineCount()
		return true
	})

	// NB: when benchmarking, don't forget to clear the debug +
	// sanity builder flags for better performance.

	t.Log("GOMAXPROCS:           ", runtime.GOMAXPROCS(0))
	t.Log("#Source lines:        ", lineCount)
	t.Log("Load/parse/typecheck: ", t1.Sub(t0))
	t.Log("SSA create:           ", t2.Sub(t1))
	t.Log("SSA build:            ", t3.Sub(t2))

	// SSA stats:
	t.Log("#Packages:            ", numPkgs)
	t.Log("#SrcFunctions:        ", len(srcFuncs))
	t.Log("#AllFunctions:        ", len(allFuncs))
	t.Log("#Instructions:        ", numInstrs)
	t.Log("#MB AST+types:        ", int64(alloc1-alloc0)/1e6)
	t.Log("#MB SSA:              ", int64(alloc3-alloc1)/1e6)
}

// srcFunctions gathers all ssa.Functions corresponding to syntax.
// (Includes generics but excludes instances and all wrappers.)
//
// This is essentially identical to the SrcFunctions logic in
// go/analysis/passes/buildssa.
func srcFunctions(prog *ssa.Program, pkgs []*packages.Package) (res []*ssa.Function) {
	var addSrcFunc func(fn *ssa.Function)
	addSrcFunc = func(fn *ssa.Function) {
		res = append(res, fn)
		for _, anon := range fn.AnonFuncs {
			addSrcFunc(anon)
		}
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if decl, ok := decl.(*ast.FuncDecl); ok {
					obj := pkg.TypesInfo.Defs[decl.Name].(*types.Func)
					if obj == nil {
						panic("nil *types.Func: " + decl.Name.Name)
					}
					fn := prog.FuncValue(obj)
					if fn == nil {
						panic("nil *ssa.Function: " + obj.String())
					}
					addSrcFunc(fn)
				}
			}
		}
	}
	return res
}
```

## File: go/ssa/subst_test.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestSubst(t *testing.T) {
	const source = `
package P

func within(){
	// Pretend that the instantiation happens within this function.
}

type t0 int
func (t0) f()
type t1 interface{ f() }
type t2 interface{ g() }
type t3 interface{ ~int }

func Fn0[T t1](x T) T {
	x.f()
	return x
}

type A[T any] [4]T
type B[T any] []T
type C[T, S any] []struct{s S; t T}
type D[T, S any] *struct{s S; t *T}
type E[T, S any] interface{ F() (T, S) }
type F[K comparable, V any] map[K]V
type G[T any] chan *T
type H[T any] func() T
type I[T any] struct{x, y, z int; t T}
type J[T any] interface{ t1 }
type K[T any] interface{ t1; F() T }
type L[T any] interface{ F() T; J[T] }

var _ L[int] = Fn0[L[int]](nil)
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}

	var conf types.Config
	pkg, err := conf.Check("P", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	within, _ := pkg.Scope().Lookup("within").(*types.Func)
	if within == nil {
		t.Fatal("Failed to find the function within()")
	}

	for _, test := range []struct {
		expr string   // type expression of Named parameterized type
		args []string // type expressions of args for named
		want string   // expected underlying value after substitution
	}{
		{"A", []string{"string"}, "[4]string"},
		{"A", []string{"int"}, "[4]int"},
		{"B", []string{"int"}, "[]int"},
		{"B", []string{"int8"}, "[]int8"},
		{"C", []string{"int8", "string"}, "[]struct{s string; t int8}"},
		{"C", []string{"string", "int8"}, "[]struct{s int8; t string}"},
		{"D", []string{"int16", "string"}, "*struct{s string; t *int16}"},
		{"E", []string{"int32", "string"}, "interface{F() (int32, string)}"},
		{"F", []string{"int64", "string"}, "map[int64]string"},
		{"G", []string{"uint64"}, "chan *uint64"},
		{"H", []string{"uintptr"}, "func() uintptr"},
		{"I", []string{"t0"}, "struct{x int; y int; z int; t P.t0}"},
		{"J", []string{"t0"}, "interface{P.t1}"},
		{"K", []string{"t0"}, "interface{F() P.t0; P.t1}"},
		{"L", []string{"t0"}, "interface{F() P.t0; P.J[P.t0]}"},
		{"L", []string{"L[t0]"}, "interface{F() P.L[P.t0]; P.J[P.L[P.t0]]}"},
	} {
		// Eval() expr for its type.
		tv, err := types.Eval(fset, pkg, 0, test.expr)
		if err != nil {
			t.Fatalf("Eval(%s) failed: %v", test.expr, err)
		}
		// Eval() test.args[i] to get the i'th type arg.
		var targs []types.Type
		for _, astr := range test.args {
			tv, err := types.Eval(fset, pkg, 0, astr)
			if err != nil {
				t.Fatalf("Eval(%s) failed: %v", astr, err)
			}
			targs = append(targs, tv.Type)
		}

		T := tv.Type.(*types.Named)

		subst := makeSubster(types.NewContext(), within, T.TypeParams(), targs, true)
		sub := subst.typ(T.Underlying())
		if got := sub.String(); got != test.want {
			t.Errorf("subst{%v->%v}.typ(%s) = %v, want %v", test.expr, test.args, T.Underlying(), got, test.want)
		}
	}
}
```

## File: go/ssa/subst.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/aliases"
)

// subster defines a type substitution operation of a set of type parameters
// to type parameter free replacement types. Substitution is done within
// the context of a package-level function instantiation. *Named types
// declared in the function are unique to the instantiation.
//
// For example, given a parameterized function F
//
//	  func F[S, T any]() any {
//	    type X struct{ s S; next *X }
//		var p *X
//	    return p
//	  }
//
// calling the instantiation F[string, int]() returns an interface
// value (*X[string,int], nil) where the underlying value of
// X[string,int] is a struct{s string; next *X[string,int]}.
//
// A nil *subster is a valid, empty substitution map. It always acts as
// the identity function. This allows for treating parameterized and
// non-parameterized functions identically while compiling to ssa.
//
// Not concurrency-safe.
//
// Note: Some may find it helpful to think through some of the most
// complex substitution cases using lambda calculus inspired notation.
// subst.typ() solves evaluating a type expression E
// within the body of a function Fn[m] with the type parameters m
// once we have applied the type arguments N.
// We can succinctly write this as a function application:
//
//	((λm. E) N)
//
// go/types does not provide this interface directly.
// So what subster provides is a type substitution operation
//
//	E[m:=N]
type subster struct {
	replacements map[*types.TypeParam]types.Type // values should contain no type params
	cache        map[types.Type]types.Type       // cache of subst results
	origin       *types.Func                     // types.Objects declared within this origin function are unique within this context
	ctxt         *types.Context                  // speeds up repeated instantiations
	uniqueness   typeutil.Map                    // determines the uniqueness of the instantiations within the function
	// TODO(taking): consider adding Pos
}

// Returns a subster that replaces tparams[i] with targs[i]. Uses ctxt as a cache.
// targs should not contain any types in tparams.
// fn is the generic function for which we are substituting.
func makeSubster(ctxt *types.Context, fn *types.Func, tparams *types.TypeParamList, targs []types.Type, debug bool) *subster {
	assert(tparams.Len() == len(targs), "makeSubster argument count must match")

	subst := &subster{
		replacements: make(map[*types.TypeParam]types.Type, tparams.Len()),
		cache:        make(map[types.Type]types.Type),
		origin:       fn.Origin(),
		ctxt:         ctxt,
	}
	for i := 0; i < tparams.Len(); i++ {
		subst.replacements[tparams.At(i)] = targs[i]
	}
	return subst
}

// typ returns the type of t with the type parameter tparams[i] substituted
// for the type targs[i] where subst was created using tparams and targs.
func (subst *subster) typ(t types.Type) (res types.Type) {
	if subst == nil {
		return t // A nil subst is type preserving.
	}
	if r, ok := subst.cache[t]; ok {
		return r
	}
	defer func() {
		subst.cache[t] = res
	}()

	switch t := t.(type) {
	case *types.TypeParam:
		if r := subst.replacements[t]; r != nil {
			return r
		}
		return t

	case *types.Basic:
		return t

	case *types.Array:
		if r := subst.typ(t.Elem()); r != t.Elem() {
			return types.NewArray(r, t.Len())
		}
		return t

	case *types.Slice:
		if r := subst.typ(t.Elem()); r != t.Elem() {
			return types.NewSlice(r)
		}
		return t

	case *types.Pointer:
		if r := subst.typ(t.Elem()); r != t.Elem() {
			return types.NewPointer(r)
		}
		return t

	case *types.Tuple:
		return subst.tuple(t)

	case *types.Struct:
		return subst.struct_(t)

	case *types.Map:
		key := subst.typ(t.Key())
		elem := subst.typ(t.Elem())
		if key != t.Key() || elem != t.Elem() {
			return types.NewMap(key, elem)
		}
		return t

	case *types.Chan:
		if elem := subst.typ(t.Elem()); elem != t.Elem() {
			return types.NewChan(t.Dir(), elem)
		}
		return t

	case *types.Signature:
		return subst.signature(t)

	case *types.Union:
		return subst.union(t)

	case *types.Interface:
		return subst.interface_(t)

	case *types.Alias:
		return subst.alias(t)

	case *types.Named:
		return subst.named(t)

	case *opaqueType:
		return t // opaque types are never substituted

	default:
		panic("unreachable")
	}
}

// types returns the result of {subst.typ(ts[i])}.
func (subst *subster) types(ts []types.Type) []types.Type {
	res := make([]types.Type, len(ts))
	for i := range ts {
		res[i] = subst.typ(ts[i])
	}
	return res
}

func (subst *subster) tuple(t *types.Tuple) *types.Tuple {
	if t != nil {
		if vars := subst.varlist(t); vars != nil {
			return types.NewTuple(vars...)
		}
	}
	return t
}

type varlist interface {
	At(i int) *types.Var
	Len() int
}

// fieldlist is an adapter for structs for the varlist interface.
type fieldlist struct {
	str *types.Struct
}

func (fl fieldlist) At(i int) *types.Var { return fl.str.Field(i) }
func (fl fieldlist) Len() int            { return fl.str.NumFields() }

func (subst *subster) struct_(t *types.Struct) *types.Struct {
	if t != nil {
		if fields := subst.varlist(fieldlist{t}); fields != nil {
			tags := make([]string, t.NumFields())
			for i, n := 0, t.NumFields(); i < n; i++ {
				tags[i] = t.Tag(i)
			}
			return types.NewStruct(fields, tags)
		}
	}
	return t
}

// varlist returns subst(in[i]) or return nils if subst(v[i]) == v[i] for all i.
func (subst *subster) varlist(in varlist) []*types.Var {
	var out []*types.Var // nil => no updates
	for i, n := 0, in.Len(); i < n; i++ {
		v := in.At(i)
		w := subst.var_(v)
		if v != w && out == nil {
			out = make([]*types.Var, n)
			for j := 0; j < i; j++ {
				out[j] = in.At(j)
			}
		}
		if out != nil {
			out[i] = w
		}
	}
	return out
}

func (subst *subster) var_(v *types.Var) *types.Var {
	if v != nil {
		if typ := subst.typ(v.Type()); typ != v.Type() {
			if v.IsField() {
				return types.NewField(v.Pos(), v.Pkg(), v.Name(), typ, v.Embedded())
			}
			return types.NewParam(v.Pos(), v.Pkg(), v.Name(), typ)
		}
	}
	return v
}

func (subst *subster) union(u *types.Union) *types.Union {
	var out []*types.Term // nil => no updates

	for i, n := 0, u.Len(); i < n; i++ {
		t := u.Term(i)
		r := subst.typ(t.Type())
		if r != t.Type() && out == nil {
			out = make([]*types.Term, n)
			for j := 0; j < i; j++ {
				out[j] = u.Term(j)
			}
		}
		if out != nil {
			out[i] = types.NewTerm(t.Tilde(), r)
		}
	}

	if out != nil {
		return types.NewUnion(out)
	}
	return u
}

func (subst *subster) interface_(iface *types.Interface) *types.Interface {
	if iface == nil {
		return nil
	}

	// methods for the interface. Initially nil if there is no known change needed.
	// Signatures for the method where recv is nil. NewInterfaceType fills in the receivers.
	var methods []*types.Func
	initMethods := func(n int) { // copy first n explicit methods
		methods = make([]*types.Func, iface.NumExplicitMethods())
		for i := 0; i < n; i++ {
			f := iface.ExplicitMethod(i)
			norecv := changeRecv(f.Type().(*types.Signature), nil)
			methods[i] = types.NewFunc(f.Pos(), f.Pkg(), f.Name(), norecv)
		}
	}
	for i := 0; i < iface.NumExplicitMethods(); i++ {
		f := iface.ExplicitMethod(i)
		// On interfaces, we need to cycle break on anonymous interface types
		// being in a cycle with their signatures being in cycles with their receivers
		// that do not go through a Named.
		norecv := changeRecv(f.Type().(*types.Signature), nil)
		sig := subst.typ(norecv)
		if sig != norecv && methods == nil {
			initMethods(i)
		}
		if methods != nil {
			methods[i] = types.NewFunc(f.Pos(), f.Pkg(), f.Name(), sig.(*types.Signature))
		}
	}

	var embeds []types.Type
	initEmbeds := func(n int) { // copy first n embedded types
		embeds = make([]types.Type, iface.NumEmbeddeds())
		for i := 0; i < n; i++ {
			embeds[i] = iface.EmbeddedType(i)
		}
	}
	for i := 0; i < iface.NumEmbeddeds(); i++ {
		e := iface.EmbeddedType(i)
		r := subst.typ(e)
		if e != r && embeds == nil {
			initEmbeds(i)
		}
		if embeds != nil {
			embeds[i] = r
		}
	}

	if methods == nil && embeds == nil {
		return iface
	}
	if methods == nil {
		initMethods(iface.NumExplicitMethods())
	}
	if embeds == nil {
		initEmbeds(iface.NumEmbeddeds())
	}
	return types.NewInterfaceType(methods, embeds).Complete()
}

func (subst *subster) alias(t *types.Alias) types.Type {
	// See subster.named. This follows the same strategy.
	tparams := aliases.TypeParams(t)
	targs := aliases.TypeArgs(t)
	tname := t.Obj()
	torigin := aliases.Origin(t)

	if !declaredWithin(tname, subst.origin) {
		// t is declared outside of the function origin. So t is a package level type alias.
		if targs.Len() == 0 {
			// No type arguments so no instantiation needed.
			return t
		}

		// Instantiate with the substituted type arguments.
		newTArgs := subst.typelist(targs)
		return subst.instantiate(torigin, newTArgs)
	}

	if targs.Len() == 0 {
		// t is declared within the function origin and has no type arguments.
		//
		// Example: This corresponds to A or B in F, but not A[int]:
		//
		//     func F[T any]() {
		//       type A[S any] = struct{t T, s S}
		//       type B = T
		//       var x A[int]
		//       ...
		//     }
		//
		// This is somewhat different than *Named as *Alias cannot be created recursively.

		// Copy and substitute type params.
		var newTParams []*types.TypeParam
		for i := 0; i < tparams.Len(); i++ {
			cur := tparams.At(i)
			cobj := cur.Obj()
			cname := types.NewTypeName(cobj.Pos(), cobj.Pkg(), cobj.Name(), nil)
			ntp := types.NewTypeParam(cname, nil)
			subst.cache[cur] = ntp // See the comment "Note: Subtle" in subster.named.
			newTParams = append(newTParams, ntp)
		}

		// Substitute rhs.
		rhs := subst.typ(aliases.Rhs(t))

		// Create the fresh alias.
		//
		// Until 1.27, the result of aliases.NewAlias(...).Type() cannot guarantee it is a *types.Alias.
		// However, as t is an *alias.Alias and t is well-typed, then aliases must have been enabled.
		// Follow this decision, and always enable aliases here.
		const enabled = true
		obj := aliases.NewAlias(enabled, tname.Pos(), tname.Pkg(), tname.Name(), rhs, newTParams)

		// Substitute into all of the constraints after they are created.
		for i, ntp := range newTParams {
			bound := tparams.At(i).Constraint()
			ntp.SetConstraint(subst.typ(bound))
		}
		return obj.Type()
	}

	// t is declared within the function origin and has type arguments.
	//
	// Example: This corresponds to A[int] in F. Cases A and B are handled above.
	//     func F[T any]() {
	//       type A[S any] = struct{t T, s S}
	//       type B = T
	//       var x A[int]
	//       ...
	//     }
	subOrigin := subst.typ(torigin)
	subTArgs := subst.typelist(targs)
	return subst.instantiate(subOrigin, subTArgs)
}

func (subst *subster) named(t *types.Named) types.Type {
	// A Named type is a user defined type.
	// Ignoring generics, Named types are canonical: they are identical if
	// and only if they have the same defining symbol.
	// Generics complicate things, both if the type definition itself is
	// parameterized, and if the type is defined within the scope of a
	// parameterized function. In this case, two named types are identical if
	// and only if their identifying symbols are identical, and all type
	// arguments bindings in scope of the named type definition (including the
	// type parameters of the definition itself) are equivalent.
	//
	// Notably:
	// 1. For type definition type T[P1 any] struct{}, T[A] and T[B] are identical
	//    only if A and B are identical.
	// 2. Inside the generic func Fn[m any]() any { type T struct{}; return T{} },
	//    the result of Fn[A] and Fn[B] have identical type if and only if A and
	//    B are identical.
	// 3. Both 1 and 2 could apply, such as in
	//    func F[m any]() any { type T[x any] struct{}; return T{} }
	//
	// A subster replaces type parameters within a function scope, and therefore must
	// also replace free type parameters in the definitions of local types.
	//
	// Note: There are some detailed notes sprinkled throughout that borrow from
	// lambda calculus notation. These contain some over simplifying math.
	//
	// LC: One way to think about subster is that it is  a way of evaluating
	//   ((λm. E) N) as E[m:=N].
	// Each Named type t has an object *TypeName within a scope S that binds an
	// underlying type expression U. U can refer to symbols within S (+ S's ancestors).
	// Let x = t.TypeParams() and A = t.TypeArgs().
	// Each Named type t is then either:
	//   U              where len(x) == 0 && len(A) == 0
	//   λx. U          where len(x) != 0 && len(A) == 0
	//   ((λx. U) A)    where len(x) == len(A)
	// In each case, we will evaluate t[m:=N].
	tparams := t.TypeParams() // x
	targs := t.TypeArgs()     // A

	if !declaredWithin(t.Obj(), subst.origin) {
		// t is declared outside of Fn[m].
		//
		// In this case, we can skip substituting t.Underlying().
		// The underlying type cannot refer to the type parameters.
		//
		// LC: Let free(E) be the set of free type parameters in an expression E.
		// Then whenever m ∉ free(E), then E = E[m:=N].
		// t ∉ Scope(fn) so therefore m ∉ free(U) and m ∩ x = ∅.
		if targs.Len() == 0 {
			// t has no type arguments. So it does not need to be instantiated.
			//
			// This is the normal case in real Go code, where t is not parameterized,
			// declared at some package scope, and m is a TypeParam from a parameterized
			// function F[m] or method.
			//
			// LC: m ∉ free(A) lets us conclude m ∉ free(t). So t=t[m:=N].
			return t
		}

		// t is declared outside of Fn[m] and has type arguments.
		// The type arguments may contain type parameters m so
		// substitute the type arguments, and instantiate the substituted
		// type arguments.
		//
		// LC: Evaluate this as ((λx. U) A') where A' = A[m := N].
		newTArgs := subst.typelist(targs)
		return subst.instantiate(t.Origin(), newTArgs)
	}

	// t is declared within Fn[m].

	if targs.Len() == 0 { // no type arguments?
		assert(t == t.Origin(), "local parameterized type abstraction must be an origin type")

		// t has no type arguments.
		// The underlying type of t may contain the function's type parameters,
		// replace these, and create a new type.
		//
		// Subtle: We short circuit substitution and use a newly created type in
		// subst, i.e. cache[t]=fresh, to preemptively replace t with fresh
		// in recursive types during traversal. This both breaks infinite cycles
		// and allows for constructing types with the replacement applied in
		// subst.typ(U).
		//
		// A new copy of the Named and Typename (and constraints) per function
		// instantiation matches the semantics of Go, which treats all function
		// instantiations F[N] as having distinct local types.
		//
		// LC: x.Len()=0 can be thought of as a special case of λx. U.
		// LC: Evaluate (λx. U)[m:=N] as (λx'. U') where U'=U[x:=x',m:=N].
		tname := t.Obj()
		obj := types.NewTypeName(tname.Pos(), tname.Pkg(), tname.Name(), nil)
		fresh := types.NewNamed(obj, nil, nil)
		var newTParams []*types.TypeParam
		for i := 0; i < tparams.Len(); i++ {
			cur := tparams.At(i)
			cobj := cur.Obj()
			cname := types.NewTypeName(cobj.Pos(), cobj.Pkg(), cobj.Name(), nil)
			ntp := types.NewTypeParam(cname, nil)
			subst.cache[cur] = ntp
			newTParams = append(newTParams, ntp)
		}
		fresh.SetTypeParams(newTParams)
		subst.cache[t] = fresh
		subst.cache[fresh] = fresh
		fresh.SetUnderlying(subst.typ(t.Underlying()))
		// Substitute into all of the constraints after they are created.
		for i, ntp := range newTParams {
			bound := tparams.At(i).Constraint()
			ntp.SetConstraint(subst.typ(bound))
		}
		return fresh
	}

	// t is defined within Fn[m] and t has type arguments (an instantiation).
	// We reduce this to the two cases above:
	// (1) substitute the function's type parameters into t.Origin().
	// (2) substitute t's type arguments A and instantiate the updated t.Origin() with these.
	//
	// LC: Evaluate ((λx. U) A)[m:=N] as (t' A') where t' = (λx. U)[m:=N] and A'=A [m:=N]
	subOrigin := subst.typ(t.Origin())
	subTArgs := subst.typelist(targs)
	return subst.instantiate(subOrigin, subTArgs)
}

func (subst *subster) instantiate(orig types.Type, targs []types.Type) types.Type {
	i, err := types.Instantiate(subst.ctxt, orig, targs, false)
	assert(err == nil, "failed to Instantiate named (Named or Alias) type")
	if c, _ := subst.uniqueness.At(i).(types.Type); c != nil {
		return c.(types.Type)
	}
	subst.uniqueness.Set(i, i)
	return i
}

func (subst *subster) typelist(l *types.TypeList) []types.Type {
	res := make([]types.Type, l.Len())
	for i := 0; i < l.Len(); i++ {
		res[i] = subst.typ(l.At(i))
	}
	return res
}

func (subst *subster) signature(t *types.Signature) types.Type {
	tparams := t.TypeParams()

	// We are choosing not to support tparams.Len() > 0 until a need has been observed in practice.
	//
	// There are some known usages for types.Types coming from types.{Eval,CheckExpr}.
	// To support tparams.Len() > 0, we just need to do the following [psuedocode]:
	//   targs := {subst.replacements[tparams[i]]]}; Instantiate(ctxt, t, targs, false)

	assert(tparams.Len() == 0, "Substituting types.Signatures with generic functions are currently unsupported.")

	// Either:
	// (1)non-generic function.
	//    no type params to substitute
	// (2)generic method and recv needs to be substituted.

	// Receivers can be either:
	// named
	// pointer to named
	// interface
	// nil
	// interface is the problematic case. We need to cycle break there!
	recv := subst.var_(t.Recv())
	params := subst.tuple(t.Params())
	results := subst.tuple(t.Results())
	if recv != t.Recv() || params != t.Params() || results != t.Results() {
		return types.NewSignatureType(recv, nil, nil, params, results, t.Variadic())
	}
	return t
}

// reaches returns true if a type t reaches any type t' s.t. c[t'] == true.
// It updates c to cache results.
//
// reaches is currently only part of the wellFormed debug logic, and
// in practice c is initially only type parameters. It is not currently
// relied on in production.
func reaches(t types.Type, c map[types.Type]bool) (res bool) {
	if c, ok := c[t]; ok {
		return c
	}

	// c is populated with temporary false entries as types are visited.
	// This avoids repeat visits and break cycles.
	c[t] = false
	defer func() {
		c[t] = res
	}()

	switch t := t.(type) {
	case *types.TypeParam, *types.Basic:
		return false
	case *types.Array:
		return reaches(t.Elem(), c)
	case *types.Slice:
		return reaches(t.Elem(), c)
	case *types.Pointer:
		return reaches(t.Elem(), c)
	case *types.Tuple:
		for i := 0; i < t.Len(); i++ {
			if reaches(t.At(i).Type(), c) {
				return true
			}
		}
	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			if reaches(t.Field(i).Type(), c) {
				return true
			}
		}
	case *types.Map:
		return reaches(t.Key(), c) || reaches(t.Elem(), c)
	case *types.Chan:
		return reaches(t.Elem(), c)
	case *types.Signature:
		if t.Recv() != nil && reaches(t.Recv().Type(), c) {
			return true
		}
		return reaches(t.Params(), c) || reaches(t.Results(), c)
	case *types.Union:
		for i := 0; i < t.Len(); i++ {
			if reaches(t.Term(i).Type(), c) {
				return true
			}
		}
	case *types.Interface:
		for i := 0; i < t.NumEmbeddeds(); i++ {
			if reaches(t.Embedded(i), c) {
				return true
			}
		}
		for i := 0; i < t.NumExplicitMethods(); i++ {
			if reaches(t.ExplicitMethod(i).Type(), c) {
				return true
			}
		}
	case *types.Named, *types.Alias:
		return reaches(t.Underlying(), c)
	default:
		panic("unreachable")
	}
	return false
}
```

## File: go/ssa/task.go
```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"sync/atomic"
)

// Each task has two states: it is initially "active",
// and transitions to "done".
//
// tasks form a directed graph. An edge from x to y (with y in x.edges)
// indicates that the task x waits on the task y to be done.
// Cycles are permitted.
//
// Calling x.wait() blocks the calling goroutine until task x,
// and all the tasks transitively reachable from x are done.
//
// The nil *task is always considered done.
type task struct {
	done       chan unit      // close when the task is done.
	edges      map[*task]unit // set of predecessors of this task.
	transitive atomic.Bool    // true once it is known all predecessors are done.
}

func (x *task) isTransitivelyDone() bool { return x == nil || x.transitive.Load() }

// addEdge creates an edge from x to y, indicating that
// x.wait() will not return before y is done.
// All calls to x.addEdge(...) should happen before x.markDone().
func (x *task) addEdge(y *task) {
	if x == y || y.isTransitivelyDone() {
		return // no work remaining
	}

	// heuristic done check
	select {
	case <-x.done:
		panic("cannot add an edge to a done task")
	default:
	}

	if x.edges == nil {
		x.edges = make(map[*task]unit)
	}
	x.edges[y] = unit{}
}

// markDone changes the task's state to markDone.
func (x *task) markDone() {
	if x != nil {
		close(x.done)
	}
}

// wait blocks until x and all the tasks it can reach through edges are done.
func (x *task) wait() {
	if x.isTransitivelyDone() {
		return // already known to be done. Skip allocations.
	}

	// Use BFS to wait on u.done to be closed, for all u transitively
	// reachable from x via edges.
	//
	// This work can be repeated by multiple workers doing wait().
	//
	// Note: Tarjan's SCC algorithm is able to mark SCCs as transitively done
	// as soon as the SCC has been visited. This is theoretically faster, but is
	// a more complex algorithm. Until we have evidence, we need the more complex
	// algorithm, the simpler algorithm BFS is implemented.
	//
	// In Go 1.23, ssa/TestStdlib reaches <=3 *tasks per wait() in most schedules
	// On some schedules, there is a cycle building net/http and internal/trace/testtrace
	// due to slices functions.
	work := []*task{x}
	enqueued := map[*task]unit{x: {}}
	for i := 0; i < len(work); i++ {
		u := work[i]
		if u.isTransitivelyDone() { // already transitively done
			work[i] = nil
			continue
		}
		<-u.done // wait for u to be marked done.

		for v := range u.edges {
			if _, ok := enqueued[v]; !ok {
				enqueued[v] = unit{}
				work = append(work, v)
			}
		}
	}

	// work is transitively closed over dependencies.
	// u in work is done (or transitively done and skipped).
	// u is transitively done.
	for _, u := range work {
		if u != nil {
			x.transitive.Store(true)
		}
	}
}
```

## File: go/ssa/testutil_test.go
```go
// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file defines helper functions for SSA tests.

package ssa_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/internal/testenv"
	"golang.org/x/tools/internal/testfiles"
	"golang.org/x/tools/txtar"
)

// goMod returns a go.mod file containing a name and a go directive
// for the major version. If major < 0, use the current go toolchain
// version.
func goMod(name string, major int) []byte {
	if major < 0 {
		major = testenv.Go1Point()
	}
	return fmt.Appendf(nil, "module %s\ngo 1.%d", name, major)
}

// overlayFS returns a simple in-memory filesystem.
func overlayFS(overlay map[string][]byte) fstest.MapFS {
	// taking: Maybe loadPackages should take an overlay instead?
	fs := make(fstest.MapFS)
	for name, data := range overlay {
		fs[name] = &fstest.MapFile{Data: data}
	}
	return fs
}

// openTxtar opens a txtar file as a filesystem.
func openTxtar(t testing.TB, file string) fs.FS {
	// TODO(taking): Move to testfiles?
	t.Helper()

	ar, err := txtar.ParseFile(file)
	if err != nil {
		t.Fatal(err)
	}

	fs, err := txtar.FS(ar)
	if err != nil {
		t.Fatal(err)
	}

	return fs
}

// loadPackages copies the files in a source file system to a unique temporary
// directory and loads packages matching the given patterns from the temporary directory.
//
// TODO(69556): Migrate loader tests to loadPackages.
func loadPackages(t testing.TB, src fs.FS, patterns ...string) []*packages.Package {
	t.Helper()
	testenv.NeedsGoBuild(t) // for go/packages

	// TODO(taking): src and overlays are very similar. Overlays could have nicer paths.
	// Look into migrating src to overlays.
	dir := testfiles.CopyToTmp(t, src)

	cfg := &packages.Config{
		Dir: dir,
		Mode: packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedDeps |
			packages.NeedName |
			packages.NeedFiles |
			packages.NeedImports |
			packages.NeedCompiledGoFiles |
			packages.NeedTypes,
		Env: append(os.Environ(),
			"GO111MODULES=on",
			"GOPATH=",
			"GOWORK=off",
			"GOPROXY=off"),
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		t.Fatal(err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		t.Fatal("there were errors")
	}
	return pkgs
}

// buildPackage builds the content of a go file into:
// * a module with the same name as the package at the current go version,
// * loads the *package.Package,
// * checks that (*packages.Packages).Syntax contains one file,
// * builds the *ssa.Package (and not its dependencies), and
// * returns the built *ssa.Package and the loaded packages.Package.
//
// TODO(adonovan): factor with similar loadFile (2x) in cha/cha_test.go and vta/helpers_test.go.
func buildPackage(t testing.TB, content string, mode ssa.BuilderMode) (*ssa.Package, *packages.Package) {
	name := parsePackageClause(t, content)

	fs := overlayFS(map[string][]byte{
		"go.mod":   goMod(name, -1),
		"input.go": []byte(content),
	})
	ppkgs := loadPackages(t, fs, name)
	if len(ppkgs) != 1 {
		t.Fatalf("Expected to load 1 package from pattern %q. got %d", name, len(ppkgs))
	}
	ppkg := ppkgs[0]

	if len(ppkg.Syntax) != 1 {
		t.Fatalf("Expected 1 file in package %q. got %d", ppkg, len(ppkg.Syntax))
	}

	prog, _ := ssautil.Packages(ppkgs, mode)

	ssapkg := prog.Package(ppkg.Types)
	if ssapkg == nil {
		t.Fatalf("Failed to find ssa package for %q", ppkg.Types)
	}
	ssapkg.Build()

	return ssapkg, ppkg
}

// parsePackageClause is a test helper to extract the package name from a string
// containing the content of a go file.
func parsePackageClause(t testing.TB, content string) string {
	f, err := parser.ParseFile(token.NewFileSet(), "", content, parser.PackageClauseOnly)
	if err != nil {
		t.Fatalf("parsing the file %q failed with error: %s", content, err)
	}
	return f.Name.Name
}
```

## File: go/ssa/TODO
```
-*- text -*-

SSA Generics to-do list
===========================

DOCUMENTATION:
- Read me for internals

TYPE PARAMETERIZED GENERIC FUNCTIONS:
- sanity.go updates.
- Check source functions going to generics.
- Tests, tests, tests...

USAGE:
- Back fill users for handling ssa.InstantiateGenerics being off.
```

## File: go/ssa/typeset.go
```go
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"go/types"

	"golang.org/x/tools/internal/typeparams"
)

// Utilities for dealing with type sets.

const debug = false

// typeset is an iterator over the (type/underlying type) pairs of the
// specific type terms of the type set implied by t.
// If t is a type parameter, the implied type set is the type set of t's constraint.
// In that case, if there are no specific terms, typeset calls yield with (nil, nil).
// If t is not a type parameter, the implied type set consists of just t.
// In any case, typeset is guaranteed to call yield at least once.
func typeset(typ types.Type, yield func(t, u types.Type) bool) {
	switch typ := types.Unalias(typ).(type) {
	case *types.TypeParam, *types.Interface:
		terms := termListOf(typ)
		if len(terms) == 0 {
			yield(nil, nil)
			return
		}
		for _, term := range terms {
			u := types.Unalias(term.Type())
			if !term.Tilde() {
				u = u.Underlying()
			}
			if debug {
				assert(types.Identical(u, u.Underlying()), "Unalias(x) == under(x) for ~x terms")
			}
			if !yield(term.Type(), u) {
				break
			}
		}
		return
	default:
		yield(typ, typ.Underlying())
	}
}

// termListOf returns the type set of typ as a normalized term set. Returns an empty set on an error.
func termListOf(typ types.Type) []*types.Term {
	terms, err := typeparams.NormalTerms(typ)
	if err != nil {
		return nil
	}
	return terms
}

// typeSetIsEmpty returns true if a typeset is empty.
func typeSetIsEmpty(typ types.Type) bool {
	var empty bool
	typeset(typ, func(t, _ types.Type) bool {
		empty = t == nil
		return false
	})
	return empty
}

// isBytestring returns true if T has the same terms as interface{[]byte | string}.
// These act like a core type for some operations: slice expressions, append and copy.
//
// See https://go.dev/ref/spec#Core_types for the details on bytestring.
func isBytestring(T types.Type) bool {
	U := T.Underlying()
	if _, ok := U.(*types.Interface); !ok {
		return false
	}

	hasBytes, hasString := false, false
	ok := underIs(U, func(t types.Type) bool {
		switch {
		case isString(t):
			hasString = true
			return true
		case isByteSlice(t):
			hasBytes = true
			return true
		default:
			return false
		}
	})
	return ok && hasBytes && hasString
}

// underIs calls f with the underlying types of the type terms
// of the type set of typ and reports whether all calls to f returned true.
// If there are no specific terms, underIs returns the result of f(nil).
func underIs(typ types.Type, f func(types.Type) bool) bool {
	var ok bool
	typeset(typ, func(t, u types.Type) bool {
		ok = f(u)
		return ok
	})
	return ok
}

// indexType returns the element type and index mode of a IndexExpr over a type.
// It returns an invalid mode if the type is not indexable; this should never occur in a well-typed program.
func indexType(typ types.Type) (types.Type, indexMode) {
	switch U := typ.Underlying().(type) {
	case *types.Array:
		return U.Elem(), ixArrVar
	case *types.Pointer:
		if arr, ok := U.Elem().Underlying().(*types.Array); ok {
			return arr.Elem(), ixVar
		}
	case *types.Slice:
		return U.Elem(), ixVar
	case *types.Map:
		return U.Elem(), ixMap
	case *types.Basic:
		return tByte, ixValue // must be a string
	case *types.Interface:
		var elem types.Type
		mode := ixInvalid
		typeset(typ, func(t, _ types.Type) bool {
			if t == nil {
				return false // empty set
			}
			e, m := indexType(t)
			if elem == nil {
				elem, mode = e, m
			}
			if debug && !types.Identical(elem, e) { // if type checked, just a sanity check
				mode = ixInvalid
				return false
			}
			// Update the mode to the most constrained address type.
			mode = mode.meet(m)
			return mode != ixInvalid
		})
		return elem, mode
	}
	return nil, ixInvalid
}

// An indexMode specifies the (addressing) mode of an index operand.
//
// Addressing mode of an index operation is based on the set of
// underlying types.
// Hasse diagram of the indexMode meet semi-lattice:
//
//	ixVar     ixMap
//	  |          |
//	ixArrVar     |
//	  |          |
//	ixValue      |
//	   \        /
//	  ixInvalid
type indexMode byte

const (
	ixInvalid indexMode = iota // index is invalid
	ixValue                    // index is a computed value (not addressable)
	ixArrVar                   // like ixVar, but index operand contains an array
	ixVar                      // index is an addressable variable
	ixMap                      // index is a map index expression (acts like a variable on lhs, commaok on rhs of an assignment)
)

// meet is the address type that is constrained by both x and y.
func (x indexMode) meet(y indexMode) indexMode {
	if (x == ixMap || y == ixMap) && x != y {
		return ixInvalid
	}
	// Use int representation and return min.
	if x < y {
		return y
	}
	return x
}
```

## File: go/ssa/util.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines a number of miscellaneous utility functions.

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"os"
	"sync"
	_ "unsafe" // for go:linkname hack

	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/typeparams"
	"golang.org/x/tools/internal/typesinternal"
)

type unit struct{}

//// Sanity checking utilities

// assert panics with the mesage msg if p is false.
// Avoid combining with expensive string formatting.
func assert(p bool, msg string) {
	if !p {
		panic(msg)
	}
}

//// AST utilities

// isBlankIdent returns true iff e is an Ident with name "_".
// They have no associated types.Object, and thus no type.
func isBlankIdent(e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	return ok && id.Name == "_"
}

//// Type utilities.  Some of these belong in go/types.

// isNonTypeParamInterface reports whether t is an interface type but not a type parameter.
func isNonTypeParamInterface(t types.Type) bool {
	return !typeparams.IsTypeParam(t) && types.IsInterface(t)
}

// isBasic reports whether t is a basic type.
// t is assumed to be an Underlying type (not Named or Alias).
func isBasic(t types.Type) bool {
	_, ok := t.(*types.Basic)
	return ok
}

// isString reports whether t is exactly a string type.
// t is assumed to be an Underlying type (not Named or Alias).
func isString(t types.Type) bool {
	basic, ok := t.(*types.Basic)
	return ok && basic.Info()&types.IsString != 0
}

// isByteSlice reports whether t is of the form []~bytes.
// t is assumed to be an Underlying type (not Named or Alias).
func isByteSlice(t types.Type) bool {
	if b, ok := t.(*types.Slice); ok {
		e, _ := b.Elem().Underlying().(*types.Basic)
		return e != nil && e.Kind() == types.Byte
	}
	return false
}

// isRuneSlice reports whether t is of the form []~runes.
// t is assumed to be an Underlying type (not Named or Alias).
func isRuneSlice(t types.Type) bool {
	if b, ok := t.(*types.Slice); ok {
		e, _ := b.Elem().Underlying().(*types.Basic)
		return e != nil && e.Kind() == types.Rune
	}
	return false
}

// isBasicConvTypes returns true when the type set of a type
// can be one side of a Convert operation. This is when:
// - All are basic, []byte, or []rune.
// - At least 1 is basic.
// - At most 1 is []byte or []rune.
func isBasicConvTypes(typ types.Type) bool {
	basics, cnt := 0, 0
	ok := underIs(typ, func(t types.Type) bool {
		cnt++
		if isBasic(t) {
			basics++
			return true
		}
		return isByteSlice(t) || isRuneSlice(t)
	})
	return ok && basics >= 1 && cnt-basics <= 1
}

// isPointer reports whether t's underlying type is a pointer.
func isPointer(t types.Type) bool {
	return is[*types.Pointer](t.Underlying())
}

// isPointerCore reports whether t's core type is a pointer.
//
// (Most pointer manipulation is related to receivers, in which case
// isPointer is appropriate. tecallers can use isPointer(t).
func isPointerCore(t types.Type) bool {
	return is[*types.Pointer](typeparams.CoreType(t))
}

func is[T any](x any) bool {
	_, ok := x.(T)
	return ok
}

// recvType returns the receiver type of method obj.
func recvType(obj *types.Func) types.Type {
	return obj.Type().(*types.Signature).Recv().Type()
}

// fieldOf returns the index'th field of the (core type of) a struct type;
// otherwise returns nil.
func fieldOf(typ types.Type, index int) *types.Var {
	if st, ok := typeparams.CoreType(typ).(*types.Struct); ok {
		if 0 <= index && index < st.NumFields() {
			return st.Field(index)
		}
	}
	return nil
}

// isUntyped reports whether typ is the type of an untyped constant.
func isUntyped(typ types.Type) bool {
	// No Underlying/Unalias: untyped constant types cannot be Named or Alias.
	b, ok := typ.(*types.Basic)
	return ok && b.Info()&types.IsUntyped != 0
}

// declaredWithin reports whether an object is declared within a function.
//
// obj must not be a method or a field.
func declaredWithin(obj types.Object, fn *types.Func) bool {
	if obj.Pos() != token.NoPos {
		return fn.Scope().Contains(obj.Pos()) // trust the positions if they exist.
	}
	if fn.Pkg() != obj.Pkg() {
		return false // fast path for different packages
	}

	// Traverse Parent() scopes for fn.Scope().
	for p := obj.Parent(); p != nil; p = p.Parent() {
		if p == fn.Scope() {
			return true
		}
	}
	return false
}

// logStack prints the formatted "start" message to stderr and
// returns a closure that prints the corresponding "end" message.
// Call using 'defer logStack(...)()' to show builder stack on panic.
// Don't forget trailing parens!
func logStack(format string, args ...any) func() {
	msg := fmt.Sprintf(format, args...)
	io.WriteString(os.Stderr, msg)
	io.WriteString(os.Stderr, "\n")
	return func() {
		io.WriteString(os.Stderr, msg)
		io.WriteString(os.Stderr, " end\n")
	}
}

// newVar creates a 'var' for use in a types.Tuple.
func newVar(name string, typ types.Type) *types.Var {
	return types.NewParam(token.NoPos, nil, name, typ)
}

// anonVar creates an anonymous 'var' for use in a types.Tuple.
func anonVar(typ types.Type) *types.Var {
	return newVar("", typ)
}

var lenResults = types.NewTuple(anonVar(tInt))

// makeLen returns the len builtin specialized to type func(T)int.
func makeLen(T types.Type) *Builtin {
	lenParams := types.NewTuple(anonVar(T))
	return &Builtin{
		name: "len",
		sig:  types.NewSignatureType(nil, nil, nil, lenParams, lenResults, false),
	}
}

// receiverTypeArgs returns the type arguments to a method's receiver.
// Returns an empty list if the receiver does not have type arguments.
func receiverTypeArgs(method *types.Func) []types.Type {
	recv := method.Type().(*types.Signature).Recv()
	_, named := typesinternal.ReceiverNamed(recv)
	if named == nil {
		return nil // recv is anonymous struct/interface
	}
	ts := named.TypeArgs()
	if ts.Len() == 0 {
		return nil
	}
	targs := make([]types.Type, ts.Len())
	for i := 0; i < ts.Len(); i++ {
		targs[i] = ts.At(i)
	}
	return targs
}

// recvAsFirstArg takes a method signature and returns a function
// signature with receiver as the first parameter.
func recvAsFirstArg(sig *types.Signature) *types.Signature {
	params := make([]*types.Var, 0, 1+sig.Params().Len())
	params = append(params, sig.Recv())
	for i := 0; i < sig.Params().Len(); i++ {
		params = append(params, sig.Params().At(i))
	}
	return types.NewSignatureType(nil, nil, nil, types.NewTuple(params...), sig.Results(), sig.Variadic())
}

// instance returns whether an expression is a simple or qualified identifier
// that is a generic instantiation.
func instance(info *types.Info, expr ast.Expr) bool {
	// Compare the logic here against go/types.instantiatedIdent,
	// which also handles  *IndexExpr and *IndexListExpr.
	var id *ast.Ident
	switch x := expr.(type) {
	case *ast.Ident:
		id = x
	case *ast.SelectorExpr:
		id = x.Sel
	default:
		return false
	}
	_, ok := info.Instances[id]
	return ok
}

// instanceArgs returns the Instance[id].TypeArgs as a slice.
func instanceArgs(info *types.Info, id *ast.Ident) []types.Type {
	targList := info.Instances[id].TypeArgs
	if targList == nil {
		return nil
	}

	targs := make([]types.Type, targList.Len())
	for i, n := 0, targList.Len(); i < n; i++ {
		targs[i] = targList.At(i)
	}
	return targs
}

// Mapping of a type T to a canonical instance C s.t. types.Identical(T, C).
// Thread-safe.
type canonizer struct {
	mu    sync.Mutex
	types typeutil.Map // map from type to a canonical instance
	lists typeListMap  // map from a list of types to a canonical instance
}

func newCanonizer() *canonizer {
	c := &canonizer{}
	h := typeutil.MakeHasher()
	c.types.SetHasher(h)
	c.lists.hasher = h
	return c
}

// List returns a canonical representative of a list of types.
// Representative of the empty list is nil.
func (c *canonizer) List(ts []types.Type) *typeList {
	if len(ts) == 0 {
		return nil
	}

	unaliasAll := func(ts []types.Type) []types.Type {
		// Is there some top level alias?
		var found bool
		for _, t := range ts {
			if _, ok := t.(*types.Alias); ok {
				found = true
				break
			}
		}
		if !found {
			return ts // no top level alias
		}

		cp := make([]types.Type, len(ts)) // copy with top level aliases removed.
		for i, t := range ts {
			cp[i] = types.Unalias(t)
		}
		return cp
	}
	l := unaliasAll(ts)

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lists.rep(l)
}

// Type returns a canonical representative of type T.
// Removes top-level aliases.
//
// For performance, reasons the canonical instance is order-dependent,
// and may contain deeply nested aliases.
func (c *canonizer) Type(T types.Type) types.Type {
	T = types.Unalias(T) // remove the top level alias.

	c.mu.Lock()
	defer c.mu.Unlock()

	if r := c.types.At(T); r != nil {
		return r.(types.Type)
	}
	c.types.Set(T, T)
	return T
}

// A type for representing a canonized list of types.
type typeList []types.Type

func (l *typeList) identical(ts []types.Type) bool {
	if l == nil {
		return len(ts) == 0
	}
	n := len(*l)
	if len(ts) != n {
		return false
	}
	for i, left := range *l {
		right := ts[i]
		if !types.Identical(left, right) {
			return false
		}
	}
	return true
}

type typeListMap struct {
	hasher  typeutil.Hasher
	buckets map[uint32][]*typeList
}

// rep returns a canonical representative of a slice of types.
func (m *typeListMap) rep(ts []types.Type) *typeList {
	if m == nil || len(ts) == 0 {
		return nil
	}

	if m.buckets == nil {
		m.buckets = make(map[uint32][]*typeList)
	}

	h := m.hash(ts)
	bucket := m.buckets[h]
	for _, l := range bucket {
		if l.identical(ts) {
			return l
		}
	}

	// not present. create a representative.
	cp := make(typeList, len(ts))
	copy(cp, ts)
	rep := &cp

	m.buckets[h] = append(bucket, rep)
	return rep
}

func (m *typeListMap) hash(ts []types.Type) uint32 {
	if m == nil {
		return 0
	}
	// Some smallish prime far away from typeutil.Hash.
	n := len(ts)
	h := uint32(13619) + 2*uint32(n)
	for i := 0; i < n; i++ {
		h += 3 * m.hasher.Hash(ts[i])
	}
	return h
}

// instantiateMethod instantiates m with targs and returns a canonical representative for this method.
func (canon *canonizer) instantiateMethod(m *types.Func, targs []types.Type, ctxt *types.Context) *types.Func {
	recv := recvType(m)
	if p, ok := types.Unalias(recv).(*types.Pointer); ok {
		recv = p.Elem()
	}
	named := types.Unalias(recv).(*types.Named)
	inst, err := types.Instantiate(ctxt, named.Origin(), targs, false)
	if err != nil {
		panic(err)
	}
	rep := canon.Type(inst)
	obj, _, _ := types.LookupFieldOrMethod(rep, true, m.Pkg(), m.Name())
	return obj.(*types.Func)
}

// Exposed to ssautil using the linkname hack.
//
//go:linkname isSyntactic golang.org/x/tools/go/ssa.isSyntactic
func isSyntactic(pkg *Package) bool { return pkg.syntax }
```

## File: go/ssa/wrappers.go
```go
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// This file defines synthesis of Functions that delegate to declared
// methods; they come in three kinds:
//
// (1) wrappers: methods that wrap declared methods, performing
//     implicit pointer indirections and embedded field selections.
//
// (2) thunks: funcs that wrap declared methods.  Like wrappers,
//     thunks perform indirections and field selections. The thunk's
//     first parameter is used as the receiver for the method call.
//
// (3) bounds: funcs that wrap declared methods.  The bound's sole
//     free variable, supplied by a closure, is used as the receiver
//     for the method call.  No indirections or field selections are
//     performed since they can be done before the call.

import (
	"fmt"

	"go/token"
	"go/types"

	"golang.org/x/tools/internal/typeparams"
)

// -- wrappers -----------------------------------------------------------

// createWrapper returns a synthetic method that delegates to the
// declared method denoted by meth.Obj(), first performing any
// necessary pointer indirections or field selections implied by meth.
//
// The resulting method's receiver type is meth.Recv().
//
// This function is versatile but quite subtle!  Consider the
// following axes of variation when making changes:
//   - optional receiver indirection
//   - optional implicit field selections
//   - meth.Obj() may denote a concrete or an interface method
//   - the result may be a thunk or a wrapper.
func createWrapper(prog *Program, sel *selection) *Function {
	obj := sel.obj.(*types.Func)      // the declared function
	sig := sel.typ.(*types.Signature) // type of this wrapper

	var recv *types.Var // wrapper's receiver or thunk's params[0]
	name := obj.Name()
	var description string
	if sel.kind == types.MethodExpr {
		name += "$thunk"
		description = "thunk"
		recv = sig.Params().At(0)
	} else {
		description = "wrapper"
		recv = sig.Recv()
	}

	description = fmt.Sprintf("%s for %s", description, sel.obj)
	if prog.mode&LogSource != 0 {
		defer logStack("create %s to (%s)", description, recv.Type())()
	}
	/* method wrapper */
	return &Function{
		name:      name,
		method:    sel,
		object:    obj,
		Signature: sig,
		Synthetic: description,
		Prog:      prog,
		pos:       obj.Pos(),
		// wrappers have no syntax
		build:     (*builder).buildWrapper,
		syntax:    nil,
		info:      nil,
		goversion: "",
	}
}

// buildWrapper builds fn.Body for a method wrapper.
func (b *builder) buildWrapper(fn *Function) {
	var recv *types.Var // wrapper's receiver or thunk's params[0]
	var start int       // first regular param
	if fn.method.kind == types.MethodExpr {
		recv = fn.Signature.Params().At(0)
		start = 1
	} else {
		recv = fn.Signature.Recv()
	}

	fn.startBody()
	fn.addSpilledParam(recv)
	createParams(fn, start)

	indices := fn.method.index

	var v Value = fn.Locals[0] // spilled receiver
	if isPointer(fn.method.recv) {
		v = emitLoad(fn, v)

		// For simple indirection wrappers, perform an informative nil-check:
		// "value method (T).f called using nil *T pointer"
		if len(indices) == 1 && !isPointer(recvType(fn.object)) {
			var c Call
			c.Call.Value = &Builtin{
				name: "ssa:wrapnilchk",
				sig:  types.NewSignatureType(nil, nil, nil, types.NewTuple(anonVar(fn.method.recv), anonVar(tString), anonVar(tString)), types.NewTuple(anonVar(fn.method.recv)), false),
			}
			c.Call.Args = []Value{
				v,
				stringConst(typeparams.MustDeref(fn.method.recv).String()),
				stringConst(fn.method.obj.Name()),
			}
			c.setType(v.Type())
			v = fn.emit(&c)
		}
	}

	// Invariant: v is a pointer, either
	//   value of *A receiver param, or
	// address of  A spilled receiver.

	// We use pointer arithmetic (FieldAddr possibly followed by
	// Load) in preference to value extraction (Field possibly
	// preceded by Load).

	v = emitImplicitSelections(fn, v, indices[:len(indices)-1], token.NoPos)

	// Invariant: v is a pointer, either
	//   value of implicit *C field, or
	// address of implicit  C field.

	var c Call
	if r := recvType(fn.object); !types.IsInterface(r) { // concrete method
		if !isPointer(r) {
			v = emitLoad(fn, v)
		}
		c.Call.Value = fn.Prog.objectMethod(fn.object, b)
		c.Call.Args = append(c.Call.Args, v)
	} else {
		c.Call.Method = fn.object
		c.Call.Value = emitLoad(fn, v) // interface (possibly a typeparam)
	}
	for _, arg := range fn.Params[1:] {
		c.Call.Args = append(c.Call.Args, arg)
	}
	emitTailCall(fn, &c)
	fn.finishBody()
}

// createParams creates parameters for wrapper method fn based on its
// Signature.Params, which do not include the receiver.
// start is the index of the first regular parameter to use.
func createParams(fn *Function, start int) {
	tparams := fn.Signature.Params()
	for i, n := start, tparams.Len(); i < n; i++ {
		fn.addParamVar(tparams.At(i))
	}
}

// -- bounds -----------------------------------------------------------

// createBound returns a bound method wrapper (or "bound"), a synthetic
// function that delegates to a concrete or interface method denoted
// by obj.  The resulting function has no receiver, but has one free
// variable which will be used as the method's receiver in the
// tail-call.
//
// Use MakeClosure with such a wrapper to construct a bound method
// closure.  e.g.:
//
//	type T int          or:  type T interface { meth() }
//	func (t T) meth()
//	var t T
//	f := t.meth
//	f() // calls t.meth()
//
// f is a closure of a synthetic wrapper defined as if by:
//
//	f := func() { return t.meth() }
//
// Unlike createWrapper, createBound need perform no indirection or field
// selections because that can be done before the closure is
// constructed.
func createBound(prog *Program, obj *types.Func) *Function {
	description := fmt.Sprintf("bound method wrapper for %s", obj)
	if prog.mode&LogSource != 0 {
		defer logStack("%s", description)()
	}
	/* bound method wrapper */
	fn := &Function{
		name:      obj.Name() + "$bound",
		object:    obj,
		Signature: changeRecv(obj.Type().(*types.Signature), nil), // drop receiver
		Synthetic: description,
		Prog:      prog,
		pos:       obj.Pos(),
		// wrappers have no syntax
		build:     (*builder).buildBound,
		syntax:    nil,
		info:      nil,
		goversion: "",
	}
	fn.FreeVars = []*FreeVar{{name: "recv", typ: recvType(obj), parent: fn}} // (cyclic)
	return fn
}

// buildBound builds fn.Body for a bound method closure.
func (b *builder) buildBound(fn *Function) {
	fn.startBody()
	createParams(fn, 0)
	var c Call

	recv := fn.FreeVars[0]
	if !types.IsInterface(recvType(fn.object)) { // concrete
		c.Call.Value = fn.Prog.objectMethod(fn.object, b)
		c.Call.Args = []Value{recv}
	} else {
		c.Call.Method = fn.object
		c.Call.Value = recv // interface (possibly a typeparam)
	}
	for _, arg := range fn.Params {
		c.Call.Args = append(c.Call.Args, arg)
	}
	emitTailCall(fn, &c)
	fn.finishBody()
}

// -- thunks -----------------------------------------------------------

// createThunk returns a thunk, a synthetic function that delegates to a
// concrete or interface method denoted by sel.obj.  The resulting
// function has no receiver, but has an additional (first) regular
// parameter.
//
// Precondition: sel.kind == types.MethodExpr.
//
//	type T int          or:  type T interface { meth() }
//	func (t T) meth()
//	f := T.meth
//	var t T
//	f(t) // calls t.meth()
//
// f is a synthetic wrapper defined as if by:
//
//	f := func(t T) { return t.meth() }
func createThunk(prog *Program, sel *selection) *Function {
	if sel.kind != types.MethodExpr {
		panic(sel)
	}

	fn := createWrapper(prog, sel)
	if fn.Signature.Recv() != nil {
		panic(fn) // unexpected receiver
	}

	return fn
}

func changeRecv(s *types.Signature, recv *types.Var) *types.Signature {
	return types.NewSignatureType(recv, nil, nil, s.Params(), s.Results(), s.Variadic())
}

// A local version of *types.Selection.
// Needed for some additional control, such as creating a MethodExpr for an instantiation.
type selection struct {
	kind     types.SelectionKind
	recv     types.Type
	typ      types.Type
	obj      types.Object
	index    []int
	indirect bool
}

func toSelection(sel *types.Selection) *selection {
	return &selection{
		kind:     sel.Kind(),
		recv:     sel.Recv(),
		typ:      sel.Type(),
		obj:      sel.Obj(),
		index:    sel.Index(),
		indirect: sel.Indirect(),
	}
}

// -- instantiations --------------------------------------------------

// buildInstantiationWrapper builds the body of an instantiation
// wrapper fn. The body calls the original generic function,
// bracketed by ChangeType conversions on its arguments and results.
func (b *builder) buildInstantiationWrapper(fn *Function) {
	orig := fn.topLevelOrigin
	sig := fn.Signature

	fn.startBody()
	if sig.Recv() != nil {
		fn.addParamVar(sig.Recv())
	}
	createParams(fn, 0)

	// Create body. Add a call to origin generic function
	// and make type changes between argument and parameters,
	// as well as return values.
	var c Call
	c.Call.Value = orig
	if res := orig.Signature.Results(); res.Len() == 1 {
		c.typ = res.At(0).Type()
	} else {
		c.typ = res
	}

	// parameter of instance becomes an argument to the call
	// to the original generic function.
	argOffset := 0
	for i, arg := range fn.Params {
		var typ types.Type
		if i == 0 && sig.Recv() != nil {
			typ = orig.Signature.Recv().Type()
			argOffset = 1
		} else {
			typ = orig.Signature.Params().At(i - argOffset).Type()
		}
		c.Call.Args = append(c.Call.Args, emitTypeCoercion(fn, arg, typ))
	}

	results := fn.emit(&c)
	var ret Return
	switch res := sig.Results(); res.Len() {
	case 0:
		// no results, do nothing.
	case 1:
		ret.Results = []Value{emitTypeCoercion(fn, results, res.At(0).Type())}
	default:
		for i := 0; i < sig.Results().Len(); i++ {
			v := emitExtract(fn, results, i)
			ret.Results = append(ret.Results, emitTypeCoercion(fn, v, res.At(i).Type()))
		}
	}

	fn.emit(&ret)
	fn.currentBlock = nil

	fn.finishBody()
}
```
