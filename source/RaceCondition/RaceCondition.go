package main

import (
	"fmt"
	"sync"
)

// Main function - in this program, we'll explore Race conditions
func main() {
	var variable int = 0
	var wg sync.WaitGroup

	wg.Add(2) // We're adding 2 goroutines to our WaitGroup

	// Start 2 goroutines
	go increment(&variable, &wg)
	go increment(&variable, &wg)

	wg.Wait() // Wait for both goroutines to finish

	// Print the variable's value
	fmt.Println(variable)
}

// Goroutine definition
func increment(variable *int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	// Increment the variable 1000 times
	for i := 0; i < 1000000000; i++ {
		*variable++
	}
}
