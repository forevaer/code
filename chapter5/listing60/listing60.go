// Sample program to show what happens when the outer and inner
// type implement the same interface.
package main

import (
	"fmt"
)

// 接口
type notifier interface {
	notify()
}

// 对象
type user struct {
	name  string
	email string
}

// 实现
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// 对象
type admin struct {
	user
	level string
}

// 实现
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// 入口
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 首先执行自身，如果找不到方法，搜索扩展
	sendNotification(&ad)

	// 执行
	ad.user.notify()

	// 自身方法
	ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
