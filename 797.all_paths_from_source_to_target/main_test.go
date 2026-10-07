package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name     string
		graph    [][]int
		expected [][]int
	}{{
		name:     "example1",
		graph:    [][]int{{1, 2}, {3}, {3}, {}},
		expected: [][]int{{0, 1, 3}, {0, 2, 3}},
	}, {
		name:     "example2",
		graph:    [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
		expected: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := allPathsSourceTarget(tt.graph)

			if len(actual) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}

			for i := range actual {
				if !slices.ContainsFunc(tt.expected, func(expected []int) bool {
					return reflect.DeepEqual(actual[i], expected)
				}) {
					t.Errorf("expected %v, got %v", tt.expected, actual)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := allPathsSourceTarget2(tt.graph)

			if len(actual) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}

			for i := range actual {
				if !slices.ContainsFunc(tt.expected, func(expected []int) bool {
					return reflect.DeepEqual(actual[i], expected)
				}) {
					t.Errorf("expected %v, got %v", tt.expected, actual)
				}
			}
		})
	}
}
