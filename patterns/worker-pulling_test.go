package patterns

import (
	"context"
	"sync"
	"testing"
)

func TestCheckUrls(t *testing.T) {
	// Prepare test data
	urls := []string{"url1", "url2", "url3", "url4", "url5"}
	workers := 2
	ctx, cancel := context.WithTimeout(context.Background, timeout)
	// Create a sync.WaitGroup to track goroutine completion
	var wg sync.WaitGroup

	// Call the WorkerPooling_CheckUrls function
	// This test will just run the function to ensure it executes without errors
	WorkerPooling_CheckUrls(urls, &wg, workers)

	// Wait for all goroutines to complete
	wg.Wait()

	// If the function completes successfully without panics,
	// the test will pass. If an error occurs during execution,
	// the test will fail.
}
