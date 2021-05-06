package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, N int, M int) error {
	// Place your code here
	doneCh := make(chan struct{})
	tasksCh := make(chan Task)
	errorsCh := make(chan struct{})
	completeCh := make(chan struct{})

	go func(done chan struct{}) {
		defer close(tasksCh)
		for i := range tasks {
			select {
			case <-done:
				return
			default:
			}

			select {
			case <-done:
				return
			default:
				tasksCh <- tasks[i]
			}
		}
	}(doneCh)

	for i := 0; i < N; i++ {
		go func(done <-chan struct{}, errCh chan<- struct{}, complete chan<- struct{}) {
			for {
				select {
				case <-done:
					return
				default:
				}

				select {
				case <-done:
					return
				case task := <-tasksCh:
					if task != nil {
						err := task()
						if err != nil {
							errCh <- struct{}{}
						} else {
							complete <- struct{}{}
						}
					}
				}
			}
		}(doneCh, errorsCh, completeCh)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		complete := 0
		errs := 0
		for {
			select {
			case <-errorsCh:
				if M <= 0 {
					return
				}
				errs++
				if errs == M {
					return
				}
			case <-completeCh:
				complete++
				if complete == len(tasks) {
					return
				}
			}
		}
	}(wg)

	wg.Wait()
	close(doneCh)
	return nil
}
