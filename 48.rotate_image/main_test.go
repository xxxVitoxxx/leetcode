package main

import (
	"testing"

	"github.com/xxxVitoxxx/leetcode/testfunc"
)

func TestRotate(t *testing.T) {
	for _, tt := range []struct {
		matrix [][]int
		want   [][]int
	}{{
		matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	}, {
		matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	}} {
		matrix1 := shallowCopy2DSlice(tt.matrix)
		rotate(matrix1)
		if !testfunc.ElementsMatch(matrix1, tt.want) {
			t.Errorf("got: %v, want: %v", matrix1, tt.want)
		}

		matrix2 := shallowCopy2DSlice(tt.matrix)
		rotate2(matrix2)
		if !testfunc.ElementsMatch(matrix2, tt.want) {
			t.Errorf("got: %v, want: %v", matrix2, tt.want)
		}
	}
}

func shallowCopy2DSlice(matrix [][]int) [][]int {
	shallow := make([][]int, len(matrix))
	for i := range matrix {
		arr := make([]int, len(matrix[i]))
		copy(arr, matrix[i])
		shallow[i] = arr
	}
	return shallow
}
