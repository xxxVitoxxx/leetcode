package main

import "testing"

func TestSmallestStringWithSwaps(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		pairs    [][]int
		expected string
	}{{
		name:     "example1",
		s:        "dcab",
		pairs:    [][]int{{0, 3}, {1, 2}},
		expected: "bacd",
	}, {
		name:     "example2",
		s:        "dcab",
		pairs:    [][]int{{0, 3}, {1, 2}, {0, 2}},
		expected: "abcd",
	}, {
		name:     "example3",
		s:        "cba",
		pairs:    [][]int{{0, 1}, {1, 2}},
		expected: "abc",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := smallestStringWithSwaps(tt.s, tt.pairs)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
