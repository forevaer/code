// This sample program demonstrates how to use a channel to
package main

import (
	"log"
	"os"
	"time"

	"code/chapter7/patterns/runner"
)

// timeout
const timeout = 30 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)
	// tasks
	r.Add(createTask(), createTask(), createTask())

	// start
	if err := r.Start(); err != nil {
		switch err {
		// wait timeout
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
			// wait interrupt
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

// sleep is task
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
