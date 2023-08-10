package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// ------------------------
// NOTE FOR THE READER:
// ------------------------
// This program shows (as requested) the use of goroutines to sort a slice of integers.
// However, for small slices, it is actually slower than a regular sort.
// This is because of the overhead of creating goroutines, context switching, etc.
// If you want to see the difference, you can change the default value of the "size" flag to 100000 for example.
// For that, you can call the program with the following command:
// go run SortingGoroutines.go -size=100000
// or whichever size you want to test.
// On my PC, things even out at about 1000 elements in the slice, and goroutines are faster for larger slices.
// size=100000000 takes a few seconds to run, I dont recommend going higher than that.

func main() {

	// Define a flag. In this case, we expect an integer flag named "size" with a default value of 13.
	// The description (third argument) will be displayed in the default help message.
	var size int
	flag.IntVar(&size, "size", 13, "size of the slice to be sorted")

	// Parse the flags. This will read the user provided values, or if they're not provided, it will use the default values.
	flag.Parse()

	fmt.Println("Sorting a slice using goroutines and merge sort algorithm")
	fmt.Println("---------------------------------------------------------")
	fmt.Println()

	// We'll use the current time as a seed for the random number generator. If we don't do that, we'll always get the same random numbers.
	rand.Seed(time.Now().UnixNano())

	// Let's create a slice of integers
	slice := make([]int, size)

	// we'll fill this slice with random numbers
	randomise(slice)

	printSlice("Unsorted randomised slice", slice)

	// well divide the slice in 4 parts (about the same size, 4th part might be a bit bigger or smaller if the slice size is not a multiple of 4)

	fmt.Println("Dividing the slice in 4 parts")
	fmt.Println()

	quarter := len(slice) / 4

	slice1, slice2, slice3, slice4 := slice[:quarter], slice[quarter:quarter*2], slice[quarter*2:quarter*3], slice[quarter*3:]

	printSlice("1st slice", slice1)
	printSlice("2nd slice", slice2)
	printSlice("3rd slice", slice3)
	printSlice("4th slice", slice4)

	// we'll create 4 channels to send the sorted slices
	chans := make([]chan []int, 4)
	for i := range chans {
		chans[i] = make(chan []int)
	}

	// Start timer for regular functions
	start := time.Now()

	// we'll create 4 goroutines to sort each slice
	go Sort(slice1, chans[0])
	go Sort(slice2, chans[1])
	go Sort(slice3, chans[2])
	go Sort(slice4, chans[3])

	// we'll merge the 4 sorted slices into a new slice

	slice = merge(merge(<-chans[0], <-chans[1]),
		merge(<-chans[2], <-chans[3]))

	elapsed := time.Since(start)
	fmt.Println("Time taken to sort the 4 parts and merge them :", elapsed)
	fmt.Println()

	printSlice("Sorted slice", slice)

	//Now, just for fun, let's sort the slice using the regular sort function

	noGoroutinesSlice := make([]int, size)
	randomise(noGoroutinesSlice)

	printSlice("Unsorted randomised slice for the process without Goroutines", noGoroutinesSlice)

	// Start timer for regular functions
	start = time.Now()

	// sort the slice using the regular sort function

	sort.Ints(noGoroutinesSlice)
	elapsedStandard := time.Since(start)

	printSlice("Sorted slice using the regular sort function", noGoroutinesSlice)

	fmt.Println("Time taken to sort the slice using the regular sort function :", elapsedStandard)
	fmt.Println()

	comparePerformance(elapsed, elapsedStandard)

}

func randomise(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000)
	}
}

func printSlice(message string, slice []int) {
	if len(slice) > 50 {
		fmt.Printf("%s (%d integers in total, just showin the 50 first): %v\n", message, len(slice), slice[:50])
	} else {
		fmt.Printf("%s (%d integers): %v\n", message, len(slice), slice)
	}
	fmt.Println()
}

func Sort(slice []int, c chan []int) {

	// "My" Sort function just calls the regular sort function implemented in the sort package, and puts the result in the channel
	// This is just to show how to use channels to send data between goroutines, and so that the timing is not polluted by the time taken to sort the slice (which would not be optimised if I do it myself)
	sort.Ints(slice)

	c <- slice
}

// merge merges two slices of integers into a new slice
func merge(s1 []int, s2 []int) (s []int) {

	s = make([]int, len(s1)+len(s2))

	i, j := 0, 0

	for i < len(s1) && j < len(s2) {

		if s1[i] < s2[j] {
			s[i+j] = s1[i]
			i++
		} else {
			s[i+j] = s2[j]
			j++
		}
	}
	if i == len(s1) {
		copy(s[i+j:], s2[j:])
	} else {
		copy(s[i+j:], s1[i:])
	}
	return s
}

// comparePerformance compares the performance of the two algorithms and prints a message
func comparePerformance(elapsed1, elapsed2 time.Duration) {
	elapsed1Millis := float64(elapsed1) / float64(time.Millisecond)
	elapsed2Millis := float64(elapsed2) / float64(time.Millisecond)

	var difference float64
	var message string

	if elapsed1Millis < elapsed2Millis {
		difference = (elapsed2Millis - elapsed1Millis) / elapsed2Millis * 100
		message = fmt.Sprintf("Goroutines were faster by %.2f%%", difference)
	} else {
		difference = (elapsed1Millis - elapsed2Millis) / elapsed1Millis * 100
		message = fmt.Sprintf("Simple algorithm was faster by %.2f%%", difference)
	}

	fmt.Println(message)
}
