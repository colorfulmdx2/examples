package patterns

import (
	"fmt"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	// Step 1: Create channels with test data
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	// Add test data to the channels
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	close(ch2)

	// Step 2: Merge the channels
	chs := []chan int{ch1, ch2}
	out := FanIn_MergeChannels(chs)

	// Step 3: Collect results from the merged channel
	var result []int
	for v := range out {
		result = append(result, v)
	}
	fmt.Println(result)
}
