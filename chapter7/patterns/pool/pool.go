package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// pool
type Pool struct {
	// lock
	m sync.Mutex
	// closer
	resources chan io.Closer
	// get conn
	factory func() (io.Closer, error)
	// flag
	closed bool
}

// error
var ErrPoolClosed = errors.New("Pool has been closed.")

// constructor
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	// error create
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}
	// create
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// getConn
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			// already closed
			return nil, ErrPoolClosed
		}
		return r, nil

	// factory create
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// close
func (p *Pool) Release(r io.Closer) {
	// lock
	p.m.Lock()
	defer p.m.Unlock()

	// alive with pool and death together
	if p.closed {
		_ = r.Close()
		return
	}

	select {
	//
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	//
	default:
		log.Println("Release:", "Closing")
		_ = r.Close()
	}
}

func (p *Pool) Close() {
	// lock
	p.m.Lock()
	defer p.m.Unlock()

	// closed
	if p.closed {
		return
	}

	// set flag
	p.closed = true

	// close channel
	close(p.resources)

	// close resource
	for r := range p.resources {
		_ = r.Close()
	}
}
