知识点总结
===

## ``for``
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
## ``channel``
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
# ``goroutine``
```go
go printer(c)
```
自动开启协程执行
相较于其他语言来说
1. 天然支持并发，无需为维护线程池一类的辅助操作花费心思
2. 协程体量小，可轻松开辟上百万，消耗资源低