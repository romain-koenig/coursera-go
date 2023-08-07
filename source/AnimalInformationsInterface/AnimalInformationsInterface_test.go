package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestGetAnimalInformations(t *testing.T) {
	animals := []Animal{Cow{name: "Bessie"}, Bird{name: "John"}, Snake{name: "Alex"}}
	tests := []struct {
		name         string
		animals      []Animal
		animalName   string
		information  string
		wantResponse string
	}{
		{
			name:         "Cow Eating Test",
			animals:      animals,
			animalName:   "Bessie",
			information:  "eat",
			wantResponse: "grass\n",
		},
		{
			name:         "Bird Moving Test",
			animals:      animals,
			animalName:   "John",
			information:  "move",
			wantResponse: "fly\n",
		},

		{
			name:         "Snake Speaking Test",
			animals:      animals,
			animalName:   "Alex",
			information:  "speak",
			wantResponse: "hsss\n",
		},

		{
			name:         "Unknown Animal Test",
			animals:      animals,
			animalName:   "Bernie",
			information:  "speak",
			wantResponse: "Animal not found\n",
		},

		{
			name:         "Unknown Command Test",
			animals:      animals,
			animalName:   "Alex",
			information:  "sing",
			wantResponse: "Unknown information request\n",
		},
	}

	for _, tc := range tests {

		// Backup the real stdout.
		// Create a new pipe (reader and writer ends).
		// Set the writer end to stdout, so any print to stdout will write to the writer.
		oldStdout, r, w := BeforeTest()

		GetAnimalInformations(tc.animals, tc.animalName, tc.information)

		// Close writer end to signal to the reader that we're done.
		// Read everything from the reader end (this will be our function's output).
		// Restore real stdout.
		buf := AfterTest(w, r, oldStdout)

		if buf.String() != tc.wantResponse {
			t.Fatalf("Expected %q but got %q", tc.wantResponse, buf.String())
		}

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
			name:      "Cow Speaking Test",
			animal:    Cow{name: "Bessie"},
			wantSpeak: "moo\n",
		},

		{
			name:      "Snake Speaking Test",
			animal:    Snake{name: "Alex"},
			wantSpeak: "hsss\n",
		},
	}

	// Call the function.
	// Run test cases.
	for _, tc := range testCases {

		// Backup the real stdout.
		// Create a new pipe (reader and writer ends).
		// Set the writer end to stdout, so any print to stdout will write to the writer.
		oldStdout, r, w := BeforeTest()

		tc.animal.Speak()

		// Close writer end to signal to the reader that we're done.
		// Read everything from the reader end (this will be our function's output).
		// Restore real stdout.
		buf := AfterTest(w, r, oldStdout)

		// Check the function's output.

		if buf.String() != tc.wantSpeak {
			t.Fatalf("Expected %q but got %q", tc.wantSpeak, buf.String())
		}

	}

}

func AfterTest(w *os.File, r *os.File, oldStdout *os.File) bytes.Buffer {
	w.Close()

	var buf bytes.Buffer
	buf.ReadFrom(r)

	os.Stdout = oldStdout
	return buf
}

func BeforeTest() (*os.File, *os.File, *os.File) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w
	return oldStdout, r, w
}

func TestManageUserInput(t *testing.T) {

	animals := []Animal{Cow{name: "Bessie"}, Bird{name: "John"}, Snake{name: "Alex"}}

	var ErrInvalidAnimal = errors.New("Invalid animal")
	var ErrInvalidAnimalType = errors.New("Invalid animal type")
	var ErrInvalidCommand = errors.New("Invalid command")
	var ErrInvalidInformationRequest = errors.New("Invalid information request")
	var ErrExit = errors.New("exit")
	var ErrHelp = errors.New("help")

	tests := []struct {
		name         string
		input        string
		animals      []Animal
		wantResponse string
		wantErr      error
	}{
		{
			name:         "Valid query with available animal",
			input:        "query Bessie move",
			animals:      animals,
			wantResponse: "query Bessie move",
			wantErr:      nil,
		},
		{
			name:         "Valid query with badly written animal",
			input:        "query BeSSie move",
			animals:      animals,
			wantResponse: "query BeSSie move",
			wantErr:      nil,
		},
		{
			name:         "Valid query with unavailable animal",
			input:        "query Toby move",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidAnimal,
		},
		{
			name:         "Invalid query with available animal",
			input:        "query Alex talk",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidInformationRequest,
		},
		{
			name:         "Invalid query",
			input:        "toto Alex talk",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidCommand,
		},
		{
			name:         "Create Animal",
			input:        "newanimal Jojo snake",
			animals:      animals,
			wantResponse: "newanimal Jojo snake",
			wantErr:      nil,
		},
		{
			name:         "Create: Wrong Animal Type",
			input:        "newanimal Jojo kangaroo",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidAnimalType,
		},
		{
			name:         "Exit",
			input:        "exit",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrExit,
		},
		{
			name:         "Invalid command simple word",
			input:        "notacommand",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidCommand,
		},
		{
			name:         "Empty command",
			input:        "",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrInvalidCommand,
		},
		{
			name:         "Help command",
			input:        "help",
			animals:      animals,
			wantResponse: "",
			wantErr:      ErrHelp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ManageUserInput(tt.input, tt.animals)

			if (err != nil) && (tt.wantErr == nil || err.Error() != tt.wantErr.Error()) {
				t.Errorf("Expected error %v, but got %v", tt.wantErr, err)
			} else if (err == nil) && (tt.wantErr != nil) {
				t.Errorf("Expected error %v, but got nil", tt.wantErr)
			}

			if got != tt.wantResponse {
				t.Errorf("Expected %q but got %q", tt.wantResponse, got)
			}

			t.Logf("**** Got %q", got)

		})
	}

}
