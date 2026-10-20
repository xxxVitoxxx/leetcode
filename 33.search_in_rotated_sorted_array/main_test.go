package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{{
		name:     "example1",
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
	}, {
		name:     "example2",
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
	}, {
		name:     "example3",
		nums:     []int{1},
		target:   0,
		expected: -1,
	}, {
		name:     "example4",
		nums:     []int{1},
		target:   1,
		expected: 0,
	}, {
		name:     "example5",
		nums:     []int{1, 3},
		target:   3,
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := search(tt.nums, tt.target)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := search2(tt.nums, tt.target)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
