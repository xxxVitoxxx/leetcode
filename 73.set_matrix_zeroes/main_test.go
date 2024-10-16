package main

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	for _, tt := range []struct {
		matrix [][]int
		want   [][]int
	}{{
		matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
	}, {
		matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
		want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
	}} {
		matrix1 := shallowCopy2DSlice(tt.matrix)
		setZeroes(matrix1)
		for i := 0; i < len(matrix1); i++ {
			if !equal(tt.want[i], matrix1[i]) {
				t.Errorf("got: %v, want: %v", matrix1[i], tt.want[i])
			}
		}

		matrix2 := shallowCopy2DSlice(tt.matrix)
		setZeroes2(matrix2)
		for i := 0; i < len(matrix2); i++ {
			if !equal(tt.want[i], matrix2[i]) {
				t.Errorf("got: %v, want: %v", matrix2[i], tt.want[i])
			}
		}

		matrix3 := shallowCopy2DSlice(tt.matrix)
		setZeroes3(matrix3)
		for i := 0; i < len(matrix3); i++ {
			if !equal(tt.want[i], matrix3[i]) {
				t.Errorf("got: %v, want: %v", matrix3[i], tt.want[i])
			}
		}
	}
}

func shallowCopy2DSlice(matrix [][]int) [][]int {
	shallow := make([][]int, len(matrix))
	for i := range matrix {
		arr := make([]int, len(matrix[0]))
		copy(arr, matrix[i])
		shallow[i] = arr
	}
	return shallow
}

func equal(expect, actual []int) bool {
	if len(expect) != len(actual) {
		return false
	}

	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			return false
		}
	}
	return true
}
