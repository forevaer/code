package main

import "fmt"

type Speaker interface {
	Say()
}

type Dog struct {
}

func (dog *Dog) Say() {
	fmt.Println("旺旺")
}
