package patterns

import (
	"sync"
	"testing"
)

func TestCheckUrls(t *testing.T) {
	urls := []string{"https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost", "https://www.google.com", "https://badhost"}

	WorkerpoolingCheckurls(urls, &sync.WaitGroup{}, 3)
}
