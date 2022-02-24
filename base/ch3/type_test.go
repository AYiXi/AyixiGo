package ch3

import "testing"

type MyInt int64

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64 = int64(a)

	var c MyInt = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPointer := &a

	// go 不支持指针运算
	// aPointer = aPointer + 1
	t.Log(a, aPointer)
	t.Logf("%T %T", a, aPointer)
}
