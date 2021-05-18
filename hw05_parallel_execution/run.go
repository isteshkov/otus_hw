package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, N int, M int) error {
	// Place your code here
	var complete, errs int
	var err error
	doneCh := make(chan struct{})
	tasksCh := make(chan Task)
	resultCh := make(chan error)

	go func(done <-chan struct{}) {
		defer func() {
			close(tasksCh)
		}()
		for i := range tasks {
			select {
			case <-done:
				return
			case tasksCh <- tasks[i]:
			}
		}
	}(doneCh)

	for i := 0; i < N; i++ {
		go func() {
			for {
				select {
				case <-doneCh:
					return
				case task := <-tasksCh:
					if task != nil {
						select {
						case <-doneCh:
							return
						case resultCh <- task():
						}
					}
				}
			}
		}()
	}

	for res := range resultCh {
		if res != nil {
			if M <= 0 {
				err = ErrErrorsLimitExceeded
				break
			}
			errs++
			if errs == M {
				err = ErrErrorsLimitExceeded
				break
			}
		} else {
			complete++
			if complete == len(tasks) {
				break
			}
		}
	}
	close(doneCh)
	return err
}
