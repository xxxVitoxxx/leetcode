package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{{
		name:       "example1",
		candidates: []int{10, 1, 2, 7, 6, 1, 5},
		target:     8,
		expected:   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	}, {
		name:       "example2",
		candidates: []int{2, 5, 2, 1, 2},
		target:     5,
		expected:   [][]int{{1, 2, 2}, {5}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			slices.SortFunc(got, slices.Compare)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
