// Assignment 3: Select Timeout
//
// Goal: Wait for a channel or time.After(2 * time.Second).
//       Print which happened first.
//
// Instructions:
// 1. Create a channel that may or may not receive data
// 2. Use select with time.After for timeout
// 3. Handle both cases
//
// Run: go run main.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Fast Response ===")
	waitForResult(500 * time.Millisecond) // Response comes quickly

	fmt.Println("\n=== Slow Response (Timeout) ===")
	waitForResult(3 * time.Second) // Response too slow, will timeout
}

func waitForResult(responseDelay time.Duration) {
	result := make(chan string)

	// Simulate an async operation
	go func() {
		time.Sleep(responseDelay)
		result <- "Data received!"
	}()

	// Wait for result OR timeout
	select {
	case data := <-result:
		fmt.Println("✅ Success:", data)

	case <-time.After(2 * time.Second):
		fmt.Println("⏰ Timeout: No response within 2 seconds")
	}
}

// More realistic example: Multiple operations with timeout
func multipleOperations() {
	api1 := make(chan string)
	api2 := make(chan string)

	// Simulate two API calls
	go func() {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		api1 <- "Response from API 1"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		api2 <- "Response from API 2"
	}()

	timeout := time.After(2 * time.Second)
	results := make([]string, 0, 2)

	// Collect results until timeout
	for i := 0; i < 2; i++ {
		select {
		case r := <-api1:
			results = append(results, r)
			fmt.Println("Got:", r)
		case r := <-api2:
			results = append(results, r)
			fmt.Println("Got:", r)
		case <-timeout:
			fmt.Println("Timeout! Only got", len(results), "responses")
			return
		}
	}

	fmt.Println("All responses received:", results)
}

