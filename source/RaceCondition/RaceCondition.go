package main

import (
	"fmt"
	"sync"
	"time"
)

// In this program, we'll explore Race conditions and how to avoid them.
// First, we'll create a variable and increment it N times with 2 goroutines.
// Then, we'll do the same thing without goroutines (simple functions) and compare the results.
// Last

// Note: in the number of iterations is small enough, you might not see any difference between the 2 methods.
// Note: we set up a timer to see if the goroutines go faster than the simple functions. You may or may not see a difference.

// Note : if you run
// go run -race RaceCondition.go
// You will see:
// ==================
// WARNING: DATA RACE
// Read at 0x00c0000bc008 by goroutine 8:

// This is very important, because as stated above, if we set the number of increment to 1000, we will never see the race condition during our tests. But it's there. And it can be very hard to find.

var numberOfIncrementations = 1000000000

func main() {

	// First, we do it with goRoutines
	var goRoutinesVariable int = 0
	var wg sync.WaitGroup

	// Start timer for goroutines
	start := time.Now()

	wg.Add(2) // We're adding 2 goroutines to our WaitGroup

	// Start 2 goroutines
	go increment(&goRoutinesVariable, &wg)
	go increment(&goRoutinesVariable, &wg)

	wg.Wait() // Wait for both goroutines to finish

	// Stop timer for goroutines and print duration
	elapsed := time.Since(start)
	fmt.Println("Time taken with Goroutines:", elapsed)

	// Print the variable's value
	fmt.Printf("With Goroutines: %d (should be %d)\n", goRoutinesVariable, numberOfIncrementations*2)

	var simpleVariable int = 0

	fmt.Println("------------------------")

	// Now, we'll do it with simple functions ( = NO CONCURRENCY)

	// Start timer for regular functions
	start = time.Now()

	increment(&simpleVariable, nil)
	increment(&simpleVariable, nil)

	// Stop timer for regular functions and print duration
	elapsed = time.Since(start)
	fmt.Println("Time taken without Goroutines:", elapsed)

	// Print the variable's value
	fmt.Printf("Without Goroutines: %d (should be %d)\n", simpleVariable, numberOfIncrementations*2)

	// What we usually see here is a totally wrong answer with goroutines, and the right answer without goroutines.
	// We also usually see that it was faster with goroutines, at least on a commputer with multiuple cores.

	fmt.Println("------------------------")

	// Now, can we do it fast, with goroutines, and with the right answer?

	// Start timer for goroutines

	start = time.Now()

	var goRoutinesAndMutexVariable int = 0

	// We'll do it with a Mutex

	var mu sync.Mutex // Declare a mutex

	wg.Add(2)

	go incrementWithMutex(&goRoutinesAndMutexVariable, &wg, &mu)
	go incrementWithMutex(&goRoutinesAndMutexVariable, &wg, &mu)

	wg.Wait()

	fmt.Println("With Goroutines and Mutex: ", goRoutinesAndMutexVariable)

	// Stop timer for goroutines and print duration
	elapsed = time.Since(start)
	fmt.Println("Time taken with Goroutines and MUTEX :", elapsed)

	// Print the variable's value
	fmt.Printf("With Goroutines and MUTEX: %d (should be %d)\n", goRoutinesAndMutexVariable, numberOfIncrementations*2)

}

func increment(variable *int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // Decrement the counter when the goroutine completes
	}

	// Increment the variable 1000 times
	for i := 0; i < numberOfIncrementations; i++ {
		*variable++
	}
}

func simpleIncrement(variable *int) {

	// Increment the variable 1000 times
	for i := 0; i < numberOfIncrementations; i++ {
		*variable++
	}
}
func incrementWithMutex(variable *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < numberOfIncrementations; i++ {
		mutex.Lock() // Lock access to the variable
		*variable++
		mutex.Unlock() // Unlock access to the variable
	}
}

// Here is a result I got:

// Time taken with Goroutines: 2.143966372s
// With Goroutines: 1000853654 (should be 2000000000)
// ------------------------
// Time taken without Goroutines: 2.661362776s
// Without Goroutines: 2000000000 (should be 2000000000)
// ------------------------
// With Goroutines and Mutex:  2000000000
// Time taken with Goroutines and MUTEX : 31.479680098s
// With Goroutines and MUTEX: 2000000000 (should be 2000000000)

// So : we got the right answer with goroutines and mutex, but it was MUCH slower than without goroutines.
