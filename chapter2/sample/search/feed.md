结构对照
===
```go
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}
```
定义的结构体和转换格式的对照关系,对照名称尤其注意，可能存在差异。
<br>
如果不想显示默认(空)数据的话，可以这样进行申明
```go
type Person struct {
	Name string `json:"site"`
	Age  string `json:"link,omitempty"`
}
```

文件操作
---
```go
func main() {
	filePath := "data/data.json"
	file, err := os.Open(filePath)    // 打开文件句柄
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()              // 关闭文件句柄
	stat, _ := file.Stat()          // 文件状态
	length := stat.Size()           // 数据长度
	bs := make([]byte, length)      // 字节
	_,_ = file.Read(bs)             // 读取
	println(string(bs))
}
```
可以查看[小例子](./../note/learn_file.go)了解一下

``json``
---
```go
bs, err := json.Marshal(person)  // 对象转字节
_ = json.Unmarshal(bs, &another) // 字节解析对象
```
序列化  ：对象直接转化成字节，这个使用上面和其他的面向对象都一致<br>
反序列化：``go``中含有大量的引用传递，计算的结果不一定会返回，可能要求你指针传入对象，通过指针直接修改进行接收
详细请查看[小例子](./../note/learn_json.go)了解一下