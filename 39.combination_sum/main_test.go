package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{{
		name:       "example1",
		candidates: []int{2, 3, 6, 7},
		target:     7,
		expected:   [][]int{{2, 2, 3}, {7}},
	}, {
		name:       "example2",
		candidates: []int{2, 3, 5},
		target:     8,
		expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
	}, {
		name:       "example3",
		candidates: []int{2},
		target:     1,
		expected:   [][]int{},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := combinationSum(tt.candidates, tt.target)

			for i := range actual {
				slices.Sort(actual[i])
			}

			if len(actual) > 0 && len(tt.expected) > 0 {
				if !reflect.DeepEqual(actual, tt.expected) {
					t.Errorf("got: %v, want: %v", actual, tt.expected)
				}
			}

		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := combinationSum2(tt.candidates, tt.target)

			for i := range actual {
				slices.Sort(actual[i])
			}

			if len(actual) > 0 && len(tt.expected) > 0 {
				if !reflect.DeepEqual(actual, tt.expected) {
					t.Errorf("got: %v, want: %v", actual, tt.expected)
				}
			}

		})
	}
}
