特殊函数
---
- 入口函数（``main``）
```go
func main(){
    fmt.Println("start")
}
```
- 初始函数（``init``）
```go
func init(){
    fmt.Println("init")
}
```
标准库
---
- os
```go
stdout := os.Stdout // 标准输出
stderr := os.Stderr // 标准错误
stdin  := os.Stdin  // 标准输入
```
- log
```go
log.SetOutput(stdout) // 设置输出，可以指向文件
log.Println("")       // 同fmt输出
log.Fatalln("")       // 输出并报错
```
导入技巧
---
```go
import (
    "log"
    io "bufio"
    . "os"
    _ "encoding/json"
)
```
| 前缀      | 说明     | 用途                                                       |
| --------- | -------- | ---------------------------------------------------------- |
| ``""``    | 基本导入 | 导入                                                       |
| ``alias`` | 别名导入 | 自定义名称<br />避免同名包冲突                             |
| ``_``     | 忽略导入 | 导入未使用的包<br />使得编译不报错                         |
| ``.``     | 全部导入 | 导入指定模块全部方法(变量)<br />相当于``import package.*`` |

访问权限
---
模块之间是可见的，但是模块内部的变量/方法，如果想要外部模块也可以访问的话，需要进行特殊设置。
```go
func Say(){ // 首字母大写，外部模块可见
    fmt.Println("hello")
}

func see(){ // 首字母小写，外部模块不可见
    fmt.Println("beauty")
}
```
所以，我们经常调用的模块所使用的方法，一定都是大写。

运行
---
``default.go``在``serach``包下面，已经注册上了<br>
但是对于``rss``，没有导入对应的包，所以直接运行是没有结果的<br>
需要忽略导入一下，只是为了执行一下``init``
```go
import _ "code/chapter2/sample/matchers"
```