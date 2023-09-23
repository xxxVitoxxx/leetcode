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
		{"example1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.input, tt.target)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}
