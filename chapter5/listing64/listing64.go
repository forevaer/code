// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
)

// 别名
type alertCounter int

// 入口
func main() {
	// 方法依旧
	counter := alertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
