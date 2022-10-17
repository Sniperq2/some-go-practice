package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	numberOfTasks := len(tasks)
	t := make(chan Task, numberOfTasks)

	for i := 0; i < n; i++ { // создаем n горутин
		wg.Add(1)
		go func(t <-chan Task, wg *sync.WaitGroup, m int) {
			defer wg.Done()
			fmt.Println(<-t) // читаем из канала и "выполняем" таску
			// как-то тут надо ошибки считать и передавать в m
		}(t, &wg, m)
	}

	for k := 0; k < numberOfTasks; k++ {
		t <- tasks[k]
	}
	close(t)  // закрываем канал
	wg.Wait() // ждем пока выполнятся все таски
	return nil
}
