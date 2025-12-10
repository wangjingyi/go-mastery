// Assignment 4: The Wrapper
//
// Goal: Call a failing function. Return fmt.Errorf("db failed: %w", err).
//       Print the full chain.
//
// Instructions:
// 1. Create a chain of function calls, each wrapping errors with context
// 2. Use %w verb to wrap errors (preserves the original)
// 3. Use errors.Is to check for wrapped errors
// 4. Use errors.Unwrap to traverse the chain
//
// Key insight: Wrapping adds context while preserving the original error

package main

import (
	"errors"
	"fmt"
)

// Sentinel errors
var (
	ErrConnection = errors.New("connection failed")
	ErrTimeout    = errors.New("operation timed out")
)

// Layer 1: Database layer
func dbQuery(query string) error {
	// Simulate a connection failure
	return ErrConnection
}

// Layer 2: Repository layer - wraps DB errors
func getUserByID(id int) error {
	err := dbQuery("SELECT * FROM users WHERE id = ?")
	if err != nil {
		// Wrap with context using %w
		return fmt.Errorf("repository.getUserByID(%d): %w", id, err)
	}
	return nil
}

// Layer 3: Service layer - wraps repository errors
func fetchUserProfile(userID int) error {
	err := getUserByID(userID)
	if err != nil {
		// Add more context
		return fmt.Errorf("service.fetchUserProfile: %w", err)
	}
	return nil
}

// Layer 4: Handler layer - wraps service errors
func handleGetProfile(userID int) error {
	err := fetchUserProfile(userID)
	if err != nil {
		return fmt.Errorf("handler.handleGetProfile: %w", err)
	}
	return nil
}

func main() {
	err := handleGetProfile(42)
	
	if err != nil {
		// Print the full error chain
		fmt.Println("Full error message:")
		fmt.Println(err)
		fmt.Println()

		// errors.Is can still find the original error through the wrapping!
		if errors.Is(err, ErrConnection) {
			fmt.Println("✓ Found ErrConnection in the chain")
		}
		
		if errors.Is(err, ErrTimeout) {
			fmt.Println("✓ Found ErrTimeout in the chain")
		} else {
			fmt.Println("✗ ErrTimeout not in chain (expected)")
		}

		// Manually traverse the chain
		fmt.Println("\nError chain:")
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("  -> %v\n", e)
		}
	}
}

// Output:
// Full error message:
// handler.handleGetProfile: service.fetchUserProfile: repository.getUserByID(42): connection failed
//
// ✓ Found ErrConnection in the chain
// ✗ ErrTimeout not in chain (expected)
//
// Error chain:
//   -> handler.handleGetProfile: service.fetchUserProfile: repository.getUserByID(42): connection failed
//   -> service.fetchUserProfile: repository.getUserByID(42): connection failed
//   -> repository.getUserByID(42): connection failed
//   -> connection failed

