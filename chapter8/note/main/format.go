package main

import (
	"encoding/json"
	"fmt"
)

type Godme struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	godme := &Godme{
		Name:   "godme",
		Age:    12,
		Gender: "male",
	}
	format, _ := json.MarshalIndent(godme, "", "\t")
	fmt.Println(string(format))
}
