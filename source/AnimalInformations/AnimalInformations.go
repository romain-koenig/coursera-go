package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// We have a JSON data structure that holds information about animals.
const data = `
	[
	  {
		"name": "cow",
		"info": {
		  "Food eaten": "grass",
		  "Locomotion method": "walk",
		  "Spoken sound": "moo"
		}
	  },
	  {
		"name": "bird",
		"info": {
		  "Food eaten": "worms",
		  "Locomotion method": "fly",
		  "Spoken sound": "peep"
		}
	  },
	  {
		"name": "snake",
		"info": {
		  "Food eaten": "mice",
		  "Locomotion method": "slither",
		  "Spoken sound": "hsss"
		}
	  }
	]
`

// We create an Animal "class" to hold the information about a single animal.
type Animal struct {
	Name string
	Info map[string]string
}

// The GetAnimalInformations function takes a string as input (from the user) and returns a string as output (which will be printed on screen).
func GetAnimalInformations(request string) string {
	return "grass"
}

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

	for _, animal := range animals {
		fmt.Println("Name:", animal.Name)
		for key, value := range animal.Info {
			fmt.Println(key, ":", value)
		}
		fmt.Println("")
	}

	// Let's ask the user for the name of the animal and the information they want to know about it.
	// We'll then call GetAnimalInformations to get the information and print it on screen.

}
