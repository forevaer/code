package main

import (
	"errors"
	"fmt"
)

var (
	GError = errors.New("gError")
)

type Godme struct {
	Name string
}

func NewGodme(name string) *Godme {
	return &Godme{Name: name}
}

func (godme *Godme) Run() {
	if godme.Name == "godme" {
		panic(GError)
	}
	fmt.Printf("hello %s\n", godme.Name)
}
