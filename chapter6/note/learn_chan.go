package main

import (
	"log"
	"os"
	"sync"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Llongfile)
}

func run(prefix string, signal chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case result, running := <-signal:
			{
				if !running {
					log.Printf("%s : receive closed channel\n", prefix)
					return
				} else {
					result++
					if result > 8 {
						close(signal)
						log.Printf("%s over eight, close channel\n", prefix)
						break
					}
					signal <- result
					log.Printf("%s send plus one : %d\n", prefix, result)
				}
			}
		default:
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var signal chan int = make(chan int)
	go run("A", signal, &wg)
	go run("B", signal, &wg)
	signal <- 1
	wg.Wait()
	log.Println("over")

}
