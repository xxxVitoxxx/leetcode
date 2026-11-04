package main

import "testing"

func TestMaxCount(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		ops      [][]int
		expected int
	}{{
		name:     "example1",
		m:        3,
		n:        3,
		ops:      [][]int{{2, 2}, {3, 3}},
		expected: 4,
	}, {
		name:     "example2",
		m:        3,
		n:        3,
		ops:      [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}},
		expected: 4,
	}, {
		name:     "example3",
		m:        3,
		n:        3,
		ops:      [][]int{},
		expected: 9,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxCount(tt.m, tt.n, tt.ops)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
