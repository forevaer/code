package main

import "fmt"

type Man struct {
}

type Sayer interface {
	Say(string)
}

func (someone *Man) Say(msg string) {
	fmt.Println(msg)
}

func main() {
	var sayer Sayer = &Man{}
	sayer.Say("hello")
}
