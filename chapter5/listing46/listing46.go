// Sample program to show how you can't always get the
// address of a value.
package main

import "fmt"

// 类型重命名
type duration int

// 添加方法
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// 入口
func main() {
	duration(42).pretty()
}
