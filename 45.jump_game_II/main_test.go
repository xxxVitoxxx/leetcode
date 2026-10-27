package main

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{{
		name:     "example1",
		nums:     []int{2, 3, 1, 1, 4},
		expected: 2,
	}, {
		name:     "example2",
		nums:     []int{2, 3, 0, 1, 4},
		expected: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := jump(tt.nums)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
