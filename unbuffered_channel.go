package main

import "fmt"

func ___main() {
	ch := make(chan int) // Unbuffered

	go func() {
		ch <- 100 // Blocks until main goroutine receives
	}()

	val := <-ch      // Main goroutine blocks until value sent
	fmt.Println(val) // Prints 42

	fmt.Println("done!")
}
