// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import (
	"fmt"
)

// user
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

// 扩展
type admin struct {
	user  // Embedded Type
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

	// 内部对象调用
	ad.user.notify()

	// 即是内部对象，也是自身属性
	ad.notify()
}
