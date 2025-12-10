// Assignment 1: Ping Pong
//
// Goal: Two goroutines passing an int back and forth on a channel, incrementing it.
//
// Instructions:
// 1. Create two channels for ping and pong
// 2. One goroutine receives on ping, increments, sends on pong
// 3. Other goroutine receives on pong, increments, sends on ping
// 4. Stop after reaching a certain count
//
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan int)
	pong := make(chan int)

	// Player 1: receives from ping, sends to pong
	go func() {
		for val := range ping {
			fmt.Printf("🏓 Ping received %d\n", val)
			time.Sleep(200 * time.Millisecond) // Slow down for visibility
			pong <- val + 1
		}
		close(pong) // Close pong when ping is closed
	}()

	// Player 2: receives from pong, sends to ping
	go func() {
		for val := range pong {
			fmt.Printf("🏸 Pong received %d\n", val)
			time.Sleep(200 * time.Millisecond)
			if val >= 10 {
				close(ping) // Stop the game
				return
			}
			ping <- val + 1
		}
	}()

	// Start the game
	ping <- 1

	// Wait for game to finish
	time.Sleep(3 * time.Second)
	fmt.Println("Game over!")
}

