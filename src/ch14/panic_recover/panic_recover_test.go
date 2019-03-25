package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
	panic 通常用于不可回复的错误
	panic 退出前会执行 defer 指定的函数
	os.Exit 退出时不会调用 defer 指定的函数
	os.Exit 退出时不输出当前调用栈信息
*/

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong!"))
	// os.Exit(-1)
}
