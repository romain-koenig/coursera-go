package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ------------------------
// NOTE FOR THE READER:
// ------------------------
// This program shows (as requested) the use of goroutines to sort a slice of integers.
// However, for small slices, it is actually slower than a regular sort.
// This is because of the overhead of creating goroutines, context switching, etc.
// As it woul be impossible to type in enough numbers to see the difference, there is a buil-in auto mode
// If you want to see the difference, you can set the program to auto mode
// and change the default value of the "size" flag to a larger number 100000 for example.
// For that, you can call the program with the following command:
// go run SortingGoroutines.go - auto -size=100000
// On my PC, things even out at about 1000 elements in the slice, and goroutines are faster for larger slices.
// size=100000000 takes a few seconds to run, I dont recommend going higher than that.

func main() {

	fmt.Println("Sorting a slice using goroutines and merge sort algorithm")
	fmt.Println("---------------------------------------------------------")
	fmt.Println()

	autoMode := flag.Bool("auto", false, "generate random slice automatically")
	size := flag.Int("size", 100, "size of the slice to be sorted (only used with -auto flag)")
	flag.Parse()

	var slice []int

	if !*autoMode {
		// Prompt user for input - This is the standard case for this program, as required by the assignement
		fmt.Println("Please enter integers to sort, separated by spaces:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputs := strings.Split(input, " ")

		for _, s := range inputs {
			num, err := strconv.Atoi(s)
			if err == nil {
				slice = append(slice, num)
			} else {
				fmt.Printf("'%s' is not a valid integer and will be skipped.\n", s)
			}
		}
	} else {
		// Here we are in auto mode, we'll generate a random slice of integers
		rand.Seed(time.Now().UnixNano())
		// Generate random numbers
		slice = make([]int, *size)
		randomise(slice)

	}

	printSlice("Unsorted slice", slice)

	// well divide the slice in 4 parts (about the same size, 4th part might be a bit bigger or smaller if the slice size is not a multiple of 4)

	fmt.Println("Dividing the slice in 4 parts")
	fmt.Println()

	quarter := len(slice) / 4

	slice1, slice2, slice3, slice4 := slice[:quarter], slice[quarter:quarter*2], slice[quarter*2:quarter*3], slice[quarter*3:]

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

	if *autoMode {
		//Now, just for fun, let's sort the slice using the regular sort function
		// This is only done in auto mode as the results would not be relevant if the user entered the numbers themselves

		noGoroutinesSlice := make([]int, *size)
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

	printSlice("In a GOROUTINE - Unsorted part of the slice", slice)

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
