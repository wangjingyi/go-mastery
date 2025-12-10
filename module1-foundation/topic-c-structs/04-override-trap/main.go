// Assignment 4: The Override Trap
//
// Goal: Embed Base in Child. Give both a Describe() method.
//       Call child.Describe(). Call child.Base.Describe().
//
// Instructions:
// 1. Create Base with a Describe() method
// 2. Create Child that embeds Base
// 3. Give Child its own Describe() method (shadows Base.Describe)
// 4. Understand the difference between shadowing and overriding
//
// Key insight: Go doesn't have inheritance/override - it has SHADOWING

package main

import "fmt"

// Base struct with its own Describe method
type Base struct {
	Name string
}

func (b Base) Describe() string {
	return fmt.Sprintf("Base: %s", b.Name)
}

func (b Base) Greet() string {
	return fmt.Sprintf("Hello from Base %s", b.Name)
}

// Child embeds Base and shadows Describe
type Child struct {
	Base              // Embedded
	ChildField string
}

// Child's Describe SHADOWS (not overrides) Base's Describe
func (c Child) Describe() string {
	return fmt.Sprintf("Child: %s (child field: %s)", c.Name, c.ChildField)
}

func main() {
	child := Child{
		Base:       Base{Name: "MyEntity"},
		ChildField: "extra data",
	}

	// Calling child.Describe() uses Child's method (shadowing)
	fmt.Println(child.Describe())
	// Output: Child: MyEntity (child field: extra data)

	// You can still access the shadowed method explicitly
	fmt.Println(child.Base.Describe())
	// Output: Base: MyEntity

	// Greet() is NOT shadowed - it's promoted from Base
	fmt.Println(child.Greet())
	// Output: Hello from Base MyEntity

	// KEY DIFFERENCE FROM OOP:
	// In OOP with polymorphism, a Base reference to Child would call Child's method
	// In Go, there's no polymorphism through embedding - only through interfaces!
	
	// This would call Base.Describe, not Child.Describe:
	var base Base = child.Base
	fmt.Println("As Base:", base.Describe()) // Base: MyEntity
}

