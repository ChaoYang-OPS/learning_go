package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writeable
	Execable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	fmt.Println(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111  true true true
	// a := 1 //0001  true false false
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Execable == Execable)
	fmt.Println(a&Readable == Readable, a&Writeable == Writeable, a&Execable == Execable)
}
