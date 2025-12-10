// Assignment 3: The Race Condition
//
// Goal: Increment a global counter from 1000 goroutines.
//       Run with go run -race to detect the race.
//
// Instructions:
// 1. Create a global counter
// 2. Spawn 1000 goroutines, each incrementing the counter
// 3. Run: go run -race main.go
// 4. Observe the race condition detected by the race detector
//
// IMPORTANT: Run with: go run -race main.go

package main

import (
	"fmt"
	"sync"
)

// Global counter - UNSYNCHRONIZED (intentionally buggy)
var counter int

func main() {
	const numGoroutines = 1000

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// RACE CONDITION: Multiple goroutines read/write counter simultaneously
			// This is NOT atomic - it involves: read, increment, write
			counter++ // BUG: Data race!
		}()
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
	fmt.Printf("Expected value: %d\n", numGoroutines)
	fmt.Println()
	fmt.Println("If these don't match, you've witnessed a race condition!")
	fmt.Println("Run with: go run -race main.go to detect it.")
}

// What happens during a race:
//
// Goroutine 1              Goroutine 2
// -----------              -----------
// Read counter (0)
//                          Read counter (0)
// Increment (1)
//                          Increment (1)
// Write counter (1)
//                          Write counter (1)
//
// Result: counter = 1 (should be 2!)
//
// We'll fix this in Topic H using sync.Mutex or sync/atomic

