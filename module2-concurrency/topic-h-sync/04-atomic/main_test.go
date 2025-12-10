// Assignment 4: Atomic
//
// Goal: Replace Mutex with atomic.AddInt64 for the counter.
//       Benchmark the speed difference.
//
// Instructions:
// 1. Create both Mutex-based and Atomic-based counters
// 2. Benchmark them with go test -bench=.
// 3. Observe atomic operations are faster for simple operations
//
// Run: go test -bench=. -benchmem

package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// MutexCounter uses sync.Mutex
type MutexCounter struct {
	mu    sync.Mutex
	value int64
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// AtomicCounter uses sync/atomic
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Benchmarks

func BenchmarkMutexCounter(b *testing.B) {
	counter := &MutexCounter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
}

func BenchmarkAtomicCounter(b *testing.B) {
	counter := &AtomicCounter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
}

// Test correctness

func TestMutexCounter(t *testing.T) {
	counter := &MutexCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if counter.Value() != 1000 {
		t.Errorf("Expected 1000, got %d", counter.Value())
	}
}

func TestAtomicCounter(t *testing.T) {
	counter := &AtomicCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if counter.Value() != 1000 {
		t.Errorf("Expected 1000, got %d", counter.Value())
	}
}

// Key insight: Atomic operations are faster but limited
// Use atomic for: simple counters, flags, pointers
// Use mutex for: complex operations, multiple fields
