// Sample program to show how embedded types work with interfaces.
package main

import (
	"fmt"
)

// 接口
type notifier interface {
	notify()
}

// 结构
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

// admin
type admin struct {
	user
	level string
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

	// 指针
	sendNotification(&ad)
}

// process
func sendNotification(n notifier) {
	n.notify()
}
