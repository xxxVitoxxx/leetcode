package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{1, 2, 3, 1},
		expected: 4,
	}, {
		name:     "example2",
		nums:     []int{2, 7, 9, 3, 1},
		expected: 12,
	}, {
		name:     "example3",
		nums:     []int{0},
		expected: 0,
	}, {
		name:     "example4",
		nums:     []int{2, 1},
		expected: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob2(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rob3(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
