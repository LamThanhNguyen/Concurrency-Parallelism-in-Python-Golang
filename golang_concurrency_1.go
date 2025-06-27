package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Simulate fetching data from a source
func fetchData(source int, results chan<- int, wg *sync.WaitGroup, seed uint64, stream uint64) {
	defer wg.Done()
	// Each goroutine gets its own RNG, seeded deterministically
	rng := rand.New(rand.NewPCG(seed+uint64(source), stream+uint64(source)))
	time.Sleep(time.Duration(rng.Int64N(500)) * time.Millisecond) // simulate variable work time
	fmt.Printf("Source %d fetched\n", source)
	results <- source * 10 // return some result
}

// if you want run this file, just replace _main() to main()
// and run go run golang_concurrency_1.go
func _main() {
	seed := uint64(time.Now().UnixNano())
	stream := uint64(time.Now().UnixNano() >> 1)

	numSources := 100
	results := make(chan int, numSources)
	var wg sync.WaitGroup

	for i := 1; i <= numSources; i++ {
		wg.Add(1)
		go fetchData(i, results, &wg, seed, stream)
	}

	// Wait for all fetches to complete, then close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect all results as they come in
	sum := 0
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
		sum += result
	}

	fmt.Printf("Total sum: %d\n", sum)
}
