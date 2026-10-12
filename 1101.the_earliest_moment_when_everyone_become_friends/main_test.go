package main

import "testing"

func TestEarliestAcq(t *testing.T) {
	tests := []struct {
		name     string
		logs     [][]int
		n        int
		expected int
	}{{
		name:     "example1",
		logs:     [][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}},
		n:        6,
		expected: 20190301,
	}, {
		name:     "example2",
		logs:     [][]int{{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}},
		n:        4,
		expected: 3,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := earliestAcq(tt.logs, tt.n)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
