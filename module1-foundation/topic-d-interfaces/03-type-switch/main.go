// Assignment 3: The Type Switch
//
// Goal: Create a map[string]any. Store int, string, struct.
//       Iterate and use switch v := val.(type) to handle each.
//
// Instructions:
// 1. Create a map with string keys and any values
// 2. Store different types: int, string, bool, struct
// 3. Iterate and use type switch to handle each type
//
// Key insight: Type switches let you handle multiple types dynamically

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// map[string]any can hold values of ANY type
	data := map[string]any{
		"count":    42,
		"message":  "Hello, Go!",
		"enabled":  true,
		"ratio":    3.14159,
		"person":   Person{Name: "Alice", Age: 30},
		"numbers":  []int{1, 2, 3, 4, 5},
	}

	for key, val := range data {
		fmt.Printf("%s: ", key)
		processValue(val)
	}
}

func processValue(val any) {
	// Type switch - the Go way to handle dynamic types
	switch v := val.(type) {
	case int:
		fmt.Printf("int = %d (doubled: %d)\n", v, v*2)
	case string:
		fmt.Printf("string = %q (length: %d)\n", v, len(v))
	case bool:
		fmt.Printf("bool = %t (negated: %t)\n", v, !v)
	case float64:
		fmt.Printf("float64 = %.2f\n", v)
	case Person:
		fmt.Printf("Person: %s is %d years old\n", v.Name, v.Age)
	case []int:
		fmt.Printf("[]int with %d elements: %v\n", len(v), v)
	default:
		fmt.Printf("unknown type: %T = %v\n", v, v)
	}
}

// Type assertion (single type check)
func demonstrateTypeAssertion() {
	var val any = "hello"

	// Type assertion with ok check (safe)
	if str, ok := val.(string); ok {
		fmt.Println("It's a string:", str)
	}

	// Type assertion without ok check (panics if wrong type!)
	// str := val.(string) // This would panic if val is not a string
}

