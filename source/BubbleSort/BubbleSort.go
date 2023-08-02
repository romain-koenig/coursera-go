package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Dear reader, please note that I tried to make this code as simple as possible.
// Functions are small and do one thing.
// I hope you find it easy to read and understand.

func main() {

	numbers, err := GetNumbersToSort()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Sort the numbers.
	BubbleSort(numbers)

	// Print out the sorted numbers.
	PrintResult(numbers)
}

func PrintResult(numbers []int) {
	fmt.Println("Sorted numbers:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func GetNumbersToSort() (numbers []int, err error) {

	// Prompt the user for a series of numbers.
	// Split the input into a slice of strings on each space character.
	// Trim trailing newline
	// if the user entered more than 10 numbers, exit
	// Convert the slice of strings to a slice of integers.

	numStrings, err := ManageUserInput()
	if err != nil {
		return nil, err
	}

	numbers, err = StringsToIntegers(numStrings)
	if err != nil {
		return nil, err
	}
	return
}

func StringsToIntegers(numStrings []string) (numbers []int, err error) {
	// Create a slice of the appropriate size.
	numbers = make([]int, len(numStrings))

	for i, numStr := range numStrings {
		numbers[i], err = strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error reading numbers: ", err)
			return nil, err
		}
	}
	return numbers, nil
}

func ManageUserInput() (values []string, err error) {
	fmt.Println("Enter a series of integers (maximum 10), separated by spaces:")

	values = make([]string, 0, 10)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	input = strings.TrimSpace(input)
	values = strings.Split(input, " ")

	if len(values) > 10 {
		fmt.Println("Please enter 10 or fewer numbers.")
		err = fmt.Errorf("Too many numbers")
		return
	}
	return
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
	originalValueAtIndex := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = originalValueAtIndex
}
