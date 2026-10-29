package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{2, 3, 2},
		expected: 3,
	}, {
		name:     "example2",
		nums:     []int{1, 2, 3, 1},
		expected: 4,
	}, {
		name:     "example3",
		nums:     []int{1, 2, 3},
		expected: 3,
	}, {
		name:     "example4",
		nums:     []int{1, 2, 1, 1},
		expected: 3,
	}, {
		name:     "example5",
		nums:     []int{1, 1, 2, 1, 5},
		expected: 7,
	}, {
		name:     "example6",
		nums:     []int{1, 3, 1, 3, 100},
		expected: 103,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
