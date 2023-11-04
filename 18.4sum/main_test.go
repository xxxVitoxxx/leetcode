package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		output [][]int
	}{
		{"example1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"example2", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{"example3", []int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
		{"example4", []int{-10, 1, 9, -6, 51, 31}, -1000, [][]int{}},
		{"example5", []int{1000000, -4, 90, 799915, -10000, -532, -9}, 1800001, [][]int{{-4, 90, 799915, 1000000}}},
		{"example6", []int{1, 2}, 10, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fourSum(tt.nums, tt.target)
			assert.Len(t, res, len(tt.output))
			assert.ElementsMatch(t, res, tt.output)
		})
	}
}
