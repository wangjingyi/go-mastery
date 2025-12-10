// Assignment 5: The Nil Panic Guard
//
// Goal: Write a function accepting *User.
//       If passed nil, return an error instead of panicking.
//
// Instructions:
// 1. Implement GetUserName that safely handles nil
// 2. Return an error if user is nil
// 3. Test with both valid and nil users
//
// TODO: Implement your solution below

package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

var ErrNilUser = errors.New("user is nil")

// GetUserName safely returns the user's name
// Returns an error if user is nil instead of panicking
func GetUserName(u *User) (string, error) {
	// TODO: Check for nil and return error
	// Your code here

	return u.Name, nil // This will PANIC if u is nil!
}

func main() {
	// Test with valid user
	validUser := &User{Name: "Bob", Age: 30}
	name, err := GetUserName(validUser)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User name:", name)
	}

	// Test with nil user
	var nilUser *User = nil
	name, err = GetUserName(nilUser)
	if err != nil {
		fmt.Println("Error:", err) // Expected: Error: user is nil
	} else {
		fmt.Println("User name:", name)
	}
}

