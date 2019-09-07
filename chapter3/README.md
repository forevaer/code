读取文件
---
```go
func main() {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()
	bs := make([]byte, size)
	_, _ = file.Read(bs)
	fmt.Println(string(bs))
}
```
[demo](./note/learn_file.go)

字符串分割
---
```go
func main() {
	source := "   first second third"
	spaceSplit := strings.Fields(source)
	fmt.Println(spaceSplit)

	customerSpliter := func(rs ...rune) func(r rune) bool {
		return func(r rune) bool {
			for _, rr := range rs {
				if rr == r {
					return true
				}
			}
			return false
		}
	}
	customerResult := strings.FieldsFunc(source, customerSpliter('t', 'n', 'r', ' '))
	fmt.Println(customerResult)
}
```
[demo](./note/learn_spilt.go)