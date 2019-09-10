package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"code/chapter7/patterns/pool"
)

const (
	maxGoroutines   = 25 // the number of routines to use.
	pooledResources = 2  // number of resources in the pool
)

// closer instance
type dbConnection struct {
	ID int32
}

// implements
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// const
var idCounter int32

// static createConnect
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// create pool
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		// some not of single
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Close the pool.
	log.Println("Shutdown Program.")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	// get Conn
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	// release conn
	defer p.Release(conn)
	// sleep is work
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
