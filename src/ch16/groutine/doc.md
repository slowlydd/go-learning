### Thead vs. Groutine

1.创建时默认的stack的大小
+ JDK5 以后 Java Thread stack 默认为1M
+ Groutine 的 Stack 初始化大小为 2K

2. 和 KSE （Kernel Space Entity）的对应关系
  
+ Java Thread 是 1：1
+ Groutine 是 M:N