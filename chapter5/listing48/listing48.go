// Sample program to show how polymorphic behavior with interfaces.
package main

import (
	"fmt"
)

// 接口
type notifier interface {
	notify()
}

// 用户
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
	name  string
	email string
}

// 实现
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// 入口
func main() {
	// user
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// admin
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// exec
func sendNotification(n notifier) {
	n.notify()
}
