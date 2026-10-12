package main

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected int
	}{{
		name:     "example1",
		n:        5,
		edges:    [][]int{{0, 1}, {1, 2}, {3, 4}},
		expected: 2,
	}, {
		name:     "example2",
		n:        5,
		edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
		expected: 1,
	}, {
		name:     "example3",
		n:        5,
		edges:    [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 4}},
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countComponents(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countComponents2(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
