package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Hello(name string) string {
	return fmt.Sprintf("hello %s, my name is %s, nice to meet you !", name, man.Name)
}
func main() {
	man := Man{Name: "godme"}
	fmt.Println(man.Hello("judas"))
}
