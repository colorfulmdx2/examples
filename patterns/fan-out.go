package patterns

import (
	"net/http"
	"sync"
)

type Result struct {
	Error    error
	Response *http.Response
}

func CheckUrlsErrorHandling(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		u := url
		go func() {
			defer wg.Done()

			var result Result
			resp, err := http.Get(u)
			result = Result{Error: err, Response: resp}

			select {
			case <-done:
				return
			case results <- result:
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
