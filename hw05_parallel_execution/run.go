package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}

	var wg sync.WaitGroup

	t := make(chan Task)
	errChannel := make(chan error) // канал для записи ошибок
	var errorsCount int
	done := make(chan struct{})

	go func() {
		for _, err := range tasks {
			select {
			case t <- err:
			case <-done:
				break
			}
		}
		close(t)
		wg.Wait() // ждем пока выполнятся все таски
		close(errChannel)
	}()

	for i := 0; i < n; i++ { // создаем n горутин
		wg.Add(1)
		go func() {
			defer wg.Done() // если горутина выполнилась

			for result := range t { // в цикле по каналу читаем новую задачу
				select {
				case errChannel <- result():
				case <-done:
					return
				}
			}
		}()
	}

	for errResult := range errChannel { // читаем из канала ошибок
		if errResult != nil { // если ошибка
			errorsCount++ // то прибавляем
		}
		if errorsCount == m && m > 0 {
			close(done)
			wg.Wait()
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
