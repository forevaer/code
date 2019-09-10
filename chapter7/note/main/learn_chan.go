package main

import (
	"fmt"
	"time"
)

type Data struct {
	tasks chan int
}

func (d *Data) add(task int) {
	d.tasks <- task
}

func New() *Data {
	return &Data{
		tasks: make(chan int),
	}
}

var sleep = 5
var timeStamp = 0

func main() {
	data := New()
	go func(a *Data) {
		for s := range a.tasks {
			fmt.Println(s)
		}
	}(data)
	time.Sleep(5 * time.Second)

	for index := 0; index < 10; index++ {
		data.add(index)
	}

}
