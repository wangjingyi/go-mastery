// Assignment 5: The Heartbeat
//
// Goal: Background goroutine printing "Pulse" every 500ms.
//       Stop it when main exits.
//
// Instructions:
// 1. Create a goroutine with a ticker
// 2. Use a done channel to signal shutdown
// 3. Clean shutdown when main exits
//
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to signal the heartbeat to stop
	done := make(chan struct{})

	// Start the heartbeat goroutine
	go heartbeat(done, 500*time.Millisecond)

	// Let it run for 3 seconds
	fmt.Println("Main: Starting heartbeat for 3 seconds...")
	time.Sleep(3 * time.Second)

	// Signal the heartbeat to stop
	fmt.Println("Main: Sending shutdown signal...")
	close(done)

	// Give it time to clean up
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main: Exiting")
}

func heartbeat(done <-chan struct{}, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop() // Important: stop the ticker to free resources

	pulseCount := 0

	for {
		select {
		case <-ticker.C:
			// Ticker fired
			pulseCount++
			fmt.Printf("💓 Pulse #%d\n", pulseCount)

		case <-done:
			// Shutdown signal received
			fmt.Println("Heartbeat: Received shutdown signal, stopping...")
			return
		}
	}
}

// Key patterns demonstrated:
// 1. time.Ticker for periodic work
// 2. done channel (chan struct{}) for signaling
// 3. select for handling multiple channels
// 4. Closing a channel to broadcast (all receivers wake up)
// 5. defer ticker.Stop() to prevent resource leak

