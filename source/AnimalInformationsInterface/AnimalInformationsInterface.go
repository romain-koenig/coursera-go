package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Here is or new program. It is similar to the previous version of the program, except that we use the Animal interface instead of the Animal struct.

func main() {
	fmt.Println("Animal Informations V2 - Interfaces")
	fmt.Println("-----------------------------------")

	// Let's create an empty Animal slice, that will be used to store the animals created by the user.

	var animals []Animal

	// Here we print out some instructions so the user knows what to do. They can see it again by typing "help".
	PrintInstructions(animals)

	// Loop for user commands until "exit".
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Technical error: on reading user input ", err)
			continue
		}

		// We sanitize the user input, by doing all the checks in the ManageUserInput function.
		command, err := ManageUserInput(input, animals)

		// Error cases
		if err != nil {
			// Special case: user wants to exit the program.
			if err.Error() == "exit" {
				fmt.Println("Exiting program.")
				break
			}
			// Special case: user wants to see the instructions again.
			if err.Error() == "help" {
				PrintInstructions(animals)
				continue
			} else {
				// All other errors
				fmt.Println("Error: ", err)
			}
			continue
		}

		// We have a valid command, let's execute it.
		animals = ExecuteCommand(command, animals)
	}
}

// Let's define an interface, called Animal, that will be used to get the information about the animal.
// Eat, Move and Speak are the methods that will be used to get the information about the animal.
// GetName is the method that will be used to get the name of the animal.

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

// Let's define 3 animal types, that will each implement the Animalinterface.

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

// ----------------------------

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

// ----------------------------

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

// ----------------------------

// GetAnimalInformations function will be used to get the information about the animal the user is looking for. We give it the list of animals, the name of the animal and the information the user wants to know about it.
func GetAnimalInformations(animals []Animal, animalName string, information string) {

	// We loop through the animals array to find the animal the user is looking for.
	animalfound := false

	for _, animal := range animals {

		// We use the GetName method to get the name of the animal, and we compare it to the animalName parameter.
		// strings.ToLower is used to make the comparison case insensitive.
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

// ExecuteCommand function will be used to execute the command the user entered. We give it the command and the list of animals.
// It returns the list of animals, as it can be modified by the user commands.
// PLEASE NOTE: we pass a slice, which is kind of a pointer to the array. So we don't need to return the slice, as it is already modified. BUT, if we were to append to the slice (and we are), we would need to return it, as the slice would be copied and the original slice would not be modified. When we add new elements to the slice, and it increases its capacity, the slice is copied to a new array, and the original slice is not modified. So we need to return the slice, and assign it to the original slice.
// (dear reader, I'm sorry for the long comments, but I'm also writing to my future self, who will probably forget about this in a few months).
func ExecuteCommand(command string, animals []Animal) []Animal {
	splitCommand := strings.Split(command, " ")

	command = strings.ToLower(splitCommand[0])

	switch command {

	case "newanimal":
		animalName := splitCommand[1]
		animalType := splitCommand[2]

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
		animalName := splitCommand[1]
		info := splitCommand[2]

		GetAnimalInformations(animals, animalName, info)
		break
	default:
		fmt.Println("Unknown command")
	}
	return animals
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
	fmt.Println("Enter \"help\" to display this help again.")
	// if no animals are available, we don't display the list of available animals.
	if len(availableAnimals) == 0 {
		fmt.Println("No animals available yet.")
	} else {
		fmt.Println("Existing animals: ", Map(availableAnimals, func(animal Animal) string {
			// Here we do some type assertion (not that is really usefull, but it's a good example)
			if _, ok := animal.(Cow); ok {
				return animal.GetName() + " (cow)"
			} else if _, ok := animal.(Bird); ok {
				return animal.GetName() + " (bird)"
			} else if _, ok := animal.(Snake); ok {
				return animal.GetName() + " (snake)"
			} else {
				return ""
			}

		}))
	}

}

// Map function is used to map a slice of Animal to a slice of string. It's a helper function. It's done because GO doesn't have generics. No Map, no Filter, no Reduce. We have to do it ourselves.
func Map(data []Animal, fn func(Animal) string) []string {
	result := make([]string, len(data))
	for i, v := range data {
		result[i] = fn(v)
	}
	return result
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
