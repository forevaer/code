package main

import (
	"fmt"
	"sync"
)

func single(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 0; index < 5; index++ {
		fmt.Println(prefix, index)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go single("a", &wg)
	go single("b", &wg)
	wg.Wait()
	fmt.Println("over")
}
