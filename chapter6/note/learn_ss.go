package main

import "fmt"

func main() {
	for index := 2; index <= 10; index++ {
		if IsS(index) {
			fmt.Println(index)
		}
	}

}

func IsS(value int) bool {
	for index := 2; index < value; index++ {
		if value%index == 0 {
			return false
		}
	}
	return true
}
