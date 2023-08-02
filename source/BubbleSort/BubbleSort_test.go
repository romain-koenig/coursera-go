package main

import (
	"reflect"
	"testing"
)

func TestSortUserInput(t *testing.T) {

	testCases := []struct {
		name  string
		given []int
		want  []int
	}{
		{
			name:  "Multiple numbers in reverse order",
			given: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Empty slice",
			given: []int{},
			want:  []int{},
		},
		{
			name:  "Single value",
			given: []int{1},
			want:  []int{1},
		},
		{
			name:  "Negative numbers",
			given: []int{-5, -1, -3, -2, -4},
			want:  []int{-5, -4, -3, -2, -1},
		},
		{
			name:  "Negative and positive numbers",
			given: []int{-5, -1, 3, -2, 4},
			want:  []int{-5, -2, -1, 3, 4},
		},
		{
			name:  "Already sorted",
			given: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := make([]int, len(tc.given))
			copy(original, tc.given)

			BubbleSort(tc.given)
			if !reflect.DeepEqual(tc.given, tc.want) {
				t.Errorf("BubbleSort(%v) = %v; want %v", original, tc.given, tc.want)
			} else {
				t.Logf("BubbleSort(%v) = %v; want %v - is OK", original, tc.given, tc.want)
			}
		})
	}
}

func TestStringsToIntegers(t *testing.T) {
	// Normal Cases
	normalTestCases := []struct {
		name  string
		given []string
		want  []int
	}{
		{
			name:  "positive integers",
			given: []string{"1", "2", "3"},
			want:  []int{1, 2, 3},
		},
		{
			name:  "negative integers",
			given: []string{"-1", "-2", "-3"},
			want:  []int{-1, -2, -3},
		},
	}

	for _, tc := range normalTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := StringsToIntegers(tc.given)
			if err != nil {
				t.Fatalf("StringsToIntegers(%v) returned unexpected error: %v", tc.given, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("StringsToIntegers(%v) = %v; want %v", tc.given, got, tc.want)
			}
		})
	}

	// Error Cases
	errorTestCases := []struct {
		name  string
		given []string
	}{
		{
			name:  "non-integer strings",
			given: []string{"a", "b", "c"},
		},
	}

	for _, tc := range errorTestCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := StringsToIntegers(tc.given)
			if err == nil {
				t.Errorf("StringsToIntegers(%v) expected an error but got nil", tc.given)
			}
		})
	}
}
