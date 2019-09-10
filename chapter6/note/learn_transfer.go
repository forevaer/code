package main

import "fmt"

type Person struct {
	Name string
}

func address(obj interface{}, point bool) {
	// 如果是
	if point {
		fmt.Printf("type : %T, isPoint : %v ,\t address : %x \n", *obj.(*int), point, obj)
	} else {
		fmt.Printf("type : %T, isPoint : %v ,\t address : %x  \n", obj, point, &obj)
	}
}
func main() {
	a := 4
	fmt.Printf("type : %T, isPoint : false ,\t address : %x  \n", a, &a)
	address(a, false)
	address(&a, true)
}
