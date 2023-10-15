package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		output []int
	}{
		{"example1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"example2", []int{2, 3, 4}, 6, []int{1, 3}},
		{"example3", []int{-1, 0}, -1, []int{1, 2}},
		{"example4", []int{5, 25, 75}, 100, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			res := twoSum(tt.input, tt.target)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			res := twoSum2(tt.input, tt.target)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}
