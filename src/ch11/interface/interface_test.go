package interface_test

import "testing"

// duke type 鸭子类型

/*
	接口为非入侵性，实现不依赖于接口定义
	所以接口的定义可以包含在接口的使用者包内
*/

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
