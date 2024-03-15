package main

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name   string
		piles  []int
		h      int
		output int
	}{
		{"example1", []int{3, 6, 7, 11}, 8, 4},
		{"example2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"example3", []int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := minEatingSpeed(tt.piles, tt.h)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
