package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyComputation(id, n int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	sum := 0
	for range n {
		// Some fake CPU work, e.g., compute Fibonacci(20)
		a, b := 0, 1
		for range 20 {
			a, b = b, a+b
		}
		sum += a
	}
	fmt.Printf("Worker %d finished\n", id)
	results <- sum
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())

	numWorkers := 1000
	tasksPerWorker := 100000

	var wg sync.WaitGroup
	results := make(chan int, numWorkers)

	start := time.Now()

	for i := range numWorkers {
		wg.Add(1)
		go heavyComputation(i, tasksPerWorker, &wg, results)
	}

	wg.Wait()
	close(results)

	total := 0
	for r := range results {
		total += r
	}

	elapsed := time.Since(start)
	fmt.Printf("Total result: %d\n", total)
	fmt.Printf("Elapsed: %s (on %d CPUs)\n", elapsed, runtime.NumCPU())
}
