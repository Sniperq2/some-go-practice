package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	t := make(chan Task)
	errChannel := make(chan struct{}) // канал для записи ошибок
	for i := 0; i < n; i++ {          // создаем n горутин
		wg.Add(1)
		go func(t <-chan Task, errChannel chan<- struct{}) {
			defer wg.Done() // если горутина выполнилась

			for result := range t { // в цикле по каналу читаем новую задачу
				if err := result(); err != nil { // выполняем и если ошибка
					errChannel <- struct{}{} // пишем в канал сигнал об ошибке
				}
			}
		}(t, errChannel)
	}

	var errorsCount int
	var count int
	for _, result := range tasks {
		if errorsCount >= m { // если ошибок больше чем допустимо выходим из цикла
			break
		}
		select {
		case <-errChannel: // читаем из канала ошибок
			errorsCount++
		case t <- result:
			count++
		}
	}

	wg.Wait() // ждем пока выполнятся все таски
	close(errChannel)

	if errorsCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
