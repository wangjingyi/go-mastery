// Assignment 2: The Writer Adapter
//
// Goal: Create a ConsoleWriter struct.
//       Implement Write([]byte) (int, error).
//       Pass it to fmt.Fprintf (which expects io.Writer).
//
// Instructions:
// 1. Create ConsoleWriter struct
// 2. Implement io.Writer interface: Write(p []byte) (n int, err error)
// 3. Pass it to fmt.Fprintf and observe it works!
//
// Key insight: Go uses implicit interfaces - no "implements" keyword needed

package main

import (
	"fmt"
	"strings"
)

// ConsoleWriter writes to stdout with a prefix
type ConsoleWriter struct {
	Prefix string
}

// Write implements io.Writer interface
// This is ALL you need to satisfy io.Writer!
func (cw ConsoleWriter) Write(p []byte) (n int, err error) {
	output := cw.Prefix + string(p)
	fmt.Print(output)
	return len(p), nil
}

// UppercaseWriter transforms all text to uppercase
type UppercaseWriter struct{}

func (uw UppercaseWriter) Write(p []byte) (n int, err error) {
	upper := strings.ToUpper(string(p))
	fmt.Print(upper)
	return len(p), nil
}

func main() {
	// ConsoleWriter satisfies io.Writer implicitly
	console := ConsoleWriter{Prefix: "[LOG] "}
	
	// fmt.Fprintf expects io.Writer - our ConsoleWriter works!
	fmt.Fprintf(console, "Hello, %s!\n", "World")
	fmt.Fprintf(console, "The answer is %d\n", 42)

	fmt.Println()

	// UppercaseWriter also satisfies io.Writer
	upper := UppercaseWriter{}
	fmt.Fprintf(upper, "this will be uppercase\n")

	// The power: any function expecting io.Writer works with your types
	// No need to modify the function or use "implements" keyword
}

