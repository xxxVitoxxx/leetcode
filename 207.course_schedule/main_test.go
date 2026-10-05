package main

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			"example1",
			2,
			[][]int{{1, 0}},
			true,
		},
		{
			"example2",
			2,
			[][]int{{1, 0}, {0, 1}},
			false,
		},
		{
			"example3",
			4,
			[][]int{{3, 1}, {1, 0}, {3, 2}, {2, 0}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canFinish(tt.numCourses, tt.prerequisites)
			if tt.expected != res {
				t.Fatal("expected: ", tt.expected, " got: ", res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canFinish2(tt.numCourses, tt.prerequisites)
			if tt.expected != res {
				t.Fatal("expected: ", tt.expected, " got: ", res)
			}
		})
	}
}
