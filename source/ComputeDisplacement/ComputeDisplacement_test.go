package main

import (
	"math"
	"testing"
)

func TestGenDisplaceFn(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		name                 string
		acceleration         float64
		initialVelocity      float64
		initialDisplacement  float64
		time                 float64
		expectedDisplacement float64
	}{
		{
			name:                 "Example values",
			acceleration:         10,
			initialVelocity:      2,
			initialDisplacement:  1,
			time:                 3,
			expectedDisplacement: 52, // computed as 0.5*10*3^2 + 2*3 + 1
		},
		{
			name:                 "Example values",
			acceleration:         20,
			initialVelocity:      40,
			initialDisplacement:  120,
			time:                 90,
			expectedDisplacement: 84720, // computed as 0.5*20*90^2 + 40*90 + 120
		},
	}

	// Run test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fn := GenDisplaceFn(tc.acceleration, tc.initialVelocity, tc.initialDisplacement)
			got := fn(tc.time)
			if math.Abs(got-tc.expectedDisplacement) > 1e-6 {
				t.Errorf("Displacement = %v; want %v", got, tc.expectedDisplacement)
			}
		})
	}
}
