package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
		expected: 4,
	}, {
		name:     "example2",
		nums:     []int{0, 1, 0, 3, 2, 3},
		expected: 4,
	}, {
		name:     "example3",
		nums:     []int{7, 7, 7, 7, 7, 7},
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lengthOfLIS(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lengthOfLIS2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
