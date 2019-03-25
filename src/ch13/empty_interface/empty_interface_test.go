package empty_interface_test

/*
	空接口与断言
	1.空接口可以表示任何类型
	2.通过断言来将空接口转换为指定类型
	v, ok := p.(int) // ok=true 时转换成功
*/

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
