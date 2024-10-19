package main

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	for _, tt := range []struct {
		matrix [][]int
		want   []int
	}{{
		matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	}, {
		matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	}, {
		matrix: [][]int{{8, 9, 4, 1, 2, 5}, {10, 19, 32, 14, 3, 6}},
		want:   []int{8, 9, 4, 1, 2, 5, 6, 3, 14, 32, 19, 10},
	}, {
		matrix: [][]int{{-10, 0, -7}, {0, 9, -100}},
		want:   []int{-10, 0, -7, -100, 9, 0},
	}} {
		matrix1 := shallowCopy2DSlice(tt.matrix)
		got := spiralOrder(matrix1)
		if !equal(tt.want, got) {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}

		matrix2 := shallowCopy2DSlice(tt.matrix)
		got = spiralOrder2(matrix2)
		if !equal(tt.want, got) {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}

		matrix3 := shallowCopy2DSlice(tt.matrix)
		got = spiralOrder3(matrix3)
		if !equal(tt.want, got) {
			t.Errorf("got: %v, want: %v", got, tt.want)
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
