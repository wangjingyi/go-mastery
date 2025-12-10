// Assignment 1: The Divider
//
// Goal: Return custom error if dividing by zero. Handle it.
//
// Instructions:
// 1. Create a Divide function that returns (float64, error)
// 2. If divisor is zero, return an error
// 3. Handle the error in main
//
// TODO: Implement your solution below

package main

import (
	"errors"
	"fmt"
)

// ErrDivideByZero is our custom error for division by zero
var ErrDivideByZero = errors.New("cannot divide by zero")

// Divide performs division and returns an error for zero divisor
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	// Successful division
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Division by zero - should return error
	result, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// Using errors.Is to check for specific error
	_, err = Divide(5, 0)
	if errors.Is(err, ErrDivideByZero) {
		fmt.Println("Caught divide by zero specifically!")
	}
}

