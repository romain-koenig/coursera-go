package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// We have a JSON data structure that holds information about animals. (was a table in the exercise)
// I took the liberty to add a kangaroo to the list.
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
	  ,
	  {
		"name": "kangaroo",
		"info": {
		  "food": "grass",
		  "locomotion": "jump",
		  "noise": "boing"
		}
	  }
	]
`

// We create an Animal "class" to hold the information about a single animal.
type Animal struct {
	Name string
	Info map[string]string
}

// We then define 3 methods, that will be used to get the information about the animal. These methode use the "class" Animal as a receiver type.

// Eat method will return the food the animal eats.
func (a Animal) Eat() string {
	return a.Info["food"]
}

// Move method will return the locomotion method of the animal.
func (a Animal) Move() string {
	return a.Info["locomotion"]
}

// Speak method will return the sound the animal makes.
func (a Animal) Speak() string {
	return a.Info["noise"]
}

// GetAnimalInformations function will be used to get the information about the animal the user is looking for. We give it the list of animals, the name of the animal and the information the user wants to know about it.
func GetAnimalInformations(animals []Animal, animalName string, information string) string {

	// We loop through the animals array to find the animal the user is looking for.

	for _, animal := range animals {
		if animal.Name == animalName {
			// We found the animal the user is looking for.
			// Let's check what information the user wants to know about the animal.

			switch information {
			case "eat":
				return animal.Eat()
			case "move":
				return animal.Move()
			case "speak":
				return animal.Speak()
			default:
				return "Unknown information request"

			}
		}
	}

	return "Animal not found."

	// Note: as we sanitized the user input, error cases should not happen here. We still check for them in case the code is modified later, or called from another function.
}

// Here is where the program starts.
func main() {
	fmt.Println("Animal Informations")
	fmt.Println("-------------------")
	fmt.Println("")

	// Let's create an Animal array and unmarshal the JSON data into it.

	var animals []Animal

	err := json.Unmarshal([]byte(data), &animals)

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

	command, err := ManageUserInput(animals)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	println(GetAnimalInformations(animals, command[0], command[1]))

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

// ManageUserInput is used to get the user input and sanitize it.
func ManageUserInput(animals []Animal) (values []string, err error) {

	availableAnimals := make([]string, 0, len(animals))
	// Let's parse the animals and see which animals are available.
	for _, animal := range animals {
		// if the animal is not already in the availableAnimals array, we add it.
		if !contains(availableAnimals, animal.Name) {
			availableAnimals = append(availableAnimals, animal.Name)
		}
	}

	fmt.Println("Enter a a command: an animal and an information request (eat, move or speak).")
	fmt.Println("Available animals: ", availableAnimals)
	fmt.Printf(`> `)

	values = make([]string, 0, 2)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return nil, err
	}

	input = strings.TrimSpace(input)
	values = strings.Split(input, " ")

	if len(values) != 2 {
		fmt.Println("Please enter 2 words: an animal (cow, bird or snake) and an information request (eat, move or speak).")
		err = fmt.Errorf("Invalid request")
		return nil, err
	}

	if !contains(availableAnimals, values[0]) {
		fmt.Println("Please enter an animal.")
		fmt.Println("Available animals: ", availableAnimals)
		err = fmt.Errorf("Invalid animal")
		return nil, err
	}

	if values[1] != "eat" && values[1] != "move" && values[1] != "speak" {
		fmt.Println("Please enter an information request (eat, move or speak).")
		err = fmt.Errorf("Invalid information request")
		return nil, err
	}

	return
}
