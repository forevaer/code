package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // Number of goroutines to use.
	taskLoad         = 10 // Amount of work to process.
)

var wg sync.WaitGroup

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	// goroutine
	wg.Add(numberGoroutines)
	// create goroutine
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// create tasks
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// close tasks
	close(tasks)

	// wait work done
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	// task was done
	defer wg.Done()

	for {
		// not running, goroutine was close
		task, ok := <-tasks
		if !ok {
			// close
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// task
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// random sleep
		// this is work
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// finish task
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
