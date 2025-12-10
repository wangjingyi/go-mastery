// Package calc provides basic calculator functions
package calc

import "errors"

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNegativeInput  = errors.New("negative input not allowed")
)

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Factorial returns n! (n factorial)
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

