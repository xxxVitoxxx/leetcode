package main

import (
	"reflect"
	"testing"
)

func TestGetModifiedArray(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		updates  [][]int
		expected []int
	}{{
		name:     "example1",
		length:   5,
		updates:  [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}},
		expected: []int{-2, 0, 3, 5, 3},
	}, {
		name:     "example2",
		length:   10,
		updates:  [][]int{{2, 4, 6}, {5, 6, 8}, {1, 9, -4}},
		expected: []int{0, -4, 2, 2, 2, 4, 4, -4, -4, -4},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getModifiedArray(tt.length, tt.updates)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
