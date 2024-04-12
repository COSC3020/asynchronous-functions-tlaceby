package main

import (
	"fmt"
	"sync"
)

func countTask(array []int, start, end, key int, countsChannel chan int) {
	count := 0

	for i := start; i <= end; i++ {
		if array[i] == key {
			count++
		}
	}

	countsChannel <- count
}

func countMatches(array []int, key, threadCount int) int {
	if threadCount <= 0 {
		return -1
	}

	wg := sync.WaitGroup{}
	countsCh := make(chan int, threadCount)
	chunkSize := len(array) / threadCount

	for i := 0; i < threadCount; i++ {
		start := i * chunkSize
		end := start + chunkSize - 1

		// Make sure end is added. May not happen when large sizes of threadCount are tested.
		if i == threadCount-1 {
			end = len(array) - 1
		}

		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			countTask(array, start, end, key, countsCh)
		}(start, end) // Invoke the go thread immediatly.
	}

	wg.Wait()
	close(countsCh) // must close a channel before reading from it. Deadlock::
	total := 0
	for count := range countsCh {
		total += count
	}

	return total
}

func main() {
	arr := []int{1, 2, 5, 5, 4, 2, 3, 4, 2, 3, 2, 1, 2}

	count := countMatches(arr, 1, 5)
	fmt.Printf("Number of matches: %d\n", count)
}
