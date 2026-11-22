package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{{
		name:     "example1",
		digits:   []int{1, 2, 3},
		expected: []int{1, 2, 4},
	}, {
		name:     "example2",
		digits:   []int{4, 3, 2, 1},
		expected: []int{4, 3, 2, 2},
	}, {
		name:     "example3",
		digits:   []int{9},
		expected: []int{1, 0},
	}, {
		name:     "example4",
		digits:   []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
		expected: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
	}, {
		name:     "example5",
		digits:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		expected: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
