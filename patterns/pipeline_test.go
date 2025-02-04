package patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intStream := Generator(done, 1, 2, 3, 4)

	pipeline := Multiply(done, Add(done, Multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
}
