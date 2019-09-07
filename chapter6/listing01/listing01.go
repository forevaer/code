package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 并发任务队列
	runtime.GOMAXPROCS(1)
	// 阻塞
	var wg sync.WaitGroup
	// 2
	wg.Add(2)
	fmt.Println("Start Goroutines")
	// go
	go func() {
		// decrease
		defer wg.Done()

		// lowercase
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// go
	go func() {
		// decrease
		defer wg.Done()
		// uppercase
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// wait
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
