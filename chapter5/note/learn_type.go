package main

import "fmt"

type String string

func (s String) Show() {
	fmt.Println(s)
}
func main() {
	var name String = "godme"
	name.Show()
}
