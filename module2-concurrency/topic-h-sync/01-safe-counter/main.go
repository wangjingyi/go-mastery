// Assignment 1: Safe Counter
//
// Goal: Fix the race condition (from F3) using sync.Mutex.
//
// Instructions:
// 1. Create a Counter struct with a mutex
// 2. Lock/Unlock around the increment operation
// 3. Verify with go run -race
//
// Run: go run -race main.go

package main

import (
	"fmt"
	"sync"
)

// Counter is a thread-safe counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()         // Acquire lock
	defer c.mu.Unlock() // Release lock when function returns
	c.value++
}

// Value safely reads the counter
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	const numGoroutines = 1000

	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment() // Thread-safe now!
		}()
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter.Value())
	fmt.Printf("Expected value: %d\n", numGoroutines)

	if counter.Value() == numGoroutines {
		fmt.Println("✅ Race condition fixed!")
	} else {
		fmt.Println("❌ Still buggy")
	}
}

// Key rules for Mutex:
// 1. Always pair Lock() with Unlock()
// 2. Use defer for Unlock() to prevent deadlocks on early return/panic
// 3. Don't copy a Mutex (pass by pointer)
// 4. Keep the critical section as small as possible

