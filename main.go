package main

import (
	"fmt"
	"sort"
)

func findFreePortRanges(min, max int, busy []int) [][]int {
	// Sort busy ports to ensure correct processing
	sort.Ints(busy)

	var result [][]int
	start := min

	for _, port := range busy {
		// If the current start is less than the busy port, we have a free range
		if start < port {
			result = append(result, []int{start, port - 1})
		}
		// Update the start to the next possible port
		start = port + 1
	}

	// Add the final range if there are free ports after the last busy port
	if start <= max {
		result = append(result, []int{start, max})
	}

	return result
}

func main() {
	min := 30000
	max := 32000
	busy := []int{30100, 30200}

	freeRanges := findFreePortRanges(min, max, busy)
	fmt.Println("Free Port Ranges:", freeRanges)
}
