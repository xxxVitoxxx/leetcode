package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{{
		name:     "example1",
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	}, {
		name:     "example2",
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxProfit(tt.prices)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
