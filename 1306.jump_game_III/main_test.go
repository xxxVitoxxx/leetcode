package main

import "testing"

func TestCanReach(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		start    int
		expected bool
	}{{
		name:     "example1",
		arr:      []int{4, 2, 3, 0, 3, 1, 2},
		start:    5,
		expected: true,
	}, {
		name:     "example2",
		arr:      []int{4, 2, 3, 0, 3, 1, 2},
		start:    0,
		expected: true,
	}, {
		name:     "example3",
		arr:      []int{3, 0, 2, 1, 2},
		start:    2,
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canReach(tt.arr, tt.start)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canReach2(tt.arr, tt.start)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := canReach3(tt.arr, tt.start)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
