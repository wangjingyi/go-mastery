// Assignment 1: The Swap Function
//
// Goal: Write Swap(a, b *int). Verify variables change in main.
//
// Instructions:
// 1. Implement Swap that takes two int pointers
// 2. Swap the values they point to
// 3. Call Swap from main and verify the swap worked
//
// TODO: Implement your solution below

package main

import "fmt"

// Swap exchanges the values pointed to by a and b
func Swap(a, b *int) {
	// Your code here
}

func main() {
	x, y := 10, 20
	fmt.Printf("Before: x=%d, y=%d\n", x, y)

	Swap(&x, &y)

	fmt.Printf("After:  x=%d, y=%d\n", x, y)
	// Expected: x=20, y=10
}

