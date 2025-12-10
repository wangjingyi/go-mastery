// Assignment 1: The Shape Solver
//
// Goal: Define Shape interface (Area method).
//       Implement Circle, Rectangle. Write PrintArea(s Shape).
//
// Instructions:
// 1. Define Shape interface with Area() float64 method
// 2. Implement Circle with radius field
// 3. Implement Rectangle with width and height fields
// 4. Write PrintArea that accepts any Shape
//
// TODO: Implement your solution below

package main

import (
	"fmt"
	"math"
)

// Shape interface - any type with Area() method satisfies this
type Shape interface {
	Area() float64
}

// TODO: Implement Circle struct and its Area() method
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// Your code here
	return math.Pi * c.Radius * c.Radius
}

// TODO: Implement Rectangle struct and its Area() method
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	// Your code here
	return r.Width * r.Height
}

// PrintArea accepts ANY Shape - polymorphism through interfaces!
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 3}

	// Both satisfy the Shape interface
	PrintArea(circle)    // Area: 78.54
	PrintArea(rectangle) // Area: 30.00

	// You can store different shapes in a slice
	shapes := []Shape{circle, rectangle, Circle{Radius: 2}}
	
	fmt.Println("\nAll shapes:")
	for _, shape := range shapes {
		PrintArea(shape)
	}
}

