package main

import (
	"encoding/json"
	"fmt"
)

// 结构
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 方法
func (person *Person) equals(obj interface{}) bool {
	other := obj.(Person)
	return person.Name == other.Name && person.Age == other.Age
}

func main() {
	person := Person{
		Name: "judas",
		Age:  23,
	}
	// 序列
	bs, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(bs))
	}
	var another Person
	// 反序列
	_ = json.Unmarshal(bs, &another)
	fmt.Println(person.equals(another))
}
