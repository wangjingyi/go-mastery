// Assignment 1: The JSON Modeler
//
// Goal: Map a complex nested JSON response to Go structs using struct tags.
//       Unmarshal it.
//
// Instructions:
// 1. Define structs matching the JSON structure below
// 2. Use `json:"field_name"` tags for proper mapping
// 3. Unmarshal the JSON and print the data
//
// TODO: Implement your solution below

package main

import (
	"encoding/json"
	"fmt"
)

// Sample JSON to parse:
var sampleJSON = `{
	"user": {
		"id": 12345,
		"username": "gopher",
		"profile": {
			"full_name": "Go Gopher",
			"email": "gopher@golang.org",
			"verified": true
		},
		"posts": [
			{"id": 1, "title": "Hello Go", "likes": 100},
			{"id": 2, "title": "Concurrency Rocks", "likes": 250}
		]
	},
	"meta": {
		"request_id": "abc-123",
		"timestamp": "2024-01-15T10:30:00Z"
	}
}`

// TODO: Define your structs here
// type Post struct { ... }
// type Profile struct { ... }
// type User struct { ... }
// type Meta struct { ... }
// type Response struct { ... }

func main() {
	// TODO: Unmarshal sampleJSON into your struct
	// var response Response
	// err := json.Unmarshal([]byte(sampleJSON), &response)

	// TODO: Print the parsed data
	// fmt.Printf("Username: %s\n", response.User.Username)
	// fmt.Printf("Email: %s\n", response.User.Profile.Email)
	// fmt.Printf("First post: %s\n", response.User.Posts[0].Title)

	fmt.Println("TODO: Implement JSON parsing")
	_ = json.Unmarshal // Silence unused import
}

