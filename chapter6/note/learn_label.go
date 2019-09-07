package main

import (
	"fmt"
	"sync"
)

func show(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 0; index < 5; index++ {
		fmt.Println(prefix, index)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go show("a", &wg)
	go show("b", &wg)
	wg.Wait()
}
