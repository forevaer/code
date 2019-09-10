package main

import (
	"log"
	"sync"
	"time"

	"code/chapter7/patterns/work"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(len(names))

	for i := 0; i < 1; i++ {
		for _, name := range names {
			// create namePrinter
			np := namePrinter{
				name: name,
			}

			go func() {
				// add 2 task
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
