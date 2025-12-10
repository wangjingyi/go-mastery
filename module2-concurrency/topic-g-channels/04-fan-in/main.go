// Assignment 4: Fan-In
//
// Goal: Merge 2 channels into 1.
//
// Instructions:
// 1. Create two input channels
// 2. Create a merged output channel
// 3. Use goroutines to forward from inputs to output
// 4. Close output only when BOTH inputs are closed
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// fanIn merges multiple channels into one
func fanIn(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	merged := make(chan string)

	// Start a goroutine for each input channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	// Close merged channel when all inputs are done
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// producer creates a channel that sends n messages
func producer(name string, n int, delay time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			time.Sleep(delay)
			ch <- fmt.Sprintf("%s: message %d", name, i)
		}
	}()
	return ch
}

func main() {
	// Create two producers with different speeds
	ch1 := producer("Fast", 5, 100*time.Millisecond)
	ch2 := producer("Slow", 3, 300*time.Millisecond)

	// Merge them into one channel
	merged := fanIn(ch1, ch2)

	// Read from merged channel
	fmt.Println("Receiving merged messages:")
	for msg := range merged {
		fmt.Println("  <-", msg)
	}
	fmt.Println("All channels closed, done!")
}

// Alternative: Simple two-channel fan-in using select
func simpleFanIn(ch1, ch2 <-chan string) <-chan string {
	merged := make(chan string)

	go func() {
		defer close(merged)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil // Disable this case
					continue
				}
				merged <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil // Disable this case
					continue
				}
				merged <- v
			}
		}
	}()

	return merged
}

