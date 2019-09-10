package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// interrupt
	interrupt chan os.Signal
	// error
	complete chan error
	// timeout
	timeout <-chan time.Time
	// tasks
	tasks []func(int)
}

var (
	ErrTimeout   = errors.New("received timeout")
	ErrInterrupt = errors.New("received interrupt")
)

func New(d time.Duration) *Runner {
	// just set attributes
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// add tasks
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// connect two signal
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		// listen goroutine
		r.complete <- r.run()
	}()
	// just listen once
	select {
	// complete
	case err := <-r.complete:
		return err
	// timeout
	case <-r.timeout:
		return ErrTimeout
	}
}

// run
func (r *Runner) run() error {
	// listen
	for id, task := range r.tasks {
		// error , and return
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// run
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	// check
	select {
	// interrupt
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}
