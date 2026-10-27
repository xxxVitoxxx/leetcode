package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{{
		name:     "example1",
		nums:     []int{2, 3, 1, 1, 4},
		expected: true,
	}, {
		name:     "example2",
		nums:     []int{3, 2, 1, 0, 4},
		expected: false,
	}, {
		name:     "example3",
		nums:     []int{1},
		expected: true,
	}, {
		name:     "example4",
		nums:     []int{2, 0},
		expected: true,
	}, {
		name:     "example5",
		nums:     []int{2, 5, 0, 0},
		expected: true,
	}, {
		name:     "example6",
		nums:     []int{2, 1, 0, 0},
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canJump(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canJump2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
