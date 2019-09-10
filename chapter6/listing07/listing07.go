package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	// wg
	var wg sync.WaitGroup
	// double goroutines
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// first
	go func() {
		defer wg.Done()
		// lowerCase
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// second
	go func() {
		defer wg.Done()
		// upperCase
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
