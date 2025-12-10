// Assignment 3: Singleton
//
// Goal: Use sync.Once to ensure InitDB() runs exactly once despite concurrent calls.
//
// Instructions:
// 1. sync.Once guarantees a function runs exactly once
// 2. Even with concurrent calls, only one execution happens
// 3. All other callers wait until it completes
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// Database is a singleton connection
type Database struct {
	Host string
}

var (
	dbOnce     sync.Once
	dbInstance *Database
)

// GetDB returns the singleton database instance
func GetDB() *Database {
	dbOnce.Do(func() {
		fmt.Println("Initializing database connection...")
		time.Sleep(100 * time.Millisecond) // Simulate slow init
		dbInstance = &Database{Host: "localhost:5432"}
		fmt.Println("Database initialized!")
	})
	return dbInstance
}

func main() {
	var wg sync.WaitGroup

	// 10 goroutines all try to get the DB simultaneously
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			db := GetDB()
			fmt.Printf("Goroutine %d: Got DB with host %s\n", id, db.Host)
		}(i)
	}

	wg.Wait()

	// All goroutines got the same instance
	db1 := GetDB()
	db2 := GetDB()
	fmt.Printf("\nSame instance? %v\n", db1 == db2)
}

// Alternative: Package-level initialization (also runs once)
// var db = initDB()

// sync.Once is better when:
// 1. Initialization is expensive and might not be needed
// 2. Initialization might fail and you want to handle it
// 3. You want to defer initialization until first use (lazy loading)

// Note: sync.Once cannot be reset or rerun
// If you need retry logic, use a different pattern

