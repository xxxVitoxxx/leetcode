package main

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{"example1", []int{1, 2, 3, 4, 5}, true},
		{"example2", []int{5, 4, 3, 2, 1}, false},
		{"example3", []int{2, 1, 5, 0, 4, 6}, true},
		{"example4", []int{4, 5, 2147483647, 1, 2}, true},
		{"example5", []int{1, 5, 0, 4, 1, 3}, true},
		{"example6", []int{20, 100, 10, 12, 5, 13}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := increasingTriplet(tt.input)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
