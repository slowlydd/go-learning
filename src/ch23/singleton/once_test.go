package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleleton struct {
	data string
}

var singleInstance *Singleleton
var once sync.Once

func GetSingletonObj() *Singleleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// 这里每一个协程都会执行这个代码但是却只有第一个会创建
			// 对象，所以每次打印出来的地址都是一样的，因为没有新的对象创建
			obj := GetSingletonObj()
			fmt.Printf("%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
