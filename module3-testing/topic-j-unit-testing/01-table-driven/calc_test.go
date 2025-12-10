// Assignment 1: Table-Driven Tests
//
// Goal: Write tests using table-driven approach with subtests.
//
// Run: go test -v ./...

package calc

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		expected  int
		expectErr error
	}{
		{"normal division", 10, 2, 5, nil},
		{"integer division", 7, 2, 3, nil}, // 7/2 = 3 (truncated)
		{"divide by zero", 10, 0, 0, ErrDivisionByZero},
		{"divide zero", 0, 5, 0, nil},
		{"negative division", -10, 2, -5, nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Divide(tc.a, tc.b)

			// Check error
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("Divide(%d, %d) error = %v; want %v", tc.a, tc.b, err, tc.expectErr)
			}

			// Check result (only if no error expected)
			if tc.expectErr == nil && result != tc.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		expected  int
		expectErr error
	}{
		{"zero", 0, 1, nil},
		{"one", 1, 1, nil},
		{"five", 5, 120, nil},
		{"ten", 10, 3628800, nil},
		{"negative", -1, 0, ErrNegativeInput},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Factorial(tc.n)

			if !errors.Is(err, tc.expectErr) {
				t.Errorf("Factorial(%d) error = %v; want %v", tc.n, err, tc.expectErr)
			}

			if tc.expectErr == nil && result != tc.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tc.n, result, tc.expected)
			}
		})
	}
}

// Example of parallel subtests
func TestMultiply_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 3, 4, 12},
		{"with zero", 5, 0, 0},
		{"negatives", -3, -4, 12},
	}

	for _, tc := range tests {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Run subtests in parallel
			result := Multiply(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
