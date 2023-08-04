package main

import (
	"testing"
)

func TestGetAnimalInformations(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "Example values for COW",
			given: "cow eat",
			want:  "grass",
		},
		{
			name:  "Example values for BIRD",
			given: "bird move",
			want:  "fly",
		},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAnimalInformations(tc.given)
			if got != tc.want {
				t.Errorf("GetAnimalInformations(%v) = %v; want %v", tc.given, got, tc.want)
			}
		})
	}
}
