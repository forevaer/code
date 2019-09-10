package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	// set random seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	// 结束
	defer wg.Done()

	for {
		// 接收到
		ball, ok := <-court
		if !ok {
			// 关闭，对方认输
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 回应一个数
		n := rand.Intn(100)
		// 特定情况接不到球
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 接不到球就认输，关闭通道
			close(court)
			return
		}

		// 回球成功
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
