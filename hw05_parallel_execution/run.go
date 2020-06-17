package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error
type wStruct struct {
	w        *sync.WaitGroup
	m        *sync.Mutex
	inChan   chan Task
	stopChan chan bool
	errChan  chan bool
}

func worker(wS *wStruct) {
	defer wS.w.Done()
	wS.w.Add(1)
	for {
		wS.m.Lock()

		select {
		case t, err1 := <-wS.inChan:
			wS.m.Unlock()
			if !err1 {
				fmt.Println("Error on getting task!")
			} else {
				err2 := t()
				if err2 != nil {
					wS.errChan <- true
					fmt.Println(err2)
				}
			}

		case <-wS.stopChan:
			return
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	// Place your code here

	ws := &wStruct{}
	ws.w = &sync.WaitGroup{}
	ws.m = &sync.Mutex{}
	ws.inChan = make(chan Task)
	ws.stopChan = make(chan bool)
	ws.errChan = make(chan bool)
	errCount := 1
	defer close(ws.inChan)
	defer close(ws.stopChan)
	defer close(ws.errChan)

	for i := 0; i < N; i++ {
		go worker(ws)
	}

	for n, t := range tasks {
		fmt.Println(n, t())
		select {
		case ws.inChan <- t:
		case <-ws.errChan:
			errCount++
			fmt.Println(errCount)
			if errCount == M {
				close(ws.stopChan)
				// for i := 0; i < N; i++ {
				// 	fmt.Println("i=", i)
				// 	<-ws.stopChan
				// }
				ws.w.Wait()
				return ErrErrorsLimitExceeded
			}
		}
	}
	close(ws.stopChan)
	ws.w.Wait()
	return nil
}
