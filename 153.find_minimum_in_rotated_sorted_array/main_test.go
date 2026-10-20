package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{3, 4, 5, 1, 2},
		expected: 1,
	}, {
		name:     "example1",
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		expected: 0,
	}, {
		name:     "example1",
		nums:     []int{11, 13, 15, 17},
		expected: 11,
	}, {
		name:     "example1",
		nums:     []int{3, 1, 2},
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findMin(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
