// Assignment 2: RWMutex
//
// Goal: Create a Cache. 100 readers, 1 writer. Use RLock vs Lock.
//
// Instructions:
// 1. RWMutex allows multiple concurrent readers
// 2. Writers get exclusive access
// 3. Use RLock/RUnlock for reads, Lock/Unlock for writes
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a thread-safe key-value cache
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewCache creates a new Cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get reads from cache (uses RLock - multiple readers allowed)
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()         // Read lock - shared access
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Set writes to cache (uses Lock - exclusive access)
func (c *Cache) Set(key, value string) {
	c.mu.Lock()         // Write lock - exclusive access
	defer c.mu.Unlock()
	c.data[key] = value
}

// Delete removes from cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	// Pre-populate cache
	cache.Set("name", "Gopher")
	cache.Set("language", "Go")

	// Start 100 readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				val, _ := cache.Get("name")
				_ = val
			}
			fmt.Printf("Reader %d done\n", id)
		}(i)
	}

	// Start 1 writer that updates periodically
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			cache.Set("counter", fmt.Sprintf("%d", i))
			fmt.Printf("Writer: Set counter to %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()

	// Final read
	val, _ := cache.Get("counter")
	fmt.Printf("Final counter value: %s\n", val)

	// Key insight: RWMutex is faster than Mutex when reads >> writes
	// Multiple readers can read simultaneously
	// Writers wait for all readers to finish, and block new readers
}

