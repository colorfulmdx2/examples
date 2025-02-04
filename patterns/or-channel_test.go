package patterns

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func MakeHTTPCall(url string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)

		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			c <- fmt.Sprintf("Error: %s (%s)", err, url)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		duration := time.Since(start)
		c <- fmt.Sprintf("Response from %s in %v: %s", url, duration, string(body[:100])) // Limiting to first 100 chars
	}()
	return c
}

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-Or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func TestOrWithHttp(t *testing.T) {
	start := time.Now()

	// List of URLs to request
	fastestResponse := <-Or(
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
		MakeHTTPCall("https://catfact.ninja/fact"),
	)

	fmt.Printf("Fastest response:\n%s\n", fastestResponse)
	fmt.Printf("Total execution time: %v\n", time.Since(start))
}
