package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var value int32
	atomic.AddInt32(&value, 5)
	atomic.StoreInt32(&value, 5)
	atomic.LoadInt32(&value)

	var lock sync.Mutex
	lock.Lock()
	lock.Unlock()
}
