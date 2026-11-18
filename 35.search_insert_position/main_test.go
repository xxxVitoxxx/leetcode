package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{{
		name:     "example1",
		nums:     []int{1, 3, 5, 6},
		target:   5,
		expected: 2,
	}, {
		name:     "example1",
		nums:     []int{1, 3, 5, 6},
		target:   2,
		expected: 1,
	}, {
		name:     "example1",
		nums:     []int{1, 3, 5, 6},
		target:   7,
		expected: 4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
