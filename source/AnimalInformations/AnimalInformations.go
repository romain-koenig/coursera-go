package main

import (
	"fmt"
)

// We have a JSON data structure that holds information about animals.
var AnimalInformations = []byte(`{
	"animals": [
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
  }
`)

// We create an Animal "class" to hold the information about a single animal.
type Animal struct {
	Name string
	Info map[string]string
}

func GetAnimalInformations(request string) string {
	return "grass"
}

func main() {
	fmt.Println("Animal Informations")
	fmt.Println("-------------------")
	fmt.Println("")

}
