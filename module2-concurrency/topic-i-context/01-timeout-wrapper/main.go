// Assignment 1: Timeout Wrapper
//
// Goal: Function sleeps 5s. Context timeout 2s.
//       Return error immediately on timeout.
//
// Instructions:
// 1. Create a context with timeout
// 2. Use select to wait for work OR context cancellation
// 3. Return immediately when context is cancelled
//
// Run: go run main.go

package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation simulates a slow operation
func slowOperation(ctx context.Context) error {
	// Create a channel to signal completion
	done := make(chan struct{})

	go func() {
		// Simulate slow work (5 seconds)
		time.Sleep(5 * time.Second)
		close(done)
	}()

	// Wait for either completion or cancellation
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err() // Returns context.DeadlineExceeded or context.Canceled
	}
}

func main() {
	fmt.Println("=== Timeout Example (2s timeout, 5s work) ===")

	// Create context with 2 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to release resources!

	start := time.Now()
	err := slowOperation(ctx)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error after %v: %v\n", elapsed, err)
	} else {
		fmt.Printf("✅ Completed in %v\n", elapsed)
	}

	fmt.Println("\n=== No Timeout Example (10s timeout, 1s work) ===")

	// Now with longer timeout
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	start = time.Now()
	err = fastOperation(ctx2)
	elapsed = time.Since(start)

	if err != nil {
		fmt.Printf("❌ Error after %v: %v\n", elapsed, err)
	} else {
		fmt.Printf("✅ Completed in %v\n", elapsed)
	}
}

func fastOperation(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second) // Only 1 second
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

