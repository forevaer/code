package main

import (
	"fmt"

	"code/chapter5/listing68/counters"
)

func main() {
	// create
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
