package main

import "fmt"

type Liver struct {
	Age int
}

type Person struct {
	Liver
	Name string
	Age  int
}

func main() {
	person := Person{
		Liver: Liver{
			Age: 34,
		},
		Name: "godme",
		Age:  45,
	}
	fmt.Println(person.Age)
	fmt.Println(person.Liver.Age)
}
