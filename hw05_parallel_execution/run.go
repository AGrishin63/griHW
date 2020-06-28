package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {

	w := &sync.WaitGroup{}
	m := &sync.Mutex{}
	inChan := make(chan Task)
	stopChan := make(chan bool)
	errCount := 0
	defer close(inChan)
	defer close(stopChan)

	//Task processor
	for i := 0; i < N; i++ {
		go func() {
			defer w.Done()
			w.Add(1)
			for {
				select {
				case t, ok := <-inChan:
					if !ok {
						fmt.Println("Error on getting task.")
					} else {
						err := t()
						if err != nil {
							m.Lock()
							errCount++
							m.Unlock()
							fmt.Println(err)
						} else {
							fmt.Println("No error task ", i)
						}
					}
				case <-stopChan:
					return
				}
			}
			return
		}()
	}

	//Sending tasks to process
	for n, t := range tasks {
		if errCount < M {
			select {
			case inChan <- t:
			}
		} else {
			break
		}
	}
	fmt.Println("Все задачи выполнены.")
	for i := 0; i < N; i++ {
		stopChan <- true
	}
	w.Wait()
	fmt.Println("Конец задачи")
	if errCount < M {
		return nil
	} else {
		return ErrErrorsLimitExceeded
	}

}
