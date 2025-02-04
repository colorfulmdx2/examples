package patterns

import (
	"fmt"
	"testing"
)

func TestNewWorkerPool(t *testing.T) {
	maxWorkers := 3
	maxQueue := 100

	pool := NewWorkerPool(maxWorkers, maxQueue)
	pool.Run()
	for i := 0; i <= 9; i++ {
		err := pool.AddTask(Task{ID: i})
		if err != nil {
			fmt.Println(err)
		}
	}

	pool.Wait()
	fmt.Println("Done")

}
