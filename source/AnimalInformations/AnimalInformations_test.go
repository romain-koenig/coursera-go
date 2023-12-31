package main

import (
	"testing"
)

var constAnimals = []Animal{
	{
		Name: "cow",
		Info: map[string]string{
			"food":       "grass",
			"locomotion": "walk",
			"noise":      "moo",
		},
	},
	{
		Name: "bird",
		Info: map[string]string{
			"food":       "worms",
			"locomotion": "fly",
			"noise":      "peep",
		},
	},
	{
		Name: "snake",
		Info: map[string]string{
			"food":       "mice",
			"locomotion": "slither",
			"noise":      "hsss",
		},
	},
}

func TestGetAnimalInformations(t *testing.T) {
	tests := []struct {
		name         string
		animals      []Animal
		animalName   string
		information  string
		wantResponse string
	}{
		{
			name:         "Cow Eating Test",
			animals:      constAnimals,
			animalName:   "cow",
			information:  "eat",
			wantResponse: "grass",
		},
		{
			name:         "Bird Moving Test",
			animals:      constAnimals,
			animalName:   "bird",
			information:  "move",
			wantResponse: "fly",
		},

		{
			name:         "Snake Speaking Test",
			animals:      constAnimals,
			animalName:   "snake",
			information:  "speak",
			wantResponse: "hsss",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if gotResponse := GetAnimalInformations(tc.animals, tc.animalName, tc.information); gotResponse != tc.wantResponse {
				t.Errorf("GetAnimalInformations(%v, %v, %v) = %v, want %v", tc.animals, tc.animalName, tc.information, gotResponse, tc.wantResponse)
			}
		})
	}
}

func TestEat(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		name    string
		animal  Animal
		wantEat string
	}{
		{
			name: "Cow Eating Test",
			animal: Animal{
				Name: "cow",
				Info: map[string]string{
					"food": "grass",
				},
			},
			wantEat: "grass",
		},
		{
			name: "Snake Eating Test",
			animal: Animal{
				Name: "snake",
				Info: map[string]string{
					"food": "mice",
				},
			},
			wantEat: "mice",
		},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if gotEat := tc.animal.Eat(); gotEat != tc.wantEat {
				t.Errorf("Animal.Eat() = %v, want %v", gotEat, tc.wantEat)
			}
		})
	}
}

func TestMove(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		name     string
		animal   Animal
		wantMove string
	}{
		{
			name: "Cow Moving Test",
			animal: Animal{
				Name: "cow",
				Info: map[string]string{
					"locomotion": "walk",
				},
			},
			wantMove: "walk",
		},
		{
			name: "Snake Moving Test",
			animal: Animal{
				Name: "snake",
				Info: map[string]string{
					"locomotion": "slither",
				},
			},
			wantMove: "slither",
		},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if gotMove := tc.animal.Move(); gotMove != tc.wantMove {
				t.Errorf("Animal.Move() = %v, want %v", gotMove, tc.wantMove)
			}
		})
	}
}

func TestSpeak(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		name      string
		animal    Animal
		wantSpeak string
	}{
		{
			name: "Cow Speaking Test",
			animal: Animal{
				Name: "cow",
				Info: map[string]string{
					"noise": "moo",
				},
			},
			wantSpeak: "moo",
		},
		{
			name: "Snake Speaking Test",
			animal: Animal{
				Name: "snake",
				Info: map[string]string{
					"noise": "hsss",
				},
			},
			wantSpeak: "hsss",
		},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if gotSpeak := tc.animal.Speak(); gotSpeak != tc.wantSpeak {
				t.Errorf("Animal.Speak() = %v, want %v", gotSpeak, tc.wantSpeak)
			}
		})
	}
}
