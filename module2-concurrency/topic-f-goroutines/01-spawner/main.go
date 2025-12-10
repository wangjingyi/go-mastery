// Assignment 1: The Spawner
//
// Goal: Launch 10,000 goroutines that print "Done". Measure execution time.
//
// Instructions:
// 1. Create a loop that spawns 10,000 goroutines
// 2. Each goroutine should do some work (print or compute)
// 3. Measure how long it takes to spawn them all
// 4. Observe: goroutines are CHEAP (2kb stack each)
//
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	const numGoroutines = 10_000

	start := time.Now()

	// Spawn 10,000 goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// Each goroutine does minimal work
			_ = id * id // Some computation
		}(i)
	}

	elapsed := time.Since(start)

	fmt.Printf("Spawned %d goroutines in %v\n", numGoroutines, elapsed)
	fmt.Println("Note: Main might exit before goroutines complete!")
	fmt.Println("We'll fix this with WaitGroup in the next assignment.")

	// Give goroutines a chance to run (not a proper solution!)
	time.Sleep(100 * time.Millisecond)

	// Key insight: Creating 10,000 goroutines is FAST
	// Each goroutine has only a 2KB initial stack
	// Compare this to OS threads (typically 1-2MB stack each)!
}

