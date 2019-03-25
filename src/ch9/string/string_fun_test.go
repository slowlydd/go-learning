package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for s, part := range parts {
		t.Log(s)
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(101) // strconv.Itoa 将整形强制装换为字符串
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil { // strconv.Atoi 将字符串强制转换为整形
		t.Log(10 + i)
	}
}
