package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	}, {
		name:     "example2",
		nums:     []int{1},
		expected: 1,
	}, {
		name:     "example3",
		nums:     []int{5, 4, -1, 7, 8},
		expected: 23,
	}, {
		name:     "example4",
		nums:     []int{-1},
		expected: -1,
	}, {
		name:     "example6",
		nums:     []int{4, -1, -2, 1, 5, 4},
		expected: 11,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxSubArray(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxSubArray2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
