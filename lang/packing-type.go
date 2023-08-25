package herry

import "strings"

type Comparable interface {
	CompareTo(other interface{}) int
}

type Byte byte

func (b Byte) CompareTo(other interface{}) int {
	if c, ok := other.(Byte); ok {
		if b < c {
			return -1
		} else if b > c {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Byte")
}

type Rune rune

func (r Rune) CompareTo(other interface{}) int {
	if s, ok := other.(Rune); ok {
		if r < s {
			return -1
		} else if r > s {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Rune")
}

type Int int

func (i Int) CompareTo(other interface{}) int {
	if j, ok := other.(Int); ok {
		if i < j {
			return -1
		} else if i > j {
			return 1
		} else {
			return 0
		}
	}
	panic("not an Int")
}

type Int8 int8

func (i Int8) CompareTo(other interface{}) int {
	if j, ok := other.(Int8); ok {
		if i < j {
			return -1
		} else if i > j {
			return 1
		} else {
			return 0
		}
	}
	panic("not an Int8")
}

type Int16 int16

func (i Int16) CompareTo(other interface{}) int {
	if j, ok := other.(Int16); ok {
		if i < j {
			return -1
		} else if i > j {
			return 1
		} else {
			return 0
		}
	}
	panic("not an Int16")
}

type Int32 int32

func (i Int32) CompareTo(other interface{}) int {
	if j, ok := other.(Int32); ok {
		if i < j {
			return -1
		} else if i > j {
			return 1
		} else {
			return 0
		}
	}
	panic("not an Int32")
}

type Int64 int64

func (i Int64) CompareTo(other interface{}) int {
	if j, ok := other.(Int64); ok {
		if i < j {
			return -1
		} else if i > j {
			return 1
		} else {
			return 0
		}
	}
	panic("not an Int64")
}

type Uint uint

func (u Uint) CompareTo(other interface{}) int {
	if v, ok := other.(Uint); ok {
		if u < v {
			return -1
		} else if u > v {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Uint")
}

type Uint8 uint8

func (u Uint8) CompareTo(other interface{}) int {
	if v, ok := other.(Uint8); ok {
		if u < v {
			return -1
		} else if u > v {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Uint8")
}

type Uint16 uint16

func (u Uint16) CompareTo(other interface{}) int {
	if v, ok := other.(Uint16); ok {
		if u < v {
			return -1
		} else if u > v {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Uint16")
}

type Uint32 uint32

func (u Uint32) CompareTo(other interface{}) int {
	if v, ok := other.(Uint32); ok {
		if u < v {
			return -1
		} else if u > v {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Uint32")
}

type Uint64 uint64

func (u Uint64) CompareTo(other interface{}) int {
	if v, ok := other.(Uint64); ok {
		if u < v {
			return -1
		} else if u > v {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Uint64")
}

type Float32 float32

func (f Float32) CompareTo(other interface{}) int {
	if g, ok := other.(Float32); ok {
		if f < g {
			return -1
		} else if f > g {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Float32")
}

type Float64 float64

func (f Float64) CompareTo(other interface{}) int {
	if g, ok := other.(Float64); ok {
		if f < g {
			return -1
		} else if f > g {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Float64")
}

type Complex64 complex64

func (c Complex64) CompareTo(other interface{}) int {
	if d, ok := other.(Complex64); ok {
		if real(c) < real(d) || (real(c) == real(d) && imag(c) < imag(d)) {
			return -1
		} else if real(c) > real(d) || (real(c) == real(d) && imag(c) > imag(d)) {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Complex64")
}

type Complex128 complex128

func (c Complex128) CompareTo(other interface{}) int {
	if d, ok := other.(Complex128); ok {
		if real(c) < real(d) || (real(c) == real(d) && imag(c) < imag(d)) {
			return -1
		} else if real(c) > real(d) || (real(c) == real(d) && imag(c) > imag(d)) {
			return 1
		} else {
			return 0
		}
	}
	panic("not a Complex128")
}

type String string

func (s String) CompareTo(other interface{}) int {
	if t, ok := other.(String); ok {
		return strings.Compare(string(s), string(t))
	}
	panic("not a String")
}

type Boolean bool

func (b Boolean) CompareTo(other interface{}) int {
	if c, ok := other.(Boolean); ok {
		if b == c {
			return 0
		} else if b {
			return 1
		} else {
			return -1
		}
	}
	panic("not a Boolean")
}
