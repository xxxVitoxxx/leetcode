package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{100, 4, 200, 1, 3, 2},
		expected: 4,
	}, {
		name:     "example2",
		nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		expected: 9,
	}, {
		name:     "example3",
		nums:     []int{1, 0, 1, 2},
		expected: 3,
	}, {
		name:     "example4",
		nums:     []int{},
		expected: 0,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestConsecutive(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
