package fib

import "testing"

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1

	var c int

	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	c = 3
	t.Log(c)
}
