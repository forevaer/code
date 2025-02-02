package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton
	// start
	fmt.Printf("Runner %d Running With Baton\n", runner)

	if runner != 4 {
		// newRunner ,plus
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		// newRunner,  递归
		go Runner(baton)
	}
	// sleep
	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	// for old 2 new
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)
	// signal
	baton <- newRunner
}
