package main

import (
	"math/rand/v2"
	"testing"
)

// Coppied directly from stack overflow:
// https://stackoverflow.com/questions/23577091/generating-random-numbers-over-a-range-in-go
func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func assert(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEmptyArray(t *testing.T) {
	arr := []int{}
	got := countMatches(arr, 10, 1)
	assert(got, 0, t)
}

func TestNoThreads(t *testing.T) {
	arr := []int{}
	got := countMatches(arr, 1, 0)
	assert(got, -1, t)
}

func TestRandomArrays(t *testing.T) {
	// Generate 100 tests and run them
	for i := 0; i < 100; i++ {
		threadCount := randRange(1, 10) // betweeen 1 and 10 threads
		expected := randRange(1, 10)    // expected number is between 1 and 10
		size := randRange(25, 100)      // the size will vary from 25 to 100 random numbers
		arr := make([]int, size)
		want := 0 // number of expected which was generated

		// Fill in the new array with random numbers and count occurences of the randomly generated expected number.
		for ind := range arr {
			arr[ind] = randRange(1, 10)

			if arr[ind] == expected {
				want++
			}
		}

		assert(countMatches(arr, expected, threadCount), want, t)
	}
}
