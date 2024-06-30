package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	taskChan := make(chan Task)
	wg := sync.WaitGroup{}
	var errCount int32

	// Здесь создаем n горутин, которые ждут когда в taskChan начнут поступать данные
	for i := 0; i < n; i++ {
		wg.Add(1)
		// Эта горутина перестанет существовать если количество ошибок больше m и если канал taskChan закроется
		go func() {
			defer wg.Done()
			// range loop. Читаем, пока канал не закрыт
			for task := range taskChan {
				if atomic.LoadInt32(&errCount) >= int32(m) {
					return
				}
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()

	if atomic.LoadInt32(&errCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
