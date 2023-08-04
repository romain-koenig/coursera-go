package main

import (
	"fmt"
	"strconv"
)

// let's define a function GenDisplaceFn that returns a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.

func GenDisplaceFn(a, v0, s0 float64) func(time float64) float64 {

	// The function returned by GenDisplaceFn takes one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

	return func(t float64) float64 {

		// Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.
		// s = ½ a t2 + vot + so
		return 0.5*a*t*t + v0*t + s0

	}

}

func main() {

	// First, well ask the user for the values of acceleration, initial velocity, and initial displacement.
	// Then, we'll ask the user for a value of time and compute the displacement after that time.
	// We'll then ask the user for a new value of time and compute the displacement after that time.
	// We'll continue asking the user for values of time until they press X to exit.

	var input string

	acceleration := askForFloat("Please enter the value of acceleration: ")
	initialVelocity := askForFloat("Please enter the value of initial velocity: ")
	initialDisplacement := askForFloat("Please enter the value of initial displacement: ")
	fmt.Println("")
	fmt.Println("Acceleration: ", acceleration, " Initial Velocity: ", initialVelocity, " Initial Displacement:", initialDisplacement)

	// Now that we have the values of acceleration, initial velocity, and initial displacement, we can call GenDisplaceFn to generate a function fn which will compute displacement as a function of time.
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	// Now we will ask the user for a value of time and compute the displacement after that time, until they press X to exit.

	for {

		fmt.Println("Please enter the value of time (or press X to eXit): ")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			break
		}
		time, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number or 'X' to exit.")
			continue
		}
		fmt.Println("Displacement after time ", time, " is ", fn(time))
		fmt.Println("")
	}

}

func askForFloat(prompt string) float64 {
	var result float64
	for {
		fmt.Println(prompt)
		_, err := fmt.Scan(&result)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		break
	}
	return result
}
