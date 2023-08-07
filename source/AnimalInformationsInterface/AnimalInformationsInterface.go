package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// We have a JSON data structure that holds information about animals. (was a table in the exercise)
// We wil NOT use it in this version of the programm, let's just keep it as a reference.
const data = `
	[
	  {
		"name": "cow",
		"info": {
		  "food": "grass",
		  "locomotion": "walk",
		  "noise": "moo"
		}
	  },
	  {
		"name": "bird",
		"info": {
		  "food": "worms",
		  "locomotion": "fly",
		  "noise": "peep"
		}
	  },
	  {
		"name": "snake",
		"info": {
		  "food": "mice",
		  "locomotion": "slither",
		  "noise": "hsss"
		}
	  }
	]
`

// Let's define an interface, called Animal, that will be used to get the information about the animal.

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

// Let's define 3 animals, that will implement the Animalinterface.

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

// Let's define the methods that will implement the AnimalInterface interface.

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Cow) GetName() string {
	return c.name
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (b Bird) GetName() string {
	return b.name
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func (s Snake) GetName() string {
	return s.name
}

// GetAnimalInformations function will be used to get the information about the animal the user is looking for. We give it the list of animals, the name of the animal and the information the user wants to know about it.
func GetAnimalInformations(animals []Animal, animalName string, information string) {

	// We loop through the animals array to find the animal the user is looking for.
	animalfound := false

	for _, animal := range animals {

		if strings.ToLower(animal.GetName()) == strings.ToLower(animalName) {
			// We found the animal the user is looking for.
			// Let's check what information the user wants to know about the animal.

			animalfound = true

			switch information {
			case "eat":
				animal.Eat()
				break
			case "move":
				animal.Move()
				break
			case "speak":
				animal.Speak()
				break
			default:
				fmt.Println("Unknown information request")
			}
		}
		if animalfound {
			break
		}
	}

	if !animalfound {
		fmt.Println("Animal not found")
	}

	// Note: as we sanitized the user input, error cases should not happen here. We still check for them in case the code is modified later, or called from another function.
}

// Here is where the program starts.
func main() {
	fmt.Println("Animal Informations V2 - Interfaces")
	fmt.Println("-----------------------------------")

	// Let's create an Animal array.

	var animals []Animal

	// I create 3 random animals to fill the array, so it's simpler to test the program.
	// if you don't like it and want to start fresh, please change test to false.

	test := true
	// test := false
	if test {
		bessie := Cow{name: "Bessie"}
		tweety := Bird{name: "Tweety"}
		kobra := Snake{name: "Kobra"}

		animals = append(animals, bessie)
		animals = append(animals, tweety)
		animals = append(animals, kobra)
	}

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

	// The array is empty, the user will be able to create animals in it.

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

	PrintInstructions(animals)

	// Loop for user commands until "exit".
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		command, err := ManageUserInput(input, animals)
		if err != nil {
			if err.Error() == "exit" {
				fmt.Println("Exiting program.")
				break
			}
			if err.Error() == "help" {
				PrintInstructions(animals)
				continue
			} else {
				fmt.Println("Error: ", err)
			}
			continue

		}

		splitCommand := strings.Split(command, " ")

		switch splitCommand[0] {
		case "newanimal":
			animalName := splitCommand[1]
			animalType := splitCommand[2]
			// Create a new animal
			switch animalType {
			case "cow":
				animals = append(animals, Cow{name: animalName})
				fmt.Println("Created it!")
				break
			case "bird":
				animals = append(animals, Bird{name: animalName})
				fmt.Println("Created it!")
				break
			case "snake":
				animals = append(animals, Snake{name: animalName})
				fmt.Println("Created it!")
				break
			default:
				fmt.Println("Unknown animal type")
			}
			break
		case "query":
			// Query an animal
			// Get the animal name and the information the user wants to know about it.
			animalName := splitCommand[1]
			info := splitCommand[2]

			// Call GetAnimalInformations to get the information and print it on screen.
			GetAnimalInformations(animals, animalName, info)
			break
		default:
			fmt.Println("Unknown command")
		}
	}
}

// contains is used to check if a string is in a string array. It's just a helper function.
// Go version 1.21 should make it obsolete, please see https://tip.golang.org/doc/go1.21#slices
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// ManageUserInput function will be used to manage the user input. It will check if the input is valid, and return the command to execute. It's a helper function designed to sanitize the user input.
func ManageUserInput(input string, animals []Animal) (query string, err error) {

	//We make a list of available animals to check if the user input is valid. This list is lowercase to make the check case insensitive.
	availableAnimals := make([]string, 0, len(animals))
	for _, animal := range animals {
		if !contains(availableAnimals, strings.ToLower(animal.GetName())) {
			availableAnimals = append(availableAnimals, strings.ToLower(animal.GetName()))
		}
	}

	input = strings.TrimSpace(input)

	// Special case: the user wants to exit the program.
	if input == "exit" || input == "Exit" {
		return "", fmt.Errorf("exit")
	}

	// Special case: the user wants some help.
	if input == "help" || input == "Help" {
		return "", fmt.Errorf("help")
	}

	values := strings.Split(input, " ")

	// The user must enter 3 values. Unless they want to exit or help. Which has been checked before.
	if len(values) != 3 {
		return "", fmt.Errorf("Invalid command")
	}

	values[0] = strings.ToLower(values[0])
	values[2] = strings.ToLower(values[2])

	// Let's check if the user input is valid.
	// The command must be "newanimal" or "query".
	if values[0] != "newanimal" && values[0] != "query" {
		return "", fmt.Errorf("Invalid command")
	}

	//In case of a "query" The animal name must be known.
	if values[0] == "query" && !contains(availableAnimals, strings.ToLower(values[1])) {
		return "", fmt.Errorf("Invalid animal")
	}
	// In case of a "query" the information must be either "eat", "move" or "speak".
	if values[0] == "query" && values[2] != "eat" && values[2] != "move" && values[2] != "speak" {
		return "", fmt.Errorf("Invalid information request")
	}

	// In case of a "newanimal" the animal type must be either "cow", "bird" or "snake".
	if values[0] == "newanimal" && (values[2] != "cow" && values[2] != "bird" && values[2] != "snake") {
		return "", fmt.Errorf("Invalid animal type")
	}

	// If we're here, the input is valid. We return the command to execute.
	return values[0] + " " + values[1] + " " + values[2], nil
}

// PrintInstructions function will print the instructions to the user. Centralized here to avoid code duplication.
func PrintInstructions(availableAnimals []Animal) {
	fmt.Println("")
	fmt.Println("Enter a command followed by parameters")
	fmt.Println("newanimal <animal name> <animal type> (animal type = cow, bird or snake)")
	fmt.Println("query <animal name> <information> (information = eat, move or speak)")
	fmt.Println("Enter \"exit\" to exit the program.")
	fmt.Println("Existing animals: ", availableAnimals)
}
