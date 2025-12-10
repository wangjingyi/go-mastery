// Assignment 2: The WaitGroup
//
// Goal: Sync the 10,000 goroutines using sync.WaitGroup so main doesn't exit early.
//
// Instructions:
// 1. Create a sync.WaitGroup
// 2. Add 1 before each goroutine, Done() inside each
// 3. Wait() at the end to ensure all complete
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numGoroutines = 10_000

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment counter BEFORE spawning

		go func(id int) {
			defer wg.Done() // Decrement counter when done

			// Simulate some work
			_ = id * id
		}(i)
	}

	// Wait for ALL goroutines to complete
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("All %d goroutines completed in %v\n", numGoroutines, elapsed)

	// WaitGroup rules:
	// 1. Add() before spawning the goroutine
	// 2. Done() inside the goroutine (use defer!)
	// 3. Wait() blocks until counter reaches 0
}

// Common mistakes to avoid:

func badExample1() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(id int) {
			wg.Add(1) // BAD: Add inside goroutine - race condition!
			defer wg.Done()
			fmt.Println(id)
		}(i)
	}
	wg.Wait() // Might return before all goroutines even start
}

func badExample2() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(id)
			// BAD: Forgot wg.Done() - Wait() blocks forever!
		}(i)
	}
	// wg.Wait() // Would deadlock!
}

