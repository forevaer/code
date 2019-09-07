package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	fmt.Println(filepath.Abs(filepath.Dir(".")))
}

func main() {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()
	bs := make([]byte, size)
	_, _ = file.Read(bs)
	fmt.Println(string(bs))
}
