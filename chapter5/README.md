流拼接
---
```go
func main() {
	var b bytes.Buffer
	// 写入
	b.Write([]byte("Hello"))
	// 追加到输出流
	_, _ = fmt.Fprintf(&b, "World!")
	//   对接到输出
	_, _ = io.Copy(os.Stdout, &b)
}
```
流对接
---
```go
func main() {
	// 第一个参数为网址
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// 输入流对接少输出流，直接输出
	_, _ = io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
```
处理器
---
强调针对接口，而非具体对象，我们可以使用处理器的方式来进行函数运行<br>
```go
type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}
```

类型拼接
---
- 结构拼接
```go
type Liver struct {
	Age int
}

type Person struct {
	Liver
	Name string
}

func main() {
	person := Person{
		Liver: Liver{
			Age: 34,
		},
		Name: "godme",
	}
	fmt.Println(person.Age)
	fmt.Println(person.Liver.Age)
}

}
```
没有继承，拼接之后的元素有点奇怪<br>
1.申明的时候必须声明拼接的结构<br>
2.使用的时候可以直接访问到拼接的内部数据<br>
3.通过访问内部对象再进行字段访问也可以
> 如果结构中存在字段，拼接的结构中也有同名字段，没有指明的情况下，使用自身字段<br>
>只有当前层级查找不到，才会尝试向下查找

- 接口拼接

```go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
```
多个动作组合
- 类型扩展
```go
type String string

func (s String) Show() {
	fmt.Println(s)
}
func main() {
	var name String = "godme"
	name.Show()
}
```
如果直接引用的是基本类型，使用上没有什么变化，但有如下两个作用<br>
1. 名称标识
2. 方法扩展
一般来说，对基本类型进行方法扩展的场景并不很多，但是对基本类型进行重新标识的场景有很多，比如
``time.Duration``

接口传递
---
```go
type Man struct {
}

type Sayer interface {
	Say(string)
}

func (someone *Man) Say(msg string) {
	fmt.Println(msg)
}

func main() {
	var sayer Sayer = &Man{}
	sayer.Say("hello")
}
```
一般要使用接口的话，实例化总需要实体<br>
直接把实体传递给接口会报错，一般使用指针进行传递