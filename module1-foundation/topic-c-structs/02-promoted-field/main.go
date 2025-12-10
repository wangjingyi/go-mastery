// Assignment 2: The Promoted Field
//
// Goal: Embed BaseEntity (ID, CreatedAt) into User.
//       Access user.ID directly (field promotion).
//
// Instructions:
// 1. Create BaseEntity with ID and CreatedAt fields
// 2. Embed BaseEntity into User (without field name)
// 3. Access user.ID directly - this is field promotion!
//
// TODO: Implement your solution below

package main

import (
	"fmt"
	"time"
)

// BaseEntity contains common fields for all entities
type BaseEntity struct {
	ID        int
	CreatedAt time.Time
}

// User embeds BaseEntity - fields are "promoted"
type User struct {
	BaseEntity // Embedded (no field name) - enables promotion
	Name       string
	Email      string
}

// Product also embeds BaseEntity
type Product struct {
	BaseEntity
	Title string
	Price float64
}

func main() {
	user := User{
		BaseEntity: BaseEntity{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// Field promotion - access embedded fields directly!
	fmt.Printf("User ID: %d\n", user.ID)             // Not user.BaseEntity.ID
	fmt.Printf("Created: %v\n", user.CreatedAt)      // Promoted field
	fmt.Printf("Name: %s\n", user.Name)

	// You can still access via the embedded type
	fmt.Printf("Via BaseEntity: %d\n", user.BaseEntity.ID)

	// TODO: Create a Product and demonstrate the same promotion
}

