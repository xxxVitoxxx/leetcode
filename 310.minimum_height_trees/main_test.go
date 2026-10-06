package main

import (
	"reflect"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected []int
	}{
		{
			name:     "example1",
			n:        4,
			edges:    [][]int{{1, 0}, {1, 2}, {1, 3}},
			expected: []int{1},
		},
		{
			name:     "example2",
			n:        6,
			edges:    [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			expected: []int{3, 4},
		},
		{
			name:     "example3",
			n:        1,
			edges:    [][]int{},
			expected: []int{0},
		},
		{
			name:     "example4",
			n:        2,
			edges:    [][]int{{0, 1}},
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMinHeightTrees(tt.n, tt.edges)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, res)
			}
		})
	}
}
