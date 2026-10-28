package main

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{{
		name:     "example1",
		m:        3,
		n:        7,
		expected: 28,
	}, {
		name:     "example2",
		m:        3,
		n:        2,
		expected: 3,
	}, {
		name:     "example3",
		m:        51,
		n:        9,
		expected: 1916797311,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uniquePaths(tt.m, tt.n)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uniquePaths2(tt.m, tt.n)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uniquePaths3(tt.m, tt.n)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
