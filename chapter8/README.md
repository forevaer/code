log
---
```go
log.SetOutput(os.Stdout)
log.SetPrefix("prefix")
log.SetFlags(log.Llongfile)
/**
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
*/
log.Println()
log.Fatalln()
log.Paincln()
```

json
---
```go
json.Marshal(godme)
json.Ummarshal(bs, Godme)
json.MarshalIndent(godme,"prefix", "suffix")
```
json也可以直接使用``map``进行格式，参见
-  解析[demo](./listing29/listing29.go)
- 格式化[demo](./listing31/listing31.go)

buffer
---
```go
func main() {
	// 声明即可使用
	var b bytes.Buffer
    // 字节写入
	b.Write([]byte("Hello "))
	// Fprintf 追加写入
	fmt.Fprintf(&b, "World!")
    // 输出到指定输出
	b.WriteTo(os.Stdout)
}
```

multiWriter
```go
	dest := io.MultiWriter(os.Stdout, file)
	_, _ = io.Copy(dest, r.Body)
```
想要同时输出到多个流，可以这样来