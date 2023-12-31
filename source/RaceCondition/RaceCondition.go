package main

import (
	"fmt"
	"sync"
	"time"
)

// In this program, we'll explore Race conditions and how to avoid them.

// There is a CONCLUSION / TLDR at the end of this file. You can skip to it if you want.

// We'll create a variable and increment it N times in two separate functions. The total at the end should be 2N.

// Note: in the number of iterations is small enough, you might not see any difference between the different methods.
// Note: we set up a timer to see the difference in execution time between the different methods.

// Note : if you run
// go run -race RaceCondition.go
// You will see:
// ==================
// WARNING: DATA RACE
// Read at 0x00c0000bc008 by goroutine 8:

// This is very important, because as stated above, if we set the number of increment to a small enough number (for example 1000), we will never see the race condition problem during our tests. This flag help us to see it.

var numberOfIncrementations = 1000000000 // make it large enough

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
	// The result MAY be false

	fmt.Println("------------------------")

	// Now, we'll do it with simple functions ( = NO CONCURRENCY)
	var simpleVariable int = 0

	// Start timer for regular functions
	start = time.Now()

	increment(&simpleVariable, nil)
	increment(&simpleVariable, nil)

	// Stop timer for regular functions and print duration
	elapsed = time.Since(start)
	fmt.Println("Time taken without Goroutines:", elapsed)

	// Print the variable's value
	fmt.Printf("Without Goroutines: %d (should be %d)\n", simpleVariable, numberOfIncrementations*2)

	// The result here should be OK. It's usually a bit slower than the goroutines version. But slow and right is better than fast and wrong.

	fmt.Println("------------------------")

	// Now, can we do it fast, with goroutines, and with the right answer?

	// Start timer for goroutines

	start = time.Now()

	var goRoutinesAndMutexVariable int = 0

	// We'll do it with a Mutex

	var mutex sync.Mutex // Declare a mutex

	wg.Add(2)

	go incrementWithMutex(&goRoutinesAndMutexVariable, &wg, &mutex)
	go incrementWithMutex(&goRoutinesAndMutexVariable, &wg, &mutex)

	wg.Wait()

	fmt.Println("With Goroutines and Mutex: ", goRoutinesAndMutexVariable)

	// Stop timer for goroutines and print duration
	elapsed = time.Since(start)
	fmt.Println("Time taken with Goroutines and MUTEX :", elapsed)

	// Print the variable's value
	fmt.Printf("With Goroutines and MUTEX: %d (should be %d)\n", goRoutinesAndMutexVariable, numberOfIncrementations*2)

	// Here we should have the right result, but it's much slower than the other methods.
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

// Let's hope that the rest of the course will help us to find a better solution. As a matter of fact, based on this example, I would stick to the simple function version, because it's fast enough and it gives the right answer.

// Thanks for your time!

// ================== CONCLUSION / TLDR ==================

// Race Condition:

// A race condition is a situation where two or more concurrent threads or processes attempt to modify shared data at the same time, leading to unpredictable and often erroneous outcomes. The exact outcome depends on the relative timing of their operations.

// How it Can Occur in this code:

//     With Goroutines: When two goroutines (increment) run concurrently and attempt to increment the shared variable goRoutinesVariable, both might read and write to it almost simultaneously. Since the read-modify-write is not an atomic operation, one goroutine might overwrite the change made by another, leading to inconsistent results.

//     Result: The final value of goRoutinesVariable is unpredictable and often less than the expected value (numberOfIncrementations * 2).

//     Without Goroutines: We simply call the increment function twice sequentially. There's no concurrency, so no race condition arises.

//     Result: The final value of simpleVariable is consistent and matches the expected value.

//     With Goroutines and Mutex: By introducing a mutex, we ensure that only one goroutine at a time can increment the shared variable, effectively serializing the increment operations.

//     Result: The final value of goRoutinesAndMutexVariable is consistent and matches the expected value. However, due to the overhead of locking and unlocking the mutex for every iteration, this approach is significantly slower.
