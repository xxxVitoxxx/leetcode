package main

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{2, 3, -2, 4},
		expected: 6,
	}, {
		name:     "example2",
		nums:     []int{-2, 0, -1},
		expected: 0,
	}, {
		name:     "example1",
		nums:     []int{2, 3, -2, 4, -1},
		expected: 48,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxProduct(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxProduct2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
