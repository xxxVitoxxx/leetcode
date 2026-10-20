package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name     string
		heights  [][]int
		expected [][]int
	}{{
		name:     "example1",
		heights:  [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
		expected: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
	}, {
		name:     "example2",
		heights:  [][]int{{1}},
		expected: [][]int{{0, 0}},
	}, {
		name:     "example2",
		heights:  [][]int{{5, 5, 5, 5}, {4, 4, 4, 4}, {5, 5, 5, 5}},
		expected: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 0}, {2, 1}, {2, 2}, {2, 3}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pacificAtlantic(tt.heights)

			slices.SortFunc(actual, func(a, b []int) int {
				if a[0] != b[0] {
					return a[0] - b[0]
				}
				return a[1] - b[1]
			})

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pacificAtlantic2(tt.heights)

			slices.SortFunc(actual, func(a, b []int) int {
				if a[0] != b[0] {
					return a[0] - b[0]
				}
				return a[1] - b[1]
			})

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
