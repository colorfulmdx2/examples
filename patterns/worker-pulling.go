package patterns

import (
	"fmt"
	"sync"
)

func WorkerPooling_CheckUrls(urls []string, wg *sync.WaitGroup, workers int) {
	ch := make(chan string, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()

			for url := range ch {
				fmt.Printf("Worker %d processing URL: %s\n", workerId, url)
			}
		}(i)
	}

	for _, url := range urls {
		ch <- url
	}

	close(ch)
	wg.Wait()
}
