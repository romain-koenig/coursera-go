package main

import "fmt"

// let's define a function GenDisplaceFn that returns a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.
// For example, let's say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, v0 = 2, s0 = 1.
// I can use the following statement to call GenDisplaceFn to generate a function fn which will compute displacement as a function of time.
// fn := GenDisplaceFn(10, 2, 1)
// Then I can use the following statement to print the displacement after 3 seconds.
// fmt.Println(fn(3))
// and I can use the following statement to print the displacement after 5 seconds.
// fmt.Println(fn(5))

func GenDisplaceFn(a, v0, s0 float64) func(time float64) float64 {

	return func(t float64) float64 {

		return 0.5*a*t*t + v0*t + s0

	}

}

func main() {

	// First, well ask the user for the values of acceleration, initial velocity, and initial displacement.
	// Then, we'll ask the user for a value of time and compute the displacement after that time.
	// We'll then ask the user for a new value of time and compute the displacement after that time.

	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Println("Please enter the value of acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Println("Please enter the value of initial velocity: ")
	fmt.Scan(&initialVelocity)

	fmt.Println("Please enter the value of initial displacement: ")
	fmt.Scan(&initialDisplacement)

	fmt.Println("Please enter the value of time: ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println(fn(time))

	fmt.Println("Please enter another value of time: ")
	fmt.Scan(&time)

	fmt.Println(fn(time))

}
