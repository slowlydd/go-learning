package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	a, b := m1[3]
	t.Log(a) // key的值
	t.Log(b) // key是否存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
		t.Log(ok)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	k, v := m1[0]
	t.Log(k)
	t.Log(v)

	arr := [...]int{1, 2, 3, 4}
	for k, v := range arr { // 此处的k, v与上面的k, v是不同的作用域
		t.Log(k, v)
	}
}
