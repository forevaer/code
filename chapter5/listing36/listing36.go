// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

//  接口定义
type notifier interface {
	notify()
}

// 用户结构
type user struct {
	name  string
	email string
}

// 方法实现
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// 入口
func main() {
	// instance
	u := user{"Bill", "bill@email.com"}
	// process
	// 接收接口时使用指针比较好
	sendNotification(&u)
}

// 外部方法处理
func sendNotification(n notifier) {
	n.notify()
}
