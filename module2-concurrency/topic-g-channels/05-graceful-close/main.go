// Assignment 5: Graceful Close
//
// Goal: Producer sends 10 items then closes channel.
//       Consumer loops using range until closed.
//
// Instructions:
// 1. Producer sends items and then closes the channel
// 2. Consumer uses for-range to receive until closed
// 3. Demonstrate the "comma ok" idiom as well
//
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Using for-range (idiomatic) ===")
	rangeExample()

	fmt.Println("\n=== Using comma-ok idiom ===")
	commaOkExample()

	fmt.Println("\n=== Multiple consumers ===")
	multipleConsumers()
}

func rangeExample() {
	ch := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch) // Signal that we're done
		fmt.Println("Producer: Channel closed")
	}()

	// Consumer using for-range (cleanest way)
	for val := range ch {
		fmt.Printf("Consumer: Received %d\n", val)
	}
	fmt.Println("Consumer: Channel closed, loop exited")
}

func commaOkExample() {
	ch := make(chan string)

	go func() {
		ch <- "first"
		ch <- "second"
		close(ch)
	}()

	// Using comma-ok idiom
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed (ok=false)")
			break
		}
		fmt.Printf("Received: %s (ok=%t)\n", val, ok)
	}
}

func multipleConsumers() {
	ch := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Multiple consumers share the work
	done := make(chan struct{})

	for worker := 1; worker <= 3; worker++ {
		go func(id int) {
			for val := range ch {
				fmt.Printf("Worker %d: processed %d\n", id, val)
				time.Sleep(30 * time.Millisecond)
			}
			done <- struct{}{}
		}(worker)
	}

	// Wait for all workers
	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println("All workers done!")
}

// Important rules about closing channels:
// 1. Only the sender should close a channel
// 2. Never close a channel from the receiver side
// 3. Never close a channel twice (causes panic)
// 4. Sending on a closed channel causes panic
// 5. Receiving from a closed channel returns zero value immediately

