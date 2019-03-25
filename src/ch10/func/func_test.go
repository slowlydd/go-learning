package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 可以返回多个值, rand.Intn(10) 产生随机数
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数是一等公民， 可以作为参数和返回值
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// 延迟函数
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 不定个数的参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// defer 延迟执行函数，panic 出现无法修复的错误，
// 会在所在作用域函数的返回值返回之前执行，就算发生错误也会执行，类似finally
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
