// Assignment 4: The Mutation Bug
//
// Goal: Create a method func (u User) Birthday() (value receiver).
//       Call it. Why didn't age change? Fix it with a pointer receiver.
//
// Instructions:
// 1. Run the code below - observe the bug
// 2. Understand why value receivers don't mutate the original
// 3. Fix Birthday() to use a pointer receiver
//
// TODO: Fix the bug below

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Birthday increments the user's age
// BUG: This uses a value receiver - it operates on a COPY!
func (u User) Birthday() {
	u.Age++ // This increments the COPY, not the original
}

// TODO: Create BirthdayFixed with pointer receiver
// func (u *User) BirthdayFixed() { ... }

func main() {
	user := User{Name: "Alice", Age: 25}

	fmt.Printf("Before birthday: %s is %d\n", user.Name, user.Age)

	user.Birthday() // This doesn't work!

	fmt.Printf("After birthday:  %s is %d\n", user.Name, user.Age)
	// Expected: 26, Actual: 25 (BUG!)

	// TODO: Call BirthdayFixed and verify it works
}

