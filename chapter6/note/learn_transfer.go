package main

import "fmt"

type Person struct {
	Name string
}

func address(obj interface{}, point bool) {
	if point {
		fmt.Printf("type : %T, address : %x \n", *obj.(*int), obj)
	} else {
		fmt.Printf("type : %T, address : %x  \n", obj, &obj)
	}
}
func main() {
	a := 4
	fmt.Printf("type : %T, address : %x  \n", a, &a)
	address(a, false)
	address(&a, true)
}
