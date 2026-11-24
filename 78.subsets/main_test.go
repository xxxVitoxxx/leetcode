package main

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{{
		name:     "example1",
		nums:     []int{1, 2, 3},
		expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
	}, {
		name:     "example2",
		nums:     []int{0},
		expected: [][]int{{}, {0}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
