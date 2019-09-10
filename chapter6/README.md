参数传递
---
- 引用传递<br>
``java``等语言中，对象之间的传递，似乎都是引用传递。<br>
因而，我们能够操作对象，把结果返回出来。<br>
引用传递的好处在于，只存在一个实体，减少内存占用和复制开销
- 值传递
详情查看[demo](./note/learn_transfer.go)，可以看到。<br>
方法之间进行的是值传递，也就是说，每次的方法调用，都会进行数据拷贝操作。<br>
这样会加大内存的开销，频繁的拷贝，也会消耗系统资源。<br>
可以说，``go``把问题暴露出来了，我们应该自己去把控了。
> 关于引用问题，可以对比一下[sync](./note/learn_sync.go)和[sync2](./note/learn_sync2.go)

原子操作
---
- atomic
```go
var value int32
atomic.AddInt32(&value, 88)
atomic.StoreInt32(&value, 88)
value = atomic.LoadInt32(&value)
```
- mutex
```go
var lock sync.Mutex
lock.Lock()
lock.UnLock()
```

调度
---
```go
runtime.Gosched()
```
礼让，相当于``java``中的``yield``
退出当前的执行状态，让出时间片

接收
---
```go
data,running := <-signal
// data     传输的数据 
// running  运行状态，被关闭为false
```

