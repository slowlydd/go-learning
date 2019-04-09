### 函数

函数是一等公民

与其他主要编程语言的差异

1.可以有多个返回值
2.所有参数都是值传递：slice，map，channel会有传引用的错觉
3.函数可以作为变量的值
4.函数可以作为参数和返回值

可变参数

延迟执行， defer 函数， 在返回之前执行。类似try， catch， finally中的finally，常用于安全的释放资源，释放锁

panic 程序异常中断，通常指不可修复的

### 行为的定义和实现

封装数据和行为

结构体定义

type Employee struct {
  Id string
  Name string
  Age int
}

### Go语言的相关接口

duck type 接口的实现不依赖接口的定义


#### Go 接口

1.接口为非入侵性，实现不依赖于接口定义
2.所以接口的定义可以包含在接口使用者包内

#### 接口变量

自定义类型

type 

### 扩展与复用

package

1.基本复用模块单元
  以首字母大写来表明可被包外代码访问
2.代码的package可以和所在的目录不一致
3.同一目录里的Go代码的package要保持一致

### channel 的关闭

+ 向关闭的 channel 发送数据，会导致 panic
+ v,ok <-ch; ok 为 bool 值，true 表示正常接受，false 表示通道关闭
+ 所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 ok 值为 false。这个广播机制常被利用，进行向多个订阅者同时发送信息。如：退出信号。

### Context

+ 根 Context：通过 context.Background() 创建
+ 子 Context：通过 context.WithCancel(parentContext) 创建
  - ctx, cancel := context.WithCancel(context.Background())
+ 当前 Context 被取消时，基于他的子 context 都会被取消
+ 接收取消通知 <-ctx.Done()




