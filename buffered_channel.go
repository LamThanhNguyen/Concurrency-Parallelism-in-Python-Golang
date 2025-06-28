package main

import "fmt"

func ____main() {
	ch := make(chan int, 2) // Buffered, size 2

	ch <- 1 // Does NOT block
	ch <- 2 // Does NOT block
	// Next send will block because buffer is full
	go func() {
		ch <- 3 // Blocks until receiver makes space
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
