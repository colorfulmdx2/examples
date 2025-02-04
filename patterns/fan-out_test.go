package patterns

import (
	"fmt"
	"testing"
)

func TestCheckUrlsErrorHandling(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost"}

	for result := range CheckUrlsErrorHandling(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
