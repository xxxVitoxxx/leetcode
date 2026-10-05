package main

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      [][]int
	}{
		{
			name:          "example1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      [][]int{{0, 1}},
		},
		{
			name:          "example2",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}},
		},
		{
			name:          "example3",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      [][]int{{0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findOrder(tt.numCourses, tt.prerequisites)

			var anyCorrect bool
			for _, expected := range tt.expected {
				if reflect.DeepEqual(res, expected) {
					anyCorrect = true
					break
				}
			}
			if !anyCorrect {
				t.Errorf("expected: %v, got: %v", tt.expected, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findOrder2(tt.numCourses, tt.prerequisites)

			var anyCorrect bool
			for _, expected := range tt.expected {
				if reflect.DeepEqual(res, expected) {
					anyCorrect = true
					break
				}
			}
			if !anyCorrect {
				t.Errorf("expected: %v, got: %v", tt.expected, res)
			}
		})
	}
}
