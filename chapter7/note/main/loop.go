package main

import (
	"fmt"
)

var result = 0

func add(stop chan bool, step chan int, over chan bool) {
	for {
		select {
		case <-stop:
			over <- true
			return
		case adder := <-step:
			result += adder
		}
	}
}

func reduce(start int, stop chan bool, step chan int) {
	if start == 0 {
		stop <- true
		return
	}
	step <- start
	go reduce(start-1, stop, step)
}

func main() {
	stop := make(chan bool)
	step := make(chan int)
	over := make(chan bool)
	go add(stop, step, over)
	go reduce(2, stop, step)
	<-over
	fmt.Println(result)

}
