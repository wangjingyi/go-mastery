// Assignment 3: The Constructor Pattern
//
// Goal: Create a private struct `server`. Create a public `NewServer(port int) *server`.
//       Prevent direct initialization.
//
// Instructions:
// 1. Create a lowercase (private) struct named `server`
// 2. Create a public constructor `NewServer` that returns *server
// 3. Add validation in the constructor
// 4. This prevents users from creating invalid servers directly
//
// TODO: Implement your solution below

package main

import (
	"fmt"
	"log"
)

// server is private (lowercase) - cannot be instantiated from outside package
type server struct {
	host string
	port int
	// private fields that need proper initialization
	isRunning bool
}

// NewServer is the public constructor - the ONLY way to create a server
func NewServer(host string, port int) (*server, error) {
	// Validation logic
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid port: %d", port)
	}
	if host == "" {
		host = "localhost"
	}

	// Proper initialization
	return &server{
		host:      host,
		port:      port,
		isRunning: false,
	}, nil
}

// Start is a method on server
func (s *server) Start() error {
	s.isRunning = true
	fmt.Printf("Server started on %s:%d\n", s.host, s.port)
	return nil
}

// Address returns the full address
func (s *server) Address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func main() {
	// Correct way - using constructor
	srv, err := NewServer("0.0.0.0", 8080)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server address:", srv.Address())
	srv.Start()

	// Invalid port - constructor catches this
	_, err = NewServer("localhost", 99999)
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Note: Direct initialization like `server{port: 8080}` would work
	// in this file (same package), but NOT from other packages!
	// That's the key benefit of this pattern.
}

