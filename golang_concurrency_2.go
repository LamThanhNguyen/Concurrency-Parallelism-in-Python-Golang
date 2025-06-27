package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Job represents a unit of work.
type Job struct {
	ID int
}

// Result holds the result of processing a job.
type Result struct {
	JobID    int
	WorkerID int
	Value    int
	Duration int64
	Err      error
}

// Worker function: processes jobs and sends results.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup, seed uint64, stream uint64) {
	defer wg.Done()
	// Each goroutine gets its own RNG, seeded deterministically
	rng := rand.New(rand.NewPCG(seed+uint64(id), stream+uint64(id)))
	for job := range jobs {
		// Simulate processing time
		duration := rng.Int64N(500)
		time.Sleep(time.Duration(duration) * time.Millisecond)
		// Simulate a random error
		if rng.Float32() < 0.2 {
			results <- Result{JobID: job.ID, WorkerID: id, Err: errors.New("random failure")}
		} else {
			results <- Result{JobID: job.ID, WorkerID: id, Value: job.ID * 100, Duration: duration, Err: nil}
		}
		fmt.Printf("[Worker %d] processed job %d (sleep=%dms)\n", id, job.ID, duration)
	}
}

func main() {
	numJobs := 100
	numWorkers := 4

	seed := uint64(time.Now().UnixNano())
	stream := uint64(time.Now().UnixNano() >> 1)

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg, seed, stream)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // no more jobs

	// Wait for all workers to finish, then close results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	success, failures := 0, 0
	for res := range results {
		if res.Err != nil {
			fmt.Printf("[Job %d by worker %d] (duration: %d ms) failed: %v\n", res.JobID, res.WorkerID, res.Duration, res.Err)
			failures++
		} else {
			fmt.Printf("[Job %d by worker %d] (duration: %d ms) result: %d\n", res.JobID, res.WorkerID, res.Duration, res.Value)
			success++
		}
	}
	fmt.Printf("\nTotal: %d succeeded, %d failed\n", success, failures)
}
