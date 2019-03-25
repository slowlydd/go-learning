package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 可以对数组进行比较，要求每个位置上对应的值相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
