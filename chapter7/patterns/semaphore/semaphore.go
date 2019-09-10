package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	semaphore chan struct{}
)

type (
	readerWriter struct {
		// name
		name string

		// wg
		write sync.WaitGroup

		// chan struct
		readerControl semaphore

		// chan struct.
		shutdown chan struct{}

		// wg
		reportShutdown sync.WaitGroup

		maxReads int

		maxReaders int

		currentReads int32
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log.Println("Starting Process")

	//
	first := NewReadWriter("First", 3, 6)

	//
	second := NewReadWriter("Second", 2, 2)

	// sleep
	time.Sleep(2 * time.Second)

	// shut down
	shutdown(first, second)

	log.Println("Process Ended")
	return
}

// create ReadWrite
func NewReadWriter(name string, maxReads int, maxReaders int) *readerWriter {
	// Create a value of readerWriter and initialize.
	rw := readerWriter{
		name:          name,
		shutdown:      make(chan struct{}),
		maxReads:      maxReads,
		maxReaders:    maxReaders,
		readerControl: make(semaphore, maxReads),
	}

	// init add
	rw.reportShutdown.Add(maxReaders)
	// all reader will reader
	for goroutine := 0; goroutine < maxReaders; goroutine++ {
		go rw.reader(goroutine)
	}
	// shutdown once
	rw.reportShutdown.Add(1)
	go rw.writer()
	// return
	return &rw
}

func shutdown(readerWriters ...*readerWriter) {
	var waitShutdown sync.WaitGroup
	waitShutdown.Add(len(readerWriters))

	for _, readerWriter := range readerWriters {
		// stop own
		go readerWriter.stop(&waitShutdown)
	}

	// wait shutdown
	waitShutdown.Wait()
}

//
func (rw *readerWriter) stop(waitShutdown *sync.WaitGroup) {
	// Schedule the call to Done for once the method returns.
	defer waitShutdown.Done()

	log.Printf("%s\t: #####> Stop", rw.name)

	// close
	close(rw.shutdown)

	// report shutdown
	rw.reportShutdown.Wait()

	log.Printf("%s\t: #####> Stopped", rw.name)
}

// reader
func (rw *readerWriter) reader(reader int) {
	// report shutdown
	defer rw.reportShutdown.Done()

	for {
		select {
		// receive chan shutdown
		case <-rw.shutdown:
			log.Printf("%s\t: #> Reader Shutdown", rw.name)
			return
		// do others
		default:
			rw.performRead(reader)
		}
	}
}

func (rw *readerWriter) performRead(reader int) {
	// lock
	rw.ReadLock(reader)

	// add current readers
	count := atomic.AddInt32(&rw.currentReads, 1)

	// like reading
	log.Printf("%s\t: [%d] Start\t- [%d] Reads\n", rw.name, reader, count)
	// sleep
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// read over, decrement readers
	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: [%d] Finish\t- [%d] Reads\n", rw.name, reader, count)
	// unlock
	rw.ReadUnlock(reader)
}

//
func (rw *readerWriter) writer() {
	// wait wg
	defer rw.reportShutdown.Done()

	for {
		select {
		// write chan shutdown
		case <-rw.shutdown:
			log.Printf("%s\t: #> Writer Shutdown", rw.name)
			return
			// do something
		default:
			rw.performWrite()
		}
	}
}

//
func (rw *readerWriter) performWrite() {
	// n
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("%s\t: *****> Writing Pending\n", rw.name)

	// lock
	rw.WriteLock()

	// Simulate some writing work.
	log.Printf("%s\t: *****> Writing Start", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> Writing Finish", rw.name)

	// unlock
	rw.WriteUnlock()
}

func (rw *readerWriter) ReadLock(reader int) {
	// wait , then do others
	rw.write.Wait()

	// write control
	rw.readerControl.Acquire(1)
}

// unlock
func (rw *readerWriter) ReadUnlock(reader int) {
	// read out
	rw.readerControl.Release(1)
}

func (rw *readerWriter) WriteLock() {
	rw.write.Add(1)

	rw.readerControl.Acquire(rw.maxReads)
}

func (rw *readerWriter) WriteUnlock() {
	rw.readerControl.Release(rw.maxReads)

	rw.write.Done()
}

//
func (s semaphore) Acquire(buffers int) {
	var e struct{}
	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

// read out
func (s semaphore) Release(buffers int) {
	// read buffers
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}
