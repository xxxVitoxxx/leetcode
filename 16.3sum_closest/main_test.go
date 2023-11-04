package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		output int
	}{
		{"example1", []int{-1, 2, 1, -4}, 1, 2},
		{"example2", []int{0, 0, 0}, 1, 0},
		{"example3", []int{0, 1, 2}, 3, 3},
		{"example4", []int{-1, -4, 5, 2, -10, 19}, 10, 11},
		{"example5", []int{1, 40, 51, 30, 71, -12, -200, -21, -51, 150}, -98, -99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSumClosest(tt.nums, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSumClosest2(tt.nums, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}
}
