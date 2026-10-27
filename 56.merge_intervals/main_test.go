package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{{
		name:      "example1",
		intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
	}, {
		name:      "example2",
		intervals: [][]int{{1, 4}, {4, 5}},
		expected:  [][]int{{1, 5}},
	}, {
		name:      "example3",
		intervals: [][]int{{4, 7}, {1, 4}},
		expected:  [][]int{{1, 7}},
	}, {
		name:      "example4",
		intervals: [][]int{{1, 4}, {2, 3}},
		expected:  [][]int{{1, 4}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := merge(tt.intervals)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
