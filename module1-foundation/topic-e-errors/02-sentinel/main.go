// Assignment 2: The Sentinel
//
// Goal: Define var ErrNotFound = errors.New(...).
//       Return it. Check for it using errors.Is.
//
// Instructions:
// 1. Define sentinel errors (package-level error values)
// 2. Return them from functions
// 3. Use errors.Is to check for specific errors
//
// Key insight: Sentinel errors are compared by identity, not value

package main

import (
	"errors"
	"fmt"
)

// Sentinel errors - defined at package level
var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized access")
	ErrInvalidInput = errors.New("invalid input provided")
)

// User represents a user in our system
type User struct {
	ID   int
	Name string
}

// Database simulation
var users = map[int]User{
	1: {ID: 1, Name: "Alice"},
	2: {ID: 2, Name: "Bob"},
}

// GetUser retrieves a user by ID
func GetUser(id int) (User, error) {
	if id <= 0 {
		return User{}, ErrInvalidInput
	}
	
	user, exists := users[id]
	if !exists {
		return User{}, ErrNotFound
	}
	
	return user, nil
}

// DeleteUser removes a user (requires auth)
func DeleteUser(id int, isAdmin bool) error {
	if !isAdmin {
		return ErrUnauthorized
	}
	
	if _, exists := users[id]; !exists {
		return ErrNotFound
	}
	
	delete(users, id)
	return nil
}

func main() {
	// Test GetUser with valid ID
	user, err := GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found user: %+v\n", user)
	}

	// Test with non-existent user - use errors.Is
	_, err = GetUser(999)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User 999: Not found (caught with errors.Is)")
	}

	// Test with invalid input
	_, err = GetUser(-1)
	if errors.Is(err, ErrInvalidInput) {
		fmt.Println("User -1: Invalid input (caught with errors.Is)")
	}

	// Test unauthorized deletion
	err = DeleteUser(1, false)
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("Delete: Unauthorized (caught with errors.Is)")
	}

	// Switch on error type
	_, err = GetUser(0)
	switch {
	case errors.Is(err, ErrNotFound):
		fmt.Println("Handle not found...")
	case errors.Is(err, ErrUnauthorized):
		fmt.Println("Handle unauthorized...")
	case errors.Is(err, ErrInvalidInput):
		fmt.Println("Handle invalid input: ", err)
	case err != nil:
		fmt.Println("Unknown error:", err)
	}
}

