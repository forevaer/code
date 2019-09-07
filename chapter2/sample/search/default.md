``default.go``
---
- 对象方法
```go
type Man struct {
	Name string
}
func (man *Man) Hello(name string) string {
	return fmt.Sprintf("hello %s, my name is %s, nice to meet you !", name, man.Name)
}
func main() {
	man := Man{Name: "godme"}
	fmt.Println(man.Hello("judas"))
}
```
一般模板：
```go
func(class) funcName(parameters) (result){
}
```
不指定类型的话就是一般方法来。
这种方法使用起来和一般面向对象没有差异，但是实现上面，看起来更像是``外挂``上去的。

接口实现
---
```go
type Speaker interface {
	Say()
}

type Dog struct {
}

func (dog *Dog) Say() {
	fmt.Println("旺旺")
}
```
可以看到，我们是在不能够直观看出一个类型到底实现来哪个接口。
只要一个对象的方法名称和类型匹配上了，并且包含一个接口的全部方法，都认定该对象实现来该接口。
> 可以感受到比较麻烦的一点，正确的对象：``属性``  + ``方法`` <br>
> 方法都是外挂上去的，所以对于``struct``而言，哪怕是空实现，也是必须的<br>
>表明该队行没有属性，更主要的，是提供方法的外挂