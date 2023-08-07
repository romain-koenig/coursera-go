package main

import (
	"fmt"
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

		if animal.GetName() == animalName {
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

	// Let's create an Animal array and unmarshal the JSON data into it.

	var animals []Animal

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

	bessie := Cow{name: "Bessie"}
	tweety := Bird{name: "Tweety"}
	kobra := Snake{name: "Kobra"}

	animals = append(animals, bessie)
	animals = append(animals, tweety)
	animals = append(animals, kobra)

	GetAnimalInformations(animals, "Bessie", "eat")

	ManageUserInput("newanimal Bessie cow", animals)

	// The array is empty, the user will be able to create animals in it.

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

	// Loop for user commands until "exit".
	// for {
	// 	_, err := ManageUserInput("", animals)
	// 	if err != nil {
	// 		if err.Error() == "exit" {
	// 			fmt.Println("Exiting program.")
	// 			break
	// 		}
	// 		fmt.Println("Error: ", err)
	// 		continue
	// 	}

	// 	// result := GetAnimalInformations(animals, animalName, info)
	// 	// println(result)
	// }

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

func ManageUserInput(input string, animals []Animal) (query string, err error) {

	availableAnimals := make([]string, 0, len(animals))
	for _, animal := range animals {
		if !contains(availableAnimals, animal.GetName()) {
			availableAnimals = append(availableAnimals, animal.GetName())
		}
	}

	PrintInstructions(availableAnimals)

	input = strings.TrimSpace(input)

	// if input == "exit" || input == "Exit" {
	// 	return "", fmt.Errorf("exit")
	// }

	values := strings.Split(input, " ")

	// if len(values) != 3 {
	// 	return "", fmt.Errorf("Invalid request")
	// }

	if values[0] != "newanimal" && values[0] != "query" {
		return "", fmt.Errorf("Invalid command")
	}

	if values[0] == "query" && !contains(availableAnimals, values[1]) {
		return "", fmt.Errorf("Invalid animal")
	}

	// if values[0] == "newanimal" && (values[2] != "cow" && values[2] != "bird" && values[2] != "snake") {
	// 	return "", fmt.Errorf("Invalid animal type")
	// }

	if values[0] == "query" && values[2] != "eat" && values[2] != "move" && values[2] != "speak" {
		return "", fmt.Errorf("Invalid information request")
	}

	return values[0] + " " + values[1] + " " + values[2], nil
}

func PrintInstructions(availableAnimals []string) {
	fmt.Println("")
	fmt.Println("Enter a command followed by parameters")
	fmt.Println("newanimal <animal name> < animal type> (animal type = cow, bird or snake)")
	fmt.Println("query <animal name> <information> (information = eat, move or speak)")
	fmt.Println("Enter \"exit\" to exit the program.")
	fmt.Println("Available animals: ", availableAnimals)
	fmt.Printf(`> `)
}
