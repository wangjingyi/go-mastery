// Assignment 5: DB Loop
//
// Goal: Run a loop doing work. Check ctx.Err() every iteration to abort early.
//
// Instructions:
// 1. Long-running operation with many iterations
// 2. Check context at each iteration
// 3. Abort early if context is cancelled
//
// Run: go run main.go

package main

import (
	"context"
	"fmt"
	"time"
)

// ProcessRecords simulates processing many database records
func ProcessRecords(ctx context.Context, records []int) (int, error) {
	processed := 0

	for i, record := range records {
		// Check context at the START of each iteration
		if err := ctx.Err(); err != nil {
			return processed, fmt.Errorf("cancelled at record %d: %w", i, err)
		}

		// Simulate work for each record
		time.Sleep(100 * time.Millisecond)
		processed++

		fmt.Printf("Processed record %d (value: %d)\n", i, record)
	}

	return processed, nil
}

// Alternative: Check context using select (non-blocking)
func ProcessRecordsWithSelect(ctx context.Context, records []int) (int, error) {
	processed := 0

	for i, record := range records {
		select {
		case <-ctx.Done():
			return processed, fmt.Errorf("cancelled at record %d: %w", i, ctx.Err())
		default:
			// Continue processing
		}

		time.Sleep(100 * time.Millisecond)
		processed++
		fmt.Printf("Processed record %d (value: %d)\n", i, record)
	}

	return processed, nil
}

func main() {
	records := make([]int, 20)
	for i := range records {
		records[i] = i * 10
	}

	fmt.Println("=== With 1s timeout (should cancel early) ===")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	start := time.Now()
	processed, err := ProcessRecords(ctx, records)
	elapsed := time.Since(start)

	fmt.Printf("\nProcessed %d/%d records in %v\n", processed, len(records), elapsed)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n=== With 5s timeout (should complete) ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	start = time.Now()
	processed, err = ProcessRecords(ctx2, records[:10]) // Only 10 records
	elapsed = time.Since(start)

	fmt.Printf("\nProcessed %d records in %v\n", processed, elapsed)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// Key patterns:
// 1. Check ctx.Err() at loop start (cheap, non-blocking)
// 2. Use select for more complex scenarios
// 3. Return partial results + error on cancellation
// 4. Include iteration info in error message for debugging

