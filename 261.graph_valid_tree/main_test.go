package main

import "testing"

func TestValidTree(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{{
		name:     "example1",
		n:        5,
		edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
		expected: true,
	}, {
		name:     "example2",
		n:        5,
		edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
		expected: false,
	}, {
		name:     "example3",
		n:        5,
		edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
		expected: false,
	}, {
		name:     "example4",
		n:        4,
		edges:    [][]int{{0, 1}, {2, 3}},
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTree(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := validTree2(tt.n, tt.edges)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
