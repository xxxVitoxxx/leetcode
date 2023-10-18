package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{
			name:  "example1",
			input: []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:   "example2",
			input:  []int{0, 1, 1},
			output: [][]int{},
		},
		{
			name:   "example3",
			input:  []int{0, 0, 0},
			output: [][]int{{0, 0, 0}},
		},
		{
			name:   "example4",
			input:  []int{1, 1, -2},
			output: [][]int{{-2, 1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum(tt.input)
			for i := range res {
				exist := Exist(tt.output, res[i])
				assert.True(t, exist)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum2(tt.input)
			for i := range res {
				exist := Exist(tt.output, res[i])
				assert.True(t, exist)
			}
		})
	}
}
