// Assignment 3: The Escape Analyst
//
// Goal: Write a function returning a pointer to a local variable.
//       Run `go build -gcflags="-m"` to verify it "escapes to heap".
//
// Instructions:
// 1. Run: go build -gcflags="-m" main.go
// 2. Observe which variables escape to heap
// 3. Understand WHY they escape
//
// Key insight: Local variables that "outlive" the function must be heap-allocated

package main

import "fmt"

// createNumber returns a pointer to a local variable
// This FORCES heap allocation because the variable must outlive the function
func createNumber() *int {
	num := 42 // This will escape to heap!
	return &num
}

// createNumberNoEscape keeps the variable on stack
// Because the variable doesn't escape the function scope
func createNumberNoEscape() int {
	num := 42 // This stays on stack
	return num
}

// createSlice also causes escape
func createSlice() []int {
	data := make([]int, 100) // Escapes to heap
	return data
}

func main() {
	ptr := createNumber()
	fmt.Println("Value from heap:", *ptr)

	val := createNumberNoEscape()
	fmt.Println("Value from stack:", val)

	slice := createSlice()
	fmt.Println("Slice length:", len(slice))
}

// Run: go build -gcflags="-m" main.go
// Expected output will show:
// - "num escapes to heap" for createNumber
// - "make([]int, 100) escapes to heap" for createSlice

