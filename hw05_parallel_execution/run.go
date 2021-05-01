package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	// Place your code here
	jobsCh := make(chan Task, len(tasks))
	cancelCh := make(chan struct{})
	errorsCh := make(chan error, M)

	if M >= 0 {

	}

	go func() {

	}()

	go func() {
		for i := range tasks {
			jobsCh <- tasks[i]
		}
		close(jobsCh)
	}()

	for i := 0; i < N; i++ {
		go func(cancel chan struct{}) {
			for {
				select {
				case <-cancel:
					return
				case f := <-jobsCh:
					err := f()
					if err != nil {

					}
				}
			}
		}()
	}

	return nil
}
