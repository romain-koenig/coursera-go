package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	const size = 31
	// Let's create a slice of integers (40 integers)
	slice := make([]int, size)

	// we'll fill this slice with random numbers going from 0 to 100
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}

	fmt.Println("Unsorted slice: ", slice)

	// well divide the slice in 4 parts (about the same size)

	quarter := len(slice) / 4

	slice1, slice2, slice3, slice4 := slice[:quarter], slice[quarter:quarter*2], slice[quarter*2:quarter*3], slice[quarter*3:]

	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 3: ", slice3)
	fmt.Println("Slice 4: ", slice4)

	// we'll create 4 channels to send the sorted slices
	chans := make([]chan []int, 4)
	for i := range chans {
		chans[i] = make(chan []int)
	}

	// we'll create 4 goroutines to sort each slice
	go sort(slice1, chans[0])
	go sort(slice2, chans[1])
	go sort(slice3, chans[2])
	go sort(slice4, chans[3])

	// we'll merge the 4 sorted slices into a new slice

	slice = merge(merge(<-chans[0], <-chans[1]),
		merge(<-chans[2], <-chans[3]))

	fmt.Println("Sorted slice: ", slice)

}

func merge(s1 []int, s2 []int) (s []int) {

	// merge two slices

	// create a new slice with the size of the two slices combined
	s = make([]int, len(s1)+len(s2))

	// create two indexes, one for each slice
	i, j := 0, 0

	// loop through the two slices
	for i < len(s1) && j < len(s2) {

		// if the value of the first slice is smaller than the value of the second slice
		if s1[i] < s2[j] {

			// add the value of the first slice to the new slice
			s[i+j] = s1[i]

			// increment the index of the first slice
			i++

			// else
		} else {

			// add the value of the second slice to the new slice
			s[i+j] = s2[j]

			// increment the index of the second slice
			j++
		}
	}

	// if we reached the end of the first slice
	if i == len(s1) {

		// copy the remaining values of the second slice to the new slice
		copy(s[i+j:], s2[j:])
		// else
	} else {

		// copy the remaining values of the first slice to the new slice
		copy(s[i+j:], s1[i:])
	}

	return s

}

func sort(slice []int, c chan []int) {

	// sort the slice
	// Let's use the BubbleSort that we already implemented in another Assignement, and that was extensively tested
	BubbleSort(slice)

	c <- slice
}

// BubbleSort sorts a slice of integers using the bubble sort algorithm.

// Here is how a bubble sort works:
// 1. Start at the beginning of the list.
// 2. Compare the first two elements.
// 3. If the first is greater than the second, swap them.
// 4. Go to the next pair, and so on, continuously making sweeps of the list until sorted.
// 5. In doing so, the smaller items slowly "bubble" up to the beginning of the list.
// 6. If we ever make a sweep without swapping, the list is sorted and we can stop.

// We do it N-1 times, where N is the number of items in the list, to guarantee that it is sorted.
// If we make a sweep and make no swaps, the list is sorted and we can stop

func BubbleSort(numbers []int) {

	for i := 0; i < len(numbers); i++ {

		// At the beginning of each sweep, we setup a flag to indicate if we made a swap (we did not yet, so false)
		didSwap := false
		firstIndex := 0

		// We need to make sure that we don't go out of bounds
		for firstIndex < (len(numbers) - i - 1) {
			firstNumber := numbers[firstIndex]
			secondNumber := numbers[firstIndex+1]

			// If the first number is greater than the second number, swap them
			if firstNumber > secondNumber {
				Swap(numbers, firstIndex)
				// We made a swap, so set the flag to true
				didSwap = true
			}
			firstIndex++
		}
		// if we didn't make a swap, the list is sorted and we can stop
		if !didSwap {
			break
		}
	}
}

// Swap swaps the position of two elements in a slice of integers.
func Swap(numbers []int, index int) {
	// We need to make sure that we don't go out of bounds
	// We already checked this in BubbleSort, but we check again here for safety (if this code is later on called from another function)
	if index >= len(numbers)-1 {
		return
	}
	originalValueAtIndex := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = originalValueAtIndex
}
