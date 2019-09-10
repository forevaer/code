package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// count
	counter int
	// wg
	wg sync.WaitGroup
)

func main() {
	// double
	wg.Add(2)

	// first
	go incCounter(1)
	// second
	go incCounter(2)

	// wait
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	// done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// save counter to value
		value := counter
		fmt.Printf("%d : %d : %d\n", id, count, value)
		// yield
		runtime.Gosched()
		fmt.Printf("%d : %d : %d\n", id, count, value)
		// plus
		value++
		// counter update
		counter = value
	}
	// two goroutines use common resource counter
}
