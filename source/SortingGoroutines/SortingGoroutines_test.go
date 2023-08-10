package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{name: "typical case", input: []int{3, 1, 4, 1, 5, 9, 2}, output: []int{1, 1, 2, 3, 4, 5, 9}},
		{name: "empty slice", input: []int{}, output: []int{}},
		{name: "singleton slice", input: []int{5}, output: []int{5}},
		{name: "reversed", input: []int{9, 8, 7, 6}, output: []int{6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make(chan []int)
			go Sort(tt.input, c)

			result := <-c
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("Expected %v, but got %v", tt.output, result)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   []int
		expected []int
	}{
		{name: "typical merge", s1: []int{1, 3, 5}, s2: []int{2, 4, 6}, expected: []int{1, 2, 3, 4, 5, 6}},
		{name: "merge with empty left", s1: []int{1, 2, 3}, s2: []int{}, expected: []int{1, 2, 3}},
		{name: "merge with empty right", s1: []int{}, s2: []int{4, 5, 6}, expected: []int{4, 5, 6}},
		{name: "merge without overlap", s1: []int{7, 8}, s2: []int{9, 10}, expected: []int{7, 8, 9, 10}},
		{name: "merge with both empty", s1: []int{}, s2: []int{}, expected: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.s1, tt.s2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Merging %v and %v: expected %v, but got %v", tt.s1, tt.s2, tt.expected, result)
			}
		})
	}
}
