// Assignment 2: Worker Pool
//
// Goal: 5 workers. jobs channel. results channel. Process 100 items.
//
// Instructions:
// 1. Create a jobs channel and results channel
// 2. Start 5 worker goroutines
// 3. Send 100 jobs to the jobs channel
// 4. Collect all results
//
// Run: go run main.go

package main

import (
	"fmt"
	"time"
)

// Job represents work to be done
type Job struct {
	ID    int
	Value int
}

// Result represents the output of a job
type Result struct {
	JobID  int
	Output int
}

// worker processes jobs and sends results
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// Simulate some work
		time.Sleep(10 * time.Millisecond)

		// Process the job (square the value)
		output := job.Value * job.Value

		results <- Result{
			JobID:  job.ID,
			Output: output,
		}

		fmt.Printf("Worker %d processed job %d: %d -> %d\n", id, job.ID, job.Value, output)
	}
}

func main() {
	const numJobs = 100
	const numWorkers = 5

	jobs := make(chan Job, numJobs)       // Buffered channel
	results := make(chan Result, numJobs) // Buffered channel

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	start := time.Now()
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs) // Close to signal no more jobs

	// Collect results
	for r := 1; r <= numJobs; r++ {
		<-results // We could store these if needed
	}

	elapsed := time.Since(start)

	fmt.Printf("\nProcessed %d jobs with %d workers in %v\n", numJobs, numWorkers, elapsed)
	fmt.Printf("Sequential would take: ~%v\n", time.Duration(numJobs)*10*time.Millisecond)

	// Key insight: With 5 workers, jobs run in parallel
	// 100 jobs × 10ms = 1000ms sequential
	// With 5 workers: ~200ms (5x speedup!)
}

