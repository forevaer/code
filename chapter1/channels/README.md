``for``
 ---
 - 循环
 ```go
for index := 0; index < 10; index++ {
    fmt.Println(index)
}
```
同``java``等语言的一般``for``循环
- ``while``
```go
times := 5
func main() {
	times := 5
	for times > 0 {
		fmt.Println(times)
		times--
	}
}
```
不带条件，默认无限循环
```go
for {
    fmt.Println("forever")
}
```
- 遍历
```go
func main() {
	arr := []int{1, 2, 3, 4, 5}
	for index, data := range arr {
		fmt.Printf("%d : %d \n", index, data)
	}
}
```
[例子](./note/learn_for.go)
 ``channel``
 ---
- ``create``
```go
ch := make(chan int)
```
携带的可以是任意其他类型
- ``close``
```go
close(ch)
```
每个通道都应该有一个关闭的操作
- ``write``
```go
ch <- obj
```
- ``read``
```go
read := <-ch
```
这会引起阻塞，具体后续讨论
``goroutine``
---
```go
go printer(c)
```
自动开启协程执行
相较于其他语言来说
1. 天然支持并发，无需为维护线程池一类的辅助操作花费心思
2. 协程体量小，可轻松开辟上百万，消耗资源低

同步组
---
```go
func mian(){
    var wg sync.WaitGroup
    wg.Add(1)   //计数加
    wg.Done()   //计数减1， Add(-1  )
    wg.Wait()   //阻塞，直到计数为0
}
```
常用于协程间控制