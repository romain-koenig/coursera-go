package main

import (
	"reflect"
	"testing"
)

func TestSortUserInput(t *testing.T) {
	given := []int{3, 2, 1}
	want := []int{1, 2, 3}

	got := BubbleSort(given)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BubbleSort(%v) = %v; want %v", given, got, want)
	}
}
