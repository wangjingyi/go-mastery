# Assignment 2: Test Coverage

## Goal
Run `go test -cover` and achieve 80%+ coverage. Use `go tool cover -html` to visualize.

## Instructions

### Step 1: Create a module with functions to test
```bash
cd module3-testing/topic-j-unit-testing/02-test-coverage
go mod init github.com/yourname/coverage-demo
```

### Step 2: Create calculator.go
```go
package main

// Add returns the sum of a and b
func Add(a, b int) int {
    return a + b
}

// Subtract returns a minus b
func Subtract(a, b int) int {
    return a - b
}

// Multiply returns a times b
func Multiply(a, b int) int {
    return a * b
}

// Divide returns a divided by b
// Returns 0 if b is zero
func Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}

// IsEven returns true if n is even
func IsEven(n int) bool {
    return n%2 == 0
}

// Max returns the larger of a or b
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Factorial returns n!
func Factorial(n int) int {
    if n <= 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {}
```

### Step 3: Create calculator_test.go (incomplete tests)
```go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2,3) = %d; want 5", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5,3) = %d; want 2", result)
    }
}

// TODO: Add more tests to increase coverage!
```

### Step 4: Check initial coverage
```bash
go test -cover
```
You should see something like: `coverage: 28.6% of statements`

### Step 5: Generate coverage profile
```bash
go test -coverprofile=coverage.out
```

### Step 6: View HTML report
```bash
go tool cover -html=coverage.out
```
This opens a browser showing:
- 🟢 Green = covered lines
- 🔴 Red = uncovered lines

### Step 7: Add tests to reach 80%+ coverage

Add tests for:
- `Multiply`
- `Divide` (including the zero case!)
- `IsEven` (both true and false cases)
- `Max` (both branches)
- `Factorial`

### Step 8: Verify 80%+ coverage
```bash
go test -cover
# Should show: coverage: 80%+ of statements
```

## Key Commands
```bash
# Basic coverage
go test -cover

# Coverage profile
go test -coverprofile=coverage.out

# HTML visualization
go tool cover -html=coverage.out

# Coverage by function
go tool cover -func=coverage.out

# Coverage in specific package
go test -cover ./...
```

## Key Learnings
- Coverage measures which lines are executed by tests
- 80% is a common minimum target
- 100% coverage doesn't mean bug-free code
- Focus on testing critical paths and edge cases

