// Assignment 5: The Linter Fix
//
// Goal: Write code that intentionally breaks linter rules.
//       Fix all errors until golangci-lint passes.
//
// Instructions:
// 1. Run: golangci-lint run
// 2. Observe the errors below
// 3. Fix each error one by one
// 4. Re-run linter until clean
//
// The code below has INTENTIONAL errors. Fix them!

package main

import (
	"fmt"
	"os"
)

func main() {
	// Error 1: unused variable
	unusedVar := "I'm never used"

	// Error 2: error not checked
	os.Open("nonexistent.txt")

	// Error 3: ineffective assignment
	x := 5
	x = 10
	x = 15

	// Error 4: shadowing
	err := doSomething()
	if true {
		err := doSomethingElse()
		fmt.Println(err)
	}
	fmt.Println(err)

	fmt.Println(x)
	_ = unusedVar
}

func doSomething() error {
	return nil
}

func doSomethingElse() error {
	return nil
}

