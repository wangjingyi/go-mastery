// Assignment 3: Tree Cancel
//
// Goal: Cancel parent context. Verify child contexts are also cancelled.
//
// Instructions:
// 1. Create parent context
// 2. Derive child contexts from parent
// 3. Cancel parent - all children should be cancelled too
//
// Run: go run main.go

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create parent context with cancel
	parentCtx, parentCancel := context.WithCancel(context.Background())

	// Create child contexts derived from parent
	child1Ctx, child1Cancel := context.WithCancel(parentCtx)
	defer child1Cancel() // Good practice even though parent cancel will propagate

	child2Ctx, child2Cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer child2Cancel()

	// Create grandchild context
	grandchildCtx, grandchildCancel := context.WithCancel(child1Ctx)
	defer grandchildCancel()

	// Start workers that watch their contexts
	go worker("Parent", parentCtx)
	go worker("Child1", child1Ctx)
	go worker("Child2", child2Ctx)
	go worker("Grandchild", grandchildCtx)

	// Let them run for a bit
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🛑 Cancelling PARENT context...")
	parentCancel() // This cancels ALL derived contexts!

	// Wait to see the cancellations
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n=== Context States ===")
	fmt.Printf("Parent:     %v\n", parentCtx.Err())
	fmt.Printf("Child1:     %v\n", child1Ctx.Err())
	fmt.Printf("Child2:     %v\n", child2Ctx.Err())
	fmt.Printf("Grandchild: %v\n", grandchildCtx.Err())
}

func worker(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Context cancelled: %v\n", name, ctx.Err())
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Printf("[%s] Working...\n", name)
		}
	}
}

// Context tree structure:
//
//     Background
//         │
//     parentCtx ────────► Cancel propagates down
//      /      \
//  child1Ctx  child2Ctx
//      │
//  grandchildCtx
//
// Cancelling parent cancels ALL descendants!
// But cancelling a child does NOT affect parent or siblings.

