package main

import (
	"fmt"
)

// 结构定义
type user struct {
	name  string
	email string
}

// 输出信息
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// 替换属性
func (u *user) changeEmail(email string) {
	u.email = email
}

// 入口
func main() {
	// 创建并打印
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 创建并打印
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// 修改并打印
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 修改并打印
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
