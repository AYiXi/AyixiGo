package ch2

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestContast(t *testing.T) {
	t.Log()
}

func TestContast2(t *testing.T) {
	a := 1
	t.Log(a & Readable, a & Executable)
	t.Log(a & Readable == Readable, a & Writable == Writable, a & Executable == Executable)
} 