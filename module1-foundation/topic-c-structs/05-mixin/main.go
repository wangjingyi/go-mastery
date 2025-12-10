// Assignment 5: The Mixin
//
// Goal: Create Drivable and Flyable structs.
//       Embed both in FlyingCar. Use methods from both.
//
// Instructions:
// 1. Create Drivable struct with Drive() method
// 2. Create Flyable struct with Fly() method
// 3. Create FlyingCar embedding both (multiple embedding)
// 4. Call both Drive() and Fly() on FlyingCar
//
// This is Go's "composition over inheritance" in action!

package main

import "fmt"

// Drivable provides driving capability
type Drivable struct {
	WheelCount int
	MaxSpeed   int // km/h on ground
}

func (d Drivable) Drive() {
	fmt.Printf("Driving with %d wheels at max %d km/h\n", d.WheelCount, d.MaxSpeed)
}

func (d Drivable) Honk() {
	fmt.Println("Beep beep!")
}

// Flyable provides flying capability
type Flyable struct {
	WingSpan   float64 // meters
	MaxAltitude int    // meters
}

func (f Flyable) Fly() {
	fmt.Printf("Flying with %.1fm wingspan up to %dm\n", f.WingSpan, f.MaxAltitude)
}

func (f Flyable) Land() {
	fmt.Println("Landing safely...")
}

// FlyingCar embeds BOTH - composition/mixin pattern
type FlyingCar struct {
	Drivable
	Flyable
	Model string
}

// FlyingCar can have its own methods too
func (fc FlyingCar) Transform() {
	fmt.Printf("%s is transforming from car to plane mode!\n", fc.Model)
}

func main() {
	car := FlyingCar{
		Drivable: Drivable{
			WheelCount: 4,
			MaxSpeed:   200,
		},
		Flyable: Flyable{
			WingSpan:    8.5,
			MaxAltitude: 3000,
		},
		Model: "SkyRider X1",
	}

	fmt.Printf("=== %s ===\n", car.Model)

	// Methods from Drivable are promoted
	car.Drive()
	car.Honk()

	// Methods from Flyable are promoted
	car.Transform()
	car.Fly()
	car.Land()

	// Direct field access works too
	fmt.Printf("\nSpecs: %d wheels, %.1fm wingspan\n", car.WheelCount, car.WingSpan)
}

