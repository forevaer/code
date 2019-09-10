# go模式归纳

## 变量声明

```go
const (
    target = "godme"
)

var (
  GError = errors.New("Godme Error")
)
```

> 相较于``java``，可以当作是类变量。
>
> 因为整个文件都属于包，以包为最小包括单位，而不是类。

## 属性声明

```go
type Godme struct {
  Name string
}
```

> 这里面也有字段声明，和上面的区别在于两点。
>
> 1. 属性字段跟随对象，包变量跟随包
> 2. 属性字段方法使用，包变量随便用
>
> 也就是说，当对象方法中不得不使用的量，我们才考虑放到结构体中。
>
> 至于统一使用的变量，我们一般当作类属性。
>
> > 日志``log``的话，需要外部自定义，所以可以跟随变量
> >
> > 锁``mutex``的话，只是为了保证流程有序，跟随变量更好，避免空间占用

## 构造方法

```go
func NewGodme(name string) *Godme {
  return &Godme{
    Name:name
  }
}
```

> 构造方法不是必须，但是比较规范，作为初始化方法

## 实现方法

```go
func (godme *Godme) Run() {
  if godme.Name == target {
    panic(GError)
  }
  fmt.Printf("Hello, %s, nice to meet you !\n", godme.Name)
}
```

> 方法中的变量，大多都是属性量相关