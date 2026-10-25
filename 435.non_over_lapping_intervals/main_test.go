package main

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{{
		name:      "example1",
		intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		expected:  1,
	}, {
		name:      "example2",
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		expected:  2,
	}, {
		name:      "example3",
		intervals: [][]int{{1, 2}, {2, 3}},
		expected:  0,
	}, {
		name:      "example4",
		intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
		expected:  2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := eraseOverlapIntervals(tt.intervals)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
