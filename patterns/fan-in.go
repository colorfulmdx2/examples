package patterns

import "sync"
import "fmt"

func FanIn_MergeChannels(chs []chan int) chan int {
	out := make(chan int)
	a := 4
	fmt.Println(a)

	go func() {
		// Create a WaitGroup to track all the goroutines
		wg := &sync.WaitGroup{}

		// Iterate over all input channels
		for _, ch := range chs {
			wg.Add(1)
			// Start a goroutine for each channel
			go func(c chan int) {
				defer wg.Done()
				for i := range c {
					out <- i
				}
			}(ch)
		}

		// Close the `out` channel when all goroutines are done
		go func() {
			wg.Wait()
			close(out)
		}()
	}()

	return out
}
