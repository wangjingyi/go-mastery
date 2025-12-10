// Assignment 4: The Loop Trap
//
// Goal: Launch goroutines inside a for i := 0 loop printing i.
//       Observe they all print the same number. Fix by passing i as argument.
//
// Instructions:
// 1. First, see the bug: goroutines capture the loop variable
// 2. Then fix it by passing i as a parameter
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=== THE BUG ===")
	buggyVersion()

	fmt.Println("\n=== THE FIX ===")
	fixedVersion()
}

func buggyVersion() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// BUG: All goroutines capture the SAME variable i
			// By the time they run, the loop has finished and i == 5
			fmt.Printf("Buggy: i = %d\n", i)
		}()
	}

	wg.Wait()
	// Output: All print 5 (or close to it)
}

func fixedVersion() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) { // FIX: Pass i as parameter
			defer wg.Done()
			// Each goroutine gets its OWN copy of the value
			fmt.Printf("Fixed: id = %d\n", id)
		}(i) // Pass i here - captures the value at this moment
	}

	wg.Wait()
	// Output: 0, 1, 2, 3, 4 (in some order)
}

// Alternative fix (Go 1.22+):
// As of Go 1.22, loop variables are captured correctly per-iteration
// But it's still good practice to be explicit

func alternativeFix() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		i := i // Shadow the variable - creates new variable per iteration
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Shadow fix: i = %d\n", i)
		}()
	}

	wg.Wait()
}

