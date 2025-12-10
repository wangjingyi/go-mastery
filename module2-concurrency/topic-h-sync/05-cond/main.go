// Assignment 5: Cond
//
// Goal: Use sync.Cond to broadcast a "Start" signal to 10 waiting runners.
//
// Instructions:
// 1. sync.Cond is for complex coordination (wait/signal/broadcast)
// 2. Runners wait on a condition
// 3. Starter broadcasts to wake all runners
//
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

type Race struct {
	cond    *sync.Cond
	started bool
}

func NewRace() *Race {
	return &Race{
		cond:    sync.NewCond(&sync.Mutex{}),
		started: false,
	}
}

// WaitForStart blocks until the race starts
func (r *Race) WaitForStart(runnerID int) {
	r.cond.L.Lock()
	for !r.started {
		fmt.Printf("Runner %d: Ready and waiting...\n", runnerID)
		r.cond.Wait() // Releases lock and waits, re-acquires lock when woken
	}
	r.cond.L.Unlock()
	fmt.Printf("Runner %d: Started running! 🏃\n", runnerID)
}

// StartRace broadcasts to all waiting runners
func (r *Race) StartRace() {
	r.cond.L.Lock()
	r.started = true
	r.cond.L.Unlock()

	fmt.Println("\n🏁 START! 🏁\n")
	r.cond.Broadcast() // Wake ALL waiting goroutines
}

func main() {
	race := NewRace()
	var wg sync.WaitGroup

	// Start 10 runners
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			race.WaitForStart(id)

			// Simulate running
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Runner %d: Finished! 🎉\n", id)
		}(i)
	}

	// Wait for runners to get ready
	time.Sleep(500 * time.Millisecond)

	// Start the race!
	race.StartRace()

	wg.Wait()
	fmt.Println("\nAll runners finished!")
}

// sync.Cond methods:
// - Wait(): Releases lock, waits for signal, re-acquires lock
// - Signal(): Wakes one waiting goroutine
// - Broadcast(): Wakes ALL waiting goroutines

// Important: Always check the condition in a loop!
// for !condition {
//     cond.Wait()
// }
// This handles spurious wakeups and race conditions

