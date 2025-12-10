// Assignment 5: The Safe Recovery
//
// Goal: Write a web handler that panics.
//       Use defer and recover to catch the panic and log it instead of crashing.
//
// Instructions:
// 1. Create a handler that may panic
// 2. Create recovery middleware using defer/recover
// 3. Catch the panic, log it, and return a proper error response
//
// Key insight: recover() only works inside a deferred function

package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Simulated handler that panics on certain input
func riskyHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	
	if userID == "panic" {
		// This would crash the entire server without recovery!
		panic("something went terribly wrong!")
	}
	
	if userID == "nil" {
		// Nil pointer dereference panic
		var user *struct{ Name string }
		fmt.Fprintf(w, "User: %s", user.Name) // PANIC!
	}
	
	fmt.Fprintf(w, "Hello, user %s!", userID)
}

// RecoveryMiddleware wraps a handler and recovers from panics
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// defer + recover pattern
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace
				log.Printf("PANIC RECOVERED: %v\n", err)
				log.Printf("Stack trace:\n%s", debug.Stack())
				
				// Return a proper error response instead of crashing
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Internal Server Error")
			}
		}()
		
		// Call the actual handler
		next(w, r)
	}
}

// Demonstrate basic defer/recover without HTTP
func demonstrateRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	
	fmt.Println("About to panic...")
	panic("test panic!")
	
	// This line is never reached
	fmt.Println("After panic") //nolint:govet
}

func main() {
	fmt.Println("=== Basic Recovery Demo ===")
	demonstrateRecovery()
	fmt.Println("Program continues after recovery!")
	
	fmt.Println("\n=== HTTP Server with Recovery Middleware ===")
	fmt.Println("Starting server on :8080")
	fmt.Println("Try:")
	fmt.Println("  curl http://localhost:8080/?id=123     # Normal")
	fmt.Println("  curl http://localhost:8080/?id=panic   # Triggers panic")
	fmt.Println("  curl http://localhost:8080/?id=nil     # Nil pointer panic")
	
	// Wrap the risky handler with recovery middleware
	http.HandleFunc("/", RecoveryMiddleware(riskyHandler))
	
	// Note: In real apps, you'd use proper middleware chaining
	log.Fatal(http.ListenAndServe(":8080", nil))
}

