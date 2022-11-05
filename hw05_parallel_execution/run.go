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

	numberOfTasks := len(tasks)
	t := make(chan Task, numberOfTasks)
	errChannel := make(chan int)  // канал для записи ошибок
	exitChannel := make(chan int) // канал для выхода из программы если в него что-то записали
	for i := 0; i < n; i++ {      // создаем n горутин
		wg.Add(1)
		go func(t <-chan Task, errChannel chan<- int, exitChannel <-chan int) {
			defer wg.Done() // если горутина выполнилась

			for {
				select {
				case result := <-t: // читаем из канала новую задачу
					err := result() // выполняем
					if err != nil { // если ошибка
						select {
						case errChannel <- 1: //то пишем в канал ошибок
							err = nil
						case <-exitChannel: // читаем канал если есть что-то - значит выходим
							return
						}
					}
				case <-exitChannel: // читаем канал если есть что-то - значит выходим
					return
				}
			}
		}(t, errChannel, exitChannel)
	}

	var errorsCount int
	var count int
	for k := 0; k < numberOfTasks; k++ {
		result := tasks[k]
		select {
		case <-errChannel: //читаем из канала ошибок
			errorsCount++
		case t <- result:
			count++
		}
	}

	close(exitChannel)
	wg.Wait() // ждем пока выполнятся все таски
	close(t)  // а уже потом закрываем канал
	close(errChannel)

	if errorsCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
