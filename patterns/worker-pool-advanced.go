package patterns

/*

Клиент передаёт на вход некоторый объект с описанием задачи.

Мы должны, получив эту задачу, поставить ее в очередь на обработку.
Задача запускается в обработку (имитируем полезную работу через time.Sleep(5*time.Second)),
    если у нас есть свободные обработчики.
Как только очередная задача выполнилась – берём следующую задачу из очереди.
Если в очереди пусто, ожидаем новых задач от клиентов.

Сервис одновременно может обрабатывать не более N задач. Остальные задачи должны помещаться в очередь,
    очередь ограничена максимальным размером maxSize. Когда очередь заполнена, сервис должен вернуть ошибку

*/

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type WorkerPool struct {
	queue       chan Task
	maxWorkers  int
	maxQueue    int
	isFinishing bool
	wg          sync.WaitGroup
	once        sync.Once
	mu          sync.Mutex
}

func NewWorkerPool(maxWorkers, maxQueue int) *WorkerPool {
	return &WorkerPool{
		queue:      make(chan Task, maxQueue),
		maxWorkers: maxWorkers,
		maxQueue:   maxQueue,
	}
}

func (wp *WorkerPool) AddTask(task Task) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.isFinishing {
		return errors.New("wp is finishing")
	}

	select {
	case wp.queue <- task:
		return nil
	default:
		return errors.New("queue is full")
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go wp.Worker(i)
	}
}

func (wp *WorkerPool) Worker(id int) {
	defer wp.wg.Done()

	for task := range wp.queue {
		fmt.Println(task.ID, " is in work")
		time.Sleep(time.Second)
	}
}

func (wp *WorkerPool) Wait() {
	wp.once.Do(func() {
		close(wp.queue)

		wp.mu.Lock()
		wp.isFinishing = true
		wp.mu.Unlock()
		wp.wg.Wait()
	})
}

//func main() {
//
//	maxWorkers := 3
//	maxQueue := 100
//
//	pool := patterns.NewWorkerPool(maxWorkers, maxQueue)
//	pool.Run()
//	for i := 0; i <= 9; i++ {
//		err := pool.AddTask(patterns.Task{ID: i})
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	pool.Wait()
//	fmt.Println("Done")
//
//}
